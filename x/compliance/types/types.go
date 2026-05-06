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
