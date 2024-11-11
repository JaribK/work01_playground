package handlers

import (
	"work01/internal/entities"
	"work01/internal/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HttpUserHandler struct {
	userUseCase usecases.UserUseCase
}

func NewHttpUserHandler(useCase usecases.UserUseCase) *HttpUserHandler {
	return &HttpUserHandler{userUseCase: useCase}
}

func (h *HttpUserHandler) CreateUserHandler(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	user.ID = uuid.New()
	if err := h.userUseCase.CreateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "create user successful.",
	})
}

func (h *HttpUserHandler) GetUserByIdHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID.",
		})
	}

	role, err := h.userUseCase.GetUserById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "user not found.",
		})
	}

	return c.JSON(role)
}

func (h *HttpUserHandler) GetAllUsersHandler(c *fiber.Ctx) error {
	users, err := h.userUseCase.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to fetch users.",
		})
	}

	return c.JSON(users)
}

func (h *HttpUserHandler) UpdateUserHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID.",
		})
	}

	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	user.ID = id
	if err := h.userUseCase.UpdateUser(user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "update user successful.",
		"ID user": user.ID,
	})
}

func (h *HttpUserHandler) DeleteUserHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID.",
		})
	}

	if err := h.userUseCase.DeleteUser(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "delete user successful.",
		"ID user": id,
	})
}
