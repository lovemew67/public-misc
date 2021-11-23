package sqlite

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/lovemew67/public-misc/golang-sample/domainv1"
	"github.com/lovemew67/public-misc/golang-sample/gen/go/proto"
	"github.com/spf13/viper"
)

type JobV1SQLiteRepositorier struct{}

// from scheduler or http handler
func Insert(job domainv1.Job) error {
	db := sqlitedb
	db = db.Create(&job)
	err := db.Error

	return err
}

// for dispatcher
func QueryReadyTask() ([]domainv1.Job, error) {
	tasks := make([]domainv1.Job, 1)

	queryDelay := viper.GetInt("app.retry.delay")
	maxRetryCount := viper.GetInt("app.retry.max_count")

	query := fmt.Sprintf("status >= %d AND status < %d AND processing = %d AND ((retry_count = %d) OR (retry_count < %d AND updated_at <= datetime('now', 'localtime', '-%d minutes')))",
		domainv1.JobGeneralStatusReady,
		domainv1.JobGeneralStatusFinished,
		domainv1.JobProcessStatusStopped,
		maxRetryCount,
		maxRetryCount,
		queryDelay,
	)

	db := sqlitedb
	db = db.Where(query)
	db = db.Limit(1)
	db = db.Find(&tasks)
	err := db.Error

	return tasks, err
}

func UpdateProcessStatusToOngoing(id int) error {
	db := sqlitedb
	db = db.Model(domainv1.Job{})
	db = db.Where("id = ?", id)
	db = db.Updates(domainv1.Job{
		Processing: domainv1.JobProcessStatusOngoing,
	})
	err := db.Error

	return err
}

func CancelTaskByID(id int) error {
	db := sqlitedb
	db = db.Model(domainv1.Job{})
	db = db.Where("id = ?", id)
	db = db.Updates(domainv1.Job{
		Status:     domainv1.JobGeneralStatusCancelled.ToInt(),
		Processing: domainv1.JobProcessStatusStopped,
	})
	err := db.Error

	return err
}

// for task handler
func RemoveFromTaskQueue(task *domainv1.Job) error {
	db := sqlitedb
	db = db.Model(domainv1.Job{})
	db = db.Where("id = ?", task.ID)
	db = db.Updates(domainv1.Job{
		Processing:   domainv1.JobProcessStatusStopped,
		RetryCount:   task.RetryCount - 1,
		InternalData: task.InternalData,
	})
	err := db.Error

	return err
}

func UpdateTaskStatusStillOngoing(task *domainv1.Job) error {
	db := sqlitedb
	db = db.Model(domainv1.Job{})
	db = db.Where("id = ?", task.ID)
	db = db.Updates(domainv1.Job{
		Status:       task.Status,
		InternalData: task.InternalData,
		Processing:   domainv1.JobProcessStatusOngoing,
	})
	err := db.Error

	return err
}

func UpdateTaskStatusToStopped(task *domainv1.Job) error {
	db := sqlitedb
	db = db.Model(domainv1.Job{})
	db = db.Where("id = ?", task.ID)
	db = db.Updates(domainv1.Job{
		Status:       task.Status,
		InternalData: task.InternalData,
		Processing:   domainv1.JobProcessStatusStopped,
	})
	err := db.Error

	return err
}

func NewJobV1SQLiteRepositorier(ctx cornerstone.Context) (result *JobV1SQLiteRepositorier, err error) {
	funcName := "NewJobV1SQLiteRepositorier"

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

	task := &proto.StaffV1{}
	if hasTable := sqlitedb.HasTable(task); hasTable {
		cornerstone.Infof(ctx, "[%s] continue to reuse the table: %s", funcName, jobV1TableName)
		db.AutoMigrate(&proto.StaffV1{})
		return
	}

	if err = sqlitedb.CreateTable(task).Error; err != nil {
		return
	}

	result = &JobV1SQLiteRepositorier{}
	return
}
