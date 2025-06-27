package repos

import "twiteer/internal/domain/models"

type ReactionRepo interface {
	LikeTret(userID string, tretID string) error
	DislikeTret(userID string, tretID string) error
	GetLikedUsers(tretID string) ([]models.User, error)
	GetDislikedUsers(tretID string) ([]models.User, error)
	GetLikedTrets(UserID string) ([]models.Tret, error)
	GetDislikedTrets(UserID string) ([]models.Tret, error)
}
