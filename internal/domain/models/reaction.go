package models

type Reaction struct {
	IsLike bool
	UserID uint
	TretID uint

	User User
	Tret Tret
}
