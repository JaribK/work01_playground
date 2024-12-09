package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"
	"work01/internal/entities"

	"github.com/go-redis/cache/v9"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type (
	RoleRepository interface {
		GetById(ctx context.Context, id uuid.UUID) (*entities.Role, []entities.FeatureInRole, error)
		GetAllDefault() ([]entities.Role, error)
		RoleNameIsAlreadyExitsUpdate(roleId uuid.UUID, roleName string) (bool, error)
		RoleNameIsAlreadyExits(roleName string) (bool, error)
		GetRoleLevelOfRoleUserByUserId(id uuid.UUID) (*entities.ResRoleLevel, error)
		GetAllFetureDefault() ([]entities.Feature, error)
		GetAllModify(ctx context.Context) ([]entities.ResAllRoleDetails, error)
		Create(role *entities.Role, roleFeatures []entities.RoleFeature) error
		Update(ctx context.Context, role *entities.Role, roleFeatures []entities.RoleFeature) error
		Delete(id uuid.UUID, delBy uuid.UUID) error
		CheckRoleHaveUserUsed(roleId uuid.UUID) (bool, error)
	}

	roleRepository struct {
		db         *gorm.DB
		redisCache *cache.Cache
	}
)

func NewRoleRepository(db *gorm.DB, redisClient *redis.Client) RoleRepository {
	c := cache.New(&cache.Options{
		Redis:      redisClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	return &roleRepository{db: db, redisCache: c}
}

func (r *roleRepository) Create(role *entities.Role, roleFeatures []entities.RoleFeature) error {
	if err := r.db.Create(&role).Error; err != nil {
		return err
	}

	for i := range roleFeatures {
		roleFeatures[i].ID = uuid.New()
		roleFeatures[i].RoleId = role.ID
	}

	if err := r.db.Create(&roleFeatures).Error; err != nil {
		return err
	}

	return nil
}

func (r *roleRepository) GetById(ctx context.Context, id uuid.UUID) (*entities.Role, []entities.FeatureInRole, error) {
	var roleOjb entities.Role
	var roleFeatureDetails []entities.FeatureInRole
	cacheKey := fmt.Sprintf("role:%s", id)

	if err := r.redisCache.Get(ctx, cacheKey, roleOjb); err == nil {
		return &roleOjb, roleFeatureDetails, nil
	}

	if err := r.db.Preload("Features").Where("id=?", id).First(&roleOjb).Error; err != nil {
		return nil, nil, err
	}

	for _, rf := range roleOjb.Features {
		var roleFeatures entities.RoleFeature
		if err := r.db.Where("role_id = ? AND feature_id = ?", roleOjb.ID, rf.ID).Preload("Feature").First(&roleFeatures).Error; err != nil {
			return nil, nil, err
		}
		roleFeatureDetails = append(roleFeatureDetails, entities.FeatureInRole{
			FeatureId:   roleFeatures.FeatureId,
			FeatureName: roleFeatures.Feature.Name,
			IsAdd:       roleFeatures.IsAdd,
			IsView:      roleFeatures.IsView,
			IsEdit:      roleFeatures.IsEdit,
			IsDelete:    roleFeatures.IsDelete,
		})
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: roleOjb,
		TTL:   time.Minute * 10,
	}); err != nil {
		return nil, nil, err
	}

	return &roleOjb, roleFeatureDetails, nil
}

func (r *roleRepository) GetAllModify(ctx context.Context) ([]entities.ResAllRoleDetails, error) {
	var roleOjbs []entities.Role
	var roleRes []entities.ResAllRoleDetails

	cacheKey := "roles_list"

	if err := r.redisCache.Get(ctx, cacheKey, roleOjbs); err == nil {
		return roleRes, nil
	}

	if err := r.db.Preload("Features").Preload("Users").Find(&roleOjbs).Error; err != nil {
		return nil, err
	}

	for _, role := range roleOjbs {
		roleRes = append(roleRes, entities.ResAllRoleDetails{
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

func (r *roleRepository) GetAllFetureDefault() ([]entities.Feature, error) {
	var features []entities.Feature

	if err := r.db.Find(&features).Error; err != nil {
		return nil, err
	}

	return features, nil
}

func (r *roleRepository) Update(ctx context.Context, role *entities.Role, roleFeatures []entities.RoleFeature) error {
	if err := r.db.Where("id=?", role.ID).Updates(&role).Error; err != nil {
		return err
	}

	for _, rf := range roleFeatures {
		if err := r.db.Where("role_id = ? AND feature_id = ?", role.ID, rf.FeatureId).Updates(&rf).Error; err != nil {
			return err
		}
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
	var roleRes []entities.ResAllRoleDetails

	cacheKey2 := "roles_list"

	if err := r.redisCache.Delete(ctx, cacheKey2); err != nil {
		return err
	}

	if err := r.db.Preload("Users").Find(&roleOjbs).Error; err != nil {
		return err
	}

	for _, role := range roleOjbs {
		roleRes = append(roleRes, entities.ResAllRoleDetails{
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

func (r *roleRepository) RoleNameIsAlreadyExitsUpdate(roleId uuid.UUID, roleName string) (bool, error) {
	var role entities.Role
	if err := r.db.Where("name=? AND id != ?", roleName, roleId).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *roleRepository) RoleNameIsAlreadyExits(roleName string) (bool, error) {
	var role entities.Role
	if err := r.db.Where("name=?", roleName).First(&role).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

func (r *roleRepository) GetRoleLevelOfRoleUserByUserId(id uuid.UUID) (*entities.ResRoleLevel, error) {
	var user entities.User
	if err := r.db.Preload("Role").Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}

	res := entities.ResRoleLevel{
		RoleLevel: user.Role.Level,
	}

	return &res, nil
}

func (r *roleRepository) CheckRoleHaveUserUsed(roleId uuid.UUID) (bool, error) {
	var roleOjb entities.Role
	if err := r.db.Preload("Features").Preload("Users").Where("id=?", roleId).Find(&roleOjb).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}

	if len(roleOjb.Users) > 0 {
		return true, nil
	}

	return false, nil
}
