// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: error_user.proto

package v1

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

type ErrorType int32

const (
	ErrorType_UserNotFound ErrorType = 0
	ErrorType_ArgError     ErrorType = 1
)

// Enum value maps for ErrorType.
var (
	ErrorType_name = map[int32]string{
		0: "UserNotFound",
		1: "ArgError",
	}
	ErrorType_value = map[string]int32{
		"UserNotFound": 0,
		"ArgError":     1,
	}
)

func (x ErrorType) Enum() *ErrorType {
	p := new(ErrorType)
	*p = x
	return p
}

func (x ErrorType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorType) Descriptor() protoreflect.EnumDescriptor {
	return file_error_user_proto_enumTypes[0].Descriptor()
}

func (ErrorType) Type() protoreflect.EnumType {
	return &file_error_user_proto_enumTypes[0]
}

func (x ErrorType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorType.Descriptor instead.
func (ErrorType) EnumDescriptor() ([]byte, []int) {
	return file_error_user_proto_rawDescGZIP(), []int{0}
}

var File_error_user_proto protoreflect.FileDescriptor

var file_error_user_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2a,
	0x2b, 0x0a, 0x09, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x10, 0x0a, 0x0c,
	0x55, 0x73, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x00, 0x12, 0x0c,
	0x0a, 0x08, 0x41, 0x72, 0x67, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x10, 0x01, 0x42, 0x24, 0x0a, 0x0b,
	0x61, 0x70, 0x69, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x13, 0x62,
	0x6c, 0x6f, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_error_user_proto_rawDescOnce sync.Once
	file_error_user_proto_rawDescData = file_error_user_proto_rawDesc
)

func file_error_user_proto_rawDescGZIP() []byte {
	file_error_user_proto_rawDescOnce.Do(func() {
		file_error_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_error_user_proto_rawDescData)
	})
	return file_error_user_proto_rawDescData
}

var file_error_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_error_user_proto_goTypes = []interface{}{
	(ErrorType)(0), // 0: api.user.v1.ErrorType
}
var file_error_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_error_user_proto_init() }
func file_error_user_proto_init() {
	if File_error_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_error_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_error_user_proto_goTypes,
		DependencyIndexes: file_error_user_proto_depIdxs,
		EnumInfos:         file_error_user_proto_enumTypes,
	}.Build()
	File_error_user_proto = out.File
	file_error_user_proto_rawDesc = nil
	file_error_user_proto_goTypes = nil
	file_error_user_proto_depIdxs = nil
}
