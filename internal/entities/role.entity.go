package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID          uuid.UUID      `json:"id" gorm:"type:uuid;primaryKey;"`
	Name        string         `json:"name" gorm:"not null;unique"`
	Level       int32          `json:"level" gorm:"not null;default:0"`
	CreatedAt   time.Time      `json:"createdAt"`
	CreatedBy   uuid.UUID      `json:"createdBy,omitempty" gorm:"type:uuid"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	UpdatedBy   uuid.UUID      `json:"updatedBy" gorm:"type:uuid"`
	DeletedAt   gorm.DeletedAt `json:"-"`
	DeletedBy   *uuid.UUID     `json:"-" gorm:"type:uuid;index;"`
	Permissions []Permission   `json:"permissions" gorm:"many2many:role_permissions;"`
}

type RolePermission struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	RoleID uuid.UUID `json:"roleId" gorm:"type:uuid;primaryKey;"`
	// Role Role
	PermissionID uuid.UUID `json:"permissionId" gorm:"type:uuid;primaryKey;"`
	// Permission Permission
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
