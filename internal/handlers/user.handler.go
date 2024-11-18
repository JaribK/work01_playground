package handlers

import (
	"strconv"
	"work01/internal/entities"
	"work01/internal/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HttpUserHandler struct {
	userUseCase usecases.UserUsecase
}

func NewHttpUserHandler(useCase usecases.UserUsecase) *HttpUserHandler {
	return &HttpUserHandler{userUseCase: useCase}
}

func (h *HttpUserHandler) CreateUserHandler(c *fiber.Ctx) error {
	var user entities.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	creBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	user.ID = uuid.New()
	user.CreatedBy = creBy
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
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	size, err := strconv.Atoi(c.Query("size", "10"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	users, err := h.userUseCase.GetAllUsers(page, size)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)
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

	updBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	user.ID = id
	user.UpdatedBy = updBy
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

	delBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	if err := h.userUseCase.DeleteUser(id, delBy); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "delete user successful.",
		"ID user": id,
	})
}

// func (h *HttpUserHandler) LoginHandler(c *fiber.Ctx) error {
// 	var requests struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}

// 	if err := c.BodyParser(&requests); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "invalid request",
// 		})
// 	}

// 	user, token, err := h.userUseCase.Login(requests.Email, requests.Password)

// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}

// 	return c.JSON(fiber.Map{
// 		"message": "Login successful",
// 		"user":    user,
// 		"token":   token,
// 	})
// }
