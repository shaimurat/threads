package usecases

import (
	"twiteer/internal/domain/models"
	"twiteer/internal/domain/repos"
)

type ReactionUC struct {
	reactionRepository repos.ReactionRepo
}

func NewReactionUC(reactionRepo repos.ReactionRepo) *ReactionUC {
	return &ReactionUC{reactionRepository: reactionRepo}
}

func (uc *ReactionUC) LikeTret(userID string, tretID string) error {
	return uc.reactionRepository.LikeTret(userID, tretID)
}
func (uc *ReactionUC) DislikeTret(userID string, tretID string) error {
	return uc.reactionRepository.DislikeTret(userID, tretID)
}

func (uc *ReactionUC) GetLikedUsers(tretID string) ([]models.User, error) {
	return uc.reactionRepository.GetLikedUsers(tretID)
}
func (uc *ReactionUC) GetDislikedUsers(tretID string) ([]models.User, error) {
	return uc.reactionRepository.GetDislikedUsers(tretID)
}

func (uc *ReactionUC) GetLikedTrets(userID string) ([]models.Tret, error) {
	return uc.reactionRepository.GetLikedTrets(userID)
}
func (uc *ReactionUC) GetDislikedTrets(userID string) ([]models.Tret, error) {
	return uc.reactionRepository.GetDislikedTrets(userID)
}
