package service

import (
	"context"
	"time"
	"task/internal/repository"
	db "task/db/sqlc"
	"errors"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// Calculate age dynamically from DOB
func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

// Create user
func (s *UserService) CreateUser(ctx context.Context, name, dob string) (db.User, error) {
	if name == "" {
		return db.User{}, errors.New("name is required")
	}
	return s.repo.CreateUser(ctx, name, dob)
}

// Get user by ID
func (s *UserService) GetUserByID(ctx context.Context, id int32) (map[string]interface{}, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Add dynamic age
	result := map[string]interface{}{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Format("2006-01-02"),
		"age":  calculateAge(user.Dob),
	}

	return result, nil
}

// List users
func (s *UserService) ListUsers(ctx context.Context) ([]map[string]interface{}, error) {
	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, u := range users {
		result = append(result, map[string]interface{}{
			"id":   u.ID,
			"name": u.Name,
			"dob":  u.Dob.Format("2006-01-02"),
			"age":  calculateAge(u.Dob),
		})
	}

	return result, nil
}

// Update user
func (s *UserService) UpdateUser(ctx context.Context, id int32, name, dob string) (db.User, error) {
	if name == "" {
		return db.User{}, errors.New("name is required")
	}
	return s.repo.UpdateUser(ctx, id, name, dob)
}

// Delete user
func (s *UserService) DeleteUser(ctx context.Context, id int32) error {
	return s.repo.DeleteUser(ctx, id)
}
