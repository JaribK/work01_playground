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

type roleRepository struct {
	db         *gorm.DB
	redisCache *cache.Cache
}

type RoleRepository interface {
	GetById(ctx context.Context, id uuid.UUID) (*entities.Role, error)
	GetAllDefault() ([]entities.Role, error)
	GetAllModify(ctx context.Context) ([]models.ResRoleDetails, error)
	Create(role *entities.Role) error
	Update(ctx context.Context, role *entities.Role) error
	Delete(id uuid.UUID, delBy uuid.UUID) error
}

func NewRoleRepository(db *gorm.DB, redisClient *redis.Client) RoleRepository {
	c := cache.New(&cache.Options{
		Redis:      redisClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	return &roleRepository{db: db, redisCache: c}
}

func (r *roleRepository) Create(role *entities.Role) error {
	if err := r.db.Create(&role).Error; err != nil {
		return err
	}
	return nil
}

func (r *roleRepository) GetById(ctx context.Context, id uuid.UUID) (*entities.Role, error) {
	var roleOjb entities.Role
	cacheKey := fmt.Sprintf("role:%s", id)

	if err := r.redisCache.Get(ctx, cacheKey, roleOjb); err == nil {
		return &roleOjb, nil
	}

	if err := r.db.Preload("Features").Where("id=?", id).First(&roleOjb).Error; err != nil {
		return nil, err
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: roleOjb,
		TTL:   time.Minute * 10,
	}); err != nil {
		return nil, err
	}

	return &roleOjb, nil
}

func (r *roleRepository) GetAllModify(ctx context.Context) ([]models.ResRoleDetails, error) {
	var roleOjbs []entities.Role
	var roleRes []models.ResRoleDetails

	cacheKey := "roles_list"

	if err := r.redisCache.Get(ctx, cacheKey, roleOjbs); err == nil {
		return roleRes, nil
	}

	if err := r.db.Preload("Features").Preload("Users").Find(&roleOjbs).Error; err != nil {
		return nil, err
	}

	for _, role := range roleOjbs {
		roleRes = append(roleRes, models.ResRoleDetails{
			RoleID:     role.ID,
			RoleName:   role.Name,
			RoleLevel:  role.Level,
			NumberUser: int32(len(role.Users)),
		})
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: roleRes,
		TTL:   time.Minute * 10,
	}); err != nil {
		return nil, err
	}

	return roleRes, nil
}

func (r *roleRepository) GetAllDefault() ([]entities.Role, error) {
	var roleOjbs []entities.Role

	if err := r.db.Preload("Features").Preload("Users").Find(&roleOjbs).Error; err != nil {
		return nil, err
	}

	return roleOjbs, nil
}

func (r *roleRepository) Update(ctx context.Context, role *entities.Role) error {
	if err := r.db.Where("id=?", role.ID).Updates(&role).Error; err != nil {
		return err
	}

	cacheKey1 := fmt.Sprintf("role:%s", role.ID)
	if err := r.redisCache.Delete(ctx, cacheKey1); err != nil {
		return err
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey1,
		Value: role,
		TTL:   time.Minute * 10,
	}); err != nil {
		return err
	}

	var roleOjbs []entities.Role
	var roleRes []models.ResRoleDetails

	cacheKey2 := "roles_list"

	if err := r.redisCache.Delete(ctx, cacheKey2); err != nil {
		return err
	}

	if err := r.db.Preload("Users").Find(&roleOjbs).Error; err != nil {
		return err
	}

	for _, role := range roleOjbs {
		roleRes = append(roleRes, models.ResRoleDetails{
			RoleID:     role.ID,
			RoleName:   role.Name,
			RoleLevel:  role.Level,
			NumberUser: int32(len(role.Users)),
		})
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey2,
		Value: roleRes,
		TTL:   time.Minute * 10,
	}); err != nil {
		return err
	}

	return nil
}

func (r *roleRepository) Delete(id uuid.UUID, delBy uuid.UUID) error {
	if err := r.db.Model(&entities.Role{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_by": delBy,
	}).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&entities.Role{}, id).Error; err != nil {
		return err
	}
	return nil
}
