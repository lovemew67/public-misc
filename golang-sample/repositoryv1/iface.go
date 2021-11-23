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
	Insert(job domainv1.Job) error
	QueryReadyTask() ([]domainv1.Job, error)
	UpdateProcessStatusToOngoing(id int) error
	CancelTaskByID(id int) error
	RemoveFromTaskQueue(task *domainv1.Job) error
	UpdateTaskStatusStillOngoing(task *domainv1.Job) error
	UpdateTaskStatusToStopped(task *domainv1.Job) error
}
