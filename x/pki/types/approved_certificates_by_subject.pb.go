// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: zigbeealliance/distributedcomplianceledger/pki/approved_certificates_by_subject.proto

package types

import (
	fmt "fmt"
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

type ApprovedCertificatesBySubject struct {
	Subject       string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	SubjectKeyIds []string `protobuf:"bytes,2,rep,name=subjectKeyIds,proto3" json:"subjectKeyIds,omitempty"`
	SchemaVersion uint32   `protobuf:"varint,3,opt,name=schemaVersion,proto3" json:"schemaVersion,omitempty"`
}

func (m *ApprovedCertificatesBySubject) Reset()         { *m = ApprovedCertificatesBySubject{} }
func (m *ApprovedCertificatesBySubject) String() string { return proto.CompactTextString(m) }
func (*ApprovedCertificatesBySubject) ProtoMessage()    {}
func (*ApprovedCertificatesBySubject) Descriptor() ([]byte, []int) {
	return fileDescriptor_00fee9a70661a177, []int{0}
}
func (m *ApprovedCertificatesBySubject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ApprovedCertificatesBySubject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ApprovedCertificatesBySubject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ApprovedCertificatesBySubject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApprovedCertificatesBySubject.Merge(m, src)
}
func (m *ApprovedCertificatesBySubject) XXX_Size() int {
	return m.Size()
}
func (m *ApprovedCertificatesBySubject) XXX_DiscardUnknown() {
	xxx_messageInfo_ApprovedCertificatesBySubject.DiscardUnknown(m)
}

var xxx_messageInfo_ApprovedCertificatesBySubject proto.InternalMessageInfo

func (m *ApprovedCertificatesBySubject) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *ApprovedCertificatesBySubject) GetSubjectKeyIds() []string {
	if m != nil {
		return m.SubjectKeyIds
	}
	return nil
}

func (m *ApprovedCertificatesBySubject) GetSchemaVersion() uint32 {
	if m != nil {
		return m.SchemaVersion
	}
	return 0
}

func init() {
	proto.RegisterType((*ApprovedCertificatesBySubject)(nil), "zigbeealliance.distributedcomplianceledger.pki.ApprovedCertificatesBySubject")
}

func init() {
	proto.RegisterFile("zigbeealliance/distributedcomplianceledger/pki/approved_certificates_by_subject.proto", fileDescriptor_00fee9a70661a177)
}

var fileDescriptor_00fee9a70661a177 = []byte{
	// 260 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x0a, 0xad, 0xca, 0x4c, 0x4f,
	0x4a, 0x4d, 0x4d, 0xcc, 0xc9, 0xc9, 0x4c, 0xcc, 0x4b, 0x4e, 0xd5, 0x4f, 0xc9, 0x2c, 0x2e, 0x29,
	0xca, 0x4c, 0x2a, 0x2d, 0x49, 0x4d, 0x49, 0xce, 0xcf, 0x2d, 0x80, 0x88, 0xe6, 0xa4, 0xa6, 0xa4,
	0xa7, 0x16, 0xe9, 0x17, 0x64, 0x67, 0xea, 0x27, 0x16, 0x14, 0x14, 0xe5, 0x97, 0xa5, 0xa6, 0xc4,
	0x27, 0xa7, 0x16, 0x95, 0x64, 0xa6, 0x65, 0x26, 0x27, 0x96, 0xa4, 0x16, 0xc7, 0x27, 0x55, 0xc6,
	0x17, 0x97, 0x26, 0x65, 0xa5, 0x26, 0x97, 0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0xe9, 0xa1,
	0x1a, 0xab, 0x87, 0xc7, 0x58, 0xbd, 0x82, 0xec, 0x4c, 0xa5, 0x56, 0x46, 0x2e, 0x59, 0x47, 0xa8,
	0xd1, 0xce, 0x48, 0x26, 0x3b, 0x55, 0x06, 0x43, 0xcc, 0x15, 0x92, 0xe0, 0x62, 0x87, 0x5a, 0x21,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe3, 0x0a, 0xa9, 0x70, 0xf1, 0x42, 0x99, 0xde, 0xa9,
	0x95, 0x9e, 0x29, 0xc5, 0x12, 0x4c, 0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0xa8, 0x82, 0x60, 0x55, 0xc9,
	0x19, 0xa9, 0xb9, 0x89, 0x61, 0xa9, 0x45, 0xc5, 0x99, 0xf9, 0x79, 0x12, 0xcc, 0x0a, 0x8c, 0x1a,
	0xbc, 0x41, 0xa8, 0x82, 0x4e, 0x71, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0,
	0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10,
	0xe5, 0x92, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x0f, 0xf1, 0x9c, 0x2e,
	0xb6, 0x40, 0xd3, 0x45, 0x78, 0x4f, 0x17, 0x1a, 0x6c, 0x15, 0xe0, 0x80, 0x2b, 0xa9, 0x2c, 0x48,
	0x2d, 0x4e, 0x62, 0x03, 0x07, 0x8f, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x01, 0x11, 0xca, 0xc4,
	0x77, 0x01, 0x00, 0x00,
}

func (m *ApprovedCertificatesBySubject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ApprovedCertificatesBySubject) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ApprovedCertificatesBySubject) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.SchemaVersion != 0 {
		i = encodeVarintApprovedCertificatesBySubject(dAtA, i, uint64(m.SchemaVersion))
		i--
		dAtA[i] = 0x18
	}
	if len(m.SubjectKeyIds) > 0 {
		for iNdEx := len(m.SubjectKeyIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.SubjectKeyIds[iNdEx])
			copy(dAtA[i:], m.SubjectKeyIds[iNdEx])
			i = encodeVarintApprovedCertificatesBySubject(dAtA, i, uint64(len(m.SubjectKeyIds[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintApprovedCertificatesBySubject(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintApprovedCertificatesBySubject(dAtA []byte, offset int, v uint64) int {
	offset -= sovApprovedCertificatesBySubject(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ApprovedCertificatesBySubject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovApprovedCertificatesBySubject(uint64(l))
	}
	if len(m.SubjectKeyIds) > 0 {
		for _, s := range m.SubjectKeyIds {
			l = len(s)
			n += 1 + l + sovApprovedCertificatesBySubject(uint64(l))
		}
	}
	if m.SchemaVersion != 0 {
		n += 1 + sovApprovedCertificatesBySubject(uint64(m.SchemaVersion))
	}
	return n
}

func sovApprovedCertificatesBySubject(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozApprovedCertificatesBySubject(x uint64) (n int) {
	return sovApprovedCertificatesBySubject(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ApprovedCertificatesBySubject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowApprovedCertificatesBySubject
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
			return fmt.Errorf("proto: ApprovedCertificatesBySubject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ApprovedCertificatesBySubject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApprovedCertificatesBySubject
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
				return ErrInvalidLengthApprovedCertificatesBySubject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApprovedCertificatesBySubject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectKeyIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApprovedCertificatesBySubject
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
				return ErrInvalidLengthApprovedCertificatesBySubject
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthApprovedCertificatesBySubject
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectKeyIds = append(m.SubjectKeyIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SchemaVersion", wireType)
			}
			m.SchemaVersion = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowApprovedCertificatesBySubject
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
			skippy, err := skipApprovedCertificatesBySubject(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthApprovedCertificatesBySubject
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
func skipApprovedCertificatesBySubject(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowApprovedCertificatesBySubject
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
					return 0, ErrIntOverflowApprovedCertificatesBySubject
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
					return 0, ErrIntOverflowApprovedCertificatesBySubject
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
				return 0, ErrInvalidLengthApprovedCertificatesBySubject
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupApprovedCertificatesBySubject
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthApprovedCertificatesBySubject
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthApprovedCertificatesBySubject        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowApprovedCertificatesBySubject          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupApprovedCertificatesBySubject = fmt.Errorf("proto: unexpected end of group")
)
