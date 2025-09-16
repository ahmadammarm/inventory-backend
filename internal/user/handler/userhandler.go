package handler

import (
	"github.com/ahmadammarm/inventory-backend/internal/user/dto"
	"github.com/ahmadammarm/inventory-backend/internal/user/model"
	"github.com/ahmadammarm/inventory-backend/internal/user/service"
	"github.com/ahmadammarm/inventory-backend/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"
)

type UserHandler struct {
	userService service.UserService
	validator   *validator.Validate
}

func (handler *UserHandler) SignupUser(context *fiber.Ctx) error {
	user := new(model.User)

	if err := context.BodyParser(user); err != nil {
		slog.Error("Error parsing request body: %v", err)
		return context.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid request body",
			Success: false,
			Code:    fiber.StatusBadRequest,
			Errors:  err.Error(),
		})
	}

	if err := handler.userService.SignupUser(user); err != nil {
		if err.Error() == "email already exists" {
			return context.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
				Message: "Email already exists",
				Success: false,
				Code:    fiber.StatusBadRequest,
				Errors:  "duplicate_email",
			})
		}

		slog.Error("Error signup user: %v", err)
		return context.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Message: "Internal Server Error",
			Success: false,
			Code:    fiber.StatusInternalServerError,
			Errors:  err.Error(),
		})
	}

	userResponse:= dto.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	return context.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Message: "User registered successfully",
		Success: true,
		Code:    fiber.StatusOK,
		Data:    userResponse,
	})
}

func (handler *UserHandler) UserRouters(router fiber.Router) {
	router.Post("/user/signup", handler.SignupUser)
}

func NewUserHandler(userService service.UserService, val *validator.Validate) *UserHandler {
	return &UserHandler{userService, val}
}
