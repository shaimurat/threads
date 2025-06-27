package models

import "time"

type User struct {
	ID        string
	Username  string
	Name      string
	Surname   string
	Email     string
	Password  string
	AvatarUrl string
	UpdatedAt time.Time
	CreatedAt time.Time
}
