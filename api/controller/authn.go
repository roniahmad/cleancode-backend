package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"wetees.com/domain"
	"wetees.com/internal/response"
	"wetees.com/internal/vars"
	"wetees.com/internal/wmvalidator"
)

type AuthController struct {
	Uc     domain.UserUsecase
	Conf   *domain.Config
	Logger zerolog.Logger
}

/*
RegisterUser Handler
For a new user to be able to use all features,they have to be registered.
Registration form contains:

	Name
	Email
	Password
*/
func (ctr *AuthController) RegisterUser(c *fiber.Ctx) error {
	request := domain.User{}

	err := c.BodyParser(&request)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	if err = wmvalidator.Validate(request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	if err = ctr.Uc.Register(c.Context(), &request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	return response.SendSuccessResponse(c, ctr.Conf.Mode, ctr.Logger, vars.RegisterUserSucceed)
}

/*
Login Handler
*/
func (ctr *AuthController) Login(c *fiber.Ctx) error {
	var (
		request = domain.Login{}
		data    = map[string]interface{}{}
	)

	err := c.BodyParser(&request)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	if err = wmvalidator.Validate(request); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}

	data, err = ctr.Uc.Login(c.Context(), request)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, vars.ErrLoginFailed)
	}

	return response.SendDataResponse(c, ctr.Conf.Mode, ctr.Logger, data)
}
