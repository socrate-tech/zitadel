// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: zitadel/protoc_gen_zitadel/v2/options.proto

package protoc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Options struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthOption   *AuthOption         `protobuf:"bytes,1,opt,name=auth_option,json=authOption,proto3" json:"auth_option,omitempty"`
	HttpResponse *CustomHTTPResponse `protobuf:"bytes,2,opt,name=http_response,json=httpResponse,proto3" json:"http_response,omitempty"`
}

func (x *Options) Reset() {
	*x = Options{}
	mi := &file_zitadel_protoc_gen_zitadel_v2_options_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Options) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Options) ProtoMessage() {}

func (x *Options) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_protoc_gen_zitadel_v2_options_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Options.ProtoReflect.Descriptor instead.
func (*Options) Descriptor() ([]byte, []int) {
	return file_zitadel_protoc_gen_zitadel_v2_options_proto_rawDescGZIP(), []int{0}
}

func (x *Options) GetAuthOption() *AuthOption {
	if x != nil {
		return x.AuthOption
	}
	return nil
}

func (x *Options) GetHttpResponse() *CustomHTTPResponse {
	if x != nil {
		return x.HttpResponse
	}
	return nil
}

type AuthOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permission string `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
	OrgField   string `protobuf:"bytes,3,opt,name=org_field,json=orgField,proto3" json:"org_field,omitempty"`
}

func (x *AuthOption) Reset() {
	*x = AuthOption{}
	mi := &file_zitadel_protoc_gen_zitadel_v2_options_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AuthOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthOption) ProtoMessage() {}

func (x *AuthOption) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_protoc_gen_zitadel_v2_options_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthOption.ProtoReflect.Descriptor instead.
func (*AuthOption) Descriptor() ([]byte, []int) {
	return file_zitadel_protoc_gen_zitadel_v2_options_proto_rawDescGZIP(), []int{1}
}

func (x *AuthOption) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *AuthOption) GetOrgField() string {
	if x != nil {
		return x.OrgField
	}
	return ""
}

type CustomHTTPResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SuccessCode int32 `protobuf:"varint,1,opt,name=success_code,json=successCode,proto3" json:"success_code,omitempty"`
}

func (x *CustomHTTPResponse) Reset() {
	*x = CustomHTTPResponse{}
	mi := &file_zitadel_protoc_gen_zitadel_v2_options_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CustomHTTPResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomHTTPResponse) ProtoMessage() {}

func (x *CustomHTTPResponse) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_protoc_gen_zitadel_v2_options_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomHTTPResponse.ProtoReflect.Descriptor instead.
func (*CustomHTTPResponse) Descriptor() ([]byte, []int) {
	return file_zitadel_protoc_gen_zitadel_v2_options_proto_rawDescGZIP(), []int{2}
}

func (x *CustomHTTPResponse) GetSuccessCode() int32 {
	if x != nil {
		return x.SuccessCode
	}
	return 0
}

var file_zitadel_protoc_gen_zitadel_v2_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*Options)(nil),
		Field:         50001,
		Name:          "zitadel.protoc_gen_zitadel.v2.options",
		Tag:           "bytes,50001,opt,name=options",
		Filename:      "zitadel/protoc_gen_zitadel/v2/options.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional zitadel.protoc_gen_zitadel.v2.Options options = 50001;
	E_Options = &file_zitadel_protoc_gen_zitadel_v2_options_proto_extTypes[0]
)

var File_zitadel_protoc_gen_zitadel_v2_options_proto protoreflect.FileDescriptor

var file_zitadel_protoc_gen_zitadel_v2_options_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x7a,
	0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65,
	0x6e, 0x5f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x32, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad,
	0x01, 0x0a, 0x07, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4a, 0x0a, 0x0b, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x29, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x32, 0x2e,
	0x41, 0x75, 0x74, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x56, 0x0a, 0x0d, 0x68, 0x74, 0x74, 0x70, 0x5f, 0x72,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x31, 0x2e,
	0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67,
	0x65, 0x6e, 0x5f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x52, 0x0c, 0x68, 0x74, 0x74, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4f,
	0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a,
	0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09,
	0x6f, 0x72, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6f, 0x72, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22,
	0x37, 0x0a, 0x12, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x48, 0x54, 0x54, 0x50, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x3a, 0x62, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xd1, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x7a, 0x69,
	0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x5f, 0x67, 0x65, 0x6e,
	0x5f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x32, 0x2e, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x36, 0x5a, 0x34,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64,
	0x65, 0x6c, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x76, 0x32, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zitadel_protoc_gen_zitadel_v2_options_proto_rawDescOnce sync.Once
	file_zitadel_protoc_gen_zitadel_v2_options_proto_rawDescData = file_zitadel_protoc_gen_zitadel_v2_options_proto_rawDesc
)

func file_zitadel_protoc_gen_zitadel_v2_options_proto_rawDescGZIP() []byte {
	file_zitadel_protoc_gen_zitadel_v2_options_proto_rawDescOnce.Do(func() {
		file_zitadel_protoc_gen_zitadel_v2_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_zitadel_protoc_gen_zitadel_v2_options_proto_rawDescData)
	})
	return file_zitadel_protoc_gen_zitadel_v2_options_proto_rawDescData
}

var file_zitadel_protoc_gen_zitadel_v2_options_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_zitadel_protoc_gen_zitadel_v2_options_proto_goTypes = []any{
	(*Options)(nil),                    // 0: zitadel.protoc_gen_zitadel.v2.Options
	(*AuthOption)(nil),                 // 1: zitadel.protoc_gen_zitadel.v2.AuthOption
	(*CustomHTTPResponse)(nil),         // 2: zitadel.protoc_gen_zitadel.v2.CustomHTTPResponse
	(*descriptorpb.MethodOptions)(nil), // 3: google.protobuf.MethodOptions
}
var file_zitadel_protoc_gen_zitadel_v2_options_proto_depIdxs = []int32{
	1, // 0: zitadel.protoc_gen_zitadel.v2.Options.auth_option:type_name -> zitadel.protoc_gen_zitadel.v2.AuthOption
	2, // 1: zitadel.protoc_gen_zitadel.v2.Options.http_response:type_name -> zitadel.protoc_gen_zitadel.v2.CustomHTTPResponse
	3, // 2: zitadel.protoc_gen_zitadel.v2.options:extendee -> google.protobuf.MethodOptions
	0, // 3: zitadel.protoc_gen_zitadel.v2.options:type_name -> zitadel.protoc_gen_zitadel.v2.Options
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	3, // [3:4] is the sub-list for extension type_name
	2, // [2:3] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_zitadel_protoc_gen_zitadel_v2_options_proto_init() }
func file_zitadel_protoc_gen_zitadel_v2_options_proto_init() {
	if File_zitadel_protoc_gen_zitadel_v2_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zitadel_protoc_gen_zitadel_v2_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_zitadel_protoc_gen_zitadel_v2_options_proto_goTypes,
		DependencyIndexes: file_zitadel_protoc_gen_zitadel_v2_options_proto_depIdxs,
		MessageInfos:      file_zitadel_protoc_gen_zitadel_v2_options_proto_msgTypes,
		ExtensionInfos:    file_zitadel_protoc_gen_zitadel_v2_options_proto_extTypes,
	}.Build()
	File_zitadel_protoc_gen_zitadel_v2_options_proto = out.File
	file_zitadel_protoc_gen_zitadel_v2_options_proto_rawDesc = nil
	file_zitadel_protoc_gen_zitadel_v2_options_proto_goTypes = nil
	file_zitadel_protoc_gen_zitadel_v2_options_proto_depIdxs = nil
}
