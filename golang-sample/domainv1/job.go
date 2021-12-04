package domainv1

import (
	"encoding/json"
	"time"
)

// unlike mongo, sqlite index shared accross tables

type Job struct {
	ID           string `json:"id" gorm:"column:id;primary_key"`
	CID          string `json:"cid"           gorm:"column:cid"`
	InternalData string `json:"internalData"  gorm:"column:internal_data"`
	RetryCount   int    `json:"retryCount"    gorm:"column:retry_count;index:retry_count"`
	Status       int    `json:"status"        gorm:"column:status;index:status"`

	Type       JobType          `json:"type" gorm:"column:type;index:type"`
	Processing JobProcessStatus `json:"processing" grom:"column:processing;index:processing"`

	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at"`
	// UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;index:undated_at"`
}

type JobInternalData struct {
	FailedReasons map[int]string         `json:"failedReasons"`
	Metadata      map[string]interface{} `json:"metadata"`
}

func NewJob(maxRetry int, jobType JobType, metadata map[string]interface{}) *Job {
	now := time.Now().UTC()
	internalData := JobInternalData{
		FailedReasons: map[int]string{},
	}
	if metadata != nil {
		internalData.Metadata = metadata
	} else {
		internalData.Metadata = map[string]interface{}{}
	}
	internalDataBytes, _ := json.Marshal(internalData)
	return &Job{
		InternalData: string(internalDataBytes),
		RetryCount:   maxRetry,
		Status:       JobGeneralStatusReady.ToInt(),

		Type:       jobType,
		Processing: JobProcessStatusStopped,

		CreatedAt: now,
		UpdatedAt: now,
	}
}
