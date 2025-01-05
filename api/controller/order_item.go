package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"wetees.com/domain"
	"wetees.com/internal/response"
	"wetees.com/internal/vars"
	"wetees.com/internal/wmvalidator"
)

type OrderItemController struct {
	Uc     domain.OrderItemUsecase
	Conf   *domain.Config
	Logger zerolog.Logger
}

/*
Add Order Item Handler
*/
func (ctr *OrderItemController) AddOrderItem(c *fiber.Ctx) error {
	request := domain.ModifyQuantityOrderItem{}

	if err := c.BodyParser(&request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	if err := wmvalidator.Validate(request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	item := domain.OrderItem{
		OrderId:   request.OrderId,
		ProductId: request.ProductId,
		Quantity:  request.Quantity,
	}

	err := ctr.Uc.AddChart(c.Context(), &item)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	return response.SendSuccessResponse(c, ctr.Conf.Mode, ctr.Logger, vars.CreateOrderSucceed)
}

/*
Cancel Order Handler
*/
func (ctr *OrderItemController) DeleteOrderItem(c *fiber.Ctx) error {
	request := domain.RemoveOrderItem{}

	if err := c.BodyParser(&request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	if err := wmvalidator.Validate(request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	if err := ctr.Uc.DeleteItem(c.Context(), request.OrderId, request.ProductId); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, vars.ErrRemoveOrderItemFailed)
	}

	return response.SendSuccessResponse(c, ctr.Conf.Mode, ctr.Logger, vars.RemoveOrderItemSucceed)
}

/*
Update Quantity Order Item Handler
*/
func (ctr *OrderItemController) UpdateQuantityOrderItem(c *fiber.Ctx) error {
	request := domain.ModifyQuantityOrderItem{}

	if err := c.BodyParser(&request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	if err := wmvalidator.Validate(request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	if err := ctr.Uc.ModifyQuantity(c.Context(), request.OrderId, request.ProductId, request.Quantity); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, vars.ErrUpdateQtyFailed)
	}

	return response.SendSuccessResponse(c, ctr.Conf.Mode, ctr.Logger, vars.UpdateQtySucceed)
}
