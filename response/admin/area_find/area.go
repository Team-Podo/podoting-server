package area_find

import "github.com/Team-Podo/podoting-server/repository"

type Area struct {
	ID                 uint   `json:"id"`
	Name               string `json:"name"`
	BackgroundImageUrl string `json:"backgroundImageUrl"`
	CreatedAt          string `json:"createdAt"`
	UpdatedAt          string `json:"updatedAt"`
}

func ParseResponseForm(area *repository.Area) Area {
	return Area{
		ID:                 area.ID,
		Name:               area.Name,
		BackgroundImageUrl: getBackgroundImageUrl(area.BackgroundImage),
		CreatedAt:          area.CreatedAt.String(),
		UpdatedAt:          area.UpdatedAt.String(),
	}
}

func getBackgroundImageUrl(f *repository.File) string {
	if f == nil {
		return ""
	}

	url := f.FullPath()

	return url
}
