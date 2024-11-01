# Upgrade Module

<!-- markdownlint-disable MD036 -->

### PROPOSE_UPGRADE

**Status: Implemented**

Proposes an upgrade plan with the given name at the given height.

- Parameters:
  - name: `string` - upgrade plan name
  - upgrade-height: `int64` -  upgrade plan height (positive non-zero)
  - upgrade-info: `optional(string)` - upgrade plan info (for node admins to
      read). Recommended format is an os/architecture -> application binary URL
      map as a JSON under `binaries` key where each URL should include the
      corresponding checksum as `checksum` query parameter with the value in the
      format `type:value` where `type` is `sha256` or `sha512` and `value` is
      the actual checksum value. For example:

```json
{
  "binaries": {
    "linux/amd64":"https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/v0.7.0/dcld?checksum=sha256:aec070645fe53ee3b3763059376134f058cc337247c978add178b6ccdfb0019f"
  }
}
```

- In State: `dclupgrade/ProposedUpgrade/value/<name>`
- Who can send:
  - Trustee
- Number of required approvals:
  - greater than 2/3 of Trustees (proposal by a Trustee is also counted as an approval)
- CLI command minimal:

```bash
dcld tx dclupgrade propose-upgrade --name=<string> --upgrade-height=<int64> --from=<account>
```

- CLI command full:

```bash
dcld tx dclupgrade propose-upgrade --name=<string> --upgrade-height=<int64> --upgrade-info=<string> --from=<account>
```

> **_Note:_**  If the current upgrade proposal is out of date(when the current network height is greater than the proposed upgrade height), we can resubmit the upgrade proposal with the same name.

### APPROVE_UPGRADE

**Status: Implemented**

Approves the proposed upgrade plan with the given name. It also can be used for revote (i.e. change vote from reject to approve)

- Parameters:
  - name: `string` - upgrade plan name
- In State: `upgrade/0x0`
- Who can send:
  - Trustee
- Number of required approvals:
  - greater than 2/3 of Trustees (proposal by a Trustee is also counted as an approval)
- CLI command:

```bash
dcld tx dclupgrade approve-upgrade --name=<string> --from=<account>
```

### REJECT_UPGRADE

**Status: Implemented**

Rejects the proposed upgrade plan with the given name. It also can be used for revote (i.e. change vote from approve to reject)

If proposed upgrade has only proposer's approval and no rejects then proposer can send this transaction to remove the proposal

- Paramaters:
  - name: `string` - upgrade plan name
- In State: `RejectUpgrade/value/<name>`
- Who can send:
  - Trustee
- Number of required rejects:
  - more than 1/3 of Trustees
- CLI command:

```bash
dcld tx dclupgrade reject-upgrade --name=<string> --from=<account>
```

### GET_PROPOSED_UPGRADE

**Status: Implemented**

Gets the proposed upgrade plan with the given name.

- Parameters:
  - name: `string` - upgrade plan name
- CLI command:

```bash
dcld query dclupgrade proposed-upgrade --name=<string>
```

- REST API:
  - GET `/dcl/dclupgrade/proposed-upgrades/{name}`

### GET_APPROVED_UPGRADE

**Status: Implemented**

Gets the approved upgrade plan with the given name.

- Parameters:
  - name: `string` - upgrade plan name
- CLI command:

```bash
dcld query dclupgrade approved-upgrade --name=<string>
```

- REST API:
  - GET `/dcl/dclupgrade/approved-upgrades/{name}`

### GET_REJECTED_UPGRADE

**Status: Implemented**

Gets the rejected upgrade plan with the given name.

- Parameters:
  - name: `string` - upgrade plan name
- CLI command:

```bash
dcld query dclupgrade rejected-upgrade --name=<string>
```

- REST API:
  - GET `/dcl/dclupgrade/rejected-upgrades/{name}`

### GET_ALL_PROPOSED_UPGRADES

**Status: Implemented**

Gets all the proposed upgrade plans.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:

```bash
dcld query dclupgrade all-proposed-upgrades
```

- REST API:
  - GET `/dcl/dclupgrade/proposed-upgrades`

### GET_ALL_APPROVED_UPGRADES

**Status: Implemented**

Gets all the approved upgrade plans.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:

```bash
dcld query dclupgrade all-approved-upgrades
```

- REST API:
  - GET `/dcl/dclupgrade/approved-upgrades`

### GET_ALL_REJECTED_UPGRADES

**Status: Implemented**

Gets all the rejected upgrade plans.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:

```bash
dcld query dclupgrade all-rejected-upgrades
```

- REST API:
  - GET `/dcl/dclupgrade/rejected-upgrades`

### GET_UPGRADE_PLAN

**Status: Implemented**

Gets the currently scheduled upgrade plan, if it exists.

- CLI command:

```bash
dcld query upgrade plan
```

- REST API:
  - GET `/cosmos/upgrade/v1beta1/current_plan`

### GET_APPLIED_UPGRADE

**Status: Implemented**

Returns the header for the block at which the upgrade with the given name was
applied, if it was previously executed on the chain. This helps a client
determine which binary was valid over a given range of blocks, as well as gives
more context to understand past migrations.

- Parameters:
  - `string` - upgrade name
- CLI command:

```bash
dcld query upgrade applied <string>
```

- REST API:
  - GET `/cosmos/upgrade/v1beta1/applied_plan/{name}`

### GET_MODULE_VERSIONS

**Status: Implemented**

Gets a list of module names and their respective consensus versions. Following
the command with a specific module name will return only that module's
information.

- Parameters:
  - `optional(string)` - module name
- CLI command minimal:

```bash
dcld query upgrade module_versions
```

- CLI command full:

```bash
dcld query upgrade module_versions <string>
```

- REST API:
  - GET `/cosmos/upgrade/v1beta1/module_versions`

<!-- markdownlint-enable MD036 -->
