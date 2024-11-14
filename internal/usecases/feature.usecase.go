package usecases

import (
	"work01/internal/entities"
	"work01/internal/repositories"

	"github.com/google/uuid"
)

type FeatureUsecase interface {
	CreateFeature(feature entities.Feature) error
	GetFeatureById(id uuid.UUID) (*entities.Feature, error)
	GetAllFeatures() ([]entities.Feature, error)
	UpdateFeature(feature entities.Feature) error
	DeleteFeature(id uuid.UUID) error
}

type featureUsecase struct {
	repo repositories.FeatureRepository
}

func NewFeatureUsecase(repo repositories.FeatureRepository) FeatureUsecase {
	return &featureUsecase{repo: repo}
}

func (s *featureUsecase) CreateFeature(feature entities.Feature) error {
	err := s.repo.Create(&feature)
	if err != nil {
		return err
	}
	return nil
}

func (s *featureUsecase) GetFeatureById(id uuid.UUID) (*entities.Feature, error) {
	feature, err := s.repo.GetById(id)
	if err != nil {
		return nil, err
	}

	return feature, nil
}

func (s *featureUsecase) GetAllFeatures() ([]entities.Feature, error) {
	features, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}
	return features, nil
}

func (s *featureUsecase) UpdateFeature(feature entities.Feature) error {
	err := s.repo.Update(&feature)
	if err != nil {
		return err
	}

	return nil
}

func (s *featureUsecase) DeleteFeature(id uuid.UUID) error {
	err := s.repo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
