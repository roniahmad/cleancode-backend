package route

import (
	"database/sql"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"wetees.com/api/controller"
	"wetees.com/domain"
	"wetees.com/repository"
	"wetees.com/usecase"
)

func NewOrderItemRoute(conf *domain.Config, timeout time.Duration, logger zerolog.Logger, db *sql.DB, app fiber.Router) {
	ur := repository.NewOrderItemRepository(db)

	ctr := &controller.OrderItemController{
		Uc:     usecase.NewOrderItemUsecase(ur, conf, timeout),
		Conf:   conf,
		Logger: logger,
	}

	app.Post("/additem", ctr.AddOrderItem)
	app.Put("/updateqty", ctr.UpdateQuantityOrderItem)
	app.Delete("/deleteitem", ctr.DeleteOrderItem)
}
