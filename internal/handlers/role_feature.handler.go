package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"work01/internal/entities"
	"work01/internal/usecases"
)

type HttpRoleFeatureHandler struct {
	rolePermissionUseCase usecases.RoleFeatureUsecase
}

func NewHttpRoleFeatureHandler(useCase usecases.RoleFeatureUsecase) *HttpRoleFeatureHandler {
	return &HttpRoleFeatureHandler{rolePermissionUseCase: useCase}
}

func (h *HttpRoleFeatureHandler) CreateRoleFeatureHandler(c *fiber.Ctx) error {
	var rolePermission entities.RoleFeature
	if err := c.BodyParser(&rolePermission); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	rolePermission.ID = uuid.New()
	if err := h.rolePermissionUseCase.CreateRoleFeature(rolePermission); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "create rolePermission successful.",
	})

}

func (h *HttpRoleFeatureHandler) GetRoleFeatureByIdHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID.",
		})
	}

	rolePermission, err := h.rolePermissionUseCase.GetRoleFeatureById(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "role not found.",
		})
	}

	return c.JSON(rolePermission)
}

func (h *HttpRoleFeatureHandler) GetAllRoleFeaturesHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	rolePermissions, err := h.rolePermissionUseCase.GetAllRoleFeatures(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(rolePermissions)
}

func (h *HttpRoleFeatureHandler) UpdateRoleFeatureHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var rolePermission entities.RoleFeature
	if err := c.BodyParser(&rolePermission); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	rolePermission.ID = id

	if err := h.rolePermissionUseCase.UpdateRoleFeature(rolePermission); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":           "update rolePermission successful.",
		"ID rolePermission": rolePermission.ID,
	})
}

func (h *HttpRoleFeatureHandler) DeleteRoleFeatureHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	if err := h.rolePermissionUseCase.DeleteRoleFeature(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":           "detele role successful.",
		"ID rolePermission": id,
	})
}
