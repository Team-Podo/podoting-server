package person_find

import "github.com/Team-Podo/podoting-server/repository"

type Person struct {
	ID           uint    `json:"id"`
	Name         string  `json:"name"`
	Birth        *string `json:"birth"`
	ProfileImage *string `json:"profileImage"`
	CreatedAt    string  `json:"createdAt"`
	UpdatedAt    string  `json:"updatedAt"`
}

func ParseResponseForm(p *repository.Person) *Person {
	return &Person{
		ID:           p.ID,
		Name:         p.Name,
		Birth:        getBirth(p),
		ProfileImage: getProfileImage(p.ProfileImage),
		CreatedAt:    p.CreatedAt.String(),
		UpdatedAt:    p.UpdatedAt.String(),
	}
}

func getProfileImage(p *repository.File) *string {
	if p == nil {
		return nil
	}

	profileImage := p.FullPath()

	return &profileImage
}

func getBirth(p *repository.Person) *string {
	if p.Birth == nil {
		return nil
	}

	birth := p.Birth.Format("2006-01-02")

	return &birth
}
