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

func IntermediateCertificateNoVid(address sdk.AccAddress) types.Certificate {
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
