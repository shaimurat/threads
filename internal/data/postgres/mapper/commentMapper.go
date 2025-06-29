package mapper

import (
	postgreModels "twiteer/internal/data/postgres/models"
	"twiteer/internal/domain/models"
)

func ToComment(c postgreModels.CommentDoc) models.Comment {
	return models.Comment{
		ID:        c.ID,
		Text:      c.Text,
		Likes:     c.Likes,
		Dislikes:  c.Dislikes,
		UpdatedAt: c.UpdatedAt,
		CreatedAt: c.CreatedAt,
		UserID:    c.UserID,
		TretID:    c.TretID,
		User:      ToUser(c.User),
		Tret:      ToTret(c.Tret),
	}
}

func FromComment(c models.Comment) postgreModels.CommentDoc {
	return postgreModels.CommentDoc{
		ID:        c.ID,
		Text:      c.Text,
		Likes:     c.Likes,
		Dislikes:  c.Dislikes,
		UpdatedAt: c.UpdatedAt,
		CreatedAt: c.CreatedAt,
		UserID:    c.UserID,
		TretID:    c.TretID,
		User:      FromUser(c.User),
		Tret:      FromTret(c.Tret),
	}
}
