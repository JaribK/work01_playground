// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: internal/proto/auth.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PermissionDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeatureId string `protobuf:"bytes,1,opt,name=feature_id,json=featureId,proto3" json:"feature_id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// string parent_menu_id = 3; // Optional, use string to represent UUID
	MenuIcon   string `protobuf:"bytes,4,opt,name=menu_icon,json=menuIcon,proto3" json:"menu_icon,omitempty"`
	MenuNameTh string `protobuf:"bytes,5,opt,name=menu_name_th,json=menuNameTh,proto3" json:"menu_name_th,omitempty"`
	MenuNameEn string `protobuf:"bytes,6,opt,name=menu_name_en,json=menuNameEn,proto3" json:"menu_name_en,omitempty"`
	MenuSlug   string `protobuf:"bytes,7,opt,name=menu_slug,json=menuSlug,proto3" json:"menu_slug,omitempty"`
	MenuSeqNo  string `protobuf:"bytes,8,opt,name=menu_seq_no,json=menuSeqNo,proto3" json:"menu_seq_no,omitempty"`
	IsActive   bool   `protobuf:"varint,9,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	// bool create_access = 10;
	// bool read_access = 11;
	// bool update_access = 12;
	// bool delete_access = 13;
	ParentMenuId *string `protobuf:"bytes,14,opt,name=parent_menu_id,json=parentMenuId,proto3,oneof" json:"parent_menu_id,omitempty"`
	IsAdd        bool    `protobuf:"varint,15,opt,name=is_add,json=isAdd,proto3" json:"is_add,omitempty"`
	IsView       bool    `protobuf:"varint,16,opt,name=is_view,json=isView,proto3" json:"is_view,omitempty"`
	IsEdit       bool    `protobuf:"varint,17,opt,name=is_edit,json=isEdit,proto3" json:"is_edit,omitempty"`
	IsDelete     bool    `protobuf:"varint,18,opt,name=is_delete,json=isDelete,proto3" json:"is_delete,omitempty"`
}

func (x *PermissionDTO) Reset() {
	*x = PermissionDTO{}
	mi := &file_internal_proto_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PermissionDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PermissionDTO) ProtoMessage() {}

func (x *PermissionDTO) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PermissionDTO.ProtoReflect.Descriptor instead.
func (*PermissionDTO) Descriptor() ([]byte, []int) {
	return file_internal_proto_auth_proto_rawDescGZIP(), []int{0}
}

func (x *PermissionDTO) GetFeatureId() string {
	if x != nil {
		return x.FeatureId
	}
	return ""
}

func (x *PermissionDTO) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PermissionDTO) GetMenuIcon() string {
	if x != nil {
		return x.MenuIcon
	}
	return ""
}

func (x *PermissionDTO) GetMenuNameTh() string {
	if x != nil {
		return x.MenuNameTh
	}
	return ""
}

func (x *PermissionDTO) GetMenuNameEn() string {
	if x != nil {
		return x.MenuNameEn
	}
	return ""
}

func (x *PermissionDTO) GetMenuSlug() string {
	if x != nil {
		return x.MenuSlug
	}
	return ""
}

func (x *PermissionDTO) GetMenuSeqNo() string {
	if x != nil {
		return x.MenuSeqNo
	}
	return ""
}

func (x *PermissionDTO) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *PermissionDTO) GetParentMenuId() string {
	if x != nil && x.ParentMenuId != nil {
		return *x.ParentMenuId
	}
	return ""
}

func (x *PermissionDTO) GetIsAdd() bool {
	if x != nil {
		return x.IsAdd
	}
	return false
}

func (x *PermissionDTO) GetIsView() bool {
	if x != nil {
		return x.IsView
	}
	return false
}

func (x *PermissionDTO) GetIsEdit() bool {
	if x != nil {
		return x.IsEdit
	}
	return false
}

func (x *PermissionDTO) GetIsDelete() bool {
	if x != nil {
		return x.IsDelete
	}
	return false
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId            string           `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Email             string           `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	FirstName         string           `protobuf:"bytes,3,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName          string           `protobuf:"bytes,4,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	PhoneNumber       string           `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Avatar            string           `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`
	RoleName          string           `protobuf:"bytes,7,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	RoleLevel         int32            `protobuf:"varint,8,opt,name=role_level,json=roleLevel,proto3" json:"role_level,omitempty"`
	TwoFactorEnabled  bool             `protobuf:"varint,9,opt,name=two_factor_enabled,json=twoFactorEnabled,proto3" json:"two_factor_enabled,omitempty"`
	TwoFactorVerified bool             `protobuf:"varint,10,opt,name=two_factor_verified,json=twoFactorVerified,proto3" json:"two_factor_verified,omitempty"`
	TwoFactorAuthUrl  string           `protobuf:"bytes,11,opt,name=two_factor_auth_url,json=twoFactorAuthUrl,proto3" json:"two_factor_auth_url,omitempty"`
	TwoFactorToken    string           `protobuf:"bytes,12,opt,name=two_factor_token,json=twoFactorToken,proto3" json:"two_factor_token,omitempty"`
	Permissions       []*PermissionDTO `protobuf:"bytes,13,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_internal_proto_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_internal_proto_auth_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *User) GetRoleLevel() int32 {
	if x != nil {
		return x.RoleLevel
	}
	return 0
}

func (x *User) GetTwoFactorEnabled() bool {
	if x != nil {
		return x.TwoFactorEnabled
	}
	return false
}

func (x *User) GetTwoFactorVerified() bool {
	if x != nil {
		return x.TwoFactorVerified
	}
	return false
}

func (x *User) GetTwoFactorAuthUrl() string {
	if x != nil {
		return x.TwoFactorAuthUrl
	}
	return ""
}

func (x *User) GetTwoFactorToken() string {
	if x != nil {
		return x.TwoFactorToken
	}
	return ""
}

func (x *User) GetPermissions() []*PermissionDTO {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type LoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginRequest) Reset() {
	*x = LoginRequest{}
	mi := &file_internal_proto_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginRequest) ProtoMessage() {}

func (x *LoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginRequest.ProtoReflect.Descriptor instead.
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_auth_proto_rawDescGZIP(), []int{2}
}

func (x *LoginRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message      string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	AccessToken  string `protobuf:"bytes,2,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,3,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	User         *User  `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *LoginResponse) Reset() {
	*x = LoginResponse{}
	mi := &file_internal_proto_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginResponse) ProtoMessage() {}

func (x *LoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginResponse.ProtoReflect.Descriptor instead.
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_auth_proto_rawDescGZIP(), []int{3}
}

func (x *LoginResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *LoginResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *LoginResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *LoginResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type LogoutRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token  string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *LogoutRequest) Reset() {
	*x = LogoutRequest{}
	mi := &file_internal_proto_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutRequest) ProtoMessage() {}

func (x *LogoutRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutRequest.ProtoReflect.Descriptor instead.
func (*LogoutRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_auth_proto_rawDescGZIP(), []int{4}
}

func (x *LogoutRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LogoutRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type LogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *LogoutResponse) Reset() {
	*x = LogoutResponse{}
	mi := &file_internal_proto_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutResponse) ProtoMessage() {}

func (x *LogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutResponse.ProtoReflect.Descriptor instead.
func (*LogoutResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_auth_proto_rawDescGZIP(), []int{5}
}

func (x *LogoutResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RefreshTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *RefreshTokenRequest) Reset() {
	*x = RefreshTokenRequest{}
	mi := &file_internal_proto_auth_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenRequest) ProtoMessage() {}

func (x *RefreshTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_auth_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokenRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_auth_proto_rawDescGZIP(), []int{6}
}

func (x *RefreshTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RefreshTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
}

func (x *RefreshTokenResponse) Reset() {
	*x = RefreshTokenResponse{}
	mi := &file_internal_proto_auth_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenResponse) ProtoMessage() {}

func (x *RefreshTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_auth_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenResponse.ProtoReflect.Descriptor instead.
func (*RefreshTokenResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_auth_proto_rawDescGZIP(), []int{7}
}

func (x *RefreshTokenResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

var File_internal_proto_auth_proto protoreflect.FileDescriptor

var file_internal_proto_auth_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xa1, 0x03, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x44, 0x54, 0x4f, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6e, 0x75, 0x5f,
	0x69, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6e, 0x75,
	0x49, 0x63, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x5f, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x6e, 0x75,
	0x4e, 0x61, 0x6d, 0x65, 0x54, 0x68, 0x12, 0x20, 0x0a, 0x0c, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x5f, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65,
	0x6e, 0x75, 0x4e, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6e, 0x75,
	0x5f, 0x73, 0x6c, 0x75, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6e,
	0x75, 0x53, 0x6c, 0x75, 0x67, 0x12, 0x1e, 0x0a, 0x0b, 0x6d, 0x65, 0x6e, 0x75, 0x5f, 0x73, 0x65,
	0x71, 0x5f, 0x6e, 0x6f, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x6e, 0x75,
	0x53, 0x65, 0x71, 0x4e, 0x6f, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x65, 0x12, 0x29, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x6e,
	0x75, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0c, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a,
	0x06, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69,
	0x73, 0x41, 0x64, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x76, 0x69, 0x65, 0x77, 0x18,
	0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x56, 0x69, 0x65, 0x77, 0x12, 0x17, 0x0a,
	0x07, 0x69, 0x73, 0x5f, 0x65, 0x64, 0x69, 0x74, 0x18, 0x11, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06,
	0x69, 0x73, 0x45, 0x64, 0x69, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x18, 0x12, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x6d,
	0x65, 0x6e, 0x75, 0x5f, 0x69, 0x64, 0x22, 0xd7, 0x03, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a,
	0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x4c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x2c, 0x0a, 0x12, 0x74, 0x77, 0x6f, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f,
	0x65, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x74,
	0x77, 0x6f, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x64, 0x12,
	0x2e, 0x0a, 0x13, 0x74, 0x77, 0x6f, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x74, 0x77,
	0x6f, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12,
	0x2d, 0x0a, 0x13, 0x74, 0x77, 0x6f, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x77,
	0x6f, 0x46, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x41, 0x75, 0x74, 0x68, 0x55, 0x72, 0x6c, 0x12, 0x28,
	0x0a, 0x10, 0x74, 0x77, 0x6f, 0x5f, 0x66, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x77, 0x6f, 0x46, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x36, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0d, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x44, 0x54, 0x4f, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x22, 0x40, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x22, 0x92, 0x01, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x3e, 0x0a, 0x0d, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x3a, 0x0a, 0x13, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x39, 0x0a, 0x14, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xc3, 0x01, 0x0a, 0x0d, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x35, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x0e, 0x5a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x30, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_auth_proto_rawDescOnce sync.Once
	file_internal_proto_auth_proto_rawDescData = file_internal_proto_auth_proto_rawDesc
)

func file_internal_proto_auth_proto_rawDescGZIP() []byte {
	file_internal_proto_auth_proto_rawDescOnce.Do(func() {
		file_internal_proto_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_auth_proto_rawDescData)
	})
	return file_internal_proto_auth_proto_rawDescData
}

var file_internal_proto_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_internal_proto_auth_proto_goTypes = []any{
	(*PermissionDTO)(nil),        // 0: proto.PermissionDTO
	(*User)(nil),                 // 1: proto.User
	(*LoginRequest)(nil),         // 2: proto.LoginRequest
	(*LoginResponse)(nil),        // 3: proto.LoginResponse
	(*LogoutRequest)(nil),        // 4: proto.LogoutRequest
	(*LogoutResponse)(nil),       // 5: proto.LogoutResponse
	(*RefreshTokenRequest)(nil),  // 6: proto.RefreshTokenRequest
	(*RefreshTokenResponse)(nil), // 7: proto.RefreshTokenResponse
}
var file_internal_proto_auth_proto_depIdxs = []int32{
	0, // 0: proto.User.permissions:type_name -> proto.PermissionDTO
	1, // 1: proto.LoginResponse.user:type_name -> proto.User
	2, // 2: proto.Authorization.Login:input_type -> proto.LoginRequest
	4, // 3: proto.Authorization.Logout:input_type -> proto.LogoutRequest
	6, // 4: proto.Authorization.RefreshToken:input_type -> proto.RefreshTokenRequest
	3, // 5: proto.Authorization.Login:output_type -> proto.LoginResponse
	5, // 6: proto.Authorization.Logout:output_type -> proto.LogoutResponse
	7, // 7: proto.Authorization.RefreshToken:output_type -> proto.RefreshTokenResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_proto_auth_proto_init() }
func file_internal_proto_auth_proto_init() {
	if File_internal_proto_auth_proto != nil {
		return
	}
	file_internal_proto_auth_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_proto_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_auth_proto_goTypes,
		DependencyIndexes: file_internal_proto_auth_proto_depIdxs,
		MessageInfos:      file_internal_proto_auth_proto_msgTypes,
	}.Build()
	File_internal_proto_auth_proto = out.File
	file_internal_proto_auth_proto_rawDesc = nil
	file_internal_proto_auth_proto_goTypes = nil
	file_internal_proto_auth_proto_depIdxs = nil
}
