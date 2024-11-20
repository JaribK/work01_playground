package models

import "github.com/google/uuid"

type PermissionDTO struct {
	ID           uuid.UUID  `json:"featureId"`
	Name         string     `json:"name"`
	ParentMenuId *uuid.UUID ` json:"parentMenuId,omitempty"`
	MenuIcon     string     `json:"menuIcon"`
	MenuNameTh   string     `json:"menuNameTh"`
	MenuNameEn   string     `json:"menuNameEn"`
	MenuSlug     string     `json:"menuSlug"`
	MenuSeqNo    string     `json:"menuSeqNo"`
	IsActive     bool       `json:"isActive"`
	CreateAccess bool       `json:"createAccess"`
	ReadAccess   bool       `json:"readAccess"`
	UpdateAccess bool       `json:"updateAccess"`
	DeleteAccess bool       `json:"deleteAccess"`
}
