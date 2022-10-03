package character_get

import "github.com/Team-Podo/podoting-server/repository"

type Character struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func ParseResponseForm(c []repository.Character) []Character {
	characters := make([]Character, len(c))

	for i, character := range c {
		characters[i] = Character{
			ID:        character.ID,
			Name:      character.Name,
			CreatedAt: character.CreatedAt.String(),
			UpdatedAt: character.UpdatedAt.String(),
		}
	}

	return characters
}
