package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"work01/internal/entities"
	"work01/internal/helpers"
	"work01/internal/usecases"
)

type (
	HttpAuthorizationHandler interface {
		RefreshToken(c *fiber.Ctx) error
		LoginHandler(c *fiber.Ctx) error
		LogoutHandler(c *fiber.Ctx) error
		CreateAuthorizationHandler(c *fiber.Ctx) error
		GetAuthorizationByIdHandler(c *fiber.Ctx) error
		GetAllAuthorizationsHandler(c *fiber.Ctx) error
		UpdateAuthorizationHandler(c *fiber.Ctx) error
		DeleteAuthorizationHandler(c *fiber.Ctx) error
	}

	httpAuthorizationHandler struct {
		authorizationUsecase usecases.AuthorizationUsecase
	}
)

func NewHttpAuthorizationHandler(useCase usecases.AuthorizationUsecase) HttpAuthorizationHandler {
	return &httpAuthorizationHandler{authorizationUsecase: useCase}
}

func (h *httpAuthorizationHandler) RefreshToken(c *fiber.Ctx) error {
	var req entities.RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	newAccessToken, err := h.authorizationUsecase.RefreshToken(req.RefreshToken)
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusUnauthorized, "Unauthorization", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": newAccessToken,
	})
}

func (h *httpAuthorizationHandler) LoginHandler(c *fiber.Ctx) error {
	var requests struct {
		Identifier string `json:"identifier"`
		Password   string `json:"password"`
	}

	if err := c.BodyParser(&requests); err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	user, token, err := h.authorizationUsecase.Login(requests.Identifier, requests.Password)

	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusUnauthorized, "Unauthorization", err.Error())
	}

	userDTO, err := h.authorizationUsecase.GetUserDataById(user.ID)
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusNotFound, "User Not Found", err.Error())
	}

	res := entities.ResLogin{
		Message:      "Login successful",
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		User:         userDTO,
	}

	return c.Status(fiber.StatusOK).JSON(res)
}

func (h *httpAuthorizationHandler) LogoutHandler(c *fiber.Ctx) error {
	tokenString := c.Get("Authorization")
	if tokenString == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Authorization token is required",
		})
	}

	if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
		tokenString = tokenString[7:]
	}

	userID, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	err = h.authorizationUsecase.Logout(userID, tokenString)
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Logged out successful",
	})
}

func (h *httpAuthorizationHandler) CreateAuthorizationHandler(c *fiber.Ctx) error {
	var auth entities.Authorization
	if err := c.BodyParser(&auth); err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	creBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	auth.ID = uuid.New()
	auth.CreatedBy = creBy
	if err := h.authorizationUsecase.CreateAuthorization(auth); err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":        "create auth successful",
		"created authId": auth.ID,
	})
}

func (h *httpAuthorizationHandler) GetAuthorizationByIdHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	auth, err := h.authorizationUsecase.GetAuthorizationById(id)
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusNotFound, "Auth Not Found", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(auth)
}

func (h *httpAuthorizationHandler) GetAllAuthorizationsHandler(c *fiber.Ctx) error {
	auths, err := h.authorizationUsecase.GetAllAuthorizations()
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(auths)
}

func (h *httpAuthorizationHandler) UpdateAuthorizationHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	var auth entities.Authorization
	if err := c.BodyParser(&auth); err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	updBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	auth.ID = id
	auth.UpdatedBy = updBy
	if err := h.authorizationUsecase.UpdateAuthorization(auth); err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "update auth successful",
		"updated authId": auth.ID,
	})
}

func (h *httpAuthorizationHandler) DeleteAuthorizationHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	delBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	if err := h.authorizationUsecase.DeleteAuthorization(id, delBy); err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "detele auth successful",
		"deleted authId": id,
	})
}
