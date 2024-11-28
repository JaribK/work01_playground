package models

import "github.com/google/uuid"

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
	MenuIcon     string     `json:"menuIcon"`
	MenuNameTh   string     `json:"menuNameTh"`
	MenuNameEn   string     `json:"menuNameEn"`
	MenuSlug     string     `json:"menuSlug"`
	MenuSeqNo    string     `json:"menuSeqNo"`
	IsActive     bool       `json:"isActive"`
	IsAdd        *bool      `json:"isAdd"`
	IsView       *bool      `json:"isView"`
	IsEdit       *bool      `json:"isEdit"`
	IsDelete     *bool      `json:"isDelete"`
}
