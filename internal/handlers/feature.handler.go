package handlers

import (
	"mime/multipart"
	"work01/internal/entities"
	"work01/internal/helpers"
	"work01/internal/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type (
	HttpFeatureHandler interface {
		CreateFeatureHandler(c *fiber.Ctx) error
		GetFeatureByIdHandler(c *fiber.Ctx) error
		GetAllFeaturePermissionsHandler(c *fiber.Ctx) error
		GetRefFeatureHandler(c *fiber.Ctx) error
		GetAllFeaturesDefaultHandler(c *fiber.Ctx) error
		UpdateFeatureHandler(c *fiber.Ctx) error
		DeleteFeatureHandler(c *fiber.Ctx) error
	}

	httpFeatureHandler struct {
		featureUseCase usecases.FeatureUsecase
	}
)

func NewHttpFeatureHandler(useCase usecases.FeatureUsecase) HttpFeatureHandler {
	return &httpFeatureHandler{featureUseCase: useCase}
}

func (h *httpFeatureHandler) CreateFeatureHandler(c *fiber.Ctx) error {
	var feature entities.Feature
	if err := c.BodyParser(&feature); err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	file, err := c.FormFile("menuIcon")
	var Iconfile *multipart.FileHeader
	if err == nil {
		Iconfile = file
	}

	feature.ID = uuid.New()
	if err := h.featureUseCase.CreateFeature(feature, Iconfile); err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":           "create feature successful.",
		"created featureId": feature.ID,
	})
}

func (h *httpFeatureHandler) GetFeatureByIdHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	feature, err := h.featureUseCase.GetFeatureById(ctx, id)
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusNotFound, "Feature Not Found", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(feature)
}

func (h *httpFeatureHandler) GetAllFeaturePermissionsHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	features, err := h.featureUseCase.GetAllRoleFeatures(ctx)
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(features)
}

func (h *httpFeatureHandler) GetRefFeatureHandler(c *fiber.Ctx) error {
	features, err := h.featureUseCase.GetRefFeatures()
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(features)
}

func (h *httpFeatureHandler) GetAllFeaturesDefaultHandler(c *fiber.Ctx) error {
	features, err := h.featureUseCase.GetAllFeaturesDefault()
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(features)
}

func (h *httpFeatureHandler) UpdateFeatureHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	var feature entities.Feature
	if err := c.BodyParser(&feature); err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	file, err := c.FormFile("menuIcon")
	var Iconfile *multipart.FileHeader
	if err == nil {
		Iconfile = file
	}

	feature.ID = id

	if err := h.featureUseCase.UpdateFeature(ctx, feature, Iconfile); err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":           "update feature successful.",
		"updated featureId": feature.ID,
	})
}

func (h *httpFeatureHandler) DeleteFeatureHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	if err := h.featureUseCase.DeleteFeature(id); err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":           "detele feature successful.",
		"deleted featureId": id,
	})
}
