// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: internal/proto/test_user.proto

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

type CreateUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName   string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	PhoneNumber string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password    string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *CreateUserReq) Reset() {
	*x = CreateUserReq{}
	mi := &file_internal_proto_test_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserReq) ProtoMessage() {}

func (x *CreateUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_test_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserReq.ProtoReflect.Descriptor instead.
func (*CreateUserReq) Descriptor() ([]byte, []int) {
	return file_internal_proto_test_user_proto_rawDescGZIP(), []int{0}
}

func (x *CreateUserReq) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreateUserReq) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *CreateUserReq) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *CreateUserReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type CreateUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *CreateUserRes) Reset() {
	*x = CreateUserRes{}
	mi := &file_internal_proto_test_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRes) ProtoMessage() {}

func (x *CreateUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_test_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRes.ProtoReflect.Descriptor instead.
func (*CreateUserRes) Descriptor() ([]byte, []int) {
	return file_internal_proto_test_user_proto_rawDescGZIP(), []int{1}
}

func (x *CreateUserRes) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type GetUserByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetUserByIdReq) Reset() {
	*x = GetUserByIdReq{}
	mi := &file_internal_proto_test_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdReq) ProtoMessage() {}

func (x *GetUserByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_test_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdReq.ProtoReflect.Descriptor instead.
func (*GetUserByIdReq) Descriptor() ([]byte, []int) {
	return file_internal_proto_test_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserByIdReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetUserByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FirstName   string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	IsActive    bool   `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Avatar      string `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar,omitempty"`
	RoleName    string `protobuf:"bytes,8,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	// string role_id = 9;
	RoleId *string `protobuf:"bytes,10,opt,name=role_id,json=roleId,proto3,oneof" json:"role_id,omitempty"`
}

func (x *GetUserByIdRes) Reset() {
	*x = GetUserByIdRes{}
	mi := &file_internal_proto_test_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserByIdRes) ProtoMessage() {}

func (x *GetUserByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_test_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserByIdRes.ProtoReflect.Descriptor instead.
func (*GetUserByIdRes) Descriptor() ([]byte, []int) {
	return file_internal_proto_test_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetUserByIdRes) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetUserByIdRes) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetUserByIdRes) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetUserByIdRes) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUserByIdRes) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *GetUserByIdRes) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *GetUserByIdRes) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *GetUserByIdRes) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *GetUserByIdRes) GetRoleId() string {
	if x != nil && x.RoleId != nil {
		return *x.RoleId
	}
	return ""
}

type GetAllUserReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page     int32  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size     int32  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	RoleId   string `protobuf:"bytes,3,opt,name=role_id,json=roleId,proto3" json:"role_id,omitempty"`
	IsActive string `protobuf:"bytes,4,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *GetAllUserReq) Reset() {
	*x = GetAllUserReq{}
	mi := &file_internal_proto_test_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllUserReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserReq) ProtoMessage() {}

func (x *GetAllUserReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_test_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserReq.ProtoReflect.Descriptor instead.
func (*GetAllUserReq) Descriptor() ([]byte, []int) {
	return file_internal_proto_test_user_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllUserReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllUserReq) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *GetAllUserReq) GetRoleId() string {
	if x != nil {
		return x.RoleId
	}
	return ""
}

func (x *GetAllUserReq) GetIsActive() string {
	if x != nil {
		return x.IsActive
	}
	return ""
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CurrentPage int32 `protobuf:"varint,1,opt,name=current_page,json=currentPage,proto3" json:"current_page,omitempty"`
	TotalPage   int32 `protobuf:"varint,2,opt,name=total_page,json=totalPage,proto3" json:"total_page,omitempty"`
	Size        int32 `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	TotalItems  int32 `protobuf:"varint,4,opt,name=total_items,json=totalItems,proto3" json:"total_items,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	mi := &file_internal_proto_test_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_test_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_internal_proto_test_user_proto_rawDescGZIP(), []int{5}
}

func (x *Pagination) GetCurrentPage() int32 {
	if x != nil {
		return x.CurrentPage
	}
	return 0
}

func (x *Pagination) GetTotalPage() int32 {
	if x != nil {
		return x.TotalPage
	}
	return 0
}

func (x *Pagination) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *Pagination) GetTotalItems() int32 {
	if x != nil {
		return x.TotalItems
	}
	return 0
}

type UsersDTO struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FirstName   string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	IsActive    bool   `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
	Avatar      string `protobuf:"bytes,7,opt,name=avatar,proto3" json:"avatar,omitempty"`
	RoleName    string `protobuf:"bytes,8,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
}

func (x *UsersDTO) Reset() {
	*x = UsersDTO{}
	mi := &file_internal_proto_test_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UsersDTO) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersDTO) ProtoMessage() {}

func (x *UsersDTO) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_test_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersDTO.ProtoReflect.Descriptor instead.
func (*UsersDTO) Descriptor() ([]byte, []int) {
	return file_internal_proto_test_user_proto_rawDescGZIP(), []int{6}
}

func (x *UsersDTO) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UsersDTO) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UsersDTO) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UsersDTO) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UsersDTO) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UsersDTO) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

func (x *UsersDTO) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UsersDTO) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

type GetAllUserRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// string user_id = 1;
	// string first_name = 2;
	// string last_name = 3;
	// string email = 4;
	// string phone_number = 5;
	// bool is_active = 6;
	// string avatar = 7;
	// string role_name = 8;
	Users      []*UsersDTO `protobuf:"bytes,9,rep,name=users,proto3" json:"users,omitempty"`
	Pagination *Pagination `protobuf:"bytes,10,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (x *GetAllUserRes) Reset() {
	*x = GetAllUserRes{}
	mi := &file_internal_proto_test_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllUserRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserRes) ProtoMessage() {}

func (x *GetAllUserRes) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_test_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserRes.ProtoReflect.Descriptor instead.
func (*GetAllUserRes) Descriptor() ([]byte, []int) {
	return file_internal_proto_test_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetAllUserRes) GetUsers() []*UsersDTO {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *GetAllUserRes) GetPagination() *Pagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type UpdateUserByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FirstName   string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName    string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	PhoneNumber string `protobuf:"bytes,5,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	IsActive    bool   `protobuf:"varint,6,opt,name=is_active,json=isActive,proto3" json:"is_active,omitempty"`
}

func (x *UpdateUserByIdReq) Reset() {
	*x = UpdateUserByIdReq{}
	mi := &file_internal_proto_test_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserByIdReq) ProtoMessage() {}

func (x *UpdateUserByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_test_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserByIdReq.ProtoReflect.Descriptor instead.
func (*UpdateUserByIdReq) Descriptor() ([]byte, []int) {
	return file_internal_proto_test_user_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateUserByIdReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UpdateUserByIdReq) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UpdateUserByIdReq) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *UpdateUserByIdReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUserByIdReq) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *UpdateUserByIdReq) GetIsActive() bool {
	if x != nil {
		return x.IsActive
	}
	return false
}

type UpdateUserByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *UpdateUserByIdRes) Reset() {
	*x = UpdateUserByIdRes{}
	mi := &file_internal_proto_test_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserByIdRes) ProtoMessage() {}

func (x *UpdateUserByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_test_user_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserByIdRes.ProtoReflect.Descriptor instead.
func (*UpdateUserByIdRes) Descriptor() ([]byte, []int) {
	return file_internal_proto_test_user_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateUserByIdRes) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type DeleteUserByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteUserByIdReq) Reset() {
	*x = DeleteUserByIdReq{}
	mi := &file_internal_proto_test_user_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserByIdReq) ProtoMessage() {}

func (x *DeleteUserByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_test_user_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserByIdReq.ProtoReflect.Descriptor instead.
func (*DeleteUserByIdReq) Descriptor() ([]byte, []int) {
	return file_internal_proto_test_user_proto_rawDescGZIP(), []int{10}
}

func (x *DeleteUserByIdReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type DeleteUserByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DeleteUserByIdRes) Reset() {
	*x = DeleteUserByIdRes{}
	mi := &file_internal_proto_test_user_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserByIdRes) ProtoMessage() {}

func (x *DeleteUserByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_test_user_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserByIdRes.ProtoReflect.Descriptor instead.
func (*DeleteUserByIdRes) Descriptor() ([]byte, []int) {
	return file_internal_proto_test_user_proto_rawDescGZIP(), []int{11}
}

func (x *DeleteUserByIdRes) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_internal_proto_test_user_proto protoreflect.FileDescriptor

var file_internal_proto_test_user_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x27, 0x0a, 0x0d, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x9a,
	0x02, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1c, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x22, 0x6d, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6c, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x83, 0x01, 0x0a, 0x0a, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0b, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73,
	0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x22, 0xea, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x73, 0x44, 0x54, 0x4f, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x69, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x25,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x44, 0x54, 0x4f, 0x52, 0x05,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbe, 0x01, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x09,
	0x69, 0x73, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x08, 0x69, 0x73, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x2b, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x2b, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x32, 0xce, 0x02, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x47, 0x72, 0x70, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12,
	0x3b, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x15,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x12, 0x44, 0x0a, 0x0e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x12, 0x18,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x73, 0x42, 0x0e, 0x5a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x30, 0x31, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_test_user_proto_rawDescOnce sync.Once
	file_internal_proto_test_user_proto_rawDescData = file_internal_proto_test_user_proto_rawDesc
)

func file_internal_proto_test_user_proto_rawDescGZIP() []byte {
	file_internal_proto_test_user_proto_rawDescOnce.Do(func() {
		file_internal_proto_test_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_test_user_proto_rawDescData)
	})
	return file_internal_proto_test_user_proto_rawDescData
}

var file_internal_proto_test_user_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_internal_proto_test_user_proto_goTypes = []any{
	(*CreateUserReq)(nil),     // 0: proto.CreateUserReq
	(*CreateUserRes)(nil),     // 1: proto.CreateUserRes
	(*GetUserByIdReq)(nil),    // 2: proto.GetUserByIdReq
	(*GetUserByIdRes)(nil),    // 3: proto.GetUserByIdRes
	(*GetAllUserReq)(nil),     // 4: proto.GetAllUserReq
	(*Pagination)(nil),        // 5: proto.Pagination
	(*UsersDTO)(nil),          // 6: proto.UsersDTO
	(*GetAllUserRes)(nil),     // 7: proto.GetAllUserRes
	(*UpdateUserByIdReq)(nil), // 8: proto.UpdateUserByIdReq
	(*UpdateUserByIdRes)(nil), // 9: proto.UpdateUserByIdRes
	(*DeleteUserByIdReq)(nil), // 10: proto.DeleteUserByIdReq
	(*DeleteUserByIdRes)(nil), // 11: proto.DeleteUserByIdRes
}
var file_internal_proto_test_user_proto_depIdxs = []int32{
	6,  // 0: proto.GetAllUserRes.users:type_name -> proto.UsersDTO
	5,  // 1: proto.GetAllUserRes.pagination:type_name -> proto.Pagination
	0,  // 2: proto.UserGrpcService.CreateUser:input_type -> proto.CreateUserReq
	2,  // 3: proto.UserGrpcService.GetUserById:input_type -> proto.GetUserByIdReq
	4,  // 4: proto.UserGrpcService.GetAllUser:input_type -> proto.GetAllUserReq
	8,  // 5: proto.UserGrpcService.UpdateUserById:input_type -> proto.UpdateUserByIdReq
	10, // 6: proto.UserGrpcService.DeleteUserById:input_type -> proto.DeleteUserByIdReq
	1,  // 7: proto.UserGrpcService.CreateUser:output_type -> proto.CreateUserRes
	3,  // 8: proto.UserGrpcService.GetUserById:output_type -> proto.GetUserByIdRes
	7,  // 9: proto.UserGrpcService.GetAllUser:output_type -> proto.GetAllUserRes
	9,  // 10: proto.UserGrpcService.UpdateUserById:output_type -> proto.UpdateUserByIdRes
	11, // 11: proto.UserGrpcService.DeleteUserById:output_type -> proto.DeleteUserByIdRes
	7,  // [7:12] is the sub-list for method output_type
	2,  // [2:7] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_internal_proto_test_user_proto_init() }
func file_internal_proto_test_user_proto_init() {
	if File_internal_proto_test_user_proto != nil {
		return
	}
	file_internal_proto_test_user_proto_msgTypes[3].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_proto_test_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_test_user_proto_goTypes,
		DependencyIndexes: file_internal_proto_test_user_proto_depIdxs,
		MessageInfos:      file_internal_proto_test_user_proto_msgTypes,
	}.Build()
	File_internal_proto_test_user_proto = out.File
	file_internal_proto_test_user_proto_rawDesc = nil
	file_internal_proto_test_user_proto_goTypes = nil
	file_internal_proto_test_user_proto_depIdxs = nil
}