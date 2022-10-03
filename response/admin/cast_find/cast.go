package cast_find

import "github.com/Team-Podo/podoting-server/repository"

type Cast struct {
	ID           uint       `json:"id"`
	ProfileImage *string    `json:"profileImage"`
	Character    *Character `json:"character"`
	Person       *Person    `json:"person"`
	CreatedAt    string     `json:"createdAt"`
	UpdatedAt    string     `json:"updatedAt"`
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
		ID:           cast.ID,
		ProfileImage: getProfileImage(cast.ProfileImage),
		Character:    getCharacter(cast.Character),
		Person:       getPerson(cast.Person),
		CreatedAt:    cast.CreatedAt.String(),
		UpdatedAt:    cast.UpdatedAt.String(),
	}
}

func getProfileImage(f *repository.File) *string {
	if f == nil {
		return nil
	}

	url := f.FullPath()

	return &url
}

func getCharacter(c *repository.Character) *Character {
	if c == nil {
		return nil
	}

	return &Character{
		ID:        c.ID,
		Name:      c.Name,
		CreatedAt: c.CreatedAt.String(),
		UpdatedAt: c.UpdatedAt.String(),
	}
}

func getPerson(c *repository.Person) *Person {
	if c == nil {
		return nil
	}

	return &Person{
		ID:        c.ID,
		Name:      c.Name,
		Birth:     c.Birth.Format("2006-01-02"),
		CreatedAt: c.CreatedAt.String(),
		UpdatedAt: c.UpdatedAt.String(),
	}
}
