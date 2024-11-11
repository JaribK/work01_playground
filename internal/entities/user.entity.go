package entities

import (
	"time"

	"github.com/google/uuid"
)

// `gorm:"unique;not null;type:varchar(100);default:null"`
type User struct {
	ID                 uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	FirstName          string    `json:"first_name" gorm:"type:varchar;not null"`
	LastName           string    `json:"last_name" gorm:"type:varchar;not null"`
	Email              string    `json:"email" gorm:"type:varchar;not null"`
	PhoneNumber        string    `json:"phone_number" gorm:"type:varchar;not null"`
	Password           string    `json:"password" gorm:"type:varchar;not null"`
	Avatar             string    `json:"avatar" gorm:"type:varchar;"`
	TwoFactorEnabled   bool      `json:"two_facter_enabled" gorm:"not null;default:false"`
	TwoFactorVerified  bool      `json:"two_facter_verified" gorm:"not null;default:false"`
	TwoFactorToken     string    `json:"two_facter_token" gorm:"type:varchar"`
	TwoFactorAuthUrl   string    `json:"two_facter_auth_url" gorm:"type:varchar"`
	RoleId             uuid.UUID `json:"role_id" gorm:"type:uuid"`
	Role               Role
	ForgotPasswordCode string     `json:"forgot_password_code" gorm:"type:varchar"`
	IsActive           bool       `json:"is_active" gorm:"default:true"`
	CreatedAt          time.Time  `json:"created_at"`
	CreatedBy          uuid.UUID  `json:"created_by" gorm:"type:uuid"`
	UpdatedAt          time.Time  `json:"updated_at"`
	UpdatedBy          uuid.UUID  `json:"updated_by" gorm:"type:uuid"`
	DeletedAt          *time.Time `json:"deleted_at"`
	DeletedBy          *uuid.UUID `json:"deleted_by" gorm:"type:uuid;index"`
}
