package model

import "time"

type Status string

const (
	Created       string = "new"
	Inproccessing string = "proccesing"
	Finished      string = "done"
)

type Task struct {
	ID       int64     `json:"id"`
	Title    string    `json:"title"`
	CreateAt time.Time `json:"create_at"`
	Status
	UpdateAt time.Time `json:"update_at"`
}
