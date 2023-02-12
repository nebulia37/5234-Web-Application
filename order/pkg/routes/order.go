package routes

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	amqp "github.com/rabbitmq/amqp091-go"

	"cse5234/order/pkg/domain"
)

const (
	ParamOrderID   = "orderID"
	ParamOrderName = "name"
)

// OrderRouter serves all the routes related to orders
func (handler Handler) OrderRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", handler.GetOrders) // GET /orders
	r.Post("/", handler.AddOrder) // POST /orders

	r.Route("/{orderID}", func(r chi.Router) {
		r.Use(handler.OrderCtx)      // Load the *Order on the request context
		r.Get("/", handler.GetOrder) // GET /orders/123
		// r.Put("/", handler.UpdateOrder)    // PUT /orders/123
		r.Delete("/", handler.DeleteOrder) // DELETE /orders/123
	})

	return r
}

type OrderRequest struct {
	*domain.Order
}

// Bind preprocesses the request for some basic error checking
func (req *OrderRequest) Bind(r *http.Request) error {
	// Return an error to avoid a nil pointer dereference.
	if req.Order == nil {
		return errors.New("missing required fields")
	}

	return nil
}

type OrderResponse struct {
	*domain.Order
}

// Render preprocess the response before it's sent to the wire
func (rsp *OrderResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// do nothing
	return nil
}

func NewOrderResponse(order *domain.Order) *OrderResponse {
	resp := OrderResponse{
		Order: order,
	}

	return &resp
}

func NewOrderListResponse(orders []*domain.Order) []render.Renderer {
	list := []render.Renderer{}

	for _, order := range orders {
		list = append(list, NewOrderResponse(order))
	}

	return list
}

// OrderCtx middleware is used to load an Order object from
// the URL parameters passed through as the request. In case
// the Order could not be found, we stop here and return a 404.
func (handler Handler) OrderCtx(next http.Handler) http.Handler {
	const (
		Base = 10
		Bit  = 64
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var order *domain.Order
		var orderID int64
		var err error

		// find the orderID from URL params
		if param := chi.URLParam(r, ParamOrderID); len(param) > 0 {
			orderID, err = strconv.ParseInt(param, Base, Bit)
			if err != nil {
				render.Render(w, r, ErrInvalidRequest(err))
				return
			}
			order, err = handler.Service.GetOrderByID(r.Context(), orderID)
		} else {
			render.Render(w, r, ErrNotFound())
			return
		}

		if err != nil {
			render.Render(w, r, ErrNotFound())
			return
		}

		ctx := context.WithValue(r.Context(), orderKey, order)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (handler Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	order := r.Context().Value(orderKey).(*domain.Order)

	if err := handler.Service.DeleteOrderByID(r.Context(), order.ID); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusOK)
}

// func (handler Handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
// 	order := r.Context().Value(orderKey).(*domain.Order)

// 	req := OrderRequest{}

// 	// unmarshal new order from request
// 	if err := render.Bind(r, &req); err != nil {
// 		render.Render(w, r, ErrInvalidRequest(err))
// 		return
// 	}

// 	// add the new order to database
// 	newOrder := req.Order
// 	if err := handler.Service.UpdateOrderByID(r.Context(), order.ID, newOrder); err != nil {
// 		render.Render(w, r, ErrInvalidRequest(err))
// 		return
// 	}

// 	render.Status(r, http.StatusOK)
// }

func (handler Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	order := r.Context().Value(orderKey).(*domain.Order)

	if err := render.Render(w, r, NewOrderResponse(order)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (handler Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
	// query the database for list of orders
	orders, err := handler.Service.GetOrders(r.Context())
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	// render response
	if err := render.RenderList(w, r, NewOrderListResponse(orders)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (handler Handler) AddOrder(w http.ResponseWriter, r *http.Request) {
	req := OrderRequest{&domain.Order{}}

	// unmarshal new order from request
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	// add the new order to database
	order := req.Order
	if err := handler.Service.AddOrder(r.Context(), order); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	// encode order details into json
	shipmentReq, _ := json.Marshal(order)

	// initiate shipment
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := handler.ShipmentChannel.PublishWithContext(
		ctx,
		"",
		handler.RoutingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        shipmentReq,
		},
	)
	if err != nil {
		log.Printf("RabbitMQ Error: %v", err)
	}

	// render response
	render.Status(r, http.StatusOK)
	render.Render(w, r, NewOrderResponse(order))
}
