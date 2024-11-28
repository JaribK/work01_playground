package repositories

import (
	"context"
	"fmt"
	"time"
	"work01/internal/entities"
	"work01/internal/models"

	"github.com/go-redis/cache/v9"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type featureRepository struct {
	db         *gorm.DB
	redisCache *cache.Cache
}

type FeatureRepository interface {
	Create(user *entities.Feature) error
	GetById(ctx context.Context, id uuid.UUID) (*models.FeatureDTO, error)
	GetAllDefault() ([]entities.Feature, error)
	GetAllRoleFeatures(ctx context.Context) ([]models.FeatureDTO, error)
	Update(ctx context.Context, feature *entities.Feature) error
	Delete(id uuid.UUID) error
}

func NewFeatureRepository(db *gorm.DB, redisClient *redis.Client) FeatureRepository {
	c := cache.New(&cache.Options{
		Redis:      redisClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	return &featureRepository{db: db, redisCache: c}
}

func (r *featureRepository) Create(feature *entities.Feature) error {
	if err := r.db.Create(&feature).Error; err != nil {
		return err
	}
	return nil
}

func (r *featureRepository) GetById(ctx context.Context, id uuid.UUID) (*models.FeatureDTO, error) {
	var obj entities.RoleFeature
	var featureRole models.FeatureDTO

	cacheKey := fmt.Sprintf("feature:%s", id)
	if err := r.redisCache.Get(ctx, cacheKey, &featureRole); err == nil {
		return &featureRole, nil
	}

	if err := r.db.Model(&entities.RoleFeature{}).Preload("Feature").Joins("LEFT JOIN features ON role_features.feature_id = features.id").Where("role_features.feature_id=?", id).First(&obj).Error; err != nil {
		return nil, err
	}

	featureRole = models.FeatureDTO{
		FeatureDTOID: obj.Feature.ID,
		FeatureName:  obj.Feature.Name,
		IsView:       obj.IsView,
		IsAdd:        obj.IsAdd,
		IsEdit:       obj.IsEdit,
		IsDelete:     obj.IsDelete,
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: featureRole,
		TTL:   time.Minute * 10,
	}); err != nil {
		return nil, err
	}

	return &featureRole, nil
}

func (r *featureRepository) GetAllDefault() ([]entities.Feature, error) {
	var features []entities.Feature

	if err := r.db.Find(&features).Error; err != nil {
		return nil, err
	}

	return features, nil
}

func (r *featureRepository) Update(ctx context.Context, feature *entities.Feature) error {
	if err := r.db.Where("id=?", feature.ID).Updates(&feature).Error; err != nil {
		return err
	}

	cacheKey1 := fmt.Sprintf("feature:%s", feature.ID)
	if err := r.redisCache.Delete(ctx, cacheKey1); err != nil {
		return nil
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey1,
		Value: feature,
		TTL:   time.Minute * 10,
	}); err != nil {
		return err
	}

	// //query again

	cacheKey2 := "role_feature_list"
	if err := r.redisCache.Delete(ctx, cacheKey2); err != nil {
		return err
	}

	var roleFeatures []entities.RoleFeature
	var featureRole []models.FeatureDTO

	if err := r.db.Model(&entities.RoleFeature{}).Preload("Feature").Joins("LEFT JOIN features ON role_features.feature_id = features.id").Find(&roleFeatures).Error; err != nil {
		return err
	}

	for _, obj := range roleFeatures {
		featureRole = append(featureRole, models.FeatureDTO{
			FeatureDTOID: obj.Feature.ID,
			FeatureName:  obj.Feature.Name,
			IsView:       obj.IsView,
			IsAdd:        obj.IsAdd,
			IsEdit:       obj.IsEdit,
			IsDelete:     obj.IsDelete,
		})
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey2,
		Value: featureRole,
		TTL:   time.Minute * 10,
	}); err != nil {
		return err
	}

	return nil
}

func (r *featureRepository) GetAllRoleFeatures(ctx context.Context) ([]models.FeatureDTO, error) {
	var featureRole []models.FeatureDTO
	var roleFeatures []entities.RoleFeature

	cacheKey1 := "role_feature_list"
	if err := r.redisCache.Get(ctx, cacheKey1, &featureRole); err == nil {
		return featureRole, nil
	}

	if err := r.db.Model(&entities.RoleFeature{}).Preload("Feature").Joins("LEFT JOIN features ON role_features.feature_id = features.id").Find(&roleFeatures).Error; err != nil {
		return nil, err
	}

	for _, obj := range roleFeatures {
		featureRole = append(featureRole, models.FeatureDTO{
			FeatureDTOID: obj.Feature.ID,
			FeatureName:  obj.Feature.Name,
			IsView:       obj.IsView,
			IsAdd:        obj.IsAdd,
			IsEdit:       obj.IsEdit,
			IsDelete:     obj.IsDelete,
		})
	}

	cacheKey2 := "role_feature_list"
	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey2,
		Value: featureRole,
		TTL:   time.Minute * 10,
	}); err != nil {
		return nil, err
	}

	return featureRole, nil
}

func (r *featureRepository) Delete(id uuid.UUID) error {
	if err := r.db.Delete(&entities.Feature{}, id).Error; err != nil {
		return err
	}
	return nil
}
