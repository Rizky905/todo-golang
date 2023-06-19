package model

import "time"

type Todo struct {
	ID         string    `json:"id"`
	Title      string    `json:"title"`
	IsDone     bool      `json:"is_done"`
	CreatedAdd time.Time `json:"created_at"`
}
