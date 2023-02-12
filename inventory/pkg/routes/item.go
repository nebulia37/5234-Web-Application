package routes

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"cse5234/inventory/pkg/domain"
)

const (
	ParamItemID   = "itemID"
	ParamItemName = "name"
)

// ItemRouter serves all the routes related to items
func (handler Handler) ItemRouter() chi.Router {
	r := chi.NewRouter()

	r.Get("/", handler.GetItems) // GET /items
	r.Post("/", handler.AddItem) // POST /items

	r.Route("/{itemID}", func(r chi.Router) {
		r.Use(handler.ItemCtx)            // Load the *Item on the request context
		r.Get("/", handler.GetItem)       // GET /items/123
		r.Put("/", handler.UpdateItem)    // PUT /images/123
		r.Delete("/", handler.DeleteItem) // DELETE /images/123
	})

	return r
}

type ItemRequest struct {
	*domain.Item
}

// Bind preprocesses the request for some basic error checking
func (req *ItemRequest) Bind(r *http.Request) error {
	// Return an error to avoid a nil pointer dereference.
	if req.Item == nil {
		return errors.New("missing required fields")
	}

	return nil
}

type ItemResponse struct {
	*domain.Item
}

// Render preprocess the response before it's sent to the wire
func (rsp *ItemResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// do nothing
	return nil
}

func NewItemResponse(item *domain.Item) *ItemResponse {
	resp := ItemResponse{
		Item: item,
	}

	return &resp
}

func NewItemListResponse(items []*domain.Item) []render.Renderer {
	list := []render.Renderer{}

	for _, item := range items {
		list = append(list, NewItemResponse(item))
	}

	return list
}

// ItemCtx middleware is used to load an Item object from
// the URL parameters passed through as the request. In case
// the Item could not be found, we stop here and return a 404.
func (handler Handler) ItemCtx(next http.Handler) http.Handler {
	const (
		Base = 10
		Bit  = 64
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var item *domain.Item
		var itemID int64
		var err error

		// find the itemID from URL params
		if param := chi.URLParam(r, ParamItemID); len(param) > 0 {
			itemID, err = strconv.ParseInt(param, Base, Bit)
			if err != nil {
				render.Render(w, r, ErrInvalidRequest(err))
				return
			}
			item, err = handler.Service.GetItemByID(r.Context(), itemID)
		} else {
			render.Render(w, r, ErrNotFound())
			return
		}

		if err != nil {
			render.Render(w, r, ErrNotFound())
			return
		}

		ctx := context.WithValue(r.Context(), itemKey, item)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (handler Handler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	item := r.Context().Value(itemKey).(*domain.Item)

	if err := handler.Service.DeleteItemByID(r.Context(), item.ID); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusOK)
}

func (handler Handler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	item := r.Context().Value(itemKey).(*domain.Item)

	req := ItemRequest{}

	// unmarshal new item from request
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	// add the new item to database
	newItem := req.Item
	if err := handler.Service.UpdateItemByID(r.Context(), item.ID, newItem); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	render.Status(r, http.StatusOK)
}

func (handler Handler) GetItem(w http.ResponseWriter, r *http.Request) {
	item := r.Context().Value(itemKey).(*domain.Item)

	if err := render.Render(w, r, NewItemResponse(item)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (handler Handler) GetItems(w http.ResponseWriter, r *http.Request) {
	// query the database for list of items
	items, err := handler.Service.GetItems(r.Context())
	if err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

	// render response
	if err := render.RenderList(w, r, NewItemListResponse(items)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func (handler Handler) AddItem(w http.ResponseWriter, r *http.Request) {
	req := ItemRequest{&domain.Item{}}

	// unmarshal new item from request
	if err := render.Bind(r, &req); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	// add the new item to database
	item := req.Item
	if err := handler.Service.AddItem(r.Context(), item); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		return
	}

	// render response
	render.Status(r, http.StatusOK)
	render.Render(w, r, NewItemResponse(item))
}
