1. [CERTIFY_MODEL](#certify_model)
2. [UPDATE_COMPLIANCE_INFO](#update_complince_info)
3. [DELETE_COMPLIANCE_INFO](#delete_complince_info)
4. [REVOKE_MODEL_CERTIFICATION](#revoke_model_certification)
5. [PROVISION_MODEL](#provision_model)
6. [GET_CERTIFIED_MODEL](#get_certified_model)
7. [GET_REVOKED_MODEL](#get_revoked_model)
8. [GET_PROVISIONAL_MODEL](#get_provisional_model)
9. [GET_COMPLIANCE_INFO](#get_compliance_info)
10. [GET_DEVICE_SOFTWARE_COMPLIANCE](#get_device_software_compliance)
11. [GET_ALL_CERTIFIED_MODELS](#get_all_certified_models)
12. [GET_ALL_REVOKED_MODELS](#get_all_revoked_models)
13. [GET_ALL_PROVISIONAL_MODELS](#get_all_provisional_models)
14. [GET_ALL_COMPLIANCE_INFO_RECORDS](#get_all_compliance_info_records)
15. [GET_ALL_DEVICE_SOFTWARE_COMPLIANCES](#get_all_device_software_compliances)

## [CERTIFY_MODEL](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/docs/transactions/compliance.md#certify_model)	

### CLI command	
CLI command: `https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/docs/transactions/compliance.md#certify_model`

* CLI command send	
     * Valid command	
          * command exists/relevant	
     * Invalid command	
          * access is denied to execute command	
          * incorrect command syntax	

* Сommand result	
     * CERTIFY_MODEL command completed successfully ⇒ attests compliance of the Model Version to the ZB or Matter standard
          * ADD_MODEL and ADD_MODEL_VERSION command completed successfully	a record of a Model and Model Version is present on the ledger corresponding Model and Model Version is present on the ledger	
          * command is called for compliant device for use cases where compliance is tracked on ledger	
          * revocation is  tracked on the ledger to remove a Model Version from the revocation list	
     * CERTIFY_MODEL command failed ⇒ does not attests compliance of the Model Version to the ZB or Matter standard
          * ADD_MODEL command was not executed	
          * record about a model has been removed	
          * ADD_MODEL command executed, but ADD_MODEL_VERSION command was not executed	
          * record about a model version has been removed	
          * corresponding Model Version not present on the ledger	
          * corresponding Model Version is  present on the ledger, but corresponding Model not present on the ledger	
          * command not called for compliant device for use cases where compliance is tracked on ledger	
          * revocation is not tracked on the ledger to remove a Model Version from the revocation list	

* Role (Who can send)
     * Positive:
          * CertificationCenter
     * Negative:	
          * Trustee	
          * Vendor 	
          * VendorAdmin 	 	
          * NodeAdmin 	

* Parameters:	
     * vid (Vendor ID) - uint16
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * pid (Product ID) - uint16 
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * softwareVersion (Software Version) - uint32 
          * Positive:	
               * value exists	
               * value >= 0	
               * integer value format
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	(MAX = 4 294 967 295)
     * softwareVersionSting (Software Version Sting) - string 
          * Positive:	
               * value exists	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * certificationDate (Certification Date) - string 
          * Positive:	
               * string matches the format	2019-10-12T07:20:50.52Z
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * certificationType (Certification Type) - string 
          * Positive:	
               * existing value	
               * valid type	
                    * zigbee	
                    * matter	
                    * access control	
                    * product security	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * cdCertificateId (CD Certificate ID) - string 
          * Positive:	
               * existing value	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * reason (Reason) - optional(string)
          * Positive:	
               * empty value	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * cDVersionNumber (CD Version Number) - optional(uint32)
          * Positive:	
               * value exists	
               * value > 0	
               * empty value	
               * format	
          * Negative:
               * value =< 0
               * string value format		
               * length > MAX	MAX = 4 294 967 295
     * familyId (Family ID) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * supportedClusters (Supported Clusters) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * compliantPlatformUsed (Compliant Platform Used) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * compliantPlatformVersion (Compliant Platform Version) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * OSVersion (OS Version) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * certificationRoute (Certification Route) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * programType (Program Type) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * programTypeVersion (Program Type Version) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * transport (Transport) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * parentChild (Parent Child) - optional(string)
          * Positive:	
               * supported types	
               * parent	
               * child	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * certificationIDOfSoftwareComponent (Certification ID Of Software Component) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * schemaVersion (Schema Version) - optional(uint16)
          * Positive:	
               * value = 0	
               * integer value format
               * empty value	
          * Negative:	
               * length > MAX	MAX = 65535

### REST API 	
REST API command: POST /cosmos/tx/v1beta1/txs: [MsgCertifyModel](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/proto/zigbeealliance/distributedcomplianceledger/compliance/tx.proto#L21)

* REST API command send	
     * Valid command	
          * correct HTTP method	
          * request is authorized	
          * uses valid credentials/role	
     * Invalid command	
          * incorrect request	
          * server side error	

* Сommand result	
     * CERTIFY_MODEL command completed successfully ⇒ attests compliance of the Model Version to the ZB or Matter standard
          * ADD_MODEL and ADD_MODEL_VERSION command completed successfully	a record of a Model and Model Version is present on the ledger corresponding Model and Model Version is present on the ledger	
          * command is called for compliant device for use cases where compliance is tracked on ledger	
          * revocation is  tracked on the ledger to remove a Model Version from the revocation list	
     * CERTIFY_MODEL command failed ⇒ does not attests compliance of the Model Version to the ZB or Matter standard
          * ADD_MODEL command was not executed	
          * record about a model has been removed	
          * ADD_MODEL command executed, but ADD_MODEL_VERSION command was not executed	
          * record about a model version has been removed	
          * corresponding Model Version not present on the ledger	
          * corresponding Model Version is  present on the ledger, but corresponding Model not present on the ledger	
          * command not called for compliant device for use cases where compliance is tracked on ledger	
          * revocation is not tracked on the ledger to remove a Model Version from the revocation list	

* Role (Who can send)
     * Positive:
          * CertificationCenter
     * Negative:	
          * Trustee	
          * Vendor 	
          * VendorAdmin 	 	
          * NodeAdmin 	

* Parameters:	
     * vid (Vendor ID) - uint16
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * pid (Product ID) - uint16 
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * softwareVersion (Software Version) - uint32 
          * Positive:	
               * value exists	
               * value >= 0	
               * integer value format
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	(MAX = 4 294 967 295)
     * softwareVersionSting (Software Version Sting) - string 
          * Positive:	
               * value exists	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * certificationDate (Certification Date) - string 
          * Positive:	
               * string matches the format	2019-10-12T07:20:50.52Z
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * certificationType (Certification Type) - string 
          * Positive:	
               * existing value	
               * valid type	
                    * zigbee	
                    * matter	
                    * access control	
                    * product security	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * cdCertificateId (CD Certificate ID) - string 
          * Positive:	
               * existing value	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * reason (Reason) - optional(string)
          * Positive:	
               * empty value	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * cDVersionNumber (CD Version Number) - optional(uint32)
          * Positive:	
               * value exists	
               * value > 0	
               * empty value	
               * format	
          * Negative:
               * value =< 0
               * string value format		
               * length > MAX	MAX = 4 294 967 295
     * familyId (Family ID) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * supportedClusters (Supported Clusters) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * compliantPlatformUsed (Compliant Platform Used) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * compliantPlatformVersion (Compliant Platform Version) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * OSVersion (OS Version) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * certificationRoute (Certification Route) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * programType (Program Type) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * programTypeVersion (Program Type Version) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * transport (Transport) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * parentChild (Parent Child) - optional(string)
          * Positive:	
               * supported types	
               * parent	
               * child	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * certificationIDOfSoftwareComponent (Certification ID Of Software Component) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * schemaVersion (Schema Version) - optional(uint16)
          * Positive:	
               * value = 0	
               * integer value format
               * empty value	
          * Negative:	
               * length > MAX	MAX = 65535

## [UPDATE_COMPLIANCE_INFO]( https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/docs/transactions/compliance.md#update_compliance_info)	

### CLI command
CLI command: `dcld tx compliance update-compliance-info`	

* CLI command send	
     * Valid command	
          * command exists/relevant	
     * Invalid command	
          * access is denied to execute command	
          * incorrect command syntax	

* Сommand result	
     * UPDATE_COMPLIANCE_INFO command completed successfully ⇒ updates a compliance info by VID, PID, Software Version and Certification Type
          * CERTIFY_MODEL command completed successfully for the selected compliance	
          * VID, PID, Software Version and Certification Type value are indicated correctly	
     * UPDATE_COMPLIANCE_INFO command failed	⇒ does not updates a compliance info by VID, PID, Software Version and Certification Type
          * CERTIFY_MODEL command was not executed for the selected compliance	
          * incorrect VID  /PID/Software Version/Certification Type	
          * VID, PID, Software Version values ​​are correct, Certification Type values ​​are incorrect	
          * VID, PID, Certification Type values ​​are correct, Software Version values ​​are incorrect	
          * VID, Software Version, Certification Type values ​​are correct, PID values ​​are incorrect	
          * Software Versio, PID, Certification Type values ​​are correct, VID values ​​are incorrect	

* Role (Who can send)
     * Positive:
          * CertificationCenter
     * Negative:
          * Trustee	
          * Vendor 	
          * VendorAdmin 	 	
          * NodeAdmin 	

* Parameters:	
     * vid (Vendor ID) - uint16
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	MAX = 65535
     * pid (Product ID) - uint16 
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	MAX = 65535
     * softwareVersion (Software Version) - uint32 
          * Positive:	
               * value exists	
               * value >= 0	
               * integer value format
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	MAX = 4 294 967 295
     * certificationType (Certification Type) - string 
          * Positive:	
               * existing value	
               * valid type	
                    * zigbee	
                    * matter	
                    * access control	
                    * product security	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * certificationDate (Certification Date) - optional(string)
          * Positive:	
               * string matches the format > 2019-10-12T07:20:50.52Z
               * text value format
               * MIN < length < MAX	
               * empty value	
          * Negative:	
               * nonexistent value	
               * length > MAX	
     * cdCertificateId (CD Certificate ID) - string 
          * Positive:	
               * existing value	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * reason (Reason) - optional(string)
          * Positive:	
               * empty value	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * cDVersionNumber (CD Version Number) - optional(uint32)
          * Positive:	
               * value exists	
               * value > 0	
               * empty value	
               * format	
          * Negative:
               * value =< 0
               * string value format		
               * length > MAX	(MAX = 4 294 967 295)
               * CD version number does not match model version	
     * familyId (Family ID) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * supportedClusters (Supported Clusters) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * compliantPlatformUsed (Compliant Platform Used) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * compliantPlatformVersion (Compliant Platform Version) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * OSVersion (OS Version) = optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * certificationRoute (Certification Route) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * programType (Program Type) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * programTypeVersion (Program Type Version) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * transport (Transport) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * parentChild (Parent Child) - optional(string)
          * Positive:	
               * supported types	
                    * parent	
                    * child	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
     * length > MAX	
     * certificationIDOfSoftwareComponent (Certification ID Of Software Component) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * schemaVersion (Schema Version) - optional(uint16)
          * Positive:	
               * value = 0	
               * integer value format
               * empty value	
          * Negative:	
               * length > MAX	(MAX = 65535)

### REST API 	
REST API command: POST /cosmos/tx/v1beta1/txs: [MsgUpdateComplianceInfo](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/proto/zigbeealliance/distributedcomplianceledger/compliance/tx.proto#L93)

* REST API command send	
     * Valid command	
          * correct HTTP method	
          * request is authorized	
          * uses valid credentials/role	
     * Invalid command	
          * incorrect request	
          * server side error	

* Сommand result	
     * UPDATE_COMPLIANCE_INFO command completed successfully ⇒ updates a compliance info by VID, PID, Software Version and Certification Type
          * CERTIFY_MODEL command completed successfully for the selected compliance	
          * VID, PID, Software Version and Certification Type value are indicated correctly	
     * UPDATE_COMPLIANCE_INFO command failed	⇒ does not updates a compliance info by VID, PID, Software Version and Certification Type
          * CERTIFY_MODEL command was not executed for the selected compliance	
          * incorrect VID  /PID/Software Version/Certification Type	
          * VID, PID, Software Version values ​​are correct, Certification Type values ​​are incorrect	
          * VID, PID, Certification Type values ​​are correct, Software Version values ​​are incorrect	
          * VID, Software Version, Certification Type values ​​are correct, PID values ​​are incorrect	
          * Software Versio, PID, Certification Type values ​​are correct, VID values ​​are incorrect	

* Role (Who can send)
     * Positive:
          * CertificationCenter
     * Negative:
          * Trustee	
          * Vendor 	
          * VendorAdmin 	 	
          * NodeAdmin 	

* Parameters:	
     * vid (Vendor ID) - uint16
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	MAX = 65535
     * pid (Product ID) - uint16 
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	MAX = 65535
     * softwareVersion (Software Version) - uint32 
          * Positive:	
               * value exists	
               * value >= 0	
               * integer value format
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	MAX = 4 294 967 295
     * certificationType (Certification Type) - string 
          * Positive:	
               * existing value	
               * valid type	
                    * zigbee	
                    * matter	
                    * access control	
                    * product security	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * certificationDate (Certification Date) - optional(string)
          * Positive:	
               * string matches the format > 2019-10-12T07:20:50.52Z
               * text value format
               * MIN < length < MAX	
               * empty value	
          * Negative:	
               * nonexistent value	
               * length > MAX	
     * cdCertificateId (CD Certificate ID) - string 
          * Positive:	
               * existing value	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * reason (Reason) - optional(string)
          * Positive:	
               * empty value	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * cDVersionNumber (CD Version Number) - optional(uint32)
          * Positive:	
               * value exists	
               * value > 0	
               * empty value	
               * format	
          * Negative:
               * value =< 0
               * string value format		
               * length > MAX	(MAX = 4 294 967 295)
               * CD version number does not match model version	
     * familyId (Family ID) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * supportedClusters (Supported Clusters) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * compliantPlatformUsed (Compliant Platform Used) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * compliantPlatformVersion (Compliant Platform Version) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * OSVersion (OS Version) = optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * certificationRoute (Certification Route) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * programType (Program Type) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * programTypeVersion (Program Type Version) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * transport (Transport) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * parentChild (Parent Child) - optional(string)
          * Positive:	
               * supported types	
                    * parent	
                    * child	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
     * length > MAX	
     * certificationIDOfSoftwareComponent (Certification ID Of Software Component) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * schemaVersion (Schema Version) - optional(uint16)
          * Positive:	
               * value = 0	
               * integer value format
               * empty value	
          * Negative:	
               * length > MAX	(MAX = 65535)

## [DELETE_COMPLIANCE_INFO](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/docs/transactions/compliance.md#delete_compliance_info)	
### CLI command	
CLI command: `dcld tx compliance delete-compliance-info --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee|access control|product security> --from=<account>`

* CLI command send	
     * Valid command	
          * command exists/relevant	
     * Invalid command	
          * access is denied to execute command	
          * incorrect command syntax	

* Сommand result	
     * DELETE_COMPLIANCE_INFO command completed successfully ⇒ delete compliance of the Model Version to the ZB or Matter standard
          * CERTIFY_MODEL command completed successfully for the selected compliance	
          * there is at least one compliance of the Model Version to the ZB or Matter standard	
          * corresponding Compliance Info is present on the ledger	
     * DELETE_COMPLIANCE_INFO command failed	⇒ does not delete compliance of the Model Version to the ZB or Matter standard
          * CERTIFY_MODEL command was not executed for the selected compliance	
          * there is not one compliance of the Model Version to the ZB or Matter standard	
          * corresponding Compliance Info is not present on the ledger	

* Role (Who can send)
     * Positive:
          * CertificationCenter 
     * Negative:	
          * Trustee	error
          * Vendor 	error
          * VendorAdmin 	error
          * NodeAdmin 	error

* Parameters:	
     * vid (Vendor ID) - uint16
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * pid (Product ID) - uint16 
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * softwareVersion (Software Version) - uint32 
          * Positive:	
               * value exists	
               * value >= 0	
               * integer value format
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	(MAX = 4 294 967 295)
     * certificationType (Certification Type) - string 
          * Positive:	
               * existing value	
               * valid type	
                    * zigbee	
                    * matter	
                    * access control	
                    * product security	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	

### REST API 	
REST API command: POST /cosmos/tx/v1beta1/txs: [MsgUpdateComplianceInfo](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/proto/zigbeealliance/distributedcomplianceledger/compliance/tx.proto#L121)

* REST API command send	
     * Valid command	
          * correct HTTP method	
          * request is authorized	
          * uses valid credentials/role	
     * Invalid command	
          * incorrect request	
          * server side error	

* Сommand result	
     * DELETE_COMPLIANCE_INFO command completed successfully ⇒ delete compliance of the Model Version to the ZB or Matter standard
          * CERTIFY_MODEL command completed successfully for the selected compliance	
          * there is at least one compliance of the Model Version to the ZB or Matter standard	
          * corresponding Compliance Info is present on the ledger	
     * DELETE_COMPLIANCE_INFO command failed	⇒ does not delete compliance of the Model Version to the ZB or Matter standard
          * CERTIFY_MODEL command was not executed for the selected compliance	
          * there is not one compliance of the Model Version to the ZB or Matter standard	
          * corresponding Compliance Info is not present on the ledger	

* Role (Who can send)
     * Positive:
          * CertificationCenter 
     * Negative:	
          * Trustee	error
          * Vendor 	error
          * VendorAdmin 	error
          * NodeAdmin 	error

* Parameters:	
     * vid (Vendor ID) - uint16
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * pid (Product ID) - uint16 
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * softwareVersion (Software Version) - uint32 
          * Positive:	
               * value exists	
               * value >= 0	
               * integer value format
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	(MAX = 4 294 967 295)
     * certificationType (Certification Type) - string 
          * Positive:	
               * existing value	
               * valid type	
                    * zigbee	
                    * matter	
                    * access control	
                    * product security	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	

## [REVOKE_MODEL_CERTIFICATION](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/docs/transactions/compliance.md#revoke_model_certification)	

### CLI command	
CLI command: `dcld tx compliance revoke-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee|access control|product security> --revocationDate=<rfc3339 encoded date> --reason=<string> --from=<account>`

* CLI command send	
     * Valid command	
          * command exists/relevant	
     * Invalid command	
          * access is denied to execute command	
          * incorrect command syntax	

* Сommand result	
     * REVOKE_MODEL_CERTIFICATION command completed successfully	⇒ revoke compliance of the Model Version to the ZB or Matter standard
          * CERTIFY_MODEL command completed successfully for the selected compliance	
          * corresponding Model and Model Version not present on the ledger	
          * corresponding Model and Model Version present on the ledger	
          * only corresponding Model present on the ledger	
          * only corresponding Model Version present on the ledger	
          * compliance result is written on the ledger	
          * only revocation list is stored on the ledger	
     * REVOKE_MODEL_CERTIFICATION command failed	⇒ does not revoke compliance of the Model Version to the ZB or Matter standard
          * CERTIFY_MODEL command was not executed for the selected compliance	
          * compliance result is not written on the ledger	
          * not only revocation list is stored on the ledger	

* Role (Who can send)
     * Positive:
          * CertificationCenter
     * Negative:	
          * Trustee	
          * Vendor 	
          * VendorAdmin 	
          * NodeAdmin 	

* Parameters:	
     * vid (Vendor ID) - uint16
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * pid (Product ID) - uint16 
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * softwareVersion (Software Version) - uint32 
          * Positive:	
               * value exists	
               * value >= 0	
               * integer value format
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	(MAX = 4 294 967 295)
     * softwareVersionSting (Software Version Sting) - string 
          * Positive:	
               * value exists	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * revocationDate (Revocation Date) - string 
          * Positive:	
               * format date	
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * certificationType (Certification Type) - string 
          * Positive:	
               * existing value	
               * valid type	
                    * zigbee	
                    * matter	
                    * access control	
                    * product security	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * reason (Reason) - optional(string)
          * Positive:	
               * text value format
               * empty value	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * schemaVersion (Schema Version) - optional(uint16)
          * Positive:	
               * value = 0	
               * integer value format
               * empty value	
          * Negative:	
               * length > MAX	(MAX = 65535)

### REST API 	
REST API command: POST /cosmos/tx/v1beta1/txs: [MsgRevokeModel](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/proto/zigbeealliance/distributedcomplianceledger/compliance/tx.proto#L49)

* REST API command send	
     * Valid command	
          * correct HTTP method	
          * request is authorized	
     * Invalid command	
          * incorrect request	
          * server side error	

* Сommand result	
     * REVOKE_MODEL_CERTIFICATION command completed successfully	⇒ revoke compliance of the Model Version to the ZB or Matter standard
          * CERTIFY_MODEL command completed successfully for the selected compliance	
          * corresponding Model and Model Version not present on the ledger	
          * corresponding Model and Model Version present on the ledger	
          * only corresponding Model present on the ledger	
          * only corresponding Model Version present on the ledger	
          * compliance result is written on the ledger	
          * only revocation list is stored on the ledger	
     * REVOKE_MODEL_CERTIFICATION command failed	⇒ does not revoke compliance of the Model Version to the ZB or Matter standard
          * CERTIFY_MODEL command was not executed for the selected compliance	
          * compliance result is not written on the ledger	
          * not only revocation list is stored on the ledger	

* Role (Who can send)
     * Positive:
          * CertificationCenter
     * Negative:	
          * Trustee	
          * Vendor 	
          * VendorAdmin 	
          * NodeAdmin 	

* Parameters:	
     * vid (Vendor ID) - uint16
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * pid (Product ID) - uint16 
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * softwareVersion (Software Version) - uint32 
          * Positive:	
               * value exists	
               * value >= 0	
               * integer value format
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	(MAX = 4 294 967 295)
     * softwareVersionSting (Software Version Sting) - string 
          * Positive:	
               * value exists	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * revocationDate (Revocation Date) - string 
          * Positive:	
               * format date	
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * certificationType (Certification Type) - string 
          * Positive:	
               * existing value	
               * valid type	
                    * zigbee	
                    * matter	
                    * access control	
                    * product security	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * reason (Reason) - optional(string)
          * Positive:	
               * text value format
               * empty value	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * schemaVersion (Schema Version) - optional(uint16)
          * Positive:	
               * value = 0	
               * integer value format
               * empty value	
          * Negative:	
               * length > MAX	(MAX = 65535)

## [PROVISION_MODEL](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/docs/transactions/compliance.md#provision_model)	

### CLI command	
CLI command: `dcld tx compliance provision-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<matter|zigbee|access control|product security> --provisionalDate=<rfc3339 encoded date> --from=<account>`

* CLI command send	
     * Valid command	
          * command exists/relevant	
     * Invalid command	
          * access is denied to execute command	
          * incorrect command syntax	

* Сommand result	
     * PROVISION_MODEL command completed successfully ⇒ sets provisional state for the Model Version
          * ADD_MODEL and ADD_MODEL_VERSION command completed successfully	
          * corresponding Model and Model Version is present on the ledger	
          * certification record is missing from on the ledger (certified or revoked)	
     * PROVISION_MODEL command failed ⇒ does not sets provisional state for the Model Version
          * ADD_MODEL command was not executed	
          * record about a model has been removed	
          * ADD_MODEL command executed, but ADD_MODEL_VERSION command was not executed	
          * record about a model version has been removed	
          * corresponding Model not present on the ledger	
          * corresponding Model is present on the ledger, but  corresponding Model Version not present on the ledger	
          * there is already a certification record on the ledger (certified)	
          * there is already a certification record on the ledger (revoked)	

* Role (Who can send)
     * Positive:
          * CertificationCenter 	
     * Negative:	
          * Trustee	error
          * Vendor 	error
          * VendorAdmin 	error
          * NodeAdmin 	error

* Parameters:	
     * vid (Vendor ID) - uint16
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * pid (Product ID) - uint16 
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * softwareVersion (Software Version) - uint32 
          * Positive:	
               * value exists	
               * value >= 0	
               * integer value format
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	(MAX = 4 294 967 295)
     * softwareVersionSting (Software Version Sting) - string 
          * Positive:	
               * value exists	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * provisionalDate (Provisional Date) - string 
          * Positive:	
               * string matches the format	2019-10-12T07:20:50.52Z
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * certificationType (Certification Type) - string 
          * Positive:	
               * existing value	
               * valid type	
                    * zigbee	
                    * matter	
                    * access control	
                    * product security	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * cdCertificateId (CD Certificate ID) - string 
          * Positive:	
               * existing value	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * reason (Reason) - optional(string)
          * Positive:	
               * empty value	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * cDVersionNumber (CD Version Number) - optional(uint32)
          * Positive:	
               * value exists	
               * value > 0	
               * empty value	
               * format	
          * Negative:
               * value =< 0
               * string value format		
               * length > MAX	(MAX = 4 294 967 295)
     * familyId (Family ID) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * supportedClusters (Supported Clusters) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * compliantPlatformUsed (Compliant Platform Used) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * compliantPlatformVersion (Compliant Platform Version) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * OSVersion (OS Version) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     certificationRoute (Certification Route) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * programType (Program Type) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * programTypeVersion (Program Type Version) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * transport (Transport) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * parentChild (Parent Child) - optional(string)
          * Positive:	
               * supported types	
                    * parent	
                    * child	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * certificationIDOfSoftwareComponent (Certification ID Of Software Component) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * schemaVersion (Schema Version) - optional(uint16)
          * Positive:	
               * value = 0	
               * integer value format
               * empty value	
          * Negative:	
               * length > MAX	(MAX = 65535)

### REST API 	
REST API command: POST /cosmos/tx/v1beta1/txs: [MsgRevokeModel](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/proto/zigbeealliance/distributedcomplianceledger/compliance/tx.proto#L65)

* REST API command send	
     * Valid command	
          * correct HTTP method	
          * request is authorized	
          * uses valid credentials/role	
     * Invalid command	
          * incorrect request	
          * server side error	

* Сommand result	
     * PROVISION_MODEL command completed successfully ⇒ sets provisional state for the Model Version
          * ADD_MODEL and ADD_MODEL_VERSION command completed successfully	
          * corresponding Model and Model Version is present on the ledger	
          * certification record is missing from on the ledger (certified or revoked)	
     * PROVISION_MODEL command failed ⇒ does not sets provisional state for the Model Version
          * ADD_MODEL command was not executed	
          * record about a model has been removed	
          * ADD_MODEL command executed, but ADD_MODEL_VERSION command was not executed	
          * record about a model version has been removed	
          * corresponding Model not present on the ledger	
          * corresponding Model is present on the ledger, but  corresponding Model Version not present on the ledger	
          * there is already a certification record on the ledger (certified)	
          * there is already a certification record on the ledger (revoked)	

* Role (Who can send)
     * Positive:
          * CertificationCenter 	
     * Negative:	
          * Trustee	error
          * Vendor 	error
          * VendorAdmin 	error
          * NodeAdmin 	error

* Parameters:	
     * vid (Vendor ID) - uint16
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * pid (Product ID) - uint16 
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * softwareVersion (Software Version) - uint32 
          * Positive:	
               * value exists	
               * value >= 0	
               * integer value format
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	(MAX = 4 294 967 295)
     * softwareVersionSting (Software Version Sting) - string 
          * Positive:	
               * value exists	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * provisionalDate (Provisional Date) - string 
          * Positive:	
               * string matches the format	2019-10-12T07:20:50.52Z
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * certificationType (Certification Type) - string 
          * Positive:	
               * existing value	
               * valid type	
                    * zigbee	
                    * matter	
                    * access control	
                    * product security	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * cdCertificateId (CD Certificate ID) - string 
          * Positive:	
               * existing value	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	
     * reason (Reason) - optional(string)
          * Positive:	
               * empty value	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * cDVersionNumber (CD Version Number) - optional(uint32)
          * Positive:	
               * value exists	
               * value > 0	
               * empty value	
               * format	
          * Negative:
               * value =< 0
               * string value format		
               * length > MAX	(MAX = 4 294 967 295)
     * familyId (Family ID) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * supportedClusters (Supported Clusters) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * compliantPlatformUsed (Compliant Platform Used) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * compliantPlatformVersion (Compliant Platform Version) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * OSVersion (OS Version) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     certificationRoute (Certification Route) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * programType (Program Type) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * programTypeVersion (Program Type Version) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * transport (Transport) - optional(string)
          * Positive:	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * parentChild (Parent Child) - optional(string)
          * Positive:	
               * supported types	
                    * parent	
                    * child	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * certificationIDOfSoftwareComponent (Certification ID Of Software Component) - optional(string)
          * Positive:	
               * value exists	
               * empty value	
               * format	
               * MIN < length < MAX	
          * Negative:	
               * length > MAX	
     * schemaVersion (Schema Version) - optional(uint16)
          * Positive:	
               * value = 0	
               * integer value format
               * empty value	
          * Negative:	
               * length > MAX	(MAX = 65535)

## [GET_CERTIFIED_MODEL]( https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/docs/transactions/compliance.md#get_certified_model)	

### CLI command	
CLI command: `dcld query compliance certified-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter|access control|product security>`

* CLI command send	
     * Valid command	
          * command exists/relevant	
     * Invalid command	
          * incorrect command syntax	

* Сommand result	
     * GET_CERTIFIED_MODEL command completed successfully ⇒ gets a structure containing the Model Version / Certification Type key (vid, pid, softwareVersion, certificationType) and a flag (value) indicating whether the given Model Version is compliant to certificationType standard
          * ADD_MODEL and ADD_MODEL_VERSION command completed successfully	
          * a record of a Model and Model Version is present on the ledger	
          * compliance is tracked on the ledger	
          * Model Version was certified earlier	
          * compliance information is found on ledger and it's in certified state	
     * GET_CERTIFIED_MODEL command failed ⇒ does not gets a structure containing the Model Version / Certification Type key (vid, pid, softwareVersion, certificationType) and a flag (value) indicating whether the given Model Version is compliant to certificationType standard
          * ADD_MODEL command was not executed	
          * record about a model has been removed	
          * ADD_MODEL command executed, but ADD_MODEL_VERSION command was not executed	
          * record about a model version has been removed	
          * a record of a Model Version is not present on the ledger	
          * compliance is not tracked on the ledger	
          * compliance  information is not found on ledger	
          * compliance information is found on ledger but it is not certified state	
          * Model Version was never certified earlier	

* Role (Who can send)
     * Positive:	
          * Trustee	
          * Vendor 	
          * VendorAdmin 	
          * CertificationCenter 	
          * NodeAdmin 	

* Parameters:	
     * vid (Vendor ID) - uint16
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * pid (Product ID) - uint16 
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * softwareVersion (Software Version) - uint32 
          * Positive:	
               * value exists	
               * value >= 0	
               * integer value format
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	(MAX = 4 294 967 295)
     * certificationType (Certification Type) - string 
          * Positive:	
               * existing value	
               * valid type	
                    * zigbee	
                    * matter	
                    * access control	
                    * product security	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	

### REST API
REST API command: `GET /dcl/compliance/certified-models/{vid}/{pid}/{software_version}/{certification_type}` 	

* REST API command send	
     * Valid command	
          * correct HTTP method	
          * request is authorized	
          * uses valid credentials/role	
     * Invalid command	
          * incorrect request	
          * server side error	

* Сommand result	
     * GET_CERTIFIED_MODEL command completed successfully ⇒ gets a structure containing the Model Version / Certification Type key (vid, pid, softwareVersion, certificationType) and a flag (value) indicating whether the given Model Version is compliant to certificationType standard
          * ADD_MODEL and ADD_MODEL_VERSION command completed successfully	
          * a record of a Model and Model Version is present on the ledger	
          * compliance is tracked on the ledger	
          * Model Version was certified earlier	
          * compliance information is found on ledger and it's in certified state	
     * GET_CERTIFIED_MODEL command failed ⇒ does not gets a structure containing the Model Version / Certification Type key (vid, pid, softwareVersion, certificationType) and a flag (value) indicating whether the given Model Version is compliant to certificationType standard
          * ADD_MODEL command was not executed	
          * record about a model has been removed	
          * ADD_MODEL command executed, but ADD_MODEL_VERSION command was not executed	
          * record about a model version has been removed	
          * a record of a Model Version is not present on the ledger	
          * compliance is not tracked on the ledger	
          * compliance  information is not found on ledger	
          * compliance information is found on ledger but it is not certified state	
          * Model Version was never certified earlier	

* Role (Who can send)
     * Positive:	
          * Trustee	
          * Vendor 	
          * VendorAdmin 	
          * CertificationCenter 	
          * NodeAdmin 	

* Parameters:	
     * vid (Vendor ID) - uint16
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * pid (Product ID) - uint16 
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * softwareVersion (Software Version) - uint32 
          * Positive:	
               * value exists	
               * value >= 0	
               * integer value format
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	(MAX = 4 294 967 295)
     * certificationType (Certification Type) - string 
          * Positive:	
               * existing value	
               * valid type	
                    * zigbee	
                    * matter	
                    * access control	
                    * product security	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	

## [GET_REVOKED_MODEL](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/docs/transactions/compliance.md#get_revoked_model)	

### CLI command	
CLI command: `dcld query compliance revoked-model --vid=<uint16> --pid=<uint16> --softwareVersion=<uint32> --certificationType=<zigbee|matter|access control|product security>`

* CLI command send	
     * Valid command	
          * command exists/relevant	
     * Invalid command	
          * incorrect command syntax	

* Сommand result	
     * GET_REVOKED_MODEL command completed successfully ⇒ gets a structure containing the Model Version / Certification Type key (vid, pid, softwareVersion, certificationType) and a flag (value) indicating whether the given Model Version is revoked for certificationType standard
          * ADD_MODEL and ADD_MODEL_VERSION command completed successfully	
          * a record of a Model and Model Version is present on the ledger	
          * only revocation is tracked on the ledger	
          * Model Version was certified or revoked earlier	
          * compliance information is found on ledger and it's in revoked state	
     * GET_REVOKED_MODEL command failed	⇒ does not gets a structure containing the Model Version / Certification Type key (vid, pid, softwareVersion, certificationType) and a flag (value) indicating whether the given Model Version is revoked for certificationType standard
          * ADD_MODEL command was not executed	
          * record about a model has been removed	
          * ADD_MODEL command executed, but ADD_MODEL_VERSION command was not executed	
          * record about a model version has been removed	
          * not only revocation is tracked on the ledger	
          * Model Version was never certified or revoked earlier	
          * compliance information is not found on ledger	
          * compliance information is found on ledger but it is not revoked state	

* Role (Who can send)	
     * Positive:
          * Trustee	
          * Vendor 	
          * VendorAdmin 	
          * CertificationCenter 	
          * NodeAdmin 	

* Parameters:	
     * vid (Vendor ID) - uint16
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * pid (Product ID) - uint16 
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * softwareVersion (Software Version) - uint32 
          * Positive:	
               * value exists	
               * value >= 0	
               * integer value format
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	(MAX = 4 294 967 295)
     * certificationType (Certification Type) - string 
          * Positive:	
               * existing value	
               * valid type	
                    * zigbee	
                    * matter	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	

### REST API 	
REST API command: `GET /dcl/compliance/revoked-models/{vid}/{pid}/{software_version}/{certification_type}` 

* REST API command send	
     * Valid command	
          * correct HTTP method	
          * request is authorized	
          * uses valid credentials/role	
     * Invalid command	
          * incorrect request	
          * server side error	

* Сommand result	
     * GET_REVOKED_MODEL command completed successfully ⇒ gets a structure containing the Model Version / Certification Type key (vid, pid, softwareVersion, certificationType) and a flag (value) indicating whether the given Model Version is revoked for certificationType standard
          * ADD_MODEL and ADD_MODEL_VERSION command completed successfully	
          * a record of a Model and Model Version is present on the ledger	
          * only revocation is tracked on the ledger	
          * Model Version was certified or revoked earlier	
          * compliance information is found on ledger and it's in revoked state	
     * GET_REVOKED_MODEL command failed	⇒ does not gets a structure containing the Model Version / Certification Type key (vid, pid, softwareVersion, certificationType) and a flag (value) indicating whether the given Model Version is revoked for certificationType standard
          * ADD_MODEL command was not executed	
          * record about a model has been removed	
          * ADD_MODEL command executed, but ADD_MODEL_VERSION command was not executed	
          * record about a model version has been removed	
          * not only revocation is tracked on the ledger	
          * Model Version was never certified or revoked earlier	
          * compliance information is not found on ledger	
          * compliance information is found on ledger but it is not revoked state	

* Role (Who can send)	
     * Positive:
          * Trustee	
          * Vendor 	
          * VendorAdmin 	
          * CertificationCenter 	
          * NodeAdmin 	

* Parameters:	
     * vid (Vendor ID) - uint16
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * pid (Product ID) - uint16 
          * Positive:	
               * unique value	
               * value > 0	
               * integer value format
          * Negative:	
               * empty value
               * value =< 0
               * string value format		
               * nonexistent value	
               * length > MAX	(MAX = 65535)
     * softwareVersion (Software Version) - uint32 
          * Positive:	
               * value exists	
               * value >= 0	
               * integer value format
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX	(MAX = 4 294 967 295)
     * certificationType (Certification Type) - string 
          * Positive:	
               * existing value	
               * valid type	
                    * zigbee	
                    * matter	
               * text value format
               * MIN < length < MAX	
          * Negative:	
               * empty value	
               * nonexistent value	
               * length > MAX

## [GET_PROVISIONAL_MODEL]()	

### CLI command	

* CLI command send	
* Valid command	
* command exists/relevant	
* Invalid command	
* access is denied to execute command	
* incorrect command syntax	
* Сommand result	
* GET_PROVISIONAL_MODEL command completed successfully	gets a structure containing the Model Version / Certification Type key (vid, pid, softwareVersion, certificationType) and a flag (value) indicating whether the given Model Version is revoked for certificationType standard
* ADD_MODEL and ADD_MODEL_VERSION command completed successfully	
* a record of a Model and Model Version is present on the ledger	
* Model Version was certified or revoked earlier	
* compliance information is found on the ledger and it's in provisional state	
* GET_PROVISIONAL_MODEL command failed	does not gets a structure containing the Model Version / Certification Type key (vid, pid, softwareVersion, certificationType) and a flag (value) indicating whether the given Model Version is revoked for certificationType standard
* ADD_MODEL command was not executed	
* record about a model has been removed	
* ADD_MODEL command executed, but ADD_MODEL_VERSION command was not executed	
* record about a model version has been removed	
* Model Version was never certified or revoked earlier	
* compliance information is not found on ledger	
* compliance information is found on ledger but it is not provisional state	
* Role (Who can send)	
* Trustee	
* Vendor 	
* VendorAdmin 	
* CertificationCenter 	
* NodeAdmin 	
* Parameters:	
* vid (Vendor ID)	uint16
     * Positive:	
* unique value	
* value > 0	
* integer value format
     * Negative:	
* empty value
* value =< 0
* string value format		
* nonexistent value	
* length > MAX	MAX = 65535
* pid (Product ID)	uint16 
     * Positive:	
* unique value	
* value > 0	
* integer value format
     * Negative:	
* empty value
* value =< 0
* string value format		
* nonexistent value	
* length > MAX	MAX = 65535
* softwareVersion (Software Version)	uint32 
     * Positive:	
* value exists	
* value >= 0	
* integer value format
     * Negative:	
* empty value	
* nonexistent value	
* length > MAX	MAX = 4 294 967 295
* certificationType (Certification Type)	string 
     * Positive:	
* existing value	
* valid type	
* zigbee	
* matter	
* text value format
* MIN < length < MAX	
     * Negative:	
* empty value	
* nonexistent value	
* length > MAX	

### REST API 	

* REST API command send	
* Valid command	
* correct HTTP method	
* request is authorized	
* uses valid credentials/role	
* Invalid command	
* incorrect request	
* server side error	

Сommand result	
GET_PROVISIONAL_MODEL command completed successfully	gets a structure containing the Model Version / Certification Type key (vid, pid, softwareVersion, certificationType) and a flag (value) indicating whether the given Model Version is revoked for certificationType standard
ADD_MODEL and ADD_MODEL_VERSION command completed successfully	
a record of a Model and Model Version is present on the ledger	
Model Version was certified or revoked earlier	
compliance information is found on the ledger and it's in provisional state	
GET_PROVISIONAL_MODEL command failed	does not gets a structure containing the Model Version / Certification Type key (vid, pid, softwareVersion, certificationType) and a flag (value) indicating whether the given Model Version is revoked for certificationType standard
ADD_MODEL command was not executed	
record about a model has been removed	
ADD_MODEL command executed, but ADD_MODEL_VERSION command was not executed	
record about a model version has been removed	
Model Version was never certified or revoked earlier	
compliance information is not found on ledger	
compliance information is found on ledger but it is not provisional state	
Role (Who can send)	
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
pid (Product ID)	uint16 
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
softwareVersion (Software Version)	uint32 
     * Positive:	
value exists	
value >= 0	
integer value format
     * Negative:	
empty value	
nonexistent value	
length > MAX	MAX = 4 294 967 295
certificationType (Certification Type)	string 
     * Positive:	
existing value	
valid type	
zigbee	
matter	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	

## GET_COMPLIANCE_INFO	

### CLI command	

* CLI command send	
* Valid command	
* command exists/relevant	
* Invalid command	
* access is denied to execute command	
* incorrect command syntax	
* Сommand result	
* GET_COMPLIANCE_INFO command completed successfully	compliance information associated with the Model Version and Certification Type (identified by the vid, pid, softwareVersion and certification_type)
* CERTIFY_MODEL command completed successfully for the selected compliance	
* compliance information is found in store	
* GET_COMPLIANCE_INFO command failed	does not gets compliance information associated with the Model Version and Certification Type (identified by the vid, pid, softwareVersion and certification_type)
* CERTIFY_MODEL command was not executed for the selected compliance	
* compliance information is not found in store	
* Role (Who can send)	
* Trustee	
* Vendor 	
* VendorAdmin 	
* CertificationCenter 	
* NodeAdmin 	
* Parameters:	
* vid (Vendor ID)	uint16
     * Positive:	
* unique value	
* value > 0	
* integer value format
     * Negative:	
* empty value
* value =< 0
* string value format		
* nonexistent value	
* length > MAX	MAX = 65535
* pid (Product ID)	uint16 
     * Positive:	
* unique value	
* value > 0	
* integer value format
     * Negative:	
* empty value	
* value =< 0
* string value format	
* nonexistent value	
* length > MAX	MAX = 65535
* softwareVersion (Software Version)	uint32 
     * Positive:	
* value exists	
* value >= 0	
* integer value format
     * Negative:	
* empty value	
* nonexistent value	
* length > MAX	MAX = 4 294 967 295
* certificationType (Certification Type)	string 
     * Positive:	
* existing value	
* valid type	
* zigbee	
* matter	
* text value format
* MIN < length < MAX	
     * Negative:	
* empty value	
* nonexistent value	
* length > MAX	

### REST API 	

* REST API command send	
* Valid command	
* correct HTTP method	
* request is authorized	
* uses valid credentials/role	
* Invalid command	
* incorrect request	
* server side error	

Сommand result	
GET_COMPLIANCE_INFO command completed successfully	compliance information associated with the Model Version and Certification Type (identified by the vid, pid, softwareVersion and certification_type)
CERTIFY_MODEL command completed successfully for the selected compliance	
compliance information is found in store	
GET_COMPLIANCE_INFO command failed	does not gets compliance information associated with the Model Version and Certification Type (identified by the vid, pid, softwareVersion and certification_type)
CERTIFY_MODEL command was not executed for the selected compliance	
compliance information is not found in store	
Role (Who can send)	
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
pid (Product ID)	uint16 
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
softwareVersion (Software Version)	uint32 
     * Positive:	
value exists	
value >= 0	
integer value format
     * Negative:	
empty value	
nonexistent value	
length > MAX	MAX = 4 294 967 295
certificationType (Certification Type)	string 
     * Positive:	
existing value	
valid type	
zigbee	
matter	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	

## GET_DEVICE_SOFTWARE_COMPLIANCE	

### CLI command	

* CLI command send	
* Valid command	
* command exists/relevant	
* Invalid command	
* access is denied to execute command	
* incorrect command syntax	
* Сommand result	
* GET_COMPLIANCE_INFO command completed successfully	gets device software compliance associated with the cDCertificateId
* device software compliance is found in store	
* cDCertificateId value is correct	
* GET_COMPLIANCE_INFO command failed	does not gets device software compliance associated with the cDCertificateId
* cDCertificateId value is incorrect	
* device software compliance is not found in store	
* Role (Who can send)	
* Trustee	
* Vendor 	
* VendorAdmin 	
* CertificationCenter 	
* NodeAdmin 	
* Parameters:	
* cdCertificateId (CD Certificate ID)	string 
     * Positive:	
* existing value	
* text value format
* MIN < length < MAX	
     * Negative:	
* empty value	
* nonexistent value	
* length > MAX	

### REST API 	

* REST API command send	
* Valid command	
* correct HTTP method	
* request is authorized	
* uses valid credentials/role	
* Invalid command	
* incorrect request	
* server side error	

Сommand result	
GET_COMPLIANCE_INFO command completed successfully	gets device software compliance associated with the cDCertificateId
device software compliance is found in store	
cDCertificateId value is correct	
GET_COMPLIANCE_INFO command failed	does not gets device software compliance associated with the cDCertificateId
cDCertificateId value is incorrect	
device software compliance is not found in store	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
cdCertificateId (CD Certificate ID)	string 
     * Positive:	
existing value	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	

## GET_ALL_CERTIFIED_MODELS	

### CLI command	

* CLI command send	
* Valid command	
* command exists/relevant	
* Invalid command	
* access is denied to execute command	
* incorrect command syntax	
* Сommand result	
* GET_ALL_CERTIFIED_MODELS command completed successfully	gets all compliant Model Versions for all vendors (vids)
* compliance is tracked on ledger	
* there is at least one compliant Model Versions for all vendors	
* GET_ALL_CERTIFIED_MODELS command failed	does not gets all compliant Model Versions for all vendors (vids)
* compliance is not tracked on ledger	
* there is not one compliant Model Versions for all vendors	
* Role (Who can send)	
* Trustee	error
* Vendor 	error
* VendorAdmin 	error
* CertificationCenter 	
* NodeAdmin 	error
* Parameters:	
* count-total	optional(bool)
     * Positive:	
* empty value	
* value state	
* TRUE (-1)	
* FALSE (0)	
     * Negative:	
* value is not bool	
* limit 	optional(uint)
     * Positive:	
* value exists	
* empty value	
     * Negative:	
* value > 100	
* offset 	optional(uint)
     * Positive:	
* value exists	
* empty value	
     * Negative:	
* value < 0	
* page 	optional(uint)
     * Positive:	
* value exists	
* empty value	
     * Negative:	
* value < 0	
* page-key	optional(string)
     * Positive:	
* empty value	
* value exists	
     * Negative:	
* length < MIN	
* reverse	optional(bool)
     * Positive:	
* empty value	
* value state	
* TRUE (-1)	
* FALSE (0)	
     * Negative:	
* value is not bool	

### REST API 	

* REST API command send	
* Valid command	
* correct HTTP method	
* request is authorized	
* uses valid credentials/role	
* Invalid command	
* incorrect request	
* server side error	

Сommand result	
GET_ALL_CERTIFIED_MODELS command completed successfully	gets all compliant Model Versions for all vendors (vids)
compliance is tracked on ledger	
there is at least one compliant Model Versions for all vendors	
GET_ALL_CERTIFIED_MODELS command failed	does not gets all compliant Model Versions for all vendors (vids)
compliance is not tracked on ledger	
there is not one compliant Model Versions for all vendors	
Role (Who can send)	
Trustee	error
Vendor 	error
VendorAdmin 	error
CertificationCenter 	
NodeAdmin 	error
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

## GET_ALL_REVOKED_MODELS	

### CLI command	

* CLI command send	
* Valid command	
* command exists/relevant	
* Invalid command	
* access is denied to execute command	
* incorrect command syntax	
* Сommand result	
* GET_ALL_REVOKED_MODELS command completed successfully	gets all revoked Model Versions for all vendors (vids)
* only revocation is tracked on the ledger	
* there is at least one compliant revoked Model Versions for all vendors	
* GET_ALL_REVOKED_MODELS command failed	does not gets all revoked Model Versions for all vendors (vids)
* not only revocation is tracked on the ledger	
* there is not one compliant revoked Model Versions for all vendors	
* Role (Who can send)	
* Trustee	error
* Vendor 	error
* VendorAdmin 	error
* CertificationCenter 	
* NodeAdmin 	error
* Parameters:	
* count-total	optional(bool)
     * Positive:	
* empty value	
* value state	
* TRUE (-1)	
* FALSE (0)	
     * Negative:	
* value is not bool	
* limit 	optional(uint)
     * Positive:	
* value exists	
* empty value	
     * Negative:	
* value > 100	
* offset 	optional(uint)
     * Positive:	
* value exists	
* empty value	
     * Negative:	
* value < 0	
* page 	optional(uint)
     * Positive:	
* value exists	
* empty value	
     * Negative:	
* value < 0	
* page-key	optional(string)
     * Positive:	
* empty value	
* value exists	
     * Negative:	
* length < MIN	
* reverse	optional(bool)
     * Positive:	
* empty value	
* value state	
* TRUE (-1)	
* FALSE (0)	
     * Negative:	
* value is not bool	

### REST API 	

* REST API command send	
* Valid command	
* correct HTTP method	
* request is authorized	
* uses valid credentials/role	
* Invalid command	
* incorrect request	
* server side error	
*
Сommand result	
GET_ALL_REVOKED_MODELS command completed successfully	gets all revoked Model Versions for all vendors (vids)
only revocation is tracked on the ledger	
there is at least one compliant revoked Model Versions for all vendors	
GET_ALL_REVOKED_MODELS command failed	does not gets all revoked Model Versions for all vendors (vids)
not only revocation is tracked on the ledger	
there is not one compliant revoked Model Versions for all vendors	
Role (Who can send)	
Trustee	error
Vendor 	error
VendorAdmin 	error
CertificationCenter 	
NodeAdmin 	error
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

## GET_ALL_PROVISIONAL_MODELS	

### CLI command	

* CLI command send	
* Valid command	
* command exists/relevant	
* Invalid command	
* access is denied to execute command	
* incorrect command syntax	
* Сommand result	
* GET_ALL_PROVISIONAL_MODELS command completed successfully	gets all Model Versions in provisional state for all vendors (vids)
* there is at least one Model Versions in provisional state for all vendors	
* GET_ALL_PROVISIONAL_MODELS command failed	does not gets all Model Versions in provisional state for all vendors (vids)
* there is at least one Model Versions in provisional state for all vendors	
* Role (Who can send)	
* Trustee	error
* Vendor 	error
* VendorAdmin 	error
* CertificationCenter 	
* NodeAdmin 	error
* Parameters:	
* count-total	optional(bool)
     * Positive:	
* empty value	
* value state	
* TRUE (-1)	
* FALSE (0)	
     * Negative:	
* value is not bool	
* limit 	optional(uint)
     * Positive:	
* value exists	
* empty value	
     * Negative:	
* value > 100	
* offset 	optional(uint)
     * Positive:	
* value exists	
* empty value	
     * Negative:	
* value < 0	
* page 	optional(uint)
     * Positive:	
* value exists	
* empty value	
     * Negative:	
* value < 0	
* page-key	optional(string)
     * Positive:	
* empty value	
* value exists	
     * Negative:	
* length < MIN	
* reverse	optional(bool)
     * Positive:	
* empty value	
* value state	
* TRUE (-1)	
* FALSE (0)	
     * Negative:	
* value is not bool	

### REST API 	

* REST API command send	
* Valid command	
* correct HTTP method	
* request is authorized	
* uses valid credentials/role	
* Invalid command	
* incorrect request	
* server side error	

Сommand result	
GET_ALL_PROVISIONAL_MODELS command completed successfully	gets all Model Versions in provisional state for all vendors (vids)
there is at least one Model Versions in provisional state for all vendors	
GET_ALL_PROVISIONAL_MODELS command failed	does not gets all Model Versions in provisional state for all vendors (vids)
there is at least one Model Versions in provisional state for all vendors	
Role (Who can send)	
Trustee	error
Vendor 	error
VendorAdmin 	error
CertificationCenter 	
NodeAdmin 	error
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

## GET_ALL_COMPLIANCE_INFO_RECORDS	

### CLI command	

* CLI command send	
* Valid command	
* command exists/relevant	
* Invalid command	
* access is denied to execute command	
* incorrect command syntax	
* Сommand result	
* GET_ALL_COMPLIANCE_INFO_RECORDS command completed successfully	gets all stored compliance information records for all vendors (vids)
* there is at least one stored compliance information records for all vendors	
* GET_ALL_COMPLIANCE_INFO_RECORDS command failed	does not gets all stored compliance information records for all vendors (vids)
* there is at least onestored compliance information records for all vendors	
* Role (Who can send)	
* Trustee	error
* Vendor 	error
* VendorAdmin 	error
* CertificationCenter 	
* NodeAdmin 	error
* Parameters:	
* count-total	optional(bool)
     * Positive:	
* empty value	
* value state	
* TRUE (-1)	
* FALSE (0)	
     * Negative:	
* value is not bool	
* limit 	optional(uint)
     * Positive:	
* value exists	
* empty value	
     * Negative:	
* value > 100	
* offset 	optional(uint)
     * Positive:	
* value exists	
* empty value	
     * Negative:	
* value < 0	
* page 	optional(uint)
     * Positive:	
* value exists	
* empty value	
     * Negative:	
* value < 0	
* page-key	optional(string)
     * Positive:	
* empty value	
* value exists	
     * Negative:	
* length < MIN	
* reverse	optional(bool)
     * Positive:	
* empty value	
* value state	
* TRUE (-1)	
* FALSE (0)	
     * Negative:	
* value is not bool	

### REST API 	

* REST API command send	
* Valid command	
* correct HTTP method	
* request is authorized	
* uses valid credentials/role	
* Invalid command	
* incorrect request	
* server side error	

Сommand result	
GET_ALL_COMPLIANCE_INFO_RECORDS command completed successfully	gets all stored compliance information records for all vendors (vids)
there is at least one stored compliance information records for all vendors	
GET_ALL_COMPLIANCE_INFO_RECORDS command failed	does not gets all stored compliance information records for all vendors (vids)
there is at least onestored compliance information records for all vendors	
Role (Who can send)	
Trustee	error
Vendor 	error
VendorAdmin 	error
CertificationCenter 	
NodeAdmin 	error
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

## GET_ALL_DEVICE_SOFTWARE_COMPLIANCES	

### CLI command	

* CLI command send	
* Valid command	
* command exists/relevant	
* Invalid command	
* access is denied to execute command	
* incorrect command syntax	
* Сommand result	
* GET_ALL_DEVICE_SOFTWARE_COMPLIANCES command completed successfully	gets all stored device software compliance's
* there is at least one stored device software compliance's	
* GET_ALL_DEVICE_SOFTWARE_COMPLIANCES command failed	does not gets all stored device software compliance's
* there is at least one stored device software compliance's	
* Role (Who can send)	
* Trustee	error
* Vendor 	error
* VendorAdmin 	error
* CertificationCenter 	
* NodeAdmin 	error
* Parameters:	
* count-total	optional(bool)
     * Positive:	
* empty value	
* value state	
* TRUE (-1)	
* FALSE (0)	
     * Negative:	
* value is not bool	
* limit 	optional(uint)
     * Positive:	
* value exists	
* empty value	
     * Negative:	
* value > 100	
* offset 	optional(uint)
     * Positive:	
* value exists	
* empty value	
     * Negative:	
* value < 0	
* page 	optional(uint)
     * Positive:	
* value exists	
* empty value	
     * Negative:	
* value < 0	
* page-key	optional(string)
     * Positive:	
* empty value	
* value exists	
     * Negative:	
* length < MIN	
* reverse	optional(bool)
     * Positive:	
* empty value	
* value state	
* TRUE (-1)	
* FALSE (0)	
     * Negative:	
* value is not bool	

### REST API 	
* REST API command send	
* Valid command	
* correct HTTP method	
* request is authorized	
* uses valid credentials/role	
* Invalid command	
* incorrect request	
* server side error	

Сommand result	
GET_ALL_DEVICE_SOFTWARE_COMPLIANCES command completed successfully	gets all stored device software compliance's
there is at least one stored device software compliance's	
GET_ALL_DEVICE_SOFTWARE_COMPLIANCES command failed	does not gets all stored device software compliance's
there is at least one stored device software compliance's	
Role (Who can send)	
Trustee	error
Vendor 	error
VendorAdmin 	error
CertificationCenter 	
NodeAdmin 	error
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