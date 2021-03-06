// Code generated by protoc-gen-go. DO NOT EDIT.
// source: Models.proto

package services

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

//商品模型
type ProdModel struct {
	ProdId               int32    `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	ProdName             string   `protobuf:"bytes,2,opt,name=prod_name,json=prodName,proto3" json:"prod_name,omitempty"`
	ProdPrice            float32  `protobuf:"fixed32,3,opt,name=prod_price,json=prodPrice,proto3" json:"prod_price,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdModel) Reset()         { *m = ProdModel{} }
func (m *ProdModel) String() string { return proto.CompactTextString(m) }
func (*ProdModel) ProtoMessage()    {}
func (*ProdModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b05f67b8e9f86a, []int{0}
}

func (m *ProdModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdModel.Unmarshal(m, b)
}
func (m *ProdModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdModel.Marshal(b, m, deterministic)
}
func (m *ProdModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdModel.Merge(m, src)
}
func (m *ProdModel) XXX_Size() int {
	return xxx_messageInfo_ProdModel.Size(m)
}
func (m *ProdModel) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdModel.DiscardUnknown(m)
}

var xxx_messageInfo_ProdModel proto.InternalMessageInfo

func (m *ProdModel) GetProdId() int32 {
	if m != nil {
		return m.ProdId
	}
	return 0
}

func (m *ProdModel) GetProdName() string {
	if m != nil {
		return m.ProdName
	}
	return ""
}

func (m *ProdModel) GetProdPrice() float32 {
	if m != nil {
		return m.ProdPrice
	}
	return 0
}

//主订单模型
type OrderMain struct {
	OrderId              int32                `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	OrderNo              string               `protobuf:"bytes,2,opt,name=order_no,json=orderNo,proto3" json:"order_no,omitempty"`
	UserId               int32                `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	OrderMoney           float32              `protobuf:"fixed32,4,opt,name=order_money,json=orderMoney,proto3" json:"order_money,omitempty"`
	OrderTime            *timestamp.Timestamp `protobuf:"bytes,5,opt,name=order_time,json=orderTime,proto3" json:"order_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OrderMain) Reset()         { *m = OrderMain{} }
func (m *OrderMain) String() string { return proto.CompactTextString(m) }
func (*OrderMain) ProtoMessage()    {}
func (*OrderMain) Descriptor() ([]byte, []int) {
	return fileDescriptor_96b05f67b8e9f86a, []int{1}
}

func (m *OrderMain) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderMain.Unmarshal(m, b)
}
func (m *OrderMain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderMain.Marshal(b, m, deterministic)
}
func (m *OrderMain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderMain.Merge(m, src)
}
func (m *OrderMain) XXX_Size() int {
	return xxx_messageInfo_OrderMain.Size(m)
}
func (m *OrderMain) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderMain.DiscardUnknown(m)
}

var xxx_messageInfo_OrderMain proto.InternalMessageInfo

func (m *OrderMain) GetOrderId() int32 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *OrderMain) GetOrderNo() string {
	if m != nil {
		return m.OrderNo
	}
	return ""
}

func (m *OrderMain) GetUserId() int32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *OrderMain) GetOrderMoney() float32 {
	if m != nil {
		return m.OrderMoney
	}
	return 0
}

func (m *OrderMain) GetOrderTime() *timestamp.Timestamp {
	if m != nil {
		return m.OrderTime
	}
	return nil
}

func init() {
	proto.RegisterType((*ProdModel)(nil), "services.ProdModel")
	proto.RegisterType((*OrderMain)(nil), "services.OrderMain")
}

func init() { proto.RegisterFile("Models.proto", fileDescriptor_96b05f67b8e9f86a) }

var fileDescriptor_96b05f67b8e9f86a = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x90, 0xc1, 0x6e, 0xc2, 0x30,
	0x10, 0x44, 0x65, 0x28, 0x10, 0x2f, 0x3d, 0xf9, 0x52, 0x97, 0xaa, 0x22, 0xe2, 0x94, 0x53, 0x90,
	0xda, 0x53, 0x3f, 0x21, 0x87, 0x50, 0x14, 0xf5, 0x4e, 0x03, 0xde, 0x22, 0x4b, 0x38, 0x1b, 0xd9,
	0xa1, 0x52, 0x7f, 0xac, 0xdf, 0x57, 0x79, 0x5d, 0x94, 0xe3, 0x9b, 0xd9, 0xd1, 0xd8, 0x03, 0xf7,
	0x35, 0x19, 0xbc, 0x84, 0xb2, 0xf7, 0x34, 0x90, 0xca, 0x02, 0xfa, 0x6f, 0x7b, 0xc2, 0xb0, 0x5a,
	0x9f, 0x89, 0xce, 0x17, 0xdc, 0xb2, 0x7e, 0xbc, 0x7e, 0x6d, 0x07, 0xeb, 0x30, 0x0c, 0xad, 0xeb,
	0xd3, 0xe9, 0xe6, 0x13, 0xe4, 0xde, 0x93, 0xe1, 0xb8, 0x7a, 0x80, 0x45, 0xef, 0xc9, 0x1c, 0xac,
	0xd1, 0x22, 0x17, 0xc5, 0xac, 0x99, 0x47, 0xac, 0x8c, 0x7a, 0x02, 0xc9, 0x46, 0xd7, 0x3a, 0xd4,
	0x93, 0x5c, 0x14, 0xb2, 0xc9, 0xa2, 0xb0, 0x6b, 0x1d, 0xaa, 0x67, 0x00, 0x36, 0x7b, 0x6f, 0x4f,
	0xa8, 0xa7, 0xb9, 0x28, 0x26, 0x0d, 0x9f, 0xef, 0xa3, 0xb0, 0xf9, 0x15, 0x20, 0xdf, 0xbd, 0x41,
	0x5f, 0xb7, 0xb6, 0x53, 0x8f, 0x90, 0x51, 0x84, 0xb1, 0x63, 0xc1, 0x5c, 0x99, 0xd1, 0xea, 0xe8,
	0xbf, 0x23, 0x59, 0x3b, 0x8a, 0x0f, 0xbb, 0x86, 0x14, 0x9a, 0xa6, 0x87, 0x45, 0xac, 0x8c, 0x5a,
	0xc3, 0x32, 0x65, 0x1c, 0x75, 0xf8, 0xa3, 0xef, 0xb8, 0x1c, 0x58, 0xaa, 0xa3, 0xa2, 0xde, 0x20,
	0xd1, 0x21, 0x7e, 0x5c, 0xcf, 0x72, 0x51, 0x2c, 0x5f, 0x56, 0x65, 0x5a, 0xa5, 0xbc, 0xad, 0x52,
	0x7e, 0xdc, 0x56, 0x69, 0x24, 0x5f, 0x47, 0x3e, 0xce, 0xd9, 0x7e, 0xfd, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0xf0, 0xfd, 0x30, 0x89, 0x5c, 0x01, 0x00, 0x00,
}
