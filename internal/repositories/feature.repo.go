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
	GetAll(ctx context.Context) ([]entities.Feature, error)
	GetAllFeaturePermission(ctx context.Context) ([]models.FeatureDTO, error)
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
	var permission entities.Permission
	var featurePermission models.FeatureDTO

	cacheKey := fmt.Sprintf("feature:%s", id)
	if err := r.redisCache.Get(ctx, cacheKey, &featurePermission); err == nil {
		return &featurePermission, nil
	}

	err := r.db.Model(&entities.Permission{}).Preload("Feature").Joins("RIGHT JOIN features ON permissions.feature_id = features.id").Where("permissions.feature_id=?", id).First(&permission).Error
	if err != nil {
		return nil, err
	}

	featurePermission = models.FeatureDTO{
		FeatureDTOID: permission.Feature.ID,
		FeatureName:  permission.Feature.Name,
		IsView:       permission.ReadAccess,
		IsAdd:        permission.CreateAccess,
		IsEdit:       permission.UpdateAccess,
		IsDelete:     permission.DeleteAccess,
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: featurePermission,
		TTL:   time.Minute * 10,
	}); err != nil {
		return nil, err
	}

	return &featurePermission, nil
}

func (r *featureRepository) GetAll(ctx context.Context) ([]entities.Feature, error) {
	var features []entities.Feature

	cacheKey := fmt.Sprintln("feature_list")
	if err := r.redisCache.Get(ctx, cacheKey, &features); err == nil {
		return features, nil
	}

	if err := r.db.Find(&features).Error; err != nil {
		return nil, err
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: features,
		TTL:   time.Minute * 10,
	}); err != nil {
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

	cacheKey2 := "permission_feature_list"
	if err := r.redisCache.Delete(ctx, cacheKey2); err != nil {
		return err
	}

	var permissions []entities.Permission

	if err := r.db.Model(&entities.Permission{}).Preload("Feature").Joins("LEFT JOIN features ON permissions.feature_id = features.id").Find(&permissions).Error; err != nil {
		return err
	}

	var featurePermission []models.FeatureDTO
	for _, permission := range permissions {
		featurePermission = append(featurePermission, models.FeatureDTO{
			FeatureDTOID: permission.Feature.ID,
			FeatureName:  permission.Feature.Name,
			IsView:       permission.ReadAccess,
			IsAdd:        permission.CreateAccess,
			IsEdit:       permission.UpdateAccess,
			IsDelete:     permission.DeleteAccess,
		})
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey2,
		Value: featurePermission,
		TTL:   time.Minute * 10,
	}); err != nil {
		return err
	}

	return nil
}

func (r *featureRepository) GetAllFeaturePermission(ctx context.Context) ([]models.FeatureDTO, error) {
	var featurePermission []models.FeatureDTO
	var permissions []entities.Permission

	cacheKey1 := "permission_feature_list"
	if err := r.redisCache.Get(ctx, cacheKey1, &featurePermission); err == nil {
		return featurePermission, nil
	}

	if err := r.db.Model(&entities.Permission{}).Preload("Feature").Joins("LEFT JOIN features ON permissions.feature_id = features.id").Find(&permissions).Error; err != nil {
		return nil, err
	}

	for _, permission := range permissions {
		featurePermission = append(featurePermission, models.FeatureDTO{
			FeatureDTOID: permission.Feature.ID,
			FeatureName:  permission.Feature.Name,
			IsView:       permission.ReadAccess,
			IsAdd:        permission.CreateAccess,
			IsEdit:       permission.UpdateAccess,
			IsDelete:     permission.DeleteAccess,
		})
	}

	cacheKey2 := "permission_feature_list"
	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey2,
		Value: featurePermission,
		TTL:   time.Minute * 10,
	}); err != nil {
		return nil, err
	}

	return featurePermission, nil
}

func (r *featureRepository) Delete(id uuid.UUID) error {
	if err := r.db.Delete(&entities.Feature{}, id).Error; err != nil {
		return err
	}
	return nil
}
