package usecases

import (
	"context"
	"fmt"
	"work01/internal/entities"
	"work01/internal/repositories"

	"github.com/google/uuid"
)

type (
	RoleUsecase interface {
		CreateRole(role entities.Role, roleFeatures []entities.RoleFeature) error
		GetRoleById(ctx context.Context, id uuid.UUID) (*entities.ResRoleDetails, error)
		GetAllRolesDefault() ([]entities.Role, error)
		GetAllRolesModify(ctx context.Context) ([]entities.ResAllRoleDetails, error)
		GetAllRolesDropdown(ctx context.Context) ([]entities.ResAllRoleDropDown, error)
		UpdateRole(ctx context.Context, role *entities.Role, roleFeatures []entities.RoleFeature) error
		DeleteRole(ctx context.Context, id uuid.UUID, delBy uuid.UUID) error
	}

	roleUsecase struct {
		repo repositories.RoleRepository
	}
)

func NewRoleUsecase(repo repositories.RoleRepository) RoleUsecase {
	return &roleUsecase{repo: repo}
}

func (s *roleUsecase) CreateRole(role entities.Role, roleFeatures []entities.RoleFeature) error {
	if role.Name == "" {
		return fmt.Errorf("role name cannot be empty on create")
	}
	if err := s.ValidateBodyRole(role.Name, role.Level, role.CreatedBy); err != nil {
		return err
	}

	if err := s.repo.Create(&role, roleFeatures); err != nil {
		return err
	}

	return nil
}

func (s *roleUsecase) GetRoleById(ctx context.Context, id uuid.UUID) (*entities.ResRoleDetails, error) {
	role, roleFeatures, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	res := entities.ResRoleDetails{
		RoleID:    role.ID,
		RoleLevel: role.Level,
		RoleName:  role.Name,
		Features:  roleFeatures,
	}

	return &res, nil
}

func (s *roleUsecase) GetAllRolesDefault() ([]entities.Role, error) {
	roles, err := s.repo.GetAllDefault()
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (s *roleUsecase) GetAllRolesModify(ctx context.Context) ([]entities.ResAllRoleDetails, error) {
	roles, err := s.repo.GetAllModify(ctx)
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (s *roleUsecase) GetAllRolesDropdown(ctx context.Context) ([]entities.ResAllRoleDropDown, error) {
	roles, err := s.repo.GetAllModify(ctx)
	if err != nil {
		return nil, err
	}

	var roleRes []entities.ResAllRoleDropDown
	for _, role := range roles {
		roleRes = append(roleRes, entities.ResAllRoleDropDown{
			RoleID:   role.RoleID,
			RoleName: role.RoleName,
		})
	}

	return roleRes, nil
}

func (s *roleUsecase) UpdateRole(ctx context.Context, role *entities.Role, roleFeatures []entities.RoleFeature) error {
	if role.Name == "" {
		return fmt.Errorf("role name cannot be empty on update")
	}

	roleSelect, _, err := s.repo.GetById(ctx, role.ID)
	if err != nil {
		return err
	}

	userRoleLevel, err := s.repo.GetRoleLevelOfRoleUserByUserId(role.UpdatedBy)
	if err != nil {
		return err
	}

	if role.Level > userRoleLevel.RoleLevel {
		return fmt.Errorf("you can not modify your role level to higher than tour level")
	}

	if err := s.ValidateUpdateBodyRole(role.ID, role.Level, role.Name, roleSelect.Level, role.UpdatedBy); err != nil {
		return err
	}

	err = s.repo.Update(ctx, role, roleFeatures)
	if err != nil {
		return err
	}

	return nil
}

func (s *roleUsecase) DeleteRole(ctx context.Context, id uuid.UUID, delBy uuid.UUID) error {
	role, _, err := s.repo.GetById(ctx, id)
	if err != nil {
		return err
	}

	if role.Name == "Super Administrator" {
		return fmt.Errorf("can't remove role super administrator")
	}

	if err := s.ValidateBodyRoleDelete(id, role.Level, delBy); err != nil {
		return err
	}

	err = s.repo.Delete(id, delBy)
	if err != nil {
		return err
	}

	return nil
}

func (s *roleUsecase) ValidateBodyRole(roleName string, roleLevel int32, manageBy uuid.UUID) error {

	if roleName != "" {
		checkName, _ := s.repo.RoleNameIsAlreadyExits(roleName)
		if checkName {
			return fmt.Errorf("the role name alredy exists")
		}
	}

	userRoleLevel, err := s.repo.GetRoleLevelOfRoleUserByUserId(manageBy)
	if err != nil {
		return err
	}

	if userRoleLevel.RoleLevel <= roleLevel {
		return fmt.Errorf("the role level you hold must be higher than the role level you are attempting to create")
	}

	if roleLevel > 100 || roleLevel < 0 {
		return fmt.Errorf("the role level can be set from 0 to 100")
	}

	return nil
}

func (s *roleUsecase) ValidateUpdateBodyRole(roleId uuid.UUID, roleLevelCurr int32, roleName string, roleLevel int32, manageBy uuid.UUID) error {

	if roleName != "" {
		checkName, _ := s.repo.RoleNameIsAlreadyExitsUpdate(roleId, roleName)
		if checkName {
			return fmt.Errorf("the role name alredy exists")
		}
	}

	userRoleLevel, err := s.repo.GetRoleLevelOfRoleUserByUserId(manageBy)
	if err != nil {
		return err
	}

	if userRoleLevel.RoleLevel <= roleLevel {
		return fmt.Errorf("the role level you hold must be higher than the role level you are attempting to manage")
	}

	if userRoleLevel.RoleLevel <= roleLevelCurr {
		return fmt.Errorf("you role level can not up level this role to more than or equal your level")
	}

	if roleLevel > 100 || roleLevel < 0 {
		return fmt.Errorf("the role level can be set from 0 to 100")
	}

	return nil
}

func (s *roleUsecase) ValidateBodyRoleDelete(roleId uuid.UUID, roleLevel int32, manageBy uuid.UUID) error {

	userRoleLevel, err := s.repo.GetRoleLevelOfRoleUserByUserId(manageBy)
	if err != nil {
		return err
	}

	if userRoleLevel.RoleLevel <= roleLevel {
		return fmt.Errorf("the role level you hold must be higher than the role level you are attempting to manage")
	}

	checkHaveUser, _ := s.repo.CheckRoleHaveUserUsed(roleId)
	if checkHaveUser {
		return fmt.Errorf("can not delete the role that have user in used")
	}

	return nil
}
