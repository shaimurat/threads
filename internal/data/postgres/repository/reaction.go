package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"strconv"
	"twiteer/internal/data/postgres/mapper"
	postgreModels "twiteer/internal/data/postgres/models"
	"twiteer/internal/domain/models"
)

type ReactionRepoSQL struct {
	db *gorm.DB
}

func NewReactionRepoSQL(db *gorm.DB) *ReactionRepoSQL {
	return &ReactionRepoSQL{db: db}
}
func (r *ReactionRepoSQL) setReaction(userIDStr, tretIDStr string, isLike bool) error {
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		log.Printf("[ERROR] Invalid userID: %v", err)
		return err
	}
	tretID, err := strconv.ParseUint(tretIDStr, 10, 64)
	if err != nil {
		log.Printf("[ERROR] Invalid tretID: %v", err)
		return err
	}

	reaction := postgreModels.ReactionDoc{
		UserID: uint(userID),
		TretID: uint(tretID),
		IsLike: isLike,
	}

	log.Printf("[SET REACTION] userID=%d tretID=%d isLike=%v", userID, tretID, isLike)

	// Delete opposite reaction (if exists)
	if err := r.db.
		Where("user_id = ? AND tret_id = ? AND is_like <> ?", userID, tretID, isLike).
		Delete(&postgreModels.ReactionDoc{}).Error; err != nil {
		log.Printf("[ERROR] Deleting opposite reaction: %v", err)
		return err
	}

	// Insert or update reaction
	err = r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "user_id"}, {Name: "tret_id"}},
		DoUpdates: clause.AssignmentColumns([]string{"is_like"}),
	}).Create(&reaction).Error

	if err != nil {
		log.Printf("[ERROR] Upserting reaction: %v", err)
		return err
	}

	log.Printf("[SUCCESS] Reaction set for userID=%d tretID=%d", userID, tretID)
	return nil
}

func (r *ReactionRepoSQL) LikeTret(userID string, tretID string) error {
	return r.setReaction(userID, tretID, true)
}

func (r *ReactionRepoSQL) DislikeTret(userID string, tretID string) error {
	return r.setReaction(userID, tretID, false)
}

func (r *ReactionRepoSQL) GetLikedUsers(tretID string) ([]models.User, error) {
	return r.getUsersByReaction(tretID, true)
}

func (r *ReactionRepoSQL) GetDislikedUsers(tretID string) ([]models.User, error) {
	return r.getUsersByReaction(tretID, false)
}
func (r *ReactionRepoSQL) getUsersByReaction(tretIDStr string, isLike bool) ([]models.User, error) {
	tretID, err := strconv.ParseUint(tretIDStr, 10, 64)
	if err != nil {
		return nil, err
	}

	var reactions []postgreModels.ReactionDoc
	err = r.db.Preload("User").Where("tret_id = ? AND is_like = ?", tretID, isLike).Find(&reactions).Error
	if err != nil {
		return nil, err
	}

	var users []models.User
	for _, r := range reactions {
		users = append(users, mapper.ToUser(r.User))
	}
	return users, nil
}
func (r *ReactionRepoSQL) GetLikedTrets(userID string) ([]models.Tret, error) {
	return r.getTretsByReaction(userID, true)
}

func (r *ReactionRepoSQL) GetDislikedTrets(userID string) ([]models.Tret, error) {
	return r.getTretsByReaction(userID, false)
}
func (r *ReactionRepoSQL) getTretsByReaction(userIDStr string, isLike bool) ([]models.Tret, error) {
	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return nil, err
	}

	var reactions []postgreModels.ReactionDoc
	err = r.db.Preload("Tret.User").Where("user_id = ? AND is_like = ?", userID, isLike).Find(&reactions).Error
	if err != nil {
		return nil, err
	}

	var trets []models.Tret
	for _, r := range reactions {
		trets = append(trets, mapper.ToTret(r.Tret))
	}
	return trets, nil
}
