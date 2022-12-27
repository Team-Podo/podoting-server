package response

import "github.com/Team-Podo/podoting-server/repository"

type CreateManyCast struct {
	ID          uint   `json:"id"`
	PersonID    uint   `json:"personID"`
	CharacterID uint   `json:"characterID"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

func GetCreateManyResponse(casts []repository.Cast) []CreateManyCast {
	var response []CreateManyCast
	for _, cast := range casts {
		response = append(response, CreateManyCast{
			ID:          cast.ID,
			PersonID:    cast.PersonID,
			CharacterID: cast.CharacterID,
			CreatedAt:   cast.CreatedAt.String(),
			UpdatedAt:   cast.UpdatedAt.String(),
		})
	}

	return response
}
