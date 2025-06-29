package postgreModels

import "time"

type UserDoc struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Username  string    `gorm:"unique" json:"username"`
	Name      string    `gorm:"not null" json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `gorm:"not null" json:"email"`
	Password  string    `gorm:"not null" json:"password"`
	AvatarUrl string    `json:"avatarUrl"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;autoUpdateTime" json:"updatedAt"`
}
