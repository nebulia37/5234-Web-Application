package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"cse5234/inventory/pkg/domain"
)

func (service Service) DeleteItemByID(ctx context.Context, id int64) error {
	// start a transaction
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("DeleteItemByID [%d]: %v", id, err)
	}
	defer tx.Rollback()

	// execute the query
	result, err := tx.ExecContext(ctx,
		`DELETE FROM items 
		WHERE id=?`,
		id)

	// check if the query failed
	if err != nil {
		return fmt.Errorf("DeleteItemByID [%d]: %v", id, err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteItemByID [%d]: %v", id, err)
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("DeleteItemByID [%d]: %v", id, err)
	}

	return nil
}

func (service Service) UpdateItemByID(ctx context.Context, id int64, newItem *domain.Item) error {
	if newItem == nil {
		return fmt.Errorf("UpdateItemByID [%d]: null pointer error", id)
	}

	// start a transaction
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("UpdateItemByID [%d]: %v", id, err)
	}
	defer tx.Rollback()

	// execute the query
	result, err := tx.ExecContext(ctx,
		`UPDATE items 
		SET img=?, title=?, price=?, quantity=?
		WHERE id=?`,
		newItem.Image,
		newItem.Title,
		newItem.Price,
		newItem.Quantity,
		id)

	// check if the query failed
	if err != nil {
		return fmt.Errorf("UpdateItemByID [%d]: %v", id, err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("UpdateItemByID [%d]: %v", id, err)
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("UpdateItemByID [%d]: %v", id, err)
	}

	return nil
}

func (service Service) GetItemByID(ctx context.Context, id int64) (*domain.Item, error) {
	item := domain.Item{}

	// start a transaction
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("GetItemByID[%d]: %v", id, err)
	}
	defer tx.Rollback()

	// execute the query
	row := tx.QueryRowContext(ctx,
		`SELECT items.id, items.img, items.title, items.price, items.quantity 
		FROM items WHERE id=?`,
		id)

	// parse response
	if err := row.Scan(
		&item.ID,
		&item.Image,
		&item.Title,
		&item.Price,
		&item.Quantity,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("GetItemByID[%d]: no such item", id)
		}

		return nil, fmt.Errorf("GetItemByID[%d]: %v", id, err)
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("GetItemByID[%d]: %v", id, err)
	}

	return &item, nil
}

func (service Service) GetItems(ctx context.Context) ([]*domain.Item, error) {
	// the slice to hold the data from database query
	items := []*domain.Item{}

	// start a transaction
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("GetItems: %v", err)
	}
	defer tx.Rollback()

	// execute the query
	rows, err := tx.QueryContext(ctx,
		`SELECT items.id, items.img, items.title, items.price, items.quantity 
		FROM items ORDER BY id DESC`)

	// check if the query failed
	if err != nil {
		return nil, fmt.Errorf("GetItems: %v", err)
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
			return nil, fmt.Errorf("GetItems: %v", err)
		}

		// append to the return slice
		items = append(items, &item)
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("GetItems: %v", err)
	}

	return items, nil
}

func (service Service) AddItem(ctx context.Context, item *domain.Item) error {
	if item == nil {
		return fmt.Errorf("AddItem : null pointer error")
	}

	// start a transaction
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("AddItem [%s]: %v", item.Title, err)
	}
	defer tx.Rollback()

	// execute the query
	result, err := tx.ExecContext(ctx,
		`INSERT INTO items (img, title, price, quantity) 
		VALUES (?, ?, ?, ?)`,
		item.Image,
		item.Title,
		item.Price,
		item.Quantity,
	)

	// check if the query failed
	if err != nil {
		return fmt.Errorf("AddItem [%s]: %v", item.Title, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("AddItem [%s]: %v", item.Title, err)
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("AddItem [%s]: %v", item.Title, err)
	}

	item.ID = id

	return nil
}
