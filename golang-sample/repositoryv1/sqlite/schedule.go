package sqlite

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/lovemew67/public-misc/golang-sample/domainv1"
	"github.com/lovemew67/public-misc/golang-sample/repositoryv1"
)

var (
	_ repositoryv1.ScheduleV1Repository = &ScheduleV1SQLiteRepositorier{}
)

type ScheduleV1SQLiteRepositorier struct{}

func (s *ScheduleV1SQLiteRepositorier) CreateSchedule(schedule *domainv1.Schedule) (result *domainv1.Schedule, err error) {
	id := uuid.New().String()
	schedule.ID = id
	now := time.Now()
	schedule.CreatedAt = now
	schedule.UpdatedAt = now
	db := sqlitedb
	db = db.Create(schedule)
	err = db.Error
	if err != nil {
		return
	}

	result, err = s.GetSchedule(id)
	return
}

func (s *ScheduleV1SQLiteRepositorier) CountTotalSchedules() (result int, err error) {
	db := sqlitedb
	db = db.Model(domainv1.Schedule{})
	db = db.Count(&result)
	err = db.Error
	return
}

func (s *ScheduleV1SQLiteRepositorier) GetSchedule(id string) (schedule *domainv1.Schedule, err error) {
	scheduleList := make([]*domainv1.Schedule, 1)
	db := sqlitedb
	db = db.Where("id = ?", id)
	db = db.Limit(1)
	db = db.Find(&scheduleList)
	err = db.Error
	if len(scheduleList) != 0 {
		schedule = scheduleList[0]
	} else {
		err = fmt.Errorf("not found")
	}
	return
}

func (s *ScheduleV1SQLiteRepositorier) QueryAllSchedulesWithOffsetAndLimit(offset, limit int) (scheduleList []*domainv1.Schedule, err error) {
	scheduleList = make([]*domainv1.Schedule, limit)
	db := sqlitedb
	db = db.Offset(offset)
	db = db.Limit(limit)
	db = db.Find(&scheduleList)
	err = db.Error
	return
}

func (s *ScheduleV1SQLiteRepositorier) PatchSchedule(id string, schedule *domainv1.Schedule) (err error) {
	now := time.Now()
	schedule.UpdatedAt = now
	db := sqlitedb
	db = db.Model(domainv1.Schedule{ID: id})
	db = db.Update(schedule)
	err = db.Error
	return
}

func (s *ScheduleV1SQLiteRepositorier) DeleteSchedule(id string) (err error) {
	schedule := domainv1.Schedule{ID: id}
	db := sqlitedb
	db = db.Delete(domainv1.Schedule{}, &schedule)
	err = db.Error
	return
}

func NewScheduleV1SQLiteRepositorier(ctx cornerstone.Context) (result *ScheduleV1SQLiteRepositorier, err error) {
	funcName := "NewScheduleV1SQLiteRepositorier"

	dbFilePath := fmt.Sprintf(formatSqliteDatabasePath, dataFolder)
	if err = createDirIfNotExist(ctx, dataFolder); err != nil {
		return
	}
	if err = createFileIfNotExist(ctx, dbFilePath); err != nil {
		return
	}

	db, err := gorm.Open(dialect, dbFilePath)
	if err != nil {
		return
	}
	sqlitedb = db

	schedule := &domainv1.Schedule{}
	if hasTable := sqlitedb.HasTable(schedule); hasTable {
		cornerstone.Infof(ctx, "[%s] continue to reuse the table: %s", funcName, scheduleV1TableName)
		db.AutoMigrate(&domainv1.Schedule{})
		return
	}

	if err = sqlitedb.CreateTable(schedule).Error; err != nil {
		return
	}

	result = &ScheduleV1SQLiteRepositorier{}
	return
}
