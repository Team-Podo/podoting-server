package repository

type ScheduleCast struct {
	Schedule     *Schedule `json:"schedule" gorm:"foreignkey:ScheduleUUID"`
	ScheduleUUID string    `json:"-"`
	Cast         *Cast     `json:"cast" gorm:"foreignkey:CastID"`
	CastID       uint      `json:"-"`
}
