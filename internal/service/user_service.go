package service

import (
	"context"
	"errors"
	"strings"

	"github.com/Subham-Das-98/go-rest-api/internal/models"
	"github.com/Subham-Das-98/go-rest-api/internal/utils/pwd"
	"github.com/google/uuid"
	// "gorm.io/gorm"
)

type UserRepository interface {
	Create(context.Context, *models.User) error
	Update(context.Context, *models.User) error
	Delete(context.Context, uuid.UUID) error
	GetByID(context.Context, uuid.UUID) (*models.User, error)
	GetByEmail(context.Context, string) (*models.User, error)
	GetAll(context.Context) ([]models.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) *UserService {
	return &UserService{
		repo: r,
	}
}

func (s *UserService) CreateUser(ctx context.Context, name, email, password string) error {
	email = strings.TrimSpace(strings.ToLower(email))

	if name == "" {
		return errors.New("name is required")
	}

	if email == "" {
		return errors.New("email is required")
	}

	if password == "" {
		return errors.New("password is required")
	}

	// existingUser, err := s.repo.GetByEmail(ctx, email)

	// if err == nil && existingUser != nil {
	// 	return errors.New("email already exists")
	// }

	// if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
	// 	return err
	// }

	hash, err := pwd.RawToHash(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: hash,
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) UpdateUser(ctx context.Context, id, name, email string) error {
	if id == "" {
		return errors.New("id path value missing")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	email = strings.TrimSpace(strings.ToLower(email))

	if name == "" {
		return errors.New("name is required")
	}

	if email == "" {
		return errors.New("email is required")
	}

	user := &models.User{
		ID:    uid,
		Name:  name,
		Email: email,
	}

	err = s.repo.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, id string) error {
	if id == "" {
		return errors.New("id path value missing")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return err
	}

	err = s.repo.Delete(ctx, uid)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) GetUser(ctx context.Context, id string) (*models.User, error) {
	if id == "" {
		return nil, errors.New("id path value missing")
	}

	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return s.repo.GetByID(ctx, uid)
}

func (s *UserService) GetUsers(ctx context.Context) ([]models.User, error) {
	return s.repo.GetAll(ctx)
}
