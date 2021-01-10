package servicev1

import (
	"github.com/lovemew67/project-misc/rest-server-0/modelv1"
)

type CreateStaffV1ServiceRequest struct {
	*modelv1.Staff
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
