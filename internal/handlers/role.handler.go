package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"work01/internal/entities"
	"work01/internal/usecases"
)

type HttpRoleHandler struct {
	roleUseCase usecases.RoleUsecase
}

func NewHttpRoleHandler(useCase usecases.RoleUsecase) *HttpRoleHandler {
	return &HttpRoleHandler{roleUseCase: useCase}
}

func (h *HttpRoleHandler) CreateRoleHandler(c *fiber.Ctx) error {
	var role entities.Role
	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	creBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	role.ID = uuid.New()
	role.CreatedBy = creBy
	if err := h.roleUseCase.CreateRole(role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "create role successful.",
	})

}

func (h *HttpRoleHandler) GetRoleByIdHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID.",
		})
	}

	role, err := h.roleUseCase.GetRoleById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "role not found.",
		})
	}

	return c.JSON(role)
}

func (h *HttpRoleHandler) GetAllRolesHandler(c *fiber.Ctx) error {
	roles, err := h.roleUseCase.GetAllRoles()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to fetch roles.",
		})
	}

	return c.JSON(roles)
}

func (h *HttpRoleHandler) UpdateRoleHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var role entities.Role
	if err := c.BodyParser(&role); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	updBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	role.ID = id
	role.UpdatedBy = updBy
	if err := h.roleUseCase.UpdateRole(role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "update role successful.",
		"ID role": role.ID,
	})
}

func (h *HttpRoleHandler) DeleteRoleHandler(c *fiber.Ctx) error {
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

	if err := h.roleUseCase.DeleteRole(id, delBy); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "detele role successful.",
		"ID role": id,
	})
}
