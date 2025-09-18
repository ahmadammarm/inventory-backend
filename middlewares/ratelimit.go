package middlewares

import (
	"github.com/ahmadammarm/inventory-backend/pkg/response"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(5.0/300, 5)

func RateLimitMiddleware() fiber.Handler {

	return func(context *fiber.Ctx) error {
        
		if !limiter.Allow() {
			return context.Status(fiber.StatusTooManyRequests).JSON(response.ErrorResponse{
				Message: "Too many requests, please try again later",
				Success: false,
				Code:    fiber.StatusTooManyRequests,
				Errors:  "Rate limit exceeded",
			})
		}

		return context.Next()
	}
}
