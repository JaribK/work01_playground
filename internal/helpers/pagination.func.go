package helpers

import (
	"math"
	"work01/internal/models"
)

func Pagiante(page, size int, total int64, items interface{}) models.Pagination {
	totalPage := int(math.Ceil(float64(total) / float64(size)))

	return models.Pagination{
		Page:      page,
		TotalPage: totalPage,
		Size:      size,
		Total:     total,
		Items:     items,
	}
}
