// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/dialogflow/v2beta1/knowledge_base.proto

package dialogflow

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	field_mask "google.golang.org/genproto/protobuf/field_mask"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Represents knowledge base resource.
//
// Note: The `projects.agent.knowledgeBases` resource is deprecated;
// only use `projects.knowledgeBases`.
type KnowledgeBase struct {
	// The knowledge base resource name.
	// The name must be empty when creating a knowledge base.
	// Format: `projects/<Project ID>/knowledgeBases/<Knowledge Base ID>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The display name of the knowledge base. The name must be 1024
	// bytes or less; otherwise, the creation request fails.
	DisplayName          string   `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KnowledgeBase) Reset()         { *m = KnowledgeBase{} }
func (m *KnowledgeBase) String() string { return proto.CompactTextString(m) }
func (*KnowledgeBase) ProtoMessage()    {}
func (*KnowledgeBase) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdbbb049e0ce16a, []int{0}
}

func (m *KnowledgeBase) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KnowledgeBase.Unmarshal(m, b)
}
func (m *KnowledgeBase) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KnowledgeBase.Marshal(b, m, deterministic)
}
func (m *KnowledgeBase) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KnowledgeBase.Merge(m, src)
}
func (m *KnowledgeBase) XXX_Size() int {
	return xxx_messageInfo_KnowledgeBase.Size(m)
}
func (m *KnowledgeBase) XXX_DiscardUnknown() {
	xxx_messageInfo_KnowledgeBase.DiscardUnknown(m)
}

var xxx_messageInfo_KnowledgeBase proto.InternalMessageInfo

func (m *KnowledgeBase) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *KnowledgeBase) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

// Request message for [KnowledgeBases.ListKnowledgeBases][google.cloud.dialogflow.v2beta1.KnowledgeBases.ListKnowledgeBases].
type ListKnowledgeBasesRequest struct {
	// Required. The project to list of knowledge bases for.
	// Format: `projects/<Project ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Optional. The maximum number of items to return in a single page. By
	// default 10 and at most 100.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// Optional. The next_page_token value returned from a previous list request.
	PageToken            string   `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListKnowledgeBasesRequest) Reset()         { *m = ListKnowledgeBasesRequest{} }
func (m *ListKnowledgeBasesRequest) String() string { return proto.CompactTextString(m) }
func (*ListKnowledgeBasesRequest) ProtoMessage()    {}
func (*ListKnowledgeBasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdbbb049e0ce16a, []int{1}
}

func (m *ListKnowledgeBasesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListKnowledgeBasesRequest.Unmarshal(m, b)
}
func (m *ListKnowledgeBasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListKnowledgeBasesRequest.Marshal(b, m, deterministic)
}
func (m *ListKnowledgeBasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListKnowledgeBasesRequest.Merge(m, src)
}
func (m *ListKnowledgeBasesRequest) XXX_Size() int {
	return xxx_messageInfo_ListKnowledgeBasesRequest.Size(m)
}
func (m *ListKnowledgeBasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListKnowledgeBasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListKnowledgeBasesRequest proto.InternalMessageInfo

func (m *ListKnowledgeBasesRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ListKnowledgeBasesRequest) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

func (m *ListKnowledgeBasesRequest) GetPageToken() string {
	if m != nil {
		return m.PageToken
	}
	return ""
}

// Response message for [KnowledgeBases.ListKnowledgeBases][google.cloud.dialogflow.v2beta1.KnowledgeBases.ListKnowledgeBases].
type ListKnowledgeBasesResponse struct {
	// The list of knowledge bases.
	KnowledgeBases []*KnowledgeBase `protobuf:"bytes,1,rep,name=knowledge_bases,json=knowledgeBases,proto3" json:"knowledge_bases,omitempty"`
	// Token to retrieve the next page of results, or empty if there are no
	// more results in the list.
	NextPageToken        string   `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListKnowledgeBasesResponse) Reset()         { *m = ListKnowledgeBasesResponse{} }
func (m *ListKnowledgeBasesResponse) String() string { return proto.CompactTextString(m) }
func (*ListKnowledgeBasesResponse) ProtoMessage()    {}
func (*ListKnowledgeBasesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdbbb049e0ce16a, []int{2}
}

func (m *ListKnowledgeBasesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListKnowledgeBasesResponse.Unmarshal(m, b)
}
func (m *ListKnowledgeBasesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListKnowledgeBasesResponse.Marshal(b, m, deterministic)
}
func (m *ListKnowledgeBasesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListKnowledgeBasesResponse.Merge(m, src)
}
func (m *ListKnowledgeBasesResponse) XXX_Size() int {
	return xxx_messageInfo_ListKnowledgeBasesResponse.Size(m)
}
func (m *ListKnowledgeBasesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListKnowledgeBasesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListKnowledgeBasesResponse proto.InternalMessageInfo

func (m *ListKnowledgeBasesResponse) GetKnowledgeBases() []*KnowledgeBase {
	if m != nil {
		return m.KnowledgeBases
	}
	return nil
}

func (m *ListKnowledgeBasesResponse) GetNextPageToken() string {
	if m != nil {
		return m.NextPageToken
	}
	return ""
}

// Request message for [KnowledgeBase.GetDocument][].
type GetKnowledgeBaseRequest struct {
	// Required. The name of the knowledge base to retrieve.
	// Format `projects/<Project ID>/knowledgeBases/<Knowledge Base ID>`.
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetKnowledgeBaseRequest) Reset()         { *m = GetKnowledgeBaseRequest{} }
func (m *GetKnowledgeBaseRequest) String() string { return proto.CompactTextString(m) }
func (*GetKnowledgeBaseRequest) ProtoMessage()    {}
func (*GetKnowledgeBaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdbbb049e0ce16a, []int{3}
}

func (m *GetKnowledgeBaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetKnowledgeBaseRequest.Unmarshal(m, b)
}
func (m *GetKnowledgeBaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetKnowledgeBaseRequest.Marshal(b, m, deterministic)
}
func (m *GetKnowledgeBaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetKnowledgeBaseRequest.Merge(m, src)
}
func (m *GetKnowledgeBaseRequest) XXX_Size() int {
	return xxx_messageInfo_GetKnowledgeBaseRequest.Size(m)
}
func (m *GetKnowledgeBaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetKnowledgeBaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetKnowledgeBaseRequest proto.InternalMessageInfo

func (m *GetKnowledgeBaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// Request message for [KnowledgeBases.CreateKnowledgeBase][google.cloud.dialogflow.v2beta1.KnowledgeBases.CreateKnowledgeBase].
type CreateKnowledgeBaseRequest struct {
	// Required. The project to create a knowledge base for.
	// Format: `projects/<Project ID>`.
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required. The knowledge base to create.
	KnowledgeBase        *KnowledgeBase `protobuf:"bytes,2,opt,name=knowledge_base,json=knowledgeBase,proto3" json:"knowledge_base,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CreateKnowledgeBaseRequest) Reset()         { *m = CreateKnowledgeBaseRequest{} }
func (m *CreateKnowledgeBaseRequest) String() string { return proto.CompactTextString(m) }
func (*CreateKnowledgeBaseRequest) ProtoMessage()    {}
func (*CreateKnowledgeBaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdbbb049e0ce16a, []int{4}
}

func (m *CreateKnowledgeBaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateKnowledgeBaseRequest.Unmarshal(m, b)
}
func (m *CreateKnowledgeBaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateKnowledgeBaseRequest.Marshal(b, m, deterministic)
}
func (m *CreateKnowledgeBaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateKnowledgeBaseRequest.Merge(m, src)
}
func (m *CreateKnowledgeBaseRequest) XXX_Size() int {
	return xxx_messageInfo_CreateKnowledgeBaseRequest.Size(m)
}
func (m *CreateKnowledgeBaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateKnowledgeBaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateKnowledgeBaseRequest proto.InternalMessageInfo

func (m *CreateKnowledgeBaseRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *CreateKnowledgeBaseRequest) GetKnowledgeBase() *KnowledgeBase {
	if m != nil {
		return m.KnowledgeBase
	}
	return nil
}

// Request message for [KnowledgeBases.DeleteKnowledgeBase][google.cloud.dialogflow.v2beta1.KnowledgeBases.DeleteKnowledgeBase].
type DeleteKnowledgeBaseRequest struct {
	// Required. The name of the knowledge base to delete.
	// Format: `projects/<Project ID>/knowledgeBases/<Knowledge Base ID>`.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Optional. Force deletes the knowledge base. When set to true, any documents
	// in the knowledge base are also deleted.
	Force                bool     `protobuf:"varint,2,opt,name=force,proto3" json:"force,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteKnowledgeBaseRequest) Reset()         { *m = DeleteKnowledgeBaseRequest{} }
func (m *DeleteKnowledgeBaseRequest) String() string { return proto.CompactTextString(m) }
func (*DeleteKnowledgeBaseRequest) ProtoMessage()    {}
func (*DeleteKnowledgeBaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdbbb049e0ce16a, []int{5}
}

func (m *DeleteKnowledgeBaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteKnowledgeBaseRequest.Unmarshal(m, b)
}
func (m *DeleteKnowledgeBaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteKnowledgeBaseRequest.Marshal(b, m, deterministic)
}
func (m *DeleteKnowledgeBaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteKnowledgeBaseRequest.Merge(m, src)
}
func (m *DeleteKnowledgeBaseRequest) XXX_Size() int {
	return xxx_messageInfo_DeleteKnowledgeBaseRequest.Size(m)
}
func (m *DeleteKnowledgeBaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteKnowledgeBaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteKnowledgeBaseRequest proto.InternalMessageInfo

func (m *DeleteKnowledgeBaseRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DeleteKnowledgeBaseRequest) GetForce() bool {
	if m != nil {
		return m.Force
	}
	return false
}

// Request message for [KnowledgeBases.UpdateKnowledgeBase][google.cloud.dialogflow.v2beta1.KnowledgeBases.UpdateKnowledgeBase].
type UpdateKnowledgeBaseRequest struct {
	// Required. The knowledge base to update.
	KnowledgeBase *KnowledgeBase `protobuf:"bytes,1,opt,name=knowledge_base,json=knowledgeBase,proto3" json:"knowledge_base,omitempty"`
	// Optional. Not specified means `update all`.
	// Currently, only `display_name` can be updated, an InvalidArgument will be
	// returned for attempting to update other fields.
	UpdateMask           *field_mask.FieldMask `protobuf:"bytes,2,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpdateKnowledgeBaseRequest) Reset()         { *m = UpdateKnowledgeBaseRequest{} }
func (m *UpdateKnowledgeBaseRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateKnowledgeBaseRequest) ProtoMessage()    {}
func (*UpdateKnowledgeBaseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9cdbbb049e0ce16a, []int{6}
}

func (m *UpdateKnowledgeBaseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateKnowledgeBaseRequest.Unmarshal(m, b)
}
func (m *UpdateKnowledgeBaseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateKnowledgeBaseRequest.Marshal(b, m, deterministic)
}
func (m *UpdateKnowledgeBaseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateKnowledgeBaseRequest.Merge(m, src)
}
func (m *UpdateKnowledgeBaseRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateKnowledgeBaseRequest.Size(m)
}
func (m *UpdateKnowledgeBaseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateKnowledgeBaseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateKnowledgeBaseRequest proto.InternalMessageInfo

func (m *UpdateKnowledgeBaseRequest) GetKnowledgeBase() *KnowledgeBase {
	if m != nil {
		return m.KnowledgeBase
	}
	return nil
}

func (m *UpdateKnowledgeBaseRequest) GetUpdateMask() *field_mask.FieldMask {
	if m != nil {
		return m.UpdateMask
	}
	return nil
}

func init() {
	proto.RegisterType((*KnowledgeBase)(nil), "google.cloud.dialogflow.v2beta1.KnowledgeBase")
	proto.RegisterType((*ListKnowledgeBasesRequest)(nil), "google.cloud.dialogflow.v2beta1.ListKnowledgeBasesRequest")
	proto.RegisterType((*ListKnowledgeBasesResponse)(nil), "google.cloud.dialogflow.v2beta1.ListKnowledgeBasesResponse")
	proto.RegisterType((*GetKnowledgeBaseRequest)(nil), "google.cloud.dialogflow.v2beta1.GetKnowledgeBaseRequest")
	proto.RegisterType((*CreateKnowledgeBaseRequest)(nil), "google.cloud.dialogflow.v2beta1.CreateKnowledgeBaseRequest")
	proto.RegisterType((*DeleteKnowledgeBaseRequest)(nil), "google.cloud.dialogflow.v2beta1.DeleteKnowledgeBaseRequest")
	proto.RegisterType((*UpdateKnowledgeBaseRequest)(nil), "google.cloud.dialogflow.v2beta1.UpdateKnowledgeBaseRequest")
}

func init() {
	proto.RegisterFile("google/cloud/dialogflow/v2beta1/knowledge_base.proto", fileDescriptor_9cdbbb049e0ce16a)
}

var fileDescriptor_9cdbbb049e0ce16a = []byte{
	// 787 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0x4d, 0x6f, 0xd3, 0x4a,
	0x14, 0xd5, 0xa4, 0xaf, 0x55, 0x3b, 0x79, 0x6d, 0x9f, 0xa6, 0x4f, 0xfd, 0x70, 0x41, 0x2d, 0x46,
	0x42, 0x55, 0xa0, 0x1e, 0x35, 0x65, 0x81, 0x52, 0x21, 0xd1, 0x0f, 0x52, 0x21, 0x3e, 0x54, 0x05,
	0x0a, 0xa2, 0x9b, 0x68, 0x92, 0x4c, 0x5c, 0x13, 0xc7, 0x63, 0x32, 0x13, 0xd2, 0x16, 0x75, 0x53,
	0xb1, 0x41, 0x48, 0x6c, 0x58, 0x23, 0x21, 0x36, 0x48, 0x88, 0x45, 0x7f, 0x07, 0x4b, 0xe0, 0x27,
	0x20, 0xfe, 0x02, 0x2c, 0x91, 0x67, 0x9c, 0x26, 0x76, 0x6c, 0x52, 0x57, 0xec, 0xe2, 0xb9, 0x73,
	0xcf, 0x9c, 0x73, 0xcf, 0xbd, 0x93, 0x81, 0x57, 0x4d, 0xc6, 0x4c, 0x9b, 0xe2, 0xb2, 0xcd, 0x9a,
	0x15, 0x5c, 0xb1, 0x88, 0xcd, 0xcc, 0xaa, 0xcd, 0x5a, 0xf8, 0x59, 0xb6, 0x44, 0x05, 0x59, 0xc2,
	0x35, 0x87, 0xb5, 0x6c, 0x5a, 0x31, 0x69, 0xb1, 0x44, 0x38, 0x35, 0xdc, 0x06, 0x13, 0x0c, 0xcd,
	0xa9, 0x2c, 0x43, 0x66, 0x19, 0x9d, 0x2c, 0xc3, 0xcf, 0xd2, 0xce, 0xf9, 0xb0, 0xc4, 0xb5, 0x30,
	0x71, 0x1c, 0x26, 0x88, 0xb0, 0x98, 0xc3, 0x55, 0xba, 0x36, 0xeb, 0x47, 0xe5, 0x57, 0xa9, 0x59,
	0xc5, 0xb4, 0xee, 0x8a, 0x7d, 0x3f, 0x38, 0x1f, 0x0e, 0x56, 0x2d, 0x6a, 0x57, 0x8a, 0x75, 0xc2,
	0x6b, 0xfe, 0x8e, 0xa9, 0x2e, 0xf0, 0xb2, 0x6d, 0x51, 0x47, 0xa8, 0x80, 0x9e, 0x87, 0xa3, 0xb7,
	0xdb, 0x74, 0xd7, 0x08, 0xa7, 0x08, 0xc1, 0x7f, 0x1c, 0x52, 0xa7, 0xd3, 0x60, 0x1e, 0x2c, 0x8c,
	0x14, 0xe4, 0x6f, 0x74, 0x01, 0xfe, 0x5b, 0xb1, 0xb8, 0x6b, 0x93, 0xfd, 0xa2, 0x8c, 0xa5, 0x64,
	0x2c, 0xed, 0xaf, 0xdd, 0x23, 0x75, 0xaa, 0x33, 0x38, 0x73, 0xc7, 0xe2, 0x22, 0x80, 0xc5, 0x0b,
	0xf4, 0x69, 0x93, 0x72, 0x81, 0x26, 0xe1, 0x90, 0x4b, 0x1a, 0xd4, 0x11, 0x3e, 0xaa, 0xff, 0x85,
	0x66, 0xe1, 0x88, 0x4b, 0x4c, 0x5a, 0xe4, 0xd6, 0x81, 0x02, 0x1d, 0x2c, 0x0c, 0x7b, 0x0b, 0xf7,
	0xad, 0x03, 0x8a, 0xce, 0x43, 0x28, 0x83, 0x82, 0xd5, 0xa8, 0x33, 0x3d, 0x20, 0x13, 0xe5, 0xf6,
	0x07, 0xde, 0x82, 0xfe, 0x16, 0x40, 0x2d, 0xea, 0x44, 0xee, 0x32, 0x87, 0x53, 0xf4, 0x08, 0x8e,
	0x07, 0x6d, 0xe0, 0xd3, 0x60, 0x7e, 0x60, 0x21, 0x9d, 0x35, 0x8c, 0x3e, 0x46, 0x18, 0x01, 0xc4,
	0xc2, 0x58, 0x2d, 0x70, 0x00, 0xba, 0x04, 0xc7, 0x1d, 0xba, 0x27, 0x8a, 0x5d, 0xdc, 0x54, 0x39,
	0x46, 0xbd, 0xe5, 0xad, 0x13, 0x7e, 0x8b, 0x70, 0x6a, 0x93, 0x06, 0xd9, 0xb5, 0xcb, 0x11, 0x51,
	0x62, 0xfd, 0x15, 0x80, 0xda, 0x7a, 0x83, 0x12, 0x41, 0x23, 0x53, 0xe2, 0x2a, 0xb8, 0x0d, 0xc7,
	0x82, 0x32, 0x25, 0x99, 0xe4, 0x2a, 0x47, 0x03, 0x2a, 0xf5, 0x3c, 0xd4, 0x36, 0xa8, 0x4d, 0x63,
	0xc8, 0x44, 0xb5, 0xc8, 0xff, 0x70, 0xb0, 0xca, 0x1a, 0x65, 0x75, 0xfe, 0x70, 0x41, 0x7d, 0xe8,
	0xc7, 0x00, 0x6a, 0xdb, 0x6e, 0x25, 0x4e, 0x55, 0x2f, 0x7b, 0xf0, 0x17, 0xd8, 0xa3, 0x15, 0x98,
	0x6e, 0xca, 0x43, 0xe5, 0x04, 0xf8, 0x15, 0xd1, 0xda, 0x98, 0xed, 0x21, 0x31, 0xf2, 0xde, 0x90,
	0xdc, 0x25, 0xbc, 0x56, 0x80, 0x6a, 0xbb, 0xf7, 0x3b, 0xfb, 0x22, 0x0d, 0xc7, 0x82, 0x3d, 0x85,
	0x7e, 0x02, 0x88, 0x7a, 0x5b, 0x0d, 0xe5, 0xfa, 0xb2, 0x8c, 0x9d, 0x08, 0x6d, 0xe5, 0x4c, 0xb9,
	0xaa, 0xb7, 0xf5, 0xdd, 0xa3, 0xaf, 0xdf, 0xdf, 0xa4, 0x4a, 0xe8, 0xf2, 0xc9, 0x8d, 0xf3, 0x5c,
	0xb5, 0xc3, 0x75, 0xb7, 0xc1, 0x9e, 0xd0, 0xb2, 0xe0, 0x38, 0x73, 0x88, 0x83, 0x7d, 0xbb, 0xb3,
	0x8c, 0x96, 0xfe, 0xb0, 0x1d, 0x13, 0x93, 0x3a, 0x22, 0x9c, 0x84, 0x7e, 0x00, 0xf8, 0x5f, 0xb8,
	0x8b, 0xd1, 0xb5, 0xbe, 0xdc, 0x63, 0x1a, 0x5f, 0x4b, 0xe8, 0x6b, 0x94, 0x50, 0xaf, 0xd9, 0xba,
	0x79, 0x07, 0x09, 0xe3, 0xcc, 0x61, 0x50, 0x68, 0x78, 0xbb, 0x94, 0xd9, 0x9b, 0x84, 0x5e, 0xa7,
	0xe0, 0x44, 0xc4, 0xf8, 0xa1, 0xfe, 0x3e, 0xc5, 0x0f, 0x6d, 0x62, 0xb9, 0x2f, 0x81, 0xd4, 0x7b,
	0x04, 0xf4, 0x24, 0xce, 0xe6, 0x42, 0x33, 0xb4, 0xb3, 0xae, 0x27, 0x77, 0x3a, 0x0c, 0x82, 0xbe,
	0x01, 0x38, 0x11, 0x71, 0x05, 0x9c, 0xa2, 0x20, 0xf1, 0x17, 0x87, 0x36, 0xd9, 0x33, 0x83, 0x37,
	0xbd, 0x7f, 0xb1, 0xb6, 0xcf, 0x99, 0x64, 0x3e, 0x67, 0xce, 0xe0, 0xf3, 0xa7, 0x14, 0x9c, 0x88,
	0xb8, 0x90, 0x4e, 0x21, 0x2b, 0xfe, 0x1a, 0x4b, 0xec, 0xf3, 0x07, 0xe5, 0xf3, 0x3b, 0x90, 0xcd,
	0x75, 0x14, 0x84, 0x1e, 0x0d, 0x7d, 0xf5, 0xf7, 0xd8, 0x5e, 0xc8, 0xde, 0x38, 0x35, 0x5a, 0x4c,
	0x79, 0xc2, 0x98, 0xda, 0xde, 0xe7, 0xd5, 0x99, 0x2e, 0x39, 0x4a, 0x26, 0x71, 0x2d, 0x6e, 0x94,
	0x59, 0xfd, 0xcb, 0xea, 0xe3, 0x5d, 0x21, 0x5c, 0x9e, 0xc3, 0xb8, 0xd5, 0x0a, 0x07, 0x31, 0x69,
	0x8a, 0x5d, 0xf5, 0x44, 0x5a, 0x74, 0x6d, 0x22, 0xaa, 0xac, 0x51, 0xbf, 0xd2, 0x6f, 0x7b, 0xe7,
	0xa8, 0xb5, 0x63, 0x00, 0x2f, 0x96, 0x59, 0xbd, 0x5f, 0x65, 0xd7, 0x50, 0xa0, 0xb4, 0x5b, 0x5e,
	0x5f, 0x6d, 0x81, 0x9d, 0x5b, 0x7e, 0x9a, 0xc9, 0x6c, 0xe2, 0x98, 0x06, 0x6b, 0x98, 0xd8, 0xa4,
	0x8e, 0xec, 0x3a, 0xdc, 0x39, 0x38, 0xf6, 0x05, 0xb7, 0xd2, 0x59, 0xfa, 0x05, 0xc0, 0xfb, 0x54,
	0x6a, 0x23, 0xff, 0x31, 0x35, 0xb7, 0xa9, 0x30, 0xd7, 0x25, 0x95, 0x8d, 0x0e, 0x95, 0x87, 0x2a,
	0xa9, 0x34, 0x24, 0xf1, 0x97, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0x47, 0x85, 0x0d, 0xf6, 0x20,
	0x0a, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// KnowledgeBasesClient is the client API for KnowledgeBases service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type KnowledgeBasesClient interface {
	// Returns the list of all knowledge bases of the specified agent.
	//
	// Note: The `projects.agent.knowledgeBases` resource is deprecated;
	// only use `projects.knowledgeBases`.
	ListKnowledgeBases(ctx context.Context, in *ListKnowledgeBasesRequest, opts ...grpc.CallOption) (*ListKnowledgeBasesResponse, error)
	// Retrieves the specified knowledge base.
	//
	// Note: The `projects.agent.knowledgeBases` resource is deprecated;
	// only use `projects.knowledgeBases`.
	GetKnowledgeBase(ctx context.Context, in *GetKnowledgeBaseRequest, opts ...grpc.CallOption) (*KnowledgeBase, error)
	// Creates a knowledge base.
	//
	// Note: The `projects.agent.knowledgeBases` resource is deprecated;
	// only use `projects.knowledgeBases`.
	CreateKnowledgeBase(ctx context.Context, in *CreateKnowledgeBaseRequest, opts ...grpc.CallOption) (*KnowledgeBase, error)
	// Deletes the specified knowledge base.
	//
	// Note: The `projects.agent.knowledgeBases` resource is deprecated;
	// only use `projects.knowledgeBases`.
	DeleteKnowledgeBase(ctx context.Context, in *DeleteKnowledgeBaseRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// Updates the specified knowledge base.
	//
	// Note: The `projects.agent.knowledgeBases` resource is deprecated;
	// only use `projects.knowledgeBases`.
	UpdateKnowledgeBase(ctx context.Context, in *UpdateKnowledgeBaseRequest, opts ...grpc.CallOption) (*KnowledgeBase, error)
}

type knowledgeBasesClient struct {
	cc *grpc.ClientConn
}

func NewKnowledgeBasesClient(cc *grpc.ClientConn) KnowledgeBasesClient {
	return &knowledgeBasesClient{cc}
}

func (c *knowledgeBasesClient) ListKnowledgeBases(ctx context.Context, in *ListKnowledgeBasesRequest, opts ...grpc.CallOption) (*ListKnowledgeBasesResponse, error) {
	out := new(ListKnowledgeBasesResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.KnowledgeBases/ListKnowledgeBases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knowledgeBasesClient) GetKnowledgeBase(ctx context.Context, in *GetKnowledgeBaseRequest, opts ...grpc.CallOption) (*KnowledgeBase, error) {
	out := new(KnowledgeBase)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.KnowledgeBases/GetKnowledgeBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knowledgeBasesClient) CreateKnowledgeBase(ctx context.Context, in *CreateKnowledgeBaseRequest, opts ...grpc.CallOption) (*KnowledgeBase, error) {
	out := new(KnowledgeBase)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.KnowledgeBases/CreateKnowledgeBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knowledgeBasesClient) DeleteKnowledgeBase(ctx context.Context, in *DeleteKnowledgeBaseRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.KnowledgeBases/DeleteKnowledgeBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *knowledgeBasesClient) UpdateKnowledgeBase(ctx context.Context, in *UpdateKnowledgeBaseRequest, opts ...grpc.CallOption) (*KnowledgeBase, error) {
	out := new(KnowledgeBase)
	err := c.cc.Invoke(ctx, "/google.cloud.dialogflow.v2beta1.KnowledgeBases/UpdateKnowledgeBase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KnowledgeBasesServer is the server API for KnowledgeBases service.
type KnowledgeBasesServer interface {
	// Returns the list of all knowledge bases of the specified agent.
	//
	// Note: The `projects.agent.knowledgeBases` resource is deprecated;
	// only use `projects.knowledgeBases`.
	ListKnowledgeBases(context.Context, *ListKnowledgeBasesRequest) (*ListKnowledgeBasesResponse, error)
	// Retrieves the specified knowledge base.
	//
	// Note: The `projects.agent.knowledgeBases` resource is deprecated;
	// only use `projects.knowledgeBases`.
	GetKnowledgeBase(context.Context, *GetKnowledgeBaseRequest) (*KnowledgeBase, error)
	// Creates a knowledge base.
	//
	// Note: The `projects.agent.knowledgeBases` resource is deprecated;
	// only use `projects.knowledgeBases`.
	CreateKnowledgeBase(context.Context, *CreateKnowledgeBaseRequest) (*KnowledgeBase, error)
	// Deletes the specified knowledge base.
	//
	// Note: The `projects.agent.knowledgeBases` resource is deprecated;
	// only use `projects.knowledgeBases`.
	DeleteKnowledgeBase(context.Context, *DeleteKnowledgeBaseRequest) (*empty.Empty, error)
	// Updates the specified knowledge base.
	//
	// Note: The `projects.agent.knowledgeBases` resource is deprecated;
	// only use `projects.knowledgeBases`.
	UpdateKnowledgeBase(context.Context, *UpdateKnowledgeBaseRequest) (*KnowledgeBase, error)
}

// UnimplementedKnowledgeBasesServer can be embedded to have forward compatible implementations.
type UnimplementedKnowledgeBasesServer struct {
}

func (*UnimplementedKnowledgeBasesServer) ListKnowledgeBases(ctx context.Context, req *ListKnowledgeBasesRequest) (*ListKnowledgeBasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListKnowledgeBases not implemented")
}
func (*UnimplementedKnowledgeBasesServer) GetKnowledgeBase(ctx context.Context, req *GetKnowledgeBaseRequest) (*KnowledgeBase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetKnowledgeBase not implemented")
}
func (*UnimplementedKnowledgeBasesServer) CreateKnowledgeBase(ctx context.Context, req *CreateKnowledgeBaseRequest) (*KnowledgeBase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateKnowledgeBase not implemented")
}
func (*UnimplementedKnowledgeBasesServer) DeleteKnowledgeBase(ctx context.Context, req *DeleteKnowledgeBaseRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteKnowledgeBase not implemented")
}
func (*UnimplementedKnowledgeBasesServer) UpdateKnowledgeBase(ctx context.Context, req *UpdateKnowledgeBaseRequest) (*KnowledgeBase, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKnowledgeBase not implemented")
}

func RegisterKnowledgeBasesServer(s *grpc.Server, srv KnowledgeBasesServer) {
	s.RegisterService(&_KnowledgeBases_serviceDesc, srv)
}

func _KnowledgeBases_ListKnowledgeBases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListKnowledgeBasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeBasesServer).ListKnowledgeBases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.KnowledgeBases/ListKnowledgeBases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeBasesServer).ListKnowledgeBases(ctx, req.(*ListKnowledgeBasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnowledgeBases_GetKnowledgeBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetKnowledgeBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeBasesServer).GetKnowledgeBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.KnowledgeBases/GetKnowledgeBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeBasesServer).GetKnowledgeBase(ctx, req.(*GetKnowledgeBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnowledgeBases_CreateKnowledgeBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateKnowledgeBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeBasesServer).CreateKnowledgeBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.KnowledgeBases/CreateKnowledgeBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeBasesServer).CreateKnowledgeBase(ctx, req.(*CreateKnowledgeBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnowledgeBases_DeleteKnowledgeBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteKnowledgeBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeBasesServer).DeleteKnowledgeBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.KnowledgeBases/DeleteKnowledgeBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeBasesServer).DeleteKnowledgeBase(ctx, req.(*DeleteKnowledgeBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KnowledgeBases_UpdateKnowledgeBase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateKnowledgeBaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KnowledgeBasesServer).UpdateKnowledgeBase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.dialogflow.v2beta1.KnowledgeBases/UpdateKnowledgeBase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KnowledgeBasesServer).UpdateKnowledgeBase(ctx, req.(*UpdateKnowledgeBaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KnowledgeBases_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.dialogflow.v2beta1.KnowledgeBases",
	HandlerType: (*KnowledgeBasesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListKnowledgeBases",
			Handler:    _KnowledgeBases_ListKnowledgeBases_Handler,
		},
		{
			MethodName: "GetKnowledgeBase",
			Handler:    _KnowledgeBases_GetKnowledgeBase_Handler,
		},
		{
			MethodName: "CreateKnowledgeBase",
			Handler:    _KnowledgeBases_CreateKnowledgeBase_Handler,
		},
		{
			MethodName: "DeleteKnowledgeBase",
			Handler:    _KnowledgeBases_DeleteKnowledgeBase_Handler,
		},
		{
			MethodName: "UpdateKnowledgeBase",
			Handler:    _KnowledgeBases_UpdateKnowledgeBase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/dialogflow/v2beta1/knowledge_base.proto",
}
