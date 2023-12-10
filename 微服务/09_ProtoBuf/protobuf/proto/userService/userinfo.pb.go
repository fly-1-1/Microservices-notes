// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: userinfo.proto

package userService

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

type PhoneType int32

const (
	PhoneType_MOBILE PhoneType = 0
	PhoneType_HOME   PhoneType = 1
	PhoneType_WORK   PhoneType = 2
)

// Enum value maps for PhoneType.
var (
	PhoneType_name = map[int32]string{
		0: "MOBILE",
		1: "HOME",
		2: "WORK",
	}
	PhoneType_value = map[string]int32{
		"MOBILE": 0,
		"HOME":   1,
		"WORK":   2,
	}
)

func (x PhoneType) Enum() *PhoneType {
	p := new(PhoneType)
	*p = x
	return p
}

func (x PhoneType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhoneType) Descriptor() protoreflect.EnumDescriptor {
	return file_userinfo_proto_enumTypes[0].Descriptor()
}

func (PhoneType) Type() protoreflect.EnumType {
	return &file_userinfo_proto_enumTypes[0]
}

func (x PhoneType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhoneType.Descriptor instead.
func (PhoneType) EnumDescriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{0}
}

type Userinfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string    `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Age      int32     `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	Type     PhoneType `protobuf:"varint,3,opt,name=type,proto3,enum=PhoneType" json:"type,omitempty"` //默认为 0 MOBILE
	Hobby    []string  `protobuf:"bytes,4,rep,name=hobby,proto3" json:"hobby,omitempty"`               //切片
}

func (x *Userinfo) Reset() {
	*x = Userinfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Userinfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Userinfo) ProtoMessage() {}

func (x *Userinfo) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Userinfo.ProtoReflect.Descriptor instead.
func (*Userinfo) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{0}
}

func (x *Userinfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *Userinfo) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Userinfo) GetType() PhoneType {
	if x != nil {
		return x.Type
	}
	return PhoneType_MOBILE
}

func (x *Userinfo) GetHobby() []string {
	if x != nil {
		return x.Hobby
	}
	return nil
}

var File_userinfo_proto protoreflect.FileDescriptor

var file_userinfo_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x6e, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0a, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x68, 0x6f,
	0x62, 0x62, 0x79, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x68, 0x6f, 0x62, 0x62, 0x79,
	0x2a, 0x2b, 0x0a, 0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a,
	0x06, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x4d,
	0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x02, 0x42, 0x0f, 0x5a,
	0x0d, 0x2e, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userinfo_proto_rawDescOnce sync.Once
	file_userinfo_proto_rawDescData = file_userinfo_proto_rawDesc
)

func file_userinfo_proto_rawDescGZIP() []byte {
	file_userinfo_proto_rawDescOnce.Do(func() {
		file_userinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_userinfo_proto_rawDescData)
	})
	return file_userinfo_proto_rawDescData
}

var file_userinfo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_userinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_userinfo_proto_goTypes = []interface{}{
	(PhoneType)(0),   // 0: PhoneType
	(*Userinfo)(nil), // 1: userinfo
}
var file_userinfo_proto_depIdxs = []int32{
	0, // 0: userinfo.type:type_name -> PhoneType
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_userinfo_proto_init() }
func file_userinfo_proto_init() {
	if File_userinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Userinfo); i {
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
			RawDescriptor: file_userinfo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_userinfo_proto_goTypes,
		DependencyIndexes: file_userinfo_proto_depIdxs,
		EnumInfos:         file_userinfo_proto_enumTypes,
		MessageInfos:      file_userinfo_proto_msgTypes,
	}.Build()
	File_userinfo_proto = out.File
	file_userinfo_proto_rawDesc = nil
	file_userinfo_proto_goTypes = nil
	file_userinfo_proto_depIdxs = nil
}
