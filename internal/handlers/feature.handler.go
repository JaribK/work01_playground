package handlers

import (
	"work01/internal/entities"
	"work01/internal/usecases"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type HttpFeatureHandler struct {
	featureUseCase usecases.FeatureUsecase
}

func NewHttpFeatureHandler(useCase usecases.FeatureUsecase) *HttpFeatureHandler {
	return &HttpFeatureHandler{featureUseCase: useCase}
}

func (h *HttpFeatureHandler) CreateFeatureHandler(c *fiber.Ctx) error {
	var feature entities.Feature
	if err := c.BodyParser(&feature); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	feature.ID = uuid.New()
	if err := h.featureUseCase.CreateFeature(feature); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "create feature successful.",
	})

}

func (h *HttpFeatureHandler) GetFeatureByIdHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID.",
		})
	}

	feature, err := h.featureUseCase.GetFeatureById(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "feature not found.",
		})
	}

	return c.JSON(feature)
}

func (h *HttpFeatureHandler) GetAllFeaturesHandler(c *fiber.Ctx) error {
	features, err := h.featureUseCase.GetAllFeatures()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Unable to fetch features.",
		})
	}

	return c.JSON(features)
}

func (h *HttpFeatureHandler) UpdateFeatureHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	var feature entities.Feature
	if err := c.BodyParser(&feature); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request.",
		})
	}

	feature.ID = id

	if err := h.featureUseCase.UpdateFeature(feature); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":    "update feature successful.",
		"ID feature": feature.ID,
	})
}

func (h *HttpFeatureHandler) DeleteFeatureHandler(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID",
		})
	}

	if err := h.featureUseCase.DeleteFeature(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message":    "detele feature successful.",
		"ID feature": id,
	})
}
