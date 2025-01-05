package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"wetees.com/domain"
	"wetees.com/internal/jwt"
	"wetees.com/internal/response"
	"wetees.com/internal/vars"
	"wetees.com/internal/wmvalidator"
)

type UserController struct {
	Uc     domain.UserUsecase
	Conf   *domain.Config
	Logger zerolog.Logger
}

/*
ChangePassword Handler
*/
func (ctr *UserController) ChangePassword(c *fiber.Ctx) error {
	request := domain.ChangePassword{}

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

	//All email sanity check should be done at client
	if err = ctr.Uc.ChangePassword(c.Context(), claims["email"].(string),
		request.OldPassword, request.NewPassword, request.ConfirmPassword); err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, vars.ErrChangePasswordFailed)
	}

	return response.SendSuccessResponse(c, ctr.Conf.Mode, ctr.Logger, vars.ChangePasswordSucceed)
}
