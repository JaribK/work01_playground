package repositories

import (
	"work01/internal/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type featureRepository struct {
	db *gorm.DB
}

type FeatureRepository interface {
	Create(user *entities.Feature) error
	GetById(id uuid.UUID) (*entities.Feature, error)
	GetAll() ([]entities.Feature, error)
	Update(user *entities.Feature) error
	Delete(id uuid.UUID) error
}

func NewFeatureRepository(db *gorm.DB) FeatureRepository {
	return &featureRepository{db: db}
}

func (r *featureRepository) Create(feature *entities.Feature) error {
	err := r.db.Create(&feature).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *featureRepository) GetById(id uuid.UUID) (*entities.Feature, error) {
	var feature entities.Feature
	err := r.db.First(&feature, id).Error
	if err != nil {
		return nil, err
	}
	return &feature, nil
}

func (r *featureRepository) GetAll() ([]entities.Feature, error) {
	var features []entities.Feature
	err := r.db.Find(&features).Error
	if err != nil {
		return nil, err
	}
	return features, nil
}

func (r *featureRepository) Update(feature *entities.Feature) error {
	err := r.db.Where("id=?", feature.ID).Updates(&feature).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *featureRepository) Delete(id uuid.UUID) error {
	err := r.db.Delete(&entities.Feature{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
