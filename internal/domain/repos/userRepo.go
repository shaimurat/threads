package repos

import "twiteer/internal/domain/models"

type UserRepo interface {
	Create(user models.User) error
	Update(user models.User) error
	Delete(id string) error
	GetUser(id string) (models.User, error)
}
