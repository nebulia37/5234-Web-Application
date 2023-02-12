package mysql

import (
	"context"
	"fmt"
)

func (service Service) UpdateOrderStatus(ctx context.Context) error {
	// start a transaction
	tx, err := service.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("UpdateOrderStatus: %v", err)
	}
	defer tx.Rollback()

	// execute the query
	result, err := tx.ExecContext(ctx,
		`UPDATE orders 
		SET order_status='CONFIRMED'
		WHERE order_status='NEW'`)

	// check if the query failed
	if err != nil {
		return fmt.Errorf("UpdateOrderStatus: %v", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("UpdateOrderStatus: %v", err)
	}

	// commit the transaction
	if err = tx.Commit(); err != nil {
		return fmt.Errorf("UpdateOrderStatus: %v", err)
	}

	return nil
}
