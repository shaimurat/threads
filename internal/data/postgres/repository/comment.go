package repository

import (
	"strconv"

	"gorm.io/gorm"

	"twiteer/internal/data/postgres/mapper"
	postgreModels "twiteer/internal/data/postgres/models"
	"twiteer/internal/domain/models"
)

type CommRepoSQL struct {
	db *gorm.DB
}

func NewCommRepoSQL(db *gorm.DB) *CommRepoSQL {
	return &CommRepoSQL{db: db}
}

func (c *CommRepoSQL) Create(com models.Comment) error {
	doc := mapper.FromComment(com)
	return c.db.Create(&doc).Error
}

func (c *CommRepoSQL) Update(com models.Comment) error {
	doc := mapper.FromComment(com)
	return c.db.Save(&doc).Error
}

func (c *CommRepoSQL) Delete(id string) error {
	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	return c.db.Delete(&postgreModels.CommentDoc{}, cid).Error
}

func (c *CommRepoSQL) LikeComment(id string) error {
	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	return c.db.Model(&postgreModels.CommentDoc{}).
		Where("id = ?", cid).
		UpdateColumn("likes", gorm.Expr("likes + ?", 1)).Error
}

func (c *CommRepoSQL) DislikeComment(id string) error {
	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	return c.db.Model(&postgreModels.CommentDoc{}).
		Where("id = ?", cid).
		UpdateColumn("dislikes", gorm.Expr("dislikes + ?", 1)).Error
}

func (c *CommRepoSQL) GetComment(id string) (models.Comment, error) {
	cid, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return models.Comment{}, err
	}

	var doc postgreModels.CommentDoc
	err = c.db.Preload("User").Preload("Tret.User").First(&doc, cid).Error
	if err != nil {
		return models.Comment{}, err
	}

	return mapper.ToComment(doc), nil
}
func (c *CommRepoSQL) GetCommentsByTretID(tretID string) ([]models.Comment, error) {
	tid, err := strconv.ParseUint(tretID, 10, 64)
	if err != nil {
		return nil, err
	}

	var docs []postgreModels.CommentDoc
	err = c.db.Preload("User").Preload("Tret.User").
		Where("tret_id = ?", tid).
		Order("created_at ASC").
		Find(&docs).Error

	if err != nil {
		return nil, err
	}

	var comments []models.Comment
	for _, doc := range docs {
		comments = append(comments, mapper.ToComment(doc))
	}
	return comments, nil
}
