package postgreModels

type ReactionDoc struct {
	IsLike bool `gorm:"not null" json:"isLike"`
	UserID uint `gorm:"not null;uniqueIndex:idx_user_tret" json:"userId"`
	TretID uint `gorm:"not null;uniqueIndex:idx_user_tret" json:"tretId"`

	User UserDoc `gorm:"foreignKey:UserID;constraint:onDelete:CASCADE" json:"user"`
	Tret TretDoc `gorm:"foreignKey:TretID;constraint:onDelete:CASCADE" json:"tret"`
}
