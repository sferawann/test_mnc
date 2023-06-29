package model

import "time"

type Account struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"id_user"`
	User      User      `json:"user"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}
