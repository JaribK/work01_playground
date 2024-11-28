package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"work01/internal/entities"
	"work01/internal/models"
	"work01/internal/usecases"
)

type HttpAuthorizationHandler struct {
	authorizationUsecase usecases.AuthorizationUsecase
}

func NewHttpAuthorizationHandler(useCase usecases.AuthorizationUsecase) *HttpAuthorizationHandler {
	return &HttpAuthorizationHandler{authorizationUsecase: useCase}
}

func (h *HttpAuthorizationHandler) RefreshToken(c *fiber.Ctx) error {
	var req models.RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	newAccessToken, err := h.authorizationUsecase.RefreshToken(req.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": newAccessToken,
	})
}

func (h *HttpAuthorizationHandler) LoginHandler(c *fiber.Ctx) error {
	var requests struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&requests); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	user, token, err := h.authorizationUsecase.Login(requests.Email, requests.Password)

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	userDTO, err := h.authorizationUsecase.GetUserDataById(user.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{})
	}

	res := models.ResLogin{
		Message:      "Login successful",
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		User:         userDTO,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *HttpAuthorizationHandler) LogoutHandler(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization token is required",
		})
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	} //ไม่แน่ใจ

	userID, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	err = h.authorizationUsecase.Logout(userID, tokenString)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logged out successfully",
	})
}

func (h *HttpAuthorizationHandler) CreateAuthorizationHandler(c *fiber.Ctx) error {
	var auth entities.Authorization
	if err := c.BodyParser(&auth); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	creBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	auth.ID = uuid.New()
	auth.CreatedBy = creBy
	if err := h.authorizationUsecase.CreateAuthorization(auth); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "create auth successful.",
		"created authId": auth.ID,
	})

}

func (h *HttpAuthorizationHandler) GetAuthorizationByIdHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID.",
		})
	}

	auth, err := h.authorizationUsecase.GetAuthorizationById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "auth not found.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(auth)
}

func (h *HttpAuthorizationHandler) GetAllAuthorizationsHandler(c *fiber.Ctx) error {
	auths, err := h.authorizationUsecase.GetAllAuthorizations()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(auths)
}

func (h *HttpAuthorizationHandler) UpdateAuthorizationHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var auth entities.Authorization
	if err := c.BodyParser(&auth); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	updBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	auth.ID = id
	auth.UpdatedBy = updBy
	if err := h.authorizationUsecase.UpdateAuthorization(auth); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "update auth successful.",
		"updated authId": auth.ID,
	})
}

func (h *HttpAuthorizationHandler) DeleteAuthorizationHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	delBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	if err := h.authorizationUsecase.DeleteAuthorization(id, delBy); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "detele auth successful.",
		"deleted authId": id,
	})
}
