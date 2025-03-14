// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v3.21.12
// source: api/tool/service.proto

package tool

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ToolGetAllRequest struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	Sort            *string                `protobuf:"bytes,1,opt,name=sort,proto3,oneof" json:"sort,omitempty"`
	ToolType        *int32                 `protobuf:"varint,2,opt,name=tool_type,json=toolType,proto3,oneof" json:"tool_type,omitempty"`
	DegreeOfWear    *int32                 `protobuf:"varint,3,opt,name=degree_of_wear,json=degreeOfWear,proto3,oneof" json:"degree_of_wear,omitempty"`
	SupplierId      *int32                 `protobuf:"varint,4,opt,name=supplier_id,json=supplierId,proto3,oneof" json:"supplier_id,omitempty"`
	Name            *string                `protobuf:"bytes,5,opt,name=name,proto3,oneof" json:"name,omitempty"`
	OnlyServiceable *bool                  `protobuf:"varint,6,opt,name=only_serviceable,json=onlyServiceable,proto3,oneof" json:"only_serviceable,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *ToolGetAllRequest) Reset() {
	*x = ToolGetAllRequest{}
	mi := &file_api_tool_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToolGetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolGetAllRequest) ProtoMessage() {}

func (x *ToolGetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tool_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolGetAllRequest.ProtoReflect.Descriptor instead.
func (*ToolGetAllRequest) Descriptor() ([]byte, []int) {
	return file_api_tool_service_proto_rawDescGZIP(), []int{0}
}

func (x *ToolGetAllRequest) GetSort() string {
	if x != nil && x.Sort != nil {
		return *x.Sort
	}
	return ""
}

func (x *ToolGetAllRequest) GetToolType() int32 {
	if x != nil && x.ToolType != nil {
		return *x.ToolType
	}
	return 0
}

func (x *ToolGetAllRequest) GetDegreeOfWear() int32 {
	if x != nil && x.DegreeOfWear != nil {
		return *x.DegreeOfWear
	}
	return 0
}

func (x *ToolGetAllRequest) GetSupplierId() int32 {
	if x != nil && x.SupplierId != nil {
		return *x.SupplierId
	}
	return 0
}

func (x *ToolGetAllRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ToolGetAllRequest) GetOnlyServiceable() bool {
	if x != nil && x.OnlyServiceable != nil {
		return *x.OnlyServiceable
	}
	return false
}

type ToolGetAllResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tools         []*Tool                `protobuf:"bytes,1,rep,name=tools,proto3" json:"tools,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ToolGetAllResponse) Reset() {
	*x = ToolGetAllResponse{}
	mi := &file_api_tool_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToolGetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolGetAllResponse) ProtoMessage() {}

func (x *ToolGetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_tool_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolGetAllResponse.ProtoReflect.Descriptor instead.
func (*ToolGetAllResponse) Descriptor() ([]byte, []int) {
	return file_api_tool_service_proto_rawDescGZIP(), []int{1}
}

func (x *ToolGetAllResponse) GetTools() []*Tool {
	if x != nil {
		return x.Tools
	}
	return nil
}

type ToolGetByIDRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ToolGetByIDRequest) Reset() {
	*x = ToolGetByIDRequest{}
	mi := &file_api_tool_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToolGetByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolGetByIDRequest) ProtoMessage() {}

func (x *ToolGetByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tool_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolGetByIDRequest.ProtoReflect.Descriptor instead.
func (*ToolGetByIDRequest) Descriptor() ([]byte, []int) {
	return file_api_tool_service_proto_rawDescGZIP(), []int{2}
}

func (x *ToolGetByIDRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ToolGetByIDResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tool          *Tool                  `protobuf:"bytes,1,opt,name=tool,proto3" json:"tool,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ToolGetByIDResponse) Reset() {
	*x = ToolGetByIDResponse{}
	mi := &file_api_tool_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToolGetByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolGetByIDResponse) ProtoMessage() {}

func (x *ToolGetByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_tool_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolGetByIDResponse.ProtoReflect.Descriptor instead.
func (*ToolGetByIDResponse) Descriptor() ([]byte, []int) {
	return file_api_tool_service_proto_rawDescGZIP(), []int{3}
}

func (x *ToolGetByIDResponse) GetTool() *Tool {
	if x != nil {
		return x.Tool
	}
	return nil
}

type ToolAddRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ToolType      int32                  `protobuf:"varint,1,opt,name=tool_type,json=toolType,proto3" json:"tool_type,omitempty"`
	DegreeOfWear  int32                  `protobuf:"varint,2,opt,name=degree_of_wear,json=degreeOfWear,proto3" json:"degree_of_wear,omitempty"`
	SupplierId    int32                  `protobuf:"varint,3,opt,name=supplier_id,json=supplierId,proto3" json:"supplier_id,omitempty"`
	Name          string                 `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Description   *string                `protobuf:"bytes,5,opt,name=description,proto3,oneof" json:"description,omitempty"`
	PurchaseDate  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=purchase_date,json=purchaseDate,proto3" json:"purchase_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ToolAddRequest) Reset() {
	*x = ToolAddRequest{}
	mi := &file_api_tool_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToolAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolAddRequest) ProtoMessage() {}

func (x *ToolAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tool_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolAddRequest.ProtoReflect.Descriptor instead.
func (*ToolAddRequest) Descriptor() ([]byte, []int) {
	return file_api_tool_service_proto_rawDescGZIP(), []int{4}
}

func (x *ToolAddRequest) GetToolType() int32 {
	if x != nil {
		return x.ToolType
	}
	return 0
}

func (x *ToolAddRequest) GetDegreeOfWear() int32 {
	if x != nil {
		return x.DegreeOfWear
	}
	return 0
}

func (x *ToolAddRequest) GetSupplierId() int32 {
	if x != nil {
		return x.SupplierId
	}
	return 0
}

func (x *ToolAddRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ToolAddRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *ToolAddRequest) GetPurchaseDate() *timestamppb.Timestamp {
	if x != nil {
		return x.PurchaseDate
	}
	return nil
}

type ToolAddResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tool          *Tool                  `protobuf:"bytes,1,opt,name=tool,proto3" json:"tool,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ToolAddResponse) Reset() {
	*x = ToolAddResponse{}
	mi := &file_api_tool_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToolAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolAddResponse) ProtoMessage() {}

func (x *ToolAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_tool_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolAddResponse.ProtoReflect.Descriptor instead.
func (*ToolAddResponse) Descriptor() ([]byte, []int) {
	return file_api_tool_service_proto_rawDescGZIP(), []int{5}
}

func (x *ToolAddResponse) GetTool() *Tool {
	if x != nil {
		return x.Tool
	}
	return nil
}

type ToolEditRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ToolType      *int32                 `protobuf:"varint,2,opt,name=tool_type,json=toolType,proto3,oneof" json:"tool_type,omitempty"`
	DegreeOfWear  *int32                 `protobuf:"varint,3,opt,name=degree_of_wear,json=degreeOfWear,proto3,oneof" json:"degree_of_wear,omitempty"`
	SupplierId    *int32                 `protobuf:"varint,4,opt,name=supplier_id,json=supplierId,proto3,oneof" json:"supplier_id,omitempty"`
	Name          *string                `protobuf:"bytes,5,opt,name=name,proto3,oneof" json:"name,omitempty"`
	Description   *string                `protobuf:"bytes,6,opt,name=description,proto3,oneof" json:"description,omitempty"`
	PurchaseDate  *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=purchase_date,json=purchaseDate,proto3,oneof" json:"purchase_date,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ToolEditRequest) Reset() {
	*x = ToolEditRequest{}
	mi := &file_api_tool_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToolEditRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolEditRequest) ProtoMessage() {}

func (x *ToolEditRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tool_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolEditRequest.ProtoReflect.Descriptor instead.
func (*ToolEditRequest) Descriptor() ([]byte, []int) {
	return file_api_tool_service_proto_rawDescGZIP(), []int{6}
}

func (x *ToolEditRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ToolEditRequest) GetToolType() int32 {
	if x != nil && x.ToolType != nil {
		return *x.ToolType
	}
	return 0
}

func (x *ToolEditRequest) GetDegreeOfWear() int32 {
	if x != nil && x.DegreeOfWear != nil {
		return *x.DegreeOfWear
	}
	return 0
}

func (x *ToolEditRequest) GetSupplierId() int32 {
	if x != nil && x.SupplierId != nil {
		return *x.SupplierId
	}
	return 0
}

func (x *ToolEditRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ToolEditRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *ToolEditRequest) GetPurchaseDate() *timestamppb.Timestamp {
	if x != nil {
		return x.PurchaseDate
	}
	return nil
}

type ToolEditResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Tool          *Tool                  `protobuf:"bytes,1,opt,name=tool,proto3" json:"tool,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ToolEditResponse) Reset() {
	*x = ToolEditResponse{}
	mi := &file_api_tool_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToolEditResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolEditResponse) ProtoMessage() {}

func (x *ToolEditResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_tool_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolEditResponse.ProtoReflect.Descriptor instead.
func (*ToolEditResponse) Descriptor() ([]byte, []int) {
	return file_api_tool_service_proto_rawDescGZIP(), []int{7}
}

func (x *ToolEditResponse) GetTool() *Tool {
	if x != nil {
		return x.Tool
	}
	return nil
}

type ToolDeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ToolDeleteRequest) Reset() {
	*x = ToolDeleteRequest{}
	mi := &file_api_tool_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToolDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolDeleteRequest) ProtoMessage() {}

func (x *ToolDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_tool_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolDeleteRequest.ProtoReflect.Descriptor instead.
func (*ToolDeleteRequest) Descriptor() ([]byte, []int) {
	return file_api_tool_service_proto_rawDescGZIP(), []int{8}
}

func (x *ToolDeleteRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ToolDeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ToolDeleteResponse) Reset() {
	*x = ToolDeleteResponse{}
	mi := &file_api_tool_service_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ToolDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ToolDeleteResponse) ProtoMessage() {}

func (x *ToolDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_tool_service_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ToolDeleteResponse.ProtoReflect.Descriptor instead.
func (*ToolDeleteResponse) Descriptor() ([]byte, []int) {
	return file_api_tool_service_proto_rawDescGZIP(), []int{9}
}

var File_api_tool_service_proto protoreflect.FileDescriptor

var file_api_tool_service_proto_rawDesc = string([]byte{
	0x0a, 0x16, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x6f, 0x6f, 0x6c, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc0, 0x02, 0x0a, 0x11, 0x54, 0x6f, 0x6f, 0x6c, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74,
	0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x09, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x01, 0x52, 0x08, 0x74, 0x6f, 0x6f, 0x6c, 0x54, 0x79,
	0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x5f,
	0x6f, 0x66, 0x5f, 0x77, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52,
	0x0c, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x4f, 0x66, 0x57, 0x65, 0x61, 0x72, 0x88, 0x01, 0x01,
	0x12, 0x24, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x03, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x2e, 0x0a, 0x10, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x05, 0x52, 0x0f, 0x6f, 0x6e, 0x6c,
	0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x07, 0x0a, 0x05, 0x5f, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x6f, 0x6f,
	0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x64, 0x65, 0x67, 0x72, 0x65,
	0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x77, 0x65, 0x61, 0x72, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x75,
	0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x42, 0x13, 0x0a, 0x11, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x61, 0x62, 0x6c, 0x65, 0x22, 0x36, 0x0a, 0x12, 0x54, 0x6f, 0x6f, 0x6c, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a,
	0x05, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74,
	0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x05, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x22,
	0x24, 0x0a, 0x12, 0x54, 0x6f, 0x6f, 0x6c, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x13, 0x54, 0x6f, 0x6f, 0x6c, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04,
	0x74, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x74, 0x6f, 0x6f,
	0x6c, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x74, 0x6f, 0x6f, 0x6c, 0x22, 0x80, 0x02, 0x0a,
	0x0e, 0x54, 0x6f, 0x6f, 0x6c, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x08, 0x74, 0x6f, 0x6f, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0e,
	0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x77, 0x65, 0x61, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x4f, 0x66, 0x57, 0x65,
	0x61, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x3f,
	0x0a, 0x0d, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0c, 0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x42,
	0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x31, 0x0a, 0x0f, 0x54, 0x6f, 0x6f, 0x6c, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x74, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0a, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x74, 0x6f,
	0x6f, 0x6c, 0x22, 0xf6, 0x02, 0x0a, 0x0f, 0x54, 0x6f, 0x6f, 0x6c, 0x45, 0x64, 0x69, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x09, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x08, 0x74, 0x6f, 0x6f,
	0x6c, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x29, 0x0a, 0x0e, 0x64, 0x65, 0x67, 0x72,
	0x65, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x77, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x48, 0x01, 0x52, 0x0c, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x4f, 0x66, 0x57, 0x65, 0x61, 0x72,
	0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x0a, 0x73, 0x75, 0x70, 0x70,
	0x6c, 0x69, 0x65, 0x72, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x25, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x12, 0x44, 0x0a, 0x0d, 0x70, 0x75, 0x72,
	0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x48, 0x05, 0x52, 0x0c,
	0x70, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x44, 0x61, 0x74, 0x65, 0x88, 0x01, 0x01, 0x42,
	0x0c, 0x0a, 0x0a, 0x5f, 0x74, 0x6f, 0x6f, 0x6c, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x42, 0x11, 0x0a,
	0x0f, 0x5f, 0x64, 0x65, 0x67, 0x72, 0x65, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x77, 0x65, 0x61, 0x72,
	0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x42, 0x07, 0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x10, 0x0a, 0x0e, 0x5f, 0x70, 0x75,
	0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x22, 0x32, 0x0a, 0x10, 0x54,
	0x6f, 0x6f, 0x6c, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1e, 0x0a, 0x04, 0x74, 0x6f, 0x6f, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x52, 0x04, 0x74, 0x6f, 0x6f, 0x6c, 0x22,
	0x23, 0x0a, 0x11, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x14, 0x0a, 0x12, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xa0, 0x03, 0x0a, 0x0b, 0x54,
	0x6f, 0x6f, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4e, 0x0a, 0x06, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x6f, 0x6f, 0x6c,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12,
	0x09, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x12, 0x56, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x18, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x6f, 0x6f,
	0x6c, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x45, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x14, 0x2e, 0x74, 0x6f, 0x6f, 0x6c,
	0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x15, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x09,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x12, 0x4d, 0x0a, 0x04, 0x45, 0x64, 0x69,
	0x74, 0x12, 0x15, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x45, 0x64, 0x69,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x2e,
	0x54, 0x6f, 0x6f, 0x6c, 0x45, 0x64, 0x69, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x32, 0x0e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x6f, 0x6f, 0x6c, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x53, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x17, 0x2e, 0x74, 0x6f, 0x6f, 0x6c, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x74, 0x6f,
	0x6f, 0x6c, 0x2e, 0x54, 0x6f, 0x6f, 0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x42, 0x5e, 0x5a,
	0x5c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x50, 0x65, 0x72, 0x65,
	0x76, 0x65, 0x72, 0x7a, 0x65, 0x76, 0x49, 0x76, 0x61, 0x6e, 0x2f, 0x72, 0x61, 0x7a, 0x72, 0x61,
	0x62, 0x6f, 0x74, 0x6b, 0x61, 0x2d, 0x62, 0x69, 0x7a, 0x6e, 0x65, 0x73, 0x2d, 0x70, 0x72, 0x69,
	0x6c, 0x6f, 0x6a, 0x65, 0x6e, 0x69, 0x69, 0x2d, 0x6b, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x65, 0x72,
	0x73, 0x6b, 0x61, 0x79, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x74, 0x6f, 0x6f, 0x6c, 0x3b, 0x74, 0x6f, 0x6f, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_tool_service_proto_rawDescOnce sync.Once
	file_api_tool_service_proto_rawDescData []byte
)

func file_api_tool_service_proto_rawDescGZIP() []byte {
	file_api_tool_service_proto_rawDescOnce.Do(func() {
		file_api_tool_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_tool_service_proto_rawDesc), len(file_api_tool_service_proto_rawDesc)))
	})
	return file_api_tool_service_proto_rawDescData
}

var file_api_tool_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_tool_service_proto_goTypes = []any{
	(*ToolGetAllRequest)(nil),     // 0: tool.ToolGetAllRequest
	(*ToolGetAllResponse)(nil),    // 1: tool.ToolGetAllResponse
	(*ToolGetByIDRequest)(nil),    // 2: tool.ToolGetByIDRequest
	(*ToolGetByIDResponse)(nil),   // 3: tool.ToolGetByIDResponse
	(*ToolAddRequest)(nil),        // 4: tool.ToolAddRequest
	(*ToolAddResponse)(nil),       // 5: tool.ToolAddResponse
	(*ToolEditRequest)(nil),       // 6: tool.ToolEditRequest
	(*ToolEditResponse)(nil),      // 7: tool.ToolEditResponse
	(*ToolDeleteRequest)(nil),     // 8: tool.ToolDeleteRequest
	(*ToolDeleteResponse)(nil),    // 9: tool.ToolDeleteResponse
	(*Tool)(nil),                  // 10: tool.Tool
	(*timestamppb.Timestamp)(nil), // 11: google.protobuf.Timestamp
}
var file_api_tool_service_proto_depIdxs = []int32{
	10, // 0: tool.ToolGetAllResponse.tools:type_name -> tool.Tool
	10, // 1: tool.ToolGetByIDResponse.tool:type_name -> tool.Tool
	11, // 2: tool.ToolAddRequest.purchase_date:type_name -> google.protobuf.Timestamp
	10, // 3: tool.ToolAddResponse.tool:type_name -> tool.Tool
	11, // 4: tool.ToolEditRequest.purchase_date:type_name -> google.protobuf.Timestamp
	10, // 5: tool.ToolEditResponse.tool:type_name -> tool.Tool
	0,  // 6: tool.ToolService.GetAll:input_type -> tool.ToolGetAllRequest
	2,  // 7: tool.ToolService.GetByID:input_type -> tool.ToolGetByIDRequest
	4,  // 8: tool.ToolService.Add:input_type -> tool.ToolAddRequest
	6,  // 9: tool.ToolService.Edit:input_type -> tool.ToolEditRequest
	8,  // 10: tool.ToolService.Delete:input_type -> tool.ToolDeleteRequest
	1,  // 11: tool.ToolService.GetAll:output_type -> tool.ToolGetAllResponse
	3,  // 12: tool.ToolService.GetByID:output_type -> tool.ToolGetByIDResponse
	5,  // 13: tool.ToolService.Add:output_type -> tool.ToolAddResponse
	7,  // 14: tool.ToolService.Edit:output_type -> tool.ToolEditResponse
	9,  // 15: tool.ToolService.Delete:output_type -> tool.ToolDeleteResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_tool_service_proto_init() }
func file_api_tool_service_proto_init() {
	if File_api_tool_service_proto != nil {
		return
	}
	file_api_tool_models_proto_init()
	file_api_tool_service_proto_msgTypes[0].OneofWrappers = []any{}
	file_api_tool_service_proto_msgTypes[4].OneofWrappers = []any{}
	file_api_tool_service_proto_msgTypes[6].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_tool_service_proto_rawDesc), len(file_api_tool_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_tool_service_proto_goTypes,
		DependencyIndexes: file_api_tool_service_proto_depIdxs,
		MessageInfos:      file_api_tool_service_proto_msgTypes,
	}.Build()
	File_api_tool_service_proto = out.File
	file_api_tool_service_proto_goTypes = nil
	file_api_tool_service_proto_depIdxs = nil
}
