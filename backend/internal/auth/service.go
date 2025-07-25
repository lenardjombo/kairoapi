//github.com/lenardjombo/kairoapi/auth

package auth

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lenardjombo/kairoapi/db/sqlc"
	"github.com/lenardjombo/kairoapi/models"
	"github.com/lenardjombo/kairoapi/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(ctx context.Context, req models.CreateUserReq) (*models.CreateUserRes, error)
	LoginUser(ctx context.Context, req models.LoginUserReq) (*models.LoginUserRes, error)
}

type service struct {
	repo UserRepository
}

func NewAuthService(repo UserRepository) AuthService {
	return &service{repo: repo}
}

func (s *service) RegisterUser(ctx context.Context, req models.CreateUserReq) (*models.CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// Validate email
	if err := utils.ValidateEmail(req.Email); err != nil {
		return nil, fmt.Errorf("invalid email format: %w", err)
	}

	// Hash password
	hashedPassword, err := utils.HashedPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	userID := uuid.New()
	now := time.Now()

	params := db.CreateUserParams{
		ID:        userID,
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		CreatedAt: now,
		UpdatedAt: now,
	}

	createdUser, err := s.repo.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	res := &models.CreateUserRes{
		ID:       createdUser.ID.String(),
		Username: createdUser.Username,
		Email:    createdUser.Email,
	}
	return res, nil
}

func (s *service) LoginUser(ctx context.Context, req models.LoginUserReq) (*models.LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := utils.ValidateEmail(req.Email); err != nil {
		return nil, fmt.Errorf("invalid email: %w", err)
	}

	foundUser, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("invalid credentials")
		}
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(req.Password)); err != nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	res := &models.LoginUserRes{
		ID:       foundUser.ID.String(),
		Username: foundUser.Username,
	}
	return res, nil
}
