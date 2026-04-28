# Mapping DCL Entities to Matter Specification

This document describes how DCL records and entities map to the Matter [specification](https://github.com/CHIP-Specifications/connectedhomeip-spec/blob/master/src/service_device_management/DistributedComplianceLedger.adoc#ref_DistributedComplianceLedger).

In DCL, different schemas and respective endpoints are used for write (txn messages) and read (query calls) requests, but the Matter specification describes only what was written and assumes that read requests are the same. To know about how write and read requests map to spec, please follow the sections below for more details.

## PKI Module

1. DCL uses a single [Certificate](../proto/zigbeealliance/distributedcomplianceledger/pki/certificate.proto) entity for all(`PAA`,`PAI`,`RCAC` and `ICAC`) certificate types mentioned in the Matter specification.
2. Associated `write/read` requests regarding the Product Attestation Authority (PAA) and Intermediate (PAI) certificate [schemas](https://github.com/CHIP-Specifications/connectedhomeip-spec/blob/master/src/service_device_management/DistributedComplianceLedger.adoc#ref_PAAAndPAICertificateSchema) can be found in the [DA certificate types section](transactions.md#x509-pki).
3. Associated `write/read` requests regarding the Operational Certificates [schemas](https://github.com/CHIP-Specifications/connectedhomeip-spec/blob/master/src/service_device_management/DistributedComplianceLedger.adoc#5-operational-trust-anchors-schema) can be found in the [NOC certificate types section](transactions.md#x509-pki).
    *   **Note:** DCL uses the terms **NOC** and **ICA**, which map to **RCAC** and **ICAC** in the Matter specification, respectively.
4. Device Attestation PKI Revocation Distribution Points [schema](https://github.com/CHIP-Specifications/connectedhomeip-spec/blob/master/src/service_device_management/DistributedComplianceLedger.adoc#9-device-attestation-pki-revocation-distribution-points-schema) correspond to [PKI Revocation Distribution Point](../proto/zigbeealliance/distributedcomplianceledger/pki/pki_revocation_distribution_point.proto) entity and the associated `write/read` requests can be found in the [revocation points section](transactions.md#x509-pki).

## Model Module

In DCL, the [Model module](transactions/model.md) is responsible to handle records regarding device and device software version models.

1. The [Device Model schema](https://github.com/CHIP-Specifications/connectedhomeip-spec/blob/master/src/service_device_management/DistributedComplianceLedger.adoc#ref_DeviceModelSchema) corresponds to the [Model](../proto/zigbeealliance/distributedcomplianceledger/model/model.proto) entity, and the associated `write/read` requests can be found in the [Model and Model Version section](transactions.md#model-and-model-version).
2. The [Device Software Version Model schema](https://github.com/CHIP-Specifications/connectedhomeip-spec/blob/master/src/service_device_management/DistributedComplianceLedger.adoc#7-devicesoftwareversionmodel-schema) corresponds to the to [Model Version](../proto/zigbeealliance/distributedcomplianceledger/model/model_version.proto) entity, and the associated `write/read` requests can be found in the [Model and Model Version section](transactions.md#model-and-model-version).

## Compliance Module

In DCL, the [Compliance module](transactions/compliance.md) is responsible to handle certification status of particular software version(`Model Version` in DCL term) for given product(`Model` in DCL term).

Below is the list of notes to consider while mapping DCL Compliance module to Matter specification:
1. The [Device Software Compliance schema](https://github.com/CHIP-Specifications/connectedhomeip-spec/blob/master/src/service_device_management/DistributedComplianceLedger.adoc#8-devicesoftwarecompliance--compliance-test-result-schema) corresponds to the [Compliance Info](../proto/zigbeealliance/distributedcomplianceledger/compliance/compliance_info.proto) entity, and the associated `write/read` requests can be found in the [Compliance section](transactions.md#compliance)
2. In DCL, specific endpoints are used (by CSA) to handle [certification status](https://github.com/CHIP-Specifications/connectedhomeip-spec/blob/master/src/service_device_management/DistributedComplianceLedger.adoc#82-softwareversioncertificationstatusenum-type) of the device software version according to the provided `certificatType`(ZB/Matter):
   - [Provision](transactions/compliance.md#PROVISION_MODEL) endpoint is used to register that particular device software is in `provisional` (going into certification testing phase) state
      * **Note:** This endpoint **cannot** be used for device software versions that are already in `compliant` or `revoked` state
   - [Certify](transactions/compliance.md#CERTIFY_MODEL) endpoint is used to register that particular device software version is in `compliant` state
      * **Note:** This endpoint **can** be used for device software versions that are already in `provisional` or `revoked` state (e.g. re-certification)
   - [Revoke](transactions/compliance.md#REVOKE_MODEL_CERTIFICATION) endpoint is used to register that particular device software version is `revoked`
     * **Note:** This endpoint **can** be used for device software versions that are already in `provisional` or `compliant` state
   - [Update](transactions/compliance.md#UPDATE_COMPLIANCE_INFO) endpoint is used to update the additional-info/metadata of a particular device software version
     * **Note:** This endpoint **cannot** be used to change the certification status
3. In DCL, there are several query endpoints can be used for specific reading purposes.
   - [Provisional Model](transactions/compliance.md#GET_PROVISIONAL_MODEL) can be used to retrieve/check the software version certification of particular device is in `provisional` state
   - [Certified Model](transactions/compliance.md#GET_CERTIFIED_MODEL) can be used to retrieve/check the software version certification of particular device in `compliant` state
   - [Revoked Model](transactions/compliance.md#GET_REVOKED_MODEL) can be used to retrieve/check the software version certification of particular device in `revoked` state
   - [Compliance Info](transactions/compliance.md#GET_COMPLIANCE_INFO) can be used to get the full compliance(including certification state) information of a particular device software version, bypassing state check query calls mentioned above
   - [Device Software Compliance](transactions/compliance.md#GET_DEVICE_SOFTWARE_COMPLIANCE) can be used to get all compliance records associated with [CD Certificate ID](https://github.com/CHIP-Specifications/connectedhomeip-spec/blob/master/src/service_device_management/DistributedComplianceLedger.adoc#83-cdcertificateid)
