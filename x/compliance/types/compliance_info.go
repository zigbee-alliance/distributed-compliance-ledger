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
}

type OptionalFields struct {
	ProgramTypeVersion                 string
	FamilyID                           string
	SupportedClusters                  string
	CompliantPlatformUsed              string
	CompliantPlatformVersion           string
	OSVersion                          string
	CertificationRoute                 string
	ProgramType                        string
	Transport                          string
	ParentChild                        string
	CertificationIDOfSoftwareComponent string
}
