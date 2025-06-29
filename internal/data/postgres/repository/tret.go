package repository

import (
	"gorm.io/gorm"
	"strconv"
	"twiteer/internal/data/postgres/mapper"
	postgreModels "twiteer/internal/data/postgres/models"
	"twiteer/internal/domain/models"
)

type TretRepoSQL struct {
	db *gorm.DB
}

func NewTretRepoSQL(db *gorm.DB) *TretRepoSQL {
	return &TretRepoSQL{db: db}
}

func (r *TretRepoSQL) Create(tret models.Tret) error {
	doc := mapper.FromTret(tret)
	result := r.db.Create(&doc)
	return result.Error
}

func (r *TretRepoSQL) Update(tret models.Tret) error {
	doc := mapper.FromTret(tret)
	result := r.db.Save(&doc)
	return result.Error
}

func (r *TretRepoSQL) Delete(id string) error {
	tretID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	result := r.db.Delete(&postgreModels.TretDoc{}, tretID)
	return result.Error
}

func (r *TretRepoSQL) GetTret(id string) (models.Tret, error) {
	tretID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return models.Tret{}, err
	}

	var doc postgreModels.TretDoc
	result := r.db.Preload("User").First(&doc, tretID)
	if result.Error != nil {
		return models.Tret{}, result.Error
	}
	return mapper.ToTret(doc), nil
}

func (r *TretRepoSQL) GetUserTrets(id string) ([]models.Tret, error) {
	userID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return nil, err
	}

	var docs []postgreModels.TretDoc
	result := r.db.Preload("User").Where("user_id = ?", userID).Find(&docs)
	if result.Error != nil {
		return nil, result.Error
	}

	var trets []models.Tret
	for _, doc := range docs {
		trets = append(trets, mapper.ToTret(doc))
	}
	return trets, nil
}

func (r *TretRepoSQL) LikeTret(id string) error {
	tretID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	return r.db.Model(&postgreModels.TretDoc{}).Where("id = ?", tretID).
		UpdateColumn("likes", gorm.Expr("likes + ?", 1)).Error
}

func (r *TretRepoSQL) DislikeTret(id string) error {
	tretID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}
	return r.db.Model(&postgreModels.TretDoc{}).Where("id = ?", tretID).
		UpdateColumn("dislikes", gorm.Expr("dislikes + ?", 1)).Error
}
