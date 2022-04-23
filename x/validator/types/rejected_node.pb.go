// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validator/rejected_node.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	proto "github.com/gogo/protobuf/proto"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

type RejectedNode struct {
	Address         string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Creator         string   `protobuf:"bytes,2,opt,name=creator,proto3" json:"creator,omitempty"`
	Approvals       []*Grant `protobuf:"bytes,3,rep,name=approvals,proto3" json:"approvals,omitempty"`
	RejectApprovals []*Grant `protobuf:"bytes,4,rep,name=rejectApprovals,proto3" json:"rejectApprovals,omitempty"`
}

func (m *RejectedNode) Reset()         { *m = RejectedNode{} }
func (m *RejectedNode) String() string { return proto.CompactTextString(m) }
func (*RejectedNode) ProtoMessage()    {}
func (*RejectedNode) Descriptor() ([]byte, []int) {
	return fileDescriptor_52c046860a4f93c9, []int{0}
}
func (m *RejectedNode) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RejectedNode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RejectedNode.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RejectedNode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RejectedNode.Merge(m, src)
}
func (m *RejectedNode) XXX_Size() int {
	return m.Size()
}
func (m *RejectedNode) XXX_DiscardUnknown() {
	xxx_messageInfo_RejectedNode.DiscardUnknown(m)
}

var xxx_messageInfo_RejectedNode proto.InternalMessageInfo

func (m *RejectedNode) GetAddress() sdk.ValAddress {
	valAddr, _ := sdk.ValAddressFromBech32(m.Address)
	return valAddr
}

func (m *RejectedNode) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *RejectedNode) GetApprovals() []*Grant {
	if m != nil {
		return m.Approvals
	}
	return nil
}

func (m *RejectedNode) GetRejectApprovals() []*Grant {
	if m != nil {
		return m.RejectApprovals
	}
	return nil
}

func init() {
	proto.RegisterType((*RejectedNode)(nil), "zigbeealliance.distributedcomplianceledger.validator.RejectedNode")
}

func init() { proto.RegisterFile("validator/rejected_node.proto", fileDescriptor_52c046860a4f93c9) }

var fileDescriptor_52c046860a4f93c9 = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0x31, 0x4b, 0xc3, 0x40,
	0x14, 0xc7, 0x9b, 0x56, 0x94, 0x46, 0x41, 0x08, 0x0a, 0xb1, 0xe0, 0x51, 0x9c, 0xba, 0xe4, 0x0e,
	0xaa, 0x9b, 0x53, 0xbb, 0x08, 0x0e, 0x0e, 0x71, 0xd2, 0xa5, 0x5c, 0xee, 0x1e, 0xf1, 0x24, 0xc9,
	0x1d, 0x77, 0xd7, 0xa2, 0x7e, 0x0a, 0x3f, 0x8c, 0x83, 0x1f, 0xc1, 0xb1, 0x38, 0x39, 0x4a, 0xf2,
	0x45, 0x24, 0xbd, 0xb4, 0x11, 0x17, 0x41, 0xc7, 0x7b, 0xf7, 0xde, 0xef, 0xc7, 0x7b, 0x7f, 0xff,
	0x78, 0x41, 0x33, 0xc1, 0xa9, 0x95, 0x9a, 0x68, 0xb8, 0x07, 0x66, 0x81, 0xcf, 0x0a, 0xc9, 0x01,
	0x2b, 0x2d, 0xad, 0x0c, 0xce, 0x9e, 0x44, 0x9a, 0x00, 0xd0, 0x2c, 0x13, 0xb4, 0x60, 0x80, 0xb9,
	0x30, 0x56, 0x8b, 0x64, 0x6e, 0x81, 0x33, 0x99, 0x2b, 0x57, 0xcd, 0x80, 0xa7, 0xa0, 0xf1, 0x86,
	0x34, 0x38, 0x62, 0xd2, 0xe4, 0xd2, 0xcc, 0x56, 0x0c, 0xe2, 0x1e, 0x0e, 0x38, 0x38, 0x6c, 0x7d,
	0xa9, 0xa6, 0x85, 0x75, 0xe5, 0x93, 0xd7, 0xae, 0xbf, 0x17, 0x37, 0xfe, 0x2b, 0xc9, 0x21, 0x18,
	0xfb, 0x3b, 0x94, 0x73, 0x0d, 0xc6, 0x84, 0xde, 0xd0, 0x1b, 0xf5, 0xa7, 0xe1, 0xfb, 0x4b, 0x74,
	0xd0, 0xa0, 0x26, 0xee, 0xe7, 0xda, 0x6a, 0x51, 0xa4, 0xf1, 0xba, 0xb1, 0x9e, 0x61, 0x1a, 0x6a,
	0x76, 0xd8, 0xfd, 0x6d, 0xa6, 0x69, 0x0c, 0x6e, 0xfc, 0x3e, 0x55, 0x4a, 0xcb, 0x05, 0xcd, 0x4c,
	0xd8, 0x1b, 0xf6, 0x46, 0xbb, 0xe3, 0x73, 0xfc, 0x97, 0xa5, 0xf1, 0x45, 0xbd, 0x4e, 0xdc, 0xd2,
	0x02, 0xf0, 0xf7, 0xdd, 0x49, 0x27, 0x1b, 0xc1, 0xd6, 0xff, 0x05, 0x3f, 0x99, 0x53, 0xfe, 0x56,
	0x22, 0x6f, 0x59, 0x22, 0xef, 0xb3, 0x44, 0xde, 0x73, 0x85, 0x3a, 0xcb, 0x0a, 0x75, 0x3e, 0x2a,
	0xd4, 0xb9, 0xbd, 0x4c, 0x85, 0xbd, 0x9b, 0x27, 0x98, 0xc9, 0x9c, 0x38, 0x63, 0xb4, 0x56, 0x92,
	0x6f, 0xca, 0xa8, 0x75, 0x46, 0x4e, 0x4a, 0x1e, 0x48, 0x1b, 0x93, 0x7d, 0x54, 0x60, 0x92, 0xed,
	0x55, 0x4e, 0xa7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x40, 0xcd, 0x09, 0x7c, 0x30, 0x02, 0x00,
	0x00,
}

func (m *RejectedNode) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RejectedNode) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RejectedNode) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RejectApprovals) > 0 {
		for iNdEx := len(m.RejectApprovals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RejectApprovals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRejectedNode(dAtA, i, uint64(size))
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
				i = encodeVarintRejectedNode(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRejectedNode(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRejectedNode(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRejectedNode(dAtA []byte, offset int, v uint64) int {
	offset -= sovRejectedNode(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RejectedNode) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRejectedNode(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRejectedNode(uint64(l))
	}
	if len(m.Approvals) > 0 {
		for _, e := range m.Approvals {
			l = e.Size()
			n += 1 + l + sovRejectedNode(uint64(l))
		}
	}
	if len(m.RejectApprovals) > 0 {
		for _, e := range m.RejectApprovals {
			l = e.Size()
			n += 1 + l + sovRejectedNode(uint64(l))
		}
	}
	return n
}

func sovRejectedNode(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRejectedNode(x uint64) (n int) {
	return sovRejectedNode(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RejectedNode) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRejectedNode
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
			return fmt.Errorf("proto: RejectedNode: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RejectedNode: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRejectedNode
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
				return ErrInvalidLengthRejectedNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRejectedNode
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
					return ErrIntOverflowRejectedNode
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
				return ErrInvalidLengthRejectedNode
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRejectedNode
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
					return ErrIntOverflowRejectedNode
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
				return ErrInvalidLengthRejectedNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRejectedNode
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
				return fmt.Errorf("proto: wrong wireType = %d for field RejectApprovals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRejectedNode
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
				return ErrInvalidLengthRejectedNode
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRejectedNode
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RejectApprovals = append(m.RejectApprovals, &Grant{})
			if err := m.RejectApprovals[len(m.RejectApprovals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRejectedNode(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRejectedNode
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
func skipRejectedNode(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRejectedNode
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
					return 0, ErrIntOverflowRejectedNode
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
					return 0, ErrIntOverflowRejectedNode
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
				return 0, ErrInvalidLengthRejectedNode
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRejectedNode
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRejectedNode
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRejectedNode        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRejectedNode          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRejectedNode = fmt.Errorf("proto: unexpected end of group")
)