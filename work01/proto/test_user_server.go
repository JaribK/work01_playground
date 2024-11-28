package proto

import (
	"context"
	"work01/internal/entities"
	"work01/internal/usecases"

	"github.com/google/uuid"
)

type userGrpcServiceServer struct {
	userUsecase usecases.UserUsecase
}

func NewUserGrpcServiceServer(usecase usecases.UserUsecase) UserGrpcServiceServer {
	return &userGrpcServiceServer{userUsecase: usecase}
}

func (userGrpcServiceServer) mustEmbedUnimplementedUserGrpcServiceServer() {}

func (s *userGrpcServiceServer) CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	var user entities.User

	user.ID = uuid.New()
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.PhoneNumber = req.PhoneNumber
	user.Password = req.Password
	user.Email = req.Email

	if err := s.userUsecase.CreateUser(user); err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		Result: "create user successful",
	}

	return res, nil
}

func (s *userGrpcServiceServer) GetUserById(ctx context.Context, req *GetUserByIdReq) (*GetUserByIdRes, error) {
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

	res := &GetUserByIdRes{
		UserId:      user.UserID.String(),
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Avatar:      *user.Avatar,
		RoleId:      roleId,
		RoleName:    user.RoleName,
	}

	return res, nil
}

func (s *userGrpcServiceServer) GetAllUser(ctx context.Context, req *GetAllUserReq) (*GetAllUserRes, error) {
	var userDTOs []*UsersDTO
	users, err := s.userUsecase.GetAllUsersNoPage()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		userDTOs = append(userDTOs, &UsersDTO{
			UserId:      user.ID.String(),
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			IsActive:    user.IsActive,
			Avatar:      *user.Avatar,
			RoleName:    user.Role.Name,
		})
	}

	res := &GetAllUserRes{
		Users: userDTOs,
	}

	return res, nil
}

func (s *userGrpcServiceServer) UpdateUserById(ctx context.Context, req *UpdateUserByIdReq) (*UpdateUserByIdRes, error) {
	var user entities.User
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}
	user.ID = userId
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.PhoneNumber = req.PhoneNumber
	user.Email = req.Email
	user.IsActive = req.IsActive
	if err := s.userUsecase.UpdateUser(ctx, user); err != nil {
		return nil, err
	}

	res := &UpdateUserByIdRes{
		Result: "update successful",
	}

	return res, nil

}

func (s *userGrpcServiceServer) DeleteUserById(ctx context.Context, req *DeleteUserByIdReq) (*DeleteUserByIdRes, error) {
	userId, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, err
	}

	if err := s.userUsecase.DeleteUser(ctx, userId, userId); err != nil {
		return nil, err
	}

	res := &DeleteUserByIdRes{
		Result: "delete successful",
	}

	return res, nil
}
