// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        (unknown)
// source: HttpMsg.proto

package netproto

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Http_Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result  *bool   `protobuf:"varint,1,req,name=result" json:"result,omitempty"`  //结果
	Message *string `protobuf:"bytes,2,req,name=message" json:"message,omitempty"` //返回结果信息
}

func (x *Http_Result) Reset() {
	*x = Http_Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HttpMsg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http_Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http_Result) ProtoMessage() {}

func (x *Http_Result) ProtoReflect() protoreflect.Message {
	mi := &file_HttpMsg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http_Result.ProtoReflect.Descriptor instead.
func (*Http_Result) Descriptor() ([]byte, []int) {
	return file_HttpMsg_proto_rawDescGZIP(), []int{0}
}

func (x *Http_Result) GetResult() bool {
	if x != nil && x.Result != nil {
		return *x.Result
	}
	return false
}

func (x *Http_Result) GetMessage() string {
	if x != nil && x.Message != nil {
		return *x.Message
	}
	return ""
}

var File_HttpMsg_proto protoreflect.FileDescriptor

var file_HttpMsg_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x0b, 0x48, 0x74, 0x74,
	0x70, 0x5f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x08, 0x12, 0x0f, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x02, 0x28, 0x09,
}

var (
	file_HttpMsg_proto_rawDescOnce sync.Once
	file_HttpMsg_proto_rawDescData = file_HttpMsg_proto_rawDesc
)

func file_HttpMsg_proto_rawDescGZIP() []byte {
	file_HttpMsg_proto_rawDescOnce.Do(func() {
		file_HttpMsg_proto_rawDescData = protoimpl.X.CompressGZIP(file_HttpMsg_proto_rawDescData)
	})
	return file_HttpMsg_proto_rawDescData
}

var file_HttpMsg_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_HttpMsg_proto_goTypes = []interface{}{
	(*Http_Result)(nil), // 0: netproto.Http_Result
}
var file_HttpMsg_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HttpMsg_proto_init() }
func file_HttpMsg_proto_init() {
	if File_HttpMsg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HttpMsg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http_Result); i {
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
			RawDescriptor: file_HttpMsg_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HttpMsg_proto_goTypes,
		DependencyIndexes: file_HttpMsg_proto_depIdxs,
		MessageInfos:      file_HttpMsg_proto_msgTypes,
	}.Build()
	File_HttpMsg_proto = out.File
	file_HttpMsg_proto_rawDesc = nil
	file_HttpMsg_proto_goTypes = nil
	file_HttpMsg_proto_depIdxs = nil
}
