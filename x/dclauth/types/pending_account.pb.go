// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dclauth/types/pending_account.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type PendingAccount struct {
	*Account `protobuf:"bytes,1,opt,name=account,proto3,embedded=account" json:"account,omitempty"`
}

func (m *PendingAccount) Reset()         { *m = PendingAccount{} }
func (m *PendingAccount) String() string { return proto.CompactTextString(m) }
func (*PendingAccount) ProtoMessage()    {}
func (*PendingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_6060dcc76949cc4b, []int{0}
}
func (m *PendingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PendingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PendingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PendingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PendingAccount.Merge(m, src)
}
func (m *PendingAccount) XXX_Size() int {
	return m.Size()
}
func (m *PendingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_PendingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_PendingAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PendingAccount)(nil), "dclauth.types.PendingAccount")
}

func init() {
	proto.RegisterFile("dclauth/types/pending_account.proto", fileDescriptor_6060dcc76949cc4b)
}

var fileDescriptor_6060dcc76949cc4b = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4e, 0x49, 0xce, 0x49,
	0x2c, 0x2d, 0xc9, 0xd0, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x2f, 0x48, 0xcd, 0x4b, 0xc9, 0xcc,
	0x4b, 0x8f, 0x4f, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x85, 0x2a, 0xd2, 0x03, 0x2b, 0x92, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xcb, 0xe8, 0x83,
	0x58, 0x10, 0x45, 0x52, 0xd2, 0xa8, 0x26, 0xa1, 0x98, 0xa0, 0xe4, 0xc1, 0xc5, 0x17, 0x00, 0x31,
	0xda, 0x11, 0x22, 0x2e, 0x64, 0xc6, 0xc5, 0x0e, 0x55, 0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d,
	0x24, 0xa6, 0x87, 0x62, 0x8b, 0x1e, 0x54, 0xa1, 0x13, 0xcb, 0x85, 0x7b, 0xf2, 0x8c, 0x41, 0x30,
	0xc5, 0x4e, 0x49, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3,
	0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xe5, 0x91, 0x9e,
	0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x5f, 0x95, 0x99, 0x9e, 0x94, 0x9a, 0xaa,
	0x9b, 0x98, 0x93, 0x93, 0x99, 0x98, 0x97, 0x9c, 0xaa, 0x9f, 0x92, 0x59, 0x5c, 0x52, 0x94, 0x99,
	0x54, 0x5a, 0x92, 0x9a, 0xa2, 0x9b, 0x9c, 0x9f, 0x5b, 0x00, 0x11, 0xd6, 0xcd, 0x49, 0x4d, 0x49,
	0x4f, 0x2d, 0xd2, 0xaf, 0xd0, 0x47, 0x71, 0x7b, 0x12, 0x1b, 0xd8, 0xd1, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xbd, 0x2c, 0x90, 0xb7, 0x1d, 0x01, 0x00, 0x00,
}

func (m *PendingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PendingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PendingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Account != nil {
		{
			size, err := m.Account.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintPendingAccount(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPendingAccount(dAtA []byte, offset int, v uint64) int {
	offset -= sovPendingAccount(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PendingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Account != nil {
		l = m.Account.Size()
		n += 1 + l + sovPendingAccount(uint64(l))
	}
	return n
}

func sovPendingAccount(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPendingAccount(x uint64) (n int) {
	return sovPendingAccount(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PendingAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPendingAccount
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
			return fmt.Errorf("proto: PendingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PendingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPendingAccount
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
				return ErrInvalidLengthPendingAccount
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPendingAccount
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Account == nil {
				m.Account = &Account{}
			}
			if err := m.Account.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPendingAccount(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPendingAccount
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
func skipPendingAccount(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPendingAccount
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
					return 0, ErrIntOverflowPendingAccount
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
					return 0, ErrIntOverflowPendingAccount
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
				return 0, ErrInvalidLengthPendingAccount
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPendingAccount
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPendingAccount
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPendingAccount        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPendingAccount          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPendingAccount = fmt.Errorf("proto: unexpected end of group")
)
