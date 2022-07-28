package model

import "time"

type Category struct {
	ID       uint      `json:"id"'`
	Name     string    `json:"name"`
	CreateAt time.Time `json:"create-at"`
	UpdateAt time.Time `json:"update_at"`
}
