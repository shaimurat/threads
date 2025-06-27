package usecases

import (
	"twiteer/internal/domain/models"
	"twiteer/internal/domain/repos"
)

type CommentUC struct {
	comRepository repos.CommentRepo
}

func NewCommentUC(comRepo repos.CommentRepo) *CommentUC {
	return &CommentUC{comRepo}
}

func (uc *CommentUC) GetComment(id string) (models.Comment, error) {
	return uc.comRepository.GetComment(id)
}
func (uc *CommentUC) Create(com models.Comment) error {
	return uc.comRepository.Create(com)
}
func (uc *CommentUC) Delete(id string) error {
	return uc.comRepository.Delete(id)
}
func (uc *CommentUC) Update(com models.Comment) error {
	return uc.comRepository.Update(com)
}
