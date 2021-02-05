package repositoryv1

import (
	"github.com/lovemew67/project-misc/grpc-gateway-1/gen/proto"
)

type StaffV1Repository interface {
	CreateStaff(*proto.StaffV1) (*proto.StaffV1, error)
	CountTotalStaff() (int, error)
	GetStaff(id string) (*proto.StaffV1, error)
	QueryAllStaffWithOffsetAndLimit(offset, limit int) ([]*proto.StaffV1, error)
	PatchStaff(string, *proto.StaffV1) error
	DeleteStaff(string) error
}
