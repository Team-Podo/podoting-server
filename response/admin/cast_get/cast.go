package cast_get

import (
	"github.com/Team-Podo/podoting-server/repository"
)

type Cast struct {
	ID           uint    `json:"id"`
	ProfileImage *string `json:"profileImage"`
	CharacterID  *uint   `json:"characterID"`
	PersonID     *uint   `json:"personID"`
	CreatedAt    string  `json:"createdAt"`
	UpdatedAt    string  `json:"updatedAt"`
}

func ParseResponseForm(casts []repository.Cast) []Cast {
	var response []Cast
	for _, cast := range casts {
		response = append(response, Cast{
			ID:           cast.ID,
			ProfileImage: getProfile(cast.ProfileImage),
			CharacterID:  getCharacterID(cast.Character),
			PersonID:     getPersonID(cast.Person),
			CreatedAt:    cast.CreatedAt.String(),
			UpdatedAt:    cast.UpdatedAt.String(),
		})
	}

	return response
}

func getProfile(c *repository.File) *string {
	if c == nil {
		return nil
	}

	url := c.FullPath()

	return &url
}

func getCharacterID(c *repository.Character) *uint {
	if c == nil {
		return nil
	}

	return &c.ID
}

func getPersonID(c *repository.Person) *uint {
	if c == nil {
		return nil
	}

	return &c.ID
}
