package auth

import (
	"context"
	// "errors"
	// "fmt"
	// "time"

	"github.com/google/uuid"
	"github.com/lenardjombo/kairoapi/db/sqlc"
	// "golang.org/x/crypto/bcrypt"
)

// AuthService defines the contract for user authentication and management.
type AuthService interface {
	RegisterUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	LoginUser(ctx context.Context, email, password string) (db.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (db.User, error)
	ListUsers(ctx context.Context) ([]db.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, arg db.UpdateUserParams) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

// service implements AuthService
type service struct {
	repo UserRepository
}