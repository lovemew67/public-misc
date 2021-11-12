package servicev1

import (
	"github.com/lovemew67/public-misc/rest-server-1/gen/proto"
	"github.com/lovemew67/public-misc/rest-server-1/repositoryv1"
)

func CreateStaffV1Service(req *CreateStaffV1ServiceRequest) (result *proto.Staff, err error) {
	result, err = repositoryv1.CreateStaff(req.Staff)
	return
}

func GetStaffV1Service(req *GetStaffV1ServiceRequest) (result *proto.Staff, err error) {
	result, err = repositoryv1.GetStaff(req.ID)
	return
}

func ListStaffV1Service(req *ListStaffV1ServiceRequest) (results []proto.Staff, total int, err error) {
	total, err = repositoryv1.CountTotalStaff()
	if err != nil {
		return
	}
	results, err = repositoryv1.QueryAllStaffWithOffsetAndLimit(req.Offset, req.Limit)
	return
}

func PatchStaffV1Service(req *PatchStaffV1ServiceRequest) (err error) {
	updater := &proto.Staff{}
	if req.Name != nil {
		updater.Name = *req.Name
	}
	if req.Email != nil {
		updater.Email = *req.Email
	}
	if req.AvatarUrl != nil {
		updater.AvatarUrl = *req.AvatarUrl
	}
	err = repositoryv1.PatchStaff(req.ID, updater)
	return
}

func DeleteStaffV1Service(req *DeleteStaffV1ServiceRequest) (err error) {
	err = repositoryv1.DeleteStaff(req.ID)
	return
}
