package repositoryv1

import (
	"github.com/lovemew67/public-misc/golang-sample/domainv1"
	"github.com/lovemew67/public-misc/golang-sample/gen/go/proto"
)

type StaffV1Repository interface {
	CreateStaff(*proto.StaffV1) (*proto.StaffV1, error)
	CountTotalStaff() (int, error)
	GetStaff(id string) (*proto.StaffV1, error)
	QueryAllStaffWithOffsetAndLimit(offset, limit int) ([]*proto.StaffV1, error)
	PatchStaff(string, *proto.StaffV1) error
	DeleteStaff(string) error
}

type JobV1Repository interface {
	CreateJob(*domainv1.Job) (*domainv1.Job, error)
	CountTotalJobs() (int, error)
	GetJob(id string) (*domainv1.Job, error)
	QueryAllJobsWithOffsetAndLimit(offset, limit int) ([]*domainv1.Job, error)
	PatchJob(string, *domainv1.Job) error
	DeleteJob(string) error

	QueryReadyTask() ([]domainv1.Job, error)
	UpdateProcessStatusToOngoing(id int) error
	CancelTaskByID(id int) error

	RemoveFromTaskQueue(task *domainv1.Job) error
	UpdateTaskStatusStillOngoing(task *domainv1.Job) error
	UpdateTaskStatusToStopped(task *domainv1.Job) error
}

type ScheduleV1Repository interface {
	CreateSchedule(*domainv1.Schedule) (*domainv1.Schedule, error)
	CountTotalSchedules() (int, error)
	GetSchedule(id string) (*domainv1.Schedule, error)
	QueryAllSchedulesWithOffsetAndLimit(offset, limit int) ([]*domainv1.Schedule, error)
	PatchSchedule(string, *domainv1.Schedule) error
	DeleteSchedule(string) error
}
