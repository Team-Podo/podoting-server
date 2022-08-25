package models

type Musical struct {
	Id          uint              `json:"id"`
	Title       string            `json:"title"`
	ThumbUrl    string            `json:"thumbUrl"`
	RunningTime string            `json:"runningTime"`
	StartDate   string            `json:"startDate"`
	EndDate     string            `json:"endDate"`
	Schedules   []MusicalSchedule `json:"schedules"`
	Cast        []struct {
		Id      int `json:"id"`
		Profile struct {
			Url string `json:"url"`
		} `json:"profile"`
		Name string `json:"name"`
		Role string `json:"role"`
	} `json:"cast"`
	Contents []struct {
		Uuid    string `json:"uuid"`
		Title   string `json:"title"`
		Content string `json:"content"`
	} `json:"contents"`
}

type MusicalSchedule struct {
	UUID string `json:"uuid"`
	Date string `json:"date"`
	Time string `json:"time"`
	Cast []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"cast"`
}
