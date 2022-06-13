package types

const (
	ZigbeeCertificationType string = "zigbee"
	MatterCertificationType string = "matter"
	FullCertificationType   string = "Full"
	CbSCertificationType    string = "CbS" // CbS - Certification by Similarity
	CTPCertificationType    string = "CTP" // CTP - Certification Transfer Program
	PFCCertificationType    string = "PFC" // PFC - Product Family Certification
)

//	List of Certification Types
type CertificationTypes []string

var CertificationTypesList = CertificationTypes{
	ZigbeeCertificationType, MatterCertificationType, FullCertificationType,
	CbSCertificationType, CTPCertificationType, PFCCertificationType,
}

func IsValidCertificationType(certificationType string) bool {
	for _, i := range CertificationTypesList {
		if i == certificationType {
			return true
		}
	}

	return false
}

const (
	CodeProvisional uint32 = 1
	CodeCertified   uint32 = 2
	CodeRevoked     uint32 = 3
)
