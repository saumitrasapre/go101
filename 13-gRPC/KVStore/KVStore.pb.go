// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0
// source: KVStore/KVStore.proto

package KVStore

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

type Key struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *Key) Reset() {
	*x = Key{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KVStore_KVStore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Key) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Key) ProtoMessage() {}

func (x *Key) ProtoReflect() protoreflect.Message {
	mi := &file_KVStore_KVStore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Key.ProtoReflect.Descriptor instead.
func (*Key) Descriptor() ([]byte, []int) {
	return file_KVStore_KVStore_proto_rawDescGZIP(), []int{0}
}

func (x *Key) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Val int32 `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KVStore_KVStore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_KVStore_KVStore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_KVStore_KVStore_proto_rawDescGZIP(), []int{1}
}

func (x *Value) GetVal() int32 {
	if x != nil {
		return x.Val
	}
	return 0
}

type KVPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Val int32  `protobuf:"varint,2,opt,name=val,proto3" json:"val,omitempty"`
}

func (x *KVPair) Reset() {
	*x = KVPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KVStore_KVStore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KVPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KVPair) ProtoMessage() {}

func (x *KVPair) ProtoReflect() protoreflect.Message {
	mi := &file_KVStore_KVStore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KVPair.ProtoReflect.Descriptor instead.
func (*KVPair) Descriptor() ([]byte, []int) {
	return file_KVStore_KVStore_proto_rawDescGZIP(), []int{2}
}

func (x *KVPair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *KVPair) GetVal() int32 {
	if x != nil {
		return x.Val
	}
	return 0
}

type Success struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Flag bool `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
}

func (x *Success) Reset() {
	*x = Success{}
	if protoimpl.UnsafeEnabled {
		mi := &file_KVStore_KVStore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Success) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Success) ProtoMessage() {}

func (x *Success) ProtoReflect() protoreflect.Message {
	mi := &file_KVStore_KVStore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Success.ProtoReflect.Descriptor instead.
func (*Success) Descriptor() ([]byte, []int) {
	return file_KVStore_KVStore_proto_rawDescGZIP(), []int{3}
}

func (x *Success) GetFlag() bool {
	if x != nil {
		return x.Flag
	}
	return false
}

var File_KVStore_KVStore_proto protoreflect.FileDescriptor

var file_KVStore_KVStore_proto_rawDesc = []byte{
	0x0a, 0x15, 0x4b, 0x56, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x4b, 0x56, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x4b, 0x56, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x22, 0x17, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x19, 0x0a, 0x05, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x76, 0x61, 0x6c, 0x22, 0x2c, 0x0a, 0x06, 0x4b, 0x56, 0x50, 0x61, 0x69, 0x72, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x76,
	0x61, 0x6c, 0x22, 0x1d, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x66, 0x6c, 0x61,
	0x67, 0x32, 0x5c, 0x0a, 0x07, 0x4b, 0x56, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x25, 0x0a, 0x03,
	0x47, 0x65, 0x74, 0x12, 0x0c, 0x2e, 0x4b, 0x56, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4b, 0x65,
	0x79, 0x1a, 0x0e, 0x2e, 0x4b, 0x56, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x0f, 0x2e, 0x4b, 0x56, 0x53,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x4b, 0x56, 0x50, 0x61, 0x69, 0x72, 0x1a, 0x10, 0x2e, 0x4b, 0x56,
	0x53, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x00, 0x42,
	0x0b, 0x5a, 0x09, 0x2e, 0x2f, 0x4b, 0x56, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_KVStore_KVStore_proto_rawDescOnce sync.Once
	file_KVStore_KVStore_proto_rawDescData = file_KVStore_KVStore_proto_rawDesc
)

func file_KVStore_KVStore_proto_rawDescGZIP() []byte {
	file_KVStore_KVStore_proto_rawDescOnce.Do(func() {
		file_KVStore_KVStore_proto_rawDescData = protoimpl.X.CompressGZIP(file_KVStore_KVStore_proto_rawDescData)
	})
	return file_KVStore_KVStore_proto_rawDescData
}

var file_KVStore_KVStore_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_KVStore_KVStore_proto_goTypes = []interface{}{
	(*Key)(nil),     // 0: KVStore.Key
	(*Value)(nil),   // 1: KVStore.Value
	(*KVPair)(nil),  // 2: KVStore.KVPair
	(*Success)(nil), // 3: KVStore.Success
}
var file_KVStore_KVStore_proto_depIdxs = []int32{
	0, // 0: KVStore.KVStore.Get:input_type -> KVStore.Key
	2, // 1: KVStore.KVStore.Put:input_type -> KVStore.KVPair
	1, // 2: KVStore.KVStore.Get:output_type -> KVStore.Value
	3, // 3: KVStore.KVStore.Put:output_type -> KVStore.Success
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_KVStore_KVStore_proto_init() }
func file_KVStore_KVStore_proto_init() {
	if File_KVStore_KVStore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_KVStore_KVStore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Key); i {
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
		file_KVStore_KVStore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_KVStore_KVStore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KVPair); i {
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
		file_KVStore_KVStore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Success); i {
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
			RawDescriptor: file_KVStore_KVStore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_KVStore_KVStore_proto_goTypes,
		DependencyIndexes: file_KVStore_KVStore_proto_depIdxs,
		MessageInfos:      file_KVStore_KVStore_proto_msgTypes,
	}.Build()
	File_KVStore_KVStore_proto = out.File
	file_KVStore_KVStore_proto_rawDesc = nil
	file_KVStore_KVStore_proto_goTypes = nil
	file_KVStore_KVStore_proto_depIdxs = nil
}
