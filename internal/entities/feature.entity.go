package entities

import (
	"github.com/google/uuid"
)

type Feature struct {
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey;"`
	Name         string     `json:"name" gorm:"type:varchar"`
	ParentMenuId *uuid.UUID `json:"parentMenuId"`
	// ParentMenu   *Feature   `json:"parentMenu"`
	MenuIcon   string `json:"menuIcon" gorm:"type:varchar"`
	MenuNameTh string `json:"menuNameTh" gorm:"type:varchar"`
	MenuNameEn string `json:"menuNameEn" gorm:"type:varchar"`
	MenuSlug   string `json:"menuSlug" gorm:"type:varchar"`
	MenuSeqNo  string `json:"menuSeqNo" gorm:"type:varchar"`
	IsActive   *bool  `json:"isActive" gorm:"default:true"`
	Roles      []Role `json:"-" gorm:"many2many:role_features;"`
}

type FeatureDTO struct {
	FeatureDTOID uuid.UUID `json:"featureId"`
	FeatureName  string    `json:"featureName"`
	IsView       *bool     `json:"isView"`
	IsAdd        *bool     `json:"isAdd"`
	IsEdit       *bool     `json:"isEdit"`
	IsDelete     *bool     `json:"isDelete"`
}

type FeatureDTODetails struct {
	ID           uuid.UUID  `json:"featureId"`
	Name         string     `json:"name"`
	ParentMenuId *uuid.UUID ` json:"parentMenuId,omitempty"`
	MenuIcon     *string    `json:"menuIcon"`
	MenuNameTh   string     `json:"menuNameTh"`
	MenuNameEn   string     `json:"menuNameEn"`
	MenuSlug     string     `json:"menuSlug"`
	MenuSeqNo    string     `json:"menuSeqNo"`
	IsActive     *bool      `json:"isActive"`
	IsAdd        *bool      `json:"isAdd"`
	IsView       *bool      `json:"isView"`
	IsEdit       *bool      `json:"isEdit"`
	IsDelete     *bool      `json:"isDelete"`
}

type RefFeatureDTO struct {
	FeatureDTOID uuid.UUID `json:"featureId"`
	FeatureName  string    `json:"featureName"`
	IsView       *bool     `json:"isView"`
	IsAdd        *bool     `json:"isAdd"`
	IsEdit       *bool     `json:"isEdit"`
	IsDelete     *bool     `json:"isDelete"`
}

type ResMenuIcon struct {
	MenuIcon string `json:"menuIcon"`
}
