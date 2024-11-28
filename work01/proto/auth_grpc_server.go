package proto

import (
	"context"
	"work01/internal/usecases"

	"github.com/google/uuid"
)

type authorizationServer struct {
	authUsecase usecases.AuthorizationUsecase
}

func NewAuthorizationServer(authUsecase usecases.AuthorizationUsecase) AuthorizationServer {
	return &authorizationServer{authUsecase}
}

func (authorizationServer) mustEmbedUnimplementedAuthorizationServer() {}

func (s *authorizationServer) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
	var mergedPermissions []*PermissionDTO
	user, token, err := s.authUsecase.Login(req.Email, req.Password)
	if err != nil {
		return nil, err
	}

	userDTO, err := s.authUsecase.GetUserDataById(user.ID)
	if err != nil {
		return nil, err
	}

	for _, Feature := range userDTO.Features {
		var parentId *string
		if Feature.ParentMenuId != nil {
			idStr := Feature.ParentMenuId.String()
			parentId = &idStr
		} else {
			parentId = nil
		}
		mergedPermissions = append(mergedPermissions, &PermissionDTO{
			FeatureId:    Feature.ID.String(),
			Name:         Feature.Name,
			ParentMenuId: parentId,
			MenuIcon:     Feature.MenuIcon,
			MenuNameTh:   Feature.MenuNameTh,
			MenuNameEn:   Feature.MenuNameEn,
			MenuSlug:     Feature.MenuSlug,
			MenuSeqNo:    Feature.MenuSeqNo,
			IsActive:     Feature.IsActive,
			IsAdd:        *Feature.IsAdd,
			IsView:       *Feature.IsView,
			IsEdit:       *Feature.IsEdit,
			IsDelete:     *Feature.IsDelete,
		})
	}

	res := LoginResponse{
		Message:      "Login Successful",
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		User: &User{
			UserId:            user.ID.String(),
			Email:             user.Email,
			FirstName:         user.FirstName,
			LastName:          user.LastName,
			PhoneNumber:       user.PhoneNumber,
			Avatar:            *user.Avatar,
			RoleName:          user.Role.Name,
			RoleLevel:         user.Role.Level,
			TwoFactorAuthUrl:  user.TwoFactorAuthUrl,
			TwoFactorEnabled:  user.TwoFactorEnabled,
			TwoFactorToken:    user.TwoFactorToken,
			TwoFactorVerified: user.TwoFactorVerified,
			Permissions:       mergedPermissions,
		},
	}

	return &res, nil

}

func (s *authorizationServer) Logout(ctx context.Context, req *LogoutRequest) (*LogoutResponse, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	err = s.authUsecase.Logout(userId, req.Token)
	if err != nil {
		return nil, err
	}

	res := LogoutResponse{
		Message: "Logout Successful",
	}

	return &res, nil
}

func (s *authorizationServer) RefreshToken(ctx context.Context, req *RefreshTokenRequest) (*RefreshTokenResponse, error) {
	newAccessToken, err := s.authUsecase.RefreshToken(req.RefreshToken)
	if err != nil {
		return nil, err
	}

	res := RefreshTokenResponse{
		AccessToken: newAccessToken,
	}

	return &res, nil
}
