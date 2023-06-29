package model

import "time"

type History struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"id_account"`
	Account   Account   `json:"account"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}
