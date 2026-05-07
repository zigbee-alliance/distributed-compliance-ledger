package types

import (
	"regexp"
	"strings"
)

const (
	ZigbeeCertificationType string = "zigbee"
	MatterCertificationType string = "matter"
	AliroCertificationType  string = "aliro"
)

// List of Certification Types.
type CertificationTypes []string

var CertificationTypesList = CertificationTypes{ZigbeeCertificationType, MatterCertificationType, AliroCertificationType}

func IsValidCertificationType(certificationType string) bool {
	for _, i := range CertificationTypesList {
		if i == certificationType {
			return true
		}
	}

	return false
}

const (
	EndProductProgramType        string = "endProduct"
	SoftwareComponentProgramType string = "softwareComponent"
	CompliantPlatformProgramType string = "compliantPlatform"
)

// List of supported Program Types.
type ProgramTypes []string

var ProgramTypesList = ProgramTypes{EndProductProgramType, SoftwareComponentProgramType, CompliantPlatformProgramType}

func IsValidProgramType(programType string) bool {
	for _, i := range ProgramTypesList {
		if i == programType {
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

const (
	ParentPFCCertificationRoute  = "parent"
	ChildPFCCertificationRoute   = "child"
	DefaultPFCCertificationRoute = ""
)

// List of PFC Certification Routes.
type PFCCertificationRoutes []string

var PFCCertificationRouteList = PFCCertificationRoutes{ParentPFCCertificationRoute, ChildPFCCertificationRoute, DefaultPFCCertificationRoute}

func IsValidPFCCertificationRoute(certificationRoute string) bool {
	for _, i := range PFCCertificationRouteList {
		if i == certificationRoute {
			return true
		}
	}

	return false
}

const (
	CertificationRouteFullTested  = "fullTested"
	CertificationRouteSimilarity  = "similarity"
	CertificationRouteRapidRecert = "rapid-recert"
	CertificationRouteFastTrack   = "fastTrack"
	CertificationRouteCtp         = "ctp"
	CertificationRouteFamily      = "family"
	CertificationRoutePortfolio   = "portfolio"
)

// List of Certification Routes.
type CertificationRoutes []string

var CertificationRoutesList = CertificationRoutes{CertificationRouteFullTested, CertificationRouteSimilarity, CertificationRouteRapidRecert, CertificationRouteFastTrack, CertificationRouteCtp, CertificationRouteFamily, CertificationRoutePortfolio}

func IsValidCertificationRoute(certificationRoute string) bool {
	for _, i := range CertificationRoutesList {
		if i == certificationRoute {
			return true
		}
	}

	return false
}

var familyIDRegex = regexp.MustCompile(`^[a-zA-Z0-9]*$`)

func IsValidFamilyID(id string) bool {
	if id == "" {
		return false
	}

	return familyIDRegex.MatchString(id)
}

const (
	TransportThread    = "thread"
	TransportWifi      = "wi-fi"
	TransportEthernet  = "ethernet"
	TransportBluetooth = "bluetooth"
	TransportNFC       = "nfc"
)

// Transports List of supported transports.
type Transports []string

var TransportsList = Transports{TransportThread, TransportWifi, TransportEthernet, TransportBluetooth, TransportNFC}

// IsValidTransport reports whether transport is a comma-separated list of
// supported transport values.
func IsValidTransport(transport string) bool {
	seen := make(map[string]struct{})
	for _, item := range strings.Split(transport, ",") {
		if !isSupportedTransport(item) {
			return false
		}
		if _, ok := seen[item]; ok {
			return false
		}
		seen[item] = struct{}{}
	}

	return true
}

func isSupportedTransport(transport string) bool {
	for _, i := range TransportsList {
		if i == transport {
			return true
		}
	}

	return false
}

var supportedClusterRegex = regexp.MustCompile(`^0x[0-9a-fA-F]{1,4}$`)

// IsValidSupportedClusters reports whether supportedClusters is a comma-separated
// list of hexadecimal cluster IDs (e.g. "0x0003,0x0004,0x0006"). Each entry must
// be 0x followed by 1–4 hex digits.
func IsValidSupportedClusters(supportedClusters string) bool {
	if supportedClusters == "" {
		return true
	}

	seen := make(map[string]struct{})
	for _, item := range strings.Split(supportedClusters, ",") {
		if !supportedClusterRegex.MatchString(item) {
			return false
		}
		key := strings.ToLower(item)
		if _, ok := seen[key]; ok {
			return false
		}
		seen[key] = struct{}{}
	}

	return true
}
