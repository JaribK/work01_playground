package models

import "github.com/google/uuid"

type ResRoleDropDown struct {
	RoleID   uuid.UUID `json:"roleId"`
	RoleName string    `json:"roleName"`
}

type ResRoleDetails struct {
	RoleID     uuid.UUID `json:"roleId"`
	RoleName   string    `json:"roleName"`
	RoleLevel  int32     `json:"roleLevel"`
	NumberUser int32     `json:"numberUser"`
}
