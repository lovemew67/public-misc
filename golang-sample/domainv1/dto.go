package domainv1

import (
	"github.com/lovemew67/public-misc/golang-sample/gen/go/proto"
)

// staff v1

type CreateStaffV1ServiceRequest struct {
	*proto.StaffV1
}

type GetStaffV1ServiceRequest struct {
	ID string
}

type ListStaffV1ServiceRequest struct {
	Offset int `form:"offset" url:"offset"`
	Limit  int `form:"limit" url:"limit"`
}

type PatchStaffV1ServiceRequest struct {
	ID string

	Name      *string `json:"name"`
	Email     *string `json:"email"`
	AvatarUrl *string `json:"avatar_url"`
}

type DeleteStaffV1ServiceRequest struct {
	ID string
}

// job v1

type CreateJobV1Request struct {
	*Job
}

type ListJobV1Request struct {
	Offset int `form:"offset" url:"offset"`
	Limit  int `form:"limit" url:"limit"`
}

type PatchJobV1Request struct {
	ID string

	RetryCount *int              `json:"retry_count"`
	Status     *int              `json:"status"`
	Type       *JobType          `json:"type"`
	Processing *JobProcessStatus `json:"processing"`
}

type DeleteJobV1Request struct {
	ID string
}

// schedule v1

type CreateScheduleV1Request struct {
	*Schedule
}

type ListScheduleV1Request struct {
	Offset int `form:"offset" url:"offset"`
	Limit  int `form:"limit" url:"limit"`
}

type PatchScheduleV1Request struct {
	ID string

	TimeInHourUTC *int     `json:"time_in_hour_utc"`
	Enable        *bool    `json:"enable"`
	Type          *JobType `json:"type"`
}

type DeleteScheduleV1Request struct {
	ID string
}
