package res

import "github.com/Team-Podo/podoting-server/repository"

type MainPageResponse struct {
	Performances []Performance `json:"performances"`
}

type Performance struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	ThumbUrl    *string `json:"thumbUrl"`
	RunningTime string  `json:"runningTime"`
	StartDate   string  `json:"startDate"`
	EndDate     string  `json:"endDate"`
	Rating      string  `json:"rating"`
	Place       *Place  `json:"place"`
}

type Place struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

func (m MainPageResponse) Of(performances []repository.Performance) MainPageResponse {
	var response MainPageResponse
	response.Performances = make([]Performance, len(performances))
	for i := 0; i < len(performances); i++ {
		p := performances[i]
		response.Performances[i] = Performance{
			ID:          p.ID,
			Title:       p.Title,
			ThumbUrl:    p.GetThumbnailURL(),
			RunningTime: p.RunningTime,
			StartDate:   p.StartDate,
			EndDate:     p.EndDate,
			Rating:      p.Rating,
			Place:       getPlace(p.Place),
		}
	}

	return response
}

func getPlace(place *repository.Place) *Place {
	if place == nil {
		return nil
	}

	return &Place{
		ID:    place.ID,
		Name:  place.Name,
		Image: place.Name,
	}
}
