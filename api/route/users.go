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

func NewUserRoute(conf *domain.Config, timeout time.Duration, logger zerolog.Logger, db *sql.DB, app fiber.Router) {
	ur := repository.NewUserRepository(db)

	ctr := &controller.UserController{
		Uc:     usecase.NewUserUsecase(ur, conf, timeout),
		Conf:   conf,
		Logger: logger,
	}

	app.Put("/changepassword", ctr.ChangePassword)
}
