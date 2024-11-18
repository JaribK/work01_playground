package entities

import (
	"github.com/google/uuid"
)

type Permission struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	FeatureId    uuid.UUID `json:"featureId" gorm:"type:uuid;not null"`
	Feature      Feature   `json:"features"`
	CreateAccess bool      `json:"createAccess" gorm:"default:false;"`
	ReadAccess   bool      `json:"readAccess" gorm:"default:false;"`
	UpdateAccess bool      `json:"updateAccess" gorm:"default:false;"`
	DeleteAccess bool      `json:"deleteAccess" gorm:"default:false;"`
	// Roles        []*Role   `gorm:"many2many:role_permissions;"`
}
