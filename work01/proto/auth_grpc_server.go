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

	for _, permission := range user.Role.Permissions {
		var parentId *string
		if permission.Feature.ParentMenuId != nil {
			idStr := permission.Feature.ParentMenuId.String()
			parentId = &idStr
		} else {
			parentId = nil
		}
		mergedPermissions = append(mergedPermissions, &PermissionDTO{
			FeatureId:    permission.Feature.ID.String(),
			Name:         permission.Feature.Name,
			ParentMenuId: parentId,
			MenuIcon:     permission.Feature.MenuIcon,
			MenuNameTh:   permission.Feature.MenuNameTh,
			MenuNameEn:   permission.Feature.MenuNameEn,
			MenuSlug:     permission.Feature.MenuSlug,
			MenuSeqNo:    permission.Feature.MenuSeqNo,
			IsActive:     permission.Feature.IsActive,
			CreateAccess: permission.CreateAccess,
			ReadAccess:   permission.ReadAccess,
			UpdateAccess: permission.UpdateAccess,
			DeleteAccess: permission.DeleteAccess,
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
			Avatar:            user.Avatar,
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
