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

type authorizationRepository struct {
	db         *gorm.DB
	redisCache *cache.Cache
}

type AuthorizationRepository interface {
	Create(auth *entities.Authorization) error
	GetById(id uuid.UUID) (*entities.Authorization, error)
	GetAll() ([]entities.Authorization, error)
	Update(auth *entities.Authorization) error
	Delete(id uuid.UUID, deleteBy uuid.UUID) error
	GetUserById(id uuid.UUID) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	CheckAuthorizationByUserID(id uuid.UUID) bool
	GetAuthorizationByUserID(id uuid.UUID) (*entities.Authorization, error)
	DeleteAuthorizationByUserId(id uuid.UUID, tokenString string, ttl time.Duration) error
	GetAuthorizationByRefreshToken(refreshToken string) (*entities.Authorization, error)
}

func NewAuthorizationRepository(db *gorm.DB, redisClient *redis.Client) AuthorizationRepository {
	c := cache.New(&cache.Options{
		Redis:      redisClient,
		LocalCache: cache.NewTinyLFU(1000, time.Minute),
	})
	return &authorizationRepository{db: db, redisCache: c}
}

func (r *authorizationRepository) Create(auth *entities.Authorization) error {
	if err := r.db.Create(&auth).Error; err != nil {
		return err
	}

	return nil
}

func (r *authorizationRepository) GetById(id uuid.UUID) (*entities.Authorization, error) {
	var auth entities.Authorization
	if err := r.db.First(&auth, id).Error; err != nil {
		return nil, err
	}

	return &auth, nil
}

func (r *authorizationRepository) GetAll() ([]entities.Authorization, error) {
	var auths []entities.Authorization
	if err := r.db.Find(&auths).Error; err != nil {
		return nil, err
	}

	return auths, nil
}

func (r *authorizationRepository) Update(auth *entities.Authorization) error {
	if err := r.db.Where("id=?", auth.ID).Updates(&auth).Error; err != nil {
		return err
	}

	return nil
}

// delete with path
func (r *authorizationRepository) Delete(id uuid.UUID, deleteBy uuid.UUID) error {
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

func (r *authorizationRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Preload("Role").Where("email=?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authorizationRepository) GetUserById(id uuid.UUID) (*entities.User, error) {
	var user entities.User
	if err := r.db.Preload("Role.Permissions.Feature").Where("id=?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *authorizationRepository) CheckAuthorizationByUserID(id uuid.UUID) bool {
	var auth entities.Authorization
	if err := r.db.Where("user_id = ?", id).First(&auth).Error; err != nil {
		return false
	}
	return true
}

func (r *authorizationRepository) GetAuthorizationByUserID(id uuid.UUID) (*entities.Authorization, error) {
	var auth entities.Authorization
	if err := r.db.Where("user_id = ?", id).First(&auth).Error; err != nil {
		return nil, err
	}
	return &auth, nil
}

func (r *authorizationRepository) GetAuthorizationByRefreshToken(refreshToken string) (*entities.Authorization, error) {
	var auth entities.Authorization
	if err := r.db.Where("refresh_token = ?", refreshToken).First(&auth).Error; err != nil {
		return nil, err
	}
	return &auth, nil
}

// for logout
func (r *authorizationRepository) DeleteAuthorizationByUserId(id uuid.UUID, tokenString string, ttl time.Duration) error {
	cacheKey := fmt.Sprintf("blocked:%s", tokenString)

	if err := r.redisCache.Set(&cache.Item{
		Ctx:   context.Background(),
		Key:   cacheKey,
		Value: tokenString,
		TTL:   ttl,
	}); err != nil {
		return fmt.Errorf("failed to set token in Redis: %w", err)
	}

	if err := r.db.Model(&entities.Authorization{}).Where("user_id = ?", id).Updates(map[string]interface{}{
		"access_token":  "",
		"refresh_token": "",
	}).Error; err != nil {
		return err
	}

	return nil
}
