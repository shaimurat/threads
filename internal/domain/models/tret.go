package models

import "time"

type Tret struct {
	ID          uint
	Theme       string
	Description string
	Likes       uint
	Dislikes    uint
	UpdatedAt   time.Time
	CreatedAt   time.Time
	UserID      uint

	User User
}
