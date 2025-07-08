
   * [All Certificates (DA, NOC)](#all_certificates)
      * [GET_CERT](#get_cert)
      * [GET_ALL_CERTS](#get_all_certs)
      * [GET_ALL_CERTS_BY_SUBJECT](#get_all_certs_by_subjects)
      * [GET_ALL_CERTS_BY_SKID](#get_all_certs_by_skid)
      * [GET_CHILD_CERTS](#get_child_certs)
   * [Device Attestation Certificates](#device_attestation_certificate)
      * [PROPOSE_ADD_PAA](#propose_add_paa)
      * [APPROVE_ADD_PAA](#approve_add_paa)
      * [REJECT_ADD_PAA](#reject_add_paa)
      * [PROPOSE_REVOKE_PAA](#propose_revoke_paa)
      * [APPROVE_REVOKE_PAA](#approve_revoke_paa)
      * [ASSIGN_VID_TO_PAA](#assign_vid_to_paa)
      * [ADD_REVOCATION_DISTRIBUTION_POINT](#add_revocation_distribution_point)
      * [UPDATE_REVOCATION_DISTRIBUTION_POINT](#update_revocation_distribution_point)
      * [DELETE_REVOCATION_DISTRIBUTION_POINT](#delete_revocation_distribution_point)
      * [ADD_PAI](#add_pai)
      * [REVOKE_PAI](#revoke_pai)
      * [REMOVE_PAI](#remove_pai)
      * [GET_DA_CERT](#get_da_cert)
      * [GET_REVOKED_DA_CERT](#get_revoked_da_cert)
      * [GET_DA_CERTS_BY_SKID](#get_da_certs_by_skid)
      * [GET_DA_CERTS_BY_SUBJECT](#get_da_certs_by_subject)
      * [GET_ALL_DA_CERTS](#get_all_da_certs)
      * [GET_ALL_REVOKED_DA_CERTS](#get_all_revoked_da_certs)
      * [GET_PKI_REVOCATION_DISTRIBUTION_POINT](#get_pki_revocation_distribution_point)
      * [GET_PKI_REVOCATION_DISTRIBUTION_POINTS_BY_SUBJECT_KEY_ID](#get_pki_revocation_distribution_points_by_subject_key_id)
      * [GET_ALL_PKI_REVOCATION_DISTRIBUTION_POINT](#get_all_pki_revocation_distribution_point)
      * [GET_PROPOSED_PAA](#get_proposed_paa)
      * [GET_REJECTED_PAA](#get_rejected_paa)
      * [GET_PROPOSED_PAA_TO_REVOKE](#get_proposed_paa_to_revoke)
      * [GET_ALL_PAA](#get_all_paa)
      * [GET_ALL_REVOKED_PAA](#get_all_revoked_paa)
      * [GET_ALL_PROPOSED_PAA](#get_all_proposed_paa)
      * [GET_ALL_REJECTED_PAA](#get_all_rejected_paa)
      * [GET_ALL_PROPOSED_PAA_TO_REVOKE](#get_all_proposed_paa_to_revoke)
   * [E2E (NOC): RCAC, ICAC](#e2e_(noc):_rcac,_icac)
      * [ADD_NOC_ROOT (RCAC)](#add_noc_root_(rcac))
      * [REVOKE_NOC_ROOT (RCAC)](#revoke_noc_root_(rcac))
      * [REMOVE_NOC_ROOT (RCAC)](#remove_noc_root(rcac))
      * [ADD_NOC_ICA (ICAC)](#add_noc_ica_(icac))
      * [REVOKE_NOC_ICA (ICAC)](#revoke_noc_ica_(icac))
      * [REMOVE_NOC_ICA (ICAC)](#remove_noc_ica_(icac))
      * [GET_NOC_CERT](#get_noc_cert)
      * [GET_NOC_ROOT_BY_VID (RCACs)](#get_noc_root_by_vid_(rcacs))
      * [GET_NOC_BY_VID_AND_SKID (RCACs/ICACs)](#get_noc_by_vid_and_skid_(rcacs/icacs))
      * [GET_NOC_ICA_BY_VID (ICACs)](#get_noc_ica_by_vid_(icacs))
      * [GET_NOC_CERTS_BY_SUBJECT](#get_noc_certs_by_subject)
      * [GET_REVOKED_NOC_ROOT (RCAC)](#get_revoked_noc_root_(rcac))
      * [GET_REVOKED_NOC_ICA (ICAC)](#get_revoked_noc_ica_(icac))
      * [GET_ALL_NOC (RCACs/ICACs)](#get_all_noc_(rcacs/icacs))
      * [GET_ALL_NOC_ROOT (RCACs)](#get_all_noc_root_(rcacs))
      * [GET_ALL_NOC_ICA (ICACs)](#get_all_noc_ica_(icacs))
      * [GET_ALL_REVOKED_NOC_ROOT (RCACs)](#get_all_revoked_noc_root-(rcacs))
      * [GET_ALL_REVOKED_NOC_ICA (ICACs)](#get_all_revoked_noc_ica_(icacs))

# All Certificates (DA, NOC)	
## [GET_CERT]()

### CLI command  
CLI command: `dcld query certificate get-cert --subject=<string> --subject-key-id=<string>`

* CLI command send  
     * Valid command  
          * command exists/relevant  
          * query works for all types of certificates:  
               * PAA  
               * PAI  
               * RCAC  
               * ICAC  
     * Invalid command  
          * access is denied to execute command  
          * incorrect command syntax  

* Command result  
     * GET_CERT command completed successfully ⇒ gets a certificate by the given subject and subject key ID attributes  
          * certificate with given subject and subject key ID attributes exists  
     * GET_CERT command failed ⇒ does not get a certificate  
          * certificate with given subject and subject key ID attributes does not exist  

* Role (Who can send)  
     * Trustee  
     * Vendor  
     * VendorAdmin  
     * CertificationCenter  
     * NodeAdmin  

* Parameters:  
     * **subject (Subject)** - string  
          * Positive:  
               * string matches the format  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * **subject_key_id (Subject Key ID)** - string  
          * Positive:  
               * formatted string (e.g., `5A:88:0E:6C:36:53:D0:7F:...`)  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  

---

### REST API  
REST API command: `GET /certificates/{subject}/{subjectKeyId}`

* REST API command send  
     * Valid command  
          * correct HTTP method  
          * request is authorized  
          * uses valid credentials/role  
          * works for all certificate types:  
               * PAA  
               * PAI  
               * RCAC  
               * ICAC  
     * Invalid command  
          * incorrect request  
          * server side error  

* Command result  
     * GET_CERT command completed successfully ⇒ gets a certificate by the given subject and subject key ID attributes  
          * certificate with given subject and subject key ID attributes exists  
     * GET_CERT command failed ⇒ does not get a certificate  
          * certificate with given subject and subject key ID attributes does not exist  

* Role (Who can send)  
     * Trustee  
     * Vendor  
     * VendorAdmin  
     * CertificationCenter  
     * NodeAdmin  

* Parameters:
     * **subject (Subject)** - string  
          * Positive:  
               * string matches the format  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * **subject_key_id (Subject Key ID)** - string  
          * Positive:  
               * formatted string (e.g., `5A:88:0E:6C:36:53:D0:7F:...`)  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX 

## [GET_ALL_CERTS]()

### CLI command  
CLI command: `dcld query certificate all-certs [--count-total=<bool>] [--limit=<uint>] [--offset=<uint>] [--page=<uint>] [--page-key=<string>] [--reverse=<bool>]`

* CLI command send  
     * Valid command  
          * command exists/relevant  
          * query works for all types of certificates  
               * PAA  
               * PAI  
               * RCAC  
               * ICAC  
     * Invalid command  
          * access is denied to execute command  
          * incorrect command syntax  

* Command result  
     * GET_ALL_CERTS command completed successfully ⇒ gets all certificates  
          * there is at least one certificate  
     * GET_ALL_CERTS command failed ⇒ does not get all certificates  
          * there is not a single certificate  

* Role (Who can send)  
     * Trustee  
     * Vendor  
     * VendorAdmin  
     * CertificationCenter  
     * NodeAdmin  

* Parameters:  
     * **count-total** (optional, bool)  
          * Positive:  
               * empty value  
               * value state  
               * TRUE (-1)  
               * FALSE (0)  
          * Negative:  
               * value is not bool  
     * **limit** (optional, uint)  
          * Positive:  
               * value exists  
               * empty value  
          * Negative:  
               * value > 100  
     * **offset** (optional, uint)  
          * Positive:  
               * value exists  
               * empty value  
          * Negative:  
               * value < 0  
     * **page** (optional, uint)  
          * Positive:  
               * value exists  
               * empty value  
          * Negative:  
               * value < 0  
     * **page-key** (optional, string)  
          * Positive:  
               * empty value  
               * value exists  
          * Negative:  
               * length < MIN  
     * **reverse** (optional, bool)  
          * Positive:  
               * empty value  
               * value state  
               * TRUE (-1)  
               * FALSE (0)  
          * Negative:  
               * value is not bool  

---

### REST API  
REST API command: `GET /certificates`

* REST API command send  
     * Valid command  
          * correct HTTP method  
          * request is authorized  
          * uses valid credentials/role  
          * query works for all types of certificates  
               * PAA  
               * PAI  
               * RCAC  
               * ICAC  
     * Invalid command  
          * incorrect request  
          * server side error  

* Command result  
     * GET_ALL_CERTS command completed successfully ⇒ gets all certificates  
          * there is at least one certificate  
     * GET_ALL_CERTS command failed ⇒ does not get all certificates  
          * there is not a single certificate  

* Role (Who can send)  
     * Trustee  
     * Vendor  
     * VendorAdmin  
     * CertificationCenter  
     * NodeAdmin  

* Parameters: 
     * **count-total** (optional, bool)  
          * Positive:  
               * empty value  
               * value state  
               * TRUE (-1)  
               * FALSE (0)  
          * Negative:  
               * value is not bool  
     * **limit** (optional, uint)  
          * Positive:  
               * value exists  
               * empty value  
          * Negative:  
               * value > 100  
     * **offset** (optional, uint)  
          * Positive:  
               * value exists  
               * empty value  
          * Negative:  
               * value < 0  
     * **page** (optional, uint)  
          * Positive:  
               * value exists  
               * empty value  
          * Negative:  
               * value < 0  
     * **page-key** (optional, string)  
          * Positive:  
               * empty value  
               * value exists  
          * Negative:  
               * length < MIN  
     * **reverse** (optional, bool)  
          * Positive:  
               * empty value  
               * value state  
               * TRUE (-1)  
               * FALSE (0)  
          * Negative:  
               * value is not bool  

## [GET_ALL_CERTS_BY_SUBJECT]()

### CLI command  
CLI command: `dcld query certificate all-certs-by-subject --subject=<string>`

* CLI command send  
     * Valid command  
          * command exists/relevant  
          * query works for all types of certificates  
               * PAA  
               * PAI  
               * RCAC  
               * ICAC  
     * Invalid command  
          * access is denied to execute command  
          * incorrect command syntax  

* Command result  
     * GET_ALL_CERTS_BY_SUBJECT command completed successfully ⇒ gets all certificates associated with a subject  
          * there is at least one certificate associated with a subject  
     * GET_ALL_CERTS_BY_SUBJECT command failed ⇒ does not get all certificates associated with a subject  
          * there is not one certificate associated with a subject  

* Role (Who can send)  
     * Trustee  
     * Vendor  
     * VendorAdmin  
     * CertificationCenter  
     * NodeAdmin  

* Parameters:  
     * **subject (Subject)** - string  
          * Positive:  
               * string matches the format  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  

---

### REST API  
REST API command: `GET /certificates/by-subject/{subject}`

* REST API command send  
     * Valid command  
          * correct HTTP method  
          * request is authorized  
          * uses valid credentials/role  
          * query works for all types of certificates  
               * PAA  
               * PAI  
               * RCAC  
               * ICAC  
     * Invalid command  
          * incorrect request  
          * server side error  

* Command result  
     * GET_ALL_CERTS_BY_SUBJECT command completed successfully ⇒ gets all certificates associated with a subject  
          * there is at least one certificate associated with a subject  
     * GET_ALL_CERTS_BY_SUBJECT command failed ⇒ does not get all certificates associated with a subject  
          * there is not one certificate associated with a subject  

* Role (Who can send)  
     * Trustee  
     * Vendor  
     * VendorAdmin  
     * CertificationCenter  
     * NodeAdmin  

* Parameters:
     * **subject (Subject)** - string  
          * Positive:  
               * string matches the format  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  

## [GET_ALL_CERTS_BY_SKID]()

### CLI command  
CLI command: `dcld query certificate all-certs-by-skid --subject-key-id=<string>`

* CLI command send  
     * Valid command  
          * command exists/relevant  
          * query works for all types of certificates  
               * PAA  
               * PAI  
               * RCAC  
               * ICAC  
     * Invalid command  
          * access is denied to execute command  
          * incorrect command syntax  

* Command result  
     * GET_ALL_CERTS_BY_SKID command completed successfully ⇒ gets all certificates by the given subject key ID attribute  
          * there is at least one certificate by the given subject key ID attribute  
     * GET_ALL_CERTS_BY_SKID command failed ⇒ does not get certificates  
          * there is not one certificate by the given subject key ID attribute  

* Role (Who can send)  
     * Trustee  
     * Vendor  
     * VendorAdmin  
     * CertificationCenter  
     * NodeAdmin  

* Parameters:  
     * **subject_key_id (Subject Key ID)** - string  
          * Positive:  
               * string matches the format (e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`)  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  

---

### REST API  
REST API command: `GET /certificates/by-skid/{subjectKeyId}`

* REST API command send  
     * Valid command  
          * correct HTTP method  
          * request is authorized  
          * uses valid credentials/role  
          * query works for all types of certificates  
               * PAA  
               * PAI  
               * RCAC  
               * ICAC  
     * Invalid command  
          * incorrect request  
          * server side error  

* Command result  
     * GET_ALL_CERTS_BY_SKID command completed successfully ⇒ gets all certificates by the given subject key ID attribute  
          * there is at least one certificate by the given subject key ID attribute  
     * GET_ALL_CERTS_BY_SKID command failed ⇒ does not get certificates  
          * there is not one certificate by the given subject key ID attribute  

* Role (Who can send)  
     * Trustee  
     * Vendor  
     * VendorAdmin  
     * CertificationCenter  
     * NodeAdmin  

* Parameters:
     * **subject_key_id (Subject Key ID)** - string  
          * Positive:  
               * string matches the format (e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`)  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  

## [GET_CHILD_CERTS]()

### CLI command  
CLI command: `dcld query certificate child-certs --subject=<string> --subject-key-id=<string>`

* CLI command send  
     * Valid command  
          * command exists/relevant  
          * query works for child certificates  
               * PAI  
               * NOC_ICA  
     * Invalid command  
          * access is denied to execute command  
          * incorrect command syntax  

* Command result  
     * GET_CHILD_CERTS command completed successfully ⇒ gets all child certificates for the given certificate  
          * there is at least one child certificate for the given certificate  
     * GET_CHILD_CERTS command failed ⇒ does not get all child certificates  
          * there is not one child certificate for the given certificate  

* Role (Who can send)  
     * Trustee  
     * Vendor  
     * VendorAdmin  
     * CertificationCenter  
     * NodeAdmin  

* Parameters:  
     * **subject (Subject)** - string  
          * Positive:  
               * string matches the format  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * **subject_key_id (Subject Key ID)** - string  
          * Positive:  
               * string matches the format (e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`)  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  

---

### REST API  
REST API command: `GET /certificates/children/{subject}/{subjectKeyId}`

* REST API command send  
     * Valid command  
          * correct HTTP method  
          * request is authorized  
          * uses valid credentials/role  
          * query works for child certificates  
               * PAI  
               * NOC_ICA  
     * Invalid command  
          * incorrect request  
          * server side error  

* Command result  
     * GET_CHILD_CERTS command completed successfully ⇒ gets all child certificates for the given certificate  
          * there is at least one child certificate for the given certificate  
     * GET_CHILD_CERTS command failed ⇒ does not get all child certificates  
          * there is not one child certificate for the given certificate  

* Role (Who can send)  
     * Trustee  
     * Vendor  
     * VendorAdmin  
     * CertificationCenter  
     * NodeAdmin  

* Parameters: 
     * **subject (Subject)** - string  
          * Positive:  
               * string matches the format  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * **subject_key_id (Subject Key ID)** - string  
          * Positive:  
               * string matches the format (e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`)  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX   

# Device Attestation Certificates (DA): PAA, PAI	
## [PROPOSE_ADD_PAA]()
### CLI command  
CLI command: `dcld tx certificate propose-add-paa --cert=<string> --vid=<uint16> [--info=<string>] [--time=<int64>] [--schema-version=<uint16>]`

* CLI command send  
     * Valid command  
          * command exists/relevant  
          * query works for the following certificates  
               * PAA  
     * Invalid command  
          * access is denied to execute command  
          * incorrect command syntax  

* Command result  
     * PROPOSE_ADD_PAA command completed successfully ⇒ proposes a new PAA (self-signed root certificate)  
          * sufficient number of approvals is received → certificate added  
          * provided certificate is root  
               * Issuer == Subject  
               * Authority Key Identifier == Subject Key Identifier  
          * no existing Proposed certificate with the same <Certificate's Subject>:<Certificate's Subject Key ID> combination  
          * certificates with the same <Certificate's Subject>:<Certificate's Subject Key ID> combination already exist  
          * the existing certificate is not a NOC certificate  
          * sender matches the owner of the existing certificates  
          * no existing certificate with the same <Certificate's Issuer>:<Certificate's Serial Number> combination  
          * signature (self-signature) and expiration date are valid  
     * PROPOSE_ADD_PAA command failed ⇒ does not propose a new PAA  
          * insufficient number of approvals received → certificate is in Pending state  
          * certificate not added  
          * user tries to edit certificate → PAA certificate is immutable  
          * provided certificate is not root  
               * Issuer != Subject, but Authority Key Identifier == Subject Key Identifier  
               * Issuer == Subject, but Authority Key Identifier != Subject Key Identifier  
               * Issuer != Subject and Authority Key Identifier != Subject Key Identifier  
          * existing Proposed certificate with same Subject:SKID combination  
          * existing certificate is a NOC certificate  
          * sender does not match the owner of the existing certificates  
          * existing certificate with same Issuer:Serial Number  
          * signature (self-signature) is not valid  
          * signature valid but expiration is not  
          * both signature and expiration are invalid  

* Role  

     * Who can send
          * Positive: 
               * Trustee 
          * Negative:   
               * Vendor  
               * VendorAdmin  
               * CertificationCenter  
               * NodeAdmin 

     * Who can revoke 
          * Positive:
               * Trustee  
                    * owner → if there was 1 signature for the certificate  
                    * quorum → if more than 1 signature for the certificate   
          * Negative: 
               * Vendor   
               * VendorAdmin   
               * CertificationCenter  
               * NodeAdmin 

* Parameters:  
     * **cert (Certificate)** - string  
          * Positive:  
               * value exists  
               * contains PEM string or path to file  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * **info (Information/Notes)** - optional string  
          * Positive:  
               * empty value  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * length > MAX (4096 characters)  
     * **time (Proposal Time)** - optional int64  
          * Positive:  
               * default value is current time  
               * empty value  
               * integer value format  
          * Negative:  
               * length > MAX (9_223_372_036_854_775_807)  
     * **vid (Vendor ID)** - uint16  
          * Positive:  
               * unique combination  
               * value > 0  
               * integer value format  
               * Vendor ID matches Certificate's vid field for VID-scoped PAA  
          * Negative:  
               * empty value  
               * value <= 0  
               * string value format  
               * length > MAX (65535)  
     * **schemaVersion (Schema Version)** - optional uint16  
          * Positive:  
               * value = 0  
               * integer value format  
               * empty value  
          * Negative:  
               * length > MAX (65535)  

---

### REST API  
REST API command: `POST /certificates/propose-add-paa`

* REST API command send  
     * Valid command  
          * correct HTTP method  
          * request is authorized  
          * uses valid credentials/role  
          * query works for the following certificates  
               * PAA  
     * Invalid command  
          * incorrect request  
          * server side error  

* Command result  
     * PROPOSE_ADD_PAA command completed successfully ⇒ proposes a new PAA (self-signed root certificate)  
          * identical results as CLI above  
     * PROPOSE_ADD_PAA command failed ⇒ does not propose new PAA  
          * identical results as CLI above  

* Role  

     * Who can send
          * Positive: 
               * Trustee 
          * Negative:   
               * Vendor  
               * VendorAdmin  
               * CertificationCenter  
               * NodeAdmin 

     * Who can revoke 
          * Positive:
               * Trustee  
                    * owner → if there was 1 signature for the certificate  
                    * quorum → if more than 1 signature for the certificate   
          * Negative: 
               * Vendor   
               * VendorAdmin   
               * CertificationCenter  
               * NodeAdmin 

* Parameters:  
     * **cert (Certificate)** - string  
          * Positive:  
               * value exists  
               * contains PEM string or path to file  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * **info (Information/Notes)** - optional string  
          * Positive:  
               * empty value  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * length > MAX (4096 characters)  
     * **time (Proposal Time)** - optional int64  
          * Positive:  
               * default value is current time  
               * empty value  
               * integer value format  
          * Negative:  
               * length > MAX (9_223_372_036_854_775_807)  
     * **vid (Vendor ID)** - uint16  
          * Positive:  
               * unique combination  
               * value > 0  
               * integer value format  
               * Vendor ID matches Certificate's vid field for VID-scoped PAA  
          * Negative:  
               * empty value  
               * value <= 0  
               * string value format  
               * length > MAX (65535)  
     * **schemaVersion (Schema Version)** - optional uint16  
          * Positive:  
               * value = 0  
               * integer value format  
               * empty value  
          * Negative:  
               * length > MAX (65535) 

## [APPROVE_ADD_PAA]()

### CLI command  
CLI command: `dcld tx certificate approve-add-paa --subject=<string> --subject-key-id=<string> [--info=<string>] [--time=<int64>]`

* CLI command send  
     * Valid command  
          * command exists/relevant  
          * query works for the following certificates  
               * PAA  
     * Invalid command  
          * access is denied to execute command  
          * incorrect command syntax  

* Command result  
     * APPROVE_ADD_PAA command completed successfully ⇒ approves the proposed PAA (self-signed root certificate)  
          * PROPOSE_ADD_PAA command completed successfully  
          * command used for re-voting (i.e., change vote from reject to approve)  
          * number of approvals > 2/3 of Trustees ⇒ certificate active  
          * number of approvals = 2/3 of Trustees ⇒ certificate active  
          * the proposal to add a root certificate with the provided subject and subject_key_id was submitted first  
          * the proposed certificate hasn't been approved by the signer yet  
     * APPROVE_ADD_PAA command failed ⇒ does not approve the proposed PAA  
          * PROPOSE_ADD_PAA command failed  
          * number of approvals < 2/3 of Trustees ⇒ certificate not active  
          * the proposal to add a root certificate with the provided subject and subject_key_id was not submitted first  
          * the proposed certificate has already been approved by the signer  

* Role  

     * Who can send
          * Positive:
               * Trustee     
          * Negative: 
               * Vendor 
               * VendorAdmin  
               * CertificationCenter
               * NodeAdmin 

* Parameters:  
     * **subject (Subject)** - string  
          * Positive:  
               * string matches the format  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * **subject_key_id (Subject Key ID)** - string  
          * Positive:  
               * formatted string (e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`)  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * **info (Information/Notes)** - optional string  
          * Positive:  
               * empty value  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * length > MAX (4096 characters)  
     * **time (Proposal Time)** - optional int64  
          * Positive:  
               * default value is current time  
               * empty value  
               * integer value format  
          * Negative:  
               * length > MAX (9_223_372_036_854_775_807)  

---

### REST API  
REST API command: `POST /certificates/approve-add-paa`

* REST API command send  
     * Valid command  
          * correct HTTP method  
          * request is authorized  
          * uses valid credentials/role  
          * query works for the following certificates  
               * PAA  
     * Invalid command  
          * incorrect request  
          * server side error  

* Command result  
     * APPROVE_ADD_PAA command completed successfully ⇒ approves the proposed PAA (self-signed root certificate)  
          * identical outcomes as CLI  
     * APPROVE_ADD_PAA command failed ⇒ does not approve the proposed PAA  
          * identical outcomes as CLI  

* Role
     * Who can send
          * Positive:
               * Trustee     
          * Negative: 
               * Vendor 
               * VendorAdmin  
               * CertificationCenter
               * NodeAdmin 

* Parameters:  
     * **subject (Subject)** - string  
          * Positive:  
               * string matches the format  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * **subject_key_id (Subject Key ID)** - string  
          * Positive:  
               * formatted string (e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`)  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * **info (Information/Notes)** - optional string  
          * Positive:  
               * empty value  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * length > MAX (4096 characters)  
     * **time (Proposal Time)** - optional int64  
          * Positive:  
               * default value is current time  
               * empty value  
               * integer value format  
          * Negative:  
               * length > MAX (9_223_372_036_854_775_807)  

## REJECT_ADD_PAA

### CLI command
CLI command send

* Valid command
     * command exists/relevant
     * query works for the following certificates
          * PAA

* Invalid command
     * access is denied to execute command
     * incorrect command syntax

* Command result
     * REJECT_ADD_PAA command completed successfully ⇒ rejects the proposed PAA (self-signed root certificate)
          * PROPOSE_ADD_PAA command completed successfully
          * APPROVE_ADD_PAA command completed successfully
          * command used for re-voting ⇒ change vote from approve to reject
          * remove the proposal
          * proposed PAA certificate has only proposer's approval and no rejects
          * number of approvals greater than 1/3 of Trustees ⇒ certificate rejects
          * the proposal to add a root certificate with the provided subject and subject_key_id, submitted first
          * the proposed certificate hasn't been rejected by the signer yet
     * REJECT_ADD_PAA command failed ⇒ does not reject the proposed PAA (self-signed root certificate)
          * PROPOSE_ADD_PAA command failed
          * APPROVE_ADD_PAA command failed
          * number of approvals is less than 1/3 of Trustees ⇒ certificate is not rejected
          * number of approvals equal 1/3 of Trustees ⇒ certificate is not rejected
          * the proposal to add a root certificate with the provided subject and subject_key_id, not submitted first
          * the proposed certificate has been rejected by the signer
          * remove the proposal
          * certificate has not proposer's approval
          * certificate has only proposer's approval and rejects

* Role (Who can send)
     * Positive:
          * Trustee
     * Negative:
          * Vendor
          * VendorAdmin
          * CertificationCenter
          * NodeAdmin

* Parameters:
     * subject (Subject) - string
          * Positive:
               * string matches the format
               * text value format
               * MIN < length < MAX
          * Negative:
               * empty value
               * nonexistent value
               * length > MAX

     * subject_key_id (Subject Key ID) - string
          * Positive:
               * string matches the format (e.g., 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB)
               * text value format
               * MIN < length < MAX
          * Negative:
               * empty value
               * nonexistent value
               * length > MAX

     * info (Information/Notes) - optional(string)
          * Positive:
               * empty value
               * text value format
               * MIN < length < MAX
          * Negative:
               * length > MAX (MAX = 4096 characters)

     * time (Proposal Time) - optional(int64)
          * Positive:
               * default value ⇒ current time by default
               * empty value
               * integer value format
          * Negative:
               * length > MAX (MAX = 9 223 372 036 854 775 807)

### REST API
REST API command send

* Valid command
     * correct HTTP method
     * request is authorized
     * uses valid credentials/role
     * query works for the following certificates
          * PAA

* Invalid command
     * incorrect request
     * server side error

* Command result
     * REJECT_ADD_PAA command completed successfully ⇒ rejects the proposed PAA (self-signed root certificate)
          * PROPOSE_ADD_PAA command completed successfully
          * APPROVE_ADD_PAA command completed successfully
          * command used for re-voting ⇒ change vote from approve to reject
          * remove the proposal
          * proposed PAA certificate has only proposer's approval and no rejects
          * number of approvals greater than 1/3 of Trustees ⇒ certificate rejects
          * the proposal to add a root certificate with the provided subject and subject_key_id, submitted first
          * the proposed certificate hasn't been rejected by the signer yet
     * REJECT_ADD_PAA command failed ⇒ does not reject the proposed PAA (self-signed root certificate)
          * PROPOSE_ADD_PAA command failed
          * APPROVE_ADD_PAA command failed
          * number of approvals is less than 1/3 of Trustees ⇒ certificate is not rejected
          * number of approvals equal 1/3 of Trustees ⇒ certificate is not rejected
          * the proposal to add a root certificate with the provided subject and subject_key_id, not submitted first
          * the proposed certificate has been rejected by the signer
          * remove the proposal
          * certificate has not proposer's approval
          * certificate has only proposer's approval and rejects

* Parameters:
     * subject (Subject) - string
          * Positive:
               * string matches the format
               * text value format
               * MIN < length < MAX
          * Negative:
               * empty value
               * nonexistent value
               * length > MAX

     * subject_key_id (Subject Key ID) - string
          * Positive:
               * string matches the format (e.g., 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB)
               * text value format
               * MIN < length < MAX
          * Negative:
               * empty value
               * nonexistent value
               * length > MAX

     * info (Information/Notes) - optional(string)
          * Positive:
               * empty value
               * text value format
               * MIN < length < MAX
          * Negative:
               * length > MAX (MAX = 4096 characters)

     * time (Proposal Time) - optional(int64)
          * Positive:
               * default value ⇒ current time by default
               * empty value
               * integer value format
          * Negative:
               * length > MAX (MAX = 9 223 372 036 854 775 807)

## PROPOSE_REVOKE_PAA

### CLI command  
CLI command: `dcld tx paa propose-revoke-paa --subject=<string> --subject-key-id=<string> [--serial-number=<string>] [--revoke-child=<bool>] [--info=<string>] [--time=<int64>] --from=<account>`

* CLI command send  
     * Valid command  
          * command exists/relevant  
          * query works for the following certificates  
               * PAA  
     * Invalid command  
          * access is denied to execute command  
          * incorrect command syntax  

* Command result  
     * PROPOSE_REVOKE_PAA command completed successfully ⇒ proposes revocation of the given PAA (self-signed root certificate) by a Trustee  
          * PROPOSE_ADD_PAA command completed successfully  
          * APPROVE_ADD_PAA command completed successfully  
          * revoke-child = True ⇒ all the certificates in the chain signed by the revoked certificate will be revoked as well  
          * revoke-child = Falce ⇒ the certificates in the chain signed by the revoked certificate not be revoked  
          * sufficient number of Trustee's approvals is received ⇒ PAA certificate is revoked  
          * revoked certificate root  
               * Issuer == Subject  
               * Authority Key Identifier == Subject Key Identifier  
          * no existing Proposed certificate with the same <Certificate's Subject>:<Certificate's Subject Key ID> combination  
     * PROPOSE_REVOKE_PAA command failed ⇒ does not propose revocation of the given PAA (self-signed root certificate) by a Trustee  
          * PROPOSE_ADD_PAA command failed  
          * APPROVE_ADD_PAA command failed  
          * not sufficient number of Trustee's approvals is received ⇒ PAA certificate in the pending state  
          * there are no Trustee's approvals ⇒ PAA certificate is not revoked  
          * revoked certificate is not root  
               * Issuer != Subject and Authority Key Identifier == Subject Key Identifier  
               * Issuer != Subject and Authority Key Identifier != Subject Key Identifier  
               * Issuer == Subject and Authority Key Identifier != Subject Key Identifier  
          * existing Proposed certificate with the same <Certificate's Subject>:<Certificate's Subject Key ID> combination  

* Role (Who can send)  
     * Positive:  
          * Trustee  
     * Negative:  
          * Vendor ⇒ error  
          * VendorAdmin ⇒ error  
          * CertificationCenter ⇒ error  
          * NodeAdmin ⇒ error  

* Parameters:  
     * subject (Subject) - string  
          * Positive:  
               * string matches the format  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * subject_key_id (Subject Key ID) - string  
          * Positive:  
               * string matches the format  
                    * e.g., 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * serial-number (Serial Number) - optional(string)  
          * Positive:  
               * empty value  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * length > MAX  
     * revoke-child (Revoke Child) - optional(bool)  
          * Positive:  
               * empty value  
               * value state  
               * TRUE (-1)  
               * FALSE (0) ⇒ default value  
          * Negative:  
               * value is not bool  
     * info (Information/Notes) - optional(string)  
          * Positive:  
               * empty value  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * length > MAX ⇒ MAX = 4096 characters  
     * time (Proposal Time) - optional(int64)  
          * Positive:  
               * default value ⇒ current time by default  
               * empty value  
               * integer value format  
          * Negative:  
               * length > MAX ⇒ MAX = 9 223 372 036 854 775 807  

### REST API  
REST API command: POST /cosmos/tx/v1beta1/txs: [MsgProposeRevokePAA](https://github.com/zigbee-alliance/distributed-compliance-ledger)

* REST API command send  
     * Valid command  
          * correct HTTP method  
          * request is authorized  
          * uses valid credentials/role  
          * query works for the following certificates  
               * PAA  
     * Invalid command  
          * incorrect request  
          * server side error  

* Command result  
     * PROPOSE_REVOKE_PAA command completed successfully ⇒ proposes revocation of the given PAA (self-signed root certificate) by a Trustee  
          * PROPOSE_ADD_PAA command completed successfully  
          * APPROVE_ADD_PAA command completed successfully  
          * revoke-child = True ⇒ all the certificates in the chain signed by the revoked certificate will be revoked as well  
          * revoke-child = Falce ⇒ the certificates in the chain signed by the revoked certificate not be revoked  
          * sufficient number of Trustee's approvals is received ⇒ PAA certificate is revoked  
          * revoked certificate root  
               * Issuer == Subject  
               * Authority Key Identifier == Subject Key Identifier  
          * no existing Proposed certificate with the same <Certificate's Subject>:<Certificate's Subject Key ID> combination  
     * PROPOSE_REVOKE_PAA command failed ⇒ does not propose revocation of the given PAA (self-signed root certificate) by a Trustee  
          * PROPOSE_ADD_PAA command failed  
          * APPROVE_ADD_PAA command failed  
          * not sufficient number of Trustee's approvals is received ⇒ PAA certificate in the pending state  
          * there are no Trustee's approvals ⇒ PAA certificate is not revoked  
          * revoked certificate is not root  
               * Issuer != Subject and Authority Key Identifier == Subject Key Identifier  
               * Issuer != Subject and Authority Key Identifier != Subject Key Identifier  
               * Issuer == Subject and Authority Key Identifier != Subject Key Identifier  
          * existing Proposed certificate with the same <Certificate's Subject>:<Certificate's Subject Key ID> combination  

* Role (Who can send)  
     * Positive:  
          * Trustee  
     * Negative:  
          * Vendor ⇒ error  
          * VendorAdmin ⇒ error  
          * CertificationCenter ⇒ error  
          * NodeAdmin ⇒ error  

* Parameters:  
     * subject (Subject) - string  
          * Positive:  
               * string matches the format  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * subject_key_id (Subject Key ID) - string  
          * Positive:  
               * string matches the format  
                    * e.g., 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * serial-number (Serial Number) - optional(string)  
          * Positive:  
               * empty value  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * length > MAX  
     * revoke-child (Revoke Child) - optional(bool)  
          * Positive:  
               * empty value  
               * value state  
               * TRUE (-1)  
               * FALSE (0) ⇒ default value  
          * Negative:  
               * value is not bool  
     * info (Information/Notes) - optional(string)  
          * Positive:  
               * empty value  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * length > MAX ⇒ MAX = 4096 characters  
     * time (Proposal Time) - optional(int64)  
          * Positive:  
               * default value ⇒ current time by default  
               * empty value  
               * integer value format  
          * Negative:  
               * length > MAX ⇒ MAX = 9 223 372 036 854 775 807  


## APPROVE_REVOKE_PAA
### CLI command  
CLI command: `dcld tx paa approve-revoke-paa --subject=<string> --subject-key-id=<string> [--serial-number=<string>] [--info=<string>] [--time=<int64>] --from=<account>`

* CLI command send  
     * Valid command  
          * command exists/relevant  
          * query works for the following certificates  
               * PAA  
     * Invalid command  
          * access is denied to execute command  
          * incorrect command syntax  

* Command result  
     * APPROVE_REVOKE_PAA command completed successfully ⇒ approves the revocation of the given PAA (self-signed root certificate) by a Trustee  
          * PROPOSE_ADD_PAA command completed successfully  
          * APPROVE_ADD_PAA command completed successfully  
          * PROPOSE_REVOKE_PAA command completed successfully  
          * Number of required approvals greater than 2/3 of Trustees ⇒ revocation is applied  
          * Number of required approvals equal 2/3 of Trustees ⇒ revocation is applied  
          * the proposal to revoke a root certificate with the provided subject and subject_key_id, submitted first  
          * the proposed certificate revocation hasn't been approved by the signer yet  
     * APPROVE_REVOKE_PAA command failed ⇒ does not approve the revocation of the given PAA (self-signed root certificate) by a Trustee  
          * PROPOSE_ADD_PAA command failed  
          * APPROVE_ADD_PAA command failed  
          * PROPOSE_REVOKE_PAA command failed  
          * Number of required approvals less than 2/3 of Trustees ⇒ revocation is not applied  
          * the proposal to revoke a root certificate with the provided subject and subject_key_id, not submitted first  
          * the proposed certificate revocation has been approved by the signer  

* Role (Who can send)  
     * Positive:  
          * Trustee  
     * Negative:  
          * Vendor ⇒ error  
          * VendorAdmin ⇒ error  
          * CertificationCenter ⇒ error  
          * NodeAdmin ⇒ error  

* Parameters:  
     * subject (Subject) - string  
          * Positive:  
               * string matches the format  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * subject_key_id (Subject Key ID) - string  
          * Positive:  
               * string matches the format  
                    * e.g., 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * serial-number (Serial Number) - optional(string)  
          * Positive:  
               * empty value  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * length > MAX  
     * info (Information/Notes) - optional(string)  
          * Positive:  
               * empty value  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * length > MAX ⇒ MAX = 4096 characters  
     * time (Proposal Time) - optional(int64)  
          * Positive:  
               * default value ⇒ current time by default  
               * empty value  
               * integer value format  
          * Negative:  
               * length > MAX ⇒ MAX = 9 223 372 036 854 775 807  

### REST API  
REST API command: POST /cosmos/tx/v1beta1/txs: [MsgApproveRevokePAA](https://github.com/zigbee-alliance/distributed-compliance-ledger)

* REST API command send  
     * Valid command  
          * correct HTTP method  
          * request is authorized  
          * uses valid credentials/role  
          * query works for the following certificates  
               * PAA  
     * Invalid command  
          * incorrect request  
          * server side error  

* Command result  
     * APPROVE_REVOKE_PAA command completed successfully ⇒ approves the revocation of the given PAA (self-signed root certificate) by a Trustee  
          * PROPOSE_ADD_PAA command completed successfully  
          * APPROVE_ADD_PAA command completed successfully  
          * PROPOSE_REVOKE_PAA command completed successfully  
          * Number of required approvals greater than 2/3 of Trustees ⇒ revocation is applied  
          * Number of required approvals equal 2/3 of Trustees ⇒ revocation is applied  
          * the proposal to revoke a root certificate with the provided subject and subject_key_id, submitted first  
          * the proposed certificate revocation hasn't been approved by the signer yet  
     * APPROVE_REVOKE_PAA command failed ⇒ does not approve the revocation of the given PAA (self-signed root certificate) by a Trustee  
          * PROPOSE_ADD_PAA command failed  
          * APPROVE_ADD_PAA command failed  
          * PROPOSE_REVOKE_PAA command failed  
          * Number of required approvals less than 2/3 of Trustees ⇒ revocation is not applied  
          * the proposal to revoke a root certificate with the provided subject and subject_key_id, not submitted first  
          * the proposed certificate revocation has been approved by the signer  

* Role (Who can send)  
     * Positive:  
          * Trustee  
     * Negative:  
          * Vendor ⇒ error  
          * VendorAdmin ⇒ error  
          * CertificationCenter ⇒ error  
          * NodeAdmin ⇒ error  

* Parameters:  
     * subject (Subject) - string  
          * Positive:  
               * string matches the format  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * subject_key_id (Subject Key ID) - string  
          * Positive:  
               * string matches the format  
                    * e.g., 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * serial-number (Serial Number) - optional(string)  
          * Positive:  
               * empty value  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * length > MAX  
     * info (Information/Notes) - optional(string)  
          * Positive:  
               * empty value  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * length > MAX ⇒ MAX = 4096 characters  
     * time (Proposal Time) - optional(int64)  
          * Positive:  
               * default value ⇒ current time by default  
               * empty value  
               * integer value format  
          * Negative:  
               * length > MAX ⇒ MAX = 9 223 372 036 854 775 807  


## ASSIGN_VID_TO_PAA

### CLI command  
CLI command: `dcld tx paa assign-vid --subject=<string> --subject-key-id=<string> --vid=<uint16> --from=<account>`

* CLI command send  
     * Valid command  
          * command exists/relevant  
          * query works for the following certificates  
               * PAA  
     * Invalid command  
          * access is denied to execute command  
          * incorrect command syntax  

* Command result  
     * ASSIGN_VID_TO_PAA command completed successfully ⇒ assigns a Vendor ID (VID) to non-VID scoped PAAs (self-signed root certificate) already present on the ledger  
          * PROPOSE_ADD_PAA command completed successfully  
          * APPROVE_ADD_PAA command completed successfully  
          * PAA Certificate with the provided subject and subject_key_id exist in the ledger  
          * the PAA is a VID scoped one  
          * the vid field equal to the VID value in the PAA's subject  
     * ASSIGN_VID_TO_PAA command failed ⇒ does not assign a Vendor ID (VID) to non-VID scoped PAAs (self-signed root certificate) already present on the ledger  
          * PROPOSE_ADD_PAA command failed  
          * APPROVE_ADD_PAA command failed  
          * PAA Certificate with the provided subject and subject_key_id not exist in the ledger  
          * the PAA is a VID scoped one  
          * the vid field not equal to the VID value in the PAA's subject  

* Role (Who can send)  
     * Positive:  
          * VendorAdmin  
     * Negative:  
          * Trustee ⇒ error  
          * Vendor ⇒ error  
          * CertificationCenter ⇒ error  
          * NodeAdmin ⇒ error  

* Parameters:  
     * subject (Subject) - string  
          * Positive:  
               * string matches the format  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * subject_key_id (Subject Key ID) - string  
          * Positive:  
               * string matches the format  
                    * e.g., 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * vid (Vendor ID) - uint16  
          * Positive:  
               * value exists  
               * value > 0  
               * integer value format  
               * Vendor ID value = vid field in the VID-scoped PAA certificate  
          * Negative:  
               * empty value  
               * value =< 0  
               * string value format  
               * length > MAX ⇒ MAX = 65535  
               * nonexistent ID  

### REST API  
REST API command: POST /cosmos/tx/v1beta1/txs: [MsgAssignVidToPAA](https://github.com/zigbee-alliance/distributed-compliance-ledger)

* REST API command send  
     * Valid command  
          * correct HTTP method  
          * request is authorized  
          * uses valid credentials/role  
          * query works for the following certificates  
               * PAA  
     * Invalid command  
          * incorrect request  
          * server side error  

* Command result  
     * ASSIGN_VID_TO_PAA command completed successfully ⇒ assigns a Vendor ID (VID) to non-VID scoped PAAs (self-signed root certificate) already present on the ledger  
          * PROPOSE_ADD_PAA command completed successfully  
          * APPROVE_ADD_PAA command completed successfully  
          * PAA Certificate with the provided subject and subject_key_id exist in the ledger  
          * the PAA is a VID scoped one  
          * the vid field equal to the VID value in the PAA's subject  
     * ASSIGN_VID_TO_PAA command failed ⇒ does not assign a Vendor ID (VID) to non-VID scoped PAAs (self-signed root certificate) already present on the ledger  
          * PROPOSE_ADD_PAA command failed  
          * APPROVE_ADD_PAA command failed  
          * PAA Certificate with the provided subject and subject_key_id not exist in the ledger  
          * the PAA is a VID scoped one  
          * the vid field not equal to the VID value in the PAA's subject  

* Role (Who can send)  
     * Positive:  
          * VendorAdmin  
     * Negative:  
          * Trustee ⇒ error  
          * Vendor ⇒ error  
          * CertificationCenter ⇒ error  
          * NodeAdmin ⇒ error  

* Parameters:  
     * subject (Subject) - string  
          * Positive:  
               * string matches the format  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * subject_key_id (Subject Key ID) - string  
          * Positive:  
               * string matches the format  
                    * e.g., 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * vid (Vendor ID) - uint16  
          * Positive:  
               * value exists  
               * value > 0  
               * integer value format  
               * Vendor ID value = vid field in the VID-scoped PAA certificate  
          * Negative:  
               * empty value  
               * value =< 0  
               * string value format  
               * length > MAX ⇒ MAX = 65535  
               * nonexistent ID

## ADD_REVOCATION_DISTRIBUTION_POINT

### CLI command  
CLI command: `dcld tx pki add-revocation-distribution-point --vid=<uint16> --label=<string> --issuerSubjectKeyID=<string> --crlSignerCertificate=<string> [--crlSignerDelegator=<string>] --dataUrl=<string> [--dataFileSize=<uint64>] [--dataDigest=<string>] [--dataDigestType=<uint32>] --revocationType=<uint32> [--schemaVersion=<uint16>] --from=<account>`

* CLI command send  
     * Valid command  
          * command exists/relevant  
     * Invalid command  
          * access is denied to execute command  
          * incorrect command syntax  

* Command result  
     * ADD_REVOCATION_DISTRIBUTION_POINT command completed successfully ⇒ publishes a PKI Revocation distribution endpoint (such as RFC5280 Certificate Revocation List) owned by the Vendor  
          * PROPOSE_ADD_PAA command completed successfully  
          * APPROVE_ADD_PAA command completed successfully  
          * crlSignerCertificate is a PAA (root certificate)  
          * crlSignerCertificate is present on DCL  
          * crlSignerCertificate is a PAI (intermediate certificate)  
          * crlSignerCertificate chained back to a valid PAA (root certificate) present on DCL  
          * crlSignerCertificate is a delegated by PAA  
     * ADD_REVOCATION_DISTRIBUTION_POINT command failed ⇒ does not publish a PKI Revocation distribution endpoint (such as RFC5280 Certificate Revocation List) owned by the Vendor  
          * PROPOSE_ADD_PAA command failed  
          * APPROVE_ADD_PAA command failed  
          * crlSignerCertificate is not present on DCL  
          * crlSignerCertificate is a PAI (intermediate certificate)  
          * crlSignerCertificate is not chained back to a valid PAA (root certificate) present on DCL  
          * crlSignerCertificate is a delegated by PAA  
          * crlSignerCertificate is not chained back to a valid PAA (root certificate) present on DCL  

* Role (Who can send)  
     * Positive:  
          * Vendor  
               * vid field in the transaction (VendorID) equals the Vendor account's VID  
               * VID-scoped PAAs and PAIs: vid field in the CRLSignerCertificate's subject equals the Vendor account's VID  
               * Non-VID scoped PAAs: vid associated with the corresponding PAA on the ledger equals the Vendor account's VID  
     * Negative:  
          * Trustee ⇒ error  
          * VendorAdmin ⇒ error  
          * CertificationCenter ⇒ error  
          * NodeAdmin ⇒ error  

* Parameters:  
     * vid (Vendor ID) - uint16  
          * Positive:  
               * value exists  
               * value > 0  
               * integer value format  
               * matches Vendor account's VID and CRLSignerCertificate's subject  
          * Negative:  
               * empty value  
               * value =< 0  
               * string value format  
               * length > MAX (65535)  
               * nonexistent ID  
     * pid (Product ID) - optional(uint16)  
          * Positive:  
               * unique combination  
               * value > 0  
               * integer value format  
          * Negative:  
               * empty value  
               * value =< 0  
               * string value format  
               * length > MAX (65535)  
               * field not empty if IsPAA is true  
               * value ≠ pid field in CRLSignerCertificate  
     * isPAA (Is PAA) - bool  
          * Positive:  
               * TRUE (-1) ⇒ relates to a PAA  
               * FALSE (0)  
          * Negative:  
               * empty value  
               * not a boolean  
     * label (Label) - string  
          * Positive:  
               * value exists  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * length > MAX  
     * crlSignerCertificate - string  
          * Positive:  
               * value exists  
               * contains PEM string or path to file  
               * delegated certificate by a PAI must use crlSignerDelegator  
               * not delegated ⇒ field is direct  
          * Negative:  
               * empty value  
               * length > MAX  
     * crlSignerDelegator - optional(string)  
          * Positive:  
               * value exists  
               * contains PEM string or file  
               * delegated certificate by a PAI ⇒ must be chained to an approved cert on ledger  
               * not delegated ⇒ field can be omitted  
          * Negative:  
               * length > MAX  
     * issuerSubjectKeyID - string  
          * Positive:  
               * unique to PAA/PAI  
               * text format, MIN < length < MAX  
               * even number of uppercase hex characters  
               * e.g., 5A880E6C3653D07FB08971A3F473790930E62BDB  
          * Negative:  
               * contains whitespace or non-hex chars  
               * length > MAX  
               * empty value  
     * dataUrl - string  
          * Positive:  
               * unique for VendorID and IssuerSubjectKeyID  
               * text format, MIN < length < MAX  
               * starts with http/https  
          * Negative:  
               * format does not match RevocationType spec  
               * empty value  
               * length > MAX  
     * dataFileSize - optional(uint64)  
          * Positive:  
               * value >= 0  
               * integer value format  
          * Negative:  
               * empty value  
               * length > MAX (18,446,744,073,709,551,615)  
     * dataDigest - optional(string)  
          * Required if dataFileSize present  
          * Positive:  
               * matches format  
               * text format  
               * MIN < length < MAX  
               * empty value  
          * Negative:  
               * nonexistent  
               * length > MAX  
     * dataDigestType - optional(uint32)  
          * Required if dataDigest present  
          * Positive:  
               * value exists  
               * value > 0  
               * empty value  
               * correct format  
          * Negative:  
               * value =< 0  
               * string format  
               * length > MAX (4,294,967,295)  
     * revocationType (Revocation Type) - uint32  
          * Positive:  
               * value exists  
               * value >= 0  
               * integer format  
               * supported: 1 - RFC5280 CRL  
          * Negative:  
               * empty  
               * nonexistent  
               * length > MAX (4,294,967,295)  
     * schemaVersion - optional(uint16)  
          * Positive:  
               * value = 0  
               * integer format  
               * empty value  
          * Negative:  
               * length > MAX (65535)

### REST API  
REST API command: POST /cosmos/tx/v1beta1/txs: [MsgAddRevocationDistributionPoint](https://github.com/zigbee-alliance/distributed-compliance-ledger)

* REST API command send  
     * Valid command  
          * correct HTTP method  
          * request is authorized  
          * uses valid credentials/role  
     * Invalid command  
          * incorrect request  
          * server side error  

* Command result  
     * ADD_REVOCATION_DISTRIBUTION_POINT command completed successfully ⇒ publishes a PKI Revocation distribution endpoint (such as RFC5280 Certificate Revocation List) owned by the Vendor  
          * PROPOSE_ADD_PAA command completed successfully  
          * APPROVE_ADD_PAA command completed successfully  
          * crlSignerCertificate is a PAA (root certificate)  
          * crlSignerCertificate is present on DCL  
          * crlSignerCertificate is a PAI (intermediate certificate)  
          * crlSignerCertificate chained back to a valid PAA (root certificate) present on DCL  
          * crlSignerCertificate is a delegated by PAA  
     * ADD_REVOCATION_DISTRIBUTION_POINT command failed ⇒ does not publish a PKI Revocation distribution endpoint (such as RFC5280 Certificate Revocation List) owned by the Vendor  
          * PROPOSE_ADD_PAA command failed  
          * APPROVE_ADD_PAA command failed  
          * crlSignerCertificate is not present on DCL  
          * crlSignerCertificate is not chained back to a valid PAA  
          * crlSignerCertificate is a delegated by PAA  
          * crlSignerCertificate is not chained back to a valid PAA (root certificate) present on DCL  

* Role (Who can send)  
     * Positive:  
          * Vendor (conditions as per CLI)  
     * Negative:  
          * Trustee ⇒ error  
          * VendorAdmin ⇒ error  
          * CertificationCenter ⇒ error  
          * NodeAdmin ⇒ error

* Parameters:  
     * vid (Vendor ID) - uint16  
          * Positive:  
               * value exists  
               * value > 0  
               * integer value format  
               * matches Vendor account's VID and CRLSignerCertificate's subject  
          * Negative:  
               * empty value  
               * value =< 0  
               * string value format  
               * length > MAX (65535)  
               * nonexistent ID  
     * pid (Product ID) - optional(uint16)  
          * Positive:  
               * unique combination  
               * value > 0  
               * integer value format  
          * Negative:  
               * empty value  
               * value =< 0  
               * string value format  
               * length > MAX (65535)  
               * field not empty if IsPAA is true  
               * value ≠ pid field in CRLSignerCertificate  
     * isPAA (Is PAA) - bool  
          * Positive:  
               * TRUE (-1) ⇒ relates to a PAA  
               * FALSE (0)  
          * Negative:  
               * empty value  
               * not a boolean  
     * label (Label) - string  
          * Positive:  
               * value exists  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * length > MAX  
     * crlSignerCertificate - string  
          * Positive:  
               * value exists  
               * contains PEM string or path to file  
               * delegated certificate by a PAI must use crlSignerDelegator  
               * not delegated ⇒ field is direct  
          * Negative:  
               * empty value  
               * length > MAX  
     * crlSignerDelegator - optional(string)  
          * Positive:  
               * value exists  
               * contains PEM string or file  
               * delegated certificate by a PAI ⇒ must be chained to an approved cert on ledger  
               * not delegated ⇒ field can be omitted  
          * Negative:  
               * length > MAX  
     * issuerSubjectKeyID - string  
          * Positive:  
               * unique to PAA/PAI  
               * text format, MIN < length < MAX  
               * even number of uppercase hex characters  
               * e.g., 5A880E6C3653D07FB08971A3F473790930E62BDB  
          * Negative:  
               * contains whitespace or non-hex chars  
               * length > MAX  
               * empty value  
     * dataUrl - string  
          * Positive:  
               * unique for VendorID and IssuerSubjectKeyID  
               * text format, MIN < length < MAX  
               * starts with http/https  
          * Negative:  
               * format does not match RevocationType spec  
               * empty value  
               * length > MAX  
     * dataFileSize - optional(uint64)  
          * Positive:  
               * value >= 0  
               * integer value format  
          * Negative:  
               * empty value  
               * length > MAX (18,446,744,073,709,551,615)  
     * dataDigest - optional(string)  
          * Required if dataFileSize present  
          * Positive:  
               * matches format  
               * text format  
               * MIN < length < MAX  
               * empty value  
          * Negative:  
               * nonexistent  
               * length > MAX  
     * dataDigestType - optional(uint32)  
          * Required if dataDigest present  
          * Positive:  
               * value exists  
               * value > 0  
               * empty value  
               * correct format  
          * Negative:  
               * value =< 0  
               * string format  
               * length > MAX (4,294,967,295)  
     * revocationType (Revocation Type) - uint32  
          * Positive:  
               * value exists  
               * value >= 0  
               * integer format  
               * supported: 1 - RFC5280 CRL  
          * Negative:  
               * empty  
               * nonexistent  
               * length > MAX (4,294,967,295)  
     * schemaVersion - optional(uint16)  
          * Positive:  
               * value = 0  
               * integer format  
               * empty value  
          * Negative:  
               * length > MAX (65535)

## UPDATE_REVOCATION_DISTRIBUTION_POINT

### CLI command  
CLI command: `dcld tx pki update-revocation-distribution-point --vid=<uint16> --label=<string> --issuerSubjectKeyID=<string> [--crlSignerCertificate=<string>] [--crlSignerDelegator=<string>] --dataUrl=<string> [--dataFileSize=<uint64>] [--dataDigest=<string>] [--dataDigestType=<uint32>] [--schemaVersion=<uint16>] --from=<account>`

* CLI command send  
     * Valid command  
          * command exists/relevant  
     * Invalid command  
          * access is denied to execute command  
          * incorrect command syntax  

* Command result  
     * UPDATE_REVOCATION_DISTRIBUTION_POINT command completed successfully ⇒ updates an existing PKI Revocation distribution endpoint (such as RFC5280 Certificate Revocation List) owned by the Vendor  
          * PROPOSE_ADD_PAA command completed successfully  
          * APPROVE_ADD_PAA command completed successfully  
          * ADD_REVOCATION_DISTRIBUTION_POINT command completed successfully  
     * UPDATE_REVOCATION_DISTRIBUTION_POINT command failed ⇒ does not update an existing PKI Revocation distribution endpoint  
          * PROPOSE_ADD_PAA command failed  
          * APPROVE_ADD_PAA command failed  
          * ADD_REVOCATION_DISTRIBUTION_POINT command failed  

* Role (Who can send)  
     * Positive:  
          * Vendor  
               * vid field in transaction equals Vendor account's VID  
               * VID-scoped certs: subject VID equals Vendor account's VID  
               * Non-VID scoped PAAs: associated VID matches Vendor account's VID  
     * Negative:  
          * Trustee ⇒ error  
          * VendorAdmin ⇒ error  
          * CertificationCenter ⇒ error  
          * NodeAdmin ⇒ error  

* Parameters:  
     * vid (Vendor ID) - uint16  
          * Positive:  
               * value exists  
               * value > 0  
               * integer format  
               * matches account and certificate VID  
          * Negative:  
               * empty value  
               * value ≤ 0  
               * string format  
               * length > 65535  
               * nonexistent ID  
     * label (Label) - string  
          * Positive:  
               * value exists  
               * proper text format, MIN < length < MAX  
          * Negative:  
               * empty value  
               * length > MAX  
     * issuerSubjectKeyID - string  
          * Positive:  
               * unique for PAA/PAI  
               * text format, even-length uppercase hex (e.g. 5A880E6C...)  
               * MIN < length < MAX  
          * Negative:  
               * whitespace, non-hex chars  
               * length > MAX  
               * empty value  
     * crlSignerCertificate - optional(string)  
          * Positive:  
               * exists or empty  
               * PEM string or file path  
               * proper certificate type  
               * delegated certs require crlSignerDelegator  
          * Negative:  
               * length > MAX  
     * crlSignerDelegator - optional(string)  
          * Positive:  
               * exists or empty  
               * PEM string or file path  
               * if delegated cert, must chain back to approved cert  
          * Negative:  
               * length > MAX  
     * dataUrl - string  
          * Positive:  
               * unique for VendorID + IssuerSubjectKeyID  
               * starts with http/https  
               * MIN < length < MAX  
          * Negative:  
               * invalid format  
               * empty value  
               * length > MAX  
     * dataFileSize - optional(uint64)  
          * Positive:  
               * value ≥ 0  
               * integer format  
          * Negative:  
               * empty value  
               * length > MAX (18,446,744,073,709,551,615)  
     * dataDigest - optional(string)  
          * Required if dataFileSize is present and revocationType ≠ 1  
          * Positive:  
               * matches ISO datetime format (e.g. 2019-10-12T...)  
               * text format, MIN < length < MAX  
               * empty value allowed  
          * Negative:  
               * nonexistent value  
               * length > MAX  
     * dataDigestType - optional(uint32)  
          * Required if dataDigest present  
          * Positive:  
               * value exists and > 0  
               * empty value  
          * Negative:  
               * value ≤ 0  
               * string format  
               * length > MAX (4,294,967,295)  
     * schemaVersion - optional(uint16)  
          * Positive:  
               * value = 0  
               * empty value  
               * integer format  
          * Negative:  
               * length > MAX (65535)

### REST API  
REST API command: POST /cosmos/tx/v1beta1/txs: [MsgUpdateRevocationDistributionPoint](https://github.com/zigbee-alliance/distributed-compliance-ledger)

* REST API command send  
     * Valid command  
          * correct HTTP method  
          * request is authorized  
          * uses valid credentials/role  
     * Invalid command  
          * incorrect request  
          * server side error  

* Command result  
     * UPDATE_REVOCATION_DISTRIBUTION_POINT command completed successfully ⇒ updates an existing PKI Revocation distribution endpoint (such as RFC5280 Certificate Revocation List) owned by the Vendor  
          * PROPOSE_ADD_PAA command completed successfully  
          * APPROVE_ADD_PAA command completed successfully  
          * ADD_REVOCATION_DISTRIBUTION_POINT command completed successfully  
     * UPDATE_REVOCATION_DISTRIBUTION_POINT command failed ⇒ does not update an existing PKI Revocation distribution endpoint  
          * PROPOSE_ADD_PAA command failed  
          * APPROVE_ADD_PAA command failed  
          * ADD_REVOCATION_DISTRIBUTION_POINT command failed  

* Role (Who can send)  
     * Positive:  
          * Vendor  
     * Negative:  
          * Trustee ⇒ error  
          * VendorAdmin ⇒ error  
          * CertificationCenter ⇒ error  
          * NodeAdmin ⇒ error

* Parameters:  
     * vid (Vendor ID) - uint16  
          * Positive:  
               * value exists  
               * value > 0  
               * integer format  
               * matches account and certificate VID  
          * Negative:  
               * empty value  
               * value ≤ 0  
               * string format  
               * length > 65535  
               * nonexistent ID  
     * label (Label) - string  
          * Positive:  
               * value exists  
               * proper text format, MIN < length < MAX  
          * Negative:  
               * empty value  
               * length > MAX  
     * issuerSubjectKeyID - string  
          * Positive:  
               * unique for PAA/PAI  
               * text format, even-length uppercase hex (e.g. 5A880E6C...)  
               * MIN < length < MAX  
          * Negative:  
               * whitespace, non-hex chars  
               * length > MAX  
               * empty value  
     * crlSignerCertificate - optional(string)  
          * Positive:  
               * exists or empty  
               * PEM string or file path  
               * proper certificate type  
               * delegated certs require crlSignerDelegator  
          * Negative:  
               * length > MAX  
     * crlSignerDelegator - optional(string)  
          * Positive:  
               * exists or empty  
               * PEM string or file path  
               * if delegated cert, must chain back to approved cert  
          * Negative:  
               * length > MAX  
     * dataUrl - string  
          * Positive:  
               * unique for VendorID + IssuerSubjectKeyID  
               * starts with http/https  
               * MIN < length < MAX  
          * Negative:  
               * invalid format  
               * empty value  
               * length > MAX  
     * dataFileSize - optional(uint64)  
          * Positive:  
               * value ≥ 0  
               * integer format  
          * Negative:  
               * empty value  
               * length > MAX (18,446,744,073,709,551,615)  
     * dataDigest - optional(string)  
          * Required if dataFileSize is present and revocationType ≠ 1  
          * Positive:  
               * matches ISO datetime format (e.g. 2019-10-12T...)  
               * text format, MIN < length < MAX  
               * empty value allowed  
          * Negative:  
               * nonexistent value  
               * length > MAX  
     * dataDigestType - optional(uint32)  
          * Required if dataDigest present  
          * Positive:  
               * value exists and > 0  
               * empty value  
          * Negative:  
               * value ≤ 0  
               * string format  
               * length > MAX (4,294,967,295)  
     * schemaVersion - optional(uint16)  
          * Positive:  
               * value = 0  
               * empty value  
               * integer format  
          * Negative:  
               * length > MAX (65535)

## DELETE_REVOCATION_DISTRIBUTION_POINT

### CLI command  
CLI command: `dcld tx pki delete-revocation-distribution-point --vid=<uint16> --label=<string> --issuerSubjectKeyID=<string> --from=<account>`

* CLI command send  
     * Valid command  
          * command exists/relevant  
     * Invalid command  
          * access is denied to execute command  
          * incorrect command syntax  

* Command result  
     * DELETE_REVOCATION_DISTRIBUTION_POINT command completed successfully ⇒ deletes a PKI Revocation distribution endpoint (such as RFC5280 Certificate Revocation List) owned by the Vendor  
          * PROPOSE_ADD_PAA command completed successfully  
          * APPROVE_ADD_PAA command completed successfully  
          * ADD_REVOCATION_DISTRIBUTION_POINT command completed successfully  
     * DELETE_REVOCATION_DISTRIBUTION_POINT command failed ⇒ does not delete a PKI Revocation distribution endpoint (such as RFC5280 Certificate Revocation List) owned by the Vendor  
          * PROPOSE_ADD_PAA command failed  
          * APPROVE_ADD_PAA command failed  
          * ADD_REVOCATION_DISTRIBUTION_POINT command failed  

* Role (Who can send)  
     * Positive:  
          * Vendor  
               * vid field in transaction equals Vendor account's VID  
               * VID-scoped certs: subject VID equals Vendor account's VID  
               * Non-VID scoped PAAs: associated VID matches Vendor account's VID  
     * Negative:  
          * Trustee ⇒ error  
          * VendorAdmin ⇒ error  
          * CertificationCenter ⇒ error  
          * NodeAdmin ⇒ error  

* Parameters:  
     * vid (Vendor ID) - uint16  
          * Positive:  
               * value exists  
               * value > 0  
               * integer format  
               * matches account and certificate VID  
          * Negative:  
               * empty value  
               * value ≤ 0  
               * string format  
               * length > 65535  
               * nonexistent ID  
     * label (Label) - string  
          * Positive:  
               * value exists  
               * proper text format, MIN < length < MAX  
          * Negative:  
               * empty value  
               * length > MAX  
     * issuerSubjectKeyID - string  
          * Positive:  
               * unique for PAA/PAI  
               * text format, even-length uppercase hex (e.g. 5A880E6C...)  
               * MIN < length < MAX  
               * must be provided using crlSignerDelegator field if delegated by PAI  
          * Negative:  
               * whitespace, non-hex chars  
               * length > MAX  
               * empty value

### REST API  
REST API command: POST /cosmos/tx/v1beta1/txs: [MsgDeleteRevocationDistributionPoint](https://github.com/zigbee-alliance/distributed-compliance-ledger)

* REST API command send  
     * Valid command  
          * correct HTTP method  
          * request is authorized  
          * uses valid credentials/role  
     * Invalid command  
          * incorrect request  
          * server side error  

* Command result  
     * DELETE_REVOCATION_DISTRIBUTION_POINT command completed successfully ⇒ deletes a PKI Revocation distribution endpoint (such as RFC5280 Certificate Revocation List) owned by the Vendor  
          * PROPOSE_ADD_PAA command completed successfully  
          * APPROVE_ADD_PAA command completed successfully  
          * ADD_REVOCATION_DISTRIBUTION_POINT command completed successfully  
     * DELETE_REVOCATION_DISTRIBUTION_POINT command failed ⇒ does not delete a PKI Revocation distribution endpoint  
          * PROPOSE_ADD_PAA command failed  
          * APPROVE_ADD_PAA command failed  
          * ADD_REVOCATION_DISTRIBUTION_POINT command failed  

* Role (Who can send)  
     * Positive:  
          * Vendor (see CLI role conditions)  
     * Negative:  
          * Trustee ⇒ error  
          * VendorAdmin ⇒ error  
          * CertificationCenter ⇒ error  
          * NodeAdmin ⇒ error

* Parameters:  
     * vid (Vendor ID) - uint16  
          * Positive:  
               * value exists  
               * value > 0  
               * integer format  
               * matches account and certificate VID  
          * Negative:  
               * empty value  
               * value ≤ 0  
               * string format  
               * length > 65535  
               * nonexistent ID  
     * label (Label) - string  
          * Positive:  
               * value exists  
               * proper text format, MIN < length < MAX  
          * Negative:  
               * empty value  
               * length > MAX  
     * issuerSubjectKeyID - string  
          * Positive:  
               * unique for PAA/PAI  
               * text format, even-length uppercase hex (e.g. 5A880E6C...)  
               * MIN < length < MAX  
               * must be provided using crlSignerDelegator field if delegated by PAI  
          * Negative:  
               * whitespace, non-hex chars  
               * length > MAX  
               * empty value

## ADD_PAI

### CLI command  
CLI command: `dcld tx pki add-pai --cert=<string> [--certificate-schema-version=<uint16>] --from=<account>`

* CLI command send  
     * Valid command  
          * command exists/relevant  
          * query works for the following certificates  
               * PAI  
     * Invalid command  
          * access is denied to execute command  
          * incorrect command syntax  

* Command result  
     * ADD_PAI command completed successfully ⇒ adds a PAI (intermediate certificate) signed by a chain of certificates already present on the ledger  
          * PROPOSE_ADD_PAA command completed successfully  
          * APPROVE_ADD_PAA command completed successfully  
          * provided certificate is not root  
               * Issuer ≠ Subject  
               * Authority Key Identifier ≠ Subject Key Identifier  
          * no existing certificate with the same <Issuer>:<Serial Number>  
          * no certificate with the same <Subject>:<Subject Key ID>  
          * existing certificate is not a NOC  
          * sender's VID matches existing certificate's owner VID  
          * signature and expiration date are valid  
          * parent certificate already stored and valid chain to root exists  
          * root is VID scoped, provided certificate is also VID scoped  
               * VIDs in root and provided cert match and equal to sender VID  
          * root is not VID scoped but has associated VID  
               * provided cert is VID or non-VID scoped  
               * if VID scoped, its VID matches root's associated VID and sender VID  
          * multiple certificates refer to the same <Subject>:<Subject Key ID>  
     * ADD_PAI command failed ⇒ does not add the certificate  
          * PROPOSE_ADD_PAA command failed  
          * APPROVE_ADD_PAA command failed  
          * certificate is root (Issuer == Subject)  
          * Authority Key Identifier == Subject Key Identifier  
          * existing certificate with same <Issuer>:<Serial Number>  
          * duplicate <Subject>:<Subject Key ID>  
          * certificate is a NOC  
          * sender's VID does not match certificate owner's VID  
          * signature or expiration invalid  
               * one or both invalid  
          * parent cert not stored or chain to root cannot be built  
          * root is VID scoped, provided cert is not VID scoped  
               * VIDs in root and provided cert do not match  
               * VIDs do not match sender's VID  
          * root is non-VID scoped without associated VID  
          * provided cert is not VID scoped  
               * if VID scoped, its VID does not match root's associated VID and sender VID  

* Role (Who can send)  
     * Positive:  
          * Vendor  
     * Negative:  
          * Trustee ⇒ error  
          * VendorAdmin ⇒ error  
          * CertificationCenter ⇒ error  
          * NodeAdmin ⇒ error  

* Parameters:  
     * cert (Certificate) - string  
          * Positive:  
               * value exists  
               * contains PEM string or valid file path  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * certificate-schema-version - optional(uint16)  
          * Positive:  
               * value = 0  
               * integer value format  
               * empty value  
          * Negative:  
               * length > MAX (65535)

### REST API  
REST API command: POST /cosmos/tx/v1beta1/txs: [MsgAddPAI](https://github.com/zigbee-alliance/distributed-compliance-ledger)

* REST API command send  
     * Valid command  
          * correct HTTP method  
          * request is authorized  
          * uses valid credentials/role  
          * query works for the following certificates  
               * PAI  
     * Invalid command  
          * incorrect request  
          * server side error  

* Command result  
     * ADD_PAI command completed successfully ⇒ adds a PAI (intermediate certificate) signed by a chain of certificates already present on the ledger  
          * PROPOSE_ADD_PAA command completed successfully  
          * APPROVE_ADD_PAA command completed successfully  
          * provided certificate is not root  
               * Issuer ≠ Subject  
               * Authority Key Identifier ≠ Subject Key Identifier  
          * no existing certificate with the same <Issuer>:<Serial Number>  
          * no certificate with the same <Subject>:<Subject Key ID>  
          * existing certificate is not a NOC  
          * sender's VID matches existing certificate's owner VID  
          * signature and expiration date are valid  
          * parent certificate already stored and valid chain to root exists  
          * root is VID scoped, provided certificate is also VID scoped  
               * VIDs in root and provided cert match and equal to sender VID  
          * root is not VID scoped but has associated VID  
               * provided cert is VID or non-VID scoped  
               * if VID scoped, its VID matches root's associated VID and sender VID  
          * multiple certificates refer to the same <Subject>:<Subject Key ID>  
     * ADD_PAI command failed ⇒ does not add the certificate  
          * PROPOSE_ADD_PAA command failed  
          * APPROVE_ADD_PAA command failed  
          * certificate is root (Issuer == Subject)  
          * Authority Key Identifier == Subject Key Identifier  
          * existing certificate with same <Issuer>:<Serial Number>  
          * duplicate <Subject>:<Subject Key ID>  
          * certificate is a NOC  
          * sender's VID does not match certificate owner's VID  
          * signature or expiration invalid  
               * one or both invalid  
          * parent cert not stored or chain to root cannot be built  
          * root is VID scoped, provided cert is not VID scoped  
               * VIDs in root and provided cert do not match  
               * VIDs do not match sender's VID  
          * root is non-VID scoped without associated VID  
          * provided cert is not VID scoped  
               * if VID scoped, its VID does not match root's associated VID and sender VID  

* Role (Who can send)  
     * Positive:  
          * Vendor  
     * Negative:  
          * Trustee ⇒ error  
          * VendorAdmin ⇒ error  
          * CertificationCenter ⇒ error  
          * NodeAdmin ⇒ error

* Parameters:  
     * cert (Certificate) - string  
          * Positive:  
               * value exists  
               * contains PEM string or valid file path  
               * text value format  
               * MIN < length < MAX  
          * Negative:  
               * empty value  
               * nonexistent value  
               * length > MAX  
     * certificate-schema-version - optional(uint16)  
          * Positive:  
               * value = 0  
               * integer value format  
               * empty value  
          * Negative:  
               * length > MAX (65535)

## REVOKE_PAI

### CLI command  
CLI command: `dcld tx pki revoke-pai --subject=<string> --subject-key-id=<string> [--serial-number=<string>] [--revoke-child=<bool>] [--info=<string>] [--time=<int64>] --from=<account>`

* CLI command send  
     * Valid command  
          * command exists/relevant  
          * query works for the following certificates  
               * PAI  
     * Invalid command  
          * access is denied to execute command  
          * incorrect command syntax  

* Command result  
     * REVOKE_PAI command completed successfully ⇒ revokes the given PAI (intermediate certificate)  
          * PROPOSE_ADD_PAA command completed successfully  
          * APPROVE_ADD_PAA command completed successfully  
          * ADD_PAI command completed successfully  
          * revoke-child = True ⇒ all certificates signed by the revoked certificate will be revoked  
          * revoke-child = False ⇒ certificates in the chain will not be revoked  
          * PAI certificate with the provided subject and subject_key_id exists on the ledger  
     * REVOKE_PAI command failed ⇒ does not revoke the given PAI (intermediate certificate)  
          * PROPOSE_ADD_PAA command failed  
          * APPROVE_ADD_PAA command failed  
          * ADD_PAI command failed  
          * PAI certificate with the provided subject and subject_key_id does not exist  

* Role (Who can send)  
     * Positive:  
          * Vendor  
               * sender's VID matches the VID of the certificate owner  
     * Negative:  
          * Trustee ⇒ error  
          * VendorAdmin ⇒ error  
          * CertificationCenter ⇒ error  
          * NodeAdmin ⇒ error  

* Parameters:  
     * subject (Subject) - string  
          * Positive:  
               * string matches format  
               * valid text format  
               * MIN < length < MAX  
          * Negative:  
               * empty or nonexistent value  
               * length > MAX  
     * subject_key_id (Subject Key ID) - string  
          * Positive:  
               * string matches format  
               * valid text format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`  
               * MIN < length < MAX  
          * Negative:  
               * empty or nonexistent value  
               * length > MAX  
     * serial-number (Serial Number) - optional(string)  
          * Positive:  
               * empty or valid text format  
               * MIN < length < MAX  
          * Negative:  
               * length > MAX  
     * revoke-child (Revoke Child) - optional(bool)  
          * Positive:  
               * empty value or valid state  
               * TRUE (-1) ⇒ all child certs revoked  
               * FALSE (0) ⇒ default  
          * Negative:  
               * value is not a boolean  
     * info (Information/Notes) - optional(string)  
          * Positive:  
               * empty or valid text format  
               * MIN < length < MAX  
          * Negative:  
               * length > MAX (4096 characters)  
     * time (Proposal Time) - optional(int64)  
          * Positive:  
               * default value ⇒ current time  
               * empty value or valid integer format  
          * Negative:  
               * length > MAX (9,223,372,036,854,775,807)

### REST API  
REST API command: POST /cosmos/tx/v1beta1/txs: [MsgRevokePAI](https://github.com/zigbee-alliance/distributed-compliance-ledger)

* REST API command send  
     * Valid command  
          * correct HTTP method  
          * request is authorized  
          * uses valid credentials/role  
          * query works for the following certificates  
               * PAI  
     * Invalid command  
          * incorrect request  
          * server side error  

* Command result  
     * REVOKE_PAI command completed successfully ⇒ revokes the given PAI (intermediate certificate)  
          * PROPOSE_ADD_PAA command completed successfully  
          * APPROVE_ADD_PAA command completed successfully  
          * ADD_PAI command completed successfully  
          * revoke-child = True ⇒ all certificates signed by the revoked certificate will be revoked  
          * revoke-child = False ⇒ certificates in the chain will not be revoked  
          * PAI certificate with the provided subject and subject_key_id exists on the ledger  
     * REVOKE_PAI command failed ⇒ does not revoke the given PAI (intermediate certificate)  
          * PROPOSE_ADD_PAA command failed  
          * APPROVE_ADD_PAA command failed  
          * ADD_PAI command failed  
          * PAI certificate with the provided subject and subject_key_id does not exist  

* Role (Who can send)  
     * Positive:  
          * Vendor  
     * Negative:  
          * Trustee ⇒ error  
          * VendorAdmin ⇒ error  
          * CertificationCenter ⇒ error  
          * NodeAdmin ⇒ error

* Parameters:  
     * subject (Subject) - string  
          * Positive:  
               * string matches format  
               * valid text format  
               * MIN < length < MAX  
          * Negative:  
               * empty or nonexistent value  
               * length > MAX  
     * subject_key_id (Subject Key ID) - string  
          * Positive:  
               * string matches format  
               * valid text format, e.g., `5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB`  
               * MIN < length < MAX  
          * Negative:  
               * empty or nonexistent value  
               * length > MAX  
     * serial-number (Serial Number) - optional(string)  
          * Positive:  
               * empty or valid text format  
               * MIN < length < MAX  
          * Negative:  
               * length > MAX  
     * revoke-child (Revoke Child) - optional(bool)  
          * Positive:  
               * empty value or valid state  
               * TRUE (-1) ⇒ all child certs revoked  
               * FALSE (0) ⇒ default  
          * Negative:  
               * value is not a boolean  
     * info (Information/Notes) - optional(string)  
          * Positive:  
               * empty or valid text format  
               * MIN < length < MAX  
          * Negative:  
               * length > MAX (4096 characters)  
     * time (Proposal Time) - optional(int64)  
          * Positive:  
               * default value ⇒ current time  
               * empty value or valid integer format  
          * Negative:  
               * length > MAX (9,223,372,036,854,775,807)

## REMOVE_PAI	
### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
PAI	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
REMOVE_PAI command completed successfully	completely removes the given PAI (intermediate certificate) from both the approved and revoked certificates list
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_PAI command completed successfully	
a PAI Certificate with the provided subject and subject_key_id exist in the ledger	
REMOVE_PAI command failed	does not removes the given PAI (intermediate certificate) from both the approved and revoked certificates list
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_PAI command failed	
a PAI Certificate with the provided subject and subject_key_id not exist in the ledger	
Role	
 Who can send	
Trustee	error
Vendor 	
sender's VID match the VID of the removing certificate's owner	
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
serial-number (Serial Number)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
PAI	
Invalid command	
incorrect request	
server side error	
Сommand result	
REMOVE_PAI command completed successfully	completely removes the given PAI (intermediate certificate) from both the approved and revoked certificates list
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_PAI command completed successfully	
a PAI Certificate with the provided subject and subject_key_id exist in the ledger	
REMOVE_PAI command failed	does not removes the given PAI (intermediate certificate) from both the approved and revoked certificates list
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_PAI command failed	
a PAI Certificate with the provided subject and subject_key_id not exist in the ledger	
Role	
 Who can send	
Trustee	error
Vendor 	
sender's VID match the VID of the removing certificate's owner	
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
serial-number (Serial Number)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
## GET_DA_CERT	
### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
PAA	
PAI	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_DA_CERT command completed successfully	gets a DA certificate by the given subject and subject key ID attributes. 
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_PAI command completed successfully	
a PAI Certificate with the provided subject and subject_key_id exist in the ledger	
GET_DA_CERT command failed	does not gets a DA certificate by the given subject and subject key ID attributes. 
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_PAI command failed	
a PAI Certificate with the provided subject and subject_key_id not exist in the ledger	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
PAA	
PAI	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_DA_CERT command completed successfully	gets a DA certificate by the given subject and subject key ID attributes. 
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_PAI command completed successfully	
GET_DA_CERT command failed	does not gets a DA certificate by the given subject and subject key ID attributes. 
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_PAI command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
## GET_REVOKED_DA_CERT	
### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
PAA	
PAI	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_REVOKED_DA_CERT command completed successfully	gets a revoked DA certificate by the given subject and subject key ID attributes
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_PAI command completed successfully	
GET_REVOKED_DA_CERT command failed	does not gets a revoked DA certificate by the given subject and subject key ID attributes
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_PAI command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
PAA	
PAI	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_REVOKED_DA_CERT command completed successfully	gets a revoked DA certificate by the given subject and subject key ID attributes
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_PAI command completed successfully	
GET_REVOKED_DA_CERT command failed	does not gets a revoked DA certificate by the given subject and subject key ID attributes
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_PAI command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
## GET_DA_CERTS_BY_SKID	
### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
PAA	
PAI	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_DA_CERTS_BY_SKID command completed successfully	gets all DA certificates by the given subject key ID attribute
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_PAI command completed successfully	
GET_DA_CERTS_BY_SKID command failed	does not gets all DA certificates by the given subject key ID attribute
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_PAI command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
PAA	
PAI	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_DA_CERTS_BY_SKID command completed successfully	gets all DA certificates by the given subject key ID attribute
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_PAI command completed successfully	
GET_DA_CERTS_BY_SKID command failed	does not gets all DA certificates by the given subject key ID attribute
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_PAI command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
## GET_DA_CERTS_BY_SUBJECT	
### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
PAA	
PAI	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_DA_CERTS_BY_SUBJECT command completed successfully	gets all DA certificates associated with a subject
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_PAI command completed successfully	
GET_DA_CERTS_BY_SUBJECT command failed	does not gets all DA certificates associated with a subject
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_PAI command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
PAA	
PAI	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_DA_CERTS_BY_SUBJECT command completed successfully	gets all DA certificates associated with a subject
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_PAI command completed successfully	
GET_DA_CERTS_BY_SUBJECT command failed	does not gets all DA certificates associated with a subject
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_PAI command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
## GET_ALL_DA_CERTS	
### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
PAA	
PAI	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_DA_CERTS command completed successfully	gets all DA certificates
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_PAI command completed successfully	
GET_ALL_DA_CERTS command failed	does not gets all DA certificates
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_PAI command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
PAA	
PAI	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ALL_DA_CERTS command completed successfully	gets all DA certificates
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_PAI command completed successfully	
GET_ALL_DA_CERTS command failed	does not gets all DA certificates
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_PAI command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
### GET_ALL_REVOKED_DA_CERTS	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
PAA	
PAI	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_REVOKED_DA_CERTS command completed successfully	gets all revoked DA certificates
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_PAI command completed successfully	
GET_ALL_REVOKED_DA_CERTS command failed	does not gets all revoked DA certificates
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_PAI command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
PAA	
PAI	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ALL_REVOKED_DA_CERTS command completed successfully	gets all revoked DA certificates
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_PAI command completed successfully	
GET_ALL_REVOKED_DA_CERTS command failed	does not gets all revoked DA certificates
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_PAI command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
### GET_PKI_REVOCATION_DISTRIBUTION_POINT	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_PKI_REVOCATION_DISTRIBUTION_POINT command completed successfully	gets a revocation distribution point (such as RFC5280 Certificate Revocation List) identified by (VendorID, Label, IssuerSubjectKeyID) unique combination
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_REVOCATION_DISTRIBUTION_POINT command completed successfully	
GET_PKI_REVOCATION_DISTRIBUTION_POINT command failed	does not gets a revocation distribution point (such as RFC5280 Certificate Revocation List) identified by (VendorID, Label, IssuerSubjectKeyID) unique combination
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_REVOCATION_DISTRIBUTION_POINT command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
vid (Vendor ID)	uint16
     * Positive:	
value exists	
value > 0	
integer value format
     * Negative:	
empty value	
value =< 0
string value format	
length > MAX	MAX = 65535
nonexistent ID	
label (Label)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
empty value	
issuerSubjectKeyID (Issuer Subject Key ID)	string 
     * Positive:	
unique value for PAA/PAI	
text value format
MIN < length < MAX	
value consist of even number of uppercase hexadecimal characters ([0-9A-F])	for example, 5A880E6C3653D07FB08971A3F473790930E62BDB
     * Negative:	
contains whitespace	
contains non-hexadecimal characters	
length > MAX	
empty value	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_PKI_REVOCATION_DISTRIBUTION_POINT command completed successfully	gets a revocation distribution point (such as RFC5280 Certificate Revocation List) identified by (VendorID, Label, IssuerSubjectKeyID) unique combination
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_REVOCATION_DISTRIBUTION_POINT command completed successfully	
GET_PKI_REVOCATION_DISTRIBUTION_POINT command failed	does not gets a revocation distribution point (such as RFC5280 Certificate Revocation List) identified by (VendorID, Label, IssuerSubjectKeyID) unique combination
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_REVOCATION_DISTRIBUTION_POINT command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
vid (Vendor ID)	uint16
     * Positive:	
value exists	
value > 0	
integer value format
     * Negative:	
empty value	
value =< 0
string value format	
length > MAX	MAX = 65535
nonexistent ID	
label (Label)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
empty value	
issuerSubjectKeyID (Issuer Subject Key ID)	string 
     * Positive:	
unique value for PAA/PAI	
text value format
MIN < length < MAX	
value consist of even number of uppercase hexadecimal characters ([0-9A-F])	for example, 5A880E6C3653D07FB08971A3F473790930E62BDB
     * Negative:	
contains whitespace	
contains non-hexadecimal characters	
length > MAX	
empty value	
### GET_PKI_REVOCATION_DISTRIBUTION_POINTS_BY_SUBJECT_KEY_ID	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_PKI_REVOCATION_DISTRIBUTION_POINTS_BY_SUBJECT_KEY_ID command completed successfully	gets a list of revocation distribution point (such as RFC5280 Certificate Revocation List) identified by IssuerSubjectKeyID
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_REVOCATION_DISTRIBUTION_POINT command completed successfully	
GET_PKI_REVOCATION_DISTRIBUTION_POINTS_BY_SUBJECT_KEY_ID command failed	does not gets a a list of revocation distribution point (such as RFC5280 Certificate Revocation List) identified by IssuerSubjectKeyID
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_REVOCATION_DISTRIBUTION_POINT command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
issuerSubjectKeyID (Issuer Subject Key ID)	string 
     * Positive:	
unique value for PAA/PAI	
text value format
MIN < length < MAX	
value consist of even number of uppercase hexadecimal characters ([0-9A-F])	for example, 5A880E6C3653D07FB08971A3F473790930E62BDB
     * Negative:	
contains whitespace	
contains non-hexadecimal characters	
length > MAX	
empty value	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_PKI_REVOCATION_DISTRIBUTION_POINTS_BY_SUBJECT_KEY_ID command completed successfully	gets a list of revocation distribution point (such as RFC5280 Certificate Revocation List) identified by IssuerSubjectKeyID
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_REVOCATION_DISTRIBUTION_POINT command completed successfully	
GET_PKI_REVOCATION_DISTRIBUTION_POINTS_BY_SUBJECT_KEY_ID command failed	does not gets a a list of revocation distribution point (such as RFC5280 Certificate Revocation List) identified by IssuerSubjectKeyID
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_REVOCATION_DISTRIBUTION_POINT command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
issuerSubjectKeyID (Issuer Subject Key ID)	string 
     * Positive:	
unique value for PAA/PAI	
text value format
MIN < length < MAX	
value consist of even number of uppercase hexadecimal characters ([0-9A-F])	for example, 5A880E6C3653D07FB08971A3F473790930E62BDB
     * Negative:	
contains whitespace	
contains non-hexadecimal characters	
length > MAX	
empty value	
### GET_ALL_PKI_REVOCATION_DISTRIBUTION_POINT	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_PKI_REVOCATION_DISTRIBUTION_POINT command completed successfully	gets a list of all revocation distribution points (such as RFC5280 Certificate Revocation List)
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_REVOCATION_DISTRIBUTION_POINT command completed successfully	
GET_ALL_PKI_REVOCATION_DISTRIBUTION_POINT command failed	does not gets a list of all revocation distribution points (such as RFC5280 Certificate Revocation List)
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_REVOCATION_DISTRIBUTION_POINT command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ALL_PKI_REVOCATION_DISTRIBUTION_POINT command completed successfully	gets a list of all revocation distribution points (such as RFC5280 Certificate Revocation List)
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
ADD_REVOCATION_DISTRIBUTION_POINT command completed successfully	
GET_ALL_PKI_REVOCATION_DISTRIBUTION_POINT command failed	does not gets a list of all revocation distribution points (such as RFC5280 Certificate Revocation List)
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
ADD_REVOCATION_DISTRIBUTION_POINT command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
### GET_PROPOSED_PAA	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for all types of certificates	
PAA	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_PROPOSED_PAA command completed successfully	gets a proposed but not approved PAA certificate with the given subject and subject key ID attributes
PROPOSE_ADD_PAA command completed successfully	
GET_PROPOSED_PAA command failed	does not gets a proposed but not approved PAA certificate with the given subject and subject key ID attributes
PROPOSE_ADD_PAA command failed	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for all types of certificates	
PAA	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_PROPOSED_PAA command completed successfully	gets a proposed but not approved PAA certificate with the given subject and subject key ID attributes
PROPOSE_ADD_PAA command completed successfully	
GET_PROPOSED_PAA command failed	does not gets a proposed but not approved PAA certificate with the given subject and subject key ID attributes
PROPOSE_ADD_PAA command failed	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
### GET_REJECTED_PAA	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for all types of certificates	
PAA	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_REJECTED_PAA command completed successfully	gets a rejected PAA certificate with the given subject and subject key ID attributes
PROPOSE_ADD_PAA command completed successfully, but PAA certificate rejected	
there is at least one rejected PAA certificate	
GET_REJECTED_PAA command failed	does not gets a rejected PAA certificate with the given subject and subject key ID attributes
PROPOSE_ADD_PAA command failed	
there are not one rejected PAA certificates	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for all types of certificates	
PAA	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_REJECTED_PAA command completed successfully	gets a rejected PAA certificate with the given subject and subject key ID attributes
PROPOSE_ADD_PAA command completed successfully, but PAA certificate rejected	
there is at least one rejected PAA certificate	
GET_REJECTED_PAA command failed	does not gets a rejected PAA certificate with the given subject and subject key ID attributes
PROPOSE_ADD_PAA command failed	
there are not one rejected PAA certificates	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
### GET_PROPOSED_PAA_TO_REVOKE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
PAA 	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_PROPOSED_PAA_TO_REVOKE command completed successfully	gets a proposed but not approved PAA certificate to be revoked
PROPOSE_ADD_PAA command completed successfully, but PAA certificate not approved	
there is at least one not approved  PAA certificate	
GET_PROPOSED_PAA_TO_REVOKE command failed	does not gets a proposed but not approved PAA certificate to be revoked
PROPOSE_ADD_PAA command failed	
there are not one not approved PAA certificates	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
serial-number (Serial Number)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
PAA 	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_PROPOSED_PAA_TO_REVOKE command completed successfully	gets a proposed but not approved PAA certificate to be revoked
PROPOSE_ADD_PAA command completed successfully, but PAA certificate not approved	
there is at least one not approved  PAA certificate	
GET_PROPOSED_PAA_TO_REVOKE command failed	does not gets a proposed but not approved PAA certificate to be revoked
PROPOSE_ADD_PAA command failed	
there are not one not approved PAA certificates	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
serial-number (Serial Number)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
### GET_ALL_PAA	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
PAA 	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_PAA command completed successfully	gets all approved PAA certificates
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
there is at least one approved  PAA certificate	
GET_ALL_PAA command failed	does not gets all approved PAA certificates
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
there are not one approved PAA certificates	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
PAA 	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ALL_PAA command completed successfully	gets all approved PAA certificates
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
there is at least one approved  PAA certificate	
GET_ALL_PAA command failed	does not gets all approved PAA certificates
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
there are not one approved PAA certificates	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
### GET_ALL_REVOKED_PAA	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
PAA 	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_REVOKED_PAA command completed successfully	gets all revoked PAA certificates
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
PROPOSE_REVOKE_PAA command completed successfully	
APPROVE_REVOKE_PAA command completed successfully	
there is at least one revoked PAA certificates	
GET_ALL_REVOKED_PAA command failed	does not gets all revoked PAA certificates
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
PROPOSE_REVOKE_PAA command failed	
APPROVE_REVOKE_PAA command failed	
there are not one revoked PAA certificates	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
PAA 	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ALL_REVOKED_PAA command completed successfully	gets all revoked PAA certificates
PROPOSE_ADD_PAA command completed successfully	
APPROVE_ADD_PAA command completed successfully	
PROPOSE_REVOKE_PAA command completed successfully	
APPROVE_REVOKE_PAA command completed successfully	
there is at least one revoked PAA certificates	
GET_ALL_REVOKED_PAA command failed	does not gets all revoked PAA certificates
PROPOSE_ADD_PAA command failed	
APPROVE_ADD_PAA command failed	
PROPOSE_REVOKE_PAA command failed	
APPROVE_REVOKE_PAA command failed	
there are not one revoked PAA certificates	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
### GET_ALL_PROPOSED_PAA	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
PAA 	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_PROPOSED_PAA command completed successfully	gets all proposed but not approved root certificates
PROPOSE_ADD_PAA command completed successfully, but not approved root certificates	
there is at least one proposed but not approved root certificates	
GET_ALL_PROPOSED_PAA command failed	does not gets all proposed but not approved root certificates
PROPOSE_ADD_PAA command failed	
there are not one proposed but not approved root certificates	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
PAA 	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ALL_PROPOSED_PAA command completed successfully	gets all proposed but not approved root certificates
PROPOSE_ADD_PAA command completed successfully, but not approved root certificates	
there is at least one proposed but not approved root certificates	
GET_ALL_PROPOSED_PAA command failed	does not gets all proposed but not approved root certificates
PROPOSE_ADD_PAA command failed	
there are not one proposed but not approved root certificates	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
### GET_ALL_REJECTED_PAA	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
PAA 	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_REJECTED_PAA command completed successfully	gets all rejected root certificates
PROPOSE_ADD_PAA command completed successfully, but root certificates rejected 	
there is at least one rejected root certificates	
GET_ALL_REJECTED_PAA command failed	does not gets all rejected root certificates
PROPOSE_ADD_PAA command failed	
there are not one rejected root certificates	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
PAA 	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ALL_REJECTED_PAA command completed successfully	gets all rejected root certificates
PROPOSE_ADD_PAA command completed successfully, but root certificates rejected 	
there is at least one rejected root certificates	
GET_ALL_REJECTED_PAA command failed	does not gets all rejected root certificates
PROPOSE_ADD_PAA command failed	
there are not one rejected root certificates	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
### GET_ALL_PROPOSED_PAA_TO_REVOKE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
PAA 	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_PROPOSED_PAA_TO_REVOKE command completed successfully	gets all proposed but not approved root certificates to be revoked
PROPOSE_ADD_PAA command completed successfully, but root certificates not approved 	
there is at least one proposed but not approved root certificates	
GET_ALL_PROPOSED_PAA_TO_REVOKE command failed	does not gets all proposed but not approved root certificates to be revoked
PROPOSE_ADD_PAA command failed	
there are not one proposed but not approved root certificates	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
PAA 	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ALL_PROPOSED_PAA_TO_REVOKE command completed successfully	gets all proposed but not approved root certificates to be revoked
PROPOSE_ADD_PAA command completed successfully, but root certificates not approved 	
there is at least one proposed but not approved root certificates	
GET_ALL_PROPOSED_PAA_TO_REVOKE command failed	does not gets all proposed but not approved root certificates to be revoked
PROPOSE_ADD_PAA command failed	
there are not one proposed but not approved root certificates	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
# E2E (NOC): RCAC, ICAC	
### ADD_NOC_ROOT (RCAC)	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
RCAC	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
ADD_NOC_ROOT (RCAC) command completed successfully	adds a NOC root certificate (RCAC) owned by the Vendor
the provided certificate be a root certificate (RCAC)	
Issuer == Subject	
"Authority Key Identifier == Subject Key Identifier
"	
no existing certificate with the same <Certificate's Issuer>:<Certificate's Serial Number> combination	
certificates with the same <Certificate's Subject>:<Certificate's Subject Key ID> combination already exist	
the existing certificate NOC root certificate (RCAC)	
"the sender's VID match the vid field of the existing certificates
"	
the signature (self-signature) and expiration date must be valid	
ADD_NOC_ROOT (RCAC) command failed	does not adds a NOC root certificate (RCAC) owned by the Vendor
the provided certificate not a root certificate (RCAC)	
Issuer!= Subject and Authority Key Identifier == Subject Key Identifier	
"Issuer!= Subject  and Authority Key Identifier != Subject Key Identifier
"	
"Issuer == Subject  and Authority Key Identifier == Subject Key Identifier
"	
existing certificate with the same <Certificate's Issuer>:<Certificate's Serial Number> combination	
certificates with the same <Certificate's Subject>:<Certificate's Subject Key ID> combination already exist	
the existing certificate is not NOC root certificate (RCAC)	
"the sender's VID is not match the vid field of the existing certificates
"	
the signature (self-signature) and expiration date not be valid	
the signature (self-signature) is valid and expiration date not be valid	
the signature (self-signature) not be valid and expiration date valid	
Role	
 Who can send	
Trustee	error
Vendor 	
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
cert (Certificate)	string 
     * Positive:	
value exists	
contain a PEM string	
contain path to a file containing the data	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
schemaVersion (Schema Version)	optional(uint16)
     * Positive:	
value = 0	
integer value format
empty value	
     * Negative:	
length > MAX	MAX = 65535
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
RCAC	
Invalid command	
incorrect request	
server side error	
Сommand result	
ADD_NOC_ROOT (RCAC) command completed successfully	adds a NOC root certificate (RCAC) owned by the Vendor
the provided certificate be a root certificate (RCAC)	
Issuer == Subject	
"Authority Key Identifier == Subject Key Identifier
"	
no existing certificate with the same <Certificate's Issuer>:<Certificate's Serial Number> combination	
certificates with the same <Certificate's Subject>:<Certificate's Subject Key ID> combination already exist	
the existing certificate NOC root certificate (RCAC)	
"the sender's VID match the vid field of the existing certificates
"	
the signature (self-signature) and expiration date must be valid	
ADD_NOC_ROOT (RCAC) command failed	does not adds a NOC root certificate (RCAC) owned by the Vendor
the provided certificate not a root certificate (RCAC)	
Issuer!= Subject and Authority Key Identifier == Subject Key Identifier	
"Issuer!= Subject  and Authority Key Identifier != Subject Key Identifier
"	
"Issuer == Subject  and Authority Key Identifier == Subject Key Identifier
"	
existing certificate with the same <Certificate's Issuer>:<Certificate's Serial Number> combination	
certificates with the same <Certificate's Subject>:<Certificate's Subject Key ID> combination already exist	
the existing certificate is not NOC root certificate (RCAC)	
"the sender's VID is not match the vid field of the existing certificates
"	
the signature (self-signature) and expiration date not be valid	
the signature (self-signature) is valid and expiration date not be valid	
the signature (self-signature) not be valid and expiration date valid	
Role	
 Who can send	
Trustee	error
Vendor 	
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
cert (Certificate)	string 
     * Positive:	
value exists	
contain a PEM string	
contain path to a file containing the data	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
schemaVersion (Schema Version)	optional(uint16)
     * Positive:	
value = 0	
integer value format
empty value	
     * Negative:	
length > MAX	MAX = 65535
### REVOKE_NOC_ROOT (RCAC)	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
RCAC 	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
REVOKE_NOC_ROOT (RCAC) command completed successfully	revokes a NOC root certificate (RCAC) owned by the Vendor
ADD_NOC_ROOT (RCAC) command completed successfully	
a NOC Root Certificate (RCAC) with the provided subject and subject_key_id exist in the ledger	
REVOKE_NOC_ROOT (RCAC) command failed	does not revokes a NOC root certificate (RCAC) owned by the Vendor
ADD_NOC_ROOT (RCAC) command failed	
a NOC Root Certificate (RCAC) with the provided subject and subject_key_id is not exist in the ledger	
Role	
 Who can send	
Trustee	error
Vendor 	
Vid field associated with the corresponding NOC root certificate (RCAC) on the ledger equal to the Vendor account's VID	
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
serial-number (Serial Number)	optional(string)
     * Positive:	
empty value	
transaction will revoke all certificates that match the given subject and subject_key_id combination	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
revoke-child (Revoke Child)	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	then all certificates in the chain signed by the revoked certificate (intermediate, leaf) are revoked as well
FALSE (0)	only the current root cert is revoked (default value)
     * Negative:	
value is not bool	
info (Information/Notes)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	MAX=4096 characters
time (Proposal Time)	optional(int64)
     * Positive:	
default value	current time by default
empty value	
integer value format
     * Negative:	
length > MAX	MAX = 9 223 372 036 854 775 807
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
RCAC 	
Invalid command	
incorrect request	
server side error	
Сommand result	
REVOKE_NOC_ROOT (RCAC) command completed successfully	revokes a NOC root certificate (RCAC) owned by the Vendor
ADD_NOC_ROOT (RCAC) command completed successfully	
a NOC Root Certificate (RCAC) with the provided subject and subject_key_id exist in the ledger	
REVOKE_NOC_ROOT (RCAC) command failed	does not revokes a NOC root certificate (RCAC) owned by the Vendor
ADD_NOC_ROOT (RCAC) command failed	
a NOC Root Certificate (RCAC) with the provided subject and subject_key_id is not exist in the ledger	
Role	
 Who can send	
Trustee	error
Vendor 	
Vid field associated with the corresponding NOC root certificate (RCAC) on the ledger equal to the Vendor account's VID	
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
serial-number (Serial Number)	optional(string)
     * Positive:	
empty value	
transaction will revoke all certificates that match the given subject and subject_key_id combination	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
revoke-child (Revoke Child)	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	then all certificates in the chain signed by the revoked certificate (intermediate, leaf) are revoked as well
FALSE (0)	only the current root cert is revoked (default value)
     * Negative:	
value is not bool	
info (Information/Notes)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	MAX=4096 characters
time (Proposal Time)	optional(int64)
     * Positive:	
default value	current time by default
empty value	
integer value format
     * Negative:	
length > MAX	MAX = 9 223 372 036 854 775 807
### REMOVE_NOC_ROOT (RCAC)	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
RCAC 	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
REMOVE_NOC_ROOT (RCAC) command completed successfully	completely removes the given NOC root certificate (RCAC) owned by the Vendor from the ledger
ADD_NOC_ROOT (RCAC) command completed successfully	
a NOC Root Certificate (RCAC) with the provided subject and subject_key_id exist in the ledger	
REMOVE_NOC_ROOT (RCAC) command failed	does not completely removes the given NOC root certificate (RCAC) owned by the Vendor from the ledger
ADD_NOC_ROOT (RCAC) command failed	
a NOC Root Certificate (RCAC) with the provided subject and subject_key_id not exist in the ledger	
Role	
 Who can send	
Trustee	error
Vendor 	
Vid field associated with the corresponding NOC certificate on the ledger equal to the Vendor account's VID	
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
serial-number (Serial Number)	optional(string)
     * Positive:	
empty value	
transaction will revoke all certificates that match the given subject and subject_key_id combination	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
RCAC 	
Invalid command	
incorrect request	
server side error	
Сommand result	
REMOVE_NOC_ROOT (RCAC) command completed successfully	completely removes the given NOC root certificate (RCAC) owned by the Vendor from the ledger
ADD_NOC_ROOT (RCAC) command completed successfully	
a NOC Root Certificate (RCAC) with the provided subject and subject_key_id exist in the ledger	
REMOVE_NOC_ROOT (RCAC) command failed	does not completely removes the given NOC root certificate (RCAC) owned by the Vendor from the ledger
ADD_NOC_ROOT (RCAC) command failed	
a NOC Root Certificate (RCAC) with the provided subject and subject_key_id not exist in the ledger	
Role	
 Who can send	
Trustee	error
Vendor 	
Vid field associated with the corresponding NOC certificate on the ledger equal to the Vendor account's VID	
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
serial-number (Serial Number)	optional(string)
     * Positive:	
empty value	
transaction will revoke all certificates that match the given subject and subject_key_id combination	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
### ADD_NOC_ICA (ICAC)	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
ICAC	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
ADD_NOC_ICA (ICAC) command completed successfully	adds a NOC ICA certificate (ICAC) owned by the Vendor signed by a chain of certificates which must be already present on the ledger
ADD_NOC_ROOT (RCAC) command completed successfully	
the certificate chain is already present in the ledger	
"the provided certificate a non-root certificate
"	
Issuer != Subject	
Authority Key Identifier != Subject Key Identifier	
the root certificate e a NOC certificate and added by the same vendor	
isNoc field of the root certificate set to true	
VID of root certificate == VID of account	
no existing certificate with the same <Certificate's Issuer>:<Certificate's Serial Number> combination	
certificates with the same <Certificate's Subject>:<Certificate's Subject Key ID> combination already exist	
the existing certificate NOC non-root certificate	
the sender's VID match the vid field of the existing certificates	
the signature and expiration date valid	
ADD_NOC_ICA (ICAC) command failed	does not adds a NOC ICA certificate (ICAC) owned by the Vendor signed by a chain of certificates which must be already present on the ledger
ADD_NOC_ROOT (RCAC) command failed	
the certificate chain is not present in the ledger	
"the provided certificate a root certificate
"	
Issuer != Subject and Authority Key Identifier == Subject Key Identifier	
Issuer == Subject and Authority Key Identifier != Subject Key Identifier	
Issuer == Subject and Authority Key Identifier == Subject Key Identifier	
the root certificate is not NOC certificate and added by the same vendor	
isNoc field of the root certificate is not set to true	
VID of root certificate != VID of account	
existing certificate with the same <Certificate's Issuer>:<Certificate's Serial Number> combination	
certificates with the same <Certificate's Subject>:<Certificate's Subject Key ID> combination already exist	
the existing certificate NOC root certificate	
the sender's VID does not match the vid field of the existing certificates	
the signature and expiration date invalid	
the signature invalid and expiration date valid	
the signature valid and expiration date invalid	
Role	
 Who can send	
Trustee	error
Vendor 	
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
cert (Certificate)	string 
     * Positive:	
value exists	
contain a PEM string	
contain path to a file containing the data	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
certificate-schema-version	optional(uint16)
     * Positive:	
value = 0	
integer value format
empty value	
     * Negative:	
length > MAX	MAX = 65535
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
ICAC	
Invalid command	
incorrect request	
server side error	
Сommand result	
ADD_NOC_ICA (ICAC) command completed successfully	adds a NOC ICA certificate (ICAC) owned by the Vendor signed by a chain of certificates which must be already present on the ledger
ADD_NOC_ROOT (RCAC) command completed successfully	
the certificate chain is already present in the ledger	
"the provided certificate a non-root certificate
"	
Issuer != Subject	
Authority Key Identifier != Subject Key Identifier	
the root certificate e a NOC certificate and added by the same vendor	
isNoc field of the root certificate set to true	
VID of root certificate == VID of account	
no existing certificate with the same <Certificate's Issuer>:<Certificate's Serial Number> combination	
certificates with the same <Certificate's Subject>:<Certificate's Subject Key ID> combination already exist	
the existing certificate NOC non-root certificate	
the sender's VID match the vid field of the existing certificates	
the signature and expiration date valid	
ADD_NOC_ICA (ICAC) command failed	does not adds a NOC ICA certificate (ICAC) owned by the Vendor signed by a chain of certificates which must be already present on the ledger
ADD_NOC_ROOT (RCAC) command failed	
the certificate chain is not present in the ledger	
"the provided certificate a root certificate
"	
Issuer != Subject and Authority Key Identifier == Subject Key Identifier	
Issuer == Subject and Authority Key Identifier != Subject Key Identifier	
Issuer == Subject and Authority Key Identifier == Subject Key Identifier	
the root certificate is not NOC certificate and added by the same vendor	
isNoc field of the root certificate is not set to true	
VID of root certificate != VID of account	
existing certificate with the same <Certificate's Issuer>:<Certificate's Serial Number> combination	
certificates with the same <Certificate's Subject>:<Certificate's Subject Key ID> combination already exist	
the existing certificate NOC root certificate	
the sender's VID does not match the vid field of the existing certificates	
the signature and expiration date invalid	
the signature invalid and expiration date valid	
the signature valid and expiration date invalid	
Role	
 Who can send	
Trustee	error
Vendor 	
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
cert (Certificate)	string 
     * Positive:	
value exists	
contain a PEM string	
contain path to a file containing the data	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
certificate-schema-version	optional(uint16)
     * Positive:	
value = 0	
integer value format
empty value	
     * Negative:	
length > MAX	MAX = 65535
### REVOKE_NOC_ICA (ICAC)	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
ICAC	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
REVOKE_NOC_ICA (ICAC) command completed successfully	revokes a NOC ICA certificate (ICAC) owned by the Vendor
ADD_NOC_ROOT (RCAC) command completed successfully	
a NOC Certificate with the provided subject and subject_key_id exist in the ledger	
REVOKE_NOC_ICA (ICAC) command failed	does not revokes a NOC ICA certificate (ICAC) owned by the Vendor
ADD_NOC_ROOT (RCAC) command failed	
a NOC Certificate with the provided subject and subject_key_id is not exist in the ledger	
Role	
 Who can send	
Trustee	error
Vendor 	
Vid field associated with the corresponding NOC certificate on the ledger must be equal to the Vendor account's VID	
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
serial-number (Serial Number)	optional(string)
     * Positive:	
empty value	
transaction will revoke all certificates that match the given subject and subject_key_id combination	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
revoke-child (Revoke Child)	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	all certificates in the chain signed by the revoked certificate (leaf) are revoked as well
FALSE (0)	only the current cert is revoked (default value)
     * Negative:	
value is not bool	
info (Information/Notes)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	MAX=4096 characters
time (Proposal Time)	optional(int64)
     * Positive:	
default value	current time by default
empty value	
integer value format
     * Negative:	
length > MAX	MAX = 9 223 372 036 854 775 807
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
ICAC	
Invalid command	
incorrect request	
server side error	
Сommand result	
REVOKE_NOC_ICA (ICAC) command completed successfully	revokes a NOC ICA certificate (ICAC) owned by the Vendor
ADD_NOC_ROOT (RCAC) command completed successfully	
a NOC Certificate with the provided subject and subject_key_id exist in the ledger	
REVOKE_NOC_ICA (ICAC) command failed	does not revokes a NOC ICA certificate (ICAC) owned by the Vendor
ADD_NOC_ROOT (RCAC) command failed	
a NOC Certificate with the provided subject and subject_key_id is not exist in the ledger	
Role	
 Who can send	
Trustee	error
Vendor 	
Vid field associated with the corresponding NOC certificate on the ledger must be equal to the Vendor account's VID	
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
serial-number (Serial Number)	optional(string)
     * Positive:	
empty value	
transaction will revoke all certificates that match the given subject and subject_key_id combination	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
revoke-child (Revoke Child)	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	all certificates in the chain signed by the revoked certificate (leaf) are revoked as well
FALSE (0)	only the current cert is revoked (default value)
     * Negative:	
value is not bool	
info (Information/Notes)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	MAX=4096 characters
time (Proposal Time)	optional(int64)
     * Positive:	
default value	current time by default
empty value	
integer value format
     * Negative:	
length > MAX	MAX = 9 223 372 036 854 775 807
### REMOVE_NOC_ICA (ICAC)	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for  the following certificates	
ICAC	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
REMOVE_NOC_ICA (ICAC) command completed successfully	completely removes the given NOC ICA (ICAC) owned by the Vendor from the ledger
ADD_NOC_ROOT (RCAC) command completed successfully	
a NOC ICA Certificate (ICAC) with the provided subject and subject_key_id exist in the ledger	
REMOVE_NOC_ICA (ICAC) command failed	does not removes the given NOC ICA (ICAC) owned by the Vendor from the ledger
ADD_NOC_ROOT (RCAC) command failed	
a NOC ICA Certificate (ICAC) with the provided subject and subject_key_id not exist in the ledger	
Role	
 Who can send	
Trustee	error
Vendor 	
Vid field associated with the corresponding NOC certificate on the ledger equal to the Vendor account's VID	
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
serial-number (Serial Number)	optional(string)
     * Positive:	
empty value	
transaction will revoke all certificates that match the given subject and subject_key_id combination	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for  the following certificates	
ICAC	
Invalid command	
incorrect request	
server side error	
Сommand result	
REMOVE_NOC_ICA (ICAC) command completed successfully	completely removes the given NOC ICA (ICAC) owned by the Vendor from the ledger
ADD_NOC_ROOT (RCAC) command completed successfully	
a NOC ICA Certificate (ICAC) with the provided subject and subject_key_id exist in the ledger	
REMOVE_NOC_ICA (ICAC) command failed	does not removes the given NOC ICA (ICAC) owned by the Vendor from the ledger
ADD_NOC_ROOT (RCAC) command failed	
a NOC ICA Certificate (ICAC) with the provided subject and subject_key_id not exist in the ledger	
Role	
 Who can send	
Trustee	error
Vendor 	
Vid field associated with the corresponding NOC certificate on the ledger equal to the Vendor account's VID	
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
serial-number (Serial Number)	optional(string)
     * Positive:	
empty value	
transaction will revoke all certificates that match the given subject and subject_key_id combination	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
### GET_NOC_CERT	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for all types of Noc certificates	
NOC_ROOT	
NOC_ICA	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_NOC_CERT command completed successfully	gets a NOC certificate by the given subject and subject key ID attributes
ADD_NOC_ROOT (RCAC) command completed successfully	
GET_NOC_CERT command failed	does not gets a NOC certificate by the given subject and subject key ID attributes
ADD_NOC_ROOT (RCAC) command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for all types of Noc certificates	
NOC_ROOT	
NOC_ICA	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_NOC_CERT command completed successfully	gets a NOC certificate by the given subject and subject key ID attributes
ADD_NOC_ROOT (RCAC) command completed successfully	
GET_NOC_CERT command failed	does not gets a NOC certificate by the given subject and subject key ID attributes
ADD_NOC_ROOT (RCAC) command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
### GET_NOC_ROOT_BY_VID (RCACs)	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for all types of Noc certificates	
RCACs	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_NOC_ROOT_BY_VID (RCACs) command completed successfully	retrieve NOC root certificates (RCACs) associated with a specific VID
ADD_NOC_ROOT (RCAC) command completed successfully	
GET_NOC_ROOT_BY_VID (RCACs) command failed	does not retrieve NOC root certificates (RCACs) associated with a specific VID
ADD_NOC_ROOT (RCAC) command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
vid (Vendor ID)	uint16
     * Positive:	
unique value	
value > 0	
integer value format
     * Negative:	
empty value
value =< 0
string value format		
nonexistent value	
length > MAX	MAX = 65535
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for all types of Noc certificates	
RCACs	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_NOC_ROOT_BY_VID (RCACs) command completed successfully	retrieve NOC root certificates (RCACs) associated with a specific VID
ADD_NOC_ROOT (RCAC) command completed successfully	
GET_NOC_ROOT_BY_VID (RCACs) command failed	does not retrieve NOC root certificates (RCACs) associated with a specific VID
ADD_NOC_ROOT (RCAC) command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
vid (Vendor ID)	uint16
     * Positive:	
unique value	
value > 0	
integer value format
     * Negative:	
empty value
value =< 0
string value format		
nonexistent value	
length > MAX	MAX = 65535
### GET_NOC_BY_VID_AND_SKID (RCACs/ICACs)	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for all types of Noc certificates	
RCACs	
ICACs	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_NOC_BY_VID_AND_SKID (RCACs/ICACs) command completed successfully	retrieve NOC (Root/ICA) certificates (RCACs/ICACs) associated with a specific VID and subject key ID
ADD_NOC_ROOT (RCAC) command completed successfully	
GET_NOC_BY_VID_AND_SKID (RCACs/ICACs) command failed	does not retrieve NOC (Root/ICA) certificates (RCACs/ICACs) associated with a specific VID and subject key ID
ADD_NOC_ROOT (RCAC) command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
vid (Vendor ID)	uint16
     * Positive:	
unique value	
value > 0	
integer value format
     * Negative:	
empty value
value =< 0
string value format		
nonexistent value	
length > MAX	MAX = 65535
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for all types of Noc certificates	
RCACs	
ICACs	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_NOC_BY_VID_AND_SKID (RCACs/ICACs) command completed successfully	retrieve NOC (Root/ICA) certificates (RCACs/ICACs) associated with a specific VID and subject key ID
ADD_NOC_ROOT (RCAC) command completed successfully	
GET_NOC_BY_VID_AND_SKID (RCACs/ICACs) command failed	does not retrieve NOC (Root/ICA) certificates (RCACs/ICACs) associated with a specific VID and subject key ID
ADD_NOC_ROOT (RCAC) command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
vid (Vendor ID)	uint16
     * Positive:	
unique value	
value > 0	
integer value format
     * Negative:	
empty value
value =< 0
string value format		
nonexistent value	
length > MAX	MAX = 65535
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
### GET_NOC_ICA_BY_VID (ICACs)	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for all types of Noc certificates	
ICACs	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_NOC_ICA_BY_VID (ICACs) command completed successfully	retrieve NOC ICA certificates (ICACs) associated with a specific VID
ADD_NOC_ROOT (RCAC) command completed successfully	
GET_NOC_ICA_BY_VID (ICACs) command failed	does not retrieve NOC ICA certificates (ICACs) associated with a specific VID
ADD_NOC_ROOT (RCAC) command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
vid (Vendor ID)	uint16
     * Positive:	
unique value	
value > 0	
integer value format
     * Negative:	
empty value
value =< 0
string value format		
nonexistent value	
length > MAX	MAX = 65535
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for all types of Noc certificates	
ICACs	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_NOC_ICA_BY_VID (ICACs) command completed successfully	retrieve NOC ICA certificates (ICACs) associated with a specific VID
ADD_NOC_ROOT (RCAC) command completed successfully	
GET_NOC_ICA_BY_VID (ICACs) command failed	does not retrieve NOC ICA certificates (ICACs) associated with a specific VID
ADD_NOC_ROOT (RCAC) command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
vid (Vendor ID)	uint16
     * Positive:	
unique value	
value > 0	
integer value format
     * Negative:	
empty value
value =< 0
string value format		
nonexistent value	
length > MAX	MAX = 65535
### GET_NOC_CERTS_BY_SUBJECT	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for both types of certificates	
NOC_ROOT	
NOC_ICA	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_NOC_CERTS_BY_SUBJECT command completed successfully	gets all NOC certificates associated with a subject
ADD_NOC_ROOT (RCAC) command completed successfully	
GET_NOC_CERTS_BY_SUBJECT command failed	does not gets all NOC certificates associated with a subject
ADD_NOC_ROOT (RCAC) command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for both types of certificates	
NOC_ROOT	
NOC_ICA	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_NOC_CERTS_BY_SUBJECT command completed successfully	gets all NOC certificates associated with a subject
ADD_NOC_ROOT (RCAC) command completed successfully	
GET_NOC_CERTS_BY_SUBJECT command failed	does not gets all NOC certificates associated with a subject
ADD_NOC_ROOT (RCAC) command failed	
Role	
 Who can send	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
### GET_REVOKED_NOC_ROOT (RCAC)	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for following certificates	
RCAC	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_REVOKED_NOC_ROOT (RCAC) command completed successfully	gets a revoked NOC root certificate (RCAC) by the given subject and subject key ID attributes
ADD_NOC_ROOT (RCAC) command completed successfully	
GET_REVOKED_NOC_ROOT (RCAC) command failed	does not gets a revoked NOC root certificate (RCAC) by the given subject and subject key ID attributes
ADD_NOC_ROOT (RCAC) command failed	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for all types of certificates	
RCAC	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_REVOKED_NOC_ROOT (RCAC) command completed successfully	gets a revoked NOC root certificate (RCAC) by the given subject and subject key ID attributes
ADD_NOC_ROOT (RCAC) command completed successfully	
GET_REVOKED_NOC_ROOT (RCAC) command failed	does not gets a revoked NOC root certificate (RCAC) by the given subject and subject key ID attributes
ADD_NOC_ROOT (RCAC) command failed	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
### GET_REVOKED_NOC_ICA (ICAC)	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for following certificates	
ICAC	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_REVOKED_NOC_ICA (ICAC) command completed successfully	gets a revoked NOC ica certificate (ICAC) by the given subject and subject key ID attributes
ADD_NOC_ROOT (RCAC) command completed successfully	
GET_REVOKED_NOC_ICA (ICAC) command failed	does not a revoked NOC ica certificate (ICAC) by the given subject and subject key ID attributes
ADD_NOC_ROOT (RCAC) command failed	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for all types of certificates	
ICAC	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_REVOKED_NOC_ICA (ICAC) command completed successfully	gets a revoked NOC ica certificate (ICAC) by the given subject and subject key ID attributes
ADD_NOC_ROOT (RCAC) command completed successfully	
GET_REVOKED_NOC_ICA (ICAC) command failed	does not a revoked NOC ica certificate (ICAC) by the given subject and subject key ID attributes
ADD_NOC_ROOT (RCAC) command failed	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
subject (Subject)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
subject_key_id (Subject Key ID)	string 
     * Positive:	
string matches the format	for example: 5A:88:0E:6C:36:53:D0:7F:B0:89:71:A3:F4:73:79:09:30:E6:2B:DB
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
### GET_ALL_NOC (RCACs/ICACs)	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for following certificates	
RCACs 	
ICACs	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_NOC (RCACs/ICACs) command completed successfully	retrieve a list of all of NOC certificates (RCACs of ICACs)
ADD_NOC_ROOT (RCAC) command completed successfully	
there is at least one NOC certificates (RCACs of ICACs)	
GET_ALL_NOC (RCACs/ICACs) command failed	does not retrieve a list of all of NOC certificates (RCACs of ICACs)
ADD_NOC_ROOT (RCAC) command failed	
there are not one NOC certificates (RCACs of ICACs)	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for following certificates	
RCACs 	
ICACs	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ALL_NOC (RCACs/ICACs) command completed successfully	retrieve a list of all of NOC certificates (RCACs of ICACs)
ADD_NOC_ROOT (RCAC) command completed successfully	
there is at least one NOC certificates (RCACs of ICACs)	
GET_ALL_NOC (RCACs/ICACs) command failed	does not retrieve a list of all of NOC certificates (RCACs of ICACs)
ADD_NOC_ROOT (RCAC) command failed	
there are not one NOC certificates (RCACs of ICACs)	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
### GET_ALL_NOC_ROOT (RCACs)	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for following certificates	
RCACs 	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_NOC_ROOT (RCACs) command completed successfully	retrieve a list of all of NOC root certificates (RCACs)
ADD_NOC_ROOT (RCAC) command completed successfully	
there is at least one NOC root certificates (RCACs)	
GET_ALL_NOC_ROOT (RCACs) command failed	does not retrieve a list of all of NOC root certificates (RCACs)
ADD_NOC_ROOT (RCAC) command failed	
there are not one NOC root certificates (RCACs)	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for following certificates	
RCACs 	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ALL_NOC_ROOT (RCACs) command completed successfully	retrieve a list of all of NOC root certificates (RCACs)
ADD_NOC_ROOT (RCAC) command completed successfully	
there is at least one NOC root certificates (RCACs)	
GET_ALL_NOC_ROOT (RCACs) command failed	does not retrieve a list of all of NOC root certificates (RCACs)
ADD_NOC_ROOT (RCAC) command failed	
there are not one NOC root certificates (RCACs)	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
### GET_ALL_NOC_ICA (ICACs)	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for following certificates	
ICACs	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_NOC_ICA (ICACs) command completed successfully	retrieve a list of all of NOC ICA certificates (ICACs)
ADD_NOC_ROOT (RCAC) command completed successfully	
there is at least one NOC ICA certificates (ICACs)	
GET_ALL_NOC_ICA (ICACs) command failed	does not retrieve a list of all of NOC ICA certificates (ICACs)
ADD_NOC_ROOT (RCAC) command failed	
there are not one NOC ICA certificates (ICACs)	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for following certificates	
ICACs	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ALL_NOC_ICA (ICACs) command completed successfully	retrieve a list of all of NOC ICA certificates (ICACs)
ADD_NOC_ROOT (RCAC) command completed successfully	
there is at least one NOC ICA certificates (ICACs)	
GET_ALL_NOC_ICA (ICACs) command failed	does not retrieve a list of all of NOC ICA certificates (ICACs)
ADD_NOC_ROOT (RCAC) command failed	
there are not one NOC ICA certificates (ICACs)	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
### GET_ALL_REVOKED_NOC_ROOT (RCACs)	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for following certificates	
RCACs	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_REVOKED_NOC_ROOT (RCACs) command completed successfully	gets all revoked NOC root certificates (RCACs)
ADD_NOC_ROOT (RCAC) command completed successfully	
there is at least one revoked NOC root certificates (RCACs)	
GET_ALL_REVOKED_NOC_ROOT (RCACs) command failed	does not gets all revoked NOC root certificates (RCACs)
ADD_NOC_ROOT (RCAC) command failed	
there are not one revoked NOC root certificates (RCACs)	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for following certificates	
RCACs	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ALL_REVOKED_NOC_ROOT (RCACs) command completed successfully	gets all revoked NOC root certificates (RCACs)
ADD_NOC_ROOT (RCAC) command completed successfully	
there is at least one revoked NOC root certificates (RCACs)	
GET_ALL_REVOKED_NOC_ROOT (RCACs) command failed	does not gets all revoked NOC root certificates (RCACs)
ADD_NOC_ROOT (RCAC) command failed	
there are not one revoked NOC root certificates (RCACs)	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
### GET_ALL_REVOKED_NOC_ICA (ICACs)	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
query works for following certificates	
ICACs	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_REVOKED_NOC_ICA (ICACs) command completed successfully	gets all revoked NOC ica certificates (ICACs)
ADD_NOC_ROOT (RCAC) command completed successfully	
there is at least one revoked NOC ica certificates (ICACs)	
GET_ALL_REVOKED_NOC_ICA (ICACs) command failed	does not gets all revoked NOC ica certificates (ICACs)
ADD_NOC_ROOT (RCAC) command failed	
there are not one revoked NOC ica certificates (ICACs)	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
#### REST API 	
REST API command send	
Valid command	
correct HTTP method	
request is authorized	
uses valid credentials/role	
query works for following certificates	
ICACs	
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ALL_REVOKED_NOC_ICA (ICACs) command completed successfully	gets all revoked NOC ica certificates (ICACs)
ADD_NOC_ROOT (RCAC) command completed successfully	
there is at least one revoked NOC ica certificates (ICACs)	
GET_ALL_REVOKED_NOC_ICA (ICACs) command failed	does not gets all revoked NOC ica certificates (ICACs)
ADD_NOC_ROOT (RCAC) command failed	
there are not one revoked NOC ica certificates (ICACs)	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
count-total	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	
limit 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value > 100	
offset 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page 	optional(uint)
     * Positive:	
value exists	
empty value	
     * Negative:	
value < 0	
page-key	optional(string)
     * Positive:	
empty value	
value exists	
     * Negative:	
length < MIN	
reverse	optional(bool)
     * Positive:	
empty value	
value state	
TRUE (-1)	
FALSE (0)	
     * Negative:	
value is not bool	      