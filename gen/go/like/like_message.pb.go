// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.3
// source: like_message.proto

package like

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

type ExistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *ExistsRequest) Reset() {
	*x = ExistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistsRequest) ProtoMessage() {}

func (x *ExistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistsRequest.ProtoReflect.Descriptor instead.
func (*ExistsRequest) Descriptor() ([]byte, []int) {
	return file_like_message_proto_rawDescGZIP(), []int{0}
}

func (x *ExistsRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type ExistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,3,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *ExistsResponse) Reset() {
	*x = ExistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExistsResponse) ProtoMessage() {}

func (x *ExistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExistsResponse.ProtoReflect.Descriptor instead.
func (*ExistsResponse) Descriptor() ([]byte, []int) {
	return file_like_message_proto_rawDescGZIP(), []int{1}
}

func (x *ExistsResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

type LikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *LikeRequest) Reset() {
	*x = LikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeRequest) ProtoMessage() {}

func (x *LikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeRequest.ProtoReflect.Descriptor instead.
func (*LikeRequest) Descriptor() ([]byte, []int) {
	return file_like_message_proto_rawDescGZIP(), []int{2}
}

func (x *LikeRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type LikeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *LikeResponse) Reset() {
	*x = LikeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeResponse) ProtoMessage() {}

func (x *LikeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeResponse.ProtoReflect.Descriptor instead.
func (*LikeResponse) Descriptor() ([]byte, []int) {
	return file_like_message_proto_rawDescGZIP(), []int{3}
}

func (x *LikeResponse) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

type CountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Count int32 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *CountResponse) Reset() {
	*x = CountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountResponse) ProtoMessage() {}

func (x *CountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_like_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountResponse.ProtoReflect.Descriptor instead.
func (*CountResponse) Descriptor() ([]byte, []int) {
	return file_like_message_proto_rawDescGZIP(), []int{4}
}

func (x *CountResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type CountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId string `protobuf:"bytes,1,opt,name=post_id,json=postId,proto3" json:"post_id,omitempty"`
}

func (x *CountRequest) Reset() {
	*x = CountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_like_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountRequest) ProtoMessage() {}

func (x *CountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_like_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountRequest.ProtoReflect.Descriptor instead.
func (*CountRequest) Descriptor() ([]byte, []int) {
	return file_like_message_proto_rawDescGZIP(), []int{5}
}

func (x *CountRequest) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

var File_like_message_proto protoreflect.FileDescriptor

var file_like_message_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x69, 0x6b, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x22, 0x28, 0x0a, 0x0d, 0x45, 0x78,
	0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70,
	0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f,
	0x73, 0x74, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x0e, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x22, 0x26,
	0x0a, 0x0b, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22, 0x27, 0x0a, 0x0c, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x22,
	0x25, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x27, 0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x42,
	0x0c, 0x5a, 0x0a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x6b, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_like_message_proto_rawDescOnce sync.Once
	file_like_message_proto_rawDescData = file_like_message_proto_rawDesc
)

func file_like_message_proto_rawDescGZIP() []byte {
	file_like_message_proto_rawDescOnce.Do(func() {
		file_like_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_like_message_proto_rawDescData)
	})
	return file_like_message_proto_rawDescData
}

var file_like_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_like_message_proto_goTypes = []any{
	(*ExistsRequest)(nil),  // 0: like.ExistsRequest
	(*ExistsResponse)(nil), // 1: like.ExistsResponse
	(*LikeRequest)(nil),    // 2: like.LikeRequest
	(*LikeResponse)(nil),   // 3: like.LikeResponse
	(*CountResponse)(nil),  // 4: like.CountResponse
	(*CountRequest)(nil),   // 5: like.CountRequest
}
var file_like_message_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_like_message_proto_init() }
func file_like_message_proto_init() {
	if File_like_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_like_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ExistsRequest); i {
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
		file_like_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ExistsResponse); i {
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
		file_like_message_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*LikeRequest); i {
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
		file_like_message_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*LikeResponse); i {
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
		file_like_message_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*CountResponse); i {
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
		file_like_message_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*CountRequest); i {
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
			RawDescriptor: file_like_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_like_message_proto_goTypes,
		DependencyIndexes: file_like_message_proto_depIdxs,
		MessageInfos:      file_like_message_proto_msgTypes,
	}.Build()
	File_like_message_proto = out.File
	file_like_message_proto_rawDesc = nil
	file_like_message_proto_goTypes = nil
	file_like_message_proto_depIdxs = nil
}