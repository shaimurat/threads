package repository

import "twiteer/internal/domain/models"

type TretRepoSQL struct{}

func (r *TretRepoSQL) Create(tret models.Tret) error {
	//TODO implement me
	panic("implement me")
}

func (r *TretRepoSQL) Update(tret models.Tret) error {
	//TODO implement me
	panic("implement me")
}

func (r *TretRepoSQL) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (r *TretRepoSQL) GetTret(id string) (models.Tret, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TretRepoSQL) GetUserTrets(id string) ([]models.Tret, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TretRepoSQL) LikeTret(id string) error {
	//TODO implement me
	panic("implement me")
}

func (r *TretRepoSQL) DislikeTret(id string) error {
	//TODO implement me
	panic("implement me")
}
