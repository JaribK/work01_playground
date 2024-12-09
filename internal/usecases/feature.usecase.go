package usecases

import (
	"context"
	"mime/multipart"
	"work01/internal/entities"
	"work01/internal/repositories"
	"work01/pkg/minio"

	"github.com/google/uuid"
)

type (
	FeatureUsecase interface {
		CreateFeature(feature entities.Feature, fileHeader *multipart.FileHeader) error
		GetFeatureById(ctx context.Context, id uuid.UUID) (*entities.FeatureDTO, error)
		GetRefFeatures() ([]entities.RefFeatureDTO, error)
		GetAllFeaturesDefault() ([]entities.Feature, error)
		GetAllRoleFeatures(ctx context.Context) ([]entities.FeatureDTO, error)
		UpdateFeature(ctx context.Context, feature entities.Feature, fileHeader *multipart.FileHeader) error
		DeleteFeature(id uuid.UUID) error
	}

	featureUsecase struct {
		repo repositories.FeatureRepository
	}
)

func NewFeatureUsecase(repo repositories.FeatureRepository) FeatureUsecase {
	return &featureUsecase{repo: repo}
}

func (s *featureUsecase) CreateFeature(feature entities.Feature, fileHeader *multipart.FileHeader) error {
	if fileHeader != nil {
		avatarURL, err := minio.UploadAvatar(fileHeader)
		if err != nil {
			return err
		}
		feature.MenuIcon = avatarURL
	}

	if err := s.repo.Create(&feature); err != nil {
		return err
	}
	return nil
}

func (s *featureUsecase) GetFeatureById(ctx context.Context, id uuid.UUID) (*entities.FeatureDTO, error) {
	feature, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return feature, nil
}

func (s *featureUsecase) GetRefFeatures() ([]entities.RefFeatureDTO, error) {
	features, err := s.repo.RefForFeature()
	if err != nil {
		return nil, err
	}
	return features, nil
}

func (s *featureUsecase) GetAllRoleFeatures(ctx context.Context) ([]entities.FeatureDTO, error) {
	features, err := s.repo.GetAllRoleFeatures(ctx)
	if err != nil {
		return nil, err
	}
	return features, nil
}

func (s *featureUsecase) GetAllFeaturesDefault() ([]entities.Feature, error) {
	features, err := s.repo.GetAllDefault()
	if err != nil {
		return nil, err
	}
	return features, nil
}

func (s *featureUsecase) UpdateFeature(ctx context.Context, feature entities.Feature, fileHeader *multipart.FileHeader) error {
	menuIcon, err := s.repo.GetMenuIconByFeatureId(feature.ID)
	if err != nil {
		return err
	}

	if fileHeader != nil {
		menuIconURL, err := minio.UploadAvatarUpdate(fileHeader, menuIcon.MenuIcon)
		if err != nil {
			return err
		}
		feature.MenuIcon = menuIconURL
	}
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
