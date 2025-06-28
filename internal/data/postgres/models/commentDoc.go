package models

import "time"

type CommentDoc struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Text      string    `gorm:"not null"`
	Likes     uint      `gorm:"default:0"`
	Dislikes  uint      `gorm:"default:0"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP; autoUpdateTime"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UserID    uint      `gorm:"not null"`
	TretID    uint      `gorm:"not null"`

	User UserDoc `gorm:"foreignKey:UserID; constraint:onCascade:DELETE"`
	Tret TretDoc `gorm:"foreignKey:TretID; constraint:onCascade:DELETE"`
}
