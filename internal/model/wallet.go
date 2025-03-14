package model

import "time"

type Wallet struct {
	Id        int64     `json:"id"`
	UserId    string    `json:"user_id"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
