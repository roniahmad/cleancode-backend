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

func NewOrderRoute(conf *domain.Config, timeout time.Duration, logger zerolog.Logger, db *sql.DB, app fiber.Router) {
	ur := repository.NewOrderRepository(db)

	ctr := &controller.OrderController{
		Uc:     usecase.NewOrderUsecase(ur, conf, timeout),
		Conf:   conf,
		Logger: logger,
	}

	app.Post("/order", ctr.CreateOrder)
	app.Put("/cancelorder", ctr.CancelOrder)
	app.Put("/checkout", ctr.CheckoutOrder)
}
