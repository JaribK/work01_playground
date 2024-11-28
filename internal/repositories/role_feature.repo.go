package repositories

import (
	"context"
	"time"
	"work01/internal/entities"

	"github.com/go-redis/cache/v9"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type roleFeatureRepository struct {
	db         *gorm.DB
	redisCache *cache.Cache
}

type RoleFeatureRepository interface {
	Create(roleFeature *entities.RoleFeature) error
	GetById(ctx context.Context, id uuid.UUID) (*entities.RoleFeature, error)
	GetAll(ctx context.Context) ([]entities.RoleFeature, error)
	Update(roleFeature *entities.RoleFeature) error
	Delete(id uuid.UUID) error
}

func NewRoleFeatureRepository(db *gorm.DB, redisClient *redis.Client) RoleFeatureRepository {
	c := cache.New(&cache.Options{
		Redis:      redisClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	return &roleFeatureRepository{db: db, redisCache: c}
}

func (r *roleFeatureRepository) Create(roleFeature *entities.RoleFeature) error {
	if err := r.db.Create(&roleFeature).Error; err != nil {
		return err
	}
	return nil
}

func (r *roleFeatureRepository) GetById(ctx context.Context, id uuid.UUID) (*entities.RoleFeature, error) {
	var roleFeature entities.RoleFeature

	if err := r.db.First(&roleFeature, id).Error; err != nil {
		return nil, err
	}

	return &roleFeature, nil
}

func (r *roleFeatureRepository) GetAll(ctx context.Context) ([]entities.RoleFeature, error) {
	var roleFeature []entities.RoleFeature

	if err := r.db.Preload("Feature").Find(&roleFeature).Error; err != nil {
		return nil, err
	}

	return roleFeature, nil
}

func (r *roleFeatureRepository) Update(roleFeature *entities.RoleFeature) error {
	if err := r.db.Where("id=?", roleFeature.ID).Updates(&roleFeature).Error; err != nil {
		return err
	}
	return nil
}

func (r *roleFeatureRepository) Delete(id uuid.UUID) error {
	if err := r.db.Delete(&entities.RoleFeature{}, id).Error; err != nil {
		return err
	}
	return nil
}
