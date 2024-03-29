// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validator/proposed_disable_validator.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type ProposedDisableValidator struct {
	Address   string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Creator   string   `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Approvals []*Grant `protobuf:"bytes,3,rep,name=approvals,proto3" json:"approvals,omitempty"`
	Rejects   []*Grant `protobuf:"bytes,4,rep,name=rejects,proto3" json:"rejects,omitempty"`
}

func (m *ProposedDisableValidator) Reset()         { *m = ProposedDisableValidator{} }
func (m *ProposedDisableValidator) String() string { return proto.CompactTextString(m) }
func (*ProposedDisableValidator) ProtoMessage()    {}
func (*ProposedDisableValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7ffedaeb03ca643, []int{0}
}
func (m *ProposedDisableValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposedDisableValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposedDisableValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposedDisableValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposedDisableValidator.Merge(m, src)
}
func (m *ProposedDisableValidator) XXX_Size() int {
	return m.Size()
}
func (m *ProposedDisableValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposedDisableValidator.DiscardUnknown(m)
}

var xxx_messageInfo_ProposedDisableValidator proto.InternalMessageInfo

func (m *ProposedDisableValidator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *ProposedDisableValidator) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ProposedDisableValidator) GetApprovals() []*Grant {
	if m != nil {
		return m.Approvals
	}
	return nil
}

func (m *ProposedDisableValidator) GetRejects() []*Grant {
	if m != nil {
		return m.Rejects
	}
	return nil
}

func init() {
	proto.RegisterType((*ProposedDisableValidator)(nil), "zigbeealliance.distributedcomplianceledger.validator.ProposedDisableValidator")
}

func init() {
	proto.RegisterFile("validator/proposed_disable_validator.proto", fileDescriptor_f7ffedaeb03ca643)
}

var fileDescriptor_f7ffedaeb03ca643 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xbb, 0xad, 0x58, 0xba, 0xde, 0x16, 0x85, 0xd8, 0x43, 0x28, 0x9e, 0x8a, 0xb0, 0x09,
	0x54, 0x6f, 0x9e, 0x2c, 0x82, 0xe0, 0x49, 0x2a, 0x0a, 0x7a, 0x29, 0xd9, 0x64, 0x58, 0x23, 0x69,
	0x13, 0x92, 0xb4, 0xa8, 0x4f, 0xe1, 0xc3, 0xf4, 0x21, 0x3c, 0x16, 0x4f, 0x1e, 0xa5, 0x7d, 0x11,
	0xd9, 0x66, 0xb7, 0xeb, 0x4d, 0xd0, 0xe3, 0xce, 0xcc, 0xff, 0x7d, 0xb3, 0x99, 0xf8, 0x78, 0xce,
	0x94, 0x14, 0xcc, 0x6b, 0x4b, 0x8d, 0xd5, 0x46, 0x3b, 0x10, 0x63, 0x21, 0x1d, 0xcb, 0x14, 0x8c,
	0xb7, 0x2d, 0x62, 0xac, 0xf6, 0x3a, 0x39, 0x7d, 0x95, 0x79, 0x06, 0xc0, 0x94, 0x92, 0x6c, 0xca,
	0x81, 0x08, 0xe9, 0xbc, 0x95, 0xd9, 0xcc, 0x83, 0xe0, 0x7a, 0x62, 0x42, 0x55, 0x81, 0xc8, 0xc1,
	0x92, 0x6d, 0xb6, 0x7b, 0xc8, 0xb5, 0x9b, 0x68, 0x37, 0xde, 0x30, 0x68, 0xf8, 0x08, 0xc0, 0xee,
	0x41, 0x2d, 0xcf, 0x2d, 0x9b, 0xfa, 0x50, 0x3e, 0x5a, 0x34, 0x63, 0x74, 0x5d, 0x2e, 0x73, 0x11,
	0x76, 0xb9, 0xab, 0x06, 0x93, 0x41, 0xdc, 0x66, 0x42, 0x58, 0x70, 0x0e, 0x45, 0xbd, 0xa8, 0xdf,
	0x19, 0xa2, 0x8f, 0x45, 0xba, 0x5f, 0x62, 0xcf, 0x43, 0xe7, 0xc6, 0x5b, 0x39, 0xcd, 0x47, 0xd5,
	0x60, 0x91, 0xe1, 0x16, 0x8a, 0x38, 0x6a, 0xfe, 0x96, 0x29, 0x07, 0x93, 0xfb, 0xb8, 0xc3, 0x8c,
	0xb1, 0x7a, 0xce, 0x94, 0x43, 0xad, 0x5e, 0xab, 0xbf, 0x37, 0x38, 0x23, 0x7f, 0x79, 0x00, 0x72,
	0x59, 0xfc, 0xda, 0xa8, 0xa6, 0x25, 0xb7, 0x71, 0xdb, 0xc2, 0x13, 0x70, 0xef, 0xd0, 0xce, 0xff,
	0xc1, 0x15, 0x6b, 0x28, 0xde, 0x57, 0x38, 0x5a, 0xae, 0x70, 0xf4, 0xb5, 0xc2, 0xd1, 0xdb, 0x1a,
	0x37, 0x96, 0x6b, 0xdc, 0xf8, 0x5c, 0xe3, 0xc6, 0xc3, 0x55, 0x2e, 0xfd, 0xe3, 0x2c, 0x23, 0x5c,
	0x4f, 0x68, 0x30, 0xa5, 0x95, 0x8a, 0xfe, 0x50, 0xa5, 0xb5, 0x2b, 0x0d, 0x32, 0xfa, 0x4c, 0xeb,
	0x13, 0xf9, 0x17, 0x03, 0x2e, 0xdb, 0xdd, 0xdc, 0xe8, 0xe4, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xb6,
	0xd3, 0xb2, 0xf7, 0x39, 0x02, 0x00, 0x00,
}

func (m *ProposedDisableValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposedDisableValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposedDisableValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Rejects) > 0 {
		for iNdEx := len(m.Rejects) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rejects[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposedDisableValidator(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Approvals) > 0 {
		for iNdEx := len(m.Approvals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Approvals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposedDisableValidator(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintProposedDisableValidator(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintProposedDisableValidator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposedDisableValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposedDisableValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProposedDisableValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovProposedDisableValidator(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovProposedDisableValidator(uint64(l))
	}
	if len(m.Approvals) > 0 {
		for _, e := range m.Approvals {
			l = e.Size()
			n += 1 + l + sovProposedDisableValidator(uint64(l))
		}
	}
	if len(m.Rejects) > 0 {
		for _, e := range m.Rejects {
			l = e.Size()
			n += 1 + l + sovProposedDisableValidator(uint64(l))
		}
	}
	return n
}

func sovProposedDisableValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposedDisableValidator(x uint64) (n int) {
	return sovProposedDisableValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposedDisableValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposedDisableValidator
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
			return fmt.Errorf("proto: ProposedDisableValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposedDisableValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedDisableValidator
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
				return ErrInvalidLengthProposedDisableValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedDisableValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedDisableValidator
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
				return ErrInvalidLengthProposedDisableValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedDisableValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approvals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedDisableValidator
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
				return ErrInvalidLengthProposedDisableValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposedDisableValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Approvals = append(m.Approvals, &Grant{})
			if err := m.Approvals[len(m.Approvals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rejects", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedDisableValidator
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
				return ErrInvalidLengthProposedDisableValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposedDisableValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rejects = append(m.Rejects, &Grant{})
			if err := m.Rejects[len(m.Rejects)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposedDisableValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposedDisableValidator
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
func skipProposedDisableValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposedDisableValidator
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
					return 0, ErrIntOverflowProposedDisableValidator
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
					return 0, ErrIntOverflowProposedDisableValidator
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
				return 0, ErrInvalidLengthProposedDisableValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposedDisableValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposedDisableValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposedDisableValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposedDisableValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposedDisableValidator = fmt.Errorf("proto: unexpected end of group")
)
