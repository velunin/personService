// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: personsApi.proto

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

type GetPersonsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetPersonsRequest) Reset() {
	*x = GetPersonsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personsApi_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonsRequest) ProtoMessage() {}

func (x *GetPersonsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personsApi_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonsRequest.ProtoReflect.Descriptor instead.
func (*GetPersonsRequest) Descriptor() ([]byte, []int) {
	return file_personsApi_proto_rawDescGZIP(), []int{0}
}

func (x *GetPersonsRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetPersonsRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPersonRequest) Reset() {
	*x = GetPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personsApi_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonRequest) ProtoMessage() {}

func (x *GetPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personsApi_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonRequest.ProtoReflect.Descriptor instead.
func (*GetPersonRequest) Descriptor() ([]byte, []int) {
	return file_personsApi_proto_rawDescGZIP(), []int{1}
}

func (x *GetPersonRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPersonsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Persons []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
}

func (x *GetPersonsResponse) Reset() {
	*x = GetPersonsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personsApi_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonsResponse) ProtoMessage() {}

func (x *GetPersonsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_personsApi_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonsResponse.ProtoReflect.Descriptor instead.
func (*GetPersonsResponse) Descriptor() ([]byte, []int) {
	return file_personsApi_proto_rawDescGZIP(), []int{2}
}

func (x *GetPersonsResponse) GetPersons() []*Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

type GetPersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Person *Person `protobuf:"bytes,1,opt,name=person,proto3" json:"person,omitempty"`
}

func (x *GetPersonResponse) Reset() {
	*x = GetPersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personsApi_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPersonResponse) ProtoMessage() {}

func (x *GetPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_personsApi_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPersonResponse.ProtoReflect.Descriptor instead.
func (*GetPersonResponse) Descriptor() ([]byte, []int) {
	return file_personsApi_proto_rawDescGZIP(), []int{3}
}

func (x *GetPersonResponse) GetPerson() *Person {
	if x != nil {
		return x.Person
	}
	return nil
}

type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	IsBlocked bool   `protobuf:"varint,4,opt,name=is_blocked,json=isBlocked,proto3" json:"is_blocked,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personsApi_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_personsApi_proto_msgTypes[4]
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
	return file_personsApi_proto_rawDescGZIP(), []int{4}
}

func (x *Person) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Person) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *Person) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *Person) GetIsBlocked() bool {
	if x != nil {
		return x.IsBlocked
	}
	return false
}

type CreatePersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *CreatePersonRequest) Reset() {
	*x = CreatePersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personsApi_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonRequest) ProtoMessage() {}

func (x *CreatePersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personsApi_proto_msgTypes[5]
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
	return file_personsApi_proto_rawDescGZIP(), []int{5}
}

func (x *CreatePersonRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *CreatePersonRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type CreatePersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreatePersonResponse) Reset() {
	*x = CreatePersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personsApi_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePersonResponse) ProtoMessage() {}

func (x *CreatePersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_personsApi_proto_msgTypes[6]
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
	return file_personsApi_proto_rawDescGZIP(), []int{6}
}

func (x *CreatePersonResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type RenamePersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FirstName string `protobuf:"bytes,2,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,3,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
}

func (x *RenamePersonRequest) Reset() {
	*x = RenamePersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personsApi_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenamePersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenamePersonRequest) ProtoMessage() {}

func (x *RenamePersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personsApi_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenamePersonRequest.ProtoReflect.Descriptor instead.
func (*RenamePersonRequest) Descriptor() ([]byte, []int) {
	return file_personsApi_proto_rawDescGZIP(), []int{7}
}

func (x *RenamePersonRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RenamePersonRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *RenamePersonRequest) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

type RenamePersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RenamePersonResponse) Reset() {
	*x = RenamePersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personsApi_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RenamePersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RenamePersonResponse) ProtoMessage() {}

func (x *RenamePersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_personsApi_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RenamePersonResponse.ProtoReflect.Descriptor instead.
func (*RenamePersonResponse) Descriptor() ([]byte, []int) {
	return file_personsApi_proto_rawDescGZIP(), []int{8}
}

type BlockPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BlockPersonRequest) Reset() {
	*x = BlockPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personsApi_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockPersonRequest) ProtoMessage() {}

func (x *BlockPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personsApi_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockPersonRequest.ProtoReflect.Descriptor instead.
func (*BlockPersonRequest) Descriptor() ([]byte, []int) {
	return file_personsApi_proto_rawDescGZIP(), []int{9}
}

func (x *BlockPersonRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type BlockPersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BlockPersonResponse) Reset() {
	*x = BlockPersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personsApi_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockPersonResponse) ProtoMessage() {}

func (x *BlockPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_personsApi_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockPersonResponse.ProtoReflect.Descriptor instead.
func (*BlockPersonResponse) Descriptor() ([]byte, []int) {
	return file_personsApi_proto_rawDescGZIP(), []int{10}
}

type UnblockPersonRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UnblockPersonRequest) Reset() {
	*x = UnblockPersonRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personsApi_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnblockPersonRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnblockPersonRequest) ProtoMessage() {}

func (x *UnblockPersonRequest) ProtoReflect() protoreflect.Message {
	mi := &file_personsApi_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnblockPersonRequest.ProtoReflect.Descriptor instead.
func (*UnblockPersonRequest) Descriptor() ([]byte, []int) {
	return file_personsApi_proto_rawDescGZIP(), []int{11}
}

func (x *UnblockPersonRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type UnblockPersonResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnblockPersonResponse) Reset() {
	*x = UnblockPersonResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_personsApi_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnblockPersonResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnblockPersonResponse) ProtoMessage() {}

func (x *UnblockPersonResponse) ProtoReflect() protoreflect.Message {
	mi := &file_personsApi_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnblockPersonResponse.ProtoReflect.Descriptor instead.
func (*UnblockPersonResponse) Descriptor() ([]byte, []int) {
	return file_personsApi_proto_rawDescGZIP(), []int{12}
}

var File_personsApi_proto protoreflect.FileDescriptor

var file_personsApi_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69, 0x22, 0x41,
	0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x07, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x52, 0x07, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x22, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a,
	0x0a, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69, 0x2e, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x52, 0x06, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x22, 0x73, 0x0a, 0x06, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x22,
	0x51, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x61, 0x0a, 0x13, 0x52, 0x65,
	0x6e, 0x61, 0x6d, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x16, 0x0a,
	0x14, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x0a, 0x12, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x55, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a, 0x15, 0x55, 0x6e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xef, 0x03, 0x0a, 0x0a, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x41,
	0x70, 0x69, 0x12, 0x4b, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73,
	0x12, 0x1d, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1e, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x48, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x1c, 0x2e, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x65, 0x72, 0x73, 0x6f,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x70, 0x65, 0x72, 0x73,
	0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0c,
	0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x1f, 0x2e, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d, 0x65,
	0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69, 0x2e, 0x52, 0x65, 0x6e, 0x61, 0x6d,
	0x65, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4e, 0x0a, 0x0b, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x1e,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69, 0x2e, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x54, 0x0a, 0x0d, 0x55, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x12, 0x20, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69, 0x2e, 0x55, 0x6e,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x41, 0x70, 0x69, 0x2e,
	0x55, 0x6e, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x22, 0x5a, 0x0c, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0xaa, 0x02, 0x11, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_personsApi_proto_rawDescOnce sync.Once
	file_personsApi_proto_rawDescData = file_personsApi_proto_rawDesc
)

func file_personsApi_proto_rawDescGZIP() []byte {
	file_personsApi_proto_rawDescOnce.Do(func() {
		file_personsApi_proto_rawDescData = protoimpl.X.CompressGZIP(file_personsApi_proto_rawDescData)
	})
	return file_personsApi_proto_rawDescData
}

var file_personsApi_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_personsApi_proto_goTypes = []interface{}{
	(*GetPersonsRequest)(nil),     // 0: personsApi.GetPersonsRequest
	(*GetPersonRequest)(nil),      // 1: personsApi.GetPersonRequest
	(*GetPersonsResponse)(nil),    // 2: personsApi.GetPersonsResponse
	(*GetPersonResponse)(nil),     // 3: personsApi.GetPersonResponse
	(*Person)(nil),                // 4: personsApi.Person
	(*CreatePersonRequest)(nil),   // 5: personsApi.CreatePersonRequest
	(*CreatePersonResponse)(nil),  // 6: personsApi.CreatePersonResponse
	(*RenamePersonRequest)(nil),   // 7: personsApi.RenamePersonRequest
	(*RenamePersonResponse)(nil),  // 8: personsApi.RenamePersonResponse
	(*BlockPersonRequest)(nil),    // 9: personsApi.BlockPersonRequest
	(*BlockPersonResponse)(nil),   // 10: personsApi.BlockPersonResponse
	(*UnblockPersonRequest)(nil),  // 11: personsApi.UnblockPersonRequest
	(*UnblockPersonResponse)(nil), // 12: personsApi.UnblockPersonResponse
}
var file_personsApi_proto_depIdxs = []int32{
	4,  // 0: personsApi.GetPersonsResponse.persons:type_name -> personsApi.Person
	4,  // 1: personsApi.GetPersonResponse.person:type_name -> personsApi.Person
	0,  // 2: personsApi.PersonsApi.GetPersons:input_type -> personsApi.GetPersonsRequest
	1,  // 3: personsApi.PersonsApi.GetPerson:input_type -> personsApi.GetPersonRequest
	5,  // 4: personsApi.PersonsApi.CreatePerson:input_type -> personsApi.CreatePersonRequest
	7,  // 5: personsApi.PersonsApi.RenamePerson:input_type -> personsApi.RenamePersonRequest
	9,  // 6: personsApi.PersonsApi.BlockPerson:input_type -> personsApi.BlockPersonRequest
	11, // 7: personsApi.PersonsApi.UnblockPerson:input_type -> personsApi.UnblockPersonRequest
	2,  // 8: personsApi.PersonsApi.GetPersons:output_type -> personsApi.GetPersonsResponse
	3,  // 9: personsApi.PersonsApi.GetPerson:output_type -> personsApi.GetPersonResponse
	6,  // 10: personsApi.PersonsApi.CreatePerson:output_type -> personsApi.CreatePersonResponse
	8,  // 11: personsApi.PersonsApi.RenamePerson:output_type -> personsApi.RenamePersonResponse
	10, // 12: personsApi.PersonsApi.BlockPerson:output_type -> personsApi.BlockPersonResponse
	12, // 13: personsApi.PersonsApi.UnblockPerson:output_type -> personsApi.UnblockPersonResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_personsApi_proto_init() }
func file_personsApi_proto_init() {
	if File_personsApi_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_personsApi_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonsRequest); i {
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
		file_personsApi_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonRequest); i {
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
		file_personsApi_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonsResponse); i {
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
		file_personsApi_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPersonResponse); i {
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
		file_personsApi_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_personsApi_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_personsApi_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_personsApi_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenamePersonRequest); i {
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
		file_personsApi_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RenamePersonResponse); i {
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
		file_personsApi_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockPersonRequest); i {
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
		file_personsApi_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockPersonResponse); i {
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
		file_personsApi_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnblockPersonRequest); i {
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
		file_personsApi_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnblockPersonResponse); i {
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
			RawDescriptor: file_personsApi_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_personsApi_proto_goTypes,
		DependencyIndexes: file_personsApi_proto_depIdxs,
		MessageInfos:      file_personsApi_proto_msgTypes,
	}.Build()
	File_personsApi_proto = out.File
	file_personsApi_proto_rawDesc = nil
	file_personsApi_proto_goTypes = nil
	file_personsApi_proto_depIdxs = nil
}
