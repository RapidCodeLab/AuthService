// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: proto/user.proto

package user_grpc

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

type UserStatus int32

const (
	UserStatus_NEW    UserStatus = 0
	UserStatus_ACTIVE UserStatus = 1
	UserStatus_BANNED UserStatus = 2
)

// Enum value maps for UserStatus.
var (
	UserStatus_name = map[int32]string{
		0: "NEW",
		1: "ACTIVE",
		2: "BANNED",
	}
	UserStatus_value = map[string]int32{
		"NEW":    0,
		"ACTIVE": 1,
		"BANNED": 2,
	}
)

func (x UserStatus) Enum() *UserStatus {
	p := new(UserStatus)
	*p = x
	return p
}

func (x UserStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_user_proto_enumTypes[0].Descriptor()
}

func (UserStatus) Type() protoreflect.EnumType {
	return &file_proto_user_proto_enumTypes[0]
}

func (x UserStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserStatus.Descriptor instead.
func (UserStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{0}
}

type UserRoles int32

const (
	UserRoles_SUPER   UserRoles = 0
	UserRoles_ADMIN   UserRoles = 1
	UserRoles_REGULAR UserRoles = 2
)

// Enum value maps for UserRoles.
var (
	UserRoles_name = map[int32]string{
		0: "SUPER",
		1: "ADMIN",
		2: "REGULAR",
	}
	UserRoles_value = map[string]int32{
		"SUPER":   0,
		"ADMIN":   1,
		"REGULAR": 2,
	}
)

func (x UserRoles) Enum() *UserRoles {
	p := new(UserRoles)
	*p = x
	return p
}

func (x UserRoles) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRoles) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_user_proto_enumTypes[1].Descriptor()
}

func (UserRoles) Type() protoreflect.EnumType {
	return &file_proto_user_proto_enumTypes[1]
}

func (x UserRoles) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRoles.Descriptor instead.
func (UserRoles) EnumDescriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{1}
}

type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Role     int64  `protobuf:"varint,3,opt,name=role,proto3" json:"role,omitempty"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserRequest) GetRole() int64 {
	if x != nil {
		return x.Role
	}
	return 0
}

type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email      string      `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	UserStatus UserStatus  `protobuf:"varint,3,opt,name=user_status,json=userStatus,proto3,enum=user_grpc.UserStatus" json:"user_status,omitempty"`
	UserRoles  []UserRoles `protobuf:"varint,4,rep,packed,name=user_roles,json=userRoles,proto3,enum=user_grpc.UserRoles" json:"user_roles,omitempty"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserResponse) GetUserStatus() UserStatus {
	if x != nil {
		return x.UserStatus
	}
	return UserStatus_NEW
}

func (x *UserResponse) GetUserRoles() []UserRoles {
	if x != nil {
		return x.UserRoles
	}
	return nil
}

var File_proto_user_proto protoreflect.FileDescriptor

var file_proto_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x22, 0x53, 0x0a,
	0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x72, 0x6f,
	0x6c, 0x65, 0x22, 0xa1, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x36, 0x0a, 0x0b, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x33, 0x0a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x18,
	0x04, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x52, 0x09, 0x75, 0x73, 0x65,
	0x72, 0x52, 0x6f, 0x6c, 0x65, 0x73, 0x2a, 0x2d, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x07, 0x0a, 0x03, 0x4e, 0x45, 0x57, 0x10, 0x00, 0x12, 0x0a, 0x0a,
	0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x42, 0x41, 0x4e,
	0x4e, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x2e, 0x0a, 0x09, 0x55, 0x73, 0x65, 0x72, 0x52, 0x6f, 0x6c,
	0x65, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x55, 0x50, 0x45, 0x52, 0x10, 0x00, 0x12, 0x09, 0x0a,
	0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x47, 0x55,
	0x4c, 0x41, 0x52, 0x10, 0x02, 0x32, 0x85, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x47, 0x50,
	0x52, 0x43, 0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70,
	0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d,
	0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0c, 0x5a,
	0x0a, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_user_proto_rawDescOnce sync.Once
	file_proto_user_proto_rawDescData = file_proto_user_proto_rawDesc
)

func file_proto_user_proto_rawDescGZIP() []byte {
	file_proto_user_proto_rawDescOnce.Do(func() {
		file_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_proto_rawDescData)
	})
	return file_proto_user_proto_rawDescData
}

var file_proto_user_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_user_proto_goTypes = []interface{}{
	(UserStatus)(0),      // 0: user_grpc.UserStatus
	(UserRoles)(0),       // 1: user_grpc.UserRoles
	(*UserRequest)(nil),  // 2: user_grpc.UserRequest
	(*UserResponse)(nil), // 3: user_grpc.UserResponse
}
var file_proto_user_proto_depIdxs = []int32{
	0, // 0: user_grpc.UserResponse.user_status:type_name -> user_grpc.UserStatus
	1, // 1: user_grpc.UserResponse.user_roles:type_name -> user_grpc.UserRoles
	2, // 2: user_grpc.UserGPRC.GetUser:input_type -> user_grpc.UserRequest
	2, // 3: user_grpc.UserGPRC.CreateUser:input_type -> user_grpc.UserRequest
	3, // 4: user_grpc.UserGPRC.GetUser:output_type -> user_grpc.UserResponse
	3, // 5: user_grpc.UserGPRC.CreateUser:output_type -> user_grpc.UserResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_user_proto_init() }
func file_proto_user_proto_init() {
	if File_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_user_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_user_proto_goTypes,
		DependencyIndexes: file_proto_user_proto_depIdxs,
		EnumInfos:         file_proto_user_proto_enumTypes,
		MessageInfos:      file_proto_user_proto_msgTypes,
	}.Build()
	File_proto_user_proto = out.File
	file_proto_user_proto_rawDesc = nil
	file_proto_user_proto_goTypes = nil
	file_proto_user_proto_depIdxs = nil
}
