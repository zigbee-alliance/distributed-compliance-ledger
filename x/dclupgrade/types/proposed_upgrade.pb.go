// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zigbeealliance/distributedcomplianceledger/dclupgrade/proposed_upgrade.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/x/upgrade/types"
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

type ProposedUpgrade struct {
	Plan      types.Plan `protobuf:"bytes,1,opt,name=plan,proto3" json:"plan"`
	Creator   string     `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Approvals []*Grant   `protobuf:"bytes,3,rep,name=approvals,proto3" json:"approvals,omitempty"`
	Rejects   []*Grant   `protobuf:"bytes,4,rep,name=rejects,proto3" json:"rejects,omitempty"`
}

func (m *ProposedUpgrade) Reset()         { *m = ProposedUpgrade{} }
func (m *ProposedUpgrade) String() string { return proto.CompactTextString(m) }
func (*ProposedUpgrade) ProtoMessage()    {}
func (*ProposedUpgrade) Descriptor() ([]byte, []int) {
	return fileDescriptor_0325dd5a8e62f2be, []int{0}
}
func (m *ProposedUpgrade) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposedUpgrade) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposedUpgrade.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposedUpgrade) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposedUpgrade.Merge(m, src)
}
func (m *ProposedUpgrade) XXX_Size() int {
	return m.Size()
}
func (m *ProposedUpgrade) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposedUpgrade.DiscardUnknown(m)
}

var xxx_messageInfo_ProposedUpgrade proto.InternalMessageInfo

func (m *ProposedUpgrade) GetPlan() types.Plan {
	if m != nil {
		return m.Plan
	}
	return types.Plan{}
}

func (m *ProposedUpgrade) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *ProposedUpgrade) GetApprovals() []*Grant {
	if m != nil {
		return m.Approvals
	}
	return nil
}

func (m *ProposedUpgrade) GetRejects() []*Grant {
	if m != nil {
		return m.Rejects
	}
	return nil
}

func init() {
	proto.RegisterType((*ProposedUpgrade)(nil), "zigbeealliance.distributedcomplianceledger.dclupgrade.ProposedUpgrade")
}

func init() {
	proto.RegisterFile("zigbeealliance/distributedcomplianceledger/dclupgrade/proposed_upgrade.proto", fileDescriptor_0325dd5a8e62f2be)
}

var fileDescriptor_0325dd5a8e62f2be = []byte{
	// 361 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x91, 0xbf, 0x4e, 0x2a, 0x41,
	0x14, 0xc6, 0x77, 0x81, 0x5c, 0xc2, 0x52, 0xdc, 0x64, 0x43, 0xb1, 0x97, 0xdc, 0xac, 0xc4, 0x58,
	0xd0, 0xec, 0x6c, 0xc0, 0x68, 0x65, 0x03, 0x8d, 0x85, 0x16, 0x04, 0xa3, 0x05, 0x0d, 0x99, 0xdd,
	0x39, 0x19, 0xd7, 0x0c, 0x3b, 0x93, 0x99, 0x81, 0xa8, 0x4f, 0xe1, 0xab, 0x98, 0xf8, 0x10, 0x94,
	0xc4, 0xca, 0xca, 0x18, 0x78, 0x11, 0x03, 0x33, 0x1b, 0x24, 0x31, 0x16, 0xc4, 0x6e, 0xcf, 0x9f,
	0xfd, 0x9d, 0xef, 0x9b, 0xcf, 0xbb, 0x7c, 0xcc, 0x68, 0x02, 0x80, 0x19, 0xcb, 0x70, 0x9e, 0x42,
	0x4c, 0x32, 0xa5, 0x65, 0x96, 0x4c, 0x35, 0x90, 0x94, 0x4f, 0x84, 0xe9, 0x32, 0x20, 0x14, 0x64,
	0x4c, 0x52, 0x36, 0x15, 0x54, 0x62, 0x02, 0xb1, 0x90, 0x5c, 0x70, 0x05, 0x64, 0x6c, 0x1b, 0x48,
	0x48, 0xae, 0xb9, 0x7f, 0xb2, 0x4b, 0x43, 0x3f, 0xd0, 0xd0, 0x96, 0xd6, 0x6c, 0x50, 0x4e, 0xf9,
	0x86, 0x10, 0xaf, 0xbf, 0x0c, 0xac, 0xf9, 0x2f, 0xe5, 0x6a, 0xc2, 0xd5, 0xd8, 0x0c, 0x4c, 0x61,
	0x47, 0x47, 0xa6, 0x8a, 0x0b, 0x39, 0xb3, 0x4e, 0x02, 0x1a, 0x77, 0xe2, 0x1d, 0x35, 0xcd, 0xde,
	0x7e, 0xde, 0xa8, 0xc4, 0xb9, 0x36, 0x88, 0xc3, 0xe7, 0x92, 0xf7, 0x77, 0x60, 0xbd, 0x5e, 0x9b,
	0xb9, 0x7f, 0xea, 0x55, 0x04, 0xc3, 0x79, 0xe0, 0xb6, 0xdc, 0x76, 0xbd, 0xfb, 0x1f, 0x59, 0x65,
	0xc5, 0x6d, 0xab, 0x05, 0x0d, 0x18, 0xce, 0xfb, 0x95, 0xf9, 0xfb, 0x81, 0x33, 0xdc, 0xec, 0xfb,
	0x5d, 0xaf, 0x9a, 0x4a, 0xc0, 0x9a, 0xcb, 0xa0, 0xd4, 0x72, 0xdb, 0xb5, 0x7e, 0xf0, 0xfa, 0x12,
	0x35, 0xec, 0xdf, 0x3d, 0x42, 0x24, 0x28, 0x75, 0xa5, 0x65, 0x96, 0xd3, 0x61, 0xb1, 0xe8, 0x8f,
	0xbc, 0x1a, 0x16, 0x42, 0xf2, 0x19, 0x66, 0x2a, 0x28, 0xb7, 0xca, 0xed, 0x7a, 0xf7, 0x0c, 0xed,
	0xf5, 0xc8, 0xe8, 0x7c, 0x6d, 0x6b, 0xb8, 0xc5, 0xf9, 0x37, 0x5e, 0x55, 0xc2, 0x1d, 0xa4, 0x5a,
	0x05, 0x95, 0x5f, 0x20, 0x17, 0xb0, 0x3e, 0xcc, 0x97, 0xa1, 0xbb, 0x58, 0x86, 0xee, 0xc7, 0x32,
	0x74, 0x9f, 0x56, 0xa1, 0xb3, 0x58, 0x85, 0xce, 0xdb, 0x2a, 0x74, 0x46, 0x17, 0x34, 0xd3, 0xb7,
	0xd3, 0x04, 0xa5, 0x7c, 0x12, 0x9b, 0x53, 0xd1, 0x77, 0xe1, 0x44, 0xdb, 0x63, 0x91, 0x8d, 0xe7,
	0xfe, 0x6b, 0x40, 0xfa, 0x41, 0x80, 0x4a, 0xfe, 0x6c, 0x12, 0x3a, 0xfe, 0x0c, 0x00, 0x00, 0xff,
	0xff, 0x3c, 0x63, 0x17, 0x3a, 0xc2, 0x02, 0x00, 0x00,
}

func (m *ProposedUpgrade) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposedUpgrade) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposedUpgrade) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
				i = encodeVarintProposedUpgrade(dAtA, i, uint64(size))
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
				i = encodeVarintProposedUpgrade(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintProposedUpgrade(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Plan.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintProposedUpgrade(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintProposedUpgrade(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposedUpgrade(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProposedUpgrade) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Plan.Size()
	n += 1 + l + sovProposedUpgrade(uint64(l))
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovProposedUpgrade(uint64(l))
	}
	if len(m.Approvals) > 0 {
		for _, e := range m.Approvals {
			l = e.Size()
			n += 1 + l + sovProposedUpgrade(uint64(l))
		}
	}
	if len(m.Rejects) > 0 {
		for _, e := range m.Rejects {
			l = e.Size()
			n += 1 + l + sovProposedUpgrade(uint64(l))
		}
	}
	return n
}

func sovProposedUpgrade(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposedUpgrade(x uint64) (n int) {
	return sovProposedUpgrade(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposedUpgrade) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposedUpgrade
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
			return fmt.Errorf("proto: ProposedUpgrade: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposedUpgrade: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Plan", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedUpgrade
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
				return ErrInvalidLengthProposedUpgrade
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposedUpgrade
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Plan.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedUpgrade
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
				return ErrInvalidLengthProposedUpgrade
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedUpgrade
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
					return ErrIntOverflowProposedUpgrade
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
				return ErrInvalidLengthProposedUpgrade
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposedUpgrade
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
					return ErrIntOverflowProposedUpgrade
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
				return ErrInvalidLengthProposedUpgrade
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposedUpgrade
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
			skippy, err := skipProposedUpgrade(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposedUpgrade
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
func skipProposedUpgrade(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposedUpgrade
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
					return 0, ErrIntOverflowProposedUpgrade
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
					return 0, ErrIntOverflowProposedUpgrade
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
				return 0, ErrInvalidLengthProposedUpgrade
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposedUpgrade
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposedUpgrade
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposedUpgrade        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposedUpgrade          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposedUpgrade = fmt.Errorf("proto: unexpected end of group")
)
