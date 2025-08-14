package model

import "time"

type Task struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	CreateAt time.Time `json:"create_at"`
}
