package models

import (
	"work01/internal/entities"

	"github.com/google/uuid"
)

type ResAllUserDTOs struct {
	UserID      uuid.UUID `json:"userId"`
	Email       string    `json:"email"`
	FullName    string    `json:"fullName"`
	PhoneNumber string    `json:"phoneNumber"`
	IsActive    bool      `json:"isActive"`
	Avatar      string    `json:"avatar"`
	RoleName    string    `json:"roleName"`
}

type ResUserDTO struct {
	UserID            uuid.UUID `json:"userId"`
	Email             string    `json:"email"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	PhoneNumber       string    `json:"phoneNumber"`
	Avatar            string    `json:"avatar"`
	RoleName          string    `json:"roleName"`
	RoleLevel         int32     `json:"roleLevel"`
	TwoFactorAuthUrl  string    `json:"twoFactorAuthUrl"`
	TwoFactorEnabled  bool      `json:"twoFactorEnabled"`
	TwoFactorToken    string    `json:"twoFactorToken"`
	TwoFactorVerified bool      `json:"twoFactorVerified"`
	Permissions       []entities.Permission
}