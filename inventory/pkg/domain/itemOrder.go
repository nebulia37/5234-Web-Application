package domain

import "context"

type ItemOrder struct {
	ID             int64 `json:"id,omitempty"`
	ItemID         int64 `json:"itemID"`
	OrderID        int64 `json:"orderID"`
	OrderQuantity  int64 `json:"orderQuantity"`
	OrderFulfilled bool  `json:"orderFulfilled"`
}

type ItemOrderService interface {
	PlaceOrderOnItem(ctx context.Context, itemID int64, orderID int64) error
	GetItemsFromOrder(ctx context.Context, orderID int64) ([]*Item, error)
	RemoveOrderFromItem(ctx context.Context, itemID int64, orderID int64) error
}
