package model

import "time"

type Account struct {
	Id        int64              `json:"id"`
	UserId    int64              `json:"user_id"`
	Balance   map[string]float64 `json:"balance"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
