package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"work01/internal/entities"
	"work01/internal/helpers"
	"work01/internal/usecases"
)

type (
	httpRoleHandler struct {
		roleUseCase usecases.RoleUsecase
	}

	HttpRoleHandler interface {
		CreateRoleHandler(c *fiber.Ctx) error
		GetRoleByIdHandler(c *fiber.Ctx) error
		GetAllRolesModifyHandler(c *fiber.Ctx) error
		GetAllRolesDefaultHandler(c *fiber.Ctx) error
		GetAllRolesDropdownHandler(c *fiber.Ctx) error
		UpdateRoleHandler(c *fiber.Ctx) error
		DeleteRoleHandler(c *fiber.Ctx) error
	}
)

func NewHttpRoleHandler(useCase usecases.RoleUsecase) HttpRoleHandler {
	return &httpRoleHandler{roleUseCase: useCase}
}

func (h *httpRoleHandler) CreateRoleHandler(c *fiber.Ctx) error {
	var roleReq entities.ReqRoleCreate
	if err := c.BodyParser(&roleReq); err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	creBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	role := entities.Role{
		Name:  roleReq.Name,
		Level: roleReq.Level,
	}

	var roleFeatures []entities.RoleFeature
	for _, f := range roleReq.Features {
		roleFeatures = append(roleFeatures, entities.RoleFeature{
			FeatureId: f.FeatureId,
			IsAdd:     f.IsAdd,
			IsView:    f.IsView,
			IsEdit:    f.IsEdit,
			IsDelete:  f.IsDelete,
		})
	}

	role.ID = uuid.New()
	role.CreatedBy = creBy
	if err := h.roleUseCase.CreateRole(role, roleFeatures); err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message":      "create role successful.",
		"created role": role.Name,
		"roleId":       role.ID,
	})
}

func (h *httpRoleHandler) GetRoleByIdHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	role, err := h.roleUseCase.GetRoleById(ctx, id)
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusNotFound, "Role Not Found", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(role)
}

func (h *httpRoleHandler) GetAllRolesModifyHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	roles, err := h.roleUseCase.GetAllRolesModify(ctx)
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(roles)
}

func (h *httpRoleHandler) GetAllRolesDefaultHandler(c *fiber.Ctx) error {
	roles, err := h.roleUseCase.GetAllRolesDefault()
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(roles)
}

func (h *httpRoleHandler) GetAllRolesDropdownHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	roles, err := h.roleUseCase.GetAllRolesDropdown(ctx)
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(roles)
}

func (h *httpRoleHandler) UpdateRoleHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	var roleReq entities.ReqRoleUpdate
	if err := c.BodyParser(&roleReq); err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	updBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	role := entities.Role{
		Name:  roleReq.Name,
		Level: roleReq.Level,
	}

	var roleFeatures []entities.RoleFeature
	for _, f := range roleReq.Features {
		roleFeatures = append(roleFeatures, entities.RoleFeature{
			FeatureId: f.FeatureId,
			IsAdd:     f.IsAdd,
			IsView:    f.IsView,
			IsEdit:    f.IsEdit,
			IsDelete:  f.IsDelete,
		})
	}

	role.ID = id
	role.UpdatedBy = updBy
	if err := h.roleUseCase.UpdateRole(ctx, &role, roleFeatures); err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "update role successful.",
		"updated roleId": role.ID,
	})
}

func (h *httpRoleHandler) DeleteRoleHandler(c *fiber.Ctx) error {
	ctx := c.Context()
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return helpers.ErrResponse(c, fiber.StatusBadRequest, "Bad Request", err.Error())
	}

	delBy, err := uuid.Parse(c.Locals("userId").(string))
	if err != nil {
		return err
	}

	if err := h.roleUseCase.DeleteRole(ctx, id, delBy); err != nil {
		return helpers.ErrResponse(c, fiber.StatusInternalServerError, "Internal Server Error", err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":        "detele role successful.",
		"deleted roleId": id,
	})
}
