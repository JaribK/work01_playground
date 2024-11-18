package models

import "github.com/google/uuid"

type ResRole struct {
	RoleID   uuid.UUID `json:"roleId"`
	RoleName string    `json:"roleName"`
}
