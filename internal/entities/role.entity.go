package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role struct {
	ID        uuid.UUID       `json:"id" gorm:"type:uuid;primaryKey;"`
	Name      string          `json:"name" gorm:"not null;"`
	Level     int32           `json:"level" gorm:"not null;default:0"`
	CreatedAt time.Time       `json:"createdAt"`
	CreatedBy uuid.UUID       `json:"createdBy,omitempty" gorm:"type:uuid"`
	UpdatedAt time.Time       `json:"updatedAt"`
	UpdatedBy uuid.UUID       `json:"updatedBy" gorm:"type:uuid"`
	DeletedAt *gorm.DeletedAt `json:"-"`
	DeletedBy *uuid.UUID      `json:"-" gorm:"type:uuid;index;"`
	Users     []User          `json:"-"`
	Features  []Feature       `json:"features" gorm:"many2many:role_features;"`
}

type RoleFeature struct {
	ID     uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;"`
	RoleId uuid.UUID `json:"roleId" gorm:"type:uuid;primaryKey;"`
	// Role Role
	FeatureId uuid.UUID `json:"featureId" gorm:"type:uuid;primaryKey;"`
	Feature   Feature   `json:"features"`
	IsAdd     *bool     `json:"isAdd" gorm:"default:false;"`
	IsView    *bool     `json:"isView" gorm:"default:false;"`
	IsEdit    *bool     `json:"isEdit" gorm:"default:false;"`
	IsDelete  *bool     `json:"isDelete" gorm:"default:false;"`
}

type ResAllRoleDropDown struct {
	RoleID   uuid.UUID `json:"roleId"`
	RoleName string    `json:"roleName"`
}

type ResAllRoleDetails struct {
	RoleID     uuid.UUID `json:"roleId"`
	RoleName   string    `json:"roleName"`
	RoleLevel  int32     `json:"roleLevel"`
	NumberUser int32     `json:"numberUser"`
}

type ReqRoleCreate struct {
	Name     string          `json:"name"`
	Level    int32           `json:"level"`
	Features []FeatureInRole `json:"features"`
}

type FeatureInRole struct {
	FeatureId   uuid.UUID `json:"featureId"`
	FeatureName string    `json:"featureName"`
	IsAdd       *bool     `json:"isAdd"`
	IsView      *bool     `json:"isView"`
	IsEdit      *bool     `json:"isEdit"`
	IsDelete    *bool     `json:"isDelete"`
}

type ReqRoleUpdate struct {
	Name     string                `json:"name"`
	Level    int32                 `json:"level"`
	Features []FeatureInRoleUpdate `json:"features"`
}

type FeatureInRoleUpdate struct {
	FeatureId   uuid.UUID `json:"featureId"`
	FeatureName string    `json:"featureName"`
	IsAdd       *bool     `json:"isAdd"`
	IsView      *bool     `json:"isView"`
	IsEdit      *bool     `json:"isEdit"`
	IsDelete    *bool     `json:"isDelete"`
}

type ResRoleLevel struct {
	RoleLevel int32
}

type ResRoleDetails struct {
	RoleID    uuid.UUID       `json:"roleId"`
	RoleName  string          `json:"roleName"`
	RoleLevel int32           `json:"roleLevel"`
	Features  []FeatureInRole `json:"features"`
}
