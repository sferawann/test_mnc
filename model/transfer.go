package model

import "time"

type Transfer struct {
	ID            int64     `json:"id"`
	FromAccountID int64     `json:"from_account_id"`
	FromAccount   Account   `json:"from_account"`
	ToAccountID   int64     `json:"to_account_id"`
	ToAccount     Account   `json:"to_account"`
	Amount        float64   `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}
