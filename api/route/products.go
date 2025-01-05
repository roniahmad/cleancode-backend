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

func NewProductRoute(conf *domain.Config, timeout time.Duration, logger zerolog.Logger, db *sql.DB, app fiber.Router) {
	ur := repository.NewProductRepository(db)

	ctr := &controller.ProductController{
		Uc:     usecase.NewProductUsecase(ur, conf, timeout),
		Conf:   conf,
		Logger: logger,
	}

	app.Get("/viewproduct", ctr.ViewProduct)
	app.Get("/searchproduct", ctr.SearchProduct)
}
