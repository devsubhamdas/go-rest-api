package repository

import (
	"context"

	"github.com/Subham-Das-98/go-rest-api/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Save(user).Error
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, id).Error
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]models.User, error) {
	var users []models.User
	err := r.db.WithContext(ctx).Find(&users).Error
	return users, err
}
