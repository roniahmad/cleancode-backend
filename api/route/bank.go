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

func NewBankRoute(conf *domain.Config, timeout time.Duration, logger zerolog.Logger, db *sql.DB, app fiber.Router) {
	ur := repository.NewTransactionRepository(db)

	ctr := &controller.BankController{
		Uc:     usecase.NewTransactionUsecase(ur, conf, timeout),
		Conf:   conf,
		Logger: logger,
	}

	app.Post("/deposit", ctr.Deposit)
	app.Post("/withdraws", ctr.Withdraws)
}
