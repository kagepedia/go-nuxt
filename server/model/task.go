package model

import "time"

type Task struct {
	Id       int64
	Task     string
	CreateAt time.Time
	UpdateAt time.Time
}
