7. [Upgrade](#upgrade)
   * [PROPOSE_UPGRADE](#propose_upgrade)
   * [APPROVE_UPGRADE](#approve_upgrade)
   * [REJECT_UPGRADE](#reject_upgrade)
   * [GET_PROPOSED_UPGRADE](#get_proposed_upgrade)
   * [GET_APPROVED_UPGRADE](#get_approved_upgrade)
   * [GET_REJECTED_UPGRADE](#get_rejected_upgrade)
   * [GET_ALL_PROPOSED_UPGRADES](#get_all_proposed_upgrades)
   * [GET_ALL_APPROVED_UPGRADES](#get_all_approved_upgrades)
   * [GET_ALL_REJECTED_UPGRADES](#get_all_rejected_upgrades)
   * [GET_UPGRADE_PLAN](#get_upgrade_plan)
   * [GET_APPLIED_UPGRADE](#get_applied_upgrade)
   * [GET_MODULE_VERSIONS](#get_module_versions)

### PROPOSE_UPGRADE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
PROPOSE_UPGRADE command completed successfully	proposes an upgrade plan with the given name at the given height
number of  approvals greater than 2/3 of Trustees	upgrade approvals
upgrade proposal with the same name	
current upgrade proposal is out of date (when the current network height is greater than the proposed upgrade height)	
PROPOSE_UPGRADE command failed	does not proposes an upgrade plan with the given name at the given height
number of  approvals equals 2/3 of Trustees	upgrade not approvals
number of  approvals less than 2/3 of Trustees	upgrade not approvals
upgrade proposal with the same name	
current upgrade proposal is not out of date 	
Role (Who can send)	
Trustee	
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
name (Name)	string 
     * Positive:	
value exists	
current upgrade proposal is out of date	
current upgrade proposal is не out of date	error 
text value format
MIN < length < MAX	
     * Negative:	
empty value	
length > MAX	
upgrade-height (Upgrade Height)	int64
     * Positive:	
value > 0	
integer value format
     * Negative:	
empty value	
value =< 0
string value format	
length > MAX	MAX = 9 223 372 036 854 775 807
upgrade-info (Upgrade Info)	optional(string)
     * Positive:	
empty value	
text value format
value format	os/architecture
URL format	each URL include the corresponding checksum as checksum query parameter with the value in the format type:value
MIN < length < MAX	
     * Negative:	
length > MAX	
сontains spaces or line breaks	
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
PROPOSE_UPGRADE command completed successfully	proposes an upgrade plan with the given name at the given height
number of  approvals greater than 2/3 of Trustees	upgrade approvals
upgrade proposal with the same name	
current upgrade proposal is out of date (when the current network height is greater than the proposed upgrade height)	
PROPOSE_UPGRADE command failed	does not proposes an upgrade plan with the given name at the given height
number of  approvals equals 2/3 of Trustees	upgrade not approvals
number of  approvals less than 2/3 of Trustees	upgrade not approvals
upgrade proposal with the same name	
current upgrade proposal is not out of date 	
Role (Who can send)	
Trustee	
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
name (Name)	string 
     * Positive:	
value exists	
current upgrade proposal is out of date	
current upgrade proposal is не out of date	error 
text value format
MIN < length < MAX	
     * Negative:	
empty value	
length > MAX	
upgrade-height (Upgrade Height)	int64
     * Positive:	
value > 0	
integer value format
     * Negative:	
empty value	
value =< 0
string value format	
length > MAX	MAX = 9 223 372 036 854 775 807
upgrade-info (Upgrade Info)	optional(string)
     * Positive:	
empty value	
text value format
value format	os/architecture
URL format	each URL include the corresponding checksum as checksum query parameter with the value in the format type:value
MIN < length < MAX	
     * Negative:	
length > MAX	
сontains spaces or line breaks	
### APPROVE_UPGRADE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
APPROVE_UPGRADE command completed successfully	aproves the proposed upgrade plan with the given name
PROPOSE_UPGRADE command completed successfully	
number of  approvals greater than 2/3 of Trustees	upgrade approvals
APPROVE_UPGRADE command failed	does not aproves the proposed upgrade plan with the given name
PROPOSE_UPGRADEcommand failed	
number of  approvals equals 2/3 of Trustees	upgrade not approvals
number of  approvals less than 2/3 of Trustees	upgrade not approvals
Role (Who can send)	
Trustee	
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
name (Name)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
length > MAX	
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
APPROVE_UPGRADE command completed successfully	aproves the proposed upgrade plan with the given name
PROPOSE_UPGRADE command completed successfully	
number of  approvals greater than 2/3 of Trustees	upgrade approvals
APPROVE_UPGRADE command failed	does not aproves the proposed upgrade plan with the given name
PROPOSE_UPGRADEcommand failed	
number of  approvals equals 2/3 of Trustees	upgrade not approvals
number of  approvals less than 2/3 of Trustees	upgrade not approvals
Role (Who can send)	
Trustee	
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
name (Name)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
length > MAX	
### REJECT_UPGRADE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
REJECT_UPGRADE command completed successfully	rejects the proposed upgrade plan with the given name
PROPOSE_UPGRADE command completed successfully	
remove the proposal	
proposed upgrade has only proposer's approval and no rejects	
number of rejects more than 1/3 of Trustees	upgrade rejects 
REJECT_UPGRADE command failed	does not rejects the proposed upgrade plan with the given name
PROPOSE_UPGRADE command failed	
remove the proposal	
proposed upgrade has not proposer's approval	
proposed upgrade has only proposer's approval and rejects	
number of rejects equals 1/3 of Trustees	upgrade not rejects 
number of rejects less than 1/3 of Trustees	upgrade not rejects 
Role (Who can send)	
Trustee	
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
name (Name)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
length > MAX	
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
REJECT_UPGRADE command completed successfully	rejects the proposed upgrade plan with the given name
PROPOSE_UPGRADE command completed successfully	
remove the proposal	
proposed upgrade has only proposer's approval and no rejects	
number of rejects more than 1/3 of Trustees	upgrade rejects 
REJECT_UPGRADE command failed	does not rejects the proposed upgrade plan with the given name
PROPOSE_UPGRADE command failed	
remove the proposal	
proposed upgrade has not proposer's approval	
proposed upgrade has only proposer's approval and rejects	
number of rejects equals 1/3 of Trustees	upgrade not rejects 
number of rejects less than 1/3 of Trustees	upgrade not rejects 
Role (Who can send)	
Trustee	
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
name (Name)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
length > MAX	
### GET_PROPOSED_UPGRADE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_PROPOSED_UPGRADE command completed successfully	gets the proposed upgrade plan with the given name
PROPOSE_UPGRADE command completed successfully	
there is at least one proposed upgrade plan	
GET_PROPOSED_UPGRADE command failed	does not gets the proposed upgrade plan with the given name
PROPOSE_UPGRADE command failed	
there is not a one proposed upgrade plan	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
name (Name)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
length > MAX	
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
GET_PROPOSED_UPGRADE command completed successfully	gets the proposed upgrade plan with the given name
PROPOSE_UPGRADE command completed successfully	
there is at least one proposed upgrade plan	
GET_PROPOSED_UPGRADE command failed	does not gets the proposed upgrade plan with the given name
PROPOSE_UPGRADE command failed	
there is not a one proposed upgrade plan	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
name (Name)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
length > MAX	
### GET_APPROVED_UPGRADE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_APPROVED_UPGRADE command completed successfully	gets the approved upgrade plan with the given name
PROPOSE_UPGRADE command completed successfully	
APPROVED_UPGRADE command completed successfully	
there is at least one approved upgrade plan	
GET_APPROVED_UPGRADE command failed	does not gets the approved upgrade plan with the given name
PROPOSE_UPGRADE command failed	
APPROVED_UPGRADE command failed	
there is not a one approved upgrade plan	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
name (Name)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
length > MAX	
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
GET_APPROVED_UPGRADE command completed successfully	gets the approved upgrade plan with the given name
PROPOSE_UPGRADE command completed successfully	
APPROVED_UPGRADE command completed successfully	
there is at least one approved upgrade plan	
GET_APPROVED_UPGRADE command failed	does not gets the approved upgrade plan with the given name
PROPOSE_UPGRADE command failed	
APPROVED_UPGRADE command failed	
there is not a one approved upgrade plan	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
name (Name)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
length > MAX	
### GET_REJECTED_UPGRADE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_REJECTED_UPGRADE command completed successfully	gets the rejected upgrade plan with the given name
PROPOSE_UPGRADE command completed successfully	
REJECTED_UPGRADE command completed successfully	
there is at least one rejected upgrade plan	
GET_REJECTED_UPGRADE command failed	does not gets the rejected upgrade plan with the given name
PROPOSE_UPGRADE command failed	
REJECTED_UPGRADE command failed	
there is not a one rejected upgrade plan	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
name (Name)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
length > MAX	
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
GET_REJECTED_UPGRADE command completed successfully	gets the rejected upgrade plan with the given name
PROPOSE_UPGRADE command completed successfully	
REJECTED_UPGRADE command completed successfully	
there is at least one rejected upgrade plan	
GET_REJECTED_UPGRADE command failed	does not gets the rejected upgrade plan with the given name
PROPOSE_UPGRADE command failed	
REJECTED_UPGRADE command failed	
there is not a one rejected upgrade plan	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
name (Name)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
length > MAX	
### GET_ALL_PROPOSED_UPGRADES	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_PROPOSED_UPGRADES command completed successfully	gets all the proposed upgrade plans
PROPOSE_UPGRADE command completed successfully	
there is at least one proposed upgrade plans	
GET_ALL_PROPOSED_UPGRADES command failed	does not gets all the proposed upgrade plans
PROPOSE_UPGRADE command failed	
there is not a one proposed upgrade plans	
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
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ALL_PROPOSED_UPGRADES command completed successfully	gets all the proposed upgrade plans
PROPOSE_UPGRADE command completed successfully	
there is at least one proposed upgrade plans	
GET_ALL_PROPOSED_UPGRADES command failed	does not gets all the proposed upgrade plans
PROPOSE_UPGRADE command failed	
there is not a one proposed upgrade plans	
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
### GET_ALL_APPROVED_UPGRADES	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_APPROVED_UPGRADES command completed successfully	gets all the approved upgrade plans
PROPOSE_UPGRADE command completed successfully	
APPROVED_UPGRADE command completed successfully	
there is at least one approved upgrade plan	
GET_ALL_APPROVED_UPGRADES command failed	does not gets all the approved upgrade plans
PROPOSE_UPGRADE command failed	
APPROVED_UPGRADE command failed	
there is not a one approved upgrade plan	
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
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ALL_APPROVED_UPGRADES command completed successfully	gets all the approved upgrade plans
PROPOSE_UPGRADE command completed successfully	
APPROVED_UPGRADE command completed successfully	
there is at least one approved upgrade plan	
GET_ALL_APPROVED_UPGRADES command failed	does not gets all the approved upgrade plans
PROPOSE_UPGRADE command failed	
APPROVED_UPGRADE command failed	
there is not a one approved upgrade plan	
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
### GET_ALL_REJECTED_UPGRADES	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_REJECTED_UPGRADES command completed successfully	gets all the rejected upgrade plans
PROPOSE_UPGRADE command completed successfully	
REJECTED_UPGRADE command completed successfully	
there is at least one rejected upgrade plan	
GET_ALL_REJECTED_UPGRADES command failed	does not gets all the rejected upgrade plans
PROPOSE_UPGRADE command failed	
REJECTED_UPGRADE command failed	
there is not a one rejected upgrade plan	
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
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ALL_REJECTED_UPGRADES command completed successfully	gets all the rejected upgrade plans
PROPOSE_UPGRADE command completed successfully	
REJECTED_UPGRADE command completed successfully	
there is at least one rejected upgrade plan	
GET_ALL_REJECTED_UPGRADES command failed	does not gets all the rejected upgrade plans
PROPOSE_UPGRADE command failed	
REJECTED_UPGRADE command failed	
there is not a one rejected upgrade plan	
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
### GET_UPGRADE_PLAN	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_UPGRADE_PLAN command completed successfully	gets the currently scheduled upgrade plan, if it exists
currently scheduled upgrade plan exists	
GET_UPGRADE_PLAN command failed	does not gets the currently scheduled upgrade plan, if it exists
currently scheduled upgrade plan is not exists	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
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
GET_UPGRADE_PLAN command completed successfully	gets the currently scheduled upgrade plan, if it exists
currently scheduled upgrade plan exists	
GET_UPGRADE_PLAN command failed	does not gets the currently scheduled upgrade plan, if it exists
currently scheduled upgrade plan is not exists	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
### GET_APPLIED_UPGRADE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_APPLIED_UPGRADE command completed successfully	returns the header for the block at which the upgrade with the given name was applied, if it was previously executed on the chain
in the header for the block the upgrade with the given name was applied	
upgrade with the given name was previously executed on the chain	
GET_APPLIED_UPGRADE command failed	does not gets returns the header for the block at which the upgrade with the given name was applied, if it was previously executed on the chain
the upgrade with the given name was not applied in the header for the block	
upgrade with the given name was not previously executed on the chain	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
upgrade name (Upgrade Name)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
length > MAX	
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
GET_APPLIED_UPGRADE command completed successfully	returns the header for the block at which the upgrade with the given name was applied, if it was previously executed on the chain
in the header for the block the upgrade with the given name was applied	
upgrade with the given name was previously executed on the chain	
GET_APPLIED_UPGRADE command failed	does not gets returns the header for the block at which the upgrade with the given name was applied, if it was previously executed on the chain
the upgrade with the given name was not applied in the header for the block	
upgrade with the given name was not previously executed on the chain	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
upgrade name (Upgrade Name)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
length > MAX	
### GET_MODULE_VERSIONS	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_MODULE_VERSIONS command completed successfully	gets a list of module names and their respective consensus versions.
module version is specified	
a specific module is specified	return only that module's information
specific module name not specified	returns list of module names
GET_MODULE_VERSIONS command failed	does not gets a list of module names and their respective consensus versions.
module version is not specified	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
module name (Module Name)	optional(string)
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
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_MODULE_VERSIONS command completed successfully	gets a list of module names and their respective consensus versions.
module version is specified	
a specific module is specified	return only that module's information
specific module name not specified	returns list of module names
GET_MODULE_VERSIONS command failed	does not gets a list of module names and their respective consensus versions.
module version is not specified	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
module name (Module Name)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX