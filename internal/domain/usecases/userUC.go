package usecases

import (
	"twiteer/internal/domain/models"
	"twiteer/internal/domain/repos"
)

type UserUC struct {
	userRepository repos.UserRepo
}

func NewUserUC(userRepository repos.UserRepo) *UserUC {
	return &UserUC{userRepository}
}

func (uc *UserUC) Create(user models.User) error {
	return uc.userRepository.Create(user)
}
func (uc *UserUC) Update(user models.User) error {
	return uc.userRepository.Update(user)
}
func (uc *UserUC) Delete(id string) error {
	return uc.userRepository.Delete(id)
}
func (uc *UserUC) GetUser(id string) (models.User, error) {
	return uc.userRepository.GetUser(id)
}
