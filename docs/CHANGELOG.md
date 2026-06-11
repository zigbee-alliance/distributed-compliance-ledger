## Unreleased

### PKI — Matter R1.5 certificate-validation alignment

* `MsgProposeAddX509RootCert`, `MsgAddX509Cert`, `MsgAddNocX509RootCert`, and `MsgAddNocX509IcaCert` now reject certificates that do not satisfy the Matter R1.5 structural profile for their type (§6.2.2.3 / §6.2.2.4 / §6.2.2.5 / §6.5.12). The two overloaded handlers (`MsgAddX509Cert` for PAI + DAC, `MsgAddNocX509IcaCert` for ICAC + NOC) dispatch by the BasicConstraints `cA` flag server-side — no proto / wire changes.
* Per-add-path rules now enforced: `version == v3`, ecdsa-with-SHA256 + prime256v1, BasicConstraints critical with the correct `cA`, KeyUsage critical with the profile-specific bits (CA: `keyCertSign + cRLSign` + optional `digitalSignature`; DAC / NOC: exactly `digitalSignature`), SKI required on every CA, AKI required on non-self-signed CAs, NOC requires critical EKU = `{serverAuth, clientAuth}` and RCAC / ICAC reject EKU, PAA rejects a ProductID attribute, every cert rejects duplicate matter-vid / matter-pid attributes, and `MsgAddX509Cert` enforces immediate-issuer VID / PID consistency in the DA chain.
* `VerifyCRLSignerCertFormat` BasicConstraints-criticality bug fixed — the criticality loop now tracks both BC (`2.5.29.19`) and KU (`2.5.29.15`) independently instead of silently testing only KU.
* `FormatOID` / `ToSubjectAsText` hardened: DER tag + length parsing for `#<hex>` values, accepts PrintableString and UTF8String, requires `oldKey=` exact match, and returns an error on malformed input rather than leaving the raw OID in place (which previously made `GetVidFromSubject` / `GetPidFromSubject` return 0 and silently bypass every VID / PID check).

### Compliance — schema v1 (#730)

* `MsgCertifyModel`, `MsgRevokeModel`, `MsgProvisionModel`, `MsgUpdateComplianceInfo`: `schemaVersion` must now be **`1`** (was `0`).
* `MsgCertifyModel`, `MsgProvisionModel`, `MsgUpdateComplianceInfo`: new required **`specificationVersion`** field (moved out of `ModelVersion`).
* Field-size tightening: `certificationType` max 20 (was 100); `reason` max 20480 (was 102400); `cDCertificateId` is now exactly **19 chars** (was max 64); `supportedClusters` max 256 (was 64).
* Deprecated on every compliance write path: `compliantPlatformUsed`, `compliantPlatformVersion`, `OSVersion`, `certificationIdOfSoftwareComponent`. Stored entries keep their existing values; new writes are advised to leave these empty.
* `ComplianceInfo` storage: adds `specificationVersion`; the same four fields above are marked deprecated.

### Model

* `MsgCreateModel`: `productLabel` and `partNumber` are now **required** (were optional).
* `MsgCreateModel.discoveryCapabilitiesBitmask` valid range widened to **0–30** (was 0–14).
* `MsgCreateModelVersion.otaChecksum` max 88 chars (was 64), `otaChecksumType` valid range tightened to **0–12** (was 0–65535).
* `MsgCreateModelVersion.specificationVersion` deprecated — set it via the Compliance module (`certify-model` / `provision-model` / `update-compliance-info`).
* All URL fields across `model` (and `vendorinfo`, `pki`) switched to the unified `https_url` validator. Behaviour is unchanged for users (https-only, ≤256 chars).

### PKI

* All four cert-add messages (`MsgProposeAddX509RootCert`, `MsgAddX509Cert`, `MsgAddNocX509RootCert`, `MsgAddNocX509IcaCert`) cap the inbound PEM at **20 KiB** (was ~10 MiB).
* `MsgAddPkiRevocationDistributionPoint` / `MsgUpdatePkiRevocationDistributionPoint`: explicit caps — `label` ≤ 64, `crlSignerCertificate` ≤ 2 KiB, `crlSignerDelegator` ≤ 2 KiB, `issuerSubjectKeyID` ≤ 64, `dataURL` ≤ 256 (now also accepts `http://` in addition to `https://`), `dataDigest` ≤ 128.

### Other

* `x509.CertificatePEMsEqual` switched to DER-bytes equality so the CRL revocation path tolerates whitespace / PEM-header differences.
* `x/dclupgrade`: reject-upgrade now cancels the scheduled upgrade once the reject quorum is reached.
* `URL_LIVENESS_CHECK_ENABLED` disabled in `scripts/tests-after-upgrade/{local,testnet}.env` so the upgrade-test lane is not gated on external URL liveness.

### Notes

* All in-repo test fixture chains (DA, NOC, RootCertWithVid) were regenerated with consistent keys so every CA satisfies the new strict profile and every leaf chains under its new parent.
* All certificates currently in `certs-mainnet.json` and `certs-testnet.json` were verified to pass every new `Verify…` helper (`scripts/audit_real_certs`). No on-chain cert is retroactively rejected by this change.

## [v1.4.0-dev3](https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/tag/v1.4.0-dev3) - 2024-04-24

### Dependency Changes
* Upgrade cosmos-sdk to v0.47.8
  * Migrate to CometBFT. Follow the migration instructions in the [upgrade guide](https://github.com/cosmos/cosmos-sdk/blob/main/UPGRADING.md#migration-to-cometbft-part-1).
  * Upgrade all related dependencies too 
* Upgrade Golang version to 1.20

### Breaking Changes
* Transaction broadcasting `block` mode has been removed from the updated cosmos-sdk. 
Starting from this version, dcl has only two modes: `sync` and `async`, with the default being `sync`.
In this mode, to obtain the actual result of a transaction (`txn`), an additional `query` call with the `txHash` must be executed. For example:
    `dcld query tx txHash` - where `txHash` represents the hash of the previously executed transaction."
* `starport` cli tool is no longer supported. Please use `v0.27.2` version of [ignite](https://github.com/ignite/cli).
* Due to upgrading `cosmovisor` to v1.3.0 in Docker and shell files, the node starting command has changed from `cosmovisor start` to `cosmovisor run start`
