package repository

import (
	"database/sql"
	"fmt"
	"github.com/Team-Podo/podoting-server/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ScheduleTestSuite struct {
	suite.Suite
	scheduleRepository *ScheduleRepository
	schedule           Schedule
}

func (suite *ScheduleTestSuite) SetupTest() {
	suite.scheduleRepository = &ScheduleRepository{Db: database.Gorm}
	newUUID, _ := uuid.NewUUID()
	suite.schedule = Schedule{
		UUID:        newUUID.String(),
		Performance: &Performance{ID: 1},
		Memo:        "test",
		Date:        "2022-08-02",
		Time:        sql.NullString{Valid: false},
	}
}

func (suite *ScheduleTestSuite) TestSaveSchedule() {
	schedule := suite.scheduleRepository.Save(&Schedule{
		Performance: &Performance{ID: 1},
		Memo:        "test",
		Date:        "2022-08-02",
		Time:        sql.NullString{Valid: false},
	})

	suite.scheduleRepository.Find(schedule.GetUUID())
}

func (suite *ScheduleTestSuite) TestGet() {
	schedules := suite.scheduleRepository.Get(map[string]any{})

	assert.NotNil(suite.T(), schedules)
}

func (suite *ScheduleTestSuite) TestGetScheduleByUUID() {
	schedule := suite.scheduleRepository.Save(&Schedule{
		Performance: &Performance{ID: 1},
		Memo:        "test",
		Date:        "2022-08-02",
		Time:        sql.NullString{Valid: false},
	})

	_schedule := suite.scheduleRepository.Find(schedule.GetUUID())

	assert.Equal(suite.T(), schedule.GetUUID(), _schedule.GetUUID())
}

func (suite *ScheduleTestSuite) TestUpdateSchedule() {
	schedule := suite.scheduleRepository.Save(&suite.schedule)
	assert.Equal(suite.T(), "test", schedule.GetMemo())
	fmt.Println(schedule)

	var _schedule Schedule
	_schedule.UUID = schedule.GetUUID()
	_schedule.Memo = "수정 후"

	schedule = suite.scheduleRepository.Update(&_schedule)
	assert.Equal(suite.T(), "수정 후", schedule.GetMemo())
	fmt.Println(schedule)
}

func (suite *ScheduleTestSuite) TestDeleteScheduleByUUID() {
	suite.scheduleRepository.Delete(suite.schedule.GetUUID())

	_schedule := suite.scheduleRepository.Find(suite.schedule.GetUUID())

	assert.Nil(suite.T(), _schedule)
}

func TestSaveScheduleTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleTestSuite))
}
