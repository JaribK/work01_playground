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

type permissionRepository struct {
	db         *gorm.DB
	redisCache *cache.Cache
}

type PermissionRepository interface {
	Create(permission *entities.Permission) error
	GetById(ctx context.Context, id uuid.UUID) (*entities.Permission, error)
	GetAll(ctx context.Context) ([]entities.Permission, error)
	Update(permission *entities.Permission) error
	Delete(id uuid.UUID) error
}

func NewPermissionRepository(db *gorm.DB, redisClient *redis.Client) PermissionRepository {
	c := cache.New(&cache.Options{
		Redis:      redisClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	return &permissionRepository{db: db, redisCache: c}
}

func (r *permissionRepository) Create(permission *entities.Permission) error {
	if err := r.db.Create(&permission).Error; err != nil {
		return err
	}
	return nil
}

func (r *permissionRepository) GetById(ctx context.Context, id uuid.UUID) (*entities.Permission, error) {
	var permission entities.Permission
	cacheKey := fmt.Sprintf("permission:%s", id)

	if err := r.redisCache.Get(ctx, cacheKey, &permission); err == nil {
		return &permission, nil
	}

	if err := r.db.Preload("Feature").First(&permission, id).Error; err != nil {
		return nil, err
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: permission,
		TTL:   time.Minute * 10,
	}); err != nil {
		return nil, err
	}

	return &permission, nil
}

func (r *permissionRepository) GetAll(ctx context.Context) ([]entities.Permission, error) {
	var permissions []entities.Permission
	cacheKey := fmt.Sprintln("permission_list")

	if err := r.redisCache.Get(ctx, cacheKey, &permissions); err == nil {
		return permissions, nil
	}

	if err := r.db.Preload("Feature").Find(&permissions).Error; err != nil {
		return nil, err
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: permissions,
		TTL:   time.Minute * 10,
	}); err != nil {
		return nil, err
	}

	return permissions, nil
}

func (r *permissionRepository) Update(permission *entities.Permission) error {
	if err := r.db.Where("id=?", permission.ID).Updates(&permission).Error; err != nil {
		return err
	}
	return nil
}

func (r *permissionRepository) Delete(id uuid.UUID) error {
	if err := r.db.Delete(&entities.Permission{}, id).Error; err != nil {
		return err
	}
	return nil
}
