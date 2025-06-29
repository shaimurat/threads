package postgreModels

import "time"

type TretDoc struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Theme       string    `gorm:"not null;size:50" json:"theme"`
	Description string    `gorm:"not null" json:"description"`
	Likes       uint      `gorm:"default:0" json:"likes"`
	Dislikes    uint      `gorm:"default:0" json:"dislikes"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updatedAt"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP;autoCreateTime" json:"createdAt"`
	UserID      uint      `gorm:"not null" json:"userId"`

	User UserDoc `gorm:"foreignKey:UserID;constraint:onDelete:CASCADE" json:"user"`
}
