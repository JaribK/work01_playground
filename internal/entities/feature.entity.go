package entities

import (
	"github.com/google/uuid"
)

type Feature struct {
	ID           uuid.UUID  `json:"id" gorm:"type:uuid;primaryKey;"`
	Name         string     `json:"name" gorm:"type:varchar"`
	ParentMenuId *uuid.UUID `json:"parent_menu_id"`
	ParentMenu   *Feature
	MenuIcon     string `json:"menu_icon" gorm:"type:varchar"`
	MenuNameTh   string `json:"menu_name_th" gorm:"type:varchar"`
	MenuNameEn   string `json:"menu_name_en" gorm:"type:varchar"`
	MenuSlug     string `json:"menu_slug" gorm:"type:varchar"`
	MenuSeqNo    string `json:"menu_seq_no" gorm:"type:varchar"`
	IsActive     bool   `json:"is_active"`
}
