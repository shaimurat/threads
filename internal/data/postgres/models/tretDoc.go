package models

import "time"

type TretDoc struct {
	ID          uint      `gorm:"primary_key; auto_increment;"`
	Theme       string    `gorm:"not null; size:50;"`
	Description string    `gorm:"not null"`
	Likes       uint      `gorm:"default:0"`
	Dislikes    uint      `gorm:"default:0"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP; autoUpdateTime"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP; autoCreateTime"`
	UserID      uint      `gorm:"not null"`

	User UserDoc `gorm:"foreignKey:UserID; constraint:onCascade:DELETE"`
}
