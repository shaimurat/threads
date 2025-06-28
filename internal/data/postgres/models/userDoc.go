package models

import "time"

type UserDoc struct {
	ID        uint   `gorm:"primary_key; auto_increment"`
	Username  string `gorm:"unique"`
	Name      string `gorm:"not null"`
	Surname   string
	Email     string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	AvatarUrl string    `gorm:""`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP; autoCreateTime"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP; autoUpdateTime"`
}
