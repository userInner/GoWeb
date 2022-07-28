package model

import (
	"github.com/google/uuid"
	"time"
)

type Post struct {
	ID         uuid.UUID `json:"id"`
	UserId     uint      `json:"user_id"`
	CategoryId uint      `json:"category_id"`
	Category   *Category `json:"catagory"`
	Title      string    `json:"title"`
	HeadImg    string    `json:"head_img"`
	Content    string    `json:"content"`
	CreateAt   time.Time `json:"create-at"`
	UpdateAt   time.Time `json:"update_at"`
}
