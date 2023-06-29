package model

import "time"

type Session struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"id_user"`
	User      User      `json:"user"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
}
