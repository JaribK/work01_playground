package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Authorization struct {
	ID           uuid.UUID       `json:"id" gorm:"type:uuid"`
	UserId       uuid.UUID       `json:"userId" gorm:"type:uuid"`
	AccessToken  string          `json:"accessToken"`
	RefreshToken string          `json:"refreshToken"`
	CreatedAt    time.Time       `json:"createdAt"`
	CreatedBy    uuid.UUID       `json:"createdBy,omitempty" gorm:"type:uuid"`
	UpdatedAt    time.Time       `json:"updatedAt"`
	UpdatedBy    uuid.UUID       `json:"updatedBy" gorm:"type:uuid"`
	DeletedAt    *gorm.DeletedAt `json:"-"`
	DeletedBy    *uuid.UUID      `json:"-" gorm:"type:uuid;index;"`
}
type AuthToken struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type ResLogin struct {
	Message      string      `json:"message"`
	AccessToken  string      `json:"accessToken"`
	RefreshToken string      `json:"refreshToken"`
	User         interface{} `json:"user"`
}
