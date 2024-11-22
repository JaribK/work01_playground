package usecases

import (
	"context"
	"work01/internal/entities"
	"work01/internal/models"
	"work01/internal/repositories"

	"github.com/google/uuid"
)

type FeatureUsecase interface {
	CreateFeature(feature entities.Feature) error
	GetFeatureById(ctx context.Context, id uuid.UUID) (*models.FeatureDTO, error)
	GetAllFeatures(ctx context.Context) ([]models.FeatureDTO, error)
	UpdateFeature(ctx context.Context, feature entities.Feature) error
	DeleteFeature(id uuid.UUID) error
}

type featureUsecase struct {
	repo repositories.FeatureRepository
}

func NewFeatureUsecase(repo repositories.FeatureRepository) FeatureUsecase {
	return &featureUsecase{repo: repo}
}

func (s *featureUsecase) CreateFeature(feature entities.Feature) error {
	if err := s.repo.Create(&feature); err != nil {
		return err
	}
	return nil
}

func (s *featureUsecase) GetFeatureById(ctx context.Context, id uuid.UUID) (*models.FeatureDTO, error) {
	feature, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return feature, nil
}

func (s *featureUsecase) GetAllFeatures(ctx context.Context) ([]models.FeatureDTO, error) {
	features, err := s.repo.GetAllFeaturePermission(ctx)
	if err != nil {
		return nil, err
	}
	return features, nil
}

func (s *featureUsecase) UpdateFeature(ctx context.Context, feature entities.Feature) error {
	if err := s.repo.Update(ctx, &feature); err != nil {
		return err
	}

	return nil
}

func (s *featureUsecase) DeleteFeature(id uuid.UUID) error {
	if err := s.repo.Delete(id); err != nil {
		return err
	}

	return nil
}
