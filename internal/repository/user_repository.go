package repository

import (
	"context"

	db "task/db/sqlc"
)

type UserRepository struct {
	q *db.Queries
}

func NewUserRepository(q *db.Queries) *UserRepository {
	return &UserRepository{q: q}
}

// Create a new user
func (r *UserRepository) CreateUser(ctx context.Context, name string, dob string) (db.User, error) {
	// Parse dob into time.Time
	parsedDob, err := ParseDate(dob)
	if err != nil {
		return db.User{}, err
	}

	params := db.CreateUserParams{
		Name: name,
		Dob:  parsedDob,
	}

	return r.q.CreateUser(ctx, params)
}

// Get user by ID
func (r *UserRepository) GetUserByID(ctx context.Context, id int32) (db.User, error) {
	return r.q.GetUserByID(ctx, id)
}

// List all users
func (r *UserRepository) ListUsers(ctx context.Context) ([]db.User, error) {
	return r.q.ListUsers(ctx)
}

// Update user
func (r *UserRepository) UpdateUser(ctx context.Context, id int32, name, dob string) (db.User, error) {
	parsedDob, err := ParseDate(dob)
	if err != nil {
		return db.User{}, err
	}

	params := db.UpdateUserParams{
		ID:   id,
		Name: name,
		Dob:  parsedDob,
	}

	return r.q.UpdateUser(ctx, params)
}

// Delete user
func (r *UserRepository) DeleteUser(ctx context.Context, id int32) error {
	return r.q.DeleteUser(ctx, id)
}
