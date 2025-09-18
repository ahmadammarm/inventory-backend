package handler

import (
	"github.com/ahmadammarm/inventory-backend/internal/user/dto"
	"github.com/ahmadammarm/inventory-backend/internal/user/model"
	"github.com/ahmadammarm/inventory-backend/internal/user/service"
	"github.com/ahmadammarm/inventory-backend/middlewares"
	"github.com/ahmadammarm/inventory-backend/pkg/inputvalidator"
	"github.com/ahmadammarm/inventory-backend/pkg/response"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService service.UserService
	validator   *validator.Validate
}

func (handler *UserHandler) SignupUser(context *fiber.Ctx) error {

	user := new(model.User)

	if err := context.BodyParser(user); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid request body",
			Success: false,
			Code:    fiber.StatusBadRequest,
			Errors:  err.Error(),
		})
	}


    if err := handler.validator.Struct(user); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Validation failed",
			Success: false,
			Code:    fiber.StatusBadRequest,
			Errors:  inputvalidator.TranslateErrorMessage(err),
		})
	}


    if err := handler.userService.SignupUser(user); err != nil {

        if inputvalidator.IsDuplicateEntryError(err) || err.Error() == "email already exists" {
			return context.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
				Message: "Email already exists",
				Success: false,
				Code:    fiber.StatusBadRequest,
				Errors:  map[string]string{"email": "Email already exists"},
			})
		}


        return context.Status(fiber.StatusInternalServerError).JSON(response.ErrorResponse{
			Message: "Failed to register user",
			Success: false,
			Code:    fiber.StatusInternalServerError,
			Errors:  err.Error(),
		})
	}


    userResponse := dto.UserResponse{
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



func (handler *UserHandler) SigninUser(context *fiber.Ctx) error {
	userReq := new(model.User)

	if err := context.BodyParser(userReq); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Invalid request body",
			Success: false,
			Code:    fiber.StatusBadRequest,
			Errors:  err.Error(),
		})
	}

	if err := handler.validator.Struct(userReq); err != nil {
		return context.Status(fiber.StatusBadRequest).JSON(response.ErrorResponse{
			Message: "Validation failed",
			Success: false,
			Code:    fiber.StatusBadRequest,
			Errors:  inputvalidator.TranslateErrorMessage(err),
		})
	}

	user, token, err := handler.userService.SigninUser(userReq)
	if err != nil {
		return context.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
			Message: "Invalid email or password",
			Success: false,
			Code:    fiber.StatusUnauthorized,
			Errors:  err.Error(),
		})
	}

	userResponse := dto.UserJWTResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Token: token,
	}

	return context.Status(fiber.StatusOK).JSON(response.SuccessResponse{
		Message: "User signed in successfully",
		Success: true,
		Code:    fiber.StatusOK,
		Data:    userResponse,
	})
}

func (handler *UserHandler) UserRouters(router fiber.Router) {
	router.Post("/user/signin", middlewares.RateLimitMiddleware(), handler.SigninUser)
	router.Post("/user/signup", middlewares.RateLimitMiddleware(), handler.SignupUser)
}

func NewUserHandler(userService service.UserService, val *validator.Validate) *UserHandler {
	return &UserHandler{userService, val}
}
