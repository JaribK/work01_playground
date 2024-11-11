package entities

import (
	"github.com/google/uuid"
)

type Permission struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	FeatureId    uuid.UUID `json:"feature_id:not null"`
	CreateAccess bool      `json:"create_access" gorm:"default:false;"`
	ReadAccess   bool      `json:"read_access" gorm:"default:false;"`
	UpdateAccess bool      `json:"update_access" gorm:"default:false;"`
	DeleteAccess bool      `json:"delete_access" gorm:"default:false;"`
	Roles        []Role    `gorm:"many2many:role_permissions;"`
}
