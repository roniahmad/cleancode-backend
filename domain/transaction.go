package domain

import (
	"context"
	"database/sql"
)

type Transaction struct {
	ID        int          `json:"id"`
	AccNumber string       `json:"acc_number"`
	TransType string       `json:"trans_type"`
	Amount    float64      `json:"amount"`
	TransDate sql.NullTime `json:"trans_date"`
}

type TransactionRepository interface {
	GetBalanceAvailable(c context.Context, accNumber string) (float64, error)
	UpdateBalance(c context.Context, accNumber string, balance float64) error
	Deposit(c context.Context, trans *Transaction) (id int, err error)
	Withdraws(c context.Context, trans *Transaction) (id int, err error)
}

type TransactionUsecase interface {
	Deposit(c context.Context, trans *Transaction) (id int, err error)
	Withdraws(c context.Context, trans *Transaction) (id int, err error)
}
