package middleware

import (
	"strings"

	"github.com/gofiber/fiber/v3"
)

func validateToken(token string) bool {
	// У реальному додатку тут буде перевірка токена (наприклад, запит до сервісу авторизації)
	// Для емуляції просто перевіримо, чи дорівнює токен "valid-token"
	return token == "valid-token"
}

func AuthMiddleware(c fiber.Ctx) error {
	// Отримуємо токен з заголовка Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization header is missing",
		})
	}

	// Перевіряємо формат заголовка (наприклад, "Bearer token")
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid Authorization header format",
		})
	}

	token := parts[1]

	if !validateToken(token) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid token",
		})
	}

	return c.Next()
}
