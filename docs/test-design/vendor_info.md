1. [ADD_VENDOR_INFO](#add_vendor_info)
2. [UPDATE_VENDOR_INFO](#update_vendor_info)
3. [GET_VENDOR_INFO](#get_vendor_info)
4. [GET_ALL_VENDOR_INFO](#get_all_vendor_info)

## [ADD_VENDOR_INFO](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/docs/transactions/vendor-info.md#add_vendor_info)

### CLI command
CLI command: `dcld tx vendorinfo add-vendor --vid=<uint16> --vendorName=<string> --companyLegalName=<string> --companyPreferredName=<string> --vendorLandingPageURL=<string> --from=<account>`

* CLI command send
   * Positive:
      * command exists/relevant
   * Negative:
      * access is denied to execute the command
      * multiple commands are sent with the same Vendor ID
      * incorrect command syntax

* Сommand result
   * Positive:
      * ADD_VENDOR_INFO command completed successfully **⇒** adds a record about a Vendor
   * Negative:
      * ADD_VENDOR_INFO command failed **⇒** does not add a record about a Vendor
 
* Role (Who can send)
   * Positive:
      * Vendor (vendor role the matching Vendor ID)
      * VendorAdmin 
   * Negative:
      * Trustee
      * Vendor (vendor role does not match Vendor ID)
      * CertificationCenter
      * NodeAdmin
  
* Parameters:
   * vid (Vendor ID) - uint16:
      * Positive:
         * value exists
         * value > 0
         * integer value format
      * Negative 
         * empty value
         * value =< 0
         * string value format
         * length > MAX (MAX = 65535)
         * nonexistent ID
   
   * vendorName (Vendor name) - string 
      * Positive:
         * text value format	
         * MIN < length < MAX	
      * Negative 
         * empty value
         * length > MAX (MAX = 128)
      
   * companyLegalName (Company Legal Name) - string 
      * Positive:
         * text value format	
         * MIN < length < MAX	
      * Negative
        * empty value	
        * length > MAX (MAX = 256)

   * companyPreferredName (Company Preferred Name)	optional(string)
      * Positive:
         * empty value	
         * text value format	
         * MIN < length < MAX	
      * Negative            
         * length > MAX
       
   * vendorLandingPageURL (Vendor Landing Page URL)	optional(string)
      * Positive:
         * value exists	
         * empty value	
         * text value format	
         * MIN < length < MAX	
      * Negative              
         * length > MAX	
         * сontains spaces or line breaks
         * value is not URL
         * URL starts not with https:
         * URL can't be non-http
   
   * schemaVersion (Schema Version)	optional(uint16)
      * Positive:
         * value = 0	
         * integer value format	
         * empty value	
      * Negative 	
         * length > MAX	(MAX = 65535)
         * value >< 0
         * string value format

### REST API 
POST: `/cosmos/tx/v1beta1/txs`[NewMsgCreateVendorInfo](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/proto/zigbeealliance/distributedcomplianceledger/vendorinfo/tx.proto#L18)

* REST API command send
   * Positive:
      * correct HTTP method
      * request is authorized
      * uses valid credentials/role
   * Negative:
      * incorrect request
      * multiple commands are sent with the same Vendor ID
      * server-side error

* Сommand result
   * Positive:
      * ADD_VENDOR_INFO command completed successfully **⇒** adds a record about a Vendor
   * Negative:
      * ADD_VENDOR_INFO command failed **⇒** does not add a record about a Vendor

* Role (Who can send)
   * Positive:
      * Vendor (vendor role the matching Vendor ID)
      * VendorAdmin 
   * Negative:
      * Trustee
      * Vendor (vendor role does not match Vendor ID)
      * CertificationCenter
      * NodeAdmin 

* Parameters:
   
   * vid (Vendor ID) - uint16:
      * Positive:
         * value exists
         * value > 0
         * integer value format
      * Negative 
         * empty value
         * value =< 0
         * string value format
         * length > MAX (MAX = 65535)
         * nonexistent ID
   
   * vendorName (Vendor name) - string 
      * Positive:
         * text value format	
         * MIN < length < MAX	
      * Negative 
         * empty value	
         * length > MAX (MAX = 128)	

   * companyLegalName (Company Legal Name) - string 
      * Positive:
         * text value format	
         * MIN < length < MAX	
      * Negative
        * empty value	
        * length > MAX (MAX = 256)

   * companyPreferredName (Company Preferred Name) - optional(string)
      * Positive:
         * empty value	
         * text value format	
         * MIN < length < MAX	
      * Negative            
         * length > MAX	
   
   * vendorLandingPageURL (Vendor Landing Page URL) - optional(string)
      * Positive:
         * value exists	
         * empty value	
         * text value format	
         * MIN < length < MAX	
      * Negative              
         * length > MAX	
         * сontains spaces or line breaks	
         * value is not URL
         * URL starts not with https:
         * URL can't be non-http
   
   * schemaVersion (Schema Version) - optional(uint16)
      * Positive:
         * value = 0	
         * integer value format	
         * empty value	
      * Negative 	
         * length > MAX	(MAX = 65535)
         * value >< 0
         * string value format

## [UPDATE_VENDOR_INFO](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/docs/transactions/vendor-info.md#update_vendor_info)

### CLI command
CLI command: `dcld tx vendorinfo update-vendor --vid=<uint16> ... --from=<account>`

* CLI command send
   * Positive:
      * command exists/relevant
   * Negative:
      * access is denied to execute the command
      * incorrect command syntax

* Сommand result
   * Positive:
      * UPDATE_VENDOR_INFO command completed successfully **⇒** updates a record about a Vendor
         * ADD_VENDOR_INFO command completed successfully
         * There is at least one record about a Vendor
   * Negative:
      * UPDATE_VENDOR_INFO command failed **⇒** does not update a record about a Vendor
         * ADD_VENDOR_INFO command was not executed
         * There is no one record about a Vendor

* Role (Who can send)
   * Positive:
      * Vendor (vendor role the matching Vendor ID)
      * VendorAdmin 
   * Negative:
      * Trustee
      * Vendor (vendor role does not match Vendor ID)
      * CertificationCenter
      * NodeAdmin 

* Parameters:
   
   * vid (Vendor ID) - uint16:
      * Positive:
         * value exists
         * value > 0
         * integer value format
      * Negative 
         * empty value
         * value =< 0
         * string value format
         * length > MAX (MAX = 65535)
         * nonexistent ID
   
   * vendorName (Vendor name) - string 
      * Positive:	
         * text value format	
         * MIN < length < MAX	
      * Negative 	
         * empty value
         * length > MAX	(MAX = 128)
   * companyLegalName (Company Legal Name) - string 
      * Positive:	
         * text value format	
         * MIN < length < MAX	
      * Negative 	
         * empty value
         * length > MAX	(MAX = 256)
   
   * companyPreferredName (Company Preferred Name) - optional(string)
      * Positive:	
         * empty value	
         * text value format	
         * MIN < length < MAX	
      * Negative 	
         * length > MAX	
   
   * vendorLandingPageURL (Vendor Landing Page URL) - optional(string)
      * Positive:	
         * value exists	
         * empty value	
         * text value format	
         * MIN < length < MAX	
      * Negative 	
         * length > MAX	
         * сontains spaces or line breaks
         * value is not URL
         * URL starts not with https:
         * URL can't be non-http	
   
   * schemaVersion (Schema Version) - optional(uint16)
      * Positive:	
         * value = 0	
         * integer value format	
         * empty value	
      * Negative 	
         * length > MAX	(MAX = 65535)
         * value >< 0
         * string value format

### REST API 
POST: `/cosmos/tx/v1beta1/txs`[MsgUpdateVendorInfo](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/proto/zigbeealliance/distributedcomplianceledger/vendorinfo/tx.proto#L30)

* REST API command send
   * Positive:
      * correct HTTP method
      * request is authorized
      * uses valid credentials/role
   * Negative:
      * incorrect request
      * server-side error

* Сommand result
   * Positive:
      * UPDATE_VENDOR_INFO command completed successfully **⇒** updates a record about a Vendor
         * ADD_VENDOR_INFO command completed successfully
         * There is at least one record about a Vendor
   * Negative:
      * UPDATE_VENDOR_INFO command failed **⇒** does not update a record about a Vendor
         * ADD_VENDOR_INFO command was not executed
         * There is no one record about a Vendor

* Role (Who can send)
   * Positive:
      * Vendor (vendor role the matching Vendor ID)
      * VendorAdmin 
   * Negative:
      * Trustee
      * Vendor (vendor role does not match Vendor ID)
      * CertificationCenter
      * NodeAdmin 

* Parameters:
   * vid (Vendor ID) - uint16:
      * Positive:
         * value exists
         * value > 0
         * integer value format
      * Negative 
         * empty value
         * value =< 0
         * string value format
         * length > MAX (MAX = 65535)
         * nonexistent ID
   
   * vendorName (Vendor name) - string 
      * Positive:	
         * text value format	
         * MIN < length < MAX	
      * Negative 	
         * empty value
         * length > MAX	(MAX = 128)
   
   * companyLegalName (Company Legal Name) - string 
      * Positive:	
         * text value format	
         * MIN < length < MAX	
      * Negative 	
         * empty value
         * length > MAX	
   
   * companyPreferredName (Company Preferred Name) - optional(string)
      * Positive:	
         * empty value	
         * text value format	
         * MIN < length < MAX	
      * Negative 	
         * length > MAX	(MAX = 256)
   
   * vendorLandingPageURL (Vendor Landing Page URL) - optional(string)
      * Positive:	
         * value exists	
         * empty value	
         * text value format	
         * MIN < length < MAX	
      * Negative 	
         * length > MAX	
         * сontains spaces or line breaks
         * value is not URL
         * URL starts not with https:
         * URL can't be non-http	
   
   * schemaVersion (Schema Version) - optional(uint16)
      * Positive:	
         * value = 0	
         * integer value format	
         * empty value	
      * Negative 	
         * length > MAX	(MAX = 65535)
         * value >< 0
         * string value format

## [GET_VENDOR_INFO](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/docs/transactions/vendor-info.md#get_vendor_info)

### CLI command
CLI command: `dcld query vendorinfo vendor --vid=<uint16>`

* CLI command send
   * Positive:
      * command exists/relevant
   * Negative:
      * incorrect command syntax

* Сommand result
   * Positive:
      * GET_VENDOR_INFO command completed successfully **⇒** gets a Vendor Info for the given vid (vendor ID)
         * there is at least one Vendor Info for the given vid (vendor ID)
   * Negative:
      * GET_VENDOR_INFO command failed **⇒** does not gets a Vendor Info for the given vid (vendor ID)
         * there is not one Vendor Info for the given vid (vendor ID)

* Role (Who can send)
   * Positive:
      * Trustee
      * Vendor 
      * VendorAdmin 
      * CertificationCenter 
      * NodeAdmin 

* Parameters:
   * vid (Vendor ID) - uint16:
      * Positive:
         * value exists
         * value > 0
         * integer value format
      * Negative 
         * empty value
         * value =< 0
         * string value format
         * length > MAX (MAX = 65535)
         * nonexistent ID

#### REST API 
GET: `/dcl/vendorinfo/vendors/{vid}`

* REST API command send
   * Positive:
      * correct HTTP method
      * request is authorized
      * uses valid credentials/role
   * Negative:
      * incorrect request
      * server side error

* Сommand result
   * Positive:
      * GET_VENDOR_INFO command completed successfully **⇒** gets a Vendor Info for the given vid (vendor ID)
         * there is at least one Vendor Info for the given vid (vendor ID)
   * Negative:
      * GET_VENDOR_INFO command failed **⇒** does not gets a Vendor Info for the given vid (vendor ID)
         * there is not one Vendor Info for the given vid (vendor ID)

* Role (Who can send)
   * Positive:
      * Trustee
      * Vendor 
      * VendorAdmin 
      * CertificationCenter 
      * NodeAdmin 

* Parameters:
   * vid (Vendor ID) - uint16:
      * Positive:
         * value exists
         * value > 0
         * integer value format
      * Negative 
         * empty value
         * value =< 0
         * string value format
         * length > MAX (MAX = 65535)
         * nonexistent ID

## [GET_ALL_VENDOR_INFO](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/docs/transactions/vendor-info.md#get_all_vendor_info)

### CLI command
CLI command: `dcld query vendorinfo all-vendors`

* CLI command send
   * Positive:
      * command exists/relevant
   * Negative:
      * incorrect command syntax
      * queries are performed through a Light Client Proxy

* Сommand result
   * Positive:
      * GET_ALL_VENDOR_INFO command completed successfully **⇒** gets information about all vendors for all VIDs
         * there is at least one Vendor for all VIDs
   * Negative:
      * GET_ALL_VENDOR_INFO command failed **⇒** does not information about all vendors for all VIDs
         * there is not one  Vendor for all VIDs

* Role (Who can send)
   * Positive:
      * Trustee
      * Vendor 
      * VendorAdmin 
      * CertificationCenter 
      * NodeAdmin 

* Parameters:
   
   * count-total - optional(bool)
      * Positive:
         * empty value
         * value state
            * TRUE (-1)
            * FALSE (0)
      * Negative 
         * value is not bool
   
   * limit - optional(uint)
      * Positive:	
         * value exists	
         * empty value
         * value =< 100	
      * Negative	
         * value > 100	
   
   * offset - optional(uint)
      * Positive:	
         * value exists	
         * empty value
         * value >= 0		
      * Negative	
         * value < 0	
   
   * page - optional(uint)
      * Positive:	
         * value exists	
         * empty value
         * value >= 0	
      * Negative	
         * value < 0	
   
   * page-key - optional(string)
      * Positive:	
         * empty value	
         * value exists	
      * Negative	
         * length < MIN
         * length > MAX
   
   * reverse - optional(bool)
      * Positive:	
         * empty value	
         * value state	
            * TRUE (-1)	
            * FALSE (0)	
      * Negative	
         * value is not bool	

### REST API 
GET: `/dcl/vendorinfo/vendors`

* REST API command send
   * Positive:
      * correct HTTP method
      * request is authorized
      * uses valid credentials/role
   * Negative:
      * incorrect request
      * queries are performed through a Light Client Proxy
      * server side error

* Сommand result
   * Positive:
      * GET_ALL_VENDOR_INFO command completed successfully **⇒** gets information about all vendors for all VIDs
         * there is at least one Vendor for all VIDs
   * Negative:
      * GET_ALL_VENDOR_INFO command failed **⇒** does not information about all vendors for all VIDs
         * there is not one  Vendor for all VIDs

* Role (Who can send)
   * Positive:
      * Trustee
      * Vendor 
      * VendorAdmin 
      * CertificationCenter 
      * NodeAdmin 

* Parameters:
   
   * count-total - optional(bool)
      * Positive:
         * empty value
         * value state
            * TRUE (-1)
            * FALSE (0)
      * Negative 
         * value is not bool
   
   * limit - optional(uint)
      * Positive:	
         * value exists	
         * empty value
         * value < 100	
      * Negative	
         * value > 100	
   
   * offset - optional(uint)
      * Positive:	
         * value exists	
         * empty value
         * value >= 0	
      * Negative	
         * value < 0	
   * page - optional(uint)
      * Positive:	
         * value exists	
         * empty value
         * value >= 0	
      * Negative	
         * value < 0	
   
   * page-key - optional(string)
      * Positive:	
         * empty value	
         * value exists	
      * Negative	
         * length < MIN
         * length > MAX	
   
   * reverse - optional(bool)
      * Positive:	
         * empty value	
         * value state	
            * TRUE (-1)	
            * FALSE (0)	
      * Negative	
         * value is not bool	