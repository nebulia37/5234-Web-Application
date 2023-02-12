package routes

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/google/uuid"

	"cse5234/payment/pkg/domain"
)

const (
	ParamPaymentID   = "paymentID"
	ParamPaymentName = "name"
)

// OrderRouter serves all the routes related to orders
func (handler Handler) PaymentRouter() chi.Router {
	r := chi.NewRouter()

	//r.Get("/", handler.GetPayments) // GET /orders
	r.Post("/", handler.AddPayment) // POST /orders

	//r.Route("/{orderID}", func(r chi.Router) {
	//r.Use(handler.PaymentCtx)            // Load the *Order on the request context
	//r.Get("/", handler.GetPayment)       // GET /orders/123
	//r.Put("/", handler.UpdateOrder)    // PUT /orders/123
	//r.Delete("/", handler.DeleteOrder) // DELETE /orders/123
	//})

	return r
}

type PaymentRequest struct {
	*domain.Payment
}

// Bind preprocesses the request for some basic error checking
func (req *PaymentRequest) Bind(r *http.Request) error {
	// Return an error to avoid a nil pointer dereference.
	if req.Payment == nil {
		return errors.New("missing required fields")
	}

	return nil
}

type PaymentResponse struct {
	*domain.Payment
}

// Render preprocess the response before it's sent to the wire
func (rsp *PaymentResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// do nothing
	return nil
}

func NewPaymentResponse(payment *domain.Payment) *PaymentResponse {
	resp := PaymentResponse{
		Payment: payment,
	}

	return &resp
}

func NewPaymentListResponse(payments []*domain.Payment) []render.Renderer {
	list := []render.Renderer{}

	for _, order := range payments {
		list = append(list, NewPaymentResponse(order))
	}

	return list
}

// OrderCtx middleware is used to load an Order object from
// the URL parameters passed through as the request. In case
// the Order could not be found, we stop here and return a 404.
// func (handler Handler) PaymentCtx(next http.Handler) http.Handler {
// 	const (
// 		Base = 10
// 		Bit  = 64
// 	)

// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		var payment *domain.Payment
// 		var paymentID int64
// 		var err error

// 		// find the orderID from URL params
// 		if param := chi.URLParam(r, ParamPaymentID); len(param) > 0 {
// 			paymentID, err = strconv.ParseInt(param, Base, Bit)
// 			if err != nil {
// 				render.Render(w, r, ErrInvalidRequest(err))
// 				return
// 			}
// 			payment, err = handler.Service.GetOrderByID(r.Context(), paymentID)
// 		} else {
// 			render.Render(w, r, ErrNotFound())
// 			return
// 		}

// 		if err != nil {
// 			render.Render(w, r, ErrNotFound())
// 			return
// 		}

// 		ctx := context.WithValue(r.Context(), PaymentKey, payment)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

// func (handler Handler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
// 	order := r.Context().Value(paymentKey).(*domain.Payment)

// 	if err := handler.Service.DeleteOrderByID(r.Context(), order.ID); err != nil {
// 		render.Render(w, r, ErrInvalidRequest(err))
// 		return
// 	}

// 	render.Status(r, http.StatusOK)
// }

// func (handler Handler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
// 	order := r.Context().Value(paymentKey).(*domain.Payment)

// 	req := PaymentRequest{}

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

// func (handler Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
// 	order := r.Context().Value(paymentKey).(*domain.Payment)

// 	if err := render.Render(w, r, NewPaymentResponse(order)); err != nil {
// 		render.Render(w, r, ErrRender(err))
// 		return
// 	}
// }

// func (handler Handler) GetOrders(w http.ResponseWriter, r *http.Request) {
// 	// query the database for list of orders
// 	orders, err := handler.Service.GetOrders(r.Context())
// 	if err != nil {
// 		render.Render(w, r, ErrRender(err))
// 		return
// 	}

// 	// render response
// 	if err := render.RenderList(w, r, NewOrderListResponse(orders)); err != nil {
// 		render.Render(w, r, ErrRender(err))
// 		return
// 	}
// }

func (handler Handler) AddPayment(w http.ResponseWriter, r *http.Request) {
	req := PaymentRequest{&domain.Payment{}}

	// unmarshal new order from request
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	req.Confirmation = uuid.New().String()
	log.Printf("Confirmed Payment: %s", req.Confirmation)

	// // add the new order to database
	// payment := req.Payment
	// if err := handler.Service.AddOrder(r.Context(), order); err != nil {
	// 	render.Render(w, r, ErrInvalidRequest(err))
	// 	return
	// }

	// render response
	render.Status(r, http.StatusOK)
	render.Render(w, r, NewPaymentResponse(req.Payment))
}
