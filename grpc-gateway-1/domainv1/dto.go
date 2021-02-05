package domainv1

import (
	"github.com/lovemew67/project-misc/grpc-gateway-1/gen/proto"
)

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
