package cast_find

import "github.com/Team-Podo/podoting-server/repository"

type Cast struct {
	ID        uint       `json:"id"`
	Profile   string     `json:"profile"`
	Character *Character `json:"character"`
	Person    *Person    `json:"person"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
}

type Character struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type Person struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Birth     string `json:"birth"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func ParseResponseForm(cast *repository.Cast) Cast {
	return Cast{
		ID:        cast.ID,
		Profile:   cast.ProfileImage.FullPath(),
		Character: getCharacter(cast.Character),
		Person:    getPerson(cast.Person),
		CreatedAt: cast.CreatedAt.String(),
		UpdatedAt: cast.UpdatedAt.String(),
	}
}

func getCharacter(c *repository.Character) *Character {
	return &Character{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt.String(),
		UpdatedAt: c.UpdatedAt.String(),
	}
}

func getPerson(c *repository.Person) *Person {
	return &Person{
		ID:        c.ID,
		Name:      c.Name,
		Birth:     c.Birth.Format("2006-01-02"),
		CreatedAt: c.CreatedAt.String(),
		UpdatedAt: c.UpdatedAt.String(),
	}
}
