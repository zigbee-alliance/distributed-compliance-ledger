# Compliance

<!-- markdownlint-disable MD036 -->

### CERTIFY_MODEL

**Status: Implemented**

Attests compliance of the Model Version to the ZB or Matter standard.

`REVOKE_MODEL_CERTIFICATION` should be used for revoking (disabling) the compliance.
It's possible to call `CERTIFY_MODEL` for revoked model versions to enable them back.

The corresponding Model and Model Version are not required to be present in the ledger. It can be added later by Vendors.

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
  - certificationType: `string` - Certification program applied to the model. Currently, 'zigbee' or 'matter' types are supported
  - cdCertificateId: `string` - Connectivity Standards Alliance certification's certificate ID applied to the model certification
  - reason `optional(string)` - optional comment describing the reason of the certification
  - cDVersionNumber `optional(uint32)` - optional field describing the CD version number
  - familyId `optional(string)` - optional field describing the product family to which the certified model belongs. The possible value should start with the prefix `FAM` and be followed by an alphanumeric character (e.g. `FAM123456`).
  - supportedClusters `optional(string)` - optional field describing the cluster IDs supported by the application. Supported cluster IDs are `0x0003`, `0x0004`, `0x0006`, `0x0062`, `0x0008`, and `0x0406`.
  - compliantPlatformUsed `optional(string)` - optional field describing the Certification ID of the compliant platform used with the product.
  - compliantPlatformVersion `optional(string)` - optional field describing the compliant platform version
  - OSName `optional(string)` - optional field describing the operating system name and version running on the device at the time of certification
  - certificationRoute `optional(string)` - optional field describing the various certification paths, such as Fully Tested, Certification by Similarity, Family/Portfolio Certification, Certification Transfer etc. Supported values are  `fullTested`, `similarity`, `rapid-recert`, `fastTrack`, `ctp`, `family`, and `portfolio`
  - productType `optional(string)` - optional field describing the product type. Supported values are `endProduct`, `softwareComponent` or `compliantPlatform`.
- certificationTypeVersion `optional(string)` - optional field describing the version of the certification program (see `certificationType` for supported programs)
  - transport `optional(string)` - optional field describing the underlying communication technology the device uses to connect and exchange data. Supported values are `thread`, `wi-fi`, `ethernet`, and `bluetooth`
  - parentChild `optional(string)` - optional field describing the parent vs. child characteristic when using the Product Family Certification or Portfolio Certification Program. Supported values are `parent` and `child`.
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State:
  - `compliance/ComplianceInfo/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
  - `compliance/CertifiedModel/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
- Who can send:
  - CertificationCenter
- CLI command:
  - `dcld tx compliance certify-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string>  --certificationType=<matter|zigbee> --certificationDate=<rfc3339 encoded date> --cdCertificateId=<string> --from=<account>`
- CLI command full:
  - `dcld tx compliance certify-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --softwareVersionString=<string>  --certificationType=<matter|zigbee> --certificationDate=<rfc3339 encoded date> --cdCertificateId=<string> --reason=<string> --cDVersionNumber=<uint32> --familyId=<string> --supportedClusters=<string> --compliantPlatformUsed=<string> --compliantPlatformVersion=<string> --OSName=<string> --certificationRoute=<string> --productType=<string> --certificationTypeVersion=<string> --transport=<string> --parentChild=<string> --from=<account>`

### UPDATE_COMPLIANCE_INFO

**Status: Implemented**

Updates a compliance info by VID, PID, Software Version and Certification Type.


- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - certificationType: `string` - certification program applied to the model. Currently, 'zigbee' or 'matter' types are supported
  - certificationDate: `optional(string)` - the date of model certification (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z. If not provided, the existing value will be preserved.
  - cdCertificateId: `optional(string)` - Connectivity Standards Alliance certification's certificate ID applied to the model certification. If not provided, the existing value will be preserved.
  - reason `optional(string)` - optional comment describing the reason of the certification
  - cDVersionNumber `optional(string)` - optional field (a uint32-parsable string) describing the CD version number, must be the same with the associated model version
  - familyId `optional(string)` - optional field describing the product family to which the certified model belongs. The possible value should start with the prefix `FAM` and be followed by an alphanumeric character (e.g. `FAM123456`).
  - supportedClusters `optional(string)` - optional field describing the cluster IDs supported by the application. Supported cluster IDs are `0x0003`, `0x0004`, `0x0006`, `0x0062`, `0x0008`, and `0x0406`.
  - compliantPlatformUsed `optional(string)` - optional field describing the Certification ID of the compliant platform used with the product.
  - compliantPlatformVersion `optional(string)` - optional field describing the compliant platform version
  - OSName `optional(string)` - optional field describing the operating system name and version running on the device at the time of certification
  - certificationRoute `optional(string)` - optional field describing the certification route
  - productType `optional(string)` - optional field describing the product type. Supported values are `endProduct`, `softwareComponent` or `compliantPlatform`.
  - certificationTypeVersion `optional(string)` - optional field describing the version of the certification program (see `certificationType` for supported programs)
  - transport `optional(string)` - optional field describing the underlying communication technology the device uses to connect and exchange data. Supported values are `thread`, `wi-fi`, `ethernet`, and `bluetooth`
  - parentChild `optional(string)` - optional field describing the parent vs. child characteristic when using the Product Family Certification or Portfolio Certification Program. Supported values are `parent` and `child`.
  - schemaVersion: `optional(uint16)` - schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- Who can send:
  - CertificationCenter
- CLI command:
  - `dcld tx compliance update-compliance-info`
- CLI command full:
  - `dcld tx compliance update-compliance-info --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<string> --cdVersionNumber=<string> --certificationDate=$upd_certification_date --reason=$upd_reason --cdCertificateId=$upd_cd_certificate_id --certificationRoute=$upd_certification_route --productType=$upd_program_type --certificationTypeVersion=$upd_program_type_version --compliantPlatformUsed=$upd_compliant_platform_used --compliantPlatformVersion=$upd_compliant_platform_version --transport=$upd_transport --familyId=$upd_familyID --supportedClusters=$upd_supported_clusters --OSName=$upd_os_version --parentChild=$upd_parent_child --from=$zb_account`
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
  - certificationType: `string` - Certification program applied to the model. Currently, 'zigbee' or 'matter' types are supported
- Who can send:
  - CertificationCenter
- CLI command:
  - `dcld tx compliance delete-compliance-info --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee> --from=<account>`

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
  - certificationType: `string`  - Certification program applied to the model. Currently, 'zigbee' or 'matter' types are supported
  - reason `optional(string)`  - optional comment describing the reason of revocation
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State:
  - `compliance/ComplianceInfo/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
  - `compliance/RevokedModel/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
- Who can send:
  - CertificationCenter
- CLI command:
  - `dcld tx compliance revoke-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee> --revocationDate=<rfc3339 encoded date> --reason=<string> --from=<account>`

### PROVISION_MODEL

**Status: Implemented**

Sets provisional state for the Model Version.

The corresponding Model and Model Version are not required to be present in the ledger. It can be added later by Vendors.

Can not be set if there is already a certification record on the ledger (certified or revoked).

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
  - pid: `uint16` -  model product ID (positive non-zero)
  - softwareVersion: `uint32` - model software version
  - softwareVersionSting: `string` - model software version string
  - provisionalDate: `string` - the date of model provisioning (rfc3339 encoded), for example 2019-10-12T07:20:50.52Z
  - certificationType: `string`  - certification program applied to the model. Currently, 'zigbee' or 'matter' types are supported
  - cdCertificateId: `string` - Connectivity Standards Alliance certification's certificate ID applied to the model certification
  - reason `optional(string)`  - optional comment describing the reason of revocation
  - cDVersionNumber `optional(uint32)` - optional field describing the CD version number
  - familyId `optional(string)` - optional field describing the product family to which the certified model belongs. The possible value should start with the prefix `FAM` and be followed by an alphanumeric character (e.g. `FAM123456`).
  - supportedClusters `optional(string)` - optional field describing the cluster IDs supported by the application. Supported cluster IDs are `0x0003`, `0x0004`, `0x0006`, `0x0062`, `0x0008`, and `0x0406`.
  - compliantPlatformUsed `optional(string)` - optional field describing the Certification ID of the compliant platform used with the product.
  - compliantPlatformVersion `optional(string)` - optional field describing the compliant platform 
  - OSName `optional(string)` - optional field describing the operating system name and version running on the device at the time of certification
  - certificationRoute `optional(string)` - optional field describing the certification route
  - productType `optional(string)` - optional field describing the product type. Supported values are `endProduct`, `softwareComponent` or `compliantPlatform`.
  - certificationTypeVersion `optional(string)` - optional field describing the version of the certification program (see `certificationType` for supported programs)
  - transport `optional(string)` - optional field describing the underlying communication technology the device uses to connect and exchange data. Supported values are `thread`, `wi-fi`, `ethernet`, and `bluetooth`
  - parentChild `optional(string)` - optional field describing the parent vs. child characteristic when using the Product Family Certification or Portfolio Certification Program. Supported values are `parent` and `child`.
  - schemaVersion: `optional(uint16)` - schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State:
  - `compliance/ComplianceInfo/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
  - `compliance/ProvisionalModel/value/<vid>/<pid>/<softwareVersion>/<certificationType>`
- Who can send:
  - CertificationCenter
- CLI command:
  - `dcld tx compliance provision-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee> --provisionalDate=<rfc3339 encoded date> --from=<account>`
- CLI command full:
  - `dcld tx compliance provision-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee> --provisionalDate=<rfc3339 encoded date> --cdCertificateId=<string> --reason=<string> --cDVersionNumber=<uint32> --familyId=<string> --supportedClusters=<string> --compliantPlatformUsed=<string> --compliantPlatformVersion=<string> --OSName=<string> --certificationRoute=<string> --productType=<string> --certificationTypeVersion=<string> --transport=<string> --parentChild=<string> --from=<account>`

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
  - certification_type: `string`  - Certification program applied to the model. Currently, 'zigbee' or 'matter' types are supported
- CLI command:
  - `dcld query compliance certified-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter>`
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
  - certification_type: `string`  - Certification program applied to the model. Currently, 'zigbee' or 'matter' types are supported
- CLI command:
  - `dcld query compliance revoked-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter>`
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
  - certification_type: `string`  - Certification program applied to the model. Currently, 'zigbee' or 'matter' types are supported
- CLI command:
  - `dcld query compliance provisional-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter>`
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
  - certification_type: `string`  - Certification program applied to the model. Currently, 'zigbee' or 'matter' types are supported
- CLI command:
  - `dcld query compliance compliance-info --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter>`
- REST API:
  - GET `/dcl/compliance/compliance-info/{vid}/{pid}/{software_version}/{certification_type}`

### GET_DEVICE_SOFTWARE_COMPLIANCE

**Status: Implemented**

Gets device software compliance associated with the `cDCertificateId`.

This function responds with `NotFound` (404 code) if device software compliance is not found in store.

- Parameters:
  - cDCertificateId: `string` - Connectivity Standards Alliance certification's certificate ID applied to the model certification
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
  - GET `/dcl/compliance/device-software-compliance`

<!-- markdownlint-enable MD036 -->


<!-- markdownlint-enable MD036 -->