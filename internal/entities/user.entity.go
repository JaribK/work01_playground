package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// `gorm:"unique;not null;type:varchar(100);default:null"`
type User struct {
	ID                uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey"`
	FirstName         string     `json:"firstName" gorm:"type:varchar;not null"`
	LastName          string     `json:"lastName" gorm:"type:varchar;not null"`
	Email             string     `json:"email" gorm:"type:varchar;not null"`
	PhoneNumber       string     `json:"phoneNumber" gorm:"type:varchar;not null"`
	Password          string     `json:"password" gorm:"type:varchar;not null"`
	Avatar            string     `json:"avatar" gorm:"type:varchar;"`
	TwoFactorEnabled  bool       `json:"twoFacterEnabled" gorm:"not null;default:false"`
	TwoFactorVerified bool       `json:"twoFacterVerified" gorm:"not null;default:false"`
	TwoFactorToken    string     `json:"twoFacterToken" gorm:"type:varchar"`
	TwoFactorAuthUrl  string     `json:"twoFacterAuthUrl" gorm:"type:varchar"`
	RoleId            *uuid.UUID `json:"roleId" gorm:"type:uuid"`
	// Role               Role       `json:"role"`
	ForgotPasswordCode string         `json:"forgotPasswordCode" gorm:"type:varchar"`
	IsActive           bool           `json:"isActive" gorm:"default:true"`
	CreatedAt          time.Time      `json:"createdAt"`
	CreatedBy          uuid.UUID      `json:"createdBy" gorm:"type:uuid"`
	UpdatedAt          time.Time      `json:"updatedAt"`
	UpdatedBy          uuid.UUID      `json:"updatedBy" gorm:"type:uuid"`
	DeletedAt          gorm.DeletedAt `json:"-"`
	DeletedBy          *uuid.UUID     `json:"-" gorm:"type:uuid;index;"`
}
