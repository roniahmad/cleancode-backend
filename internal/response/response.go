package response

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"wetees.com/domain"
	"wetees.com/internal/vars"
)

// Send error response
func SendErrorResponse(c *fiber.Ctx, status int, mode string, log zerolog.Logger, err error) error {
	if mode != vars.ModeProd {
		log.Debug().Msg(err.Error())
	}

	return c.Status(status).JSON(domain.MessageResponse{
		Success: false,
		Message: err.Error(),
	})
}

// Send success response
func SendSuccessResponse(c *fiber.Ctx, mode string, log zerolog.Logger, message string) error {
	if mode != vars.ModeProd {
		log.Debug().Msg(message)
	}

	return c.Status(fiber.StatusOK).JSON(domain.MessageResponse{
		Success: true,
		Message: message,
	})
}

// Send user response with data
func SendDataResponse(c *fiber.Ctx, mode string, log zerolog.Logger, data interface{}) error {
	if mode != vars.ModeProd {
		log.Debug().Msg(fmt.Sprintf("%+v", data))
	}

	return c.Status(fiber.StatusOK).JSON(domain.DataResponse{
		Success: true,
		Payload: data,
	})
}

// Send user response with paginated data
func SendPaginatedDataResponse(c *fiber.Ctx, mode string, log zerolog.Logger, data interface{}, total, page, totalPage, limit int) error {
	if mode != vars.ModeProd {
		log.Debug().Msg(fmt.Sprintf("%+v", data))
	}

	return c.Status(fiber.StatusOK).JSON(domain.PaginatedDataResponse{
		Success:    true,
		Data:       data,
		Total:      total,
		Page:       page,
		TotalPages: totalPage,
		Limit:      limit,
	})
}

// Send user checkout response with data
func SendCheckoutDataResponse(c *fiber.Ctx, mode string, log zerolog.Logger, data interface{}, totalItems int, totalPrice float64, orderId int, message string) error {
	if mode != vars.ModeProd {
		log.Debug().Msg(fmt.Sprintf("%+v", data))
	}

	return c.Status(fiber.StatusOK).JSON(domain.CheckoutDataResponse{
		Success:    true,
		Message:    message,
		Items:      data,
		TotalItems: totalItems,
		TotalPrice: int(totalPrice),
		OrderId:    orderId,
	})
}
