package models

type Musical struct {
	Id          uint              `json:"id"`
	Title       string            `json:"title"`
	ThumbUrl    string            `json:"thumbUrl"`
	RunningTime string            `json:"runningTime"`
	StartDate   string            `json:"startDate"`
	EndDate     string            `json:"endDate"`
	Rating      string            `json:"rating"`
	Place       *MusicalPlace     `json:"place"`
	Schedules   []MusicalSchedule `json:"schedules"`
	Cast        []Cast            `json:"cast"`
	Contents    []MusicalContent  `json:"contents"`
}

type MusicalPlace struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type MusicalSchedule struct {
	UUID string                `json:"uuid"`
	Date string                `json:"date"`
	Time string                `json:"time"`
	Cast []MusicalScheduleCast `json:"cast"`
}

type MusicalScheduleCast struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Cast struct {
	Id      uint    `json:"id"`
	Profile Profile `json:"profile"`
	Name    string  `json:"name"`
	Role    string  `json:"role"`
}

type Profile struct {
	Url string `json:"url"`
}

type MusicalContent struct {
	Uuid    string `json:"uuid"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
