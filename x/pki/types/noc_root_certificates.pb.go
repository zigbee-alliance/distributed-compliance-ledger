// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zigbeealliance/distributedcomplianceledger/pki/noc_root_certificates.proto

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

type NocRootCertificates struct {
	Vid           int32          `protobuf:"varint,1,opt,name=vid,proto3" json:"vid,omitempty" validate:"gte=1,lte=65535"`
	Certs         []*Certificate `protobuf:"bytes,2,rep,name=certs,proto3" json:"certs,omitempty"`
	SchemaVersion uint32         `protobuf:"varint,3,opt,name=schemaVersion,proto3" json:"schemaVersion,omitempty" validate:"gte=0,lte=65535"`
}

func (m *NocRootCertificates) Reset()         { *m = NocRootCertificates{} }
func (m *NocRootCertificates) String() string { return proto.CompactTextString(m) }
func (*NocRootCertificates) ProtoMessage()    {}
func (*NocRootCertificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_c641c0880cd46de8, []int{0}
}
func (m *NocRootCertificates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NocRootCertificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_NocRootCertificates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *NocRootCertificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NocRootCertificates.Merge(m, src)
}
func (m *NocRootCertificates) XXX_Size() int {
	return m.Size()
}
func (m *NocRootCertificates) XXX_DiscardUnknown() {
	xxx_messageInfo_NocRootCertificates.DiscardUnknown(m)
}

var xxx_messageInfo_NocRootCertificates proto.InternalMessageInfo

func (m *NocRootCertificates) GetVid() int32 {
	if m != nil {
		return m.Vid
	}
	return 0
}

func (m *NocRootCertificates) GetCerts() []*Certificate {
	if m != nil {
		return m.Certs
	}
	return nil
}

func (m *NocRootCertificates) GetSchemaVersion() uint32 {
	if m != nil {
		return m.SchemaVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*NocRootCertificates)(nil), "zigbeealliance.distributedcomplianceledger.pki.NocRootCertificates")
}

func init() {
	proto.RegisterFile("zigbeealliance/distributedcomplianceledger/pki/noc_root_certificates.proto", fileDescriptor_c641c0880cd46de8)
}

var fileDescriptor_c641c0880cd46de8 = []byte{
	// 315 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xb1, 0x4e, 0x32, 0x41,
	0x10, 0x80, 0xd9, 0x9f, 0xf0, 0x17, 0x67, 0x68, 0x4e, 0x0b, 0x42, 0xb1, 0x12, 0x2a, 0x0a, 0xb9,
	0x43, 0x09, 0x16, 0x1a, 0x12, 0x83, 0x54, 0x16, 0x26, 0x52, 0x58, 0x58, 0x48, 0xee, 0xf6, 0xc6,
	0x65, 0xc2, 0xc1, 0x5c, 0x76, 0x07, 0xa2, 0x3e, 0x85, 0x8f, 0x65, 0x49, 0x69, 0x65, 0x08, 0xbc,
	0x81, 0x4f, 0x60, 0xb8, 0x33, 0xf1, 0x48, 0x88, 0x09, 0xdd, 0x66, 0xb3, 0xf3, 0xed, 0x37, 0x9f,
	0x73, 0xf3, 0x8a, 0x3a, 0x04, 0x08, 0xe2, 0x18, 0x83, 0xa9, 0x02, 0x3f, 0x42, 0xcb, 0x06, 0xc3,
	0x19, 0x43, 0xa4, 0x68, 0x92, 0x64, 0xb7, 0x31, 0x44, 0x1a, 0x8c, 0x9f, 0x8c, 0xd1, 0x9f, 0x92,
	0x1a, 0x1a, 0x22, 0x1e, 0x2a, 0x30, 0x8c, 0x4f, 0xa8, 0x02, 0x06, 0xeb, 0x25, 0x86, 0x98, 0x5c,
	0x6f, 0x9b, 0xe5, 0xfd, 0xc1, 0xf2, 0x92, 0x31, 0x56, 0x8f, 0x34, 0x69, 0x4a, 0x47, 0xfd, 0xcd,
	0x29, 0xa3, 0x54, 0xaf, 0xf6, 0x34, 0xca, 0x89, 0x64, 0x84, 0xfa, 0x52, 0x38, 0x87, 0xb7, 0xa4,
	0x06, 0x44, 0x7c, 0x9d, 0xb3, 0x74, 0x5b, 0x4e, 0x71, 0x8e, 0x51, 0x45, 0xd4, 0x44, 0xa3, 0xd4,
	0x93, 0x5f, 0x9f, 0xc7, 0xd5, 0x79, 0x10, 0x63, 0x14, 0x30, 0x5c, 0xd4, 0x35, 0x43, 0xf7, 0xf4,
	0x24, 0x66, 0xe8, 0x9e, 0x77, 0x3a, 0xed, 0x4e, 0x7d, 0xb0, 0x79, 0xea, 0xde, 0x39, 0xa5, 0x0d,
	0xde, 0x56, 0xfe, 0xd5, 0x8a, 0x8d, 0x83, 0xb3, 0xcb, 0x3d, 0x37, 0xf4, 0x72, 0xdf, 0x0f, 0x32,
	0x92, 0xdb, 0x77, 0xca, 0x56, 0x8d, 0x60, 0x12, 0xdc, 0x83, 0xb1, 0x48, 0xd3, 0x4a, 0xb1, 0x26,
	0x1a, 0xe5, 0x5d, 0x3a, 0xad, 0xbc, 0xce, 0xf6, 0x50, 0xef, 0xf1, 0x7d, 0x25, 0xc5, 0x62, 0x25,
	0xc5, 0x72, 0x25, 0xc5, 0xdb, 0x5a, 0x16, 0x16, 0x6b, 0x59, 0xf8, 0x58, 0xcb, 0xc2, 0x43, 0x5f,
	0x23, 0x8f, 0x66, 0xa1, 0xa7, 0x68, 0xe2, 0x67, 0xb6, 0xcd, 0x5d, 0x29, 0x9b, 0xbf, 0xbe, 0xcd,
	0x9f, 0x98, 0xcf, 0x69, 0x4e, 0x7e, 0x49, 0xc0, 0x86, 0xff, 0xd3, 0x92, 0xed, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x4a, 0xba, 0x8c, 0xc7, 0x1f, 0x02, 0x00, 0x00,
}

func (m *NocRootCertificates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NocRootCertificates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NocRootCertificates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SchemaVersion != 0 {
		i = encodeVarintNocRootCertificates(dAtA, i, uint64(m.SchemaVersion))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Certs) > 0 {
		for iNdEx := len(m.Certs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Certs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNocRootCertificates(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.Vid != 0 {
		i = encodeVarintNocRootCertificates(dAtA, i, uint64(m.Vid))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintNocRootCertificates(dAtA []byte, offset int, v uint64) int {
	offset -= sovNocRootCertificates(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *NocRootCertificates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Vid != 0 {
		n += 1 + sovNocRootCertificates(uint64(m.Vid))
	}
	if len(m.Certs) > 0 {
		for _, e := range m.Certs {
			l = e.Size()
			n += 1 + l + sovNocRootCertificates(uint64(l))
		}
	}
	if m.SchemaVersion != 0 {
		n += 1 + sovNocRootCertificates(uint64(m.SchemaVersion))
	}
	return n
}

func sovNocRootCertificates(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNocRootCertificates(x uint64) (n int) {
	return sovNocRootCertificates(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *NocRootCertificates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNocRootCertificates
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
			return fmt.Errorf("proto: NocRootCertificates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NocRootCertificates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vid", wireType)
			}
			m.Vid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNocRootCertificates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Vid |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Certs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNocRootCertificates
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
				return ErrInvalidLengthNocRootCertificates
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNocRootCertificates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Certs = append(m.Certs, &Certificate{})
			if err := m.Certs[len(m.Certs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaVersion", wireType)
			}
			m.SchemaVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNocRootCertificates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SchemaVersion |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipNocRootCertificates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNocRootCertificates
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
func skipNocRootCertificates(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNocRootCertificates
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
					return 0, ErrIntOverflowNocRootCertificates
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
					return 0, ErrIntOverflowNocRootCertificates
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
				return 0, ErrInvalidLengthNocRootCertificates
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNocRootCertificates
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNocRootCertificates
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNocRootCertificates        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNocRootCertificates          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNocRootCertificates = fmt.Errorf("proto: unexpected end of group")
)
