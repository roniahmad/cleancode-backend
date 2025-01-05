package domain

import "database/sql"

type Account struct {
	ID        int          `json:"id"`
	UserId    int          `json:"user_id"`
	AccType   string       `json:"acc_type"`
	AccNumber string       `json:"acc_number"`
	Balance   float64      `json:"balance"`
	Dto       sql.NullTime `json:"dto"`
	Dtc       sql.NullTime `json:"dtc"`
	Status    string       `json:"status"`
}
