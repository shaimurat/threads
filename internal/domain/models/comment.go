package models

import "time"

type Comment struct {
	ID        uint
	Text      string
	Likes     uint
	Dislikes  uint
	UpdatedAt time.Time
	CreatedAt time.Time
	UserId    uint
	TretId    uint
}
