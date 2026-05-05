package types

const MaxComplianceHistoryItem = 20

func (d *ComplianceInfo) SetCertifiedStatus(date string, reason string, cdCertificateID string) {
	svCertificationStatus := CodeCertified
	historyItem := ComplianceHistoryItem{
		SoftwareVersionCertificationStatus: d.SoftwareVersionCertificationStatus,
		Date:                               d.Date,
		Reason:                             d.Reason,
		CDVersionNumber:                    d.CDVersionNumber,
	}
	d.updateHistory(&historyItem)
	d.SoftwareVersionCertificationStatus = svCertificationStatus
	d.Date = date
	d.Reason = reason
	d.CDCertificateId = cdCertificateID
}

func (d *ComplianceInfo) SetRevokedStatus(date string, reason string) {
	svCertificationStatus := CodeRevoked
	historyItem := ComplianceHistoryItem{
		SoftwareVersionCertificationStatus: d.SoftwareVersionCertificationStatus,
		Date:                               d.Date,
		Reason:                             d.Reason,
		CDVersionNumber:                    d.CDVersionNumber,
	}
	d.updateHistory(&historyItem)
	d.SoftwareVersionCertificationStatus = svCertificationStatus
	d.Date = date
	d.Reason = reason
}

func (d *ComplianceInfo) updateHistory(item *ComplianceHistoryItem) {
	d.History = append(d.History, item)

	if len(d.History) > MaxComplianceHistoryItem {
		// TODO Can be changed to another better logic/way to update history
		d.History = d.History[len(d.History)-MaxComplianceHistoryItem:]
	}
}

func (d *ComplianceInfo) SetOptionalFields(optionalFields *OptionalFields) {
	if optionalFields.ProgramTypeVersion != "" {
		d.ProgramTypeVersion = optionalFields.ProgramTypeVersion
	}

	if optionalFields.FamilyID != "" {
		d.FamilyId = optionalFields.FamilyID
	}

	if optionalFields.SupportedClusters != "" {
		d.SupportedClusters = optionalFields.SupportedClusters
	}

	if optionalFields.CompliantPlatformUsed != "" {
		d.CompliantPlatformUsed = optionalFields.CompliantPlatformUsed
	}

	if optionalFields.CompliantPlatformVersion != "" {
		d.CompliantPlatformVersion = optionalFields.CompliantPlatformVersion
	}

	if optionalFields.OSVersion != "" {
		d.OSVersion = optionalFields.OSVersion
	}

	if optionalFields.CertificationRoute != "" {
		d.CertificationRoute = optionalFields.CertificationRoute
	}

	if optionalFields.ProgramType != "" {
		d.ProgramType = optionalFields.ProgramType
	}

	if optionalFields.Transport != "" {
		d.Transport = optionalFields.Transport
	}

	if optionalFields.ParentChild != "" {
		d.ParentChild = optionalFields.ParentChild
	}

	if optionalFields.CertificationIDOfSoftwareComponent != "" {
		d.CertificationIdOfSoftwareComponent = optionalFields.CertificationIDOfSoftwareComponent
	}

	if optionalFields.Reason != "" {
		d.Reason = optionalFields.Reason
	}
}

type OptionalFields struct {
	FamilyID                           string
	SupportedClusters                  string
	CompliantPlatformUsed              string
	CompliantPlatformVersion           string
	CertificationIDOfSoftwareComponent string
	OSVersion                          string
	CertificationRoute                 string
	ProgramType                        string
	ProgramTypeVersion                 string
	Transport                          string
	ParentChild                        string
	Reason                             string
}
