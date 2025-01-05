package repository

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"wetees.com/domain"
)

const (
	TABLE_TRANSACTION = "transaction"
	TABLE_ACCOUNT     = "account"
)

type transactionRepository struct {
	Conn *sql.DB
}

/*
Deposit
Note : deposit should add account balance
*/
func (t *transactionRepository) Deposit(c context.Context, trans *domain.Transaction) (id int, err error) {
	var (
		stmt         *sql.Stmt
		result       sql.Result
		lastInsertId int64
	)

	query := fmt.Sprintf(`INSERT INTO %s (acc_number, trans_type, amount, trans_date)	
		VALUES(?, ?, ?, ?)`, TABLE_TRANSACTION)

	if stmt, err = t.Conn.PrepareContext(c, query); err != nil {
		return 0, err
	}

	if result, err = stmt.ExecContext(c, &trans.AccNumber, &trans.TransType, &trans.Amount, time.Now().Local().UTC()); err != nil {
		return 0, err
	}

	lastInsertId, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertId), nil
}

// Check if there is enough balance available
func (t *transactionRepository) GetBalanceAvailable(c context.Context, accNumber string) (float64, error) {
	var balance float64 = 0
	query := fmt.Sprintf("SELECT balance FROM %s WHERE acc_number=?", TABLE_ACCOUNT)
	if err := t.Conn.QueryRowContext(c, query, accNumber).Scan(&balance); err != nil {
		return 0, err
	}

	return balance, nil
}

// Update Balance
func (t *transactionRepository) UpdateBalance(c context.Context, accNumber string, balance float64) (err error) {
	var (
		stmt   *sql.Stmt
		result sql.Result
	)

	query := fmt.Sprintf("UPDATE %s set balance=? WHERE acc_number=?", TABLE_ACCOUNT)
	if stmt, err = t.Conn.PrepareContext(c, query); err != nil {
		return
	}

	if result, err = stmt.ExecContext(c, balance, accNumber); err != nil {
		return
	}

	_, err = result.RowsAffected()

	return
}

/*
Withdraws
Note : withdraws should check available account balance and reduce account balance
*/
func (t *transactionRepository) Withdraws(c context.Context, trans *domain.Transaction) (id int, err error) {
	var (
		stmt         *sql.Stmt
		result       sql.Result
		lastInsertId int64
	)

	query := fmt.Sprintf(`INSERT INTO %s (acc_number, trans_type, amount, trans_date)	
		VALUES(?, ?, ?, ?)`, TABLE_TRANSACTION)

	if stmt, err = t.Conn.PrepareContext(c, query); err != nil {
		return 0, err
	}

	if result, err = stmt.ExecContext(c, &trans.AccNumber, &trans.TransType, &trans.Amount, time.Now().Local().UTC()); err != nil {
		return 0, err
	}

	lastInsertId, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(lastInsertId), nil
}

func NewTransactionRepository(db *sql.DB) domain.TransactionRepository {
	return &transactionRepository{
		Conn: db,
	}
}
