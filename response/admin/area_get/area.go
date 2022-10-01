package area_get

import (
	"github.com/Team-Podo/podoting-server/repository"
)

type Area struct {
	ID                 uint    `json:"id"`
	Name               string  `json:"name"`
	BackgroundImageUrl *string `json:"backgroundImageUrl"`
	CreatedAt          string  `json:"createdAt"`
	UpdatedAt          string  `json:"updatedAt"`
}

func ParseResponseForm(area []repository.Area) []Area {
	areas := make([]Area, len(area))
	for i, a := range area {
		areas[i] = Area{
			ID:                 a.ID,
			Name:               a.Name,
			BackgroundImageUrl: getBackgroundImageUrl(a.BackgroundImage),
			CreatedAt:          a.CreatedAt.String(),
			UpdatedAt:          a.UpdatedAt.String(),
		}
	}

	return areas
}

func getBackgroundImageUrl(f *repository.File) *string {
	if f == nil {
		return nil
	}

	url := f.FullPath()

	return &url
}
