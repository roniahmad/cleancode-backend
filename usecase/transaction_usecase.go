package usecase

import (
	"context"
	"sync"
	"time"

	"wetees.com/domain"
	"wetees.com/internal/vars"
)

type transactionUsecase struct {
	repo           domain.TransactionRepository
	conf           *domain.Config
	contextTimeout time.Duration
}

var (
	mtx                    = sync.Mutex{}
	accountBalance float64 = 0
)

// Deposit
func (t *transactionUsecase) Deposit(ctx context.Context, trans *domain.Transaction) (id int, err error) {
	var (
		accNumber string  = trans.AccNumber
		amount    float64 = trans.Amount
		total     float64 = 0
		wg        sync.WaitGroup
	)

	wg.Add(2)
	//1
	chDepositId := make(chan int)
	go func(wg *sync.WaitGroup) error {
		defer wg.Done()
		trx := domain.Transaction{
			AccNumber: accNumber,
			TransType: "deposit",
			Amount:    amount,
		}
		// add deposit transaction
		id, err = t.repo.Deposit(ctx, &trx)
		if err != nil {
			return err
		}

		chDepositId <- id

		return nil
	}(&wg)
	id = <-chDepositId

	//2
	go func(wg *sync.WaitGroup) error {
		defer wg.Done()
		mtx.Lock()
		accountBalance, err = t.repo.GetBalanceAvailable(ctx, trans.AccNumber)
		if err != nil {
			return err
		}

		total = accountBalance + amount
		err = t.repo.UpdateBalance(ctx, accNumber, total)
		if err != nil {
			return err
		}
		mtx.Unlock()

		return nil
	}(&wg)

	wg.Wait()

	close(chDepositId)

	return
}

// Withdraws
func (t *transactionUsecase) Withdraws(ctx context.Context, trans *domain.Transaction) (int, error) {
	var (
		accNumber string  = trans.AccNumber
		amount    float64 = trans.Amount
		total     float64 = 0
		wg        sync.WaitGroup
	)

	wg.Add(1)

	type result struct {
		err error
		id  int
	}

	//1
	ch := make(chan result)
	go func(wg *sync.WaitGroup) error {
		defer wg.Done()

		mtx.Lock()
		accountBalance, _ = t.repo.GetBalanceAvailable(ctx, accNumber)

		total = accountBalance - amount
		if total <= 0 {
			res := new(result)
			res.err = vars.ErrNotEnoughBalance
			res.id = 0
			ch <- *res
		} else {
			trx := domain.Transaction{
				AccNumber: accNumber,
				TransType: "withdrawal",
				Amount:    amount,
			}

			// add withdrawal transaction
			id, err := t.repo.Withdraws(ctx, &trx)
			if err != nil {
				res := new(result)
				res.err = vars.ErrNotEnoughBalance
				res.id = 0
				ch <- *res
				return err
			}

			err = t.repo.UpdateBalance(ctx, accNumber, total)
			if err != nil {
				res := new(result)
				res.err = vars.ErrNotEnoughBalance
				res.id = 0
				ch <- *res
				return err
			}
			res := new(result)
			res.err = nil
			res.id = id
			ch <- *res
		}

		mtx.Unlock()

		return nil
	}(&wg)

	value := <-ch

	wg.Wait()

	close(ch)

	return value.id, value.err
}

func NewTransactionUsecase(repo domain.TransactionRepository, conf *domain.Config, timeout time.Duration) domain.TransactionUsecase {
	return &transactionUsecase{
		repo:           repo,
		conf:           conf,
		contextTimeout: timeout,
	}
}
