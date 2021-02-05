package servicev1

import (
	"github.com/lovemew67/project-misc/grpc-gateway-1/domainv1"
	"github.com/lovemew67/project-misc/grpc-gateway-1/gen/proto"
	"github.com/lovemew67/project-misc/grpc-gateway-1/repositoryv1"
)

type StaffV1Servicer struct {
	r repositoryv1.StaffV1Repository
}

func (s *StaffV1Servicer) CreateStaffV1Service(req *domainv1.CreateStaffV1ServiceRequest) (result *proto.StaffV1, err error) {
	result, err = s.r.CreateStaff(req.StaffV1)
	return
}

func (s *StaffV1Servicer) GetStaffV1Service(req *domainv1.GetStaffV1ServiceRequest) (result *proto.StaffV1, err error) {
	result, err = s.r.GetStaff(req.ID)
	return
}

func (s *StaffV1Servicer) ListStaffV1Service(req *domainv1.ListStaffV1ServiceRequest) (results []*proto.StaffV1, total int, err error) {
	total, err = s.r.CountTotalStaff()
	if err != nil {
		return
	}
	results, err = s.r.QueryAllStaffWithOffsetAndLimit(req.Offset, req.Limit)
	return
}

func (s *StaffV1Servicer) PatchStaffV1Service(req *domainv1.PatchStaffV1ServiceRequest) (err error) {
	updater := &proto.StaffV1{}
	if req.Name != nil {
		updater.Name = *req.Name
	}
	if req.Email != nil {
		updater.Email = *req.Email
	}
	if req.AvatarUrl != nil {
		updater.AvatarUrl = *req.AvatarUrl
	}
	err = s.r.PatchStaff(req.ID, updater)
	return
}

func (s *StaffV1Servicer) DeleteStaffV1Service(req *domainv1.DeleteStaffV1ServiceRequest) (err error) {
	err = s.r.DeleteStaff(req.ID)
	return
}

func NewStaffV1Servicer(_r repositoryv1.StaffV1Repository) (result *StaffV1Servicer, err error) {
	result = &StaffV1Servicer{
		r: _r,
	}
	return
}
