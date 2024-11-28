package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"work01/internal/entities"
	"work01/internal/usecases"
)

type HttpRoleFeatureHandler struct {
	roleFeatureUseCase usecases.RoleFeatureUsecase
}

func NewHttpRoleFeatureHandler(useCase usecases.RoleFeatureUsecase) *HttpRoleFeatureHandler {
	return &HttpRoleFeatureHandler{roleFeatureUseCase: useCase}
}

func (h *HttpRoleFeatureHandler) CreateRoleFeatureHandler(c *fiber.Ctx) error {
	var roleFeature entities.RoleFeature
	if err := c.BodyParser(&roleFeature); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	roleFeature.ID = uuid.New()
	if err := h.roleFeatureUseCase.CreateRoleFeature(roleFeature); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":               "create roleFeature successful.",
		"created roleFeatureId": roleFeature.ID,
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

	roleFeature, err := h.roleFeatureUseCase.GetRoleFeatureById(ctx, id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "role not found.",
		})
	}

	return c.Status(fiber.StatusOK).JSON(roleFeature)
}

func (h *HttpRoleFeatureHandler) GetAllRoleFeaturesHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	roleFeatures, err := h.roleFeatureUseCase.GetAllRoleFeatures(ctx)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(roleFeatures)
}

func (h *HttpRoleFeatureHandler) UpdateRoleFeatureHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var roleFeature entities.RoleFeature
	if err := c.BodyParser(&roleFeature); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	roleFeature.ID = id

	if err := h.roleFeatureUseCase.UpdateRoleFeature(roleFeature); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":               "update roleFeature successful.",
		"updated roleFeatureId": roleFeature.ID,
	})
}

func (h *HttpRoleFeatureHandler) DeleteRoleFeatureHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	if err := h.roleFeatureUseCase.DeleteRoleFeature(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":               "detele role successful.",
		"deleted roleFeatureId": id,
	})
}
