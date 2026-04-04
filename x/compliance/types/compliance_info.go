package types

func (d *ComplianceInfo) SetCertifiedStatus(date string, reason string, cdCertificateID string) {
	svCertificationStatus := CodeCertified
	historyItem := ComplianceHistoryItem{
		SoftwareVersionCertificationStatus: d.SoftwareVersionCertificationStatus,
		Date:                               d.Date,
		Reason:                             d.Reason,
		CDVersionNumber:                    d.CDVersionNumber,
	}
	d.History = append(d.History, &historyItem)
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
	d.History = append(d.History, &historyItem)
	d.SoftwareVersionCertificationStatus = svCertificationStatus
	d.Date = date
	d.Reason = reason
}

type ComplianceOptionalFields struct {
	CertificationTypeVersion string
	FamilyID                 string
	SupportedClusters        string
	CompliantPlatformUsed    string
	CompliantPlatformVersion string
	OSName                   string
	CertificationRoute       string
	ProductType              string
	Transport                string
	ParentChild              string
	Reason                   string
}

func (d *ComplianceInfo) SetOptionalFields(optionalFields *ComplianceOptionalFields) {
	if len(optionalFields.CertificationTypeVersion) > 0 {
		d.CertificationTypeVersion = optionalFields.CertificationTypeVersion
	}
	if len(optionalFields.FamilyID) > 0 {
		d.FamilyId = optionalFields.FamilyID
	}
	if len(optionalFields.SupportedClusters) > 0 {
		d.SupportedClusters = optionalFields.SupportedClusters
	}
	if len(optionalFields.CompliantPlatformUsed) > 0 {
		d.CompliantPlatformUsed = optionalFields.CompliantPlatformUsed
	}
	if len(optionalFields.CompliantPlatformVersion) > 0 {
		d.CompliantPlatformVersion = optionalFields.CompliantPlatformVersion
	}
	if len(optionalFields.OSName) > 0 {
		d.OSName = optionalFields.OSName
	}
	if len(optionalFields.CertificationRoute) > 0 {
		d.CertificationRoute = optionalFields.CertificationRoute
	}
	if len(optionalFields.ProductType) > 0 {
		d.ProductType = optionalFields.ProductType
	}
	if len(optionalFields.Transport) > 0 {
		d.Transport = optionalFields.Transport
	}
	if len(optionalFields.ParentChild) > 0 {
		d.ParentChild = optionalFields.ParentChild
	}
	if len(optionalFields.Reason) > 0 {
		d.Reason = optionalFields.Reason
	}
}
