package middlewares

import (
	"os"
	"strings"

	"github.com/ahmadammarm/inventory-backend/pkg/response"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)



func IsAuthenticated() fiber.Handler {
    
	return func(context *fiber.Ctx) error {
		stringToken := context.Get("Authorization")

		if stringToken == "" {
			return context.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
				Message: "Unauthorized",
				Success: false,
				Code:    fiber.StatusUnauthorized,
				Errors:  fiber.ErrUnauthorized.Message,
			})
		}

		stringToken = strings.TrimPrefix(stringToken, "Bearer ")

		token, err := jwt.Parse(stringToken, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			context.Status(fiber.StatusUnauthorized).JSON(response.ErrorResponse{
				Message: "Unauthorized",
				Success: false,
				Code:    fiber.StatusUnauthorized,
				Errors:  fiber.ErrUnauthorized.Message,
			})
		}

		context.Locals("user", token)

		return context.Next()
	}
}
