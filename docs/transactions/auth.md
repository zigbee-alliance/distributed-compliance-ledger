## AUTH

<!-- markdownlint-disable MD036 -->

### PROPOSE_ADD_ACCOUNT

**Status: Implemented**

Proposes a new Account with the given address, public key and role.

If more than 1 Trustee signature is required to add the account, the account
will be in a pending state until sufficient number of approvals is received.

- Parameters:
  - address: `string` - account address; Bech32 encoded
  - pub_key: `string` - account's Protobuf JSON encoded public key
  - vid: `optional(uint16)` - vendor ID (only needed for vendor role)
  - pid_ranges: `optional(array<uint16 range>)` - the list of product-id ranges (range item separated with "-"), comma-separated, in increasing order, associated with this account: `1-100,201-300...`
  - roles: `array<string>` - the list of roles, comma-separated, assigning to the account. Supported roles: `Vendor`, `TestHouse`, `CertificationCenter`, `Trustee`, `NodeAdmin`, `VendorAdmin`.
  - info: `optional(string)` - information/notes for the proposal
  - time: `optional(int64)` - proposal time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `dclauth/PendingAccount/value/<address>`
- Who can send:
  - Trustee
- CLI command:
  - `dcld tx auth propose-add-account --address=<bench32 encoded string> --pubkey='<protobuf JSON encoded>' --roles=<role1,role2,...> --vid=<uint16> --pid_ranges=<uint16-range,uint16-range,...> --from=<account>`

### APPROVE_ADD_ACCOUNT

**Status: Implemented**

Approves the proposed account. It also can be used for revote (i.e. change vote from reject to approve)

The account is not active until sufficient number of Trustees approve it.

- Parameters:
  - address: `string` - account address; Bech32 encoded
  - info: `optional(string)` - information/notes for the approval
  - time: `optional(int64)` - approval time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `dclauth/Account/value/<address>`
- Who can send:
  - Trustee
- Number of required approvals:
  - greater than or equal 2/3 of Trustees for account roles: `TestHouse`, `CertificationCenter`, `Trustee`, `NodeAdmin`, `VendorAdmin` (proposal by a Trustee is also counted as an approval)
  - greater than 1/3 of Trustees for account role: `Vendor` (proposal by a Trustee is also counted as an approval)
- CLI command:
  - `dcld tx auth approve-add-account --address=<bench32 encoded string> --from=<account>`

> **_Note:_**  If we are approving an account with role `Vendor`, then we need more than 1/3 of Trustees approvals.

### REJECT_ADD_ACCOUNT

**Status: Implemented**

Rejects the proposed account. It also can be used for revote (i.e. change vote from approve to reject)

If proposed account has only proposer's approval and no rejects then proposer can send this transaction to remove the proposal

The account is not reject until sufficient number of Trustees reject it.

- Parameters:
  - address: `string` - account address; Bech32 encoded
  - info: `optional(string)` - information/notes for the reject
  - time: `optional(int64)` - reject time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `dclauth/RejectedAccount/value/<address>`
- Who can send:
  - Trustee
- Number of required rejects:
  - greater than 1/3 of Trustees for account roles: `TestHouse`, `CertificationCenter`, `Trustee`, `NodeAdmin`, `VendorAdmin` (proposal by a Trustee is also counted as an approval)
  - greater than or equal 2/3 of Trustees for account role: `Vendor` (proposal by a Trustee is also counted as an approval)
- CLI command:
  - `dcld tx auth reject-add-account --address=<bench32 encoded string> --from=<account>`

### PROPOSE_REVOKE_ACCOUNT

**Status: Implemented**

Proposes revocation of the Account with the given address.

If more than 1 Trustee signature is required to revoke the account, the revocation
will be in a pending state until sufficient number of approvals is received.

- Parameters:
  - address: `string` - account address; Bech32 encoded
  - info: `optional(string)` - information/notes for the revocation proposal
  - time: `optional(int64)` - revocation proposal time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `dclauth/Account/value/<address>`
- Who can send:
  - Trustee
- CLI command:
  - `dcld tx auth propose-revoke-account --address=<bench32 encoded string> --from=<account>`

### APPROVE_REVOKE_ACCOUNT

**Status: Implemented**

Approves the proposed revocation of the account.

The account is not revoked until sufficient number of Trustees approve it.

- Parameters:
  - address: `string` - account address; Bech32 encoded
  - info: `optional(string)` - information/notes for the revocation approval
  - time: `optional(int64)` - revocation approval time (number of nanoseconds elapsed since January 1, 1970 UTC). This field cannot be specified using a CLI command and will use the current time by default.
- In State: `dclauth/Account/value/<address>`
- Who can send:
  - Trustee
- Number of required approvals:
  - greater than or equal 2/3 of Trustees (proposal by a Trustee is also counted as an approval)
- CLI command:
  - `dcld tx auth approve-revoke-account --address=<bench32 encoded string> --from=<account>`

> **_Note:_**  If revoking an account has sufficient number of Trustees approve it then this account is placed in Revoked Account.

### GET_ACCOUNT

**Status: Implemented**

Gets an accounts by the address. Revoked accounts are not returned.

- Parameters:
  - address: `string` - account address; Bech32 encoded
- CLI command:
  - `dcld query auth account --addres <bench32 encoded string>`
- REST API:
  - GET `/dcl/auth/accounts/{address}`

### GET_PROPOSED_ACCOUNT

**Status: Implemented**

Gets a proposed but not approved accounts by its address

- Parameters:
  - address: `string` - account address; Bech32 encoded
- CLI command:
  - `dcld query auth proposed-account --address <bench32 encoded string>`
- REST API:
  - GET `/dcl/auth/proposed-accounts/{address}`

### GET_REJECTED_ACCOUNT

**Status: Implemented**

Get a rejected accounts by its address

- Parameters:
  - address: `string` - account address; Bech32 encoded
- CLI command:
  - `dcld query auth rejected-account --address <bech32 encoded string>`
- REST API:
  - GET `/dcl/auth/rejected-accounts/{address}`

### GET_PROPOSED_ACCOUNT_TO_REVOKE

**Status: Implemented**

Gets a proposed but not approved accounts to be revoked by its address.

- Parameters:
  - address: `string` - account address; Bech32 encoded
- CLI command:
  - `dcld query auth proposed-account-to-revoke --address <bench32 encoded string>`
- REST API:
  - GET `/dcl/auth/proposed-revocation-accounts/{address}`

### GET_REVOKED_ACCOUNT

**Status: Implemented**

Gets a revoked account by its address.

- Parameters:
  - address: `string` - account address; Bech32 encoded
- CLI command:
  - `dcld query auth revoked-account --address <bench32 encoded string>`
- REST API:
  - GET `/dcl/auth/revoked-accounts/{address}`

### GET_ALL_ACCOUNTS

**Status: Implemented**

Gets all accounts. Revoked accounts are not returned.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query auth all-accounts`
- REST API:
  - GET `/dcl/auth/accounts`

### GET_ALL_PROPOSED_ACCOUNTS

**Status: Implemented**

Gets all proposed but not approved accounts.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query auth all-proposed-accounts`
- REST API:
  - GET `/dcl/auth/proposed-accounts`

### GET_ALL_REJECTED_ACCOUNTS

**Status: Implemented**

Get all rejected accounts.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params]
   (#common-pagination-parameters))
- CLI command:
  - `dcld query auth all-rejected-accounts`
- REST API:
  - GET `/dcl/auth/rejected-accounts`

### GET_ALL_PROPOSED_ACCOUNTS_TO_REVOKE

**Status: Implemented**

Gets all proposed but not approved accounts to be revoked.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query auth all-proposed-accounts-to-revoke`
- REST API:
  - GET `/dcl/auth/proposed-revocation-accounts`

### GET_ALL_REVOKED_ACCOUNTS

**Status: Implemented**

Gets all revoked accounts.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query auth all-revoked-accounts`
- REST API:
  - GET `/dcl/auth/revoked-accounts`

### ROTATE_KEY

**Status: Not Implemented**

Rotate's the Account's public key by the owner.

- Who can send:
  - Any role; owner


<!-- markdownlint-enable MD036 -->