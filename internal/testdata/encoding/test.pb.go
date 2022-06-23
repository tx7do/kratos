// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: encoding/test.proto

package encoding

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TestModel struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64             `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Hobby []string          `protobuf:"bytes,3,rep,name=hobby,proto3" json:"hobby,omitempty"`
	Attrs map[string]string `protobuf:"bytes,4,rep,name=attrs,proto3" json:"attrs,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *TestModel) Reset() {
	*x = TestModel{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encoding_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestModel) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestModel) ProtoMessage() {}

func (x *TestModel) ProtoReflect() protoreflect.Message {
	mi := &file_encoding_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestModel.ProtoReflect.Descriptor instead.
func (*TestModel) Descriptor() ([]byte, []int) {
	return file_encoding_test_proto_rawDescGZIP(), []int{0}
}

func (x *TestModel) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *TestModel) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TestModel) GetHobby() []string {
	if x != nil {
		return x.Hobby
	}
	return nil
}

func (x *TestModel) GetAttrs() map[string]string {
	if x != nil {
		return x.Attrs
	}
	return nil
}

type StructPb struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data     *structpb.Struct   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	DataList []*structpb.Struct `protobuf:"bytes,2,rep,name=data_list,json=dataList,proto3" json:"data_list,omitempty"`
}

func (x *StructPb) Reset() {
	*x = StructPb{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encoding_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StructPb) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructPb) ProtoMessage() {}

func (x *StructPb) ProtoReflect() protoreflect.Message {
	mi := &file_encoding_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructPb.ProtoReflect.Descriptor instead.
func (*StructPb) Descriptor() ([]byte, []int) {
	return file_encoding_test_proto_rawDescGZIP(), []int{1}
}

func (x *StructPb) GetData() *structpb.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *StructPb) GetDataList() []*structpb.Struct {
	if x != nil {
		return x.DataList
	}
	return nil
}

var File_encoding_test_proto protoreflect.FileDescriptor

var file_encoding_test_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x0a, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x68, 0x6f, 0x62, 0x62, 0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x62,
	0x62, 0x79, 0x12, 0x31, 0x0a, 0x05, 0x61, 0x74, 0x74, 0x72, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x41, 0x74, 0x74, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x61, 0x74, 0x74, 0x72, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x41, 0x74, 0x74, 0x72, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x6d, 0x0a, 0x08, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x50, 0x62, 0x12, 0x2b, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x34, 0x0a, 0x09, 0x64, 0x61, 0x74, 0x61,
	0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x0d,
	0x5a, 0x0b, 0x2e, 0x2e, 0x2f, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_encoding_test_proto_rawDescOnce sync.Once
	file_encoding_test_proto_rawDescData = file_encoding_test_proto_rawDesc
)

func file_encoding_test_proto_rawDescGZIP() []byte {
	file_encoding_test_proto_rawDescOnce.Do(func() {
		file_encoding_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_encoding_test_proto_rawDescData)
	})
	return file_encoding_test_proto_rawDescData
}

var file_encoding_test_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_encoding_test_proto_goTypes = []interface{}{
	(*TestModel)(nil),       // 0: test.test_model
	(*StructPb)(nil),        // 1: test.StructPb
	nil,                     // 2: test.test_model.AttrsEntry
	(*structpb.Struct)(nil), // 3: google.protobuf.Struct
}
var file_encoding_test_proto_depIdxs = []int32{
	2, // 0: test.test_model.attrs:type_name -> test.test_model.AttrsEntry
	3, // 1: test.StructPb.data:type_name -> google.protobuf.Struct
	3, // 2: test.StructPb.data_list:type_name -> google.protobuf.Struct
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_encoding_test_proto_init() }
func file_encoding_test_proto_init() {
	if File_encoding_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_encoding_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestModel); i {
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
		file_encoding_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StructPb); i {
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
			RawDescriptor: file_encoding_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_encoding_test_proto_goTypes,
		DependencyIndexes: file_encoding_test_proto_depIdxs,
		MessageInfos:      file_encoding_test_proto_msgTypes,
	}.Build()
	File_encoding_test_proto = out.File
	file_encoding_test_proto_rawDesc = nil
	file_encoding_test_proto_goTypes = nil
	file_encoding_test_proto_depIdxs = nil
}
