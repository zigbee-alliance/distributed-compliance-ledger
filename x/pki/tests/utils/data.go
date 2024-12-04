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

func RootCertificate(address sdk.AccAddress) types.Certificate {
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

func RootCertWithVid(address sdk.AccAddress) types.Certificate {
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

func PAACertWithNumericVid(address sdk.AccAddress) types.Certificate {
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

func PAACertWithSameSubjectID1(address sdk.AccAddress) types.Certificate {
	return types.NewRootCertificate(
		testconstants.PAACertWithSameSubjectID1,
		testconstants.PAACertWithSameSubjectID1Subject,
		testconstants.PAACertWithSameSubjectID1SubjectAsText,
		testconstants.PAACertWithSameSubjectIDSubjectID,
		testconstants.PAACertWithSameSubjectSerialNumber,
		address.String(),
		[]*types.Grant{},
		[]*types.Grant{},
		testconstants.Vid,
		testconstants.SchemaVersion,
	)
}

func PAACertWithSameSubjectID2(address sdk.AccAddress) types.Certificate {
	return types.NewRootCertificate(
		testconstants.PAACertWithSameSubjectID2,
		testconstants.PAACertWithSameSubjectID2Subject,
		testconstants.PAACertWithSameSubjectID1SubjectAsText,
		testconstants.PAACertWithSameSubjectIDSubjectID,
		testconstants.PAACertWithSameSubject2SerialNumber,
		address.String(),
		[]*types.Grant{},
		[]*types.Grant{},
		testconstants.Vid,
		testconstants.SchemaVersion,
	)
}

func RootCertWithSameSubjectAndSKID1(address sdk.AccAddress) types.Certificate {
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

func RootCertWithSameSubjectAndSKID2(address sdk.AccAddress) types.Certificate {
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

func IntermediateCertPem(address sdk.AccAddress) types.Certificate {
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

func PAICertWithNumericPidVid(address sdk.AccAddress) types.Certificate {
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

func IntermediateWithSameSubjectAndSKID1(address sdk.AccAddress) types.Certificate {
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

func IntermediateWithSameSubjectAndSKID2(address sdk.AccAddress) types.Certificate {
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

func LeafCertWithSameSubjectAndSKID(address sdk.AccAddress) types.Certificate {
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

func LeafCertPem(address sdk.AccAddress) types.Certificate {
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

func NocRootCert1(address sdk.AccAddress) types.Certificate {
	return types.NewNocRootCertificate(
		testconstants.NocRootCert1,
		testconstants.NocRootCert1Subject,
		testconstants.NocRootCert1SubjectAsText,
		testconstants.NocRootCert1SubjectKeyID,
		testconstants.NocRootCert1SerialNumber,
		address.String(),
		testconstants.Vid,
		testconstants.SchemaVersion,
	)
}

func NocRootCert1Copy(address sdk.AccAddress) types.Certificate {
	return types.NewNocRootCertificate(
		testconstants.NocRootCert1Copy,
		testconstants.NocRootCert1CopySubject,
		testconstants.NocRootCert1CopySubjectAsText,
		testconstants.NocRootCert1CopySubjectKeyID,
		testconstants.NocRootCert1CopySerialNumber,
		address.String(),
		testconstants.Vid,
		testconstants.SchemaVersion,
	)
}

func NocRootCert2(address sdk.AccAddress) types.Certificate {
	return types.NewNocRootCertificate(
		testconstants.NocRootCert2,
		testconstants.NocRootCert2Subject,
		testconstants.NocRootCert2SubjectAsText,
		testconstants.NocRootCert2SubjectKeyID,
		testconstants.NocRootCert2SerialNumber,
		address.String(),
		testconstants.Vid,
		testconstants.SchemaVersion,
	)
}

func NocCertIca1(address sdk.AccAddress) types.Certificate {
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
	)
}

func NocCert1Copy(address sdk.AccAddress) types.Certificate {
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
	)
}

func NocCert2(address sdk.AccAddress) types.Certificate {
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
	)
}

func NocLeafCert1(address sdk.AccAddress) types.Certificate {
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
	)
}
