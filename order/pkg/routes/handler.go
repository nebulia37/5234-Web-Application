package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	amqp "github.com/rabbitmq/amqp091-go"

	"cse5234/order/pkg/domain"
)

// Handler handles all API requests to camera roll
type Handler struct {
	Service         domain.Service
	ShipmentChannel *amqp.Channel
	RoutingKey      string
}

// NewHandler is the contructor method for the Handler
func NewHandler(service domain.Service, channel *amqp.Channel, routingKey string) *Handler {
	handler := Handler{
		Service:         service,
		ShipmentChannel: channel,
		RoutingKey:      routingKey,
	}

	return &handler
}

// Routes is the collection of all routes being served
func (handler Handler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"http://*", "https://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Mount("/", handler.ApiRouter())

	return r
}
