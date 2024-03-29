// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validator/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// GenesisState defines the validator module's genesis state.
type GenesisState struct {
	ValidatorList                []Validator                `protobuf:"bytes,1,rep,name=validatorList,proto3" json:"validatorList"`
	LastValidatorPowerList       []LastValidatorPower       `protobuf:"bytes,2,rep,name=lastValidatorPowerList,proto3" json:"lastValidatorPowerList"`
	ProposedDisableValidatorList []ProposedDisableValidator `protobuf:"bytes,3,rep,name=proposedDisableValidatorList,proto3" json:"proposedDisableValidatorList"`
	DisabledValidatorList        []DisabledValidator        `protobuf:"bytes,4,rep,name=disabledValidatorList,proto3" json:"disabledValidatorList"`
	RejectedValidatorList        []RejectedDisableValidator `protobuf:"bytes,5,rep,name=rejectedValidatorList,proto3" json:"rejectedValidatorList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_8143c6ee7ddaa59a, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetValidatorList() []Validator {
	if m != nil {
		return m.ValidatorList
	}
	return nil
}

func (m *GenesisState) GetLastValidatorPowerList() []LastValidatorPower {
	if m != nil {
		return m.LastValidatorPowerList
	}
	return nil
}

func (m *GenesisState) GetProposedDisableValidatorList() []ProposedDisableValidator {
	if m != nil {
		return m.ProposedDisableValidatorList
	}
	return nil
}

func (m *GenesisState) GetDisabledValidatorList() []DisabledValidator {
	if m != nil {
		return m.DisabledValidatorList
	}
	return nil
}

func (m *GenesisState) GetRejectedValidatorList() []RejectedDisableValidator {
	if m != nil {
		return m.RejectedValidatorList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "zigbeealliance.distributedcomplianceledger.validator.GenesisState")
}

func init() { proto.RegisterFile("validator/genesis.proto", fileDescriptor_8143c6ee7ddaa59a) }

var fileDescriptor_8143c6ee7ddaa59a = []byte{
	// 393 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4b, 0xcc, 0xc9,
	0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0xd2, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x32, 0xa9, 0xca, 0x4c, 0x4f, 0x4a, 0x4d, 0x4d, 0xcc, 0xc9, 0xc9, 0x4c,
	0xcc, 0x4b, 0x4e, 0xd5, 0x4b, 0xc9, 0x2c, 0x2e, 0x29, 0xca, 0x4c, 0x2a, 0x2d, 0x49, 0x4d, 0x49,
	0xce, 0xcf, 0x2d, 0x80, 0x88, 0xe6, 0xa4, 0xa6, 0xa4, 0xa7, 0x16, 0xe9, 0xc1, 0xcd, 0x90, 0x92,
	0x44, 0x18, 0x07, 0x67, 0x41, 0x0c, 0x94, 0x52, 0x41, 0x48, 0xe5, 0x24, 0x16, 0x97, 0xc4, 0xc3,
	0xb9, 0xf1, 0x05, 0xf9, 0xe5, 0xa9, 0x30, 0x55, 0x5a, 0x08, 0x55, 0x05, 0x45, 0xf9, 0x05, 0xf9,
	0xc5, 0xa9, 0x29, 0xf1, 0x29, 0x99, 0xc5, 0x89, 0x49, 0x39, 0xa9, 0xf1, 0xe8, 0x26, 0x2a, 0x21,
	0xd4, 0x42, 0x95, 0xa4, 0xe0, 0x53, 0x53, 0x94, 0x9a, 0x95, 0x9a, 0x5c, 0x82, 0x45, 0x8d, 0x48,
	0x7a, 0x7e, 0x7a, 0x3e, 0x98, 0xa9, 0x0f, 0x62, 0x41, 0x44, 0x95, 0x36, 0xb1, 0x72, 0xf1, 0xb8,
	0x43, 0x82, 0x24, 0xb8, 0x24, 0xb1, 0x24, 0x55, 0x28, 0x9b, 0x8b, 0x17, 0xae, 0xd3, 0x27, 0xb3,
	0xb8, 0x44, 0x82, 0x51, 0x81, 0x59, 0x83, 0xdb, 0xc8, 0x5e, 0x8f, 0x9c, 0x90, 0xd2, 0x0b, 0x83,
	0xb1, 0x9c, 0x58, 0x4e, 0xdc, 0x93, 0x67, 0x08, 0x42, 0x35, 0x5b, 0xa8, 0x8d, 0x91, 0x4b, 0x0c,
	0x14, 0x4c, 0x70, 0x65, 0x01, 0xa0, 0x40, 0x02, 0x5b, 0xcb, 0x04, 0xb6, 0xd6, 0x83, 0x3c, 0x6b,
	0x7d, 0x30, 0xcc, 0x84, 0xda, 0x8f, 0xc3, 0x36, 0xa1, 0x19, 0x8c, 0x5c, 0x32, 0xb0, 0x98, 0x70,
	0x81, 0x84, 0x72, 0x18, 0x4a, 0x28, 0x30, 0x83, 0x9d, 0xe3, 0x47, 0x9e, 0x73, 0x02, 0x70, 0x98,
	0x0c, 0x75, 0x14, 0x5e, 0x9b, 0x85, 0x9a, 0x19, 0xb9, 0x44, 0x61, 0x11, 0x8f, 0xea, 0x26, 0x16,
	0xb0, 0x9b, 0xdc, 0xc9, 0x73, 0x93, 0x0b, 0xba, 0x91, 0x50, 0xc7, 0x60, 0xb7, 0x4b, 0xa8, 0x8b,
	0x91, 0x4b, 0x14, 0x96, 0xb4, 0x50, 0x5d, 0xc1, 0x4a, 0x49, 0xc8, 0x04, 0x41, 0x8d, 0xc4, 0x11,
	0x32, 0xd8, 0xad, 0x74, 0x4a, 0x39, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f,
	0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28,
	0xaf, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x88, 0x83, 0x74, 0x61,
	0x2e, 0xd2, 0x47, 0x72, 0x91, 0x2e, 0xc2, 0x49, 0xba, 0x10, 0x37, 0xe9, 0x57, 0x20, 0xb2, 0xb2,
	0x7e, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0x87, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x36, 0x1c, 0xd4, 0x18, 0x3d, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RejectedValidatorList) > 0 {
		for iNdEx := len(m.RejectedValidatorList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RejectedValidatorList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.DisabledValidatorList) > 0 {
		for iNdEx := len(m.DisabledValidatorList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DisabledValidatorList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.ProposedDisableValidatorList) > 0 {
		for iNdEx := len(m.ProposedDisableValidatorList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ProposedDisableValidatorList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.LastValidatorPowerList) > 0 {
		for iNdEx := len(m.LastValidatorPowerList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LastValidatorPowerList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ValidatorList) > 0 {
		for iNdEx := len(m.ValidatorList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ValidatorList) > 0 {
		for _, e := range m.ValidatorList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.LastValidatorPowerList) > 0 {
		for _, e := range m.LastValidatorPowerList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ProposedDisableValidatorList) > 0 {
		for _, e := range m.ProposedDisableValidatorList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DisabledValidatorList) > 0 {
		for _, e := range m.DisabledValidatorList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RejectedValidatorList) > 0 {
		for _, e := range m.RejectedValidatorList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorList = append(m.ValidatorList, Validator{})
			if err := m.ValidatorList[len(m.ValidatorList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastValidatorPowerList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastValidatorPowerList = append(m.LastValidatorPowerList, LastValidatorPower{})
			if err := m.LastValidatorPowerList[len(m.LastValidatorPowerList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposedDisableValidatorList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposedDisableValidatorList = append(m.ProposedDisableValidatorList, ProposedDisableValidator{})
			if err := m.ProposedDisableValidatorList[len(m.ProposedDisableValidatorList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisabledValidatorList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisabledValidatorList = append(m.DisabledValidatorList, DisabledValidator{})
			if err := m.DisabledValidatorList[len(m.DisabledValidatorList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RejectedValidatorList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RejectedValidatorList = append(m.RejectedValidatorList, RejectedDisableValidator{})
			if err := m.RejectedValidatorList[len(m.RejectedValidatorList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
