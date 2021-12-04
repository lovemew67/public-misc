package sqlite

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"github.com/lovemew67/public-misc/cornerstone"
	"github.com/lovemew67/public-misc/golang-sample/domainv1"
	"github.com/lovemew67/public-misc/golang-sample/repositoryv1"
	"github.com/spf13/viper"
)

var (
	_ repositoryv1.JobV1Repository = &JobV1SQLiteRepositorier{}
)

type JobV1SQLiteRepositorier struct{}

func (s *JobV1SQLiteRepositorier) CreateJob(job *domainv1.Job) (result *domainv1.Job, err error) {
	id := uuid.New().String()
	job.ID = id
	now := time.Now()
	job.CreatedAt = now
	job.UpdatedAt = now
	db := sqlitedb
	db = db.Create(job)
	err = db.Error
	if err != nil {
		return
	}

	result, err = s.GetJob(id)
	return
}

func (s *JobV1SQLiteRepositorier) CountTotalJobs() (result int, err error) {
	db := sqlitedb
	db = db.Model(domainv1.Job{})
	db = db.Count(&result)
	err = db.Error
	return
}

func (s *JobV1SQLiteRepositorier) GetJob(id string) (job *domainv1.Job, err error) {
	jobList := make([]*domainv1.Job, 1)
	db := sqlitedb
	db = db.Where("id = ?", id)
	db = db.Limit(1)
	db = db.Find(&jobList)
	err = db.Error
	if len(jobList) != 0 {
		job = jobList[0]
	} else {
		err = fmt.Errorf("not found")
	}
	return
}

func (s *JobV1SQLiteRepositorier) QueryAllJobsWithOffsetAndLimit(offset, limit int) (jobList []*domainv1.Job, err error) {
	jobList = make([]*domainv1.Job, limit)
	db := sqlitedb
	db = db.Offset(offset)
	db = db.Limit(limit)
	db = db.Find(&jobList)
	err = db.Error
	return
}

func (s *JobV1SQLiteRepositorier) PatchJob(id string, job *domainv1.Job) (err error) {
	now := time.Now()
	job.UpdatedAt = now
	db := sqlitedb
	db = db.Model(domainv1.Job{ID: id})
	db = db.Update(job)
	err = db.Error
	return
}

func (s *JobV1SQLiteRepositorier) DeleteJob(id string) (err error) {
	job := domainv1.Job{ID: id}
	db := sqlitedb
	db = db.Delete(domainv1.Job{}, &job)
	err = db.Error
	return
}

// for dispatcher
func (s *JobV1SQLiteRepositorier) QueryReadyTask() ([]domainv1.Job, error) {
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

func (s *JobV1SQLiteRepositorier) UpdateProcessStatusToOngoing(id int) error {
	db := sqlitedb
	db = db.Model(domainv1.Job{})
	db = db.Where("id = ?", id)
	db = db.Updates(domainv1.Job{
		Processing: domainv1.JobProcessStatusOngoing,
	})
	err := db.Error

	return err
}

func (s *JobV1SQLiteRepositorier) CancelTaskByID(id int) error {
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
func (s *JobV1SQLiteRepositorier) RemoveFromTaskQueue(task *domainv1.Job) error {
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

func (s *JobV1SQLiteRepositorier) UpdateTaskStatusStillOngoing(task *domainv1.Job) error {
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

func (s *JobV1SQLiteRepositorier) UpdateTaskStatusToStopped(task *domainv1.Job) error {
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

	job := &domainv1.Job{}
	if hasTable := sqlitedb.HasTable(job); hasTable {
		cornerstone.Infof(ctx, "[%s] continue to reuse the table: %s", funcName, jobV1TableName)
		db.AutoMigrate(&domainv1.Job{})
		return
	}

	if err = sqlitedb.CreateTable(job).Error; err != nil {
		return
	}

	result = &JobV1SQLiteRepositorier{}
	return
}
