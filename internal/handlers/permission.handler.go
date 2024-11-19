package handlers

import (
	"work01/internal/entities"
	"work01/internal/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HttpPermissionHandler struct {
	permissionUseCase usecases.PermissionUsecase
}

func NewHttpPermissionHandler(useCase usecases.PermissionUsecase) *HttpPermissionHandler {
	return &HttpPermissionHandler{permissionUseCase: useCase}
}

func (h *HttpPermissionHandler) CreatePermissionHandler(c *fiber.Ctx) error {
	var permission entities.Permission
	if err := c.BodyParser(&permission); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	permission.ID = uuid.New()
	if err := h.permissionUseCase.CreatePermission(permission); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "create permission successful.",
	})

}

func (h *HttpPermissionHandler) GetPermissionByIdHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID.",
		})
	}

	permission, err := h.permissionUseCase.GetPermissionById(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "permission not found.",
		})
	}

	return c.JSON(permission)
}

func (h *HttpPermissionHandler) GetAllPermissionsHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	permissions, err := h.permissionUseCase.GetAllPermissions(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to fetch permissions.",
		})
	}

	return c.JSON(permissions)
}

func (h *HttpPermissionHandler) UpdatePermissionHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var permission entities.Permission
	if err := c.BodyParser(&permission); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	permission.ID = id

	if err := h.permissionUseCase.UpdatePermission(permission); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":       "update permission successful.",
		"ID permission": permission.ID,
	})
}

func (h *HttpPermissionHandler) DeletePermissionHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	if err := h.permissionUseCase.DeletePermission(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":       "detele permission successful.",
		"ID permission": id,
	})
}
