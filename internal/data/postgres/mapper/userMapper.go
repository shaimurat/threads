package mapper

import (
	postgreModels "twiteer/internal/data/postgres/models"
	"twiteer/internal/domain/models"
)

func ToUser(u postgreModels.UserDoc) models.User {
	return models.User{
		ID:        u.ID,
		Username:  u.Username,
		Name:      u.Name,
		Surname:   u.Surname,
		Email:     u.Email,
		Password:  u.Password,
		AvatarUrl: u.AvatarUrl,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func FromUser(u models.User) postgreModels.UserDoc {
	return postgreModels.UserDoc{
		ID:        u.ID,
		Username:  u.Username,
		Name:      u.Name,
		Surname:   u.Surname,
		Email:     u.Email,
		Password:  u.Password,
		AvatarUrl: u.AvatarUrl,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
