package repos

import "twiteer/internal/domain/models"

type CommentRepo interface {
	Create(com models.Comment) error
	Update(com models.Comment) error
	Delete(id string) error
	LikeComment(id string) error
	DislikeComment(id string) error
	GetComment(id string) (models.Comment, error)
	GetCommentsByTretID(tretID string) ([]models.Comment, error)
}
