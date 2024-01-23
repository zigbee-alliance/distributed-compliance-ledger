// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: model/types/genesis.proto

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

// GenesisState defines the model module's genesis state.
type GenesisState struct {
	VendorProductsList []VendorProducts `protobuf:"bytes,1,rep,name=vendorProductsList,proto3" json:"vendorProductsList"`
	ModelList          []Model          `protobuf:"bytes,2,rep,name=modelList,proto3" json:"modelList"`
	ModelVersionList   []ModelVersion   `protobuf:"bytes,3,rep,name=modelVersionList,proto3" json:"modelVersionList"`
	ModelVersionsList  []ModelVersions  `protobuf:"bytes,4,rep,name=modelVersionsList,proto3" json:"modelVersionsList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf0581302bd2e3d5, []int{0}
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

func (m *GenesisState) GetVendorProductsList() []VendorProducts {
	if m != nil {
		return m.VendorProductsList
	}
	return nil
}

func (m *GenesisState) GetModelList() []Model {
	if m != nil {
		return m.ModelList
	}
	return nil
}

func (m *GenesisState) GetModelVersionList() []ModelVersion {
	if m != nil {
		return m.ModelVersionList
	}
	return nil
}

func (m *GenesisState) GetModelVersionsList() []ModelVersions {
	if m != nil {
		return m.ModelVersionsList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "model.types.GenesisState")
}

func init() { proto.RegisterFile("model/types/genesis.proto", fileDescriptor_cf0581302bd2e3d5) }

var fileDescriptor_cf0581302bd2e3d5 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcd, 0x4f, 0x49,
	0xcd, 0xd1, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0xd6, 0x4f, 0x4f, 0xcd, 0x4b, 0x2d, 0xce, 0x2c, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x06, 0x4b, 0xe9, 0x81, 0xa5, 0xa4, 0x14, 0x91, 0xd5,
	0x95, 0xa5, 0xe6, 0xa5, 0xe4, 0x17, 0xc5, 0x17, 0x14, 0xe5, 0xa7, 0x94, 0x26, 0x97, 0x40, 0xd5,
	0x4b, 0x89, 0x23, 0x2b, 0x81, 0xe8, 0x85, 0x48, 0xc8, 0x63, 0x48, 0xc4, 0x97, 0xa5, 0x16, 0x15,
	0x67, 0xe6, 0xe7, 0x41, 0x15, 0x28, 0xe0, 0x54, 0x00, 0x33, 0x5b, 0x24, 0x3d, 0x3f, 0x3d, 0x1f,
	0xcc, 0xd4, 0x07, 0xb1, 0x20, 0xa2, 0x4a, 0xdb, 0x98, 0xb8, 0x78, 0xdc, 0x21, 0x6e, 0x0e, 0x2e,
	0x49, 0x2c, 0x49, 0x15, 0x0a, 0xe4, 0x12, 0x82, 0xb8, 0x2d, 0x00, 0xea, 0x34, 0x9f, 0xcc, 0xe2,
	0x12, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x69, 0x3d, 0x24, 0xff, 0xe8, 0x85, 0xa1, 0x28,
	0x73, 0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x08, 0x8b, 0x66, 0x21, 0x33, 0x2e, 0x4e, 0xb0, 0x3e,
	0xb0, 0x49, 0x4c, 0x60, 0x93, 0x84, 0x50, 0x4c, 0xf2, 0x05, 0xb1, 0xa1, 0x06, 0x20, 0x94, 0x0a,
	0x79, 0x73, 0x09, 0x80, 0x39, 0x61, 0x10, 0x8f, 0x80, 0xb5, 0x33, 0x83, 0xb5, 0x4b, 0x62, 0x6a,
	0x87, 0x2a, 0x82, 0x9a, 0x82, 0xa1, 0x51, 0xc8, 0x8f, 0x4b, 0x10, 0x59, 0x0c, 0xe2, 0x2d, 0x16,
	0xb0, 0x69, 0x52, 0x38, 0x4d, 0x83, 0xf9, 0x0a, 0x53, 0xab, 0x53, 0xc2, 0x89, 0x47, 0x72, 0x8c,
	0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72,
	0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xb9, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0xea, 0x57, 0x65, 0xa6, 0x27, 0xa5, 0xa6, 0xea, 0x26, 0xe6, 0xe4, 0x64, 0x26, 0xe6, 0x25,
	0xa7, 0xea, 0xa7, 0x64, 0x16, 0x97, 0x14, 0x65, 0x26, 0x95, 0x96, 0xa4, 0xa6, 0xe8, 0x26, 0xe7,
	0xe7, 0x16, 0x40, 0x84, 0x75, 0x73, 0x52, 0x53, 0xd2, 0x53, 0x8b, 0xf4, 0x2b, 0xf4, 0x91, 0x62,
	0x31, 0x89, 0x0d, 0x1c, 0x43, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x3f, 0xeb, 0x61,
	0x60, 0x02, 0x00, 0x00,
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
	if len(m.ModelVersionsList) > 0 {
		for iNdEx := len(m.ModelVersionsList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ModelVersionsList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ModelVersionList) > 0 {
		for iNdEx := len(m.ModelVersionList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ModelVersionList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ModelList) > 0 {
		for iNdEx := len(m.ModelList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ModelList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.VendorProductsList) > 0 {
		for iNdEx := len(m.VendorProductsList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VendorProductsList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.VendorProductsList) > 0 {
		for _, e := range m.VendorProductsList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ModelList) > 0 {
		for _, e := range m.ModelList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ModelVersionList) > 0 {
		for _, e := range m.ModelVersionList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ModelVersionsList) > 0 {
		for _, e := range m.ModelVersionsList {
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
				return fmt.Errorf("proto: wrong wireType = %d for field VendorProductsList", wireType)
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
			m.VendorProductsList = append(m.VendorProductsList, VendorProducts{})
			if err := m.VendorProductsList[len(m.VendorProductsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModelList", wireType)
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
			m.ModelList = append(m.ModelList, Model{})
			if err := m.ModelList[len(m.ModelList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModelVersionList", wireType)
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
			m.ModelVersionList = append(m.ModelVersionList, ModelVersion{})
			if err := m.ModelVersionList[len(m.ModelVersionList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ModelVersionsList", wireType)
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
			m.ModelVersionsList = append(m.ModelVersionsList, ModelVersions{})
			if err := m.ModelVersionsList[len(m.ModelVersionsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
