package repos

import "twiteer/internal/domain/models"

type TretRepo interface {
	Create(tret models.Tret) error
	Update(tret models.Tret) error
	Delete(id string) error
	GetTret(id string) (models.Tret, error)
	GetUserTrets(id string) ([]models.Tret, error)
	LikeTret(id string) error
	DislikeTret(id string) error
}
