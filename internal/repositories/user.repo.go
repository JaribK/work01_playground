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

type (
	UserRepository interface {
		Create(user *entities.User) error
		GetById(ctx context.Context, id uuid.UUID) (*entities.ResUserDTO, error)
		GetProfileUser(id uuid.UUID) (*entities.User, error)
		GetAvatarUserById(id uuid.UUID) (*entities.ResAvatar, error)
		GetRoleUserById(id uuid.UUID) (*entities.User, error)
		GetAllNoPage() ([]entities.ResUsersNoPage, error)
		GetAllWithPage(ctx context.Context, page, size int, roleId, isActive string, phoneNumber string, fullName string) ([]entities.ResAllUserDTOs, int64, error)
		Update(ctx context.Context, user *entities.User) error
		Delete(ctx context.Context, id uuid.UUID, deleteBy uuid.UUID) error
		GetUserByEmail(email string) (*entities.User, error)
		GetRoleByRoleId(id uuid.UUID) (*entities.Role, error)
		IsEmailExists(email string) (bool, error)
		IsPhoneExists(phone string) (bool, error)
		IsEmailExistsForUpdate(email string, id uuid.UUID) (bool, error)
		IsPhoneExistsForUpdate(phone string, id uuid.UUID) (bool, error)
		IsSuperAdministrator(id uuid.UUID) (bool, error)
		CheckThisUserHaveDataInAuth(userId uuid.UUID) (*entities.Authorization, bool, error)
		DeleteAuthAfterDeleteUser(id uuid.UUID, deleteBy uuid.UUID) error
	}

	userRepository struct {
		db         *gorm.DB
		redisCache *cache.Cache
	}
)

func NewUserRepository(db *gorm.DB, redisClient *redis.Client) UserRepository {
	c := cache.New(&cache.Options{
		Redis:      redisClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	return &userRepository{db: db, redisCache: c}
}

func (r *userRepository) Create(user *entities.User) error {

	if err := r.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetAllNoPage() ([]entities.ResUsersNoPage, error) {
	var users []entities.User
	var resUsers []entities.ResUsersNoPage
	if err := r.db.Preload("Role.Features").Find(&users).Error; err != nil {
		return nil, err
	}

	for _, user := range users {
		resUsers = append(resUsers, entities.ResUsersNoPage{
			ID:                 user.ID,
			FirstName:          user.FirstName,
			LastName:           user.LastName,
			Email:              user.Email,
			PhoneNumber:        user.PhoneNumber,
			Avatar:             returnNull(user.Avatar),
			TwoFactorEnabled:   *user.TwoFactorEnabled,
			TwoFactorVerified:  *user.TwoFactorVerified,
			TwoFactorToken:     returnNull(user.TwoFactorToken),
			TwoFactorAuthUrl:   returnNull(user.TwoFactorAuthUrl),
			RoleId:             user.RoleId,
			Role:               user.Role,
			ForgotPasswordCode: user.ForgotPasswordCode,
			IsActive:           *user.IsActive,
			CreatedAt:          user.CreatedAt,
			CreatedBy:          user.CreatedBy,
			UpdatedAt:          user.UpdatedAt,
			UpdatedBy:          user.UpdatedBy,
		})
	}

	return resUsers, nil
}

func returnNull(s string) *string {
	if s == "" {
		return nil
	}

	return &s
}

func (r *userRepository) GetRoleUserById(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	if err := r.db.Preload("Role").Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetProfileUser(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetById(ctx context.Context, id uuid.UUID) (*entities.ResUserDTO, error) {
	var user entities.User
	var userDTO entities.ResUserDTO
	var mergedPermissions []entities.FeatureDTODetails

	cacheKey := fmt.Sprintf("user:%s", id)

	if err := r.redisCache.Get(ctx, cacheKey, userDTO); err == nil {
		return &userDTO, nil
	}

	if err := r.db.Preload("Role.Features").Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}

	for _, feature := range user.Role.Features {
		var roleFeature entities.RoleFeature
		if err := r.db.Where("role_id = ? AND feature_id = ?", user.Role.ID, feature.ID).First(&roleFeature).Error; err != nil {
			return nil, err
		}

		mergedPermissions = append(mergedPermissions, entities.FeatureDTODetails{
			ID:           feature.ID,
			Name:         feature.Name,
			ParentMenuId: feature.ParentMenuId,
			MenuIcon:     returnNull(feature.MenuIcon),
			MenuNameTh:   feature.MenuNameTh,
			MenuNameEn:   feature.MenuNameEn,
			MenuSlug:     feature.MenuSlug,
			MenuSeqNo:    feature.MenuSeqNo,
			IsActive:     feature.IsActive,
			IsAdd:        roleFeature.IsAdd,
			IsView:       roleFeature.IsView,
			IsEdit:       roleFeature.IsEdit,
			IsDelete:     roleFeature.IsDelete,
		})
	}

	userDTO = entities.ResUserDTO{
		UserID:            user.ID,
		Email:             user.Email,
		FirstName:         user.FirstName,
		LastName:          user.LastName,
		PhoneNumber:       user.PhoneNumber,
		Avatar:            returnNull(user.Avatar),
		RoleId:            user.Role.ID,
		RoleName:          user.Role.Name,
		RoleLevel:         user.Role.Level,
		TwoFactorEnabled:  *user.TwoFactorEnabled,
		TwoFactorVerified: *user.TwoFactorVerified,
		TwoFactorToken:    returnNull(user.TwoFactorToken),
		TwoFactorAuthUrl:  returnNull(user.TwoFactorAuthUrl),
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

func (r *userRepository) GetAllWithPage(ctx context.Context, page, size int, roleId, isActive string, phoneNumber string, fullName string) ([]entities.ResAllUserDTOs, int64, error) {
	var users []entities.User
	var total int64
	var userDTOs []entities.ResAllUserDTOs

	cacheKey := "users_list"

	offset := (page - 1) * size

	query := r.db.Model(&entities.User{}).Preload("Role")
	if roleId != "" {
		query = query.Joins("JOIN roles ON roles.id = users.role_id").Where("roles.id = ?", roleId)
	}

	if isActive != "" {
		query = query.Where("users.is_active = ?", isActive)
	}

	if phoneNumber != "" {
		query = query.Where("users.phone_number LIKE ?", "%"+phoneNumber+"%")
	}

	if fullName != "" {
		query.Where("LOWER(CONCAT(first_name,' ',last_name)) LIKE LOWER(?)", "%"+fullName+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.redisCache.Get(ctx, cacheKey, userDTOs); err == nil {
		return userDTOs, total, nil
	}

	if err := query.Limit(size).Offset(offset).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	for _, user := range users {
		userDTOs = append(userDTOs, entities.ResAllUserDTOs{
			UserID:      user.ID,
			Email:       user.Email,
			FullName:    user.FirstName + " " + user.LastName,
			PhoneNumber: user.PhoneNumber,
			IsActive:    *user.IsActive,
			Avatar:      returnNull(user.Avatar),
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
	var userDTOs []entities.ResAllUserDTOs

	if err := r.redisCache.Delete(ctx, cacheKey2); err != nil {
		return err
	}

	if err := r.db.Preload("Role").Find(&users).Error; err != nil {
		return err
	}

	for _, user := range users {
		userDTOs = append(userDTOs, entities.ResAllUserDTOs{
			UserID:      user.ID,
			Email:       user.Email,
			FullName:    user.FirstName + " " + user.LastName,
			PhoneNumber: user.PhoneNumber,
			IsActive:    *user.IsActive,
			Avatar:      returnNull(user.Avatar),
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
	err := r.db.Model(&entities.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_by": deleteBy}).Error
	if err != nil {
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

func (r *userRepository) GetAvatarUserById(id uuid.UUID) (*entities.ResAvatar, error) {
	var user entities.User
	if err := r.db.Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}

	res := entities.ResAvatar{
		Avatar: user.Avatar,
	}

	return &res, nil
}

func (r *userRepository) GetRoleByRoleId(id uuid.UUID) (*entities.Role, error) {
	var roleOjb entities.Role

	if err := r.db.Preload("Features").Where("id=?", id).First(&roleOjb).Error; err != nil {
		return nil, err
	}

	return &roleOjb, nil
}

func (r *userRepository) CheckThisUserHaveDataInAuth(userId uuid.UUID) (*entities.Authorization, bool, error) {
	var auth entities.Authorization

	if err := r.db.Where("user_id=?", userId).First(&auth).Error; err != nil {
		return nil, false, err
	}

	return &auth, true, nil
}

func (r *userRepository) DeleteAuthAfterDeleteUser(id uuid.UUID, deleteBy uuid.UUID) error {
	if err := r.db.Model(&entities.Authorization{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_by": deleteBy,
	}).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&entities.Authorization{}, id).Error; err != nil {
		return err
	}

	return nil
}
