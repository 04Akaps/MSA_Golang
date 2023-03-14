// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        v3.21.12
// source: proto/sample.proto

package personpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PhoneNumberType int32

const (
	PhoneNumberType_MOBILE PhoneNumberType = 0
	PhoneNumberType_HOME   PhoneNumberType = 1
	PhoneNumberType_WORK   PhoneNumberType = 2
)

// Enum value maps for PhoneNumberType.
var (
	PhoneNumberType_name = map[int32]string{
		0: "MOBILE",
		1: "HOME",
		2: "WORK",
	}
	PhoneNumberType_value = map[string]int32{
		"MOBILE": 0,
		"HOME":   1,
		"WORK":   2,
	}
)

func (x PhoneNumberType) Enum() *PhoneNumberType {
	p := new(PhoneNumberType)
	*p = x
	return p
}

func (x PhoneNumberType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhoneNumberType) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_sample_proto_enumTypes[0].Descriptor()
}

func (PhoneNumberType) Type() protoreflect.EnumType {
	return &file_proto_sample_proto_enumTypes[0]
}

func (x PhoneNumberType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhoneNumberType.Descriptor instead.
func (PhoneNumberType) EnumDescriptor() ([]byte, []int) {
	return file_proto_sample_proto_rawDescGZIP(), []int{0}
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age          int32                  `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	PhoneNumbers []*PhoneNumber         `protobuf:"bytes,3,rep,name=phone_numbers,json=phoneNumbers,proto3" json:"phone_numbers,omitempty"`
	LastUpdated  *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=last_updated,json=lastUpdated,proto3" json:"last_updated,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_proto_msgTypes[0]
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
	return file_proto_sample_proto_rawDescGZIP(), []int{0}
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *Person) GetPhoneNumbers() []*PhoneNumber {
	if x != nil {
		return x.PhoneNumbers
	}
	return nil
}

func (x *Person) GetLastUpdated() *timestamppb.Timestamp {
	if x != nil {
		return x.LastUpdated
	}
	return nil
}

type PhoneNumber struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number string          `protobuf:"bytes,1,opt,name=number,proto3" json:"number,omitempty"`
	Type   PhoneNumberType `protobuf:"varint,2,opt,name=type,proto3,enum=PhoneNumberType" json:"type,omitempty"`
}

func (x *PhoneNumber) Reset() {
	*x = PhoneNumber{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PhoneNumber) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PhoneNumber) ProtoMessage() {}

func (x *PhoneNumber) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PhoneNumber.ProtoReflect.Descriptor instead.
func (*PhoneNumber) Descriptor() ([]byte, []int) {
	return file_proto_sample_proto_rawDescGZIP(), []int{1}
}

func (x *PhoneNumber) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

func (x *PhoneNumber) GetType() PhoneNumberType {
	if x != nil {
		return x.Type
	}
	return PhoneNumberType_MOBILE
}

type CreatePersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *CreatePersonRequest) Reset() {
	*x = CreatePersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonRequest) ProtoMessage() {}

func (x *CreatePersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersonRequest.ProtoReflect.Descriptor instead.
func (*CreatePersonRequest) Descriptor() ([]byte, []int) {
	return file_proto_sample_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePersonRequest) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type CreatePersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *CreatePersonResponse) Reset() {
	*x = CreatePersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonResponse) ProtoMessage() {}

func (x *CreatePersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePersonResponse.ProtoReflect.Descriptor instead.
func (*CreatePersonResponse) Descriptor() ([]byte, []int) {
	return file_proto_sample_proto_rawDescGZIP(), []int{3}
}

func (x *CreatePersonResponse) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type ReadPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PersonNumber string `protobuf:"bytes,1,opt,name=person_number,json=personNumber,proto3" json:"person_number,omitempty"`
}

func (x *ReadPersonRequest) Reset() {
	*x = ReadPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPersonRequest) ProtoMessage() {}

func (x *ReadPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPersonRequest.ProtoReflect.Descriptor instead.
func (*ReadPersonRequest) Descriptor() ([]byte, []int) {
	return file_proto_sample_proto_rawDescGZIP(), []int{4}
}

func (x *ReadPersonRequest) GetPersonNumber() string {
	if x != nil {
		return x.PersonNumber
	}
	return ""
}

type ReadPersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *ReadPersonResponse) Reset() {
	*x = ReadPersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPersonResponse) ProtoMessage() {}

func (x *ReadPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPersonResponse.ProtoReflect.Descriptor instead.
func (*ReadPersonResponse) Descriptor() ([]byte, []int) {
	return file_proto_sample_proto_rawDescGZIP(), []int{5}
}

func (x *ReadPersonResponse) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type ListPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListPersonRequest) Reset() {
	*x = ListPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPersonRequest) ProtoMessage() {}

func (x *ListPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPersonRequest.ProtoReflect.Descriptor instead.
func (*ListPersonRequest) Descriptor() ([]byte, []int) {
	return file_proto_sample_proto_rawDescGZIP(), []int{6}
}

type ListPersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *ListPersonResponse) Reset() {
	*x = ListPersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_sample_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListPersonResponse) ProtoMessage() {}

func (x *ListPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_sample_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListPersonResponse.ProtoReflect.Descriptor instead.
func (*ListPersonResponse) Descriptor() ([]byte, []int) {
	return file_proto_sample_proto_rawDescGZIP(), []int{7}
}

func (x *ListPersonResponse) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

var File_proto_sample_proto protoreflect.FileDescriptor

var file_proto_sample_proto_rawDesc = []byte{
	0x0a, 0x12, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa0, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x31, 0x0a, 0x0d, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x0c, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x4b, 0x0a, 0x0b, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x24, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e,
	0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x36, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x06,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x37, 0x0a,
	0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x35, 0x0a, 0x12, 0x52, 0x65, 0x61, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52,
	0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x50,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x35, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x07, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x2a, 0x31, 0x0a, 0x0f, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x4f, 0x42, 0x49, 0x4c, 0x45,
	0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x4d, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04,
	0x57, 0x4f, 0x52, 0x4b, 0x10, 0x02, 0x32, 0xba, 0x01, 0x0a, 0x0d, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x14, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x52, 0x65, 0x61, 0x64, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_sample_proto_rawDescOnce sync.Once
	file_proto_sample_proto_rawDescData = file_proto_sample_proto_rawDesc
)

func file_proto_sample_proto_rawDescGZIP() []byte {
	file_proto_sample_proto_rawDescOnce.Do(func() {
		file_proto_sample_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_sample_proto_rawDescData)
	})
	return file_proto_sample_proto_rawDescData
}

var file_proto_sample_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_sample_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_sample_proto_goTypes = []interface{}{
	(PhoneNumberType)(0),          // 0: PhoneNumberType
	(*Person)(nil),                // 1: Person
	(*PhoneNumber)(nil),           // 2: PhoneNumber
	(*CreatePersonRequest)(nil),   // 3: CreatePersonRequest
	(*CreatePersonResponse)(nil),  // 4: CreatePersonResponse
	(*ReadPersonRequest)(nil),     // 5: ReadPersonRequest
	(*ReadPersonResponse)(nil),    // 6: ReadPersonResponse
	(*ListPersonRequest)(nil),     // 7: ListPersonRequest
	(*ListPersonResponse)(nil),    // 8: ListPersonResponse
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
}
var file_proto_sample_proto_depIdxs = []int32{
	2,  // 0: Person.phone_numbers:type_name -> PhoneNumber
	9,  // 1: Person.last_updated:type_name -> google.protobuf.Timestamp
	0,  // 2: PhoneNumber.type:type_name -> PhoneNumberType
	1,  // 3: CreatePersonRequest.person:type_name -> Person
	1,  // 4: CreatePersonResponse.person:type_name -> Person
	1,  // 5: ReadPersonResponse.person:type_name -> Person
	1,  // 6: ListPersonResponse.person:type_name -> Person
	3,  // 7: PersonService.CreatePerson:input_type -> CreatePersonRequest
	5,  // 8: PersonService.ReadPerson:input_type -> ReadPersonRequest
	7,  // 9: PersonService.ListPerson:input_type -> ListPersonRequest
	4,  // 10: PersonService.CreatePerson:output_type -> CreatePersonResponse
	6,  // 11: PersonService.ReadPerson:output_type -> ReadPersonResponse
	8,  // 12: PersonService.ListPerson:output_type -> ListPersonResponse
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_proto_sample_proto_init() }
func file_proto_sample_proto_init() {
	if File_proto_sample_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_sample_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_proto_sample_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PhoneNumber); i {
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
		file_proto_sample_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersonRequest); i {
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
		file_proto_sample_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePersonResponse); i {
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
		file_proto_sample_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPersonRequest); i {
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
		file_proto_sample_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPersonResponse); i {
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
		file_proto_sample_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPersonRequest); i {
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
		file_proto_sample_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListPersonResponse); i {
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
			RawDescriptor: file_proto_sample_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_sample_proto_goTypes,
		DependencyIndexes: file_proto_sample_proto_depIdxs,
		EnumInfos:         file_proto_sample_proto_enumTypes,
		MessageInfos:      file_proto_sample_proto_msgTypes,
	}.Build()
	File_proto_sample_proto = out.File
	file_proto_sample_proto_rawDesc = nil
	file_proto_sample_proto_goTypes = nil
	file_proto_sample_proto_depIdxs = nil
}
