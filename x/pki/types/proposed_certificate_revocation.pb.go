// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pki/proposed_certificate_revocation.proto

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

type ProposedCertificateRevocation struct {
	Subject       string   `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	SubjectKeyId  string   `protobuf:"bytes,2,opt,name=subjectKeyId,proto3" json:"subjectKeyId,omitempty"`
	Approvals     []*Grant `protobuf:"bytes,3,rep,name=approvals,proto3" json:"approvals,omitempty"`
	SubjectAsText string   `protobuf:"bytes,4,opt,name=subjectAsText,proto3" json:"subjectAsText,omitempty"`
	SerialNumber  string   `protobuf:"bytes,5,opt,name=serialNumber,proto3" json:"serialNumber,omitempty"`
	RevokeChild   bool     `protobuf:"varint,6,opt,name=revokeChild,proto3" json:"revokeChild,omitempty"`
}

func (m *ProposedCertificateRevocation) Reset()         { *m = ProposedCertificateRevocation{} }
func (m *ProposedCertificateRevocation) String() string { return proto.CompactTextString(m) }
func (*ProposedCertificateRevocation) ProtoMessage()    {}
func (*ProposedCertificateRevocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_24b0dc6e71a9ad57, []int{0}
}
func (m *ProposedCertificateRevocation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProposedCertificateRevocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProposedCertificateRevocation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProposedCertificateRevocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProposedCertificateRevocation.Merge(m, src)
}
func (m *ProposedCertificateRevocation) XXX_Size() int {
	return m.Size()
}
func (m *ProposedCertificateRevocation) XXX_DiscardUnknown() {
	xxx_messageInfo_ProposedCertificateRevocation.DiscardUnknown(m)
}

var xxx_messageInfo_ProposedCertificateRevocation proto.InternalMessageInfo

func (m *ProposedCertificateRevocation) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *ProposedCertificateRevocation) GetSubjectKeyId() string {
	if m != nil {
		return m.SubjectKeyId
	}
	return ""
}

func (m *ProposedCertificateRevocation) GetApprovals() []*Grant {
	if m != nil {
		return m.Approvals
	}
	return nil
}

func (m *ProposedCertificateRevocation) GetSubjectAsText() string {
	if m != nil {
		return m.SubjectAsText
	}
	return ""
}

func (m *ProposedCertificateRevocation) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *ProposedCertificateRevocation) GetRevokeChild() bool {
	if m != nil {
		return m.RevokeChild
	}
	return false
}

func init() {
	proto.RegisterType((*ProposedCertificateRevocation)(nil), "zigbeealliance.distributedcomplianceledger.pki.ProposedCertificateRevocation")
}

func init() {
	proto.RegisterFile("pki/proposed_certificate_revocation.proto", fileDescriptor_24b0dc6e71a9ad57)
}

var fileDescriptor_24b0dc6e71a9ad57 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcd, 0x4a, 0xf3, 0x40,
	0x14, 0x86, 0x9b, 0xf6, 0xfb, 0xaa, 0x9d, 0x2a, 0x42, 0x56, 0x83, 0x60, 0x08, 0xc5, 0x45, 0x5d,
	0x34, 0x01, 0xc5, 0x0b, 0xd0, 0x0a, 0x22, 0x82, 0x48, 0x74, 0xe5, 0xc2, 0x32, 0x99, 0x1c, 0xd3,
	0x63, 0xd3, 0xce, 0x30, 0x33, 0x29, 0xad, 0x57, 0xe1, 0x15, 0x78, 0x3d, 0x2e, 0xbb, 0x74, 0x29,
	0xed, 0x8d, 0x48, 0x7e, 0x6a, 0xd3, 0xa5, 0xbb, 0x99, 0x77, 0xe0, 0x3c, 0xef, 0x33, 0x87, 0x9c,
	0xc8, 0x11, 0xfa, 0x52, 0x09, 0x29, 0x34, 0x44, 0x03, 0x0e, 0xca, 0xe0, 0x0b, 0x72, 0x66, 0x60,
	0xa0, 0x60, 0x2a, 0x38, 0x33, 0x28, 0x26, 0x9e, 0x54, 0xc2, 0x08, 0xdb, 0x7b, 0xc3, 0x38, 0x04,
	0x60, 0x49, 0x82, 0x6c, 0xc2, 0xc1, 0x8b, 0x50, 0x1b, 0x85, 0x61, 0x6a, 0x20, 0xe2, 0x62, 0x2c,
	0x8b, 0x34, 0x81, 0x28, 0x06, 0xe5, 0xc9, 0x11, 0x1e, 0x1e, 0x64, 0xa3, 0x63, 0xc5, 0x26, 0xa6,
	0x18, 0xd0, 0xf9, 0xa8, 0x93, 0xa3, 0xfb, 0x12, 0xd5, 0xdf, 0x90, 0x82, 0x5f, 0x90, 0x4d, 0xc9,
	0x8e, 0x4e, 0xc3, 0x57, 0xe0, 0x86, 0x5a, 0xae, 0xd5, 0x6d, 0x05, 0xeb, 0xab, 0xdd, 0x21, 0x7b,
	0xe5, 0xf1, 0x16, 0xe6, 0x37, 0x11, 0xad, 0xe7, 0xcf, 0x5b, 0x99, 0xfd, 0x40, 0x5a, 0x4c, 0x4a,
	0x25, 0xa6, 0x2c, 0xd1, 0xb4, 0xe1, 0x36, 0xba, 0xed, 0xd3, 0xf3, 0x3f, 0x96, 0xf6, 0xae, 0xb3,
	0xbe, 0xc1, 0x66, 0x8e, 0x7d, 0x4c, 0xf6, 0x4b, 0xc8, 0x85, 0x7e, 0x84, 0x99, 0xa1, 0xff, 0x72,
	0xf2, 0x76, 0x98, 0xd7, 0x03, 0x85, 0x2c, 0xb9, 0x4b, 0xc7, 0x21, 0x28, 0xfa, 0xbf, 0xac, 0x57,
	0xc9, 0x6c, 0x97, 0xb4, 0xb3, 0x3f, 0x1d, 0x41, 0x7f, 0x88, 0x49, 0x44, 0x9b, 0xae, 0xd5, 0xdd,
	0x0d, 0xaa, 0xd1, 0xe5, 0xf3, 0xe7, 0xd2, 0xb1, 0x16, 0x4b, 0xc7, 0xfa, 0x5e, 0x3a, 0xd6, 0xfb,
	0xca, 0xa9, 0x2d, 0x56, 0x4e, 0xed, 0x6b, 0xe5, 0xd4, 0x9e, 0xae, 0x62, 0x34, 0xc3, 0x34, 0xf4,
	0xb8, 0x18, 0xfb, 0x85, 0x51, 0x6f, 0xad, 0xe4, 0x57, 0x94, 0x7a, 0x1b, 0xa7, 0x5e, 0x21, 0xe5,
	0xcf, 0xfc, 0x6c, 0x0d, 0x66, 0x2e, 0x41, 0x87, 0xcd, 0x7c, 0x0f, 0x67, 0x3f, 0x01, 0x00, 0x00,
	0xff, 0xff, 0xc3, 0xe4, 0xa0, 0xad, 0xf5, 0x01, 0x00, 0x00,
}

func (m *ProposedCertificateRevocation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProposedCertificateRevocation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProposedCertificateRevocation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RevokeChild {
		i--
		if m.RevokeChild {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if len(m.SerialNumber) > 0 {
		i -= len(m.SerialNumber)
		copy(dAtA[i:], m.SerialNumber)
		i = encodeVarintProposedCertificateRevocation(dAtA, i, uint64(len(m.SerialNumber)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SubjectAsText) > 0 {
		i -= len(m.SubjectAsText)
		copy(dAtA[i:], m.SubjectAsText)
		i = encodeVarintProposedCertificateRevocation(dAtA, i, uint64(len(m.SubjectAsText)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Approvals) > 0 {
		for iNdEx := len(m.Approvals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Approvals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintProposedCertificateRevocation(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.SubjectKeyId) > 0 {
		i -= len(m.SubjectKeyId)
		copy(dAtA[i:], m.SubjectKeyId)
		i = encodeVarintProposedCertificateRevocation(dAtA, i, uint64(len(m.SubjectKeyId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintProposedCertificateRevocation(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposedCertificateRevocation(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposedCertificateRevocation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProposedCertificateRevocation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovProposedCertificateRevocation(uint64(l))
	}
	l = len(m.SubjectKeyId)
	if l > 0 {
		n += 1 + l + sovProposedCertificateRevocation(uint64(l))
	}
	if len(m.Approvals) > 0 {
		for _, e := range m.Approvals {
			l = e.Size()
			n += 1 + l + sovProposedCertificateRevocation(uint64(l))
		}
	}
	l = len(m.SubjectAsText)
	if l > 0 {
		n += 1 + l + sovProposedCertificateRevocation(uint64(l))
	}
	l = len(m.SerialNumber)
	if l > 0 {
		n += 1 + l + sovProposedCertificateRevocation(uint64(l))
	}
	if m.RevokeChild {
		n += 2
	}
	return n
}

func sovProposedCertificateRevocation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposedCertificateRevocation(x uint64) (n int) {
	return sovProposedCertificateRevocation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProposedCertificateRevocation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposedCertificateRevocation
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
			return fmt.Errorf("proto: ProposedCertificateRevocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProposedCertificateRevocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificateRevocation
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
				return ErrInvalidLengthProposedCertificateRevocation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificateRevocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectKeyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificateRevocation
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
				return ErrInvalidLengthProposedCertificateRevocation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificateRevocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectKeyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approvals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificateRevocation
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
				return ErrInvalidLengthProposedCertificateRevocation
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificateRevocation
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
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectAsText", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificateRevocation
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
				return ErrInvalidLengthProposedCertificateRevocation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificateRevocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectAsText = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SerialNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificateRevocation
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
				return ErrInvalidLengthProposedCertificateRevocation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposedCertificateRevocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SerialNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevokeChild", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposedCertificateRevocation
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.RevokeChild = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipProposedCertificateRevocation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposedCertificateRevocation
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
func skipProposedCertificateRevocation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposedCertificateRevocation
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
					return 0, ErrIntOverflowProposedCertificateRevocation
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
					return 0, ErrIntOverflowProposedCertificateRevocation
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
				return 0, ErrInvalidLengthProposedCertificateRevocation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposedCertificateRevocation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposedCertificateRevocation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposedCertificateRevocation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposedCertificateRevocation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposedCertificateRevocation = fmt.Errorf("proto: unexpected end of group")
)
