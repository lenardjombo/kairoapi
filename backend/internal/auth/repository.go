package auth

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lenardjombo/kairoapi/db/sqlc"
)

// UserRepository defines the contract for user-related database operations.
type UserRepository interface {
	CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error)
	GetUserByEmail(ctx context.Context, email string) (db.User, error)
	GetUserById(ctx context.Context, id uuid.UUID) (db.User, error)
	ListUsers(ctx context.Context) ([]db.User, error)
	UpdateUser(ctx context.Context, arg db.UpdateUserParams) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}


type userRepository struct {
	q *db.Queries
}

// NewUserRepository creates a new instance of userRepository.
func NewUserRepository(q *db.Queries) UserRepository {
	return &userRepository{q: q}
}

func (r *userRepository) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	return r.q.CreateUser(ctx, arg)
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	
	return r.q.GetUserByEmail(ctx, email)
}

func (r *userRepository) GetUserById(ctx context.Context, id uuid.UUID) (db.User, error) {
	return r.q.GetUserById(ctx, id)
}

func (r *userRepository) ListUsers(ctx context.Context) ([]db.User, error) {
	return r.q.ListUsers(ctx)
}

func (r *userRepository) UpdateUser(ctx context.Context, arg db.UpdateUserParams) error {
	affectedrows,err := r.q.UpdateUser(ctx,arg)
	if err != nil {
		return fmt.Errorf("no rows affected")
	}
	if affectedrows == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func (r *userRepository) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return r.q.DeleteUser(ctx, id)
}
