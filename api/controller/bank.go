package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"wetees.com/domain"
	"wetees.com/internal/response"
	"wetees.com/internal/vars"
	"wetees.com/internal/wmvalidator"
)

type BankController struct {
	Uc     domain.TransactionUsecase
	Conf   *domain.Config
	Logger zerolog.Logger
}

/*
Deposit Handler
*/
func (ctr *BankController) Deposit(c *fiber.Ctx) error {
	request := domain.BankTrans{}

	if err := c.BodyParser(&request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	if err := wmvalidator.Validate(request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	item := domain.Transaction{
		AccNumber: request.AccNumber,
		Amount:    request.Amount,
	}

	_, err := ctr.Uc.Deposit(c.Context(), &item)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, vars.ErrDepositFailed)
	}

	return response.SendSuccessResponse(c, ctr.Conf.Mode, ctr.Logger, vars.DepositSucceed)
}

/*
Withdraws Handler
*/
func (ctr *BankController) Withdraws(c *fiber.Ctx) error {
	request := domain.BankTrans{}

	if err := c.BodyParser(&request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	if err := wmvalidator.Validate(request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	item := domain.Transaction{
		AccNumber: request.AccNumber,
		Amount:    request.Amount,
	}

	_, err := ctr.Uc.Withdraws(c.Context(), &item)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, vars.ErrWithdrawsFailed)
	}

	return response.SendSuccessResponse(c, ctr.Conf.Mode, ctr.Logger, vars.WithdrawsSucceed)
}
