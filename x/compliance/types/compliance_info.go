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
