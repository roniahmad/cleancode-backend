package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"wetees.com/domain"
	"wetees.com/internal/response"
)

type ProductController struct {
	Uc     domain.ProductUsecase
	Conf   *domain.Config
	Logger zerolog.Logger
}

/*
View Product Handler
*/
func (ctr *ProductController) ViewProduct(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "1"))

	data, count, totalPages, err := ctr.Uc.GetProducts(c.Context(), page, limit)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}
	if page < totalPages {
		page += 1
	}

	return response.SendPaginatedDataResponse(c, ctr.Conf.Mode, ctr.Logger, data, count, page, totalPages, limit)
}

/*
Search Product Handler
*/
func (ctr *ProductController) SearchProduct(c *fiber.Ctx) error {
	item := c.Query("item", "")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "1"))

	data, count, totalPages, err := ctr.Uc.SearchProducts(c.Context(), item, page, limit)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, ctr.Conf.Mode, ctr.Logger, err)
	}
	if page < totalPages {
		page += 1
	}

	return response.SendPaginatedDataResponse(c, ctr.Conf.Mode, ctr.Logger, data, count, page, totalPages, limit)
}
