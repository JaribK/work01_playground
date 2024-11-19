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

type roleRepository struct {
	db         *gorm.DB
	redisCache *cache.Cache
}

type RoleRepository interface {
	GetById(ctx context.Context, id uuid.UUID) (*entities.Role, error)
	GetAll(ctx context.Context) ([]entities.Role, error)
	Create(role *entities.Role) error
	Update(role *entities.Role) error
	Delete(id uuid.UUID, delBy uuid.UUID) error
}

func NewRoleRepository(db *gorm.DB, redisClient *redis.Client) RoleRepository {
	c := cache.New(&cache.Options{
		Redis:      redisClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	return &roleRepository{db: db, redisCache: c}
}

func (r *roleRepository) GetById(ctx context.Context, id uuid.UUID) (*entities.Role, error) {
	var roleOjb entities.Role

	cacheKey := fmt.Sprintf("role:%s", id)

	if err := r.redisCache.Get(ctx, cacheKey, &roleOjb); err == nil {
		return &roleOjb, nil
	}

	err := r.db.First(&roleOjb, id).Error
	if err != nil {
		return nil, err
	}

	err = r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: roleOjb,
		TTL:   time.Minute * 10,
	})

	if err != nil {
		return nil, err
	}

	return &roleOjb, nil
}

func (r *roleRepository) GetAll(ctx context.Context) ([]entities.Role, error) {
	var roleOjbs []entities.Role

	cacheKey := fmt.Sprintln("roles_list")

	if err := r.redisCache.Get(ctx, cacheKey, &roleOjbs); err == nil {
		return roleOjbs, nil
	}

	err := r.db.Preload("Permissions").Preload("Users").Find(&roleOjbs).Error
	if err != nil {
		return nil, err
	}

	err = r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: roleOjbs,
		TTL:   time.Minute * 10,
	})

	if err != nil {
		return nil, err
	}

	return roleOjbs, nil
}

func (r *roleRepository) Create(role *entities.Role) error {
	err := r.db.Create(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) Update(role *entities.Role) error {
	err := r.db.Where("id=?", role.ID).Updates(&role).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) Delete(id uuid.UUID, delBy uuid.UUID) error {
	err := r.db.Model(&entities.Role{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_by": delBy,
	}).Error
	if err != nil {
		return err
	}

	err = r.db.Delete(&entities.Role{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
