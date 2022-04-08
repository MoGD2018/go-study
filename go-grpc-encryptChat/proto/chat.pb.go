// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.20.0
// source: proto/chat.proto

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

type ChatRes_MessageType int32

const (
	ChatRes_CONNECT_SUCCESS ChatRes_MessageType = 0
	ChatRes_CONNECT_FAILED  ChatRes_MessageType = 1
	ChatRes_NORMAL_MESSAGE  ChatRes_MessageType = 2
)

// Enum value maps for ChatRes_MessageType.
var (
	ChatRes_MessageType_name = map[int32]string{
		0: "CONNECT_SUCCESS",
		1: "CONNECT_FAILED",
		2: "NORMAL_MESSAGE",
	}
	ChatRes_MessageType_value = map[string]int32{
		"CONNECT_SUCCESS": 0,
		"CONNECT_FAILED":  1,
		"NORMAL_MESSAGE":  2,
	}
)

func (x ChatRes_MessageType) Enum() *ChatRes_MessageType {
	p := new(ChatRes_MessageType)
	*p = x
	return p
}

func (x ChatRes_MessageType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChatRes_MessageType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_chat_proto_enumTypes[0].Descriptor()
}

func (ChatRes_MessageType) Type() protoreflect.EnumType {
	return &file_proto_chat_proto_enumTypes[0]
}

func (x ChatRes_MessageType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChatRes_MessageType.Descriptor instead.
func (ChatRes_MessageType) EnumDescriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{1, 0}
}

// The request message containing the user's name.
type ChatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ChatReq) Reset() {
	*x = ChatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatReq) ProtoMessage() {}

func (x *ChatReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatReq.ProtoReflect.Descriptor instead.
func (*ChatReq) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{0}
}

func (x *ChatReq) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

// The response message containing the chating
type ChatRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message     string              `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	MessageType ChatRes_MessageType `protobuf:"varint,2,opt,name=message_type,json=messageType,proto3,enum=proto.ChatRes_MessageType" json:"message_type,omitempty"`
}

func (x *ChatRes) Reset() {
	*x = ChatRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_chat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChatRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatRes) ProtoMessage() {}

func (x *ChatRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_chat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatRes.ProtoReflect.Descriptor instead.
func (*ChatRes) Descriptor() ([]byte, []int) {
	return file_proto_chat_proto_rawDescGZIP(), []int{1}
}

func (x *ChatRes) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ChatRes) GetMessageType() ChatRes_MessageType {
	if x != nil {
		return x.MessageType
	}
	return ChatRes_CONNECT_SUCCESS
}

var File_proto_chat_proto protoreflect.FileDescriptor

var file_proto_chat_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x07, 0x43, 0x68, 0x61,
	0x74, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xae,
	0x01, 0x0a, 0x07, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x22, 0x4a, 0x0a, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x13, 0x0a, 0x0f, 0x43, 0x4f, 0x4e, 0x4e, 0x45, 0x43, 0x54, 0x5f, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x43, 0x4f, 0x4e, 0x4e, 0x45,
	0x43, 0x54, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x12, 0x0a, 0x0e, 0x4e,
	0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x5f, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x02, 0x32,
	0xb5, 0x01, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74, 0x12, 0x26, 0x0a, 0x04, 0x43, 0x68, 0x61, 0x74,
	0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71,
	0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73,
	0x12, 0x2a, 0x0a, 0x06, 0x43, 0x68, 0x61, 0x74, 0x49, 0x6e, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x28, 0x01, 0x12, 0x2b, 0x0a, 0x07,
	0x43, 0x68, 0x61, 0x74, 0x4f, 0x75, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x68, 0x61, 0x74, 0x52, 0x65, 0x73, 0x30, 0x01, 0x12, 0x2c, 0x0a, 0x06, 0x43, 0x68, 0x61,
	0x74, 0x49, 0x4f, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x68, 0x61, 0x74,
	0x52, 0x65, 0x73, 0x28, 0x01, 0x30, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_chat_proto_rawDescOnce sync.Once
	file_proto_chat_proto_rawDescData = file_proto_chat_proto_rawDesc
)

func file_proto_chat_proto_rawDescGZIP() []byte {
	file_proto_chat_proto_rawDescOnce.Do(func() {
		file_proto_chat_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_chat_proto_rawDescData)
	})
	return file_proto_chat_proto_rawDescData
}

var file_proto_chat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_chat_proto_goTypes = []interface{}{
	(ChatRes_MessageType)(0), // 0: proto.ChatRes.MessageType
	(*ChatReq)(nil),          // 1: proto.ChatReq
	(*ChatRes)(nil),          // 2: proto.ChatRes
}
var file_proto_chat_proto_depIdxs = []int32{
	0, // 0: proto.ChatRes.message_type:type_name -> proto.ChatRes.MessageType
	1, // 1: proto.Chat.Chat:input_type -> proto.ChatReq
	1, // 2: proto.Chat.ChatIn:input_type -> proto.ChatReq
	1, // 3: proto.Chat.ChatOut:input_type -> proto.ChatReq
	1, // 4: proto.Chat.ChatIO:input_type -> proto.ChatReq
	2, // 5: proto.Chat.Chat:output_type -> proto.ChatRes
	2, // 6: proto.Chat.ChatIn:output_type -> proto.ChatRes
	2, // 7: proto.Chat.ChatOut:output_type -> proto.ChatRes
	2, // 8: proto.Chat.ChatIO:output_type -> proto.ChatRes
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_chat_proto_init() }
func file_proto_chat_proto_init() {
	if File_proto_chat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_chat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatReq); i {
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
		file_proto_chat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChatRes); i {
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
			RawDescriptor: file_proto_chat_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_chat_proto_goTypes,
		DependencyIndexes: file_proto_chat_proto_depIdxs,
		EnumInfos:         file_proto_chat_proto_enumTypes,
		MessageInfos:      file_proto_chat_proto_msgTypes,
	}.Build()
	File_proto_chat_proto = out.File
	file_proto_chat_proto_rawDesc = nil
	file_proto_chat_proto_goTypes = nil
	file_proto_chat_proto_depIdxs = nil
}
