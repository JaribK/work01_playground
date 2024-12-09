package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"work01/internal/entities"
	"work01/internal/helpers"
	"work01/internal/usecases"
)

type (
	HttpRoleFeatureHandler interface {
		CreateRoleFeatureHandler(c *fiber.Ctx) error
		GetRoleFeatureByIdHandler(c *fiber.Ctx) error
		GetAllRoleFeaturesHandler(c *fiber.Ctx) error
		UpdateRoleFeatureHandler(c *fiber.Ctx) error
		DeleteRoleFeatureHandler(c *fiber.Ctx) error
	}

	httpRoleFeatureHandler struct {
		roleFeatureUseCase usecases.RoleFeatureUsecase
	}
)

func NewHttpRoleFeatureHandler(useCase usecases.RoleFeatureUsecase) HttpRoleFeatureHandler {
	return &httpRoleFeatureHandler{roleFeatureUseCase: useCase}
}

func (h *httpRoleFeatureHandler) CreateRoleFeatureHandler(c *fiber.Ctx) error {
	var roleFeature entities.RoleFeature
	if err := c.BodyParser(&roleFeature); err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	roleFeature.ID = uuid.New()
	if err := h.roleFeatureUseCase.CreateRoleFeature(roleFeature); err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":               "create roleFeature successful.",
		"created roleFeatureId": roleFeature.ID,
	})
}

func (h *httpRoleFeatureHandler) GetRoleFeatureByIdHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	roleFeature, err := h.roleFeatureUseCase.GetRoleFeatureById(ctx, id)
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusNotFound, "RoleFeature Not Found", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(roleFeature)
}

func (h *httpRoleFeatureHandler) GetAllRoleFeaturesHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	roleFeatures, err := h.roleFeatureUseCase.GetAllRoleFeatures(ctx)
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(roleFeatures)
}

func (h *httpRoleFeatureHandler) UpdateRoleFeatureHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	var roleFeature entities.RoleFeature
	if err := c.BodyParser(&roleFeature); err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	roleFeature.ID = id

	if err := h.roleFeatureUseCase.UpdateRoleFeature(roleFeature); err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":               "update roleFeature successful.",
		"updated roleFeatureId": roleFeature.ID,
	})
}

func (h *httpRoleFeatureHandler) DeleteRoleFeatureHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	if err := h.roleFeatureUseCase.DeleteRoleFeature(id); err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":               "detele role successful.",
		"deleted roleFeatureId": id,
	})
}
