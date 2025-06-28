package repository

import "twiteer/internal/domain/models"

type CommRepoSQL struct{}

func (r *CommRepoSQL) Create(com models.Comment) error {
	//TODO implement me
	panic("implement me")
}

func (r *CommRepoSQL) Update(com models.Comment) error {
	//TODO implement me
	panic("implement me")
}

func (r *CommRepoSQL) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (r *CommRepoSQL) GetComment(id string) (models.Comment, error) {
	//TODO implement me
	panic("implement me")
}
