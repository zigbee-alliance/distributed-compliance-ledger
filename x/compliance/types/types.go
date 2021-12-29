package types

const (
	ZigbeeCertificationType string = "zigbee"
	MatterCertificationType string = "matter"
)

//	List of Certification Types
type CertificationTypes []string

var CertificationTypesList = CertificationTypes{ZigbeeCertificationType, MatterCertificationType}

func IsValidCertificationType(certificationType string) bool {
	for _, i := range CertificationTypesList {
		if i == certificationType {
			return true
		}
	}

	return false
}

//nolint:godot
const (
	CodeProvisional uint32 = 1
	CodeCertified   uint32 = 2
	CodeRevoked     uint32 = 3
)
