// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vendorinfo/types/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreateVendorInfo struct {
	Creator              string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty" validate:"required"`
	VendorID             int32  `protobuf:"varint,2,opt,name=vendorID,proto3" json:"vendorID,omitempty" validate:"gte=1,lte=65535"`
	VendorName           string `protobuf:"bytes,3,opt,name=vendorName,proto3" json:"vendorName,omitempty" validate:"required,max=128"`
	CompanyLegalName     string `protobuf:"bytes,4,opt,name=companyLegalName,proto3" json:"companyLegalName,omitempty" validate:"required,max=256"`
	CompanyPreferredName string `protobuf:"bytes,5,opt,name=companyPreferredName,proto3" json:"companyPreferredName,omitempty" validate:"omitempty,max=256"`
	VendorLandingPageURL string `protobuf:"bytes,6,opt,name=vendorLandingPageURL,proto3" json:"vendorLandingPageURL,omitempty" validate:"omitempty,max=256,url"`
}

func (m *MsgCreateVendorInfo) Reset()         { *m = MsgCreateVendorInfo{} }
func (m *MsgCreateVendorInfo) String() string { return proto.CompactTextString(m) }
func (*MsgCreateVendorInfo) ProtoMessage()    {}
func (*MsgCreateVendorInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1cc0ab68adeeb3, []int{0}
}
func (m *MsgCreateVendorInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateVendorInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateVendorInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateVendorInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateVendorInfo.Merge(m, src)
}
func (m *MsgCreateVendorInfo) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateVendorInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateVendorInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateVendorInfo proto.InternalMessageInfo

func (m *MsgCreateVendorInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateVendorInfo) GetVendorID() int32 {
	if m != nil {
		return m.VendorID
	}
	return 0
}

func (m *MsgCreateVendorInfo) GetVendorName() string {
	if m != nil {
		return m.VendorName
	}
	return ""
}

func (m *MsgCreateVendorInfo) GetCompanyLegalName() string {
	if m != nil {
		return m.CompanyLegalName
	}
	return ""
}

func (m *MsgCreateVendorInfo) GetCompanyPreferredName() string {
	if m != nil {
		return m.CompanyPreferredName
	}
	return ""
}

func (m *MsgCreateVendorInfo) GetVendorLandingPageURL() string {
	if m != nil {
		return m.VendorLandingPageURL
	}
	return ""
}

type MsgCreateVendorInfoResponse struct {
}

func (m *MsgCreateVendorInfoResponse) Reset()         { *m = MsgCreateVendorInfoResponse{} }
func (m *MsgCreateVendorInfoResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateVendorInfoResponse) ProtoMessage()    {}
func (*MsgCreateVendorInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1cc0ab68adeeb3, []int{1}
}
func (m *MsgCreateVendorInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateVendorInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateVendorInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateVendorInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateVendorInfoResponse.Merge(m, src)
}
func (m *MsgCreateVendorInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateVendorInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateVendorInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateVendorInfoResponse proto.InternalMessageInfo

type MsgUpdateVendorInfo struct {
	Creator              string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty" validate:"required"`
	VendorID             int32  `protobuf:"varint,2,opt,name=vendorID,proto3" json:"vendorID,omitempty" validate:"gte=1,lte=65535"`
	VendorName           string `protobuf:"bytes,3,opt,name=vendorName,proto3" json:"vendorName,omitempty" validate:"omitempty,max=128"`
	CompanyLegalName     string `protobuf:"bytes,4,opt,name=companyLegalName,proto3" json:"companyLegalName,omitempty" validate:"omitempty,max=256"`
	CompanyPreferredName string `protobuf:"bytes,5,opt,name=companyPreferredName,proto3" json:"companyPreferredName,omitempty" validate:"omitempty,max=256"`
	VendorLandingPageURL string `protobuf:"bytes,6,opt,name=vendorLandingPageURL,proto3" json:"vendorLandingPageURL,omitempty" validate:"omitempty,max=256,url"`
}

func (m *MsgUpdateVendorInfo) Reset()         { *m = MsgUpdateVendorInfo{} }
func (m *MsgUpdateVendorInfo) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateVendorInfo) ProtoMessage()    {}
func (*MsgUpdateVendorInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1cc0ab68adeeb3, []int{2}
}
func (m *MsgUpdateVendorInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateVendorInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateVendorInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateVendorInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateVendorInfo.Merge(m, src)
}
func (m *MsgUpdateVendorInfo) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateVendorInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateVendorInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateVendorInfo proto.InternalMessageInfo

func (m *MsgUpdateVendorInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgUpdateVendorInfo) GetVendorID() int32 {
	if m != nil {
		return m.VendorID
	}
	return 0
}

func (m *MsgUpdateVendorInfo) GetVendorName() string {
	if m != nil {
		return m.VendorName
	}
	return ""
}

func (m *MsgUpdateVendorInfo) GetCompanyLegalName() string {
	if m != nil {
		return m.CompanyLegalName
	}
	return ""
}

func (m *MsgUpdateVendorInfo) GetCompanyPreferredName() string {
	if m != nil {
		return m.CompanyPreferredName
	}
	return ""
}

func (m *MsgUpdateVendorInfo) GetVendorLandingPageURL() string {
	if m != nil {
		return m.VendorLandingPageURL
	}
	return ""
}

type MsgUpdateVendorInfoResponse struct {
}

func (m *MsgUpdateVendorInfoResponse) Reset()         { *m = MsgUpdateVendorInfoResponse{} }
func (m *MsgUpdateVendorInfoResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateVendorInfoResponse) ProtoMessage()    {}
func (*MsgUpdateVendorInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf1cc0ab68adeeb3, []int{3}
}
func (m *MsgUpdateVendorInfoResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateVendorInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateVendorInfoResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateVendorInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateVendorInfoResponse.Merge(m, src)
}
func (m *MsgUpdateVendorInfoResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateVendorInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateVendorInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateVendorInfoResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateVendorInfo)(nil), "vendorinfo.types.MsgCreateVendorInfo")
	proto.RegisterType((*MsgCreateVendorInfoResponse)(nil), "vendorinfo.types.MsgCreateVendorInfoResponse")
	proto.RegisterType((*MsgUpdateVendorInfo)(nil), "vendorinfo.types.MsgUpdateVendorInfo")
	proto.RegisterType((*MsgUpdateVendorInfoResponse)(nil), "vendorinfo.types.MsgUpdateVendorInfoResponse")
}

func init() { proto.RegisterFile("vendorinfo/types/tx.proto", fileDescriptor_cf1cc0ab68adeeb3) }

var fileDescriptor_cf1cc0ab68adeeb3 = []byte{
	// 528 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x94, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x86, 0x6b, 0x42, 0x0b, 0xec, 0x29, 0x72, 0x73, 0x48, 0x53, 0xb0, 0x23, 0x0b, 0xa4, 0x1e,
	0x92, 0x58, 0x4d, 0x95, 0x0a, 0x55, 0x8a, 0x80, 0xc0, 0x25, 0x6a, 0x8a, 0x2a, 0x43, 0x11, 0xe2,
	0x52, 0x6d, 0xb2, 0x93, 0xed, 0x4a, 0xb6, 0xd7, 0xec, 0x6e, 0xaa, 0x84, 0xa7, 0xe0, 0x61, 0x78,
	0x08, 0x8e, 0x15, 0x27, 0x40, 0x28, 0x42, 0xc9, 0x1b, 0xf8, 0x09, 0x50, 0xd6, 0x09, 0x41, 0xb1,
	0xd3, 0x72, 0x44, 0xdc, 0xd6, 0x33, 0xf3, 0x7f, 0x33, 0xf6, 0xbf, 0x1e, 0xb4, 0x73, 0x09, 0x21,
	0xe1, 0x82, 0x85, 0x7d, 0xee, 0xaa, 0x51, 0x04, 0xd2, 0x55, 0xc3, 0x5a, 0x24, 0xb8, 0xe2, 0x66,
	0x7e, 0x99, 0xaa, 0xe9, 0x54, 0xc9, 0x49, 0x15, 0x27, 0x81, 0x73, 0x5d, 0xa3, 0x55, 0xa5, 0x02,
	0xe5, 0x94, 0xeb, 0xa3, 0x3b, 0x3b, 0xcd, 0xa3, 0x3b, 0x3d, 0x2e, 0x03, 0x2e, 0xcf, 0x93, 0x44,
	0xf2, 0x90, 0xa4, 0x9c, 0xef, 0x39, 0xb4, 0x7d, 0x22, 0xe9, 0x73, 0x01, 0x58, 0xc1, 0x1b, 0xcd,
	0x6b, 0x87, 0x7d, 0x6e, 0xb6, 0xd1, 0x9d, 0xde, 0x2c, 0xc6, 0x45, 0xd1, 0x28, 0x1b, 0x7b, 0xf7,
	0x5a, 0x6e, 0x3c, 0xb6, 0xb7, 0x2f, 0xb1, 0xcf, 0x08, 0x56, 0x70, 0xe4, 0x08, 0x78, 0x3f, 0x60,
	0x02, 0x88, 0xf3, 0xe5, 0x53, 0xb5, 0x30, 0x27, 0x3e, 0x23, 0x44, 0x80, 0x94, 0xaf, 0x94, 0x60,
	0x21, 0xf5, 0x16, 0x7a, 0xf3, 0x08, 0xdd, 0x4d, 0x06, 0x6d, 0xbf, 0x28, 0xde, 0x2a, 0x1b, 0x7b,
	0x9b, 0x2d, 0x2b, 0x1e, 0xdb, 0xa5, 0x25, 0x8b, 0x2a, 0x68, 0xee, 0x57, 0x7c, 0x05, 0xcd, 0xc3,
	0x46, 0xe3, 0xa0, 0xe1, 0x78, 0xbf, 0xeb, 0xcd, 0x27, 0x08, 0x25, 0xe7, 0x97, 0x38, 0x80, 0x62,
	0x4e, 0x4f, 0x62, 0xc7, 0x63, 0x7b, 0x37, 0x3d, 0x49, 0x25, 0xc0, 0xc3, 0xe6, 0x7e, 0xfd, 0xb1,
	0xe3, 0xfd, 0x21, 0x31, 0x8f, 0x51, 0xbe, 0xc7, 0x83, 0x08, 0x87, 0xa3, 0x0e, 0x50, 0xec, 0x6b,
	0xcc, 0xed, 0x1b, 0x31, 0xf5, 0xc6, 0xa1, 0xe3, 0xa5, 0x84, 0xe6, 0x6b, 0x54, 0x98, 0xc7, 0x4e,
	0x05, 0xf4, 0x41, 0x08, 0x20, 0x1a, 0xb8, 0xa9, 0x81, 0xe5, 0x78, 0x6c, 0xdf, 0x5f, 0x02, 0x79,
	0xc0, 0x14, 0x04, 0x91, 0x1a, 0x2d, 0x89, 0x99, 0x6a, 0xf3, 0x2d, 0x2a, 0x24, 0x03, 0x77, 0x70,
	0x48, 0x58, 0x48, 0x4f, 0x31, 0x85, 0x33, 0xaf, 0x53, 0xdc, 0xd2, 0xd4, 0x87, 0xf1, 0xd8, 0x2e,
	0x5f, 0x43, 0xad, 0x0c, 0x84, 0xef, 0x78, 0x99, 0x04, 0xe7, 0x01, 0xda, 0xcd, 0xf0, 0xd6, 0x03,
	0x19, 0xf1, 0x50, 0x82, 0xf3, 0x23, 0xf1, 0xfe, 0x2c, 0x22, 0xff, 0xa4, 0xf7, 0x4f, 0x33, 0xbc,
	0xbf, 0xf6, 0x1b, 0xa7, 0xcc, 0xef, 0xac, 0x35, 0xff, 0x66, 0xaf, 0xfe, 0x17, 0xf7, 0x57, 0xdd,
	0x5d, 0xb8, 0x5f, 0xff, 0x66, 0xa0, 0xdc, 0x89, 0xa4, 0xe6, 0x05, 0xca, 0xa7, 0xfe, 0xfe, 0x47,
	0xb5, 0xd5, 0xed, 0x53, 0xcb, 0xb8, 0x48, 0xa5, 0xea, 0x5f, 0x95, 0x2d, 0x3a, 0xce, 0x3a, 0xa5,
	0xee, 0x5a, 0x76, 0xa7, 0xd5, 0xb2, 0x35, 0x9d, 0xd6, 0xbd, 0x5b, 0x0b, 0x3e, 0x4f, 0x2c, 0xe3,
	0x6a, 0x62, 0x19, 0x3f, 0x27, 0x96, 0xf1, 0x71, 0x6a, 0x6d, 0x5c, 0x4d, 0xad, 0x8d, 0xaf, 0x53,
	0x6b, 0xe3, 0xdd, 0x31, 0x65, 0xea, 0x62, 0xd0, 0xad, 0xf5, 0x78, 0xe0, 0x7e, 0x60, 0xb4, 0x0b,
	0x50, 0xc5, 0xbe, 0xcf, 0x70, 0xd8, 0x03, 0x97, 0x30, 0xa9, 0x04, 0xeb, 0x0e, 0x14, 0x90, 0xea,
	0xcc, 0xab, 0x24, 0x5c, 0xf5, 0x81, 0x50, 0x10, 0xee, 0xd0, 0x5d, 0xdd, 0xbf, 0xdd, 0x2d, 0xbd,
	0x43, 0x0f, 0x7e, 0x05, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x78, 0xab, 0x73, 0xc7, 0x05, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateVendorInfo(ctx context.Context, in *MsgCreateVendorInfo, opts ...grpc.CallOption) (*MsgCreateVendorInfoResponse, error)
	UpdateVendorInfo(ctx context.Context, in *MsgUpdateVendorInfo, opts ...grpc.CallOption) (*MsgUpdateVendorInfoResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateVendorInfo(ctx context.Context, in *MsgCreateVendorInfo, opts ...grpc.CallOption) (*MsgCreateVendorInfoResponse, error) {
	out := new(MsgCreateVendorInfoResponse)
	err := c.cc.Invoke(ctx, "/vendorinfo.types.Msg/CreateVendorInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateVendorInfo(ctx context.Context, in *MsgUpdateVendorInfo, opts ...grpc.CallOption) (*MsgUpdateVendorInfoResponse, error) {
	out := new(MsgUpdateVendorInfoResponse)
	err := c.cc.Invoke(ctx, "/vendorinfo.types.Msg/UpdateVendorInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateVendorInfo(context.Context, *MsgCreateVendorInfo) (*MsgCreateVendorInfoResponse, error)
	UpdateVendorInfo(context.Context, *MsgUpdateVendorInfo) (*MsgUpdateVendorInfoResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateVendorInfo(ctx context.Context, req *MsgCreateVendorInfo) (*MsgCreateVendorInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVendorInfo not implemented")
}
func (*UnimplementedMsgServer) UpdateVendorInfo(ctx context.Context, req *MsgUpdateVendorInfo) (*MsgUpdateVendorInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateVendorInfo not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateVendorInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateVendorInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateVendorInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendorinfo.types.Msg/CreateVendorInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateVendorInfo(ctx, req.(*MsgCreateVendorInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateVendorInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateVendorInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateVendorInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/vendorinfo.types.Msg/UpdateVendorInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateVendorInfo(ctx, req.(*MsgUpdateVendorInfo))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "vendorinfo.types.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVendorInfo",
			Handler:    _Msg_CreateVendorInfo_Handler,
		},
		{
			MethodName: "UpdateVendorInfo",
			Handler:    _Msg_UpdateVendorInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "vendorinfo/types/tx.proto",
}

func (m *MsgCreateVendorInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateVendorInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateVendorInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VendorLandingPageURL) > 0 {
		i -= len(m.VendorLandingPageURL)
		copy(dAtA[i:], m.VendorLandingPageURL)
		i = encodeVarintTx(dAtA, i, uint64(len(m.VendorLandingPageURL)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.CompanyPreferredName) > 0 {
		i -= len(m.CompanyPreferredName)
		copy(dAtA[i:], m.CompanyPreferredName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CompanyPreferredName)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CompanyLegalName) > 0 {
		i -= len(m.CompanyLegalName)
		copy(dAtA[i:], m.CompanyLegalName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CompanyLegalName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.VendorName) > 0 {
		i -= len(m.VendorName)
		copy(dAtA[i:], m.VendorName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.VendorName)))
		i--
		dAtA[i] = 0x1a
	}
	if m.VendorID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.VendorID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateVendorInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateVendorInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateVendorInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateVendorInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateVendorInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateVendorInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VendorLandingPageURL) > 0 {
		i -= len(m.VendorLandingPageURL)
		copy(dAtA[i:], m.VendorLandingPageURL)
		i = encodeVarintTx(dAtA, i, uint64(len(m.VendorLandingPageURL)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.CompanyPreferredName) > 0 {
		i -= len(m.CompanyPreferredName)
		copy(dAtA[i:], m.CompanyPreferredName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CompanyPreferredName)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CompanyLegalName) > 0 {
		i -= len(m.CompanyLegalName)
		copy(dAtA[i:], m.CompanyLegalName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.CompanyLegalName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.VendorName) > 0 {
		i -= len(m.VendorName)
		copy(dAtA[i:], m.VendorName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.VendorName)))
		i--
		dAtA[i] = 0x1a
	}
	if m.VendorID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.VendorID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateVendorInfoResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateVendorInfoResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateVendorInfoResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateVendorInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.VendorID != 0 {
		n += 1 + sovTx(uint64(m.VendorID))
	}
	l = len(m.VendorName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CompanyLegalName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CompanyPreferredName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.VendorLandingPageURL)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateVendorInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateVendorInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.VendorID != 0 {
		n += 1 + sovTx(uint64(m.VendorID))
	}
	l = len(m.VendorName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CompanyLegalName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.CompanyPreferredName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.VendorLandingPageURL)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgUpdateVendorInfoResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateVendorInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateVendorInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateVendorInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorID", wireType)
			}
			m.VendorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VendorID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VendorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompanyLegalName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompanyLegalName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompanyPreferredName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompanyPreferredName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorLandingPageURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VendorLandingPageURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateVendorInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateVendorInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateVendorInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateVendorInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateVendorInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateVendorInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorID", wireType)
			}
			m.VendorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.VendorID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VendorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompanyLegalName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompanyLegalName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompanyPreferredName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompanyPreferredName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorLandingPageURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VendorLandingPageURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateVendorInfoResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateVendorInfoResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateVendorInfoResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
