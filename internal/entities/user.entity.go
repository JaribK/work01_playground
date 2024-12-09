package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// `gorm:"unique;not null;type:varchar(100);default:null"`
type User struct {
	ID                 uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey"`
	FirstName          string          `json:"firstName" gorm:"type:varchar;not null"`
	LastName           string          `json:"lastName" gorm:"type:varchar;not null"`
	Email              string          `json:"email" gorm:"type:varchar;not null"`
	PhoneNumber        string          `json:"phoneNumber" gorm:"type:varchar;not null"`
	Password           string          `json:"password" gorm:"type:varchar;not null"`
	Avatar             string          `json:"avatar" gorm:"type:varchar;default:null;"`
	TwoFactorEnabled   *bool           `json:"twoFactorEnabled" gorm:"not null;default:false"`
	TwoFactorVerified  *bool           `json:"twoFactorVerified" gorm:"not null;default:false"`
	TwoFactorToken     string          `json:"twoFactorToken" gorm:"type:varchar;default:null;"`
	TwoFactorAuthUrl   string          `json:"twoFactorAuthUrl" gorm:"type:varchar;default:null;"`
	RoleId             *uuid.UUID      `json:"roleId" gorm:"type:uuid"`
	Role               Role            `json:"role"`
	ForgotPasswordCode string          `json:"forgotPasswordCode" gorm:"type:varchar"`
	IsActive           *bool           `json:"isActive" gorm:"default:true"`
	CreatedAt          time.Time       `json:"createdAt"`
	CreatedBy          uuid.UUID       `json:"createdBy" gorm:"type:uuid"`
	UpdatedAt          time.Time       `json:"updatedAt"`
	UpdatedBy          uuid.UUID       `json:"updatedBy" gorm:"type:uuid"`
	DeletedAt          *gorm.DeletedAt `json:"-"`
	DeletedBy          *uuid.UUID      `json:"-" gorm:"type:uuid;index;"`
}

type ReqUser struct {
	ID                 uuid.UUID       `json:"id"`
	FirstName          string          `json:"firstName"`
	LastName           string          `json:"lastName"`
	Email              string          `json:"email"`
	PhoneNumber        string          `json:"phoneNumber"`
	Password           string          `json:"password"`
	ConfirmPassword    string          `json:"confirmPassword"`
	Avatar             string          `json:"avatar"`
	TwoFactorEnabled   *bool           `json:"twoFactorEnabled"`
	TwoFactorVerified  *bool           `json:"twoFactorVerified"`
	TwoFactorToken     string          `json:"twoFactorToken"`
	TwoFactorAuthUrl   string          `json:"twoFactorAuthUrl"`
	RoleId             *uuid.UUID      `json:"roleId"`
	ForgotPasswordCode string          `json:"forgotPasswordCode"`
	IsActive           *bool           `json:"isActive"`
	CreatedAt          time.Time       `json:"createdAt"`
	CreatedBy          uuid.UUID       `json:"createdBy"`
	UpdatedAt          time.Time       `json:"updatedAt"`
	UpdatedBy          uuid.UUID       `json:"updatedBy"`
	DeletedAt          *gorm.DeletedAt `json:"-"`
	DeletedBy          *uuid.UUID      `json:"-"`
}

type ReqChangePassword struct {
	UserId             uuid.UUID
	UpdatedBy          uuid.UUID
	NewPassword        string `json:"newPassword"`
	ConfirmNewPassword string `json:"confirmNewPassword"`
}

type ResAllUserDTOs struct {
	UserID      uuid.UUID `json:"userId"`
	Email       string    `json:"email"`
	FullName    string    `json:"fullName"`
	PhoneNumber string    `json:"phoneNumber"`
	IsActive    bool      `json:"isActive"`
	Avatar      *string   `json:"avatar"`
	RoleName    string    `json:"roleName"`
}

type ResUserDTO struct {
	UserID            uuid.UUID `json:"userId"`
	Email             string    `json:"email"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	PhoneNumber       string    `json:"phoneNumber"`
	Avatar            *string   `json:"avatar"`
	RoleId            uuid.UUID `json:"roleId"`
	RoleName          string    `json:"roleName"`
	RoleLevel         int32     `json:"roleLevel"`
	TwoFactorEnabled  bool      `json:"twoFactorEnabled"`
	TwoFactorVerified bool      `json:"twoFactorVerified"`
	TwoFactorAuthUrl  *string   `json:"twoFactorAuthUrl"`
	TwoFactorToken    *string   `json:"twoFactorToken"`
	// Permission        []entities.Permission `json:"-"`
	Features []FeatureDTODetails `json:"permissions"`
}

type ResUsersNoPage struct {
	ID                 uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey"`
	FirstName          string          `json:"firstName" gorm:"type:varchar;not null"`
	LastName           string          `json:"lastName" gorm:"type:varchar;not null"`
	Email              string          `json:"email" gorm:"type:varchar;not null"`
	PhoneNumber        string          `json:"phoneNumber" gorm:"type:varchar;not null"`
	Avatar             *string         `json:"avatar" gorm:"type:varchar;"`
	TwoFactorEnabled   bool            `json:"twoFacterEnabled" gorm:"not null;default:false"`
	TwoFactorVerified  bool            `json:"twoFacterVerified" gorm:"not null;default:false"`
	TwoFactorToken     *string         `json:"twoFacterToken" gorm:"type:varchar"`
	TwoFactorAuthUrl   *string         `json:"twoFacterAuthUrl" gorm:"type:varchar"`
	RoleId             *uuid.UUID      `json:"roleId" gorm:"type:uuid"`
	Role               Role            `json:"role"`
	ForgotPasswordCode string          `json:"forgotPasswordCode" gorm:"type:varchar"`
	IsActive           bool            `json:"isActive" gorm:"default:true"`
	CreatedAt          time.Time       `json:"createdAt"`
	CreatedBy          uuid.UUID       `json:"createdBy" gorm:"type:uuid"`
	UpdatedAt          time.Time       `json:"updatedAt"`
	UpdatedBy          uuid.UUID       `json:"updatedBy" gorm:"type:uuid"`
	DeletedAt          *gorm.DeletedAt `json:"-"`
	DeletedBy          *uuid.UUID      `json:"-" gorm:"type:uuid;index;"`
}

type ResAvatar struct {
	Avatar string
}

type ResUserProfile struct {
	UserId           uuid.UUID     `json:"userId"`
	Email            string        `json:"email"`
	FirstName        string        `json:"firstName"`
	LastName         string        `json:"lastName"`
	PhoneNumber      string        `json:"phoneNumber"`
	Avatar           *string       `json:"avatar"`
	RoleId           uuid.UUID     `json:"roleId"`
	TwoFactorEnabled bool          `json:"twoFactorEnabled"`
	IsActive         bool          `json:"isActive" gorm:"default:true"`
	CreatedAt        time.Time     `json:"createdAt"`
	UserActivity     []interface{} `json:"userActivity"`
	UserDevice       []interface{} `json:"userDevice"`
}
