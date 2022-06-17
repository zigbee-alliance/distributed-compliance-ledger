package types

func (d *ComplianceInfo) SetCertifiedStatus(date string, reason string, cdCertificationId string) {
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
	d.CDCertificationId = cdCertificationId
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

func (d *ComplianceInfo) SetOptionalFields(msg *MsgCertifyModel) {
	if msg.ProgramTypeVersion != "" {
		d.ProgramTypeVersion = msg.ProgramTypeVersion
	}

	if msg.FamilyId != "" {
		d.FamilyId = msg.FamilyId
	}

	if msg.SupportedClusters != "" {
		d.SupportedClusters = msg.SupportedClusters
	}

	if msg.CompliantPlatformUsed != "" {
		d.CompliantPlatformUsed = msg.CompliantPlatformUsed
	}

	if msg.CompliantPlatformVersion != "" {
		d.CompliantPlatformVersion = msg.CompliantPlatformVersion
	}

	if msg.OSVersion != "" {
		d.OSVersion = msg.OSVersion
	}

	if msg.CertificationRoute != "" {
		d.CertificationRoute = msg.CertificationRoute
	}

	if msg.CertificationRoute != "" {
		d.CertificationRoute = msg.CertificationRoute
	}

	if msg.ProgramType != "" {
		d.ProgramType = msg.ProgramType
	}

	if msg.Transport != "" {
		d.Transport = msg.Transport
	}

	if msg.ParentChild != "" {
		d.ParentChild = msg.ParentChild
	}
}
