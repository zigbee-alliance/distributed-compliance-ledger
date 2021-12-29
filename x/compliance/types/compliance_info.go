package types

func (d *ComplianceInfo) UpdateComplianceInfo(date string, reason string) {
	// Toggle state
	var svCertificationStatus uint32
	if d.SoftwareVersionCertificationStatus == CodeCertified {
		svCertificationStatus = CodeRevoked
	} else {
		svCertificationStatus = CodeCertified
	}

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
