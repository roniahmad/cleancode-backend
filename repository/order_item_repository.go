package repository

import (
	"context"
	"database/sql"
	"fmt"

	"wetees.com/domain"
)

type orderItemRepository struct {
	Conn *sql.DB
}

// Add Item to chart
func (o *orderItemRepository) AddItem(c context.Context, item *domain.OrderItem) (err error) {
	var (
		stmt *sql.Stmt
	)

	query := fmt.Sprintf(`INSERT INTO %s (order_id, product_id, quantity) VALUES(?, ?, ?)`, TABLE_ORDER_ITEMS)
	if stmt, err = o.Conn.PrepareContext(c, query); err != nil {
		return err
	}

	if _, err = stmt.ExecContext(c, &item.OrderId, &item.ProductId, &item.Quantity); err != nil {
		return err
	}

	return nil
}

// Delete Order Item
func (o *orderItemRepository) DeleteItem(c context.Context, orderId, productId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE order_id= ? AND product_id=?", TABLE_ORDER_ITEMS)

	stmt, err := o.Conn.PrepareContext(c, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(c, orderId, productId)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

// Modify Order Item, update quatity
func (ur *orderItemRepository) ModifyItem(c context.Context, orderId, productId, quantity int) (err error) {
	var (
		stmt   *sql.Stmt
		result sql.Result
	)

	query := fmt.Sprintf("UPDATE %s set quantity=? WHERE order_id= ? AND product_id=?", TABLE_ORDER_ITEMS)
	if stmt, err = ur.Conn.PrepareContext(c, query); err != nil {
		return
	}

	if result, err = stmt.ExecContext(c, quantity, orderId, productId); err != nil {
		return
	}

	_, err = result.RowsAffected()

	return
}

func NewOrderItemRepository(db *sql.DB) domain.OrderItemRepository {
	return &orderItemRepository{
		Conn: db,
	}
}
