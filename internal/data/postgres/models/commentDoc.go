package postgreModels

import "time"

type CommentDoc struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Text      string    `gorm:"not null" json:"text"`
	Likes     uint      `gorm:"default:0" json:"likes"`
	Dislikes  uint      `gorm:"default:0" json:"dislikes"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updatedAt"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UserID    uint      `gorm:"not null" json:"userId"`
	TretID    uint      `gorm:"not null" json:"tretId"`

	User UserDoc `gorm:"foreignKey:UserID;constraint:onDelete:CASCADE" json:"user"`
	Tret TretDoc `gorm:"foreignKey:TretID;constraint:onDelete:CASCADE" json:"tret"`
}
