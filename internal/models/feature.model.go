package models

import "github.com/google/uuid"

type FeatureDTO struct {
	FeatureDTOID uuid.UUID `json:"featureId"`
	FeatureName  string    `json:"featureName"`
	IsView       bool      `json:"isView"`
	IsAdd        bool      `json:"isAdd"`
	IsEdit       bool      `json:"isEdit"`
	IsDelete     bool      `json:"isDelete"`
}
