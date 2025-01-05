package domain

import (
	"context"
	"database/sql"
)

type (
	Order struct {
		ID        int          `json:"id"`
		UserId    int          `json:"user_id"`
		Status    string       `json:"status"`
		CreatedAt sql.NullTime `json:"created_at"`
	}

	OrderDetails struct {
		OrderId     int     `json:"order_id"`
		ProductId   int     `json:"product_id"`
		ProductName string  `json:"product_name"`
		Price       float64 `json:"price"`
		Quantity    int     `json:"quantity"`
	}
)

type OrderRepository interface {
	IsChartEmpty(c context.Context, orderId int) (bool, error)
	GetOrderItems(c context.Context, orderId int) ([]OrderDetails, error)
	CreateOrder(c context.Context, order *Order) (id int, err error)
	DeleteOrder(c context.Context, orderId, userId int) error
	CheckoutOrder(c context.Context, orderId, userId int) ([]OrderDetails, error)
}

type OrderUsecase interface {
	CreateOrder(c context.Context, userId int) (map[string]interface{}, error)
	CancelOrder(c context.Context, orderId, userId int) error
	CheckoutOrder(c context.Context, orderId, userId int) ([]OrderDetails, int, float64, error)
}
