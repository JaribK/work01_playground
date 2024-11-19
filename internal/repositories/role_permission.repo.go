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

type rolePermissionRepository struct {
	db         *gorm.DB
	redisCache *cache.Cache
}

type RolePermissionRepository interface {
	GetById(ctx context.Context, id uuid.UUID) (*entities.RolePermission, error)
	GetAll(ctx context.Context) ([]entities.RolePermission, error)
	Create(rolePermission *entities.RolePermission) error
	Update(rolePermission *entities.RolePermission) error
	Delete(id uuid.UUID) error
}

func NewRolePermissionRepository(db *gorm.DB, redisClient *redis.Client) RolePermissionRepository {
	c := cache.New(&cache.Options{
		Redis:      redisClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	return &rolePermissionRepository{db: db, redisCache: c}
}

func (r *rolePermissionRepository) GetById(ctx context.Context, id uuid.UUID) (*entities.RolePermission, error) {
	var rolePermission entities.RolePermission

	cacheKey := fmt.Sprintf("role_permission:%s", id)

	if err := r.redisCache.Get(ctx, cacheKey, &rolePermission); err == nil {
		return &rolePermission, nil
	}

	err := r.db.First(&rolePermission, id).Error
	if err != nil {
		return nil, err
	}

	err = r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: rolePermission,
		TTL:   time.Minute * 10,
	})

	if err != nil {
		return nil, err
	}

	return &rolePermission, nil
}

func (r *rolePermissionRepository) GetAll(ctx context.Context) ([]entities.RolePermission, error) {
	var rolePermission []entities.RolePermission

	cacheKey := fmt.Sprintln("role_permission_list")

	if err := r.redisCache.Get(ctx, cacheKey, &rolePermission); err == nil {
		return rolePermission, nil
	}

	err := r.db.Find(&rolePermission).Error
	if err != nil {
		return nil, err
	}

	err = r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: rolePermission,
		TTL:   time.Minute * 10,
	})

	if err != nil {
		return nil, err
	}

	return rolePermission, nil
}

func (r *rolePermissionRepository) Create(rolePermission *entities.RolePermission) error {
	err := r.db.Create(&rolePermission).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *rolePermissionRepository) Update(rolePermission *entities.RolePermission) error {
	err := r.db.Where("id=?", rolePermission.ID).Updates(&rolePermission).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *rolePermissionRepository) Delete(id uuid.UUID) error {
	err := r.db.Delete(&entities.RolePermission{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
