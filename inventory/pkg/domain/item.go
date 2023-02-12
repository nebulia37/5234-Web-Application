package domain

import (
	"context"
	"time"
)

type Item struct {
	ID        int64     `json:"id"`
	Image     string    `json:"img"`
	Title     string    `json:"title"`
	Price     string    `json:"price"`
	Quantity  int64     `json:"quantity"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type ItemService interface {
	AddItem(ctx context.Context, item *Item) error
	GetItems(ctx context.Context) ([]*Item, error)
	GetItemByID(ctx context.Context, id int64) (*Item, error)
	UpdateItemByID(ctx context.Context, id int64, newItem *Item) error
	DeleteItemByID(ctx context.Context, id int64) error
}
