package repository

import "twiteer/internal/domain/models"

type ReactionRepoSQL struct{}

func (r *ReactionRepoSQL) LikeTret(userID string, tretID string) error {
	//TODO implement me
	panic("implement me")
}

func (r *ReactionRepoSQL) DislikeTret(userID string, tretID string) error {
	//TODO implement me
	panic("implement me")
}

func (r *ReactionRepoSQL) GetLikedUsers(tretID string) ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ReactionRepoSQL) GetDislikedUsers(tretID string) ([]models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ReactionRepoSQL) GetLikedTrets(UserID string) ([]models.Tret, error) {
	//TODO implement me
	panic("implement me")
}

func (r *ReactionRepoSQL) GetDislikedTrets(UserID string) ([]models.Tret, error) {
	//TODO implement me
	panic("implement me")
}
