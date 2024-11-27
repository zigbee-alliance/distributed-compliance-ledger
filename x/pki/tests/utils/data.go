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

func CreateTestRootCert() TestCertificate {
	return TestCertificate{
		Subject:        testconstants.RootSubject,
		SubjectKeyID:   testconstants.RootSubjectKeyID,
		SerialNumber:   testconstants.RootSerialNumber,
		Issuer:         testconstants.RootIssuer,
		AuthorityKeyID: testconstants.RootSubjectKeyID,
		IsRoot:         true,
	}
}

func CreateTestRootCertWithSameSubject() TestCertificate {
	return TestCertificate{
		Subject:      testconstants.PAACertWithSameSubjectID1Subject,
		SubjectKeyID: testconstants.PAACertWithSameSubjectIDSubjectID,
		SerialNumber: testconstants.PAACertWithSameSubjectSerialNumber,
		Issuer:       testconstants.PAACertWithSameSubjectIssuer,
		IsRoot:       true,
	}
}

func CreateTestRootCertWithSameSubject2() TestCertificate {
	return TestCertificate{
		Subject:      testconstants.PAACertWithSameSubjectID2Subject,
		SubjectKeyID: testconstants.PAACertWithSameSubjectIDSubjectID,
		SerialNumber: testconstants.PAACertWithSameSubject2SerialNumber,
		Issuer:       testconstants.PAACertWithSameSubject2Issuer,
		IsRoot:       true,
	}
}

func CreateTestRootCertWithSameSubjectAndSkid1() TestCertificate {
	return TestCertificate{
		Subject:        testconstants.RootCertWithSameSubjectAndSKIDSubject,
		SubjectKeyID:   testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		SerialNumber:   testconstants.RootCertWithSameSubjectAndSKID1SerialNumber,
		Issuer:         testconstants.RootCertWithSameSubjectAndSKID1Issuer,
		AuthorityKeyID: testconstants.RootCertWithSameSubjectAndSKIDSubject,
		IsRoot:         true,
	}
}

func CreateTestRootCertWithSameSubjectAndSkid2() TestCertificate {
	return TestCertificate{
		Subject:        testconstants.RootCertWithSameSubjectAndSKIDSubject,
		SubjectKeyID:   testconstants.RootCertWithSameSubjectAndSKIDSubjectKeyID,
		SerialNumber:   testconstants.RootCertWithSameSubjectAndSKID2SerialNumber,
		Issuer:         testconstants.RootCertWithSameSubjectAndSKID2Issuer,
		AuthorityKeyID: testconstants.RootCertWithSameSubjectAndSKIDSubject,
		IsRoot:         true,
	}
}

func CreateTestIntermediateCert() TestCertificate {
	return TestCertificate{
		Subject:        testconstants.IntermediateSubject,
		SubjectKeyID:   testconstants.IntermediateSubjectKeyID,
		SerialNumber:   testconstants.IntermediateSerialNumber,
		Issuer:         testconstants.IntermediateIssuer,
		AuthorityKeyID: testconstants.IntermediateAuthorityKeyID,
		IsRoot:         false,
	}
}

func CreateTestIntermediateVidScopedCert() TestCertificate {
	return TestCertificate{
		Subject:        testconstants.PAICertWithNumericPidVidSubject,
		SubjectKeyID:   testconstants.PAICertWithNumericPidVidSubjectKeyID,
		SerialNumber:   testconstants.PAICertWithNumericPidVidSerialNumber,
		Issuer:         testconstants.PAACertWithNumericVidSubject,
		AuthorityKeyID: testconstants.PAACertWithNumericVidSubjectKeyID,
		IsRoot:         false,
	}
}

func CreateTestIntermediateCertWithSameSubjectAndSKID1() TestCertificate {
	return TestCertificate{
		Subject:        testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		SubjectKeyID:   testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		SerialNumber:   testconstants.IntermediateCertWithSameSubjectAndSKID1SerialNumber,
		Issuer:         testconstants.IntermediateCertWithSameSubjectIssuer,
		AuthorityKeyID: testconstants.IntermediateCertWithSameSubjectAuthorityKeyID,
		IsRoot:         false,
	}
}

func CreateTestIntermediateCertWithSameSubjectAndSKID2() TestCertificate {
	return TestCertificate{
		Subject:        testconstants.IntermediateCertWithSameSubjectAndSKIDSubject,
		SubjectKeyID:   testconstants.IntermediateCertWithSameSubjectAndSKIDSubjectKeyID,
		SerialNumber:   testconstants.IntermediateCertWithSameSubjectAndSKID2SerialNumber,
		Issuer:         testconstants.IntermediateCertWithSameSubjectIssuer,
		AuthorityKeyID: testconstants.IntermediateCertWithSameSubjectAuthorityKeyID,
		IsRoot:         false,
	}
}

func CreateTestLeafCertWithSameSubjectAndSKID() TestCertificate {
	return TestCertificate{
		Subject:        testconstants.LeafCertWithSameSubjectAndSKIDSubject,
		SubjectKeyID:   testconstants.LeafCertWithSameSubjectAndSKIDSubjectKeyID,
		SerialNumber:   testconstants.LeafCertWithSameSubjectAndSKIDSerialNumber,
		Issuer:         testconstants.LeafCertWithSameSubjectIssuer,
		AuthorityKeyID: testconstants.LeafCertWithSameSubjectAuthorityKeyID,
		IsRoot:         false,
	}
}

func CreateTestLeafCert() TestCertificate {
	return TestCertificate{
		Subject:        testconstants.LeafSubject,
		SubjectKeyID:   testconstants.LeafSubjectKeyID,
		SerialNumber:   testconstants.LeafSerialNumber,
		Issuer:         testconstants.LeafIssuer,
		AuthorityKeyID: testconstants.LeafAuthorityKeyID,
		IsRoot:         false,
	}
}

func CreateTestNocRoot1Cert() TestCertificate {
	return TestCertificate{
		Subject:        testconstants.NocRootCert1Subject,
		SubjectKeyID:   testconstants.NocRootCert1SubjectKeyID,
		SerialNumber:   testconstants.NocRootCert1SerialNumber,
		Issuer:         testconstants.NocRootCert1Issuer,
		AuthorityKeyID: testconstants.NocRootCert1SubjectKeyID,
		VID:            testconstants.Vid,
		IsRoot:         true,
	}
}

func CreateTestNocRoot2Cert() TestCertificate {
	return TestCertificate{
		Subject:        testconstants.NocRootCert1CopySubject,
		SubjectKeyID:   testconstants.NocRootCert1CopySubjectKeyID,
		SerialNumber:   testconstants.NocRootCert1CopySerialNumber,
		Issuer:         testconstants.NocRootCert1CopyIssuer,
		AuthorityKeyID: testconstants.NocRootCert1CopySubjectKeyID,
		VID:            testconstants.Vid,
		IsRoot:         true,
	}
}

func CreateTestNocIca1Cert() TestCertificate {
	return TestCertificate{
		Subject:        testconstants.NocCert1Subject,
		SubjectKeyID:   testconstants.NocCert1SubjectKeyID,
		SerialNumber:   testconstants.NocCert1SerialNumber,
		Issuer:         testconstants.NocCert1Issuer,
		AuthorityKeyID: testconstants.NocCert1AuthorityKeyID,
		VID:            testconstants.Vid,
		IsRoot:         false,
	}
}

func CreateTestNocIca1CertCopy() TestCertificate {
	return TestCertificate{
		Subject:        testconstants.NocCert1CopySubject,
		SubjectKeyID:   testconstants.NocCert1CopySubjectKeyID,
		SerialNumber:   testconstants.NocCert1CopySerialNumber,
		Issuer:         testconstants.NocCert1CopyIssuer,
		AuthorityKeyID: testconstants.NocCert1CopyAuthorityKeyID,
		VID:            testconstants.Vid,
		IsRoot:         false,
	}
}

func CreateTestNocLeafCert() TestCertificate {
	return TestCertificate{
		Subject:        testconstants.NocLeafCert1Subject,
		SubjectKeyID:   testconstants.NocLeafCert1SubjectKeyID,
		SerialNumber:   testconstants.NocLeafCert1SerialNumber,
		Issuer:         testconstants.NocLeafCert1Issuer,
		AuthorityKeyID: testconstants.NocLeafCert1AuthorityKeyID,
		VID:            testconstants.Vid,
		IsRoot:         false,
	}
}
