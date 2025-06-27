package models

import "time"

type Comment struct {
	ID        string
	Text      string
	Likes     uint
	Dislikes  uint
	UpdatedAt time.Time
	CreatedAt time.Time
	UserId    string
	TretId    string
}
