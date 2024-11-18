package models

type UserDTO struct {
	UserID      string `json:"userId"`
	Email       string `json:"email"`
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
	IsActive    bool   `json:"isActive"`
	Avatar      string `json:"avatar"`
	RoleName    string `json:"roleName"`
}
