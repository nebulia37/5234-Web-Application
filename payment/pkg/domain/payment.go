package domain

import (
	"context"
)

type Payment struct {
	BusinessEntityName    string `json:"be_name"`
	BusinessEntityAccount string `json:"be_account"`
	CreditCardNumber      string `json:"cc_number"`
	CreditCardName        string `json:"cc_name"`
	CreditCardExpires     string `json:"cc_expires"`
	CreditCVV             string `json:"cc_cvv"`
	Confirmation          string `json:"confirmation,omitempty"`
}

type PaymentService interface {
	AddPayment(ctx context.Context, payment *Payment) error
	// GetOrders(ctx context.Context) ([]*Order, error)
	// GetOrderByID(ctx context.Context, id int64) (*Order, error)
	// UpdateOrderByID(ctx context.Context, id int64, newOrder *Order) error
	// DeleteOrderByID(ctx context.Context, id int64) error
}
