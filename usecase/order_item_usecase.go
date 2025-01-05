package usecase

import (
	"context"
	"time"

	"wetees.com/domain"
)

type orderItemUsecase struct {
	repo           domain.OrderItemRepository
	conf           *domain.Config
	contextTimeout time.Duration
}

// Add Chart items
func (o *orderItemUsecase) AddChart(c context.Context, item *domain.OrderItem) (err error) {
	ctx, cancel := context.WithTimeout(c, o.contextTimeout)
	defer cancel()

	if err = o.repo.AddItem(ctx, item); err != nil {
		return err
	}

	return
}

// Delete Item deom chart
func (o *orderItemUsecase) DeleteItem(c context.Context, orderId int, productId int) error {
	ctx, cancel := context.WithTimeout(c, o.contextTimeout)
	defer cancel()

	err := o.repo.DeleteItem(ctx, orderId, productId)

	return err
}

// Modify Quantity
func (o *orderItemUsecase) ModifyQuantity(c context.Context, orderId int, productId int, quantity int) error {
	ctx, cancel := context.WithTimeout(c, o.contextTimeout)
	defer cancel()

	err := o.repo.ModifyItem(ctx, orderId, productId, quantity)

	return err
}

func NewOrderItemUsecase(repo domain.OrderItemRepository, conf *domain.Config, timeout time.Duration) domain.OrderItemUsecase {
	return &orderItemUsecase{
		repo:           repo,
		conf:           conf,
		contextTimeout: timeout,
	}
}
