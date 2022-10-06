package performance_get

type Performance struct {
	ID          uint        `json:"id"`
	Title       string      `json:"title"`
	MainAreaID  *uint       `json:"mainAreaID"`
	ThumbUrl    *string     `json:"thumbUrl"`
	RunningTime string      `json:"runningTime"`
	StartDate   string      `json:"startDate"`
	EndDate     string      `json:"endDate"`
	Rating      string      `json:"rating"`
	Schedules   interface{} `json:"schedules"`
	CreatedAt   string      `json:"createdAt"`
	UpdatedAt   string      `json:"updatedAt"`
}

type Product struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	File      *File  `json:"file"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type File struct {
	ID        uint   `json:"id"`
	Size      int64  `json:"Size"`
	Path      string `json:"Path"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
