package pkg

import (
	"context"
	"fmt"
	"log"
	"work01/internal/auth"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

func TokenValidationMiddleware(c *fiber.Ctx) error {
	redisClient := NewRedisClient()

	if redisClient == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Redis client not initialized",
		})
	}

	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization token is required",
		})
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	token, err := auth.ValidateToken(tokenString)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	cacheKey := fmt.Sprintf("blocked:%s", tokenString)
	blocked, err := redisClient.Get(context.Background(), cacheKey).Result()
	if err != nil && err != redis.Nil {
		log.Printf("Error fetching from Redis: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "server error"})
	}

	if blocked == tokenString {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "token blocked"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "invalid claims",
		})
	}

	userId, ok := claims["userId"]
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "userId missing from token",
		})
	}

	c.Locals("userId", userId)

	return c.Next()
}
