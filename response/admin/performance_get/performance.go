package performance_get

import "time"

type Performance struct {
	ID          uint        `json:"id"`
	RunningTime string      `json:"runningTime"`
	EndDate     string      `json:"endDate"`
	Product     *Product    `json:"product"`
	Title       string      `json:"title"`
	StartDate   string      `json:"startDate"`
	Rating      string      `json:"rating"`
	Schedules   interface{} `json:"schedules"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

type Product struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	File      *File     `json:"file"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type File struct {
	ID        uint      `json:"id"`
	Size      int64     `json:"Size"`
	Path      string    `json:"Path"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
