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

type rolePermissionRepository struct {
	db         *gorm.DB
	redisCache *cache.Cache
}

type RolePermissionRepository interface {
	Create(rolePermission *entities.RolePermission) error
	GetById(ctx context.Context, id uuid.UUID) (*entities.RolePermission, error)
	GetAll(ctx context.Context) ([]entities.RolePermission, error)
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

func (r *rolePermissionRepository) Create(rolePermission *entities.RolePermission) error {
	if err := r.db.Create(&rolePermission).Error; err != nil {
		return err
	}
	return nil
}

func (r *rolePermissionRepository) GetById(ctx context.Context, id uuid.UUID) (*entities.RolePermission, error) {
	var rolePermission entities.RolePermission

	if err := r.db.First(&rolePermission, id).Error; err != nil {
		return nil, err
	}

	return &rolePermission, nil
}

func (r *rolePermissionRepository) GetAll(ctx context.Context) ([]entities.RolePermission, error) {
	var rolePermission []entities.RolePermission

	if err := r.db.Find(&rolePermission).Error; err != nil {
		return nil, err
	}

	return rolePermission, nil
}

func (r *rolePermissionRepository) Update(rolePermission *entities.RolePermission) error {
	if err := r.db.Where("id=?", rolePermission.ID).Updates(&rolePermission).Error; err != nil {
		return err
	}
	return nil
}

func (r *rolePermissionRepository) Delete(id uuid.UUID) error {
	if err := r.db.Delete(&entities.RolePermission{}, id).Error; err != nil {
		return err
	}
	return nil
}
