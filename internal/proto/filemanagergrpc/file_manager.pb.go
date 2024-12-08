// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: internal/proto/file_manager.proto

package filemanagergrpc

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

type DeleteFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileUrl string `protobuf:"bytes,1,opt,name=file_url,json=fileUrl,proto3" json:"file_url,omitempty"`
}

func (x *DeleteFileReq) Reset() {
	*x = DeleteFileReq{}
	mi := &file_internal_proto_file_manager_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileReq) ProtoMessage() {}

func (x *DeleteFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_file_manager_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileReq.ProtoReflect.Descriptor instead.
func (*DeleteFileReq) Descriptor() ([]byte, []int) {
	return file_internal_proto_file_manager_proto_rawDescGZIP(), []int{0}
}

func (x *DeleteFileReq) GetFileUrl() string {
	if x != nil {
		return x.FileUrl
	}
	return ""
}

type DeleteFileRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteFileRes) Reset() {
	*x = DeleteFileRes{}
	mi := &file_internal_proto_file_manager_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteFileRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFileRes) ProtoMessage() {}

func (x *DeleteFileRes) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_file_manager_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFileRes.ProtoReflect.Descriptor instead.
func (*DeleteFileRes) Descriptor() ([]byte, []int) {
	return file_internal_proto_file_manager_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteFileRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type UploadFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileChunk []byte `protobuf:"bytes,1,opt,name=file_chunk,json=fileChunk,proto3" json:"file_chunk,omitempty"`
	FileName  string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	FilePath  string `protobuf:"bytes,3,opt,name=filePath,proto3" json:"filePath,omitempty"`
}

func (x *UploadFileReq) Reset() {
	*x = UploadFileReq{}
	mi := &file_internal_proto_file_manager_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileReq) ProtoMessage() {}

func (x *UploadFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_file_manager_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileReq.ProtoReflect.Descriptor instead.
func (*UploadFileReq) Descriptor() ([]byte, []int) {
	return file_internal_proto_file_manager_proto_rawDescGZIP(), []int{2}
}

func (x *UploadFileReq) GetFileChunk() []byte {
	if x != nil {
		return x.FileChunk
	}
	return nil
}

func (x *UploadFileReq) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadFileReq) GetFilePath() string {
	if x != nil {
		return x.FilePath
	}
	return ""
}

type UploadFileRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileName  string `protobuf:"bytes,1,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
	Thumbnail string `protobuf:"bytes,2,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
}

func (x *UploadFileRes) Reset() {
	*x = UploadFileRes{}
	mi := &file_internal_proto_file_manager_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UploadFileRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UploadFileRes) ProtoMessage() {}

func (x *UploadFileRes) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_file_manager_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UploadFileRes.ProtoReflect.Descriptor instead.
func (*UploadFileRes) Descriptor() ([]byte, []int) {
	return file_internal_proto_file_manager_proto_rawDescGZIP(), []int{3}
}

func (x *UploadFileRes) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

func (x *UploadFileRes) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

var File_internal_proto_file_manager_proto protoreflect.FileDescriptor

var file_internal_proto_file_manager_proto_rawDesc = []byte{
	0x0a, 0x21, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2a, 0x0a, 0x0d, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x66,
	0x69, 0x6c, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x27, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x67, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71,
	0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x12,
	0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x74, 0x68, 0x22, 0x4a, 0x0a, 0x0d, 0x55, 0x70, 0x6c, 0x6f,
	0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62,
	0x6e, 0x61, 0x69, 0x6c, 0x32, 0x81, 0x01, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x4d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x0a, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61,
	0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x12, 0x38,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x14, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x42, 0x20, 0x5a, 0x1e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_internal_proto_file_manager_proto_rawDescOnce sync.Once
	file_internal_proto_file_manager_proto_rawDescData = file_internal_proto_file_manager_proto_rawDesc
)

func file_internal_proto_file_manager_proto_rawDescGZIP() []byte {
	file_internal_proto_file_manager_proto_rawDescOnce.Do(func() {
		file_internal_proto_file_manager_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_file_manager_proto_rawDescData)
	})
	return file_internal_proto_file_manager_proto_rawDescData
}

var file_internal_proto_file_manager_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_proto_file_manager_proto_goTypes = []any{
	(*DeleteFileReq)(nil), // 0: proto.DeleteFileReq
	(*DeleteFileRes)(nil), // 1: proto.DeleteFileRes
	(*UploadFileReq)(nil), // 2: proto.UploadFileReq
	(*UploadFileRes)(nil), // 3: proto.UploadFileRes
}
var file_internal_proto_file_manager_proto_depIdxs = []int32{
	2, // 0: proto.FileManager.UploadFile:input_type -> proto.UploadFileReq
	0, // 1: proto.FileManager.DeleteFile:input_type -> proto.DeleteFileReq
	3, // 2: proto.FileManager.UploadFile:output_type -> proto.UploadFileRes
	1, // 3: proto.FileManager.DeleteFile:output_type -> proto.DeleteFileRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_file_manager_proto_init() }
func file_internal_proto_file_manager_proto_init() {
	if File_internal_proto_file_manager_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_proto_file_manager_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_file_manager_proto_goTypes,
		DependencyIndexes: file_internal_proto_file_manager_proto_depIdxs,
		MessageInfos:      file_internal_proto_file_manager_proto_msgTypes,
	}.Build()
	File_internal_proto_file_manager_proto = out.File
	file_internal_proto_file_manager_proto_rawDesc = nil
	file_internal_proto_file_manager_proto_goTypes = nil
	file_internal_proto_file_manager_proto_depIdxs = nil
}
