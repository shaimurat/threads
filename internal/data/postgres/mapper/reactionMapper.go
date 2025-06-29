package mapper

import (
	postgreModels "twiteer/internal/data/postgres/models"
	"twiteer/internal/domain/models"
)

func ToReaction(r postgreModels.ReactionDoc) models.Reaction {
	return models.Reaction{
		IsLike: r.IsLike,
		UserID: r.UserID,
		TretID: r.TretID,
		User:   ToUser(r.User),
		Tret:   ToTret(r.Tret),
	}
}

func FromReaction(r models.Reaction) postgreModels.ReactionDoc {
	return postgreModels.ReactionDoc{
		IsLike: r.IsLike,
		UserID: r.UserID,
		TretID: r.TretID,
		User:   FromUser(r.User),
		Tret:   FromTret(r.Tret),
	}
}
