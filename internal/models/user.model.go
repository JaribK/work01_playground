package models

import (
	"time"
	"work01/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
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
	UserID            uuid.UUID             `json:"userId"`
	Email             string                `json:"email"`
	FirstName         string                `json:"firstName"`
	LastName          string                `json:"lastName"`
	PhoneNumber       string                `json:"phoneNumber"`
	Avatar            string                `json:"avatar"`
	RoleId            uuid.UUID             `json:"roleId"`
	RoleName          string                `json:"roleName"`
	RoleLevel         int32                 `json:"roleLevel"`
	TwoFactorEnabled  bool                  `json:"twoFactorEnabled"`
	TwoFactorVerified bool                  `json:"twoFactorVerified"`
	TwoFactorAuthUrl  string                `json:"twoFactorAuthUrl"`
	TwoFactorToken    string                `json:"twoFactorToken"`
	Permission        []entities.Permission `json:"-"`
	Permissions       []PermissionDTO       `json:"permissions"`
}

type ResUsersNoPage struct {
	ID                 uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey"`
	FirstName          string          `json:"firstName" gorm:"type:varchar;not null"`
	LastName           string          `json:"lastName" gorm:"type:varchar;not null"`
	Email              string          `json:"email" gorm:"type:varchar;not null"`
	PhoneNumber        string          `json:"phoneNumber" gorm:"type:varchar;not null"`
	Avatar             string          `json:"avatar" gorm:"type:varchar;"`
	TwoFactorEnabled   bool            `json:"twoFacterEnabled" gorm:"not null;default:false"`
	TwoFactorVerified  bool            `json:"twoFacterVerified" gorm:"not null;default:false"`
	TwoFactorToken     string          `json:"twoFacterToken" gorm:"type:varchar"`
	TwoFactorAuthUrl   string          `json:"twoFacterAuthUrl" gorm:"type:varchar"`
	RoleId             *uuid.UUID      `json:"roleId" gorm:"type:uuid"`
	Role               entities.Role   `json:"role"`
	ForgotPasswordCode string          `json:"forgotPasswordCode" gorm:"type:varchar"`
	IsActive           bool            `json:"isActive" gorm:"default:true"`
	CreatedAt          time.Time       `json:"createdAt"`
	CreatedBy          uuid.UUID       `json:"createdBy" gorm:"type:uuid"`
	UpdatedAt          time.Time       `json:"updatedAt"`
	UpdatedBy          uuid.UUID       `json:"updatedBy" gorm:"type:uuid"`
	DeletedAt          *gorm.DeletedAt `json:"-"`
	DeletedBy          *uuid.UUID      `json:"-" gorm:"type:uuid;index;"`
}
