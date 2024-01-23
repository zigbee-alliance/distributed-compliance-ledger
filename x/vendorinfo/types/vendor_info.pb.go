// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vendorinfo/types/vendor_info.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	proto "github.com/cosmos/gogoproto/proto"
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

type VendorInfo struct {
	VendorID             int32  `protobuf:"varint,1,opt,name=vendorID,proto3" json:"vendorID,omitempty"`
	VendorName           string `protobuf:"bytes,2,opt,name=vendorName,proto3" json:"vendorName,omitempty"`
	CompanyLegalName     string `protobuf:"bytes,3,opt,name=companyLegalName,proto3" json:"companyLegalName,omitempty"`
	CompanyPreferredName string `protobuf:"bytes,4,opt,name=companyPreferredName,proto3" json:"companyPreferredName,omitempty"`
	VendorLandingPageURL string `protobuf:"bytes,5,opt,name=vendorLandingPageURL,proto3" json:"vendorLandingPageURL,omitempty"`
	Creator              string `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *VendorInfo) Reset()         { *m = VendorInfo{} }
func (m *VendorInfo) String() string { return proto.CompactTextString(m) }
func (*VendorInfo) ProtoMessage()    {}
func (*VendorInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_626e7f57c2a6c82c, []int{0}
}
func (m *VendorInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VendorInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VendorInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VendorInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VendorInfo.Merge(m, src)
}
func (m *VendorInfo) XXX_Size() int {
	return m.Size()
}
func (m *VendorInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_VendorInfo.DiscardUnknown(m)
}

var xxx_messageInfo_VendorInfo proto.InternalMessageInfo

func (m *VendorInfo) GetVendorID() int32 {
	if m != nil {
		return m.VendorID
	}
	return 0
}

func (m *VendorInfo) GetVendorName() string {
	if m != nil {
		return m.VendorName
	}
	return ""
}

func (m *VendorInfo) GetCompanyLegalName() string {
	if m != nil {
		return m.CompanyLegalName
	}
	return ""
}

func (m *VendorInfo) GetCompanyPreferredName() string {
	if m != nil {
		return m.CompanyPreferredName
	}
	return ""
}

func (m *VendorInfo) GetVendorLandingPageURL() string {
	if m != nil {
		return m.VendorLandingPageURL
	}
	return ""
}

func (m *VendorInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*VendorInfo)(nil), "vendorinfo.types.VendorInfo")
}

func init() {
	proto.RegisterFile("vendorinfo/types/vendor_info.proto", fileDescriptor_626e7f57c2a6c82c)
}

var fileDescriptor_626e7f57c2a6c82c = []byte{
	// 314 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4e, 0x02, 0x31,
	0x14, 0xc6, 0x19, 0x14, 0xd4, 0xae, 0xc8, 0x84, 0xc5, 0xc8, 0xa2, 0x21, 0xac, 0x88, 0xc9, 0xd0,
	0x04, 0x4f, 0x20, 0x71, 0x43, 0x24, 0x86, 0x8c, 0xd1, 0x85, 0x1b, 0xd2, 0x99, 0x3e, 0x6a, 0x93,
	0x99, 0x96, 0xb4, 0xc5, 0x88, 0xa7, 0xe0, 0x30, 0x1e, 0xc2, 0x25, 0x71, 0xe5, 0xd2, 0x30, 0x17,
	0x31, 0xd3, 0xe2, 0x9f, 0x28, 0xbb, 0xbe, 0xdf, 0xf7, 0xfb, 0xd2, 0xa6, 0x0f, 0xf5, 0x1e, 0x41,
	0x32, 0xa5, 0x85, 0x9c, 0x2b, 0x62, 0x57, 0x0b, 0x30, 0xc4, 0x83, 0x59, 0x45, 0x06, 0x0b, 0xad,
	0xac, 0x0a, 0x5b, 0x3f, 0xce, 0xc0, 0x39, 0x9d, 0xd3, 0x4c, 0x99, 0x42, 0x99, 0x99, 0xcb, 0x89,
	0x1f, 0xbc, 0xdc, 0x5b, 0xd7, 0x11, 0xba, 0x73, 0xfe, 0x58, 0xce, 0x55, 0xd8, 0x41, 0xc7, 0xbe,
	0x3d, 0xbe, 0x8c, 0x82, 0x6e, 0xd0, 0x6f, 0x24, 0xdf, 0x73, 0x88, 0x11, 0xf2, 0xe7, 0x6b, 0x5a,
	0x40, 0x54, 0xef, 0x06, 0xfd, 0x93, 0xe4, 0x17, 0x09, 0xcf, 0x50, 0x2b, 0x53, 0xc5, 0x82, 0xca,
	0xd5, 0x04, 0x38, 0xcd, 0x9d, 0x75, 0xe0, 0xac, 0x7f, 0x3c, 0x1c, 0xa2, 0xf6, 0x8e, 0x4d, 0x35,
	0xcc, 0x41, 0x6b, 0x60, 0xce, 0x3f, 0x74, 0xfe, 0xde, 0xac, 0xea, 0xf8, 0xdb, 0x26, 0x54, 0x32,
	0x21, 0xf9, 0x94, 0x72, 0xb8, 0x4d, 0x26, 0x51, 0xc3, 0x77, 0xf6, 0x65, 0xe1, 0x10, 0x1d, 0x65,
	0x1a, 0xa8, 0x55, 0x3a, 0x6a, 0x56, 0xda, 0x28, 0x7a, 0x7b, 0x89, 0xdb, 0xbb, 0x1f, 0xb8, 0x60,
	0x4c, 0x83, 0x31, 0x37, 0x56, 0x0b, 0xc9, 0x93, 0x2f, 0x71, 0x04, 0xaf, 0x5b, 0x1c, 0x6c, 0xb6,
	0x38, 0xf8, 0xd8, 0xe2, 0x60, 0x5d, 0xe2, 0xda, 0xa6, 0xc4, 0xb5, 0xf7, 0x12, 0xd7, 0xee, 0xaf,
	0xb8, 0xb0, 0x0f, 0xcb, 0x74, 0x90, 0xa9, 0x82, 0x3c, 0x0b, 0x9e, 0x02, 0xc4, 0x34, 0xcf, 0x05,
	0x95, 0x19, 0x10, 0x26, 0x8c, 0xd5, 0x22, 0x5d, 0x5a, 0x60, 0x71, 0xf5, 0x7c, 0x8f, 0xe3, 0x1c,
	0x18, 0x07, 0x4d, 0x9e, 0xc8, 0xdf, 0xc5, 0xa5, 0x4d, 0xb7, 0x80, 0xf3, 0xcf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xb0, 0x24, 0xfc, 0x39, 0xd3, 0x01, 0x00, 0x00,
}

func (m *VendorInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VendorInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VendorInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintVendorInfo(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.VendorLandingPageURL) > 0 {
		i -= len(m.VendorLandingPageURL)
		copy(dAtA[i:], m.VendorLandingPageURL)
		i = encodeVarintVendorInfo(dAtA, i, uint64(len(m.VendorLandingPageURL)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.CompanyPreferredName) > 0 {
		i -= len(m.CompanyPreferredName)
		copy(dAtA[i:], m.CompanyPreferredName)
		i = encodeVarintVendorInfo(dAtA, i, uint64(len(m.CompanyPreferredName)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.CompanyLegalName) > 0 {
		i -= len(m.CompanyLegalName)
		copy(dAtA[i:], m.CompanyLegalName)
		i = encodeVarintVendorInfo(dAtA, i, uint64(len(m.CompanyLegalName)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.VendorName) > 0 {
		i -= len(m.VendorName)
		copy(dAtA[i:], m.VendorName)
		i = encodeVarintVendorInfo(dAtA, i, uint64(len(m.VendorName)))
		i--
		dAtA[i] = 0x12
	}
	if m.VendorID != 0 {
		i = encodeVarintVendorInfo(dAtA, i, uint64(m.VendorID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintVendorInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovVendorInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VendorInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VendorID != 0 {
		n += 1 + sovVendorInfo(uint64(m.VendorID))
	}
	l = len(m.VendorName)
	if l > 0 {
		n += 1 + l + sovVendorInfo(uint64(l))
	}
	l = len(m.CompanyLegalName)
	if l > 0 {
		n += 1 + l + sovVendorInfo(uint64(l))
	}
	l = len(m.CompanyPreferredName)
	if l > 0 {
		n += 1 + l + sovVendorInfo(uint64(l))
	}
	l = len(m.VendorLandingPageURL)
	if l > 0 {
		n += 1 + l + sovVendorInfo(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovVendorInfo(uint64(l))
	}
	return n
}

func sovVendorInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVendorInfo(x uint64) (n int) {
	return sovVendorInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VendorInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVendorInfo
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
			return fmt.Errorf("proto: VendorInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VendorInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorID", wireType)
			}
			m.VendorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVendorInfo
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
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVendorInfo
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
				return ErrInvalidLengthVendorInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVendorInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VendorName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompanyLegalName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVendorInfo
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
				return ErrInvalidLengthVendorInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVendorInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompanyLegalName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CompanyPreferredName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVendorInfo
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
				return ErrInvalidLengthVendorInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVendorInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CompanyPreferredName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorLandingPageURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVendorInfo
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
				return ErrInvalidLengthVendorInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVendorInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VendorLandingPageURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVendorInfo
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
				return ErrInvalidLengthVendorInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVendorInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVendorInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVendorInfo
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
func skipVendorInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVendorInfo
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
					return 0, ErrIntOverflowVendorInfo
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
					return 0, ErrIntOverflowVendorInfo
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
				return 0, ErrInvalidLengthVendorInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVendorInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVendorInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVendorInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVendorInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVendorInfo = fmt.Errorf("proto: unexpected end of group")
)
