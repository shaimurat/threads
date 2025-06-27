package models

import "time"

type Tret struct {
	ID          string
	Theme       string
	Description string
	Likes       uint
	Dislikes    uint
	UpdatedAt   time.Time
	CreatedAt   time.Time
	UserID      string
}
