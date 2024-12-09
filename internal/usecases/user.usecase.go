package usecases

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"net/mail"
	"regexp"
	"unicode"

	"work01/internal/entities"
	"work01/internal/helpers"
	"work01/internal/repositories"
	"work01/pkg/minio"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type (
	UserUsecase interface {
		CreateUser(user entities.ReqUser, fileHeader *multipart.FileHeader) error
		GetUserById(ctx context.Context, id uuid.UUID) (*entities.ResUserDTO, error)
		GetUserProfileById(id uuid.UUID) (*entities.ResUserProfile, error)
		GetUserByIdCheckRole(id uuid.UUID) (*entities.User, error)
		GetAllUsersNoPage() ([]entities.ResUsersNoPage, error)
		GetAllUsersWithPage(ctx context.Context, page, size int, roleId, isActive string, phoneNumber string, fullName string) (helpers.Pagination, error)
		UpdateUser(ctx context.Context, user entities.ReqUser, fileHeader *multipart.FileHeader) error
		DeleteUser(ctx context.Context, id uuid.UUID, deleteBy uuid.UUID) error
		ChangePssword(ctx context.Context, reqPass entities.ReqChangePassword) error
	}

	userUsecase struct {
		repo repositories.UserRepository
	}
)

func NewUserUsecase(repo repositories.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (s *userUsecase) CreateUser(user entities.ReqUser, fileHeader *multipart.FileHeader) error {
	if err := s.CheckVariableToCreate(user.FirstName, user.LastName, user.Email, user.PhoneNumber, user.Password, user.ConfirmPassword); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	if user.RoleId != nil {
		var userSelect *entities.User
		var roleSelect *entities.Role

		userSelect, err := s.repo.GetRoleUserById(user.CreatedBy)
		if err != nil {
			return err
		}

		roleSelect, err = s.repo.GetRoleByRoleId(*user.RoleId)
		if err != nil {
			return err
		}

		if userSelect.Role.Level < roleSelect.Level {
			return fmt.Errorf("your role level (%d) must be higher than the role level (%d) you are attempting to create for user", userSelect.Role.Level, roleSelect.Level)
		}
	}

	if fileHeader != nil {
		avatarURL, err := minio.UploadAvatar(fileHeader)
		if err != nil {
			return err
		}
		user.Avatar = avatarURL
	}

	userStruct := &entities.User{
		ID:                 user.ID,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		Email:              user.Email,
		PhoneNumber:        user.PhoneNumber,
		Password:           user.Password,
		Avatar:             user.Avatar,
		TwoFactorEnabled:   user.TwoFactorEnabled,
		TwoFactorVerified:  user.TwoFactorVerified,
		TwoFactorToken:     user.TwoFactorToken,
		TwoFactorAuthUrl:   user.TwoFactorAuthUrl,
		RoleId:             user.RoleId,
		ForgotPasswordCode: user.ForgotPasswordCode,
		IsActive:           user.IsActive,
		CreatedAt:          user.CreatedAt,
		CreatedBy:          user.CreatedBy,
		UpdatedAt:          user.UpdatedAt,
		UpdatedBy:          user.UpdatedBy,
		DeletedAt:          user.DeletedAt,
		DeletedBy:          user.DeletedBy,
	}

	if err := s.repo.Create(userStruct); err != nil {
		return err
	}

	return nil
}

func (s *userUsecase) GetUserByIdCheckRole(id uuid.UUID) (*entities.User, error) {
	roleOfUser, err := s.repo.GetRoleUserById(id)
	if err != nil {
		return nil, err
	}

	return roleOfUser, nil
}

func (s *userUsecase) GetUserProfileById(id uuid.UUID) (*entities.ResUserProfile, error) {
	user, err := s.repo.GetProfileUser(id)
	if err != nil {
		return nil, err
	}

	var Test1 []interface{}
	var Test2 []interface{}

	res := entities.ResUserProfile{
		UserId:           user.ID,
		Email:            user.Email,
		FirstName:        user.FirstName,
		LastName:         user.LastName,
		PhoneNumber:      user.PhoneNumber,
		Avatar:           &user.Avatar,
		RoleId:           *user.RoleId,
		TwoFactorEnabled: *user.TwoFactorEnabled,
		IsActive:         *user.IsActive,
		CreatedAt:        user.CreatedAt,
		UserActivity:     Test1,
		UserDevice:       Test2,
	}

	return &res, nil
}

func (s *userUsecase) GetUserById(ctx context.Context, id uuid.UUID) (*entities.ResUserDTO, error) {
	user, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userUsecase) GetAllUsersWithPage(ctx context.Context, page, size int, roleId, isActive string, phoneNumber string, fullName string) (helpers.Pagination, error) {
	users, total, err := s.repo.GetAllWithPage(ctx, page, size, roleId, isActive, phoneNumber, fullName)
	if err != nil {
		return helpers.Pagination{}, err
	}

	return helpers.Pagiante(page, size, total, users), nil
}

func (s *userUsecase) GetAllUsersNoPage() ([]entities.ResUsersNoPage, error) {
	users, err := s.repo.GetAllNoPage()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *userUsecase) UpdateUser(ctx context.Context, user entities.ReqUser, fileHeader *multipart.FileHeader) error {

	if err := s.CheckVariableToUpdate(user.FirstName, user.LastName, user.ID, user.Email, user.PhoneNumber, user.Password, user.ConfirmPassword); err != nil {
		return err
	}

	if user.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	if user.ID != user.UpdatedBy {
		var userUpdater *entities.User
		var userUpdated *entities.User

		userUpdater, err := s.repo.GetRoleUserById(user.UpdatedBy)
		if err != nil {
			return err
		}

		userUpdated, err = s.repo.GetRoleUserById(user.ID)
		if err != nil {
			return err
		}

		if userUpdater.Role.Level < userUpdated.Role.Level {
			return fmt.Errorf("you do not have permission to update this user")
		}
	}

	if user.RoleId != nil {
		var userSelect *entities.User
		var roleSelect *entities.Role

		userSelect, err := s.repo.GetRoleUserById(user.UpdatedBy)
		if err != nil {
			return err
		}

		roleSelect, err = s.repo.GetRoleByRoleId(*user.RoleId)
		if err != nil {
			return err
		}

		if userSelect.Role.Level < roleSelect.Level {
			return fmt.Errorf("your role level (%d) must be higher than the role level (%d) you are attempting to update for user", userSelect.Role.Level, roleSelect.Level)
		}
	}

	avatar, err := s.repo.GetAvatarUserById(user.ID)
	if err != nil {
		return err
	}

	if fileHeader != nil {
		avatarURL, err := minio.UploadAvatarUpdate(fileHeader, avatar.Avatar)
		if err != nil {
			return err
		}
		user.Avatar = avatarURL
	}

	userStruct := entities.User{
		ID:                 user.ID,
		FirstName:          user.FirstName,
		LastName:           user.LastName,
		Email:              user.Email,
		PhoneNumber:        user.PhoneNumber,
		Password:           user.Password,
		Avatar:             user.Avatar,
		TwoFactorEnabled:   user.TwoFactorEnabled,
		TwoFactorVerified:  user.TwoFactorVerified,
		TwoFactorToken:     user.TwoFactorToken,
		TwoFactorAuthUrl:   user.TwoFactorAuthUrl,
		RoleId:             user.RoleId,
		ForgotPasswordCode: user.ForgotPasswordCode,
		IsActive:           user.IsActive,
		CreatedAt:          user.CreatedAt,
		CreatedBy:          user.CreatedBy,
		UpdatedAt:          user.UpdatedAt,
		UpdatedBy:          user.UpdatedBy,
		DeletedAt:          user.DeletedAt,
		DeletedBy:          user.DeletedBy,
	}

	if err := s.repo.Update(ctx, &userStruct); err != nil {
		return err
	}

	return nil
}

func (s *userUsecase) ChangePssword(ctx context.Context, reqPass entities.ReqChangePassword) error {
	if err := s.CheckVariableChangePassword(reqPass.NewPassword, reqPass.ConfirmNewPassword); err != nil {
		return err
	}

	if reqPass.NewPassword != "" {
		user, err := s.repo.GetProfileUser(reqPass.UserId)
		if err != nil {
			return err
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqPass.NewPassword)); err == nil {
			return fmt.Errorf("new password cannot be the same as the old password")
		}
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(reqPass.NewPassword), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		reqPass.NewPassword = string(hashedPassword)
	}

	if reqPass.UserId != reqPass.UpdatedBy {
		var userUpdater *entities.User
		var userUpdated *entities.User

		userUpdater, err := s.repo.GetRoleUserById(reqPass.UpdatedBy)
		if err != nil {
			return err
		}

		userUpdated, err = s.repo.GetRoleUserById(reqPass.UserId)
		if err != nil {
			return err
		}

		if userUpdater.Role.Level < userUpdated.Role.Level {
			return fmt.Errorf("you do not have permission to update this user")
		}
	}

	userStruct := entities.User{
		ID:        reqPass.UserId,
		Password:  reqPass.NewPassword,
		UpdatedBy: reqPass.UpdatedBy,
	}

	if err := s.repo.Update(ctx, &userStruct); err != nil {
		return err
	}

	return nil
}

func (s *userUsecase) DeleteUser(ctx context.Context, id uuid.UUID, deleteBy uuid.UUID) error {
	check, err := s.repo.IsSuperAdministrator(id)
	if err != nil {
		return err
	}

	userDel, err := s.repo.GetRoleUserById(id)
	if err != nil {
		return err
	}

	if !check {
		if deleteBy != id {
			var userDeleter *entities.User
			var userDeleted *entities.User
			userDeleter, err := s.repo.GetRoleUserById(deleteBy)
			if err != nil {
				return err
			}

			userDeleted, err = s.repo.GetRoleUserById(id)
			if err != nil {
				return err
			}

			if userDel.IsActive != nil && !*userDel.IsActive {
				if userDeleter.Role.Level <= userDeleted.Role.Level {
					return fmt.Errorf("you do not have permission to delete this user")
				}
			} else {
				return fmt.Errorf("can not delete user that is active")
			}
		}

		auth, isHave, _ := s.repo.CheckThisUserHaveDataInAuth(id)

		if isHave {
			err := s.repo.DeleteAuthAfterDeleteUser(auth.ID, deleteBy)
			if err != nil {
				return err
			}
		}

		if err := s.repo.Delete(ctx, id, deleteBy); err != nil {
			return err
		}

	} else {
		return fmt.Errorf("can not delete user that's have role super admin")
	}

	return nil
}

func (s *userUsecase) CheckVariableChangePassword(password string, confirmPassword string) error {
	if password == "" {
		return fmt.Errorf("not found field password")
	}

	if password != "" {
		if confirmPassword == "" {
			return fmt.Errorf("not found field confirmPassword")
		}
		if password != confirmPassword {
			return fmt.Errorf("password and confirmPassword not match")
		}
		if !isValidPassword(password) {
			return fmt.Errorf("invalid password. Please ensure your password contains at least 1 uppercase letter, 1 lowercase letter, 1 digit, 1 special character, and is between 8 and 16 characters long")
		}
	}

	return nil
}

func (s *userUsecase) CheckVariableToCreate(firstName string, lastName string, email string, phoneNumber string, password string, confirmPassword string) error {

	if firstName == "" {
		return fmt.Errorf("not found field firstName")
	}
	if lastName == "" {
		return fmt.Errorf("not found field lastName")
	}
	if email == "" {
		return fmt.Errorf("not found field email")
	}
	if phoneNumber == "" {
		return fmt.Errorf("not found field phoneNumber")
	}
	if password == "" {
		return fmt.Errorf("not found field password")
	}

	if phoneNumber != "" {
		if len(phoneNumber) < 10 || len(phoneNumber) > 10 {
			return fmt.Errorf("phoneNumber must contain 10 digits")
		}

		if !isAllDigits(phoneNumber) {
			return fmt.Errorf("phoneNumber is invalid")
		}

		phoneExists, err := s.repo.IsPhoneExists(phoneNumber)
		if err != nil {
			return err
		}

		if phoneExists {
			return fmt.Errorf("phone already exists")
		}
	}

	if email != "" {
		emailExists, err := s.repo.IsEmailExists(email)
		if err != nil {
			return err
		}

		if emailExists {
			return fmt.Errorf("email already exists")
		}

		if !isValidEmail(email) {
			return fmt.Errorf("%s is an invalid email", email)
		}

	}

	if password != "" {
		if confirmPassword == "" {
			return fmt.Errorf("not found field confirmPassword")
		}
		if password != confirmPassword {
			return fmt.Errorf("password and confirmPassword not match")
		}
		if !isValidPassword(password) {
			return fmt.Errorf("invalid password. Please ensure your password contains at least 1 uppercase letter, 1 lowercase letter, 1 digit, 1 special character, and is between 8 and 16 characters long")
		}
	}

	return nil
}

func (s *userUsecase) CheckVariableToUpdate(firstName string, lastName string, userId uuid.UUID, email string, phoneNumber string, password string, confirmPassword string) error {
	if firstName == "" {
		return fmt.Errorf("not found field firstName")
	}

	if lastName == "" {
		return fmt.Errorf("not found field lastName")
	}

	if email == "" {
		return fmt.Errorf("not found field email")
	}

	if phoneNumber == "" {
		return fmt.Errorf("not found field phoneNumber")
	}

	if password != "" {
		if confirmPassword == "" {
			return fmt.Errorf("not found field confirmPassword")
		}
		if password != confirmPassword {
			return fmt.Errorf("password and confirmPassword not match")
		}
		if !isValidPassword(password) {
			return fmt.Errorf("invalid password. Please ensure your password contains at least 1 uppercase letter, 1 lowercase letter, 1 digit, 1 special character, and is between 8 and 16 characters long")
		}
	}

	if phoneNumber != "" {
		if len(phoneNumber) < 10 || len(phoneNumber) > 10 {
			return fmt.Errorf("phoneNumber must contain 10 digits")
		}

		if !isAllDigits(phoneNumber) {
			return fmt.Errorf("phoneNumber is invalid")
		}
	}

	if email != "" {
		if !isValidEmail(email) {
			return fmt.Errorf("%s is an invalid email", email)
		}

		Exists, err := s.repo.IsEmailExistsForUpdate(email, userId)
		if err != nil {
			return err
		}

		if Exists {
			return errors.New("email already exists")
		}
	}

	if phoneNumber != "" {
		Exists, err := s.repo.IsPhoneExistsForUpdate(phoneNumber, userId)
		if err != nil {
			return err
		}

		if Exists {
			return errors.New("phone already exists")
		}
	}

	return nil
}

func isAllDigits(phone string) bool {
	for _, char := range phone {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	} else {
		return true
	}
}

func isValidPassword(password string) bool {
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_,.?":{}|<>]`).MatchString(password)
	lengthValid := len(password) >= 8 && len(password) <= 16

	return hasUpper && hasLower && hasDigit && hasSpecial && lengthValid
}
