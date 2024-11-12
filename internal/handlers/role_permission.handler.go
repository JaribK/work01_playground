package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"work01/internal/entities"
	"work01/internal/usecases"
)

type HttpRolePermissionHandler struct {
	rolePermissionUseCase usecases.RolePermissionUseCase
}

func NewHttpRolePermissionHandler(useCase usecases.RolePermissionUseCase) *HttpRolePermissionHandler {
	return &HttpRolePermissionHandler{rolePermissionUseCase: useCase}
}

func (h *HttpRolePermissionHandler) CreateRolePermissionHandler(c *fiber.Ctx) error {
	var rolePermission entities.RolePermission
	if err := c.BodyParser(&rolePermission); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	rolePermission.ID = uuid.New()
	if err := h.rolePermissionUseCase.CreateRolePermission(rolePermission); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "create rolePermission successful.",
	})

}

func (h *HttpRolePermissionHandler) GetRolePermissionByIdHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID.",
		})
	}

	rolePermission, err := h.rolePermissionUseCase.GetRolePermissionById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "role not found.",
		})
	}

	return c.JSON(rolePermission)
}

func (h *HttpRolePermissionHandler) GetAllRolePermissionsHandler(c *fiber.Ctx) error {
	rolePermissions, err := h.rolePermissionUseCase.GetAllRolePermissions()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to fetch roles.",
		})
	}

	return c.JSON(rolePermissions)
}

func (h *HttpRolePermissionHandler) UpdateRolePermissionHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var rolePermission entities.RolePermission
	if err := c.BodyParser(&rolePermission); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	rolePermission.ID = id

	if err := h.rolePermissionUseCase.UpdateRolePermission(rolePermission); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":           "update rolePermission successful.",
		"ID rolePermission": rolePermission.ID,
	})
}

func (h *HttpRolePermissionHandler) DeleteRolePermissionHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	if err := h.rolePermissionUseCase.DeleteRolePermission(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":           "detele role successful.",
		"ID rolePermission": id,
	})
}
