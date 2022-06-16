package types

func (d *ComplianceInfo) SetCertifiedStatus(date string, reason string) {
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

func (complianceInfo *ComplianceInfo) SetOptionalFields(msg *MsgCertifyModel) {
	if msg.ProgramTypeVersion != "" {
		complianceInfo.ProgramTypeVersion = msg.ProgramTypeVersion
	}

	if msg.FamilyId != "" {
		complianceInfo.FamilyId = msg.FamilyId
	}

	if msg.SupportedClusters != "" {
		complianceInfo.SupportedClusters = msg.SupportedClusters
	}

	if msg.CompliantPlatformUsed != "" {
		complianceInfo.CompliantPlatformUsed = msg.CompliantPlatformUsed
	}

	if msg.CompliantPlatformVersion != "" {
		complianceInfo.CompliantPlatformVersion = msg.CompliantPlatformVersion
	}

	if msg.OSVersion != "" {
		complianceInfo.OSVersion = msg.OSVersion
	}

	if msg.CertificationRoute != "" {
		complianceInfo.CertificationRoute = msg.CertificationRoute
	}

	if msg.CertificationRoute != "" {
		complianceInfo.CertificationRoute = msg.CertificationRoute
	}

	if msg.ProgramType != "" {
		complianceInfo.ProgramType = msg.ProgramType
	}

	if msg.Transport != "" {
		complianceInfo.Transport = msg.Transport
	}

	if msg.ParentChild != "" {
		complianceInfo.ParentChild = msg.ParentChild
	}
}
