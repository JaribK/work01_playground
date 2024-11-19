package models

type Pagination struct {
	Page      int         `json:"page"`
	TotalPage int         `json:"totalPage"`
	Size      int         `json:"size"`
	Total     int64       `json:"total"`
	Items     interface{} `json:"items"`
}