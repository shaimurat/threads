package usecases

import (
	"twiteer/internal/domain/models"
	"twiteer/internal/domain/repos"
)

type TretUC struct {
	tretRepository repos.TretRepo
}

func NewTretUC(tretRepository repos.TretRepo) *TretUC {
	return &TretUC{tretRepository: tretRepository}
}

func (uc *TretUC) Create(tret models.Tret) error {
	return uc.tretRepository.Create(tret)
}
func (uc *TretUC) Update(tret models.Tret) error {
	return uc.tretRepository.Update(tret)
}
func (uc *TretUC) Delete(id string) error {
	return uc.tretRepository.Delete(id)
}
func (uc *TretUC) GetTret(id string) (models.Tret, error) {
	return uc.tretRepository.GetTret(id)
}
func (uc *TretUC) GetUserTrets(id string) ([]models.Tret, error) {
	return uc.tretRepository.GetUserTrets(id)
}
func (uc *TretUC) LikeTret(id string) error {
	return uc.tretRepository.LikeTret(id)
}
func (uc *TretUC) DislikeTret(id string) error {
	return uc.tretRepository.DislikeTret(id)
}
