package repositories

import (
	"context"
	"fmt"
	"log"
	"time"
	"work01/internal/entities"

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
	GetById(ctx context.Context, id uuid.UUID) (*entities.User, error)
	GetAll(ctx context.Context, page, size int, roleId, isActive string) ([]entities.User, int64, error)
	Update(user *entities.User) error
	Delete(id uuid.UUID, deleteBy uuid.UUID) error
	GetUserByEmail(email string) (*entities.User, error)
	IsEmailExists(email string) (bool, error)
	IsPhoneExists(phone string) (bool, error)
	IsEmailExistsForUpdate(email string, id uuid.UUID) (bool, error)
	IsPhoneExistsForUpdate(phone string, id uuid.UUID) (bool, error)
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
	err = r.db.Create(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetById(ctx context.Context, id uuid.UUID) (*entities.User, error) {
	var user entities.User

	cacheKey := fmt.Sprintf("user:%s", id)
	if err := r.redisCache.Get(ctx, cacheKey, &user); err == nil {
		return &user, nil
	}

	err := r.db.Preload("Role.Permissions.Feature").Where("id=?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	_ = r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: user,
		TTL:   time.Minute * 10,
	})

	return &user, nil
}

func (r *userRepository) GetAll(ctx context.Context, page, size int, roleId, isActive string) ([]entities.User, int64, error) {
	var users []entities.User
	var total int64

	cacheKey := fmt.Sprintf("users:page:%d:size:%d:role:%s:isActive:%s", page, size, roleId, isActive)

	if err := r.redisCache.Get(ctx, cacheKey, &users); err == nil {
		return users, total, nil
	}

	offset := (page - 1) * size
	err := r.db.Model(&entities.User{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	query := r.db.Model(&entities.User{}).Preload("Role")
	if roleId != "" {
		query = query.Joins("JOIN roles ON roles.id = users.role_id").Where("roles.id = ?", roleId)
	}

	if isActive != "" {
		query = query.Where("users.is_active = ?", isActive)
	}

	err = query.Limit(size).Offset(offset).Find(&users).Error
	if err != nil {
		return nil, 0, err
	}

	err = r.redisCache.Set(&cache.Item{
		Ctx:   ctx,
		Key:   cacheKey,
		Value: users,
		TTL:   time.Minute * 10, // Cache expiration time (TTL)
	})

	if err != nil {
		log.Fatal(err)
	}

	return users, total, nil
}

func (r *userRepository) Update(user *entities.User) error {
	err := r.db.Where("id=?", user.ID).Updates(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(id uuid.UUID, deleteBy uuid.UUID) error {
	err := r.db.Model(&entities.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"deleted_by": deleteBy,
	}).Error
	if err != nil {
		return err
	}

	err = r.db.Delete(&entities.User{}, id).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("email=?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) IsEmailExists(email string) (bool, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *userRepository) IsPhoneExists(phone string) (bool, error) {
	var user entities.User
	err := r.db.Where("phone_number = ?", phone).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *userRepository) IsEmailExistsForUpdate(email string, id uuid.UUID) (bool, error) {
	var user entities.User
	err := r.db.Where("email = ?", email).Not("id=?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, nil
	}
	return true, nil
}

func (r *userRepository) IsPhoneExistsForUpdate(phone string, id uuid.UUID) (bool, error) {
	var user entities.User
	err := r.db.Where("phone_number = ?", phone).Not("id=?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, nil
	}

	return true, nil
}
