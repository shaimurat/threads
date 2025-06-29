package mapper

import (
	postgreModels "twiteer/internal/data/postgres/models"
	"twiteer/internal/domain/models"
)

func ToTret(t postgreModels.TretDoc) models.Tret {
	var user models.User
	if t.User.ID != 0 {
		user = ToUser(t.User)
	}
	return models.Tret{
		ID:          t.ID,
		Theme:       t.Theme,
		Description: t.Description,
		Likes:       t.Likes,
		Dislikes:    t.Dislikes,
		UpdatedAt:   t.UpdatedAt,
		CreatedAt:   t.CreatedAt,
		UserID:      t.UserID,
		User:        user,
	}
}

func FromTret(t models.Tret) postgreModels.TretDoc {
	return postgreModels.TretDoc{
		ID:          t.ID,
		Theme:       t.Theme,
		Description: t.Description,
		Likes:       t.Likes,
		Dislikes:    t.Dislikes,
		UpdatedAt:   t.UpdatedAt,
		CreatedAt:   t.CreatedAt,
		UserID:      t.UserID,
		User:        FromUser(t.User), // Assumes User is not nil
	}
}
