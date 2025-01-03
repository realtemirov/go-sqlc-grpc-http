// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: general/response.proto

package general

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	mi := &file_general_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_general_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_general_response_proto_rawDescGZIP(), []int{0}
}

type ErrorInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   string           `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	ErrorCode int32            `protobuf:"varint,2,opt,name=error_code,json=errorCode,proto3" json:"error_code,omitempty"`
	Label     *ErrorInfo_Label `protobuf:"bytes,3,opt,name=label,proto3" json:"label,omitempty"`
}

func (x *ErrorInfo) Reset() {
	*x = ErrorInfo{}
	mi := &file_general_response_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ErrorInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorInfo) ProtoMessage() {}

func (x *ErrorInfo) ProtoReflect() protoreflect.Message {
	mi := &file_general_response_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorInfo.ProtoReflect.Descriptor instead.
func (*ErrorInfo) Descriptor() ([]byte, []int) {
	return file_general_response_proto_rawDescGZIP(), []int{1}
}

func (x *ErrorInfo) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ErrorInfo) GetErrorCode() int32 {
	if x != nil {
		return x.ErrorCode
	}
	return 0
}

func (x *ErrorInfo) GetLabel() *ErrorInfo_Label {
	if x != nil {
		return x.Label
	}
	return nil
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_general_response_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_general_response_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_general_response_proto_rawDescGZIP(), []int{2}
}

func (x *Message) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ErrorInfo_Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uz string `protobuf:"bytes,1,opt,name=uz,proto3" json:"uz,omitempty"`
	Ru string `protobuf:"bytes,2,opt,name=ru,proto3" json:"ru,omitempty"`
	En string `protobuf:"bytes,3,opt,name=en,proto3" json:"en,omitempty"`
}

func (x *ErrorInfo_Label) Reset() {
	*x = ErrorInfo_Label{}
	mi := &file_general_response_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ErrorInfo_Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ErrorInfo_Label) ProtoMessage() {}

func (x *ErrorInfo_Label) ProtoReflect() protoreflect.Message {
	mi := &file_general_response_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ErrorInfo_Label.ProtoReflect.Descriptor instead.
func (*ErrorInfo_Label) Descriptor() ([]byte, []int) {
	return file_general_response_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ErrorInfo_Label) GetUz() string {
	if x != nil {
		return x.Uz
	}
	return ""
}

func (x *ErrorInfo_Label) GetRu() string {
	if x != nil {
		return x.Ru
	}
	return ""
}

func (x *ErrorInfo_Label) GetEn() string {
	if x != nil {
		return x.En
	}
	return ""
}

var File_general_response_proto protoreflect.FileDescriptor

var file_general_response_proto_rawDesc = []byte{
	0x0a, 0x16, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x6c, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xad, 0x01, 0x0a, 0x09, 0x45,
	0x72, 0x72, 0x6f, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x43, 0x6f, 0x64,
	0x65, 0x12, 0x2e, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x2e, 0x45, 0x72, 0x72, 0x6f, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x1a, 0x37, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x75, 0x7a,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x75, 0x7a, 0x12, 0x0e, 0x0a, 0x02, 0x72, 0x75,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x72, 0x75, 0x12, 0x0e, 0x0a, 0x02, 0x65, 0x6e,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x65, 0x6e, 0x22, 0x23, 0x0a, 0x07, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42,
	0x38, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61,
	0x6e, 0x64, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_general_response_proto_rawDescOnce sync.Once
	file_general_response_proto_rawDescData = file_general_response_proto_rawDesc
)

func file_general_response_proto_rawDescGZIP() []byte {
	file_general_response_proto_rawDescOnce.Do(func() {
		file_general_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_general_response_proto_rawDescData)
	})
	return file_general_response_proto_rawDescData
}

var file_general_response_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_general_response_proto_goTypes = []any{
	(*Empty)(nil),           // 0: general.Empty
	(*ErrorInfo)(nil),       // 1: general.ErrorInfo
	(*Message)(nil),         // 2: general.Message
	(*ErrorInfo_Label)(nil), // 3: general.ErrorInfo.Label
}
var file_general_response_proto_depIdxs = []int32{
	3, // 0: general.ErrorInfo.label:type_name -> general.ErrorInfo.Label
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_general_response_proto_init() }
func file_general_response_proto_init() {
	if File_general_response_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_general_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_general_response_proto_goTypes,
		DependencyIndexes: file_general_response_proto_depIdxs,
		MessageInfos:      file_general_response_proto_msgTypes,
	}.Build()
	File_general_response_proto = out.File
	file_general_response_proto_rawDesc = nil
	file_general_response_proto_goTypes = nil
	file_general_response_proto_depIdxs = nil
}