// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v3.12.4
// source: user_enum.proto

package user

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

type SortOrder int32

const (
	SortOrder_DESC SortOrder = 0
	SortOrder_ASC  SortOrder = 1
)

// Enum value maps for SortOrder.
var (
	SortOrder_name = map[int32]string{
		0: "DESC",
		1: "ASC",
	}
	SortOrder_value = map[string]int32{
		"DESC": 0,
		"ASC":  1,
	}
)

func (x SortOrder) Enum() *SortOrder {
	p := new(SortOrder)
	*p = x
	return p
}

func (x SortOrder) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SortOrder) Descriptor() protoreflect.EnumDescriptor {
	return file_user_enum_proto_enumTypes[0].Descriptor()
}

func (SortOrder) Type() protoreflect.EnumType {
	return &file_user_enum_proto_enumTypes[0]
}

func (x SortOrder) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SortOrder.Descriptor instead.
func (SortOrder) EnumDescriptor() ([]byte, []int) {
	return file_user_enum_proto_rawDescGZIP(), []int{0}
}

type Status int32

const (
	Status_ACCEPTED Status = 0
	Status_DECLINED Status = 1
	Status_PENDING  Status = 2
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "ACCEPTED",
		1: "DECLINED",
		2: "PENDING",
	}
	Status_value = map[string]int32{
		"ACCEPTED": 0,
		"DECLINED": 1,
		"PENDING":  2,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_user_enum_proto_enumTypes[1].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_user_enum_proto_enumTypes[1]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_user_enum_proto_rawDescGZIP(), []int{1}
}

var File_user_enum_proto protoreflect.FileDescriptor

var file_user_enum_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x65, 0x6e, 0x75, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x2a, 0x1e, 0x0a, 0x09, 0x53, 0x6f, 0x72, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x45, 0x53, 0x43, 0x10, 0x00, 0x12, 0x07,
	0x0a, 0x03, 0x41, 0x53, 0x43, 0x10, 0x01, 0x2a, 0x31, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x0c, 0x0a, 0x08, 0x44, 0x45, 0x43, 0x4c, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_enum_proto_rawDescOnce sync.Once
	file_user_enum_proto_rawDescData = file_user_enum_proto_rawDesc
)

func file_user_enum_proto_rawDescGZIP() []byte {
	file_user_enum_proto_rawDescOnce.Do(func() {
		file_user_enum_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_enum_proto_rawDescData)
	})
	return file_user_enum_proto_rawDescData
}

var file_user_enum_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_user_enum_proto_goTypes = []any{
	(SortOrder)(0), // 0: user.SortOrder
	(Status)(0),    // 1: user.Status
}
var file_user_enum_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_enum_proto_init() }
func file_user_enum_proto_init() {
	if File_user_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_enum_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_enum_proto_goTypes,
		DependencyIndexes: file_user_enum_proto_depIdxs,
		EnumInfos:         file_user_enum_proto_enumTypes,
	}.Build()
	File_user_enum_proto = out.File
	file_user_enum_proto_rawDesc = nil
	file_user_enum_proto_goTypes = nil
	file_user_enum_proto_depIdxs = nil
}
