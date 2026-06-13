package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	testconstants "github.com/zigbee-alliance/distributed-compliance-ledger/integration_tests/constants"
	"github.com/zigbee-alliance/distributed-compliance-ledger/x/pki/types"
)

type RootCertOptions struct {
	PemCert      string
	Info         string
	Subject      string
	SubjectKeyID string
	Vid          int32
}

func CreateTestRootCertOptions() *RootCertOptions {
	return &RootCertOptions{
		PemCert:      testconstants.RootCertPem,
		Info:         testconstants.Info,
		Subject:      testconstants.RootSubject,
		SubjectKeyID: testconstants.RootSubjectKeyID,
		Vid:          testconstants.Vid,
	}
}

func CreateRootWithVidOptions() *RootCertOptions {
	return &RootCertOptions{
		PemCert:      testconstants.RootCertWithVid,
		Info:         testconstants.Info,
		Subject:      testconstants.RootCertWithVidSubject,
		SubjectKeyID: testconstants.RootCertWithVidSubjectKeyID,
		Vid:          testconstants.RootCertWithVidVid,
	}
}

func CreatePAACertWithNumericVidOptions() *RootCertOptions {
	return &RootCertOptions{
		PemCert:      testconstants.PAACertWithNumericVid,
		Info:         testconstants.Info,
		Subject:      testconstants.PAACertWithNumericVidSubject,
		SubjectKeyID: testconstants.PAACertWithNumericVidSubjectKeyID,
		Vid:          testconstants.PAACertWithNumericVidVid,
	}
}

func CreatePAACertNoVidOptions(vid int32) *RootCertOptions {
	return &RootCertOptions{
		PemCert:      testconstants.PAACertNoVid,
		Info:         testconstants.Info,
		Subject:      testconstants.PAACertNoVidSubject,
		SubjectKeyID: testconstants.PAACertNoVidSubjectKeyID,
		Vid:          vid,
	}
}

func RootDaCertificate(address sdk.AccAddress) types.Certificate {
	return types.NewRootCertificate(
		testconstants.RootCertPem,
		testconstants.RootSubject,
		testconstants.RootSubjectAsText,
		testconstants.RootSubjectKeyID,
		testconstants.RootSerialNumber,
		address.String(),
		[]*types.Grant{},
		[]*types.Grant{},
		testconstants.Vid,
		testconstants.SchemaVersion,
	)
}

func RootDaCertificateWithVid(address sdk.AccAddress) types.Certificate {
	return types.NewRootCertificate(
		testconstants.RootCertWithVid,
		testconstants.RootCertWithVidSubject,
		testconstants.RootCertWithVidSubjectSubjectAsText,
		testconstants.RootCertWithVidSubjectKeyID,
		testconstants.RootCertWithVidSerialNumber,
		address.String(),
		[]*types.Grant{},
		[]*types.Grant{},
		testconstants.RootCertWithVidVid,
		testconstants.SchemaVersion,
	)
}

func RootDaCertificateWithNumericVid(address sdk.AccAddress) types.Certificate {
	return types.NewRootCertificate(
		testconstants.PAACertWithNumericVid,
		testconstants.PAACertWithNumericVidSubject,
		testconstants.PAACertWithNumericVidSubjectAsText,
		testconstants.PAACertWithNumericVidSubjectKeyID,
		testconstants.PAACertWithNumericVidSerialNumber,
		address.String(),
		[]*types.Grant{},
		[]*types.Grant{},
		testconstants.PAACertWithNumericVidVid,
		testconstants.SchemaVersion,
	)
}

func RootDaCertWithSameSubjectKeyID1(address sdk.AccAddress) types.Certificate {
	return types.NewRootCertificate(
		testconstants.PAACertWithSameSubjectID1,
		testconstants.PAACertWithSameSubjectID1Subject,
		testconstants.PAACertWithSameSubjectID1SubjectAsText,
		testconstants.PAACertWithSameSubjectIDSubjectKeyID,
		testconstants.PAACertWithSameSubjectSerialNumber,
		address.String(),
		[]*types.Grant{},
		[]*types.Grant{},
		testconstants.Vid,
		testconstants.SchemaVersion,
	)
}

func RootDaCertificateWithSameSubjectKeyID2(address sdk.AccAddress) types.Certificate {
	return types.NewRootCertificate(
		testconstants.PAACertWithSameSubjectID2,
		testconstants.PAACertWithSameSubjectID2Subject,
		testconstants.PAACertWithSameSubjectID1SubjectAsText,
		testconstants.PAACertWithSameSubjectIDSubjectKeyID,
		testconstants.PAACertWithSameSubject2SerialNumber,
		address.String(),
		[]*types.Grant{},
		[]*types.Grant{},
		testconstants.Vid,
		testconstants.SchemaVersion,
	)
}

func RootDaCertificateWithSameSubjectAndSKID1(address sdk.AccAddress) types.Certificate {
	return types.NewRootCertificate(
		testconstants.RootCertWithSameSubjectAndSKID1,
		testconstants.RootCertWithSameSubjectAndSKIDSubject,
		testconstants.RootCertWithSameSubjectAndSKIDSubjectAsText,
		testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		testconstants.RootCertWithSameSubjectAndSKID1SerialNumber,
		address.String(),
		[]*types.Grant{},
		[]*types.Grant{},
		testconstants.RootCertWithVidVid,
		testconstants.SchemaVersion,
	)
}

func RootDaCertificateWithSameSubjectAndSKID2(address sdk.AccAddress) types.Certificate {
	return types.NewRootCertificate(
		testconstants.RootCertWithSameSubjectAndSKID2,
		testconstants.RootCertWithSameSubjectAndSKIDSubject,
		testconstants.RootCertWithSameSubjectAndSKIDSubjectAsText,
		testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		testconstants.RootCertWithSameSubjectAndSKID2SerialNumber,
		address.String(),
		[]*types.Grant{},
		[]*types.Grant{},
		testconstants.RootCertWithVidVid,
		testconstants.SchemaVersion,
	)
}

func IntermediateDaCertificate(address sdk.AccAddress) types.Certificate {
	return types.NewNonRootCertificate(
		testconstants.IntermediateCertPem,
		testconstants.IntermediateSubject,
		testconstants.IntermediateSubjectAsText,
		testconstants.IntermediateSubjectKeyID,
		testconstants.IntermediateSerialNumber,
		testconstants.IntermediateIssuer,
		testconstants.IntermediateAuthorityKeyID,
		testconstants.RootSubject,
		testconstants.RootSubjectKeyID,
		address.String(),
		0,
		testconstants.SchemaVersion,
	)
}

func IntermediateDaCertificateWithNumericPidVid(address sdk.AccAddress) types.Certificate {
	return types.NewNonRootCertificate(
		testconstants.PAICertWithNumericPidVid,
		testconstants.PAICertWithNumericPidVidSubject,
		testconstants.PAICertWithNumericPidVidSubjectAsText,
		testconstants.PAICertWithNumericPidVidSubjectKeyID,
		testconstants.PAICertWithNumericPidVidSerialNumber,
		testconstants.PAACertWithNumericVidSubject,
		testconstants.PAACertWithNumericVidSubjectKeyID,
		testconstants.PAACertWithNumericVidSubject,
		testconstants.PAACertWithNumericVidSubjectKeyID,
		address.String(),
		0,
		testconstants.SchemaVersion,
	)
}

func IntermediateDaCertificateWithSameSubjectAndSKID1(address sdk.AccAddress) types.Certificate {
	return types.NewNonRootCertificate(
		testconstants.IntermediateWithSameSubjectAndSKID1,
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectAsText,
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		testconstants.IntermediateCertWithSameSubjectAndSKID1SerialNumber,
		testconstants.IntermediateCertWithSameSubjectIssuer,
		testconstants.IntermediateCertWithSameSubjectAuthorityKeyID,
		testconstants.IntermediateCertWithSameSubjectIssuer,
		testconstants.IntermediateCertWithSameSubjectAuthorityKeyID,
		address.String(),
		testconstants.RootCertWithVidVid,
		testconstants.SchemaVersion,
	)
}

func IntermediateDaCertificateWithSameSubjectAndSKID2(address sdk.AccAddress) types.Certificate {
	return types.NewNonRootCertificate(
		testconstants.IntermediateWithSameSubjectAndSKID2,
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectAsText,
		testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		testconstants.IntermediateCertWithSameSubjectAndSKID2SerialNumber,
		testconstants.IntermediateCertWithSameSubjectIssuer,
		testconstants.IntermediateCertWithSameSubjectAuthorityKeyID,
		testconstants.IntermediateCertWithSameSubjectIssuer,
		testconstants.IntermediateCertWithSameSubjectAuthorityKeyID,
		address.String(),
		testconstants.RootCertWithVidVid,
		testconstants.SchemaVersion,
	)
}

func LeafDaCertificateWithSameSubjectAndSKID(address sdk.AccAddress) types.Certificate {
	return types.NewNonRootCertificate(
		testconstants.LeafCertWithSameSubjectAndSKID,
		testconstants.LeafCertWithSameSubjectAndSKIDSubject,
		testconstants.LeafCertWithSameSubjectAndSKIDSubjectAsText,
		testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID,
		testconstants.LeafCertWithSameSubjectAndSKIDSerialNumber,
		testconstants.LeafCertWithSameSubjectIssuer,
		testconstants.LeafCertWithSameSubjectAuthorityKeyID,
		testconstants.IntermediateCertWithSameSubjectIssuer,
		testconstants.IntermediateCertWithSameSubjectAuthorityKeyID,
		address.String(),
		testconstants.RootCertWithVidVid,
		testconstants.SchemaVersion,
	)
}

func LeafCertificate(address sdk.AccAddress) types.Certificate {
	return types.NewNonRootCertificate(
		testconstants.LeafCertPem,
		testconstants.LeafSubject,
		testconstants.LeafSubjectAsText,
		testconstants.LeafSubjectKeyID,
		testconstants.LeafSerialNumber,
		testconstants.LeafIssuer,
		testconstants.LeafAuthorityKeyID,
		testconstants.IntermediateIssuer,
		testconstants.IntermediateAuthorityKeyID,
		address.String(),
		0,
		testconstants.SchemaVersion,
	)
}

// RootNocCertificate1 returns the on-ledger row for either an OperationalPKI
// RCAC or a VIDSignerPKI VVSC root, selected by certificateType. The two
// branches use structurally distinct PEMs because Matter R1.6 §6.5.12 mandates
// different profiles (RCAC: cA=TRUE / KU=keyCertSign+cRLSign; VVSC: cA=FALSE /
// KU=digitalSignature). Lets the table-driven OperationalPKI / VIDSignerPKI
// tests share the same fixture entry point.
func RootNocCertificate1(address sdk.AccAddress, certificateType types.CertificateType) types.Certificate {
	if certificateType == types.CertificateType_VIDSignerPKI {
		return RootVvscCertificate1(address)
	}

	return types.NewNocRootCertificate(
		testconstants.NocRootCert1,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectAsText,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1SerialNumber,
		address.String(),
		testconstants.Vid,
		testconstants.SchemaVersion,
		certificateType,
	)
}

func RootNocCertificate1Copy(address sdk.AccAddress, certificateType types.CertificateType) types.Certificate {
	if certificateType == types.CertificateType_VIDSignerPKI {
		return RootVvscCertificate1Copy(address)
	}

	return types.NewNocRootCertificate(
		testconstants.NocRootCert1Copy,
		testconstants.NocRootCert1CopySubject,
		testconstants.NocRootCert1CopySubjectAsText,
		testconstants.NocRootCert1CopySubjectKeyID,
		testconstants.NocRootCert1CopySerialNumber,
		address.String(),
		testconstants.Vid,
		testconstants.SchemaVersion,
		certificateType,
	)
}

func RootNocCertificate2(address sdk.AccAddress, certificateType types.CertificateType) types.Certificate {
	if certificateType == types.CertificateType_VIDSignerPKI {
		return RootVvscCertificate2(address)
	}

	return types.NewNocRootCertificate(
		testconstants.NocRootCert2,
		testconstants.NocRootCert2Subject,
		testconstants.NocRootCert2SubjectAsText,
		testconstants.NocRootCert2SubjectKeyID,
		testconstants.NocRootCert2SerialNumber,
		address.String(),
		testconstants.Vid,
		testconstants.SchemaVersion,
		certificateType,
	)
}

func IntermediateNocCertificate1(address sdk.AccAddress, certificateType types.CertificateType) types.Certificate {
	if certificateType == types.CertificateType_VIDSignerPKI {
		return IntermediateVvscCertificate1(address)
	}

	return types.NewNocCertificate(
		testconstants.NocCert1,
		testconstants.NocCert1Subject,
		testconstants.NocCert1SubjectAsText,
		testconstants.NocCert1SubjectKeyID,
		testconstants.NocCert1SerialNumber,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectKeyID,
		address.String(),
		testconstants.Vid,
		testconstants.SchemaVersion,
		certificateType,
	)
}

func IntermediateNocCertificate1Copy(address sdk.AccAddress, certificateType types.CertificateType) types.Certificate {
	if certificateType == types.CertificateType_VIDSignerPKI {
		return IntermediateVvscCertificate1Copy(address)
	}

	return types.NewNocCertificate(
		testconstants.NocCert1Copy,
		testconstants.NocCert1CopySubject,
		testconstants.NocCert1CopySubjectAsText,
		testconstants.NocCert1CopySubjectKeyID,
		testconstants.NocCert1CopySerialNumber,
		testconstants.NocRootCert1CopySubject,
		testconstants.NocRootCert1CopySubjectKeyID,
		testconstants.NocRootCert1CopySubject,
		testconstants.NocRootCert1CopySubjectKeyID,
		address.String(),
		testconstants.Vid,
		testconstants.SchemaVersion,
		certificateType,
	)
}

func IntermediateNocCertificate2(address sdk.AccAddress, certificateType types.CertificateType) types.Certificate {
	return types.NewNocCertificate(
		testconstants.NocCert2,
		testconstants.NocCert2Subject,
		testconstants.NocCert2SubjectAsText,
		testconstants.NocCert2SubjectKeyID,
		testconstants.NocCert2SerialNumber,
		testconstants.NocRootCert2Subject,
		testconstants.NocRootCert2SubjectKeyID,
		testconstants.NocRootCert2Subject,
		testconstants.NocRootCert2SubjectKeyID,
		address.String(),
		testconstants.Vid,
		testconstants.SchemaVersion,
		certificateType,
	)
}

func LeafNocCertificate1(address sdk.AccAddress, certificateType types.CertificateType) types.Certificate {
	if certificateType == types.CertificateType_VIDSignerPKI {
		return LeafVvscCertificate1(address)
	}

	return types.NewNocCertificate(
		testconstants.NocLeafCert1,
		testconstants.NocLeafCert1Subject,
		testconstants.NocLeafCert1SubjectAsText,
		testconstants.NocLeafCert1SubjectKeyID,
		testconstants.NocLeafCert1SerialNumber,
		testconstants.NocLeafCert1Issuer,
		testconstants.NocLeafCert1AuthorityKeyID,
		testconstants.NocRootCert2Subject,
		testconstants.NocRootCert2SubjectKeyID,
		address.String(),
		testconstants.Vid,
		testconstants.SchemaVersion,
		certificateType,
	)
}

func RootVvscCertificate1(address sdk.AccAddress) types.Certificate {
	return types.NewNocRootCertificate(
		testconstants.VvscRootCert1,
		testconstants.VvscRootCert1Subject,
		testconstants.VvscRootCert1SubjectAsText,
		testconstants.VvscRootCert1SubjectKeyID,
		testconstants.VvscRootCert1SerialNumber,
		address.String(),
		testconstants.Vid,
		testconstants.SchemaVersion,
		types.CertificateType_VIDSignerPKI,
	)
}

func RootVvscCertificate1Copy(address sdk.AccAddress) types.Certificate {
	return types.NewNocRootCertificate(
		testconstants.VvscRootCert1Copy,
		testconstants.VvscRootCert1CopySubject,
		testconstants.VvscRootCert1CopySubjectAsText,
		testconstants.VvscRootCert1CopySubjectKeyID,
		testconstants.VvscRootCert1CopySerialNumber,
		address.String(),
		testconstants.Vid,
		testconstants.SchemaVersion,
		types.CertificateType_VIDSignerPKI,
	)
}

func RootVvscCertificate2(address sdk.AccAddress) types.Certificate {
	return types.NewNocRootCertificate(
		testconstants.VvscRootCert2,
		testconstants.VvscRootCert2Subject,
		testconstants.VvscRootCert2SubjectAsText,
		testconstants.VvscRootCert2SubjectKeyID,
		testconstants.VvscRootCert2SerialNumber,
		address.String(),
		testconstants.Vid,
		testconstants.SchemaVersion,
		types.CertificateType_VIDSignerPKI,
	)
}

func IntermediateVvscCertificate1(address sdk.AccAddress) types.Certificate {
	return types.NewNocCertificate(
		testconstants.VvscIcaCert1,
		testconstants.VvscIcaCert1Subject,
		testconstants.VvscIcaCert1SubjectAsText,
		testconstants.VvscIcaCert1SubjectKeyID,
		testconstants.VvscIcaCert1SerialNumber,
		testconstants.VvscIcaCert1Issuer,
		testconstants.VvscIcaCert1AuthorityKeyID,
		testconstants.VvscRootCert1Subject,
		testconstants.VvscRootCert1SubjectKeyID,
		address.String(),
		testconstants.Vid,
		testconstants.SchemaVersion,
		types.CertificateType_VIDSignerPKI,
	)
}

func IntermediateVvscCertificate1Copy(address sdk.AccAddress) types.Certificate {
	return types.NewNocCertificate(
		testconstants.VvscIcaCert1Copy,
		testconstants.VvscIcaCert1CopySubject,
		testconstants.VvscIcaCert1CopySubjectAsText,
		testconstants.VvscIcaCert1CopySubjectKeyID,
		testconstants.VvscIcaCert1CopySerialNumber,
		testconstants.VvscIcaCert1CopyIssuer,
		testconstants.VvscIcaCert1CopyAuthorityKeyID,
		testconstants.VvscRootCert1Subject,
		testconstants.VvscRootCert1SubjectKeyID,
		address.String(),
		testconstants.Vid,
		testconstants.SchemaVersion,
		types.CertificateType_VIDSignerPKI,
	)
}

func LeafVvscCertificate1(address sdk.AccAddress) types.Certificate {
	return types.NewNocCertificate(
		testconstants.VvscLeafCert1,
		testconstants.VvscLeafCert1Subject,
		testconstants.VvscLeafCert1SubjectAsText,
		testconstants.VvscLeafCert1SubjectKeyID,
		testconstants.VvscLeafCert1SerialNumber,
		testconstants.VvscLeafCert1Issuer,
		testconstants.VvscLeafCert1AuthorityKeyID,
		testconstants.VvscRootCert1Subject,
		testconstants.VvscRootCert1SubjectKeyID,
		address.String(),
		testconstants.Vid,
		testconstants.SchemaVersion,
		types.CertificateType_VIDSignerPKI,
	)
}
