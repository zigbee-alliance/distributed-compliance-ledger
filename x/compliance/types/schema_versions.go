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

import commontypes "github.com/zigbee-alliance/distributed-compliance-ledger/x/common/types"

// Current schema version for each persisted compliance type. Bump the matching
// constant whenever a proto field is added, removed, or renamed
const (
	ComplianceInfoSchemaVersion           uint32 = 1
	CertifiedModelSchemaVersion           uint32 = 0
	ProvisionalModelSchemaVersion         uint32 = 0
	RevokedModelSchemaVersion             uint32 = 0
	DeviceSoftwareComplianceSchemaVersion uint32 = 0
	ComplianceHistoryItemSchemaVersion    uint32 = 0
)

// Compile-time assertions that each state type satisfies the SchemaVersioned interface.
var (
	_ commontypes.SchemaVersioned = (*ComplianceInfo)(nil)
	_ commontypes.SchemaVersioned = (*CertifiedModel)(nil)
	_ commontypes.SchemaVersioned = (*ProvisionalModel)(nil)
	_ commontypes.SchemaVersioned = (*RevokedModel)(nil)
	_ commontypes.SchemaVersioned = (*DeviceSoftwareCompliance)(nil)
	_ commontypes.SchemaVersioned = (*ComplianceHistoryItem)(nil)
)

func (*ComplianceInfo) CurrentSchemaVersion() uint32 { return ComplianceInfoSchemaVersion }
func (d *ComplianceInfo) SetSchemaVersion(v uint32)  { d.SchemaVersion = v }

func (*CertifiedModel) CurrentSchemaVersion() uint32 { return CertifiedModelSchemaVersion }
func (m *CertifiedModel) SetSchemaVersion(v uint32)  { m.SchemaVersion = v }

func (*ProvisionalModel) CurrentSchemaVersion() uint32 { return ProvisionalModelSchemaVersion }
func (m *ProvisionalModel) SetSchemaVersion(v uint32)  { m.SchemaVersion = v }

func (*RevokedModel) CurrentSchemaVersion() uint32 { return RevokedModelSchemaVersion }
func (m *RevokedModel) SetSchemaVersion(v uint32)  { m.SchemaVersion = v }

func (*DeviceSoftwareCompliance) CurrentSchemaVersion() uint32 {
	return DeviceSoftwareComplianceSchemaVersion
}
func (m *DeviceSoftwareCompliance) SetSchemaVersion(v uint32) { m.SchemaVersion = v }

func (*ComplianceHistoryItem) CurrentSchemaVersion() uint32 {
	return ComplianceHistoryItemSchemaVersion
}
func (m *ComplianceHistoryItem) SetSchemaVersion(v uint32) { m.SchemaVersion = v }
