# Compliance

<!-- markdownlint-disable MD036 -->

### CERTIFY_MODEL

**Status: Implemented**

Attests compliance of the Model Version to the ZB or Matter standard.

`REVOKE_MODEL_CERTIFICATION` should be used for revoking (disabling) the compliance.
It's possible to call `CERTIFY_MODEL` for revoked model versions to enable them back.

The corresponding Model and Model Version must be present on the ledger.

It must be called for every compliant device for use cases where compliance
is tracked on ledger.

It can be used for use cases where only revocation is tracked on the ledger to remove a Model Version
from the revocation list.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string
  - certificationDate: `string` - The date of model certification (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z
  - certificationType: `string` - Certification type - Currently 'zigbee', 'matter', 'access control', 'product security' types are supported
  - cdCertificateId: `string` - CD Certificate ID 
  - reason `optional(string)` - optional comment describing the reason of the certification
  - cDVersionNumber `optional(uint32)` - optional field describing the CD version number
  - familyId `optional(string)` - optional field describing the family ID
  - supportedClusters `optional(string)` - optional field describing the supported clusters
  - compliantPlatformUsed `optional(string)` - optional field describing the compliant platform used
  - compliantPlatformVersion `optional(string)` - optional field describing the compliant platform version
  - OSVersion `optional(string)` - optional field describing the OS version
  - certificationRoute `optional(string)` - optional field describing the certification route
  - programType `optional(string)` - optional field describing the program type
  - programTypeVersion `optional(string)` - optional field describing the program type version
  - transport `optional(string)` - optional field describing the transport
  - parentChild `optional(string)` - optional field describing the parent/child - Currently 'parent' and 'child' types are supported
  - certificationIDOfSoftwareComponent `optional(string)` - optional field describing the certification ID of software component
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State:
  - `compliance/ComplianceInfo/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
  - `compliance/CertifiedModel/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
- Who can send:
  - CertificationCenter
- CLI command:
  - `dcld tx compliance certify-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string>  --certificationType=<matter|zigbee|access control|product security> --certificationDate=<rfc3339 encoded date> --cdCertificateId=<string> --from=<account>`
- CLI command full:
  - `dcld tx compliance certify-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string>  --certificationType=<matter|zigbee|access control|product security> --certificationDate=<rfc3339 encoded date> --cdCertificateId=<string> --reason=<string> --cDVersionNumber=<uint32> --familyId=<string> --supportedClusters=<string> --compliantPlatformUsed=<string> --compliantPlatformVersion=<string> --OSVersion=<string> --certificationRoute=<string> --programType=<string> --programTypeVersion=<string> --transport=<string> --parentChild=<string> --certificationIDOfSoftwareComponent=<string> --from=<account>`

### UPDATE_COMPLIANCE_INFO

**Status: Implemented**

Updates a compliance info by VID, PID, Software Version and Certification Type.


- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - certificationType: `string` - Certification type - Currently 'zigbee', 'matter', 'access control', 'product security' types are supported
  - certificationDate: `optional(string)` - The date of model certification (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z
  - cdCertificateId: `optional(string)` - CD Certificate ID
  - reason `optional(string)` - optional comment describing the reason of the certification
  - cDVersionNumber `optional(string)` - optional field (a uint32-parsable string) describing the CD version number, must be the same with the associated model version
  - familyId `optional(string)` - optional field describing the family ID
  - supportedClusters `optional(string)` - optional field describing the supported clusters
  - compliantPlatformUsed `optional(string)` - optional field describing the compliant platform used
  - compliantPlatformVersion `optional(string)` - optional field describing the compliant platform version
  - OSVersion `optional(string)` - optional field describing the OS version
  - certificationRoute `optional(string)` - optional field describing the certification route
  - programType `optional(string)` - optional field describing the program type
  - programTypeVersion `optional(string)` - optional field describing the program type version
  - transport `optional(string)` - optional field describing the transport
  - parentChild `optional(string)` - optional field describing the parent/child - Currently 'parent' and 'child' types are supported
  - certificationIDOfSoftwareComponent `optional(string)` - optional field describing the certification ID of software component
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- Who can send:
  - CertificationCenter
- CLI command:
  - `dcld tx compliance update-compliance-info`
- CLI command full:
  - `dcld tx compliance update-compliance-info --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<string> --cdVersionNumber=<string> --certificationDate=$upd_certification_date --reason=$upd_reason --cdCertificateId=$upd_cd_certificate_id --certificationRoute=$upd_certification_route --programType=$upd_program_type --programTypeVersion=$upd_program_type_version --compliantPlatformUsed=$upd_compliant_platform_used --compliantPlatformVersion=$upd_compliant_platform_version --transport=$upd_transport --familyId=$upd_familyID --supportedClusters=$upd_supported_clusters --OSVersion=$upd_os_version --parentChild=$upd_parent_child --certificationIDOfSoftwareComponent=$upd_certification_id_of_software_component --from=$zb_account`
- REST API:
  - `/dcl/compliance/update-compliance-info`

### DELETE_COMPLIANCE_INFO

**Status: Implemented**

Delete compliance of the Model Version to the ZB or Matter standard.

The corresponding Compliance Info is required to be present on the ledger

- Parameters:
  - vid: `uint16` - model vendor ID (positive non-zero)
  - pid: `uint16` - model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - certificationType: `string` - Certification type - Currently 'zigbee' and 'matter', 'access control', 'product security' types are supported
- Who can send:
  - CertificationCenter
- CLI command:
  - `dcld tx compliance delete-compliance-info --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee|access control|product security> --from=<account>`

### REVOKE_MODEL_CERTIFICATION

**Status: Implemented**

Revoke compliance of the Model Version to the ZB or Matter standard.

The corresponding Model and Model Version are not required to be present on the ledger.

It can be used in cases where every compliance result
is written on the ledger (`CERTIFY_MODEL` was called), or
 cases where only revocation list is stored on the ledger.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string
  - revocationDate: `string` - The date of model revocation (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z
  - certificationType: `string`  - Certification type - Currently 'zigbee' and 'matter', 'access control', 'product security' types are supported
  - reason `optional(string)`  - optional comment describing the reason of revocation
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State:
  - `compliance/ComplianceInfo/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
  - `compliance/RevokedModel/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
- Who can send:
  - CertificationCenter
- CLI command:
  - `dcld tx compliance revoke-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee|access control|product security> --revocationDate=<rfc3339 encoded date> --reason=<string> --from=<account>`

### PROVISION_MODEL

**Status: Implemented**

Sets provisional state for the Model Version.

The corresponding Model and Model Version is required to be present in the ledger.

Can not be set if there is already a certification record on the ledger (certified or revoked).

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string
  - provisionalDate: `string` - The date of model provisioning (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z
  - certificationType: `string`  - Certification type - Currently 'zigbee' and 'matter', 'access control', 'product security' types are supported
  - cdCertificateId: `string` - CD Certificate ID 
  - reason `optional(string)`  - optional comment describing the reason of revocation
  - cDVersionNumber `optional(uint32)` - optional field describing the CD version number
  - familyId `optional(string)` - optional field describing the family ID
  - supportedClusters `optional(string)` - optional field describing the supported clusters
  - compliantPlatformUsed `optional(string)` - optional field describing the compliant platform used
  - compliantPlatformVersion `optional(string)` - optional field describing the compliant platform version
  - OSVersion `optional(string)` - optional field describing the OS version
  - certificationRoute `optional(string)` - optional field describing the certification route
  - programType `optional(string)` - optional field describing the program type
  - programTypeVersion `optional(string)` - optional field describing the program type version
  - transport `optional(string)` - optional field describing the transport
  - parentChild `optional(string)` - optional field describing the parent/child - Currently 'parent' and 'child' types are supported
  - certificationIDOfSoftwareComponent `optional(string)` - optional field describing the certification ID of software component
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State:
  - `compliance/ComplianceInfo/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
  - `compliance/ProvisionalModel/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
- Who can send:
  - CertificationCenter
- CLI command:
  - `dcld tx compliance provision-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee|access control|product security> --provisionalDate=<rfc3339 encoded date> --from=<account>`
- CLI command full:
  - `dcld tx compliance provision-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee|access control|product security> --provisionalDate=<rfc3339 encoded date> --cdCertificateId=<string> --reason=<string> --cDVersionNumber=<uint32> --familyId=<string> --supportedClusters=<string> --compliantPlatformUsed=<string> --compliantPlatformVersion=<string> --OSVersion=<string> --certificationRoute=<string> --programType=<string> --programTypeVersion=<string> --transport=<string> --parentChild=<string> --certificationIDOfSoftwareComponent=<string> --from=<account>`

### GET_CERTIFIED_MODEL

**Status: Implemented**

Gets a structure containing the Model Version / Certification Type key (`vid`, `pid`, `softwareVersion`, `certificationType`) and a flag (`value`) indicating whether the given Model Version is compliant to `certificationType` standard.

This is the aggregation of compliance and
revocation information for every vid/pid/softwareVersion/certificationType. It should be used in cases where compliance
is tracked on the ledger.

This function responds with `NotFound` (404 code) if Model Version was never certified earlier.

This function returns `true` if compliance information is found on ledger and it's in `certified` state.

You can use `GET_COMPLICE_INFO` method to get the whole compliance information.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - certification_type: `string`  - Certification type - Currently 'zigbee' and 'matter' types are supported
- CLI command:
  - `dcld query compliance certified-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter|access control|product security>`
- REST API:
  - GET `/dcl/compliance/certified-models/{vid}/{pid}/{software_version}/{certification_type}`

### GET_REVOKED_MODEL

**Status: Implemented**

Gets a structure containing the Model Version / Certification Type key (`vid`, `pid`, `softwareVersion`, `certificationType`) and a flag (`value`) indicating whether the given Model Version is revoked for `certificationType` standard.

It contains information about revocation only, so it should be used in cases
 where only revocation is tracked on the ledger.

This function responds with `NotFound` (404 code) if Model Version was never certified or revoked earlier.

This function returns `true` if compliance information is found on ledger and it's in `revoked` state.

You can use `GET_COMPLICE_INFO` method to get the whole compliance information.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - certification_type: `string`  - Certification type - Currently 'zigbee' and 'matter' types are supported
- CLI command:
  - `dcld query compliance revoked-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter|access control|product security>`
- REST API:
  - GET `/dcl/compliance/revoked-models/{vid}/{pid}/{software_version}/{certification_type}`

### GET_PROVISIONAL_MODEL

**Status: Implemented**

Gets a structure containing the Model Version / Certification Type key (`vid`, `pid`, `softwareVersion`, `certificationType`) and a flag (`value`) indicating whether the given Model Version is in provisional state for `certificationType` standard.

This function responds with `NotFound` (404 code) if Model Version was never provisioned or certified earlier.

This function returns `true` if compliance information is found on the ledger and it's in `provisional` state.

You can use `GET_COMPLICE_INFO` method to get the whole compliance information.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - certification_type: `string`  - Certification type - Currently 'zigbee' and 'matter' types are supported
- CLI command:
  - `dcld query compliance provisional-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter|access control|product security>`
- REST API:
  - GET `/dcl/compliance/provisional-models/{vid}/{pid}/{software_version}/{certification_type}`

### GET_COMPLIANCE_INFO

**Status: Implemented**

Gets compliance information associated with the Model Version and Certification Type (identified by the `vid`, `pid`, `softwareVersion` and `certification_type`).

It can be used instead of GET_CERTIFIED_MODEL / GET_REVOKED_MODEL / GET_PROVISIONAL_MODEL methods
to get the whole compliance information without additional state check.

This function responds with `NotFound` (404 code) if compliance information is not found in store.

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - certification_type: `string`  - Certification type - Currently 'zigbee' and 'matter' types are supported
- CLI command:
  - `dcld query compliance compliance-info --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter|access control|product security>`
- REST API:
  - GET `/dcl/compliance/compliance-info/{vid}/{pid}/{software_version}/{certification_type}`

### GET_DEVICE_SOFTWARE_COMPLIANCE

**Status: Implemented**

Gets device software compliance associated with the `cDCertificateId`.

This function responds with `NotFound` (404 code) if device software compliance is not found in store.

- Parameters:
  - cDCertificateId: `string` - CD Certificate ID
- CLI command:
  - `dcld query compliance device-software-compliance --cDCertificateId=<string>`
- REST API:
  - GET `/dcl/compliance/device-software-compliance/{cDCertificateId}`

### GET_ALL_CERTIFIED_MODELS

**Status: Implemented**

Gets all compliant Model Versions for all vendors (`vid`s).

This is the aggregation of compliance and
revocation information for every vid/pid. It should be used in cases where compliance is tracked on ledger.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query compliance all-certified-models`
- REST API:
  - GET `/dcl/compliance/certified-models`

### GET_ALL_REVOKED_MODELS

**Status: Implemented**

Gets all revoked Model Versions for all vendors (`vid`s).

It contains information about revocation only, so it should be used in cases
 where only revocation is tracked on the ledger.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query compliance all-revoked-models`
- REST API:
  - GET `/dcl/compliance/revoked-models`

### GET_ALL_PROVISIONAL_MODELS

**Status: Implemented**

Gets all Model Versions in provisional state for all vendors (`vid`s).

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query compliance all-provisional-models`
- REST API:
  - GET `/dcl/compliance/provisional-models`

### GET_ALL_COMPLIANCE_INFO_RECORDS

**Status: Implemented**

Gets all stored compliance information records for all vendors (`vid`s).

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query compliance all-compliance-info`
- REST API:
  - GET `/dcl/compliance/compliance-info`

### GET_ALL_DEVICE_SOFTWARE_COMPLIANCES

**Status: Implemented**

Gets all stored device software compliance's.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query compliance all-device-software-compliance`
- REST API:
  - `/dcl/compliance/device-software-compliance`

<!-- markdownlint-enable MD036 -->


<!-- markdownlint-enable MD036 -->