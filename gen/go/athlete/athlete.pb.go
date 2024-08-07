// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: athlete/athlete.proto

package athlete

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

type Athlete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CountryId string `protobuf:"bytes,3,opt,name=countryId,proto3" json:"countryId,omitempty"`
	SportType string `protobuf:"bytes,4,opt,name=sportType,proto3" json:"sportType,omitempty"`
}

func (x *Athlete) Reset() {
	*x = Athlete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_athlete_athlete_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Athlete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Athlete) ProtoMessage() {}

func (x *Athlete) ProtoReflect() protoreflect.Message {
	mi := &file_athlete_athlete_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Athlete.ProtoReflect.Descriptor instead.
func (*Athlete) Descriptor() ([]byte, []int) {
	return file_athlete_athlete_proto_rawDescGZIP(), []int{0}
}

func (x *Athlete) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Athlete) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Athlete) GetCountryId() string {
	if x != nil {
		return x.CountryId
	}
	return ""
}

func (x *Athlete) GetSportType() string {
	if x != nil {
		return x.SportType
	}
	return ""
}

type Country struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Country) Reset() {
	*x = Country{}
	if protoimpl.UnsafeEnabled {
		mi := &file_athlete_athlete_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Country) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Country) ProtoMessage() {}

func (x *Country) ProtoReflect() protoreflect.Message {
	mi := &file_athlete_athlete_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Country.ProtoReflect.Descriptor instead.
func (*Country) Descriptor() ([]byte, []int) {
	return file_athlete_athlete_proto_rawDescGZIP(), []int{1}
}

func (x *Country) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Country) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CountryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *CountryRequest) Reset() {
	*x = CountryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_athlete_athlete_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountryRequest) ProtoMessage() {}

func (x *CountryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_athlete_athlete_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountryRequest.ProtoReflect.Descriptor instead.
func (*CountryRequest) Descriptor() ([]byte, []int) {
	return file_athlete_athlete_proto_rawDescGZIP(), []int{2}
}

func (x *CountryRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CountryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CountryResponse) Reset() {
	*x = CountryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_athlete_athlete_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CountryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CountryResponse) ProtoMessage() {}

func (x *CountryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_athlete_athlete_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CountryResponse.ProtoReflect.Descriptor instead.
func (*CountryResponse) Descriptor() ([]byte, []int) {
	return file_athlete_athlete_proto_rawDescGZIP(), []int{3}
}

func (x *CountryResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AthleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CountryID string `protobuf:"bytes,2,opt,name=countryID,proto3" json:"countryID,omitempty"`
	SportType string `protobuf:"bytes,3,opt,name=sportType,proto3" json:"sportType,omitempty"`
}

func (x *AthleteRequest) Reset() {
	*x = AthleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_athlete_athlete_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AthleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AthleteRequest) ProtoMessage() {}

func (x *AthleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_athlete_athlete_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AthleteRequest.ProtoReflect.Descriptor instead.
func (*AthleteRequest) Descriptor() ([]byte, []int) {
	return file_athlete_athlete_proto_rawDescGZIP(), []int{4}
}

func (x *AthleteRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AthleteRequest) GetCountryID() string {
	if x != nil {
		return x.CountryID
	}
	return ""
}

func (x *AthleteRequest) GetSportType() string {
	if x != nil {
		return x.SportType
	}
	return ""
}

type AthleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AthleteResponse) Reset() {
	*x = AthleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_athlete_athlete_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AthleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AthleteResponse) ProtoMessage() {}

func (x *AthleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_athlete_athlete_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AthleteResponse.ProtoReflect.Descriptor instead.
func (*AthleteResponse) Descriptor() ([]byte, []int) {
	return file_athlete_athlete_proto_rawDescGZIP(), []int{5}
}

func (x *AthleteResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_athlete_athlete_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_athlete_athlete_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_athlete_athlete_proto_rawDescGZIP(), []int{6}
}

type ListAthlete struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Athlete []*Athlete `protobuf:"bytes,1,rep,name=athlete,proto3" json:"athlete,omitempty"`
}

func (x *ListAthlete) Reset() {
	*x = ListAthlete{}
	if protoimpl.UnsafeEnabled {
		mi := &file_athlete_athlete_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAthlete) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAthlete) ProtoMessage() {}

func (x *ListAthlete) ProtoReflect() protoreflect.Message {
	mi := &file_athlete_athlete_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAthlete.ProtoReflect.Descriptor instead.
func (*ListAthlete) Descriptor() ([]byte, []int) {
	return file_athlete_athlete_proto_rawDescGZIP(), []int{7}
}

func (x *ListAthlete) GetAthlete() []*Athlete {
	if x != nil {
		return x.Athlete
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_athlete_athlete_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_athlete_athlete_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_athlete_athlete_proto_rawDescGZIP(), []int{8}
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_athlete_athlete_proto protoreflect.FileDescriptor

var file_athlete_athlete_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x2f, 0x61, 0x74, 0x68, 0x6c, 0x65, 0x74,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x69, 0x0a, 0x07, 0x41, 0x74, 0x68, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x22, 0x2d, 0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x24, 0x0a, 0x0e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x60, 0x0a, 0x0e, 0x41, 0x74,
	0x68, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x44, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0x21, 0x0a, 0x0f,
	0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x31, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x07, 0x61, 0x74, 0x68, 0x6c, 0x65,
	0x74, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41, 0x74, 0x68, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x07, 0x61, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x22, 0x24, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x32, 0xcc, 0x02, 0x0a, 0x0e, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x74,
	0x68, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0f, 0x2e, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x62,
	0x79, 0x49, 0x64, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x41, 0x74, 0x68,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x08, 0x2e, 0x41,
	0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x74, 0x68,
	0x6c, 0x65, 0x74, 0x65, 0x12, 0x06, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x08, 0x2e, 0x41, 0x74,
	0x68, 0x6c, 0x65, 0x74, 0x65, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x2c, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x10, 0x2e, 0x41, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x1a, 0x09, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32,
	0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x0f, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x10, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x62, 0x79, 0x49, 0x64, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x08, 0x2e, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79,
	0x42, 0x1d, 0x5a, 0x1b, 0x61, 0x62, 0x64, 0x75, 0x61, 0x7a, 0x69, 0x6d, 0x2e, 0x61, 0x74, 0x68,
	0x6c, 0x65, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x3b, 0x61, 0x74, 0x68, 0x6c, 0x65, 0x74, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_athlete_athlete_proto_rawDescOnce sync.Once
	file_athlete_athlete_proto_rawDescData = file_athlete_athlete_proto_rawDesc
)

func file_athlete_athlete_proto_rawDescGZIP() []byte {
	file_athlete_athlete_proto_rawDescOnce.Do(func() {
		file_athlete_athlete_proto_rawDescData = protoimpl.X.CompressGZIP(file_athlete_athlete_proto_rawDescData)
	})
	return file_athlete_athlete_proto_rawDescData
}

var file_athlete_athlete_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_athlete_athlete_proto_goTypes = []any{
	(*Athlete)(nil),         // 0: Athlete
	(*Country)(nil),         // 1: Country
	(*CountryRequest)(nil),  // 2: CountryRequest
	(*CountryResponse)(nil), // 3: CountryResponse
	(*AthleteRequest)(nil),  // 4: AthleteRequest
	(*AthleteResponse)(nil), // 5: AthleteResponse
	(*Empty)(nil),           // 6: Empty
	(*ListAthlete)(nil),     // 7: ListAthlete
	(*Response)(nil),        // 8: Response
}
var file_athlete_athlete_proto_depIdxs = []int32{
	0, // 0: ListAthlete.athlete:type_name -> Athlete
	4, // 1: AthleteService.CreateAthlete:input_type -> AthleteRequest
	5, // 2: AthleteService.GetbyIdAthlete:input_type -> AthleteResponse
	6, // 3: AthleteService.GetAthlete:input_type -> Empty
	0, // 4: AthleteService.UpdateAthlete:input_type -> Athlete
	5, // 5: AthleteService.DeleteAthlete:input_type -> AthleteResponse
	2, // 6: AthleteService.CreateCountry:input_type -> CountryRequest
	3, // 7: AthleteService.GetbyIdCountry:input_type -> CountryResponse
	5, // 8: AthleteService.CreateAthlete:output_type -> AthleteResponse
	0, // 9: AthleteService.GetbyIdAthlete:output_type -> Athlete
	7, // 10: AthleteService.GetAthlete:output_type -> ListAthlete
	8, // 11: AthleteService.UpdateAthlete:output_type -> Response
	8, // 12: AthleteService.DeleteAthlete:output_type -> Response
	3, // 13: AthleteService.CreateCountry:output_type -> CountryResponse
	1, // 14: AthleteService.GetbyIdCountry:output_type -> Country
	8, // [8:15] is the sub-list for method output_type
	1, // [1:8] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_athlete_athlete_proto_init() }
func file_athlete_athlete_proto_init() {
	if File_athlete_athlete_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_athlete_athlete_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Athlete); i {
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
		file_athlete_athlete_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Country); i {
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
		file_athlete_athlete_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CountryRequest); i {
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
		file_athlete_athlete_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CountryResponse); i {
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
		file_athlete_athlete_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*AthleteRequest); i {
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
		file_athlete_athlete_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*AthleteResponse); i {
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
		file_athlete_athlete_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*Empty); i {
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
		file_athlete_athlete_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ListAthlete); i {
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
		file_athlete_athlete_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_athlete_athlete_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_athlete_athlete_proto_goTypes,
		DependencyIndexes: file_athlete_athlete_proto_depIdxs,
		MessageInfos:      file_athlete_athlete_proto_msgTypes,
	}.Build()
	File_athlete_athlete_proto = out.File
	file_athlete_athlete_proto_rawDesc = nil
	file_athlete_athlete_proto_goTypes = nil
	file_athlete_athlete_proto_depIdxs = nil
}
