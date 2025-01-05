package domain

import "context"

type OrderItem struct {
	OrderId   int `json:"order_id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type OrderItemRepository interface {
	AddItem(c context.Context, item *OrderItem) error
	DeleteItem(c context.Context, orderId, productId int) error
	ModifyItem(c context.Context, orderId, productId, quantity int) error
}

type OrderItemUsecase interface {
	AddChart(c context.Context, item *OrderItem) error
	DeleteItem(c context.Context, orderId, productId int) error
	ModifyQuantity(c context.Context, orderId, productId, quantity int) error
}
