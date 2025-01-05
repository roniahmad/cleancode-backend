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

func NewAuthnRoute(conf *domain.Config, timeout time.Duration, logger zerolog.Logger, db *sql.DB, app fiber.Router) {
	ur := repository.NewUserRepository(db)

	ctr := &controller.AuthController{
		Uc:     usecase.NewUserUsecase(ur, conf, timeout),
		Conf:   conf,
		Logger: logger,
	}

	app.Post("/login", ctr.Login)
	app.Post("/register", ctr.RegisterUser)
}
