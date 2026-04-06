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

func (d *ComplianceInfo) SetOptionalFields(optionalFields *OptionalFields) {
	if optionalFields.CertificationTypeVersion != "" {
		d.CertificationTypeVersion = optionalFields.CertificationTypeVersion
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

	if optionalFields.OSName != "" {
		d.OSName = optionalFields.OSName
	}

	if optionalFields.CertificationRoute != "" {
		d.CertificationRoute = optionalFields.CertificationRoute
	}

	if optionalFields.ProductType != "" {
		d.ProductType = optionalFields.ProductType
	}

	if optionalFields.Transport != "" {
		d.Transport = optionalFields.Transport
	}

	if optionalFields.ParentChild != "" {
		d.ParentChild = optionalFields.ParentChild
	}

	if optionalFields.Reason != "" {
		d.Reason = optionalFields.Reason
	}
}

type OptionalFields struct {
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
