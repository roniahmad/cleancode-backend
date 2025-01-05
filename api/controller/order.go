package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"wetees.com/domain"
	"wetees.com/internal/jwt"
	"wetees.com/internal/response"
	"wetees.com/internal/vars"
	"wetees.com/internal/wmvalidator"
)

type OrderController struct {
	Uc     domain.OrderUsecase
	Conf   *domain.Config
	Logger zerolog.Logger
}

/*
Create Order Handler
*/
func (ctr *OrderController) CreateOrder(c *fiber.Ctx) error {
	xAuthHeader := c.Get(fiber.HeaderAuthorization)
	authToken, _ := jwt.GetAuthToken(xAuthHeader)

	claims, err := jwt.ExtractClaimsFromToken(authToken, ctr.Conf.AccessTokenSecret)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusNotFound, ctr.Conf.Mode, ctr.Logger, err)
	}

	userId, _ := strconv.Atoi(claims["id"].(string))
	data, err := ctr.Uc.CreateOrder(c.Context(), userId)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, vars.ErrCreateOrderFailed)
	}

	return response.SendDataResponse(c, ctr.Conf.Mode, ctr.Logger, data)
}

/*
Cancel Order Handler
*/
func (ctr *OrderController) CancelOrder(c *fiber.Ctx) error {
	request := domain.CancelOrder{}

	if err := c.BodyParser(&request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	if err := wmvalidator.Validate(request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	xAuthHeader := c.Get(fiber.HeaderAuthorization)
	authToken, _ := jwt.GetAuthToken(xAuthHeader)

	claims, err := jwt.ExtractClaimsFromToken(authToken, ctr.Conf.AccessTokenSecret)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusNotFound, ctr.Conf.Mode, ctr.Logger, err)
	}

	userId, _ := strconv.Atoi(claims["id"].(string))
	err = ctr.Uc.CancelOrder(c.Context(), request.OrderId, userId)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, vars.ErrCancelOrderFailed)
	}

	return response.SendSuccessResponse(c, ctr.Conf.Mode, ctr.Logger, vars.CancelOrderSucceed)
}

/*
Checkout Order Handler
*/
func (ctr *OrderController) CheckoutOrder(c *fiber.Ctx) error {
	request := domain.CheckoutOrder{}

	if err := c.BodyParser(&request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	if err := wmvalidator.Validate(request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	xAuthHeader := c.Get(fiber.HeaderAuthorization)
	authToken, _ := jwt.GetAuthToken(xAuthHeader)

	claims, err := jwt.ExtractClaimsFromToken(authToken, ctr.Conf.AccessTokenSecret)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusNotFound, ctr.Conf.Mode, ctr.Logger, err)
	}

	userId, _ := strconv.Atoi(claims["id"].(string))
	items, totalItem, totalPrice, err := ctr.Uc.CheckoutOrder(c.Context(), request.OrderId, userId)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, vars.ErrCheckoutOrderFailed)
	}

	return response.SendCheckoutDataResponse(c, ctr.Conf.Mode, ctr.Logger, items, totalItem, totalPrice, request.OrderId, vars.StatusAWP)
}
