package domain

import (
	"context"
	"time"
)

type Order struct {
	ID                  int64     `json:"id"`
	Email               string    `json:"email"`
	CustomerName        string    `json:"customer_name"`
	AddrLine1           string    `json:"addr_line1"`
	AddrLine2           string    `json:"addr_line2"`
	AddrState           string    `json:"addr_state"`
	AddrZipcode         string    `json:"addr_zipcode"`
	CreditCardNumber    string    `json:"cc_number"`
	CreditCardName      string    `json:"cc_name"`
	CreditCardExpires   string    `json:"cc_expires"`
	CreatedAt           time.Time `json:"created_at,omitempty"`
	PaymentConfirmation string    `json:"payment_confirmation,omitempty"`
	ShipmentLabel       string    `json:"shipment_label,omitempty"`
	Status              string    `json:"status"`
	ItemID              int64     `json:"item_id,omitempty"`
	ItemCount           int64     `json:"item_count,omitempty"`
}

type OrderService interface {
	AddOrder(ctx context.Context, order *Order) error
	GetOrders(ctx context.Context) ([]*Order, error)
	GetOrderByID(ctx context.Context, id int64) (*Order, error)
	UpdatePaymentConfirmationByID(ctx context.Context, id int64, paymentConfirmation uint64) error
	UpdateShipmentLabelByID(ctx context.Context, id int64, shipmentLabel string) error
	DeleteOrderByID(ctx context.Context, id int64) error
}
