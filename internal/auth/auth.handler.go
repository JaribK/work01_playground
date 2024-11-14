package auth

import (
	"github.com/gofiber/fiber/v2"
)

type HttpAuthHandler struct {
	authUse authUsecase
}

func NewHttpAuthHandler(usecase authUsecase) *HttpAuthHandler {
	return &HttpAuthHandler{authUse: usecase}
}

func (h *HttpAuthHandler) LoginHandler(c *fiber.Ctx) error {
	var requests struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&requests); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	user, token, err := h.authUse.Login(requests.Email, requests.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user":    user,
		"token":   token,
	})
}
