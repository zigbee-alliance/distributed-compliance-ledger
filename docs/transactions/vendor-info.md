# Vendor Info Module

<!-- markdownlint-disable MD036 -->

## VENDOR INFO

### ADD_VENDOR_INFO

**Status: Implemented**

Adds a record about a Vendor.

- Parameters:
  - vid: `uint16` -  Vendor ID (positive non-zero)
  - vendorName: `string` -  Vendor name
  - companyLegalName: `string` -  Legal name of the vendor company
  - companyPreferredName: `optional(string)` -  Preferred name of the vendor company
  - vendorLandingPageURL: `optional(string)` -  URL of the vendor's landing page
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State: `vendorinfo/VendorInfo/value/<vid>`
- Who can send:
  - Account with a vendor role who has the matching Vendor ID
  - Account with a vendor admin role
- CLI command:
  - `dcld tx vendorinfo add-vendor --vid=<uint16> --vendorName=<string> --companyLegalName=<string> --companyPreferredName=<string> --vendorLandingPageURL=<string> --from=<account>`

### UPDATE_VENDOR_INFO

**Status: Implemented**

Updates a record about a Vendor.

- Parameters:
  - vid: `uint16` -  Vendor ID (positive non-zero)
  - vendorName: `optional(string)` -  Vendor name
  - companyLegalName: `optional(string)` -  Legal name of the vendor company
  - companyPreferredName: `optional(string)` -  Preferred name of the vendor company
  - vendorLandingPageURL: `optional(string)` -  URL of the vendor's landing page
  - schemaVersion: `optional(uint16)` - Schema version to support backward/forward compatability. Should be equal to 0 (default 0)
- In State: `vendorinfo/VendorInfo/value/<vid>`
- Who can send:
  - Account with a vendor role who has the matching Vendor ID
  - Account with a vendor admin role
- CLI command:
  - `dcld tx vendorinfo update-vendor --vid=<uint16> ... --from=<account>`

### GET_VENDOR_INFO

**Status: Implemented**

Gets a Vendor Info for the given `vid` (vendor ID).

- Parameters:
  - vid: `uint16` -  model vendor ID (positive non-zero)
- CLI command:
  - `dcld query vendorinfo vendor --vid=<uint16>`
- REST API:
  - GET `/dcl/vendorinfo/vendors/{vid}`

### GET_ALL_VENDOR_INFO

**Status: Implemented**

Gets information about all vendors for all VIDs.

Should be sent to trusted nodes only.

- Parameters:
  - Common pagination parameters (see [pagination-params](#common-pagination-parameters))
- CLI command:
  - `dcld query vendorinfo all-vendors`
- REST API:
  - GET `/dcl/vendorinfo/vendors`

<!-- markdownlint-enable MD036 -->
