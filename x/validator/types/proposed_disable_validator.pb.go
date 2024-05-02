// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zigbeealliance/distributedcomplianceledger/validator/proposed_disable_validator.proto

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
	return fileDescriptor_f9591474ec756c15, []int{0}
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
	proto.RegisterFile("zigbeealliance/distributedcomplianceledger/validator/proposed_disable_validator.proto", fileDescriptor_f9591474ec756c15)
}

var fileDescriptor_f9591474ec756c15 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xb1, 0x4a, 0x03, 0x31,
	0x1c, 0xc6, 0x7b, 0xad, 0x58, 0x7a, 0x6e, 0x87, 0x43, 0xec, 0x10, 0x8a, 0x53, 0x97, 0xcb, 0x41,
	0x75, 0x73, 0xd1, 0x22, 0x08, 0x4e, 0x52, 0xa9, 0xa0, 0x4b, 0xc9, 0x25, 0x7f, 0xce, 0x48, 0xda,
	0x84, 0x24, 0x2d, 0xea, 0x53, 0xf8, 0x30, 0x7d, 0x08, 0xc7, 0xe2, 0xe4, 0x28, 0xbd, 0x17, 0x91,
	0x6b, 0xee, 0x5a, 0x05, 0x51, 0xa8, 0x63, 0xf2, 0xff, 0xbe, 0xdf, 0xff, 0x4b, 0xbe, 0x70, 0xf8,
	0x2c, 0xb2, 0x14, 0x80, 0x4a, 0x29, 0xe8, 0x84, 0x41, 0xc2, 0x85, 0x75, 0x46, 0xa4, 0x53, 0x07,
	0x9c, 0xa9, 0xb1, 0xf6, 0xb7, 0x12, 0x78, 0x06, 0x26, 0x99, 0x51, 0x29, 0x38, 0x75, 0xca, 0x24,
	0xda, 0x28, 0xad, 0x2c, 0xf0, 0x11, 0x17, 0x96, 0xa6, 0x12, 0x46, 0xeb, 0x11, 0xd1, 0x46, 0x39,
	0x15, 0x1d, 0x7f, 0xc7, 0x92, 0x5f, 0xb0, 0x64, 0xed, 0x6d, 0x1f, 0x30, 0x65, 0xc7, 0xca, 0x8e,
	0x56, 0x8c, 0xc4, 0x1f, 0x3c, 0xb0, 0x7d, 0xba, 0x55, 0xce, 0xcc, 0xd0, 0x89, 0xf3, 0x84, 0xc3,
	0x79, 0x3d, 0x44, 0x57, 0x65, 0xee, 0x73, 0x1f, 0xfb, 0xa6, 0x12, 0x46, 0xbd, 0xb0, 0x49, 0x39,
	0x37, 0x60, 0x2d, 0x0a, 0x3a, 0x41, 0xb7, 0xd5, 0x47, 0x6f, 0xf3, 0x78, 0xbf, 0x4c, 0x70, 0xe6,
	0x27, 0xd7, 0xce, 0x88, 0x49, 0x36, 0xa8, 0x84, 0x85, 0x87, 0x19, 0x28, 0xec, 0xa8, 0xfe, 0x97,
	0xa7, 0x14, 0x46, 0xb7, 0x61, 0x8b, 0x6a, 0x6d, 0xd4, 0x8c, 0x4a, 0x8b, 0x1a, 0x9d, 0x46, 0x77,
	0xaf, 0x77, 0x42, 0xb6, 0xf9, 0x2b, 0x72, 0x51, 0x3c, 0x6d, 0xb0, 0xa1, 0x45, 0xc3, 0xb0, 0x69,
	0xe0, 0x01, 0x98, 0xb3, 0x68, 0xe7, 0xff, 0xe0, 0x8a, 0xd5, 0xe7, 0xaf, 0x4b, 0x1c, 0x2c, 0x96,
	0x38, 0xf8, 0x58, 0xe2, 0xe0, 0x25, 0xc7, 0xb5, 0x45, 0x8e, 0x6b, 0xef, 0x39, 0xae, 0xdd, 0x5d,
	0x66, 0xc2, 0xdd, 0x4f, 0x53, 0xc2, 0xd4, 0x38, 0xf1, 0x9b, 0xe2, 0x9f, 0xea, 0x89, 0x37, 0xbb,
	0xe2, 0xb2, 0xa0, 0xc7, 0x2f, 0x15, 0xb9, 0x27, 0x0d, 0x36, 0xdd, 0x5d, 0x75, 0x74, 0xf4, 0x19,
	0x00, 0x00, 0xff, 0xff, 0xd5, 0xca, 0x13, 0xb9, 0x8f, 0x02, 0x00, 0x00,
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
