// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: sample/proto/sample.proto

package golang

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Email_EmailDomainType int32

const (
	Email_UNKNOWN     Email_EmailDomainType = 0 // @exclude match API specifications
	Email_GMAIL       Email_EmailDomainType = 1 // google
	Email_YAHOO_JAPAN Email_EmailDomainType = 2 // yahoo japan
	Email_OUTLOOK     Email_EmailDomainType = 3 // ms outlook
	Email_OTHER       Email_EmailDomainType = 4 // other domains
)

// Enum value maps for Email_EmailDomainType.
var (
	Email_EmailDomainType_name = map[int32]string{
		0: "UNKNOWN",
		1: "GMAIL",
		2: "YAHOO_JAPAN",
		3: "OUTLOOK",
		4: "OTHER",
	}
	Email_EmailDomainType_value = map[string]int32{
		"UNKNOWN":     0,
		"GMAIL":       1,
		"YAHOO_JAPAN": 2,
		"OUTLOOK":     3,
		"OTHER":       4,
	}
)

func (x Email_EmailDomainType) Enum() *Email_EmailDomainType {
	p := new(Email_EmailDomainType)
	*p = x
	return p
}

func (x Email_EmailDomainType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Email_EmailDomainType) Descriptor() protoreflect.EnumDescriptor {
	return file_sample_proto_sample_proto_enumTypes[0].Descriptor()
}

func (Email_EmailDomainType) Type() protoreflect.EnumType {
	return &file_sample_proto_sample_proto_enumTypes[0]
}

func (x Email_EmailDomainType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Email_EmailDomainType.Descriptor instead.
func (Email_EmailDomainType) EnumDescriptor() ([]byte, []int) {
	return file_sample_proto_sample_proto_rawDescGZIP(), []int{3, 0}
}

type SampleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // sequencial id
}

func (x *SampleRequest) Reset() {
	*x = SampleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_proto_sample_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleRequest) ProtoMessage() {}

func (x *SampleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sample_proto_sample_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleRequest.ProtoReflect.Descriptor instead.
func (*SampleRequest) Descriptor() ([]byte, []int) {
	return file_sample_proto_sample_proto_rawDescGZIP(), []int{0}
}

func (x *SampleRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SampleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person string `protobuf:"bytes,1,opt,name=Person,proto3" json:"Person,omitempty"` // person info
}

func (x *SampleResponse) Reset() {
	*x = SampleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_proto_sample_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleResponse) ProtoMessage() {}

func (x *SampleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sample_proto_sample_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleResponse.ProtoReflect.Descriptor instead.
func (*SampleResponse) Descriptor() ([]byte, []int) {
	return file_sample_proto_sample_proto_rawDescGZIP(), []int{1}
}

func (x *SampleResponse) GetPerson() string {
	if x != nil {
		return x.Person
	}
	return ""
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`      // sequencial id
	Name  string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`   // his/her name
	Age   string   `protobuf:"bytes,3,opt,name=age,proto3" json:"age,omitempty"`     // his/her age
	Email []*Email `protobuf:"bytes,4,rep,name=email,proto3" json:"email,omitempty"` // his/her email
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_proto_sample_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_sample_proto_sample_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_sample_proto_sample_proto_rawDescGZIP(), []int{2}
}

func (x *Person) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

func (x *Person) GetEmail() []*Email {
	if x != nil {
		return x.Email
	}
	return nil
}

type Email struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address string                `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`                                  // email address
	Domain  Email_EmailDomainType `protobuf:"varint,2,opt,name=domain,proto3,enum=sample.Email_EmailDomainType" json:"domain,omitempty"` // domain type
}

func (x *Email) Reset() {
	*x = Email{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sample_proto_sample_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_sample_proto_sample_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Email.ProtoReflect.Descriptor instead.
func (*Email) Descriptor() ([]byte, []int) {
	return file_sample_proto_sample_proto_rawDescGZIP(), []int{3}
}

func (x *Email) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Email) GetDomain() Email_EmailDomainType {
	if x != nil {
		return x.Domain
	}
	return Email_UNKNOWN
}

var File_sample_proto_sample_proto protoreflect.FileDescriptor

var file_sample_proto_sample_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2b, 0x0a, 0x0d, 0x53, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0x22, 0x05, 0x10, 0x90,
	0x4e, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28, 0x0a, 0x0e, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x22, 0x63, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x67,
	0x65, 0x12, 0x23, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xac, 0x01, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x35, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1d, 0x2e, 0x73, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x22, 0x52, 0x0a, 0x0f, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x4d, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b,
	0x59, 0x41, 0x48, 0x4f, 0x4f, 0x5f, 0x4a, 0x41, 0x50, 0x41, 0x4e, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x4f, 0x55, 0x54, 0x4c, 0x4f, 0x4f, 0x4b, 0x10, 0x03, 0x12, 0x09, 0x0a, 0x05, 0x4f, 0x54,
	0x48, 0x45, 0x52, 0x10, 0x04, 0x32, 0x6a, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x58, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x12, 0x15, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2f, 0x7b, 0x69, 0x64,
	0x7d, 0x42, 0x29, 0x5a, 0x27, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x69, 0x6e, 0x74, 0x61, 0x6b, 0x32, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sample_proto_sample_proto_rawDescOnce sync.Once
	file_sample_proto_sample_proto_rawDescData = file_sample_proto_sample_proto_rawDesc
)

func file_sample_proto_sample_proto_rawDescGZIP() []byte {
	file_sample_proto_sample_proto_rawDescOnce.Do(func() {
		file_sample_proto_sample_proto_rawDescData = protoimpl.X.CompressGZIP(file_sample_proto_sample_proto_rawDescData)
	})
	return file_sample_proto_sample_proto_rawDescData
}

var file_sample_proto_sample_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sample_proto_sample_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_sample_proto_sample_proto_goTypes = []interface{}{
	(Email_EmailDomainType)(0), // 0: sample.Email.EmailDomainType
	(*SampleRequest)(nil),      // 1: sample.SampleRequest
	(*SampleResponse)(nil),     // 2: sample.SampleResponse
	(*Person)(nil),             // 3: sample.Person
	(*Email)(nil),              // 4: sample.Email
}
var file_sample_proto_sample_proto_depIdxs = []int32{
	4, // 0: sample.Person.email:type_name -> sample.Email
	0, // 1: sample.Email.domain:type_name -> sample.Email.EmailDomainType
	1, // 2: sample.AddressService.Search:input_type -> sample.SampleRequest
	2, // 3: sample.AddressService.Search:output_type -> sample.SampleResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sample_proto_sample_proto_init() }
func file_sample_proto_sample_proto_init() {
	if File_sample_proto_sample_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sample_proto_sample_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleRequest); i {
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
		file_sample_proto_sample_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleResponse); i {
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
		file_sample_proto_sample_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_sample_proto_sample_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Email); i {
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
			RawDescriptor: file_sample_proto_sample_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sample_proto_sample_proto_goTypes,
		DependencyIndexes: file_sample_proto_sample_proto_depIdxs,
		EnumInfos:         file_sample_proto_sample_proto_enumTypes,
		MessageInfos:      file_sample_proto_sample_proto_msgTypes,
	}.Build()
	File_sample_proto_sample_proto = out.File
	file_sample_proto_sample_proto_rawDesc = nil
	file_sample_proto_sample_proto_goTypes = nil
	file_sample_proto_sample_proto_depIdxs = nil
}
