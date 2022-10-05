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
	suite.scheduleRepository = &ScheduleRepository{DB: database.Gorm}
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
	_ = suite.scheduleRepository.Save(&suite.schedule)

	suite.scheduleRepository.Find(suite.schedule.UUID)
}

func (suite *ScheduleTestSuite) TestFindScheduleByUUID() {
	_ = suite.scheduleRepository.Save(&suite.schedule)

	_schedule := suite.scheduleRepository.Find(suite.schedule.UUID)

	assert.Equal(suite.T(), suite.schedule.UUID, _schedule.UUID)
}

func (suite *ScheduleTestSuite) TestUpdateSchedule() {
	_ = suite.scheduleRepository.Save(&suite.schedule)
	assert.Equal(suite.T(), "test", suite.schedule.Memo)
	fmt.Println(suite.schedule)

	var _schedule Schedule
	_schedule.UUID = suite.schedule.UUID
	_schedule.Memo = "수정 후"

	_ = suite.scheduleRepository.Update(&_schedule)
	assert.Equal(suite.T(), "수정 후", suite.schedule.Memo)
}

func (suite *ScheduleTestSuite) TestDeleteScheduleByUUID() {
	suite.scheduleRepository.Delete(suite.schedule.UUID)

	_schedule := suite.scheduleRepository.Find(suite.schedule.UUID)

	assert.Nil(suite.T(), _schedule)
}

func TestSaveScheduleTestSuite(t *testing.T) {
	suite.Run(t, new(ScheduleTestSuite))
}
