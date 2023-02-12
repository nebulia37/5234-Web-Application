package mysql

import (
	"context"
	"fmt"

	"cse5234/inventory/pkg/domain"
)

func (service Service) RemoveOrderFromItem(ctx context.Context, itemID int64, orderID int64) error {
	// start a transaction
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("RemoveOrderFromItem itemID[%d] orderID[%d]: %v", itemID, orderID, err)
	}
	defer tx.Rollback()

	// execute the query
	result, err := tx.ExecContext(ctx,
		`DELETE FROM item_orders
		WHERE item_id=? AND order_id=?`,
		itemID,
		orderID)

	// check if the query failed
	if err != nil {
		return fmt.Errorf("RemoveOrderFromItem itemID[%d] orderID[%d]: %v", itemID, orderID, err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("RemoveOrderFromItem itemID[%d] orderID[%d]: %v", itemID, orderID, err)
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("RemoveOrderFromItem itemID[%d] orderID[%d]: %v", itemID, orderID, err)
	}

	return nil
}

func (service Service) GetItemsFromOrder(ctx context.Context, orderID int64) ([]*domain.Item, error) {
	// the slice to hold the data from database query
	items := []*domain.Item{}

	// start a transaction
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("GetItemsFromOrder: %v", err)
	}
	defer tx.Rollback()

	// execute the query
	rows, err := tx.QueryContext(ctx,
		`SELECT items.id, items.img, items.title, items.price, items.quantity 
		FROM items_order,items WHERE item_id=id`)

	// check if the query failed
	if err != nil {
		return nil, fmt.Errorf("GetItemsFromOrder: %v", err)
	}

	defer rows.Close()

	// parse response
	for rows.Next() {
		item := domain.Item{}
		if err := rows.Scan(
			&item.ID,
			&item.Image,
			&item.Title,
			&item.Price,
			&item.Quantity); err != nil {
			return nil, fmt.Errorf("GetItemsFromOrder: %v", err)
		}

		// append to the return slice
		items = append(items, &item)
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("GetItemsFromOrder: %v", err)
	}

	return items, nil
}

func (service Service) PlaceOrderOnItem(ctx context.Context, itemID int64, orderID int64) error {
	// TODO: implement this

	return nil
}
