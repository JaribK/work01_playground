package handlers

import (
	"mime/multipart"
	"strconv"
	"work01/internal/entities"
	"work01/internal/helpers"
	"work01/internal/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type (
	HttpUserHandler interface {
		CreateUserHandler(c *fiber.Ctx) error
		GetUserByIdHandler(c *fiber.Ctx) error
		GetUserProfileByIdHandler(c *fiber.Ctx) error
		GetAllUsersWithPageHandler(c *fiber.Ctx) error
		GetAllUsersNoPageHandler(c *fiber.Ctx) error
		UpdateUserHandler(c *fiber.Ctx) error
		ChangePsswordHandler(c *fiber.Ctx) error
		DeleteUserHandler(c *fiber.Ctx) error
	}

	httpUserHandler struct {
		userUseCase usecases.UserUsecase
	}
)

func NewHttpUserHandler(useCase usecases.UserUsecase) HttpUserHandler {
	return &httpUserHandler{userUseCase: useCase}
}

func (h *httpUserHandler) CreateUserHandler(c *fiber.Ctx) error {
	var user entities.ReqUser
	if err := c.BodyParser(&user); err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request Body", err.Error())
	}

	creBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	file, err := c.FormFile("avatar")
	var avatarfile *multipart.FileHeader
	if err == nil {
		avatarfile = file
	}

	user.ID = uuid.New()
	user.CreatedBy = creBy
	if err := h.userUseCase.CreateUser(user, avatarfile); err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	userCheck, err := h.userUseCase.GetUserByIdCheckRole(user.ID)
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusNotFound, "Role Not Found", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "create user successful.",
		"email":   user.Email,
		"role":    userCheck.Role.Name,
	})
}

func (h *httpUserHandler) GetUserByIdHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	id, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request Body", err.Error())
	}

	user, err := h.userUseCase.GetUserById(ctx, id)
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusNotFound, "User Not Found", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *httpUserHandler) GetUserProfileByIdHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request Body", err.Error())
	}

	user, err := h.userUseCase.GetUserProfileById(id)
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusNotFound, "User Not Found", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *httpUserHandler) GetAllUsersWithPageHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request Body", err.Error())
	}

	size, err := strconv.Atoi(c.Query("size", "10"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request Body", err.Error())
	}

	roleId := c.Query("roleId", "")
	isActive := c.Query("isActive", "")
	phoneNumber := c.Query("phoneNumber", "")
	fullName := c.Query("fullName", "")

	users, err := h.userUseCase.GetAllUsersWithPage(ctx, page, size, roleId, isActive, phoneNumber, fullName)
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *httpUserHandler) GetAllUsersNoPageHandler(c *fiber.Ctx) error {
	users, err := h.userUseCase.GetAllUsersNoPage()
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(users)
}

func (h *httpUserHandler) UpdateUserHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request Body", err.Error())
	}

	var user entities.ReqUser
	if err := c.BodyParser(&user); err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request Body", err.Error())
	}

	updBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	file, err := c.FormFile("avatar")
	var avatarfile *multipart.FileHeader
	if err == nil {
		avatarfile = file
	}

	user.ID = id
	user.UpdatedBy = updBy
	if err := h.userUseCase.UpdateUser(ctx, user, avatarfile); err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "update user successful.",
		"updated userId": id,
	})
}

func (h *httpUserHandler) ChangePsswordHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request Body", err.Error())
	}

	var pass entities.ReqChangePassword
	if err := c.BodyParser(&pass); err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request Body", err.Error())
	}

	updBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	pass.UserId = id
	pass.UpdatedBy = updBy
	if err := h.userUseCase.ChangePssword(ctx, pass); err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":                 "update password successful.",
		"updated password userId": id,
	})
}

func (h *httpUserHandler) DeleteUserHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request Body", err.Error())
	}

	delBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	if err := h.userUseCase.DeleteUser(ctx, id, delBy); err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.JSON(fiber.Map{
		"message":        "delete user successful.",
		"deleted userId": id,
	})
}
