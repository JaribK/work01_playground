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
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepository struct {
	db         *gorm.DB
	redisCache *cache.Cache
}

type UserRepository interface {
	Create(user *entities.User) error
	GetById(ctx context.Context, id uuid.UUID) (*models.ResUserDTO, error)
	GetRoleUserById(id uuid.UUID) (*entities.User, error)
	GetAllNoPage() ([]models.ResUsersNoPage, error)
	GetAllWithPage(ctx context.Context, page, size int, roleId, isActive string) ([]models.ResAllUserDTOs, int64, error)
	Update(ctx context.Context, user *entities.User) error
	Delete(ctx context.Context, id uuid.UUID, deleteBy uuid.UUID) error
	GetUserByEmail(email string) (*entities.User, error)
	IsEmailExists(email string) (bool, error)
	IsPhoneExists(phone string) (bool, error)
	IsEmailExistsForUpdate(email string, id uuid.UUID) (bool, error)
	IsPhoneExistsForUpdate(phone string, id uuid.UUID) (bool, error)
	IsSuperAdministrator(id uuid.UUID) (bool, error)
}

func NewUserRepository(db *gorm.DB, redisClient *redis.Client) UserRepository {
	c := cache.New(&cache.Options{
		Redis:      redisClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	return &userRepository{db: db, redisCache: c}
}

func (r *userRepository) Create(user *entities.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	if err = r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetAllNoPage() ([]models.ResUsersNoPage, error) {
	var users []entities.User
	var resUsers []models.ResUsersNoPage
	if err := r.db.Preload("Role.Features").Find(&users).Error; err != nil {
		return nil, err
	}

	for _, user := range users {
		resUsers = append(resUsers, models.ResUsersNoPage{
			ID:                 user.ID,
			FirstName:          user.FirstName,
			LastName:           user.LastName,
			Email:              user.Email,
			PhoneNumber:        user.PhoneNumber,
			Avatar:             user.Avatar,
			TwoFactorEnabled:   user.TwoFactorEnabled,
			TwoFactorVerified:  user.TwoFactorVerified,
			TwoFactorToken:     *user.TwoFactorToken,
			TwoFactorAuthUrl:   *user.TwoFactorAuthUrl,
			RoleId:             user.RoleId,
			Role:               user.Role,
			ForgotPasswordCode: user.ForgotPasswordCode,
			IsActive:           user.IsActive,
			CreatedAt:          user.CreatedAt,
			CreatedBy:          user.CreatedBy,
			UpdatedAt:          user.UpdatedAt,
			UpdatedBy:          user.UpdatedBy,
		})
	}

	return resUsers, nil
}

func (r *userRepository) GetRoleUserById(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	if err := r.db.Preload("Role").Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetById(ctx context.Context, id uuid.UUID) (*models.ResUserDTO, error) {
	var user entities.User
	var userDTO models.ResUserDTO
	var roleFeature entities.RoleFeature
	var mergedPermissions []models.FeatureDTODetails

	cacheKey := fmt.Sprintf("user:%s", id)

	if err := r.redisCache.Get(ctx, cacheKey, userDTO); err == nil {
		return &userDTO, nil
	}

	if err := r.db.Preload("Role.Features").Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}

	if err := r.db.Where("role_id=?", user.Role.ID).First(&roleFeature).Error; err != nil {
		return nil, err
	}

	for _, Feature := range user.Role.Features {
		mergedPermissions = append(mergedPermissions, models.FeatureDTODetails{
			ID:           Feature.ID,
			Name:         Feature.Name,
			ParentMenuId: Feature.ParentMenuId,
			MenuIcon:     Feature.MenuIcon,
			MenuNameTh:   Feature.MenuNameTh,
			MenuNameEn:   Feature.MenuNameEn,
			MenuSlug:     Feature.MenuSlug,
			MenuSeqNo:    Feature.MenuSeqNo,
			IsActive:     Feature.IsActive,
			IsAdd:        roleFeature.IsAdd,
			IsView:       roleFeature.IsView,
			IsEdit:       roleFeature.IsEdit,
			IsDelete:     roleFeature.IsDelete,
		})
	}

	userDTO = models.ResUserDTO{
		UserID:            user.ID,
		Email:             user.Email,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		PhoneNumber:       user.PhoneNumber,
		Avatar:            user.Avatar,
		RoleId:            user.Role.ID,
		RoleName:          user.Role.Name,
		RoleLevel:         user.Role.Level,
		TwoFactorEnabled:  user.TwoFactorEnabled,
		TwoFactorVerified: user.TwoFactorVerified,
		TwoFactorAuthUrl:  *user.TwoFactorAuthUrl,
		TwoFactorToken:    *user.TwoFactorToken,
		Features:          mergedPermissions,
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: user,
		TTL:   time.Minute * 10,
	}); err != nil {
		return nil, err
	}

	return &userDTO, nil
}

func (r *userRepository) GetAllWithPage(ctx context.Context, page, size int, roleId, isActive string) ([]models.ResAllUserDTOs, int64, error) {
	var users []entities.User
	var total int64
	var userDTOs []models.ResAllUserDTOs

	cacheKey := "users_list"

	offset := (page - 1) * size
	if err := r.db.Model(&entities.User{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query := r.db.Model(&entities.User{}).Preload("Role")
	if roleId != "" {
		query = query.Joins("JOIN roles ON roles.id = users.role_id").Where("roles.id = ?", roleId)
	}

	if isActive != "" {
		query = query.Where("users.is_active = ?", isActive)
	}

	if err := r.redisCache.Get(ctx, cacheKey, userDTOs); err == nil {
		return userDTOs, total, nil
	}

	if err := query.Limit(size).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	for _, user := range users {
		userDTOs = append(userDTOs, models.ResAllUserDTOs{
			UserID:      user.ID,
			Email:       user.Email,
			FullName:    user.FirstName + " " + user.LastName,
			PhoneNumber: user.PhoneNumber,
			IsActive:    user.IsActive,
			Avatar:      user.Avatar,
			RoleName:    user.Role.Name,
		})
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: userDTOs,
		TTL:   time.Minute * 10,
	}); err != nil {
		return nil, 0, err
	}

	return userDTOs, total, nil
}

func (r *userRepository) Update(ctx context.Context, user *entities.User) error {
	if err := r.db.Where("id=?", user.ID).Updates(&user).Error; err != nil {
		return err
	}

	cacheKey1 := fmt.Sprintf("user:%s", user.ID)
	if err := r.redisCache.Delete(ctx, cacheKey1); err != nil {
		return nil
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey1,
		Value: user,
		TTL:   time.Minute * 10,
	}); err != nil {
		return err
	}

	//query again
	cacheKey2 := "users_list"
	var users []entities.User
	var userDTOs []models.ResAllUserDTOs

	if err := r.redisCache.Delete(ctx, cacheKey2); err != nil {
		return err
	}

	if err := r.db.Preload("Role").Find(&users).Error; err != nil {
		return err
	}

	for _, user := range users {
		userDTOs = append(userDTOs, models.ResAllUserDTOs{
			UserID:      user.ID,
			Email:       user.Email,
			FullName:    user.FirstName + " " + user.LastName,
			PhoneNumber: user.PhoneNumber,
			IsActive:    user.IsActive,
			Avatar:      user.Avatar,
			RoleName:    user.Role.Name,
		})
	}

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey2,
		Value: userDTOs,
		TTL:   time.Minute * 10,
	}); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(ctx context.Context, id uuid.UUID, deleteBy uuid.UUID) error {
	if err := r.db.Model(&entities.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_by": deleteBy,
	}).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&entities.User{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) IsEmailExists(email string) (bool, error) {
	var user entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *userRepository) IsPhoneExists(phone string) (bool, error) {
	var user entities.User
	if err := r.db.Where("phone_number = ?", phone).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *userRepository) IsEmailExistsForUpdate(email string, id uuid.UUID) (bool, error) {
	var user entities.User
	if err := r.db.Where("email = ?", email).Not("id=?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, nil
	}
	return true, nil
}

func (r *userRepository) IsPhoneExistsForUpdate(phone string, id uuid.UUID) (bool, error) {
	var user entities.User
	if err := r.db.Where("phone_number = ?", phone).Not("id=?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, nil
	}
	return true, nil
}

func (r *userRepository) IsSuperAdministrator(id uuid.UUID) (bool, error) {
	var user entities.User
	if err := r.db.Preload("Role").Where("id=?", id).First(&user).Error; err != nil {
		return false, nil
	}

	if user.Role.Name == "Super Administrator" {
		return true, nil
	} else {
		return false, nil
	}
}
