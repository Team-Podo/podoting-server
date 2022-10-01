package performance_find

import (
	"github.com/Team-Podo/podoting-server/repository"
)

type Performance struct {
	ID          uint       `json:"id"`
	Title       string     `json:"title"`
	ThumbUrl    *string    `json:"thumbUrl"`
	Place       *Place     `json:"place"`
	RunningTime string     `json:"runningTime"`
	StartDate   string     `json:"startDate"`
	EndDate     string     `json:"endDate"`
	Rating      string     `json:"rating"`
	Schedules   []Schedule `json:"schedules"`
	CreatedAt   string     `json:"createdAt"`
	UpdatedAt   string     `json:"updatedAt"`
}

type Place struct {
	ID        uint    `json:"id"`
	Name      string  `json:"name"`
	Address   *string `json:"address"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

type Schedule struct {
	UUID string `json:"uuid"`
	Memo string `json:"memo"`
	Date string `json:"date"`
	Time string `json:"time"`
}

func ParseResponseForm(p *repository.Performance) Performance {
	return Performance{
		ID:          p.ID,
		Place:       getPlace(p.Place),
		Title:       p.Title,
		ThumbUrl:    getThumbUrl(p.Thumbnail),
		RunningTime: p.RunningTime,
		StartDate:   p.StartDate,
		EndDate:     p.EndDate,
		Rating:      p.Rating,
		Schedules:   getSchedules(p.Schedules),
		CreatedAt:   p.CreatedAt.String(),
		UpdatedAt:   p.UpdatedAt.String(),
	}
}

func getPlace(p *repository.Place) *Place {
	if p == nil {
		return nil
	}

	var address *string

	if p.Location != nil {
		address = &p.Location.Name
	}

	return &Place{
		ID:        p.ID,
		Name:      p.Name,
		Address:   address,
		CreatedAt: p.CreatedAt.String(),
		UpdatedAt: p.UpdatedAt.String(),
	}
}

func getThumbUrl(f *repository.File) *string {
	if f == nil {
		return nil
	}

	thumbUrl := f.FullPath()
	return &thumbUrl
}

func getSchedules(s []repository.Schedule) []Schedule {
	var schedules []Schedule

	for _, v := range s {
		schedules = append(schedules, Schedule{
			UUID: v.UUID,
			Memo: v.Memo,
			Date: v.Date,
		})
	}

	return schedules
}
