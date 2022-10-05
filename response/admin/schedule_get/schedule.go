package schedule_get

import (
	"database/sql"
	"github.com/Team-Podo/podoting-server/repository"
)

type Schedule struct {
	UUID      string  `json:"uuid"`
	Memo      string  `json:"memo"`
	Open      bool    `json:"open"`
	Date      string  `json:"date"`
	Time      *string `json:"time"`
	Casts     []Cast  `json:"casts"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt string  `json:"updatedAt"`
}

type Cast struct {
	ID            uint    `json:"id"`
	CharacterName *string `json:"characterName"`
	PersonName    *string `json:"personName"`
	ProfileImage  *string `json:"profileImage"`
}

func ParseResponseFrom(schedules []repository.Schedule) []Schedule {
	var response []Schedule

	for _, schedule := range schedules {
		response = append(response, Schedule{
			UUID:      schedule.UUID,
			Memo:      schedule.Memo,
			Open:      schedule.Open,
			Date:      schedule.Date,
			Time:      getTime(schedule.Time),
			Casts:     getCasts(schedule.Casts),
			CreatedAt: schedule.CreatedAt.String(),
			UpdatedAt: schedule.UpdatedAt.String(),
		})
	}

	return response
}

func getTime(time sql.NullString) *string {
	if !time.Valid {
		return nil
	}

	return &time.String
}

func getCasts(casts []repository.Cast) []Cast {
	var response []Cast

	for _, cast := range casts {
		response = append(response, Cast{
			ID:            cast.ID,
			CharacterName: getCharacterName(cast.Character),
			PersonName:    getPersonName(cast.Person),
			ProfileImage:  getProfileImage(cast.ProfileImage),
		})
	}

	return response
}

func getCharacterName(c *repository.Character) *string {
	if c == nil {
		return nil
	}

	return &c.Name
}

func getPersonName(p *repository.Person) *string {
	if p == nil {
		return nil
	}

	return &p.Name
}

func getProfileImage(file *repository.File) *string {
	if file == nil {
		return nil
	}

	fullPath := file.FullPath()

	return &fullPath
}
