package repositories

import (
	"context"
	"fmt"
	"time"
	"work01/internal/entities"

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
	GetById(ctx context.Context, id uuid.UUID) (*entities.Feature, error)
	GetAll(ctx context.Context) ([]entities.Feature, error)
	Update(user *entities.Feature) error
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
	err := r.db.Create(&feature).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *featureRepository) GetById(ctx context.Context, id uuid.UUID) (*entities.Feature, error) {
	var feature entities.Feature

	cacheKey := fmt.Sprintf("feature:%s", id)
	if err := r.redisCache.Get(ctx, cacheKey, &feature); err == nil {
		return &feature, nil
	}

	err := r.db.First(&feature, id).Error
	if err != nil {
		return nil, err
	}

	err = r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: feature,
		TTL:   time.Minute * 10,
	})

	if err != nil {
		return nil, err
	}

	return &feature, nil
}

func (r *featureRepository) GetAll(ctx context.Context) ([]entities.Feature, error) {
	var features []entities.Feature

	cacheKey := fmt.Sprintln("feature_list")
	if err := r.redisCache.Get(ctx, cacheKey, &features); err == nil {
		return features, nil
	}

	err := r.db.Find(&features).Error
	if err != nil {
		return nil, err
	}

	err = r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: features,
		TTL:   time.Minute * 10,
	})

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
