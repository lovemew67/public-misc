package servicev1

import (
	"github.com/lovemew67/public-misc/golang-sample/domainv1"
	"github.com/lovemew67/public-misc/golang-sample/gen/go/proto"
)

type StaffV1Service interface {
	CreateStaffV1Service(*domainv1.CreateStaffV1ServiceRequest) (*proto.StaffV1, error)
	GetStaffV1Service(*domainv1.GetStaffV1ServiceRequest) (*proto.StaffV1, error)
	ListStaffV1Service(*domainv1.ListStaffV1ServiceRequest) ([]*proto.StaffV1, int, error)
	PatchStaffV1Service(*domainv1.PatchStaffV1ServiceRequest) error
	DeleteStaffV1Service(*domainv1.DeleteStaffV1ServiceRequest) error
}
