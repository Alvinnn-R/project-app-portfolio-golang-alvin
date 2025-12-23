package repository

import (
	"context"
	"errors"
	"session-19/database"
	"session-19/model"

	"go.uber.org/zap"
)

// UserRepositoryInterface defines the interface for user repository
type UserRepositoryInterface interface {
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetByID(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) error
}

// UserRepository implements UserRepositoryInterface
type UserRepository struct {
	db  database.PgxIface
	log *zap.Logger
}

// NewUserRepository creates a new user repository
func NewUserRepository(db database.PgxIface, log *zap.Logger) UserRepositoryInterface {
	return &UserRepository{
		db:  db,
		log: log,
	}
}

// GetByEmail retrieves a user by email
func (r *UserRepository) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	query := `SELECT id, email, password, name, role, created_at, updated_at FROM users WHERE email = $1`

	var user model.User
	err := r.db.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		r.log.Error("Failed to get user by email", zap.Error(err), zap.String("email", email))
		return nil, errors.New("user not found")
	}

	return &user, nil
}

// GetByID retrieves a user by ID
func (r *UserRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	query := `SELECT id, email, password, name, role, created_at, updated_at FROM users WHERE id = $1`

	var user model.User
	err := r.db.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Password,
		&user.Name,
		&user.Role,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		r.log.Error("Failed to get user by ID", zap.Error(err), zap.Int64("id", id))
		return nil, errors.New("user not found")
	}

	return &user, nil
}

// Create creates a new user
func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (email, password, name, role, created_at, updated_at) 
			  VALUES ($1, $2, $3, $4, NOW(), NOW()) RETURNING id`

	err := r.db.QueryRow(ctx, query,
		user.Email,
		user.Password,
		user.Name,
		user.Role,
	).Scan(&user.ID)

	if err != nil {
		r.log.Error("Failed to create user", zap.Error(err))
		return errors.New("failed to create user")
	}

	return nil
}

// Update updates a user
func (r *UserRepository) Update(ctx context.Context, user *model.User) error {
	query := `UPDATE users SET email = $1, name = $2, updated_at = NOW() WHERE id = $3`

	_, err := r.db.Exec(ctx, query, user.Email, user.Name, user.ID)
	if err != nil {
		r.log.Error("Failed to update user", zap.Error(err))
		return errors.New("failed to update user")
	}

	return nil
}
