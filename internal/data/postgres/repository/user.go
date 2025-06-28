package repository

import (
	"gorm.io/gorm"
	"twiteer/internal/domain/models"
)

type UserRepoSQL struct {
	db *gorm.DB
}

func (r *UserRepoSQL) Create(user models.User) error {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepoSQL) Update(user models.User) error {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepoSQL) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepoSQL) GetUser(id string) (models.User, error) {
	//TODO implement me
	panic("implement me")
}
