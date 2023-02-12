package routes

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// The key type is unexported to prevent collisions with context keys defined in
// other packages.
type key int

// context keys
const (
	itemKey key = iota
	pageIDKey
)

// ApiRouter handles RESTful API requests
func (handler Handler) ApiRouter() chi.Router {
	r := chi.NewRouter()

	r.Use(render.SetContentType(render.ContentTypeJSON))

	// public routes
	r.Group(func(r chi.Router) {
		r.Mount("/items", handler.ItemRouter())
	})

	return r
}
