5. [Auth](#auth)
   * [PROPOSE_ADD_ACCOUNT](#propose_add_account)
   * [APPROVE_ADD_ACCOUNT](#aprove_add_account)
   * [REJECT_ADD_ACCOUNT](#regect_add_account)
   * [PROPOSE_REVOKE_ACCOUNT](#propose_revoke_account)
   * [APPROVE_REVOKE_ACCOUNT](#approve_revoke_account)
   * [GET_ACCOUNT](#get_account)
   * [GET_PROPOSED_ACCOUNT](#get_proposed_account)
   * [GET_REJECTED_ACCOUNT](#get_rejected_account)
   * [GET_PROPOSED_ACCOUNT_TO_REVOKE](#get_proposed_account_to_revoke)
   * [GET_REVOKED_ACCOUNT](#get_revoked_account)
   * [GET_ALL_ACCOUNTS](#get_all_accounts)
   * [GET_ALL_PROPOSED_ACCOUNTS](#get_all_proposed_accounts)
   * [GET_ALL_REJECTED_ACCOUNTS](#get_all_rejected_accounts)
   * [GET_ALL_PROPOSED_ACCOUNTS_TO_REVOKE](#get_all_proposed_accounts_to_revoke)
   * [GET_ALL_REVOKED_ACCOUNTS](#get_all_revoked_accounts)

 ### PROPOSE_ADD_ACCOUNT	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
PROPOSE_ADD_ACCOUNT command completed successfully	proposes a new Account with the given address, public key and role
sufficient number of approvals is received	account added
insufficient number of approvals is received	account in a pending state
PROPOSE_ADD_ACCOUNT command failed	does not proposes a new Account with the given address, public key and role
approvals is not received	account not added
Role (Who can send)	
Trustee	
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
address (Address)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
pub_key (Public Key)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
vid (Vendor ID )	optional(uint16)
     * Positive:	
unique combination	
value > 0	
integer value format
nonexistent ID	
     * Negative:	
empty value
value =< 0
string value format		
length > MAX	MAX = 65535
pid_ranges (Product ID Ranges)	optional(array<uint16 range>)
     * Positive:	
unique combination	
value > 0	
integer value format
nonexistent ID	
the data stored in the array matches the expected format	
the data stored in the array does not exceed the permissible limits	
the list is displayed in ascending order	
     * Negative:	
empty value
value =< 0
string value format		
length > MAX	MAX = 65535
roles (Roles)	array<string>
     * Positive:	
support value	
Vendor	
TestHouse	
CertificationCenter	
Trustee	
NodeAdmin	
VendorAdmin	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
empty value	
all array elements are of different types	
info (information/notes)	optional(string)
     * Positive:	
value exists	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
time (proposal time)	optional(int64)
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
Invalid command	
incorrect request	
server side error	
Сommand result	
PROPOSE_ADD_ACCOUNT command completed successfully	proposes a new Account with the given address, public key and role
sufficient number of approvals is received	account added
insufficient number of approvals is received	account in a pending state
PROPOSE_ADD_ACCOUNT command failed	does not proposes a new Account with the given address, public key and role
approvals is not received	account not added
Role (Who can send)	
Trustee	
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
address (Address)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
pub_key (Public Key)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
vid (Vendor ID )	optional(uint16)
     * Positive:	
unique combination	
value > 0	
integer value format
nonexistent ID	
     * Negative:	
empty value	
value =< 0
string value format	
length > MAX	MAX = 65535
pid_ranges (Product ID Ranges)	optional(array<uint16 range>)
     * Positive:	
unique combination	
value > 0	
integer value format
nonexistent ID	
the data stored in the array matches the expected format	
the data stored in the array does not exceed the permissible limits	
the list is displayed in ascending order	
     * Negative:	
empty value	
value =< 0
string value format	
length > MAX	MAX = 65535
roles (Roles)	array<string>
     * Positive:	
support value	
Vendor	
TestHouse	
CertificationCenter	
Trustee	
NodeAdmin	
VendorAdmin	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
empty value	
all array elements are of different types	
info (information/notes)	optional(string)
     * Positive:	
value exists	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
time (proposal time)	optional(int64)
     * Positive:	
default value	current time by default
empty value	
integer value format
     * Negative:	
length > MAX	MAX = 9 223 372 036 854 775 807
### APPROVE_ADD_ACCOUNT	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
APPROVE_ADD_ACCOUNT command completed successfully	approves the proposed account
PROPOSE_ADD_ACCOUNT command completed successfully	
Number of required approvals greater than 2/3 of Trustees for account roles: TestHouse, CertificationCenter, Trustee, NodeAdmin, VendorAdmin	account is active
Number of required approvals equal 2/3 of Trustees for account roles: TestHouse, CertificationCenter, Trustee, NodeAdmin, VendorAdmin	account is active
Number of required approvals greater than 1/3 of Trustees for account role: Vendor	account is active
APPROVE_ADD_ACCOUNT command failed	does not approves the proposed account
PROPOSE_ADD_ACCOUNT command failed	
Number of required approvals less than 2/3 of Trustees for account roles: TestHouse, CertificationCenter, Trustee, NodeAdmin, VendorAdmin	account is not active
Number of required approvals less than 1/3 of Trustees for account role: Vendor	account is not active
Number of required approvals equal 1/3 of Trustees for account role: Vendor	account is not active
Role (Who can send)	
Trustee	
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
address (Address)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
info (information/notes)	optional(string)
     * Positive:	
value exists	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
time (proposal time)	optional(int64)
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
Invalid command	
incorrect request	
server side error	
Сommand result	
APPROVE_ADD_ACCOUNT command completed successfully	approves the proposed account
PROPOSE_ADD_ACCOUNT command completed successfully	
Number of required approvals greater than 2/3 of Trustees for account roles: TestHouse, CertificationCenter, Trustee, NodeAdmin, VendorAdmin	account is active
Number of required approvals equal 2/3 of Trustees for account roles: TestHouse, CertificationCenter, Trustee, NodeAdmin, VendorAdmin	account is active
Number of required approvals greater than 1/3 of Trustees for account role: Vendor	account is active
APPROVE_ADD_ACCOUNT command failed	does not approves the proposed account
PROPOSE_ADD_ACCOUNT command failed	
Number of required approvals less than 2/3 of Trustees for account roles: TestHouse, CertificationCenter, Trustee, NodeAdmin, VendorAdmin	account is not active
Number of required approvals less than 1/3 of Trustees for account role: Vendor	account is not active
Number of required approvals equal 1/3 of Trustees for account role: Vendor	account is not active
Role (Who can send)	
Trustee	
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
address (Address)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
info (information/notes)	optional(string)
     * Positive:	
value exists	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
time (proposal time)	optional(int64)
     * Positive:	
default value	current time by default
empty value	
integer value format
     * Negative:	
length > MAX	MAX = 9 223 372 036 854 775 807
### REJECT_ADD_ACCOUNT	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
REJECT_ADD_ACCOUNT command completed successfully	rejects the proposed account
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_ADD_ACCOUNT command completed successfully	
remove the proposal	
account has only proposer's approval and no rejects	
Number of required rejects greater than 1/3 of Trustees for account roles: TestHouse, CertificationCenter, Trustee, NodeAdmin, VendorAdmin	account is reject
Number of required rejects greater than 2/3 of Trustees for account role: Vendor	account is reject
Number of required rejects equal 2/3 of Trustees for account role: Vendor	account is reject
REJECT_ADD_ACCOUNT command failed	does not rejects the proposed account
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_ADD_ACCOUNT command failed	
remove the proposal	
account has not only proposer's approval and no rejects	
account has only proposer's approval and rejects	
account has not only proposer's approval and rejects	
Number of required rejects equal 1/3 of Trustees for account roles: TestHouse, CertificationCenter, Trustee, NodeAdmin, VendorAdmin	account is not reject
Number of required rejects less than 1/3 of Trustees for account roles: TestHouse, CertificationCenter, Trustee, NodeAdmin, VendorAdmin	account is not reject
Number of required rejects less than 2/3 of Trustees for account role: Vendor	account is not reject
Role (Who can send)	
Trustee	
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
address (Address)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
info (information/notes)	optional(string)
     * Positive:	
value exists	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
time (proposal time)	optional(int64)
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
Invalid command	
incorrect request	
server side error	
Сommand result	
REJECT_ADD_ACCOUNT command completed successfully	rejects the proposed account
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_ADD_ACCOUNT command completed successfully	
remove the proposal	
account has only proposer's approval and no rejects	
Number of required rejects greater than 1/3 of Trustees for account roles: TestHouse, CertificationCenter, Trustee, NodeAdmin, VendorAdmin	account is reject
Number of required rejects greater than 2/3 of Trustees for account role: Vendor	account is reject
Number of required rejects equal 2/3 of Trustees for account role: Vendor	account is reject
REJECT_ADD_ACCOUNT command failed	does not rejects the proposed account
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_ADD_ACCOUNT command failed	
remove the proposal	
account has not only proposer's approval and no rejects	
account has only proposer's approval and rejects	
account has not only proposer's approval and rejects	
Number of required rejects equal 1/3 of Trustees for account roles: TestHouse, CertificationCenter, Trustee, NodeAdmin, VendorAdmin	account is not reject
Number of required rejects less than 1/3 of Trustees for account roles: TestHouse, CertificationCenter, Trustee, NodeAdmin, VendorAdmin	account is not reject
Number of required rejects less than 2/3 of Trustees for account role: Vendor	account is not reject
Role (Who can send)	
Trustee	
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
address (Address)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
info (information/notes)	optional(string)
     * Positive:	
value exists	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
time (proposal time)	optional(int64)
     * Positive:	
default value	current time by default
empty value	
integer value format
     * Negative:	
length > MAX	MAX = 9 223 372 036 854 775 807
### PROPOSE_REVOKE_ACCOUNT	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
PROPOSE_REVOKE_ACCOUNT command completed successfully	proposes revocation of the Account with the given address
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_ADD_ACCOUNT command completed successfully	
sufficient number of approvals is received	account is revoke 
insufficient number of approvals is received	revocation in a pending state
PROPOSE_REVOKE_ACCOUNT command failed	does not proposes revocation of the Account with the given address
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_ADD_ACCOUNT command failed	
approvals is not received	revocation rejected
Role (Who can send)	
Trustee	
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
address (Address)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
info (information/notes)	optional(string)
     * Positive:	
value exists	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
time (proposal time)	optional(int64)
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
Invalid command	
incorrect request	
server side error	
Сommand result	
PROPOSE_REVOKE_ACCOUNT command completed successfully	proposes revocation of the Account with the given address
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_ADD_ACCOUNT command completed successfully	
sufficient number of approvals is received	account is revoke 
insufficient number of approvals is received	revocation in a pending state
PROPOSE_REVOKE_ACCOUNT command failed	does not proposes revocation of the Account with the given address
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_ADD_ACCOUNT command failed	
approvals is not received	revocation rejected
Role (Who can send)	
Trustee	
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
address (Address)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
info (information/notes)	optional(string)
     * Positive:	
value exists	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
time (proposal time)	optional(int64)
     * Positive:	
default value	current time by default
empty value	
integer value format
     * Negative:	
length > MAX	MAX = 9 223 372 036 854 775 807
### APPROVE_REVOKE_ACCOUNT	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
APPROVE_REVOKE_ACCOUNT command completed successfully	approves the proposed revocation of the account
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_ADD_ACCOUNT command completed successfully	
Number of required approvals greater than 2/3 of Trustees	account is revoked
Number of required approvals equal 2/3 of Trustees	account is revoked
APPROVE_REVOKE_ACCOUNT command failed	does not approves the proposed revocation of the account
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_ADD_ACCOUNT command failed	
Number of required approvals less than 2/3 of Trustees	account is not revoked
Role (Who can send)	
Trustee	
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
address (Address)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
info (information/notes)	optional(string)
     * Positive:	
value exists	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
time (proposal time)	optional(int64)
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
Invalid command	
incorrect request	
server side error	
Сommand result	
APPROVE_REVOKE_ACCOUNT command completed successfully	approves the proposed revocation of the account
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_ADD_ACCOUNT command completed successfully	
Number of required approvals greater than 2/3 of Trustees	account is revoked
Number of required approvals equal 2/3 of Trustees	account is revoked
APPROVE_REVOKE_ACCOUNT command failed	does not approves the proposed revocation of the account
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_ADD_ACCOUNT command failed	
Number of required approvals less than 2/3 of Trustees	account is not revoked
Role (Who can send)	
Trustee	
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	error
Parameters:	
address (Address)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
info (information/notes)	optional(string)
     * Positive:	
value exists	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
time (proposal time)	optional(int64)
     * Positive:	
default value	current time by default
empty value	
integer value format
     * Negative:	
length > MAX	MAX = 9 223 372 036 854 775 807
### GET_ACCOUNT	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ACCOUNT command completed successfully	gets an accounts by the address
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_ADD_ACCOUNT command completed successfully	
there is at least one account	
GET_ACCOUNT command failed	does not gets an accounts by the address
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_ADD_ACCOUNT command failed	
there are not one accounts	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
address (Address)	string 
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
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_ACCOUNT command completed successfully	gets an accounts by the address
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_ADD_ACCOUNT command completed successfully	
there is at least one account	
GET_ACCOUNT command failed	does not gets an accounts by the address
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_ADD_ACCOUNT command failed	
there are not one accounts	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
address (Address)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
### GET_PROPOSED_ACCOUNT	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
valid command result	
accounts by the address	
proposed but not approved accounts	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_PROPOSED_ACCOUNT command completed successfully	gets a proposed but not approved accounts by its address
PROPOSE_ADD_ACCOUNT command completed successfully	
there is at least one proposed but not approved accounts	
GET_PROPOSED_ACCOUNT command failed	does not gets a proposed but not approved accounts by its address
PROPOSE_ADD_ACCOUNT command failed	
there are not one proposed but not approved accounts	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
address (Address)	string 
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
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_PROPOSED_ACCOUNT command completed successfully	gets a proposed but not approved accounts by its address
PROPOSE_ADD_ACCOUNT command completed successfully	
there is at least one proposed but not approved accounts	
GET_PROPOSED_ACCOUNT command failed	does not gets a proposed but not approved accounts by its address
PROPOSE_ADD_ACCOUNT command failed	
there are not one proposed but not approved accounts	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
address (Address)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
### GET_REJECTED_ACCOUNT	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_REJECTED_ACCOUNT command completed successfully	gets a rejected accounts by its address
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_ADD_ACCOUNT command completed successfully, but account rejected 	
there is at least one rejected accounts	
GET_REJECTED_ACCOUNT command failed	does not gets a rejected accounts by its address
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_ADD_ACCOUNT command failed	
there are not one rejected accounts	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
address (Address)	string 
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
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_REJECTED_ACCOUNT command completed successfully	gets a rejected accounts by its address
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_ADD_ACCOUNT command completed successfully, but account rejected 	
there is at least one rejected accounts	
GET_REJECTED_ACCOUNT command failed	does not gets a rejected accounts by its address
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_ADD_ACCOUNT command failed	
there are not one rejected accounts	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
address (Address)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
### GET_PROPOSED_ACCOUNT_TO_REVOKE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_PROPOSED_ACCOUNT_TO_REVOKE command completed successfully	gets a proposed but not approved accounts to be revoked by its address
PROPOSE_ADD_ACCOUNT command completed successfully	
there is at least one proposed but not approved accounts	
GET_PROPOSED_ACCOUNT_TO_REVOKE command failed	does not gets a proposed but not approved accounts to be revoked by its address
PROPOSE_ADD_ACCOUNT command failed	
there are not one proposed but not approved accounts	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
address (Address)	string 
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
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_PROPOSED_ACCOUNT_TO_REVOKE command completed successfully	gets a proposed but not approved accounts to be revoked by its address
PROPOSE_ADD_ACCOUNT command completed successfully	
there is at least one proposed but not approved accounts	
GET_PROPOSED_ACCOUNT_TO_REVOKE command failed	does not gets a proposed but not approved accounts to be revoked by its address
PROPOSE_ADD_ACCOUNT command failed	
there are not one proposed but not approved accounts	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
address (Address)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
### GET_REVOKED_ACCOUNT	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_REVOKED_ACCOUNT command completed successfully	gets a revoked account by its address
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_REVOKE_ACCOUNT command completed successfully	
there is at least one revoked account	
GET_REVOKED_ACCOUNT command failed	does not gets a revoked account by its address
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_REVOKE_ACCOUNT  command failed	
there are not one revoked account	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
address (Address)	string 
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
Invalid command	
incorrect request	
server side error	
Сommand result	
GET_REVOKED_ACCOUNT command completed successfully	gets a revoked account by its address
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_REVOKE_ACCOUNT command completed successfully	
there is at least one revoked account	
GET_REVOKED_ACCOUNT command failed	does not gets a revoked account by its address
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_REVOKE_ACCOUNT  command failed	
there are not one revoked account	
Role (Who can send)	
Trustee	
Vendor 	
VendorAdmin 	
CertificationCenter 	
NodeAdmin 	
Parameters:	
address (Address)	string 
     * Positive:	
string matches the format	
text value format
MIN < length < MAX	
     * Negative:	
empty value	
nonexistent value	
length > MAX	
### GET_ALL_ACCOUNTS	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_ACCOUNTS command completed successfully	gets all accounts
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_REVOKE_ACCOUNT command completed successfully	
there is at least one account	
GET_ALL_ACCOUNTS command failed	does not gets all accounts
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_REVOKE_ACCOUNT  command failed	
there are not one account	
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
GET_ALL_ACCOUNTS command completed successfully	gets all accounts
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_REVOKE_ACCOUNT command completed successfully	
there is at least one account	
GET_ALL_ACCOUNTS command failed	does not gets all accounts
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_REVOKE_ACCOUNT  command failed	
there are not one account	
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
### GET_ALL_PROPOSED_ACCOUNTS	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_PROPOSED_ACCOUNTS command completed successfully	gets all proposed but not approved accounts
PROPOSE_ADD_ACCOUNT command completed successfully	
there is at least one proposed but not approved accounts	
GET_ALL_PROPOSED_ACCOUNTS command failed	does not gets all proposed but not approved accounts
PROPOSE_ADD_ACCOUNT command failed	
there are not one proposed but not approved accounts	
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
GET_ALL_PROPOSED_ACCOUNTS command completed successfully	gets all proposed but not approved accounts
PROPOSE_ADD_ACCOUNT command completed successfully	
there is at least one proposed but not approved accounts	
GET_ALL_PROPOSED_ACCOUNTS command failed	does not gets all proposed but not approved accounts
PROPOSE_ADD_ACCOUNT command failed	
there are not one proposed but not approved accounts	
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
### GET_ALL_REJECTED_ACCOUNTS	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_REJECTED_ACCOUNTS command completed successfully	gets all rejected accounts
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_ADD_ACCOUNT command completed successfully, but account rejected 	
there is at least one rejected accounts	
GET_ALL_REJECTED_ACCOUNTS command failed	does not getsall rejected accounts
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_ADD_ACCOUNT command failed	
there are not one rejected accounts	
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
GET_ALL_REJECTED_ACCOUNTS command completed successfully	gets all rejected accounts
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_ADD_ACCOUNT command completed successfully, but account rejected 	
there is at least one rejected accounts	
GET_ALL_REJECTED_ACCOUNTS command failed	does not getsall rejected accounts
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_ADD_ACCOUNT command failed	
there are not one rejected accounts	
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
### GET_ALL_PROPOSED_ACCOUNTS_TO_REVOKE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_PROPOSED_ACCOUNTS_TO_REVOKE command completed successfully	gets all proposed but not approved accounts to be revoked
PROPOSE_ADD_ACCOUNT command completed successfully	
there is at least one proposed but not approved accounts	
GET_ALL_PROPOSED_ACCOUNTS_TO_REVOKE command failed	does not gets all proposed but not approved accounts to be revoked
PROPOSE_ADD_ACCOUNT command failed	
there are not one proposed but not approved accounts	
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
GET_ALL_PROPOSED_ACCOUNTS_TO_REVOKE command completed successfully	gets all proposed but not approved accounts to be revoked
PROPOSE_ADD_ACCOUNT command completed successfully	
there is at least one proposed but not approved accounts	
GET_ALL_PROPOSED_ACCOUNTS_TO_REVOKE command failed	does not gets all proposed but not approved accounts to be revoked
PROPOSE_ADD_ACCOUNT command failed	
there are not one proposed but not approved accounts	
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
### GET_ALL_REVOKED_ACCOUNTS	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_REVOKED_ACCOUNTS command completed successfully	gets all revoked accounts
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_REVOKE_ACCOUNT command completed successfully	
there is at least one revoked account	
GET_ALL_REVOKED_ACCOUNTS command failed	does not gets all revoked accounts
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_REVOKE_ACCOUNT  command failed	
there are not one revoked account	
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
GET_ALL_REVOKED_ACCOUNTS command completed successfully	gets all revoked accounts
PROPOSE_ADD_ACCOUNT command completed successfully	
APPROVE_REVOKE_ACCOUNT command completed successfully	
there is at least one revoked account	
GET_ALL_REVOKED_ACCOUNTS command failed	does not gets all revoked accounts
PROPOSE_ADD_ACCOUNT command failed	
APPROVE_REVOKE_ACCOUNT  command failed	
there are not one revoked account	
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