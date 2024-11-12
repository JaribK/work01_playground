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
	IsActive   bool   `json:"isActive" gorm:"default:true"`
}
