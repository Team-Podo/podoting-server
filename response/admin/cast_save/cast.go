package cast_save

import "github.com/Team-Podo/podoting-server/repository"

type Cast struct {
	ID          uint `json:"id"`
	PersonID    uint `json:"personID"`
	CharacterID uint `json:"characterID"`
	CreatedAt   string
	UpdatedAt   string
}

func ParseResponseForm(casts []repository.Cast) []Cast {
	var response []Cast
	for _, cast := range casts {
		response = append(response, Cast{
			ID:          cast.ID,
			PersonID:    cast.PersonID,
			CharacterID: cast.CharacterID,
			CreatedAt:   cast.CreatedAt.String(),
			UpdatedAt:   cast.UpdatedAt.String(),
		})
	}

	return response
}
