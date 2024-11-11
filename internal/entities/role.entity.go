package entities

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID        uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey;"`
	Name      string     `json:"name" gorm:"not null"`
	Level     int32      `json:"level" gorm:"not null"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy uuid.UUID  `json:"created_by,omitempty" gorm:"type:uuid"`
	UpdatedAt time.Time  `json:"updated_at"`
	UpdatedBy uuid.UUID  `json:"updated_by" gorm:"type:uuid"`
	DeletedAt *time.Time `json:"deleted_at"`
	DeletedBy *uuid.UUID `json:"deleted_by" gorm:"type:uuid;index"`
	// Permissions []Permission `gorm:"many2many:role_permissions;"`
}

// type Permission struct {
// 	ID           uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
// 	FeatureId    uuid.UUID `json:"feature_id:not null"`
// 	CreateAccess bool      `json:"create_access" gorm:"default:false;"`
// 	ReadAccess   bool      `json:"read_access" gorm:"default:false;"`
// 	UpdateAccess bool      `json:"update_access" gorm:"default:false;"`
// 	DeleteAccess bool      `json:"delete_access" gorm:"default:false;"`
// 	// Roles        []Role    `gorm:"many2many:role_permissions;"`
// }

// type Feature struct {
// 	ID           uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey;"`
// 	Name         string     `json:"name" gorm:"type:varchar"`
// 	ParentMenuId *uuid.UUID `json:"parent_menu_id"`
// 	ParentMenu   *Feature
// 	MenuIcon     string `json:"menu_icon" gorm:"type:varchar"`
// 	MenuNameTh   string `json:"menu_name_th" gorm:"type:varchar"`
// 	MenuNameEn   string `json:"menu_name_en" gorm:"type:varchar"`
// 	MenuSlug     string `json:"menu_slug" gorm:"type:varchar"`
// 	MenuSeqNo    string `json:"menu_seq_no" gorm:"type:varchar"`
// 	IsActive     bool   `json:"is_active"`
// }
