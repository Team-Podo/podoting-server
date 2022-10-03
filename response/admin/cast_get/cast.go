package cast_get

import "github.com/Team-Podo/podoting-server/repository"

type Cast struct {
	ID            uint    `json:"id"`
	ProfileImage  *string `json:"profileImage"`
	CharacterName *string `json:"characterName"`
	PersonName    *string `json:"personName"`
	CreatedAt     string  `json:"createdAt"`
	UpdatedAt     string  `json:"updatedAt"`
}

func ParseResponseForm(casts []repository.Cast) []Cast {
	var response []Cast
	for _, cast := range casts {
		response = append(response, Cast{
			ID:            cast.ID,
			ProfileImage:  getProfile(cast.ProfileImage),
			CharacterName: getCharacterName(cast.Character),
			PersonName:    getPersonName(cast.Person),
			CreatedAt:     cast.CreatedAt.String(),
			UpdatedAt:     cast.UpdatedAt.String(),
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

func getCharacterName(c *repository.Character) *string {
	if c == nil {
		return nil
	}

	return &c.Name
}

func getPersonName(c *repository.Person) *string {
	if c == nil {
		return nil
	}

	return &c.Name
}
