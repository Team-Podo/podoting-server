package repository

import (
	"github.com/Team-Podo/podoting-server/query/user_seats"
	"gorm.io/gorm"
	"log"
	"time"
)

type Seat struct {
	UUID              string           `json:"uuid" gorm:"primarykey"`
	AreaBoilerplate   *AreaBoilerplate `json:"areaBoilerplate" gorm:"foreignKey:AreaBoilerplateID"`
	AreaBoilerplateID *uint            `json:"-"`
	Performance       *Performance     `json:"performance" gorm:"foreignKey:PerformanceID"`
	PerformanceID     uint             `json:"-"`
	Grade             *SeatGrade       `json:"seatGrade" gorm:"foreignkey:SeatGradeID"`
	SeatGradeID       uint             `json:"-"`
	Bookings          []SeatBooking    `json:"booking" gorm:"foreignKey:SeatUUID"`
	CreatedAt         time.Time        `json:"createdAt"`
	UpdatedAt         time.Time        `json:"updatedAt"`
	DeletedAt         *gorm.DeletedAt  `json:"-" gorm:"index"`
}

type SeatRepository struct {
	DB *gorm.DB
}

func (s *SeatRepository) GetByAreaAndPerformanceID(areaID uint, performanceID uint) []Seat {
	var seats []Seat

	err := s.DB.
		Joins("Grade").
		Preload("AreaBoilerplate", "area_id = ?", areaID).
		Preload("AreaBoilerplate.Point").
		Where(&Seat{PerformanceID: performanceID}).
		Order("area_boilerplate_id").
		Find(&seats).
		Error

	if err != nil {
		return nil
	}

	return seats
}

func (s *SeatRepository) GetByScheduleUUID(scheduleUUID string) []Seat {
	var seats []Seat
	err := s.DB.Preload("Grade").Where("schedule_uuid = ?", scheduleUUID).Find(&seats).Error

	if err != nil {
		return nil
	}

	return nil
}

func (s *SeatRepository) GetByUUID(uuid string) *Seat {
	var seat Seat
	err := s.DB.Preload("Grade").Where("uuid = ?", uuid).Find(&seat).Error

	if err != nil {
		return nil
	}

	return &seat

}

func (s *SeatRepository) GetSeatsByAreaIdAndScheduleUUID(areaId uint, scheduleUUID string) []Seat {
	var seatEntities []user_seats.Seat

	err := s.DB.
		Debug().
		Raw("select seats.uuid, ab.name `seat_name`, sg.name `grade_name`, sg.price, sg.color, p.x `point_x`, p.y `point_y`, count(sb.booked) `booked_count` from seats join area_boilerplates ab on seats.area_boilerplate_id = ab.id join seat_grades sg on seats.seat_grade_id = sg.id join points p on ab.point_id = p.id join performances perf on seats.performance_id = perf.id join schedules s on seats.performance_id = s.performance_id left join seat_bookings sb on seats.uuid = sb.seat_uuid and s.uuid = sb.schedule_uuid where ab.area_id = ? and s.uuid = ? and sb.booked = 0 group by seats.uuid", areaId, scheduleUUID).
		Scan(&seatEntities).Error

	if err != nil {
		log.Fatal(err.Error())
		return nil
	}

	var seats []Seat
	for i := range seatEntities {
		if seatEntities[i].BookedCount > 0 {
			continue
		}

		bookings := make([]SeatBooking, seatEntities[i].BookedCount)

		point := Point{
			X: seatEntities[i].PointX,
			Y: seatEntities[i].PointY,
		}

		areaBoilerplate := AreaBoilerplate{
			Name:  seatEntities[i].SeatName,
			Point: &point,
		}

		grade := SeatGrade{
			Name:  seatEntities[i].GradeName,
			Price: seatEntities[i].Price,
			Color: seatEntities[i].Color,
		}

		seats = append(seats, Seat{
			UUID:            seatEntities[i].UUID,
			AreaBoilerplate: &areaBoilerplate,
			Grade:           &grade,
			Bookings:        bookings,
		})
	}

	return seats
}

func (s *SeatRepository) SaveSeats(seats []Seat) error {
	return s.DB.
		Omit("area_boilerplate_id").
		Save(&seats).Error
}

func (s *SeatRepository) CreateSeats(seats []Seat) error {
	return s.DB.Debug().Create(&seats).Error
}
