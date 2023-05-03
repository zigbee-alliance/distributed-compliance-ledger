package types

func (m *DeviceSoftwareCompliance) IsComplianceInfoExist(
	vid int32, pid int32, softwareVersion uint32,
) (int, bool) {
	for index, info := range m.ComplianceInfo {
		if info.Vid == vid && info.Pid == pid && info.SoftwareVersion == softwareVersion {
			return index, true
		}
	}

	return -1, false
}

func (m *DeviceSoftwareCompliance) RemoveComplianceInfo(removeComplianceInfoIndex int) {
	m.ComplianceInfo = append(m.ComplianceInfo[:removeComplianceInfoIndex], m.ComplianceInfo[removeComplianceInfoIndex+1:]...)
}
