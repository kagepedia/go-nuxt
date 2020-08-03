package models

import "time"

type Registration struct {
	Pk       int64
	Id       string
	Pw       string
	Mail     string
	Is_ok    int8
	Token    string
	CreateAt time.Time
	UpdateAt time.Time
}
