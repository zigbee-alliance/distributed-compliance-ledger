// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
