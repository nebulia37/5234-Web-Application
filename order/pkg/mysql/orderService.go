package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"cse5234/order/pkg/domain"
)

func (service Service) DeleteOrderByID(ctx context.Context, id int64) error {
	// start a transaction
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("DeleteOrderByID [%d]: %v", id, err)
	}
	defer tx.Rollback()

	// execute the query
	result, err := tx.ExecContext(ctx,
		`DELETE FROM orders 
		WHERE id=?`,
		id)

	// check if the query failed
	if err != nil {
		return fmt.Errorf("DeleteOrderByID [%d]: %v", id, err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteOrderByID [%d]: %v", id, err)
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("DeleteOrderByID [%d]: %v", id, err)
	}

	return nil
}

func (service Service) UpdatePaymentConfirmationByID(ctx context.Context, id int64, confirmation uint64) error {

	// start a transaction
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("UpdatePaymentConfirmationByID [%d]: %v", id, err)
	}
	defer tx.Rollback()

	// execute the query
	result, err := tx.ExecContext(ctx,
		`UPDATE orders 
		SET payment_confirmation=? 
		WHERE id=?`,
		confirmation,
		id)

	// check if the query failed
	if err != nil {
		return fmt.Errorf("UpdatePaymentConfirmationByID [%d]: %v", id, err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("UpdatePaymentConfirmationByID [%d]: %v", id, err)
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("UpdatePaymentConfirmationByID [%d]: %v", id, err)
	}

	return nil
}

func (service Service) UpdateShipmentLabelByID(ctx context.Context, id int64, label string) error {

	// start a transaction
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("UpdateShipmentLabelByID [%d]: %v", id, err)
	}
	defer tx.Rollback()

	// execute the query
	result, err := tx.ExecContext(ctx,
		`UPDATE orders 
		SET shipment_label=? 
		WHERE id=?`,
		label,
		id)

	// check if the query failed
	if err != nil {
		return fmt.Errorf("UpdateShipmentLabelByID [%d]: %v", id, err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("UpdateShipmentLabelByID [%d]: %v", id, err)
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("UpdateShipmentLabelByID [%d]: %v", id, err)
	}

	return nil
}

func (service Service) GetOrderByID(ctx context.Context, id int64) (*domain.Order, error) {
	order := domain.Order{}

	// start a transaction
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("GetOrderByID[%d]: %v", id, err)
	}
	defer tx.Rollback()

	// execute the query
	row := tx.QueryRowContext(ctx,
		`SELECT orders.id, orders.email, orders.customer_name, orders.address_line1, orders.address_line2, orders.address_state, orders.address_zip, orders.payment_ccnumber, orders.payment_ccname, orders.payment_ccexpires, orders.created_at, orders.order_status, orders.item_id, orders.item_count
		FROM orders WHERE id=?`, id)

	// parse response
	if err := row.Scan(&order.ID, &order.Email, &order.CustomerName, &order.AddrLine1, &order.AddrLine2, &order.AddrState, &order.AddrZipcode, &order.CreditCardNumber, &order.CreditCardName, &order.CreditCardExpires, &order.CreatedAt, &order.Status, &order.ItemID, &order.ItemCount); err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("GetOrderByID[%d]: no such order", id)
		}

		return nil, fmt.Errorf("GetOrderByID[%d]: %v", id, err)
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("GetOrderByID[%d]: %v", id, err)
	}

	return &order, nil
}

// GetOrders queries the database for all the orders
func (service Service) GetOrders(ctx context.Context) ([]*domain.Order, error) {
	// order slice to hold the data from database query
	orders := []*domain.Order{}

	// start a transaction
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("GetOrders: %v", err)
	}
	defer tx.Rollback()

	// execute the query
	rows, err := tx.QueryContext(ctx,
		`SELECT orders.id, orders.email, orders.customer_name, orders.address_line1, orders.address_line2, orders.address_state, orders.address_zip, orders.payment_ccnumber, orders.payment_ccname, orders.payment_ccexpires, orders.created_at, orders.order_status, orders.item_id, orders.item_count
		FROM orders ORDER BY id`)

	// check if the query failed
	if err != nil {
		return nil, fmt.Errorf("GetOrders: %v", err)
	}

	defer rows.Close()

	// parse response
	for rows.Next() {
		order := domain.Order{}
		if err := rows.Scan(&order.ID, &order.Email, &order.CustomerName, &order.AddrLine1, &order.AddrLine2, &order.AddrState, &order.AddrZipcode, &order.CreditCardNumber, &order.CreditCardName, &order.CreditCardExpires, &order.CreatedAt, &order.Status, &order.ItemID, &order.ItemCount); err != nil {
			return nil, fmt.Errorf("GetOrders: %v", err)
		}

		// add order to the return slice
		orders = append(orders, &order)
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return nil, fmt.Errorf("GetOrders: %v", err)
	}

	return orders, nil
}

// AddOrder adds 1 order to the database,
// updating the order's ID upon success
func (service Service) AddOrder(ctx context.Context, order *domain.Order) error {
	if order == nil {
		return fmt.Errorf("AddOrder : null pointer error")
	}

	// start a transaction
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("AddOrder [%s]: %v", order.Email, err)
	}
	defer tx.Rollback()

	// execute the query
	result, err := tx.ExecContext(ctx,
		`INSERT INTO orders (email, customer_name, address_line1, address_line2, address_state, address_zip, payment_ccnumber, payment_ccname, payment_ccexpires, payment_confirm, item_id, item_count) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		order.Email,
		order.CustomerName,
		order.AddrLine1,
		order.AddrLine2,
		order.AddrState,
		order.AddrZipcode,
		order.CreditCardNumber,
		order.CreditCardName,
		order.CreditCardExpires,
		order.PaymentConfirmation,
		order.ItemID,
		order.ItemCount,
	)

	// check if the query failed
	if err != nil {
		return fmt.Errorf("AddOrder [%s]: %v", order.Email, err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("AddOrder [%s]: %v", order.Email, err)
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("AddOrder [%s]: %v", order.Email, err)
	}

	order.ID = id

	return nil
}
