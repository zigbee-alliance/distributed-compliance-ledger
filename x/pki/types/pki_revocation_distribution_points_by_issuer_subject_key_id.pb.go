// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pki/pki_revocation_distribution_points_by_issuer_subject_key_id.proto

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

type PkiRevocationDistributionPointsByIssuerSubjectKeyID struct {
	IssuerSubjectKeyID string                            `protobuf:"bytes,1,opt,name=issuerSubjectKeyID,proto3" json:"issuerSubjectKeyID,omitempty"`
	Points             []*PkiRevocationDistributionPoint `protobuf:"bytes,2,rep,name=points,proto3" json:"points,omitempty"`
}

func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) Reset() {
	*m = PkiRevocationDistributionPointsByIssuerSubjectKeyID{}
}
func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) String() string {
	return proto.CompactTextString(m)
}
func (*PkiRevocationDistributionPointsByIssuerSubjectKeyID) ProtoMessage() {}
func (*PkiRevocationDistributionPointsByIssuerSubjectKeyID) Descriptor() ([]byte, []int) {
	return fileDescriptor_304eb676640a574a, []int{0}
}
func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PkiRevocationDistributionPointsByIssuerSubjectKeyID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PkiRevocationDistributionPointsByIssuerSubjectKeyID.Merge(m, src)
}
func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) XXX_Size() int {
	return m.Size()
}
func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) XXX_DiscardUnknown() {
	xxx_messageInfo_PkiRevocationDistributionPointsByIssuerSubjectKeyID.DiscardUnknown(m)
}

var xxx_messageInfo_PkiRevocationDistributionPointsByIssuerSubjectKeyID proto.InternalMessageInfo

func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) GetIssuerSubjectKeyID() string {
	if m != nil {
		return m.IssuerSubjectKeyID
	}
	return ""
}

func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) GetPoints() []*PkiRevocationDistributionPoint {
	if m != nil {
		return m.Points
	}
	return nil
}

func init() {
	proto.RegisterType((*PkiRevocationDistributionPointsByIssuerSubjectKeyID)(nil), "zigbeealliance.distributedcomplianceledger.pki.PkiRevocationDistributionPointsByIssuerSubjectKeyID")
}

func init() {
	proto.RegisterFile("pki/pki_revocation_distribution_points_by_issuer_subject_key_id.proto", fileDescriptor_304eb676640a574a)
}

var fileDescriptor_304eb676640a574a = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0x1b, 0x85, 0x81, 0xf5, 0xd6, 0xd3, 0xf0, 0x10, 0x86, 0xa7, 0x81, 0x34, 0x05, 0xf7,
	0x0d, 0x46, 0x3d, 0x0c, 0x41, 0x46, 0xbd, 0x79, 0x30, 0x34, 0xed, 0xb3, 0x3e, 0xdb, 0x35, 0xa1,
	0x49, 0xc5, 0xf8, 0x29, 0xfc, 0x58, 0x5e, 0x84, 0x1d, 0x3d, 0x4a, 0xfb, 0x45, 0xc4, 0x46, 0xe7,
	0x40, 0x11, 0x77, 0xcc, 0xcb, 0xe3, 0xff, 0xfb, 0xbd, 0xbf, 0x7f, 0xa6, 0x4a, 0x8c, 0x54, 0x89,
	0xbc, 0x81, 0x7b, 0x99, 0xa5, 0x06, 0x65, 0xcd, 0x73, 0xd4, 0xa6, 0x41, 0xd1, 0x0e, 0x0f, 0x25,
	0xb1, 0x36, 0x9a, 0x0b, 0xcb, 0x51, 0xeb, 0x16, 0x1a, 0xae, 0x5b, 0x71, 0x07, 0x99, 0xe1, 0x25,
	0x58, 0x8e, 0x39, 0x53, 0x8d, 0x34, 0x32, 0x60, 0x8f, 0x58, 0x08, 0x80, 0xb4, 0xaa, 0x30, 0xad,
	0x33, 0x60, 0x9b, 0x08, 0xc8, 0x33, 0xb9, 0x52, 0x6e, 0x5a, 0x41, 0x5e, 0x40, 0xc3, 0x54, 0x89,
	0x47, 0x27, 0xff, 0xc2, 0xba, 0xf0, 0xe3, 0x17, 0xe2, 0xcf, 0x96, 0x25, 0x26, 0x9b, 0xd5, 0x78,
	0x6b, 0x73, 0x39, 0xf8, 0xcd, 0xed, 0x62, 0xb0, 0xbb, 0x74, 0x72, 0xe7, 0x60, 0x17, 0x71, 0xc0,
	0xfc, 0x00, 0x7f, 0x4c, 0xc7, 0x64, 0x42, 0xa6, 0x07, 0xc9, 0x2f, 0x3f, 0xc1, 0x8d, 0x3f, 0x72,
	0xd7, 0x8e, 0xf7, 0x26, 0xfb, 0xd3, 0xc3, 0xd3, 0x8b, 0x1d, 0xaf, 0x62, 0x7f, 0x4b, 0x26, 0x9f,
	0xe9, 0xf3, 0xeb, 0xe7, 0x8e, 0x92, 0x75, 0x47, 0xc9, 0x5b, 0x47, 0xc9, 0x53, 0x4f, 0xbd, 0x75,
	0x4f, 0xbd, 0xd7, 0x9e, 0x7a, 0x57, 0x71, 0x81, 0xe6, 0xb6, 0x15, 0x2c, 0x93, 0xab, 0xc8, 0xb1,
	0xc3, 0x2f, 0x78, 0xb4, 0x05, 0x0f, 0xbf, 0xe9, 0xa1, 0xc3, 0x47, 0x0f, 0x1f, 0x6d, 0x46, 0xc6,
	0x2a, 0xd0, 0x62, 0x34, 0xd4, 0x36, 0x7b, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x01, 0x2f, 0xb0, 0x5d,
	0xdc, 0x01, 0x00, 0x00,
}

func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Points) > 0 {
		for iNdEx := len(m.Points) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Points[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintPkiRevocationDistributionPointsByIssuerSubjectKeyId(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.IssuerSubjectKeyID) > 0 {
		i -= len(m.IssuerSubjectKeyID)
		copy(dAtA[i:], m.IssuerSubjectKeyID)
		i = encodeVarintPkiRevocationDistributionPointsByIssuerSubjectKeyId(dAtA, i, uint64(len(m.IssuerSubjectKeyID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPkiRevocationDistributionPointsByIssuerSubjectKeyId(dAtA []byte, offset int, v uint64) int {
	offset -= sovPkiRevocationDistributionPointsByIssuerSubjectKeyId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.IssuerSubjectKeyID)
	if l > 0 {
		n += 1 + l + sovPkiRevocationDistributionPointsByIssuerSubjectKeyId(uint64(l))
	}
	if len(m.Points) > 0 {
		for _, e := range m.Points {
			l = e.Size()
			n += 1 + l + sovPkiRevocationDistributionPointsByIssuerSubjectKeyId(uint64(l))
		}
	}
	return n
}

func sovPkiRevocationDistributionPointsByIssuerSubjectKeyId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPkiRevocationDistributionPointsByIssuerSubjectKeyId(x uint64) (n int) {
	return sovPkiRevocationDistributionPointsByIssuerSubjectKeyId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PkiRevocationDistributionPointsByIssuerSubjectKeyID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPkiRevocationDistributionPointsByIssuerSubjectKeyId
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
			return fmt.Errorf("proto: PkiRevocationDistributionPointsByIssuerSubjectKeyID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PkiRevocationDistributionPointsByIssuerSubjectKeyID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuerSubjectKeyID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPkiRevocationDistributionPointsByIssuerSubjectKeyId
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
				return ErrInvalidLengthPkiRevocationDistributionPointsByIssuerSubjectKeyId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPkiRevocationDistributionPointsByIssuerSubjectKeyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssuerSubjectKeyID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Points", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPkiRevocationDistributionPointsByIssuerSubjectKeyId
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
				return ErrInvalidLengthPkiRevocationDistributionPointsByIssuerSubjectKeyId
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPkiRevocationDistributionPointsByIssuerSubjectKeyId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Points = append(m.Points, &PkiRevocationDistributionPoint{})
			if err := m.Points[len(m.Points)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPkiRevocationDistributionPointsByIssuerSubjectKeyId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPkiRevocationDistributionPointsByIssuerSubjectKeyId
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
func skipPkiRevocationDistributionPointsByIssuerSubjectKeyId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPkiRevocationDistributionPointsByIssuerSubjectKeyId
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
					return 0, ErrIntOverflowPkiRevocationDistributionPointsByIssuerSubjectKeyId
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
					return 0, ErrIntOverflowPkiRevocationDistributionPointsByIssuerSubjectKeyId
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
				return 0, ErrInvalidLengthPkiRevocationDistributionPointsByIssuerSubjectKeyId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPkiRevocationDistributionPointsByIssuerSubjectKeyId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPkiRevocationDistributionPointsByIssuerSubjectKeyId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPkiRevocationDistributionPointsByIssuerSubjectKeyId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPkiRevocationDistributionPointsByIssuerSubjectKeyId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPkiRevocationDistributionPointsByIssuerSubjectKeyId = fmt.Errorf("proto: unexpected end of group")
)
