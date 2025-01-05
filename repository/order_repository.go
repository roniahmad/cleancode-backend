package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"wetees.com/domain"
	"wetees.com/internal/vars"
)

const (
	TABLE_ORDERS      = "orders"
	TABLE_ORDER_ITEMS = "order_items"
)

type orderRepository struct {
	Conn *sql.DB
}

// Get order items
func (uc *orderRepository) GetOrderItems(c context.Context, orderId int) ([]domain.OrderDetails, error) {
	result := []domain.OrderDetails{}

	query := fmt.Sprintf(`
		SELECT A.order_id, A.product_id, B.name, B.price, A.quantity
		FROM %s A
		LEFT JOIN products B On (A.product_id=B.id)
		WHERE  A.order_id=? `, TABLE_ORDER_ITEMS)

	rows, err := uc.Conn.QueryContext(c, query, orderId)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	for rows.Next() {
		item := domain.OrderDetails{}
		if err = rows.Scan(&item.OrderId, &item.ProductId, &item.ProductName, &item.Price, &item.Quantity); err == nil {
			result = append(result, item)
		}
	}

	return result, nil
}

// Check if chart is empty
func (uc *orderRepository) IsChartEmpty(c context.Context, orderId int) (bool, error) {
	// Calculate all rows count
	var count int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE order_id=?", TABLE_ORDER_ITEMS)
	if err := uc.Conn.QueryRowContext(c, query, orderId).Scan(&count); err != nil {
		return false, err
	}

	return count <= 0, nil
}

// Checkout Order
func (uc *orderRepository) CheckoutOrder(c context.Context, orderId int, userId int) ([]domain.OrderDetails, error) {
	orderItems := []domain.OrderDetails{}

	if empty, _ := uc.IsChartEmpty(c, orderId); empty {
		return orderItems, vars.ErrChartIsEmpty
	}

	items, _ := uc.GetOrderItems(c, orderId)

	query := fmt.Sprintf("UPDATE %s set status=? WHERE id=? AND user_id=?", TABLE_ORDERS)
	stmt, err := uc.Conn.PrepareContext(c, query)
	if err != nil {
		return orderItems, err
	}

	// Update status to AWP (Awaiting Payment)
	result, err := stmt.ExecContext(c, "AWP", orderId, userId)
	if err != nil {
		return orderItems, err
	}

	_, err = result.RowsAffected()

	return items, err
}

// Create order
func (uc *orderRepository) CreateOrder(c context.Context, user *domain.Order) (id int, err error) {
	var (
		stmt         *sql.Stmt
		result       sql.Result
		lastInsertId int64
	)

	query := fmt.Sprintf(`INSERT INTO %s (user_id, status, created_at)	
		VALUES(?, ?, ?)`, TABLE_ORDERS)

	if stmt, err = uc.Conn.PrepareContext(c, query); err != nil {
		return 0, err
	}

	if result, err = stmt.ExecContext(c, &user.UserId, &user.Status, time.Now().Local().UTC()); err != nil {
		return 0, err
	}

	lastInsertId, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertId), nil

}

// DeleteOrder implements domain.OrderRepository.
func (ur *orderRepository) DeleteOrder(c context.Context, orderId int, userId int) (err error) {
	var (
		stmt   *sql.Stmt
		result sql.Result
	)

	query := fmt.Sprintf("UPDATE %s set status=? WHERE id=? AND user_id=?", TABLE_ORDERS)
	if stmt, err = ur.Conn.PrepareContext(c, query); err != nil {
		return
	}
	// Update status to CNC (Cancelled)
	if result, err = stmt.ExecContext(c, "CNC", orderId, userId); err != nil {
		return
	}

	_, err = result.RowsAffected()

	return
}

func NewOrderRepository(db *sql.DB) domain.OrderRepository {
	return &orderRepository{
		Conn: db,
	}
}
