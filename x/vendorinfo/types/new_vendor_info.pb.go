// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: vendorinfo/new_vendor_info.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type NewVendorInfo struct {
	Index      string      `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	VendorInfo *VendorInfo `protobuf:"bytes,2,opt,name=vendorInfo,proto3" json:"vendorInfo,omitempty"`
	Creator    string      `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *NewVendorInfo) Reset()         { *m = NewVendorInfo{} }
func (m *NewVendorInfo) String() string { return proto.CompactTextString(m) }
func (*NewVendorInfo) ProtoMessage()    {}
func (*NewVendorInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_f327a20890f42060, []int{0}
}
func (m *NewVendorInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NewVendorInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NewVendorInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NewVendorInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewVendorInfo.Merge(m, src)
}
func (m *NewVendorInfo) XXX_Size() int {
	return m.Size()
}
func (m *NewVendorInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NewVendorInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NewVendorInfo proto.InternalMessageInfo

func (m *NewVendorInfo) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *NewVendorInfo) GetVendorInfo() *VendorInfo {
	if m != nil {
		return m.VendorInfo
	}
	return nil
}

func (m *NewVendorInfo) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func init() {
	proto.RegisterType((*NewVendorInfo)(nil), "zigbeealliance.distributedcomplianceledger.vendorinfo.NewVendorInfo")
}

func init() { proto.RegisterFile("vendorinfo/new_vendor_info.proto", fileDescriptor_f327a20890f42060) }

var fileDescriptor_f327a20890f42060 = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x4b, 0xcd, 0x4b,
	0xc9, 0x2f, 0xca, 0xcc, 0x4b, 0xcb, 0xd7, 0xcf, 0x4b, 0x2d, 0x8f, 0x87, 0x70, 0xe3, 0x41, 0x7c,
	0xbd, 0x82, 0xa2, 0xfc, 0x92, 0x7c, 0x21, 0xd3, 0xaa, 0xcc, 0xf4, 0xa4, 0xd4, 0xd4, 0xc4, 0x9c,
	0x9c, 0xcc, 0xc4, 0xbc, 0xe4, 0x54, 0xbd, 0x94, 0xcc, 0xe2, 0x92, 0xa2, 0xcc, 0xa4, 0xd2, 0x92,
	0xd4, 0x94, 0xe4, 0xfc, 0xdc, 0x02, 0x88, 0x68, 0x4e, 0x6a, 0x4a, 0x7a, 0x6a, 0x91, 0x1e, 0xc2,
	0x30, 0x29, 0x19, 0x24, 0x83, 0x31, 0x0c, 0x55, 0x5a, 0xc4, 0xc8, 0xc5, 0xeb, 0x97, 0x5a, 0x1e,
	0x06, 0x96, 0xf0, 0xcc, 0x4b, 0xcb, 0x17, 0x12, 0xe1, 0x62, 0xcd, 0xcc, 0x4b, 0x49, 0xad, 0x90,
	0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0x12, 0xb9, 0xb8, 0xca, 0xe0, 0x6a, 0x24,
	0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0x1c, 0xf5, 0xc8, 0x72, 0x91, 0x1e, 0xc2, 0xb2, 0x20, 0x24,
	0x43, 0x85, 0x24, 0xb8, 0xd8, 0x93, 0x8b, 0x52, 0x13, 0x4b, 0xf2, 0x8b, 0x24, 0x98, 0xc1, 0x56,
	0xc3, 0xb8, 0x4e, 0xa9, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c,
	0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xe5, 0x9d,
	0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0x71, 0x8c, 0x2e, 0xcc, 0x35,
	0xfa, 0x48, 0xae, 0xd1, 0x45, 0x38, 0x47, 0x17, 0xe2, 0x1e, 0xfd, 0x0a, 0x7d, 0xa4, 0x70, 0x29,
	0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x07, 0x89, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x12,
	0xe1, 0x5d, 0xc8, 0x8b, 0x01, 0x00, 0x00,
}

func (m *NewVendorInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NewVendorInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NewVendorInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintNewVendorInfo(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if m.VendorInfo != nil {
		{
			size, err := m.VendorInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintNewVendorInfo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintNewVendorInfo(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintNewVendorInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovNewVendorInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NewVendorInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovNewVendorInfo(uint64(l))
	}
	if m.VendorInfo != nil {
		l = m.VendorInfo.Size()
		n += 1 + l + sovNewVendorInfo(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovNewVendorInfo(uint64(l))
	}
	return n
}

func sovNewVendorInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNewVendorInfo(x uint64) (n int) {
	return sovNewVendorInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NewVendorInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNewVendorInfo
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
			return fmt.Errorf("proto: NewVendorInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NewVendorInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewVendorInfo
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
				return ErrInvalidLengthNewVendorInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewVendorInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VendorInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewVendorInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNewVendorInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNewVendorInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VendorInfo == nil {
				m.VendorInfo = &VendorInfo{}
			}
			if err := m.VendorInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNewVendorInfo
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
				return ErrInvalidLengthNewVendorInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNewVendorInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNewVendorInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNewVendorInfo
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
func skipNewVendorInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNewVendorInfo
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
					return 0, ErrIntOverflowNewVendorInfo
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
					return 0, ErrIntOverflowNewVendorInfo
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
				return 0, ErrInvalidLengthNewVendorInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNewVendorInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNewVendorInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNewVendorInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNewVendorInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNewVendorInfo = fmt.Errorf("proto: unexpected end of group")
)
