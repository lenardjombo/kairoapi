package auth

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	// "errors"
	// "fmt"
	// "time"

	"github.com/google/uuid"
	"github.com/lenardjombo/kairoapi/db/sqlc"
	"github.com/lenardjombo/kairoapi/models"
	"github.com/lenardjombo/kairoapi/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

// AuthService defines the contract for user authentication and management.
type AuthService interface {
	RegisterUser(ctx context.Context, arg models.User) (*db.User, error)
	LoginUser(ctx context.Context, email, password string) (*db.User, error)
	// GetUserById(ctx context.Context, id uuid.UUID) (db.User, error)
	// ListUsers(ctx context.Context) ([]db.User, error)
	// UpdateUser(ctx context.Context, id uuid.UUID, arg db.UpdateUserParams) error
	// DeleteUser(ctx context.Context, id uuid.UUID) error
}

// service implements AuthService
type service struct {
	repo UserRepository
}

// Initialise NewAuth
func NewAuthService(repo UserRepository) AuthService {
	return &service{repo: repo}
}

func (s *service) RegisterUser(ctx context.Context, arg models.User) (*db.User, error) {
	// Derives a new context from the incoming 'ctx' with a 1-second timeout.
	// This timeout ensures that the subsequent database `CreateUser` call
	// does not block indefinitely if the database is slow or unreachable.
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)

	defer cancel()
	//Hash the password here
	hashedpassword, err := utils.HashedPassword(arg.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password : %w", err)
	}
	arg.Password = hashedpassword

	// Create a new user in the database
	if arg.ID == uuid.Nil {
		arg.ID = uuid.New()
	}
	if arg.CreatedAt.IsZero() {
		arg.CreatedAt = time.Now()
	}
	if arg.UpdatedAt.IsZero() {
		arg.UpdatedAt = time.Now()
	}

	//validate the email
	err = utils.ValidateEmail(arg.Email)
	if err != nil {
		return nil, err
	}

	u := db.CreateUserParams{
		ID:        arg.ID,
		Email:     arg.Email,
		Password:  arg.Password, //Hashed password
		CreatedAt: arg.CreatedAt,
		UpdatedAt: arg.UpdatedAt,
	}
	r, err := s.repo.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}
	r.Password = ""
	return &r, nil
}

func (s *service) LoginUser(ctx context.Context, email, password string) (*db.User, error) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)

	defer cancel()
	// Validate email format
	err := utils.ValidateEmail(email)
	if err != nil {
		return nil,fmt.Errorf("invalid email format : %w",err)
	}

	foundUser,err := s.repo.GetUserByEmail(ctx,email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil,fmt.Errorf("invalid credentials ")
		}
		return nil,fmt.Errorf("login failed to internal error : %w",err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
    if err != nil {
		if err == bcrypt.ErrMismatchedHashAndPassword {
			return nil,fmt.Errorf("invalid credentials : ")
		}
		return nil,fmt.Errorf("login failed invalid credentials : %w",err)
	}

	//successful login
	foundUser.Password = ""

	return &foundUser, nil
}

//Todos : Refactor (User DTOs for proper separation of concerns)