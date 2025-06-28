package models

type Reaction struct {
	isLike bool    `gorm:"not null"`
	UserID uint    `gorm:"not null"`
	TretID uint    `gorm:"not null"`
	User   UserDoc `gorm:"foreignKey:UserID; constraint:onCascade:DELETE"`
	Tret   TretDoc `gorm:"foreignKey:TretID; constraint:onCascade:DELETE"`
}
