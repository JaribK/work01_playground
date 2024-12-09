package usecases

import (
	"context"
	"mime/multipart"
	"work01/internal/entities"
	"work01/internal/proto/usergrpc"

	"github.com/google/uuid"
)

type userGrpcServiceServer struct {
	userUsecase UserUsecase
	usergrpc.UnimplementedUserGrpcServiceServer
}

func NewUserGrpcServiceServer(usecase UserUsecase) usergrpc.UserGrpcServiceServer {
	return &userGrpcServiceServer{userUsecase: usecase, UnimplementedUserGrpcServiceServer: usergrpc.UnimplementedUserGrpcServiceServer{}}
}

func (s userGrpcServiceServer) CreateUser(ctx context.Context, req *usergrpc.CreateUserReq) (*usergrpc.CreateUserRes, error) {
	var user entities.ReqUser

	user.ID = uuid.New()
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.PhoneNumber = req.PhoneNumber
	user.Password = req.Password
	user.Email = req.Email

	// file, err := c.FormFile("avatar")
	var avatarfile *multipart.FileHeader
	// if err == nil {
	// 	avatarfile = file
	// }

	if err := s.userUsecase.CreateUser(user, avatarfile); err != nil {
		return nil, err
	}

	res := &usergrpc.CreateUserRes{
		Result: "create user successful",
	}

	return res, nil
}

func (s userGrpcServiceServer) GetUserById(ctx context.Context, req *usergrpc.GetUserByIdReq) (*usergrpc.GetUserByIdRes, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	user, err := s.userUsecase.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}

	var roleId *string
	if user.RoleId != uuid.Nil {
		idStr := user.RoleId.String()
		roleId = &idStr
	} else {
		roleId = nil
	}

	res := &usergrpc.GetUserByIdRes{
		UserId:      user.UserID.String(),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Avatar:      returnNullGrpc(user.Avatar),
		RoleId:      roleId,
		RoleName:    user.RoleName,
	}

	return res, nil
}

func (s userGrpcServiceServer) GetAllUser(ctx context.Context, req *usergrpc.GetAllUserReq) (*usergrpc.GetAllUserRes, error) {
	var userDTOs []*usergrpc.AllUsersDTO
	users, err := s.userUsecase.GetAllUsersWithPage(ctx, int(req.Page), int(req.Size), req.RoleId, req.IsActive, req.PhoneNumber, req.FullName)
	if err != nil {
		return nil, err
	}

	for _, user := range users.Items {
		userDTOs = append(userDTOs, &usergrpc.AllUsersDTO{
			UserId:      user.UserID.String(),
			FullName:    user.FullName,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			IsActive:    user.IsActive,
			Avatar:      returnNullGrpc(user.Avatar),
			RoleName:    user.RoleName,
		})
	}

	res := &usergrpc.GetAllUserRes{
		CurrentPage: int32(users.Page),
		Size:        int32(users.Size),
		TotalPage:   int32(users.TotalPage),
		TotalItems:  int32(users.Total),
		Users:       userDTOs,
	}

	return res, nil
}

func (s userGrpcServiceServer) UpdateUserById(ctx context.Context, req *usergrpc.UpdateUserByIdReq) (*usergrpc.UpdateUserByIdRes, error) {
	var user entities.ReqUser
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	user.ID = userId
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.PhoneNumber = req.PhoneNumber
	user.Email = req.Email
	user.IsActive = &req.IsActive

	var avatarfile *multipart.FileHeader
	if err := s.userUsecase.UpdateUser(ctx, user, avatarfile); err != nil {
		return nil, err
	}

	res := &usergrpc.UpdateUserByIdRes{
		Result: "update successful",
	}

	return res, nil

}

func (s userGrpcServiceServer) DeleteUserById(ctx context.Context, req *usergrpc.DeleteUserByIdReq) (*usergrpc.DeleteUserByIdRes, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}

	if err := s.userUsecase.DeleteUser(ctx, userId, userId); err != nil {
		return nil, err
	}

	res := &usergrpc.DeleteUserByIdRes{
		Result: "delete successful",
	}

	return res, nil
}

func returnNullGrpc(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
