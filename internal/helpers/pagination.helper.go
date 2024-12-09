package helpers

import (
	"math"
	"work01/internal/entities"
)

type Pagination struct {
	Page      int                       `json:"page"`
	TotalPage int                       `json:"totalPage"`
	Size      int                       `json:"size"`
	Total     int64                     `json:"total"`
	Items     []entities.ResAllUserDTOs `json:"items"`
}

func Pagiante(page, size int, total int64, items []entities.ResAllUserDTOs) Pagination {
	totalPage := int(math.Ceil(float64(total) / float64(size)))

	return Pagination{
		Page:      page,
		TotalPage: totalPage,
		Size:      size,
		Total:     total,
		Items:     items,
	}
}
