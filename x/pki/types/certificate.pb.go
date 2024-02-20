// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pki/certificate.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type Certificate struct {
	PemCert          string   `protobuf:"bytes,1,opt,name=pemCert,proto3" json:"pemCert,omitempty"`
	SerialNumber     string   `protobuf:"bytes,2,opt,name=serialNumber,proto3" json:"serialNumber,omitempty"`
	Issuer           string   `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	AuthorityKeyId   string   `protobuf:"bytes,4,opt,name=authorityKeyId,proto3" json:"authorityKeyId,omitempty"`
	RootSubject      string   `protobuf:"bytes,5,opt,name=rootSubject,proto3" json:"rootSubject,omitempty"`
	RootSubjectKeyId string   `protobuf:"bytes,6,opt,name=rootSubjectKeyId,proto3" json:"rootSubjectKeyId,omitempty"`
	IsRoot           bool     `protobuf:"varint,7,opt,name=isRoot,proto3" json:"isRoot,omitempty"`
	Owner            string   `protobuf:"bytes,8,opt,name=owner,proto3" json:"owner,omitempty"`
	Subject          string   `protobuf:"bytes,9,opt,name=subject,proto3" json:"subject,omitempty"`
	SubjectKeyId     string   `protobuf:"bytes,10,opt,name=subjectKeyId,proto3" json:"subjectKeyId,omitempty"`
	Approvals        []*Grant `protobuf:"bytes,11,rep,name=approvals,proto3" json:"approvals,omitempty"`
	SubjectAsText    string   `protobuf:"bytes,12,opt,name=subjectAsText,proto3" json:"subjectAsText,omitempty"`
	Rejects          []*Grant `protobuf:"bytes,13,rep,name=rejects,proto3" json:"rejects,omitempty"`
	Vid              int32    `protobuf:"varint,14,opt,name=vid,proto3" json:"vid,omitempty" validate:"gte=1,lte=65535"`
	IsNoc            bool     `protobuf:"varint,15,opt,name=isNoc,proto3" json:"isNoc,omitempty"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_2657e3d88fce7825, []int{0}
}
func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return m.Size()
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetPemCert() string {
	if m != nil {
		return m.PemCert
	}
	return ""
}

func (m *Certificate) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *Certificate) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *Certificate) GetAuthorityKeyId() string {
	if m != nil {
		return m.AuthorityKeyId
	}
	return ""
}

func (m *Certificate) GetRootSubject() string {
	if m != nil {
		return m.RootSubject
	}
	return ""
}

func (m *Certificate) GetRootSubjectKeyId() string {
	if m != nil {
		return m.RootSubjectKeyId
	}
	return ""
}

func (m *Certificate) GetIsRoot() bool {
	if m != nil {
		return m.IsRoot
	}
	return false
}

func (m *Certificate) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Certificate) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Certificate) GetSubjectKeyId() string {
	if m != nil {
		return m.SubjectKeyId
	}
	return ""
}

func (m *Certificate) GetApprovals() []*Grant {
	if m != nil {
		return m.Approvals
	}
	return nil
}

func (m *Certificate) GetSubjectAsText() string {
	if m != nil {
		return m.SubjectAsText
	}
	return ""
}

func (m *Certificate) GetRejects() []*Grant {
	if m != nil {
		return m.Rejects
	}
	return nil
}

func (m *Certificate) GetVid() int32 {
	if m != nil {
		return m.Vid
	}
	return 0
}

func (m *Certificate) GetIsNoc() bool {
	if m != nil {
		return m.IsNoc
	}
	return false
}

func init() {
	proto.RegisterType((*Certificate)(nil), "zigbeealliance.distributedcomplianceledger.pki.Certificate")
}

func init() { proto.RegisterFile("pki/certificate.proto", fileDescriptor_2657e3d88fce7825) }

var fileDescriptor_2657e3d88fce7825 = []byte{
	// 495 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x92, 0xcf, 0x8e, 0xd3, 0x3c,
	0x14, 0xc5, 0x9b, 0xaf, 0xd3, 0x76, 0xea, 0xce, 0x9f, 0x4f, 0x56, 0x41, 0xa6, 0x8b, 0x10, 0x55,
	0x08, 0x55, 0x88, 0x26, 0xc0, 0xa8, 0x2c, 0x90, 0x66, 0x31, 0x05, 0x09, 0x21, 0xa4, 0x41, 0x4a,
	0x59, 0xb1, 0x00, 0x39, 0xc9, 0x25, 0x63, 0x26, 0xa9, 0x23, 0xdb, 0x29, 0x53, 0x9e, 0x82, 0x87,
	0xe1, 0x21, 0x58, 0x8e, 0x58, 0xb1, 0x42, 0xa8, 0x7d, 0x03, 0x78, 0x01, 0x64, 0x3b, 0x55, 0x33,
	0xb0, 0x42, 0xec, 0x72, 0xce, 0xbd, 0xb9, 0xf7, 0xa7, 0x73, 0x8d, 0xae, 0x15, 0xe7, 0x2c, 0x88,
	0x41, 0x28, 0xf6, 0x96, 0xc5, 0x54, 0x81, 0x5f, 0x08, 0xae, 0x38, 0xf6, 0x3f, 0xb0, 0x34, 0x02,
	0xa0, 0x59, 0xc6, 0xe8, 0x3c, 0x06, 0x3f, 0x61, 0x52, 0x09, 0x16, 0x95, 0x0a, 0x92, 0x98, 0xe7,
	0x85, 0x75, 0x33, 0x48, 0x52, 0x10, 0x7e, 0x71, 0xce, 0x06, 0x37, 0x62, 0x2e, 0x73, 0x2e, 0xdf,
	0x98, 0xbf, 0x03, 0x2b, 0xec, 0xa8, 0x41, 0x3f, 0xe5, 0x29, 0xb7, 0xbe, 0xfe, 0xaa, 0xdc, 0x43,
	0xbd, 0x37, 0x15, 0x74, 0xae, 0xac, 0x31, 0xfc, 0xb9, 0x83, 0x7a, 0x8f, 0xb7, 0x1c, 0x98, 0xa0,
	0x4e, 0x01, 0xb9, 0x76, 0x88, 0xe3, 0x39, 0xa3, 0x6e, 0xb8, 0x91, 0x78, 0x88, 0xf6, 0x24, 0x08,
	0x46, 0xb3, 0xd3, 0x32, 0x8f, 0x40, 0x90, 0xff, 0x4c, 0xf9, 0x8a, 0x87, 0xaf, 0xa3, 0x36, 0x93,
	0xb2, 0x04, 0x41, 0x9a, 0xa6, 0x5a, 0x29, 0x7c, 0x1b, 0x1d, 0xd0, 0x52, 0x9d, 0x71, 0xc1, 0xd4,
	0xf2, 0x39, 0x2c, 0x9f, 0x25, 0x64, 0xc7, 0xd4, 0x7f, 0x73, 0xb1, 0x87, 0x7a, 0x82, 0x73, 0x35,
	0x2b, 0xa3, 0x77, 0x10, 0x2b, 0xd2, 0x32, 0x4d, 0x75, 0x0b, 0xdf, 0x41, 0xff, 0xd7, 0xa4, 0x9d,
	0xd5, 0x36, 0x6d, 0x7f, 0xf8, 0x96, 0x26, 0xe4, 0x5c, 0x91, 0x8e, 0xe7, 0x8c, 0x76, 0xc3, 0x4a,
	0x61, 0x1f, 0xb5, 0xf8, 0xfb, 0x39, 0x08, 0xb2, 0xab, 0x7f, 0x9c, 0x92, 0x2f, 0x9f, 0xc6, 0xfd,
	0x2a, 0xbb, 0x93, 0x24, 0x11, 0x20, 0xe5, 0x4c, 0x09, 0x36, 0x4f, 0x43, 0xdb, 0xa6, 0x33, 0x91,
	0x15, 0x51, 0xd7, 0x66, 0x52, 0x49, 0x93, 0x49, 0x9d, 0x04, 0x55, 0x99, 0xd4, 0x29, 0x66, 0xa8,
	0x4b, 0x8b, 0x42, 0xf0, 0x05, 0xcd, 0x24, 0xe9, 0x79, 0xcd, 0x51, 0xef, 0xc1, 0xe4, 0x2f, 0xef,
	0xec, 0x3f, 0xd5, 0x17, 0x0b, 0xb7, 0x73, 0xf0, 0x2d, 0xb4, 0x5f, 0x2d, 0x39, 0x91, 0x2f, 0xe1,
	0x42, 0x91, 0x3d, 0xb3, 0xf9, 0xaa, 0x89, 0x5f, 0xa0, 0x8e, 0x00, 0xad, 0x25, 0xd9, 0xff, 0x97,
	0xc5, 0x9b, 0x29, 0xf8, 0x1e, 0x6a, 0x2e, 0x58, 0x42, 0x0e, 0x3c, 0x67, 0xd4, 0x9a, 0xba, 0x3f,
	0xbe, 0xdd, 0x1c, 0x2c, 0x68, 0xc6, 0x12, 0xaa, 0xe0, 0xd1, 0x30, 0x55, 0x70, 0x7c, 0xff, 0x6e,
	0xa6, 0xe0, 0xf8, 0xe1, 0x64, 0x72, 0x34, 0x19, 0x86, 0xba, 0x15, 0xf7, 0x51, 0x8b, 0xc9, 0x53,
	0x1e, 0x93, 0x43, 0x73, 0x02, 0x2b, 0xa6, 0xaf, 0x3f, 0xaf, 0x5c, 0xe7, 0x72, 0xe5, 0x3a, 0xdf,
	0x57, 0xae, 0xf3, 0x71, 0xed, 0x36, 0x2e, 0xd7, 0x6e, 0xe3, 0xeb, 0xda, 0x6d, 0xbc, 0x7a, 0x92,
	0x32, 0x75, 0x56, 0x46, 0x7e, 0xcc, 0xf3, 0xc0, 0xb2, 0x8e, 0x37, 0xb0, 0x41, 0x0d, 0x76, 0xbc,
	0xa5, 0x1d, 0x5b, 0xdc, 0xe0, 0x22, 0xd0, 0x6f, 0x5b, 0x2d, 0x0b, 0x90, 0x51, 0xdb, 0x3c, 0xee,
	0xa3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x46, 0x06, 0x75, 0x03, 0x67, 0x03, 0x00, 0x00,
}

func (m *Certificate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Certificate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Certificate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsNoc {
		i--
		if m.IsNoc {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x78
	}
	if m.Vid != 0 {
		i = encodeVarintCertificate(dAtA, i, uint64(m.Vid))
		i--
		dAtA[i] = 0x70
	}
	if len(m.Rejects) > 0 {
		for iNdEx := len(m.Rejects) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Rejects[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCertificate(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x6a
		}
	}
	if len(m.SubjectAsText) > 0 {
		i -= len(m.SubjectAsText)
		copy(dAtA[i:], m.SubjectAsText)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.SubjectAsText)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.Approvals) > 0 {
		for iNdEx := len(m.Approvals) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Approvals[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCertificate(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.SubjectKeyId) > 0 {
		i -= len(m.SubjectKeyId)
		copy(dAtA[i:], m.SubjectKeyId)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.SubjectKeyId)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.Subject) > 0 {
		i -= len(m.Subject)
		copy(dAtA[i:], m.Subject)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.Subject)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x42
	}
	if m.IsRoot {
		i--
		if m.IsRoot {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x38
	}
	if len(m.RootSubjectKeyId) > 0 {
		i -= len(m.RootSubjectKeyId)
		copy(dAtA[i:], m.RootSubjectKeyId)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.RootSubjectKeyId)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.RootSubject) > 0 {
		i -= len(m.RootSubject)
		copy(dAtA[i:], m.RootSubject)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.RootSubject)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AuthorityKeyId) > 0 {
		i -= len(m.AuthorityKeyId)
		copy(dAtA[i:], m.AuthorityKeyId)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.AuthorityKeyId)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SerialNumber) > 0 {
		i -= len(m.SerialNumber)
		copy(dAtA[i:], m.SerialNumber)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.SerialNumber)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.PemCert) > 0 {
		i -= len(m.PemCert)
		copy(dAtA[i:], m.PemCert)
		i = encodeVarintCertificate(dAtA, i, uint64(len(m.PemCert)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCertificate(dAtA []byte, offset int, v uint64) int {
	offset -= sovCertificate(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Certificate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.PemCert)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.SerialNumber)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.AuthorityKeyId)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.RootSubject)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.RootSubjectKeyId)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	if m.IsRoot {
		n += 2
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	l = len(m.SubjectKeyId)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	if len(m.Approvals) > 0 {
		for _, e := range m.Approvals {
			l = e.Size()
			n += 1 + l + sovCertificate(uint64(l))
		}
	}
	l = len(m.SubjectAsText)
	if l > 0 {
		n += 1 + l + sovCertificate(uint64(l))
	}
	if len(m.Rejects) > 0 {
		for _, e := range m.Rejects {
			l = e.Size()
			n += 1 + l + sovCertificate(uint64(l))
		}
	}
	if m.Vid != 0 {
		n += 1 + sovCertificate(uint64(m.Vid))
	}
	if m.IsNoc {
		n += 2
	}
	return n
}

func sovCertificate(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCertificate(x uint64) (n int) {
	return sovCertificate(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Certificate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCertificate
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
			return fmt.Errorf("proto: Certificate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Certificate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PemCert", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PemCert = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SerialNumber", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SerialNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorityKeyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorityKeyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootSubject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootSubject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RootSubjectKeyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RootSubjectKeyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsRoot", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
			m.IsRoot = bool(v != 0)
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectKeyId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectKeyId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approvals", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Approvals = append(m.Approvals, &Grant{})
			if err := m.Approvals[len(m.Approvals)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectAsText", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectAsText = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rejects", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
				return ErrInvalidLengthCertificate
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCertificate
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Rejects = append(m.Rejects, &Grant{})
			if err := m.Rejects[len(m.Rejects)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Vid", wireType)
			}
			m.Vid = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
		case 15:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsNoc", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCertificate
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
			m.IsNoc = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCertificate(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCertificate
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
func skipCertificate(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCertificate
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
					return 0, ErrIntOverflowCertificate
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
					return 0, ErrIntOverflowCertificate
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
				return 0, ErrInvalidLengthCertificate
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCertificate
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCertificate
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCertificate        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCertificate          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCertificate = fmt.Errorf("proto: unexpected end of group")
)
