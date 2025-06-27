package repos

import "twiteer/internal/domain/models"

type CommentRepo interface {
	Create(com models.Comment) error
	Update(com models.Comment) error
	Delete(id string) error
	GetComment(id string) (models.Comment, error)
}
