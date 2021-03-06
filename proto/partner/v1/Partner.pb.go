// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/partner/v1/Partner.proto

package prixa_partner_v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "github.com/mwitkow/go-proto-validators"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
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

type PartnerResponseData struct {
	Id                   string               `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Status               string               `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	AppIds               []string             `protobuf:"bytes,4,rep,name=appIds,proto3" json:"appIds,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *PartnerResponseData) Reset()         { *m = PartnerResponseData{} }
func (m *PartnerResponseData) String() string { return proto.CompactTextString(m) }
func (*PartnerResponseData) ProtoMessage()    {}
func (*PartnerResponseData) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{0}
}

func (m *PartnerResponseData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PartnerResponseData.Unmarshal(m, b)
}
func (m *PartnerResponseData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PartnerResponseData.Marshal(b, m, deterministic)
}
func (m *PartnerResponseData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PartnerResponseData.Merge(m, src)
}
func (m *PartnerResponseData) XXX_Size() int {
	return xxx_messageInfo_PartnerResponseData.Size(m)
}
func (m *PartnerResponseData) XXX_DiscardUnknown() {
	xxx_messageInfo_PartnerResponseData.DiscardUnknown(m)
}

var xxx_messageInfo_PartnerResponseData proto.InternalMessageInfo

func (m *PartnerResponseData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PartnerResponseData) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *PartnerResponseData) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *PartnerResponseData) GetAppIds() []string {
	if m != nil {
		return m.AppIds
	}
	return nil
}

func (m *PartnerResponseData) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *PartnerResponseData) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

type CreatePartnerRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreatePartnerRequest) Reset()         { *m = CreatePartnerRequest{} }
func (m *CreatePartnerRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePartnerRequest) ProtoMessage()    {}
func (*CreatePartnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{1}
}

func (m *CreatePartnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePartnerRequest.Unmarshal(m, b)
}
func (m *CreatePartnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePartnerRequest.Marshal(b, m, deterministic)
}
func (m *CreatePartnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePartnerRequest.Merge(m, src)
}
func (m *CreatePartnerRequest) XXX_Size() int {
	return xxx_messageInfo_CreatePartnerRequest.Size(m)
}
func (m *CreatePartnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePartnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePartnerRequest proto.InternalMessageInfo

func (m *CreatePartnerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CreatePartnerResponse struct {
	Data                 *PartnerResponseData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *CreatePartnerResponse) Reset()         { *m = CreatePartnerResponse{} }
func (m *CreatePartnerResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePartnerResponse) ProtoMessage()    {}
func (*CreatePartnerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{2}
}

func (m *CreatePartnerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreatePartnerResponse.Unmarshal(m, b)
}
func (m *CreatePartnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreatePartnerResponse.Marshal(b, m, deterministic)
}
func (m *CreatePartnerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreatePartnerResponse.Merge(m, src)
}
func (m *CreatePartnerResponse) XXX_Size() int {
	return xxx_messageInfo_CreatePartnerResponse.Size(m)
}
func (m *CreatePartnerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreatePartnerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreatePartnerResponse proto.InternalMessageInfo

func (m *CreatePartnerResponse) GetData() *PartnerResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type DeletePartnerRequest struct {
	PartnerId            string   `protobuf:"bytes,1,opt,name=partnerId,proto3" json:"partnerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletePartnerRequest) Reset()         { *m = DeletePartnerRequest{} }
func (m *DeletePartnerRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePartnerRequest) ProtoMessage()    {}
func (*DeletePartnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{3}
}

func (m *DeletePartnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePartnerRequest.Unmarshal(m, b)
}
func (m *DeletePartnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePartnerRequest.Marshal(b, m, deterministic)
}
func (m *DeletePartnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePartnerRequest.Merge(m, src)
}
func (m *DeletePartnerRequest) XXX_Size() int {
	return xxx_messageInfo_DeletePartnerRequest.Size(m)
}
func (m *DeletePartnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePartnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePartnerRequest proto.InternalMessageInfo

func (m *DeletePartnerRequest) GetPartnerId() string {
	if m != nil {
		return m.PartnerId
	}
	return ""
}

type DeletePartnerResponse struct {
	Data                 *PartnerResponseData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *DeletePartnerResponse) Reset()         { *m = DeletePartnerResponse{} }
func (m *DeletePartnerResponse) String() string { return proto.CompactTextString(m) }
func (*DeletePartnerResponse) ProtoMessage()    {}
func (*DeletePartnerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{4}
}

func (m *DeletePartnerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletePartnerResponse.Unmarshal(m, b)
}
func (m *DeletePartnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletePartnerResponse.Marshal(b, m, deterministic)
}
func (m *DeletePartnerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletePartnerResponse.Merge(m, src)
}
func (m *DeletePartnerResponse) XXX_Size() int {
	return xxx_messageInfo_DeletePartnerResponse.Size(m)
}
func (m *DeletePartnerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletePartnerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeletePartnerResponse proto.InternalMessageInfo

func (m *DeletePartnerResponse) GetData() *PartnerResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type UpdatePartnerRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	PartnerId            string   `protobuf:"bytes,2,opt,name=partnerId,proto3" json:"partnerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdatePartnerRequest) Reset()         { *m = UpdatePartnerRequest{} }
func (m *UpdatePartnerRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePartnerRequest) ProtoMessage()    {}
func (*UpdatePartnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{5}
}

func (m *UpdatePartnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePartnerRequest.Unmarshal(m, b)
}
func (m *UpdatePartnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePartnerRequest.Marshal(b, m, deterministic)
}
func (m *UpdatePartnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePartnerRequest.Merge(m, src)
}
func (m *UpdatePartnerRequest) XXX_Size() int {
	return xxx_messageInfo_UpdatePartnerRequest.Size(m)
}
func (m *UpdatePartnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePartnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePartnerRequest proto.InternalMessageInfo

func (m *UpdatePartnerRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdatePartnerRequest) GetPartnerId() string {
	if m != nil {
		return m.PartnerId
	}
	return ""
}

type UpdatePartnerResponse struct {
	Data                 *PartnerResponseData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *UpdatePartnerResponse) Reset()         { *m = UpdatePartnerResponse{} }
func (m *UpdatePartnerResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePartnerResponse) ProtoMessage()    {}
func (*UpdatePartnerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{6}
}

func (m *UpdatePartnerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdatePartnerResponse.Unmarshal(m, b)
}
func (m *UpdatePartnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdatePartnerResponse.Marshal(b, m, deterministic)
}
func (m *UpdatePartnerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdatePartnerResponse.Merge(m, src)
}
func (m *UpdatePartnerResponse) XXX_Size() int {
	return xxx_messageInfo_UpdatePartnerResponse.Size(m)
}
func (m *UpdatePartnerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdatePartnerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdatePartnerResponse proto.InternalMessageInfo

func (m *UpdatePartnerResponse) GetData() *PartnerResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type GetPartnerRequest struct {
	PartnerId            string   `protobuf:"bytes,1,opt,name=partnerId,proto3" json:"partnerId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetPartnerRequest) Reset()         { *m = GetPartnerRequest{} }
func (m *GetPartnerRequest) String() string { return proto.CompactTextString(m) }
func (*GetPartnerRequest) ProtoMessage()    {}
func (*GetPartnerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{7}
}

func (m *GetPartnerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPartnerRequest.Unmarshal(m, b)
}
func (m *GetPartnerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPartnerRequest.Marshal(b, m, deterministic)
}
func (m *GetPartnerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPartnerRequest.Merge(m, src)
}
func (m *GetPartnerRequest) XXX_Size() int {
	return xxx_messageInfo_GetPartnerRequest.Size(m)
}
func (m *GetPartnerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPartnerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetPartnerRequest proto.InternalMessageInfo

func (m *GetPartnerRequest) GetPartnerId() string {
	if m != nil {
		return m.PartnerId
	}
	return ""
}

type GetPartnerResponse struct {
	Data                 *PartnerResponseData `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *GetPartnerResponse) Reset()         { *m = GetPartnerResponse{} }
func (m *GetPartnerResponse) String() string { return proto.CompactTextString(m) }
func (*GetPartnerResponse) ProtoMessage()    {}
func (*GetPartnerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{8}
}

func (m *GetPartnerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetPartnerResponse.Unmarshal(m, b)
}
func (m *GetPartnerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetPartnerResponse.Marshal(b, m, deterministic)
}
func (m *GetPartnerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetPartnerResponse.Merge(m, src)
}
func (m *GetPartnerResponse) XXX_Size() int {
	return xxx_messageInfo_GetPartnerResponse.Size(m)
}
func (m *GetPartnerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetPartnerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetPartnerResponse proto.InternalMessageInfo

func (m *GetPartnerResponse) GetData() *PartnerResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

type PageProperty struct {
	PageNo               int32    `protobuf:"varint,1,opt,name=pageNo,proto3" json:"pageNo,omitempty"`
	TotalPages           int32    `protobuf:"varint,2,opt,name=totalPages,proto3" json:"totalPages,omitempty"`
	TotalRecords         int32    `protobuf:"varint,3,opt,name=totalRecords,proto3" json:"totalRecords,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PageProperty) Reset()         { *m = PageProperty{} }
func (m *PageProperty) String() string { return proto.CompactTextString(m) }
func (*PageProperty) ProtoMessage()    {}
func (*PageProperty) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{9}
}

func (m *PageProperty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PageProperty.Unmarshal(m, b)
}
func (m *PageProperty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PageProperty.Marshal(b, m, deterministic)
}
func (m *PageProperty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PageProperty.Merge(m, src)
}
func (m *PageProperty) XXX_Size() int {
	return xxx_messageInfo_PageProperty.Size(m)
}
func (m *PageProperty) XXX_DiscardUnknown() {
	xxx_messageInfo_PageProperty.DiscardUnknown(m)
}

var xxx_messageInfo_PageProperty proto.InternalMessageInfo

func (m *PageProperty) GetPageNo() int32 {
	if m != nil {
		return m.PageNo
	}
	return 0
}

func (m *PageProperty) GetTotalPages() int32 {
	if m != nil {
		return m.TotalPages
	}
	return 0
}

func (m *PageProperty) GetTotalRecords() int32 {
	if m != nil {
		return m.TotalRecords
	}
	return 0
}

type ListPartnersRequest struct {
	Page                 int32    `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListPartnersRequest) Reset()         { *m = ListPartnersRequest{} }
func (m *ListPartnersRequest) String() string { return proto.CompactTextString(m) }
func (*ListPartnersRequest) ProtoMessage()    {}
func (*ListPartnersRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{10}
}

func (m *ListPartnersRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPartnersRequest.Unmarshal(m, b)
}
func (m *ListPartnersRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPartnersRequest.Marshal(b, m, deterministic)
}
func (m *ListPartnersRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPartnersRequest.Merge(m, src)
}
func (m *ListPartnersRequest) XXX_Size() int {
	return xxx_messageInfo_ListPartnersRequest.Size(m)
}
func (m *ListPartnersRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPartnersRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListPartnersRequest proto.InternalMessageInfo

func (m *ListPartnersRequest) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

type ListPartnersResponse struct {
	Data                 []*PartnerResponseData `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Page                 *PageProperty          `protobuf:"bytes,2,opt,name=page,proto3" json:"page,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *ListPartnersResponse) Reset()         { *m = ListPartnersResponse{} }
func (m *ListPartnersResponse) String() string { return proto.CompactTextString(m) }
func (*ListPartnersResponse) ProtoMessage()    {}
func (*ListPartnersResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6f7b43f209e465d, []int{11}
}

func (m *ListPartnersResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPartnersResponse.Unmarshal(m, b)
}
func (m *ListPartnersResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPartnersResponse.Marshal(b, m, deterministic)
}
func (m *ListPartnersResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPartnersResponse.Merge(m, src)
}
func (m *ListPartnersResponse) XXX_Size() int {
	return xxx_messageInfo_ListPartnersResponse.Size(m)
}
func (m *ListPartnersResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPartnersResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListPartnersResponse proto.InternalMessageInfo

func (m *ListPartnersResponse) GetData() []*PartnerResponseData {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *ListPartnersResponse) GetPage() *PageProperty {
	if m != nil {
		return m.Page
	}
	return nil
}

func init() {
	proto.RegisterType((*PartnerResponseData)(nil), "prixa.partner.v1.PartnerResponseData")
	proto.RegisterType((*CreatePartnerRequest)(nil), "prixa.partner.v1.CreatePartnerRequest")
	proto.RegisterType((*CreatePartnerResponse)(nil), "prixa.partner.v1.CreatePartnerResponse")
	proto.RegisterType((*DeletePartnerRequest)(nil), "prixa.partner.v1.DeletePartnerRequest")
	proto.RegisterType((*DeletePartnerResponse)(nil), "prixa.partner.v1.DeletePartnerResponse")
	proto.RegisterType((*UpdatePartnerRequest)(nil), "prixa.partner.v1.UpdatePartnerRequest")
	proto.RegisterType((*UpdatePartnerResponse)(nil), "prixa.partner.v1.UpdatePartnerResponse")
	proto.RegisterType((*GetPartnerRequest)(nil), "prixa.partner.v1.GetPartnerRequest")
	proto.RegisterType((*GetPartnerResponse)(nil), "prixa.partner.v1.GetPartnerResponse")
	proto.RegisterType((*PageProperty)(nil), "prixa.partner.v1.PageProperty")
	proto.RegisterType((*ListPartnersRequest)(nil), "prixa.partner.v1.ListPartnersRequest")
	proto.RegisterType((*ListPartnersResponse)(nil), "prixa.partner.v1.ListPartnersResponse")
}

func init() {
	proto.RegisterFile("proto/partner/v1/Partner.proto", fileDescriptor_c6f7b43f209e465d)
}

var fileDescriptor_c6f7b43f209e465d = []byte{
	// 900 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x56, 0x4f, 0x6f, 0x1b, 0x45,
	0x14, 0xef, 0x6e, 0x6c, 0xab, 0x7e, 0x8d, 0xd3, 0x74, 0xe2, 0x22, 0x77, 0x69, 0x9d, 0xd5, 0x06,
	0x92, 0x10, 0xe1, 0x1d, 0x62, 0x24, 0x54, 0x02, 0x52, 0x71, 0x9a, 0x08, 0x59, 0x8d, 0x8a, 0xb5,
	0x2d, 0x52, 0x04, 0xa7, 0xf1, 0x7a, 0x58, 0x4f, 0xb1, 0x77, 0x96, 0x9d, 0xb1, 0xd3, 0xf0, 0xe7,
	0x82, 0x84, 0xc4, 0x11, 0x95, 0x03, 0x27, 0xae, 0x7c, 0x07, 0x24, 0x0e, 0x7c, 0x07, 0x3e, 0x40,
	0xa5, 0x88, 0x0f, 0xc1, 0x81, 0x03, 0xda, 0xd9, 0xb1, 0x93, 0xb5, 0x57, 0x49, 0x68, 0x4e, 0x9e,
	0x99, 0xf7, 0xde, 0x6f, 0x7e, 0xef, 0xb7, 0xef, 0xb7, 0x5e, 0xa8, 0x47, 0x31, 0x97, 0x1c, 0x47,
	0x24, 0x96, 0x21, 0x8d, 0xf1, 0x78, 0x1b, 0x77, 0xd2, 0xa5, 0xab, 0x02, 0x68, 0x39, 0x8a, 0xd9,
	0x73, 0xe2, 0xea, 0xb8, 0x3b, 0xde, 0xb6, 0x56, 0x03, 0xce, 0x83, 0x01, 0xc5, 0x2a, 0xde, 0x1d,
	0x7d, 0x81, 0x25, 0x1b, 0x52, 0x21, 0xc9, 0x30, 0x4a, 0x4b, 0xac, 0xbb, 0x3a, 0x81, 0x44, 0x0c,
	0x93, 0x30, 0xe4, 0x92, 0x48, 0xc6, 0x43, 0xa1, 0xa3, 0x6f, 0xab, 0x1f, 0xbf, 0x11, 0xd0, 0xb0,
	0x21, 0x8e, 0x48, 0x10, 0xd0, 0x18, 0xf3, 0x48, 0x65, 0xe4, 0x64, 0xbf, 0x17, 0x30, 0xd9, 0x1f,
	0x75, 0x5d, 0x9f, 0x0f, 0xf1, 0xf0, 0x88, 0xc9, 0x2f, 0xf9, 0x11, 0x0e, 0x78, 0x43, 0x05, 0x1b,
	0x63, 0x32, 0x60, 0x3d, 0x22, 0x79, 0x2c, 0xf0, 0x74, 0x99, 0xd6, 0x39, 0xff, 0x18, 0xb0, 0xa2,
	0x1b, 0xf1, 0xa8, 0x88, 0x78, 0x28, 0xe8, 0x1e, 0x91, 0x04, 0xdd, 0x01, 0x93, 0xf5, 0x6a, 0x86,
	0x6d, 0x6c, 0x96, 0x77, 0xcb, 0x27, 0x2f, 0x57, 0x8b, 0x87, 0xc6, 0x4f, 0x46, 0xc1, 0x33, 0x59,
	0x0f, 0x59, 0x50, 0x08, 0xc9, 0x90, 0xd6, 0x4c, 0x15, 0x2c, 0x9d, 0xbc, 0x5c, 0x35, 0x0f, 0x0d,
	0x4f, 0x9d, 0xa1, 0x3a, 0x94, 0x84, 0x24, 0x72, 0x24, 0x6a, 0x0b, 0x99, 0xa8, 0x3e, 0x45, 0xaf,
	0x41, 0x89, 0x44, 0x51, 0xbb, 0x27, 0x6a, 0x05, 0x7b, 0x61, 0xb3, 0xec, 0xe9, 0x1d, 0xba, 0x0f,
	0x65, 0x3f, 0xa6, 0x44, 0xd2, 0x5e, 0x4b, 0xd6, 0x8a, 0xb6, 0xb1, 0x79, 0xa3, 0x69, 0xb9, 0xa9,
	0x3c, 0xee, 0x44, 0x3f, 0xf7, 0xe9, 0x44, 0x3f, 0xef, 0x34, 0x39, 0xa9, 0x1c, 0x45, 0x3d, 0x5d,
	0x59, 0xba, 0xb8, 0x72, 0x9a, 0xec, 0x34, 0xa1, 0xfa, 0x50, 0xc1, 0x4c, 0xfb, 0xff, 0x6a, 0x44,
	0x85, 0x9c, 0xf6, 0x67, 0xcc, 0xf7, 0xe7, 0x78, 0x70, 0x7b, 0xa6, 0x26, 0xd5, 0x0c, 0xbd, 0x0f,
	0x85, 0x1e, 0x91, 0x44, 0x15, 0xdd, 0x68, 0xbe, 0xe9, 0xce, 0x4e, 0x83, 0x9b, 0x23, 0xb2, 0xa7,
	0x4a, 0x9c, 0x07, 0x50, 0xdd, 0xa3, 0x03, 0x3a, 0xc7, 0x63, 0x03, 0xca, 0xba, 0xbe, 0x9d, 0xf3,
	0x24, 0x4e, 0x63, 0x09, 0xa9, 0x19, 0x80, 0xab, 0x93, 0xfa, 0x1c, 0xaa, 0x9f, 0x2a, 0xa5, 0x2e,
	0x2f, 0x4e, 0x96, 0xb0, 0x79, 0x3e, 0xe1, 0x19, 0xf0, 0xab, 0x13, 0xfe, 0x10, 0x6e, 0x7d, 0x4c,
	0xe5, 0xab, 0x4a, 0xf8, 0x09, 0xa0, 0xb3, 0xd5, 0x57, 0xa7, 0xf3, 0x0c, 0x16, 0x3b, 0x24, 0xa0,
	0x9d, 0x98, 0x47, 0x34, 0x96, 0xc7, 0xc9, 0xe0, 0x47, 0x24, 0xa0, 0x8f, 0xb9, 0x02, 0x2b, 0x7a,
	0x7a, 0x87, 0xea, 0x00, 0x92, 0x4b, 0x32, 0x48, 0x92, 0x85, 0x12, 0xad, 0xe8, 0x9d, 0x39, 0x41,
	0x0e, 0x2c, 0xaa, 0x9d, 0x47, 0x7d, 0x1e, 0xf7, 0x52, 0x5b, 0x15, 0xbd, 0xcc, 0x99, 0xf3, 0x16,
	0xac, 0x1c, 0x30, 0x31, 0x61, 0x2f, 0x26, 0xcd, 0x23, 0x28, 0x24, 0x97, 0xe8, 0x0b, 0xd5, 0xda,
	0xf9, 0xc1, 0x80, 0x6a, 0x36, 0x77, 0xae, 0xd5, 0x85, 0xff, 0xd9, 0x2a, 0x6a, 0xea, 0x7b, 0x4c,
	0xa5, 0x52, 0x3d, 0xaf, 0xf4, 0x54, 0x88, 0x94, 0x47, 0xf3, 0xdf, 0x22, 0x2c, 0x69, 0xc4, 0x27,
	0x34, 0x1e, 0x33, 0x9f, 0xa2, 0x5f, 0x0d, 0xa8, 0x64, 0xbc, 0x85, 0xd6, 0xe7, 0xa1, 0xf2, 0x0c,
	0x6b, 0x6d, 0x5c, 0x98, 0x97, 0x72, 0x76, 0x3e, 0x78, 0xd1, 0xaa, 0x77, 0xef, 0xc2, 0x12, 0x40,
	0x2b, 0x62, 0x8f, 0xe8, 0x71, 0x6b, 0x24, 0xfb, 0xe8, 0x1a, 0x00, 0x94, 0x76, 0x29, 0x89, 0x69,
	0x8c, 0xae, 0x7d, 0xff, 0xd7, 0xdf, 0x3f, 0x9b, 0x55, 0xe7, 0xa6, 0x7a, 0x27, 0x8f, 0xb7, 0x27,
	0xaf, 0xfb, 0x1d, 0x63, 0x0b, 0xfd, 0x66, 0x40, 0x25, 0x63, 0xb3, 0x3c, 0x7e, 0x79, 0x46, 0xce,
	0xe3, 0x97, 0xeb, 0x57, 0x67, 0xef, 0x72, 0xfc, 0xee, 0x6d, 0xbd, 0x3e, 0xc3, 0x0f, 0x7f, 0x33,
	0x1d, 0xe5, 0xef, 0xd0, 0x8f, 0x06, 0x54, 0x32, 0xf6, 0xca, 0x23, 0x9a, 0x67, 0xee, 0x3c, 0xa2,
	0xb9, 0x3e, 0x75, 0xd6, 0x15, 0x0f, 0xdb, 0x3a, 0x8f, 0x47, 0xa2, 0xd9, 0xb7, 0x00, 0xa7, 0xb6,
	0x42, 0x6b, 0xf3, 0xf0, 0x73, 0x96, 0xb5, 0xde, 0x38, 0x3f, 0x49, 0x13, 0x58, 0x4b, 0x85, 0x40,
	0xe7, 0x0a, 0xf1, 0x8b, 0x01, 0x8b, 0x67, 0x87, 0x1d, 0xe5, 0x8c, 0x75, 0x8e, 0x71, 0xac, 0xf5,
	0x8b, 0xd2, 0x34, 0x89, 0xfb, 0x97, 0x7b, 0x5c, 0xb7, 0xd0, 0xec, 0x38, 0xed, 0xfe, 0x69, 0xbe,
	0x68, 0xfd, 0x61, 0xa2, 0xdf, 0x0d, 0xb8, 0xd3, 0x49, 0xae, 0xb2, 0x35, 0xb6, 0xad, 0xcd, 0x60,
	0xb7, 0x3a, 0x6d, 0x67, 0x1f, 0x20, 0x0d, 0x3e, 0xa5, 0x7e, 0x1f, 0x6d, 0xf4, 0xa5, 0x8c, 0xc4,
	0x0e, 0xc6, 0x67, 0xfe, 0xe8, 0x15, 0xcd, 0x06, 0x61, 0x7a, 0x91, 0x7e, 0x09, 0x2c, 0x49, 0xea,
	0xf7, 0x3f, 0x4a, 0x3b, 0xf0, 0xf9, 0x70, 0xab, 0x03, 0x95, 0x14, 0xe6, 0x80, 0xf9, 0x34, 0xb1,
	0xfa, 0x83, 0x4b, 0x22, 0xe1, 0xee, 0x80, 0x77, 0xf1, 0x90, 0x08, 0x49, 0x63, 0x7c, 0xd0, 0x7e,
	0xb8, 0xff, 0xf8, 0xc9, 0xbe, 0x2b, 0x9f, 0xcb, 0xe6, 0xc2, 0xb6, 0xfb, 0x8e, 0xb5, 0x4c, 0xc3,
	0x80, 0x85, 0xb4, 0xa9, 0xb5, 0x22, 0x6c, 0xcb, 0x34, 0xcc, 0xe6, 0x32, 0x89, 0xa2, 0x01, 0xf3,
	0xd5, 0xa7, 0x09, 0x7e, 0x26, 0x78, 0xb8, 0x33, 0x77, 0xf2, 0xd9, 0x1e, 0xac, 0x65, 0x14, 0xbb,
	0x7d, 0xdd, 0xb4, 0x6e, 0x1e, 0x36, 0x14, 0xc9, 0x46, 0xab, 0xd3, 0x6e, 0x3c, 0xa2, 0xc7, 0xb6,
	0x09, 0xf7, 0xa6, 0x32, 0xae, 0x5c, 0x37, 0xad, 0x4a, 0x92, 0xca, 0x63, 0xf6, 0xb5, 0x02, 0xb2,
	0xcd, 0x6e, 0x49, 0x71, 0x7c, 0xf7, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x98, 0xea, 0xb5, 0x4c,
	0x98, 0x09, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// PartnerServiceClient is the client API for PartnerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PartnerServiceClient interface {
	CreatePartner(ctx context.Context, in *CreatePartnerRequest, opts ...grpc.CallOption) (*CreatePartnerResponse, error)
	DeletePartner(ctx context.Context, in *DeletePartnerRequest, opts ...grpc.CallOption) (*DeletePartnerResponse, error)
	UpdatePartner(ctx context.Context, in *UpdatePartnerRequest, opts ...grpc.CallOption) (*UpdatePartnerResponse, error)
	GetPartner(ctx context.Context, in *GetPartnerRequest, opts ...grpc.CallOption) (*GetPartnerResponse, error)
	ListPartners(ctx context.Context, in *ListPartnersRequest, opts ...grpc.CallOption) (*ListPartnersResponse, error)
}

type partnerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPartnerServiceClient(cc grpc.ClientConnInterface) PartnerServiceClient {
	return &partnerServiceClient{cc}
}

func (c *partnerServiceClient) CreatePartner(ctx context.Context, in *CreatePartnerRequest, opts ...grpc.CallOption) (*CreatePartnerResponse, error) {
	out := new(CreatePartnerResponse)
	err := c.cc.Invoke(ctx, "/prixa.partner.v1.PartnerService/CreatePartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) DeletePartner(ctx context.Context, in *DeletePartnerRequest, opts ...grpc.CallOption) (*DeletePartnerResponse, error) {
	out := new(DeletePartnerResponse)
	err := c.cc.Invoke(ctx, "/prixa.partner.v1.PartnerService/DeletePartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) UpdatePartner(ctx context.Context, in *UpdatePartnerRequest, opts ...grpc.CallOption) (*UpdatePartnerResponse, error) {
	out := new(UpdatePartnerResponse)
	err := c.cc.Invoke(ctx, "/prixa.partner.v1.PartnerService/UpdatePartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) GetPartner(ctx context.Context, in *GetPartnerRequest, opts ...grpc.CallOption) (*GetPartnerResponse, error) {
	out := new(GetPartnerResponse)
	err := c.cc.Invoke(ctx, "/prixa.partner.v1.PartnerService/GetPartner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *partnerServiceClient) ListPartners(ctx context.Context, in *ListPartnersRequest, opts ...grpc.CallOption) (*ListPartnersResponse, error) {
	out := new(ListPartnersResponse)
	err := c.cc.Invoke(ctx, "/prixa.partner.v1.PartnerService/ListPartners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PartnerServiceServer is the server API for PartnerService service.
type PartnerServiceServer interface {
	CreatePartner(context.Context, *CreatePartnerRequest) (*CreatePartnerResponse, error)
	DeletePartner(context.Context, *DeletePartnerRequest) (*DeletePartnerResponse, error)
	UpdatePartner(context.Context, *UpdatePartnerRequest) (*UpdatePartnerResponse, error)
	GetPartner(context.Context, *GetPartnerRequest) (*GetPartnerResponse, error)
	ListPartners(context.Context, *ListPartnersRequest) (*ListPartnersResponse, error)
}

// UnimplementedPartnerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPartnerServiceServer struct {
}

func (*UnimplementedPartnerServiceServer) CreatePartner(ctx context.Context, req *CreatePartnerRequest) (*CreatePartnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePartner not implemented")
}
func (*UnimplementedPartnerServiceServer) DeletePartner(ctx context.Context, req *DeletePartnerRequest) (*DeletePartnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePartner not implemented")
}
func (*UnimplementedPartnerServiceServer) UpdatePartner(ctx context.Context, req *UpdatePartnerRequest) (*UpdatePartnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePartner not implemented")
}
func (*UnimplementedPartnerServiceServer) GetPartner(ctx context.Context, req *GetPartnerRequest) (*GetPartnerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPartner not implemented")
}
func (*UnimplementedPartnerServiceServer) ListPartners(ctx context.Context, req *ListPartnersRequest) (*ListPartnersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPartners not implemented")
}

func RegisterPartnerServiceServer(s *grpc.Server, srv PartnerServiceServer) {
	s.RegisterService(&_PartnerService_serviceDesc, srv)
}

func _PartnerService_CreatePartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePartnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).CreatePartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prixa.partner.v1.PartnerService/CreatePartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).CreatePartner(ctx, req.(*CreatePartnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_DeletePartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePartnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).DeletePartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prixa.partner.v1.PartnerService/DeletePartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).DeletePartner(ctx, req.(*DeletePartnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_UpdatePartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePartnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).UpdatePartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prixa.partner.v1.PartnerService/UpdatePartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).UpdatePartner(ctx, req.(*UpdatePartnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_GetPartner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPartnerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).GetPartner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prixa.partner.v1.PartnerService/GetPartner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).GetPartner(ctx, req.(*GetPartnerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PartnerService_ListPartners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPartnersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PartnerServiceServer).ListPartners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/prixa.partner.v1.PartnerService/ListPartners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PartnerServiceServer).ListPartners(ctx, req.(*ListPartnersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PartnerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "prixa.partner.v1.PartnerService",
	HandlerType: (*PartnerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePartner",
			Handler:    _PartnerService_CreatePartner_Handler,
		},
		{
			MethodName: "DeletePartner",
			Handler:    _PartnerService_DeletePartner_Handler,
		},
		{
			MethodName: "UpdatePartner",
			Handler:    _PartnerService_UpdatePartner_Handler,
		},
		{
			MethodName: "GetPartner",
			Handler:    _PartnerService_GetPartner_Handler,
		},
		{
			MethodName: "ListPartners",
			Handler:    _PartnerService_ListPartners_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/partner/v1/Partner.proto",
}
