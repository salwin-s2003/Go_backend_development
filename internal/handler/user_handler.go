package handler

import (
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"task/internal/logger"
	"task/internal/service"
)

// Validator instance
var validate = validator.New()

// UserHandler handles HTTP requests for users
type UserHandler struct {
	service *service.UserService
}

// NewUserHandler creates a new handler
func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// Request structs
type CreateUserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=50"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

type UpdateUserRequest struct {
	Name string `json:"name" validate:"required,min=2,max=50"`
	Dob  string `json:"dob" validate:"required,datetime=2006-01-02"`
}

// Response struct
type UserResponse struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Dob  string `json:"dob"`
	Age  int    `json:"age"`
}

// Helper: calculate age from DOB
func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()
	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}

// Helper: map data (map[string]interface{}) to UserResponse
func mapToUserResponse(data map[string]interface{}) UserResponse {
	return UserResponse{
		ID:   data["id"].(int32),
		Name: data["name"].(string),
		Dob:  data["dob"].(string),
		Age:  data["age"].(int),
	}
}

// ----------------- Handlers -----------------

// POST /users
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Logger.Error("Failed to parse request body", logger.ZapError(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := validate.Struct(req); err != nil {
		logger.Logger.Warn("Validation failed", logger.ZapError(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.service.CreateUser(c.Context(), req.Name, req.Dob)
	if err != nil {
		logger.Logger.Error("Failed to create user", logger.ZapError(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Logger.Info("User created", logger.ZapInt32("id", user.ID), logger.ZapString("name", user.Name))
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Format("2006-01-02"),
	})
}

// PUT /users/:id
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logger.Logger.Warn("Invalid ID param", logger.ZapError(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}

	var req UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		logger.Logger.Error("Failed to parse request body", logger.ZapError(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if err := validate.Struct(req); err != nil {
		logger.Logger.Warn("Validation failed", logger.ZapError(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	user, err := h.service.UpdateUser(c.Context(), int32(id), req.Name, req.Dob)
	if err != nil {
		logger.Logger.Error("Failed to update user", logger.ZapError(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Logger.Info("User updated", logger.ZapInt32("id", user.ID), logger.ZapString("name", user.Name))
	return c.JSON(fiber.Map{
		"id":   user.ID,
		"name": user.Name,
		"dob":  user.Dob.Format("2006-01-02"),
	})
}

// GET /users/:id
func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logger.Logger.Warn("Invalid ID param", logger.ZapError(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}

	data, err := h.service.GetUserByID(c.Context(), int32(id))
	if err != nil {
		logger.Logger.Error("Failed to get user", logger.ZapError(err))
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	}

	resp := mapToUserResponse(data)
	logger.Logger.Info("User fetched", logger.ZapInt32("id", resp.ID))
	return c.JSON(resp)
}

// GET /users
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	users, err := h.service.ListUsers(c.Context())
	if err != nil {
		logger.Logger.Error("Failed to list users", logger.ZapError(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	var resp []UserResponse
	for _, u := range users {
		resp = append(resp, mapToUserResponse(u))
	}

	logger.Logger.Info("Users listed", logger.ZapInt("count", len(resp)))
	return c.JSON(resp)
}

// DELETE /users/:id
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		logger.Logger.Warn("Invalid ID param", logger.ZapError(err))
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid user ID"})
	}

	if err := h.service.DeleteUser(c.Context(), int32(id)); err != nil {
		logger.Logger.Error("Failed to delete user", logger.ZapError(err))
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	logger.Logger.Info("User deleted", logger.ZapInt("id", id)) // changed from ZapInt32
	return c.SendStatus(fiber.StatusNoContent)
}
