6. [Validator Node](#validator-node)
   * [ADD_VALIDATOR_NODE](#add_validator_node)
   * [DISABLE_VALIDATOR_NODE](#disable_validator_node)
   * [PROPOSE_DISABLE_VALIDATOR_NODE](#propose_disable_validator_node)
   * [APPROVE_DISABLE_VALIDATOR_NODE](#approve_disable_validator_node)
   * [REJECT_DISABLE_VALIDATOR_NODE](#reject_disable_validator_node)
   * [ENABLE_VALIDATOR_NODE](#enable_validator_node)
   * [GET_VALIDATOR](#get_validator)
   * [GET_ALL_VALIDATORS](#get_all_validators)
   * [GET_PROPOSED_DISABLE_VALIDATOR](#get_proposed_disable_validator)
   * [GET_ALL_PROPOSED_DISABLE_VALIDATORS](#get_all_proposed_disable_validators)
   * [GET_REJECTED_DISABLE_VALIDATOR](#get_rejected_disable_validator)
   * [GET_ALL_REJECTED_DISABLE_VALIDATORS](#get_all_rejected_disable_validators)
   * [GET_DISABLED_VALIDATOR](#get_disabled_validator)
   * [GET_ALL_DISABLED_VALIDATORS](#get_all_disabled_validators)
   * [GET_LAST_VALIDATOR_POWER](#get_last_validator_power)
   * [GET_ALL_LAST_VALIDATORS_POWER](#get_all_last_validators_power)

   ### ADD_VALIDATOR_NODE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
ADD_VALIDATOR_NODE command completed successfully	adds a new Validator node
ADD_VALIDATOR_NODE command failed	does not adds a new Validator node
Role (Who can send)	
Trustee	error
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	
Parameters:	
pubkey (Public Key)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
empty value	
moniker (Moniker)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
empty value	
identity (Identity)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
website (Website)	optional(string)
     * Positive:	
value exists	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
сontains spaces or line breaks	
details (Details)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
ip (IP)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
node-id (Node ID)	optional(string)
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
ADD_VALIDATOR_NODE command completed successfully	adds a new Validator node
ADD_VALIDATOR_NODE command failed	does not adds a new Validator node
Role (Who can send)	
Trustee	error
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	
Parameters:	
pubkey (Public Key)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
empty value	
moniker (Moniker)	string 
     * Positive:	
value exists	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
empty value	
identity (Identity)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
website (Website)	optional(string)
     * Positive:	
value exists	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
сontains spaces or line breaks	
details (Details)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
ip (IP)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
node-id (Node ID)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
### DISABLE_VALIDATOR_NODE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
DISABLE_VALIDATOR_NODE command completed successfully	disables the Validator node (removes from the validator set)
ADD_VALIDATOR_NODE command completed successfully	
there is at least one Validator node	
DISABLE_VALIDATOR_NODE command failed	does not disables the Validator node (removes from the validator set)
ADD_VALIDATOR_NODE command failed	
there is not a one Validator node	
Role (Who can send)	
Trustee	error
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	
owner	
not owner	
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
DISABLE_VALIDATOR_NODE command completed successfully	disables the Validator node (removes from the validator set)
ADD_VALIDATOR_NODE command completed successfully	
there is at least one Validator node	
DISABLE_VALIDATOR_NODE command failed	does not disables the Validator node (removes from the validator set)
ADD_VALIDATOR_NODE command failed	
there is not a one Validator node	
Role (Who can send)	
Trustee	error
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	
owner	
not owner	
### PROPOSE_DISABLE_VALIDATOR_NODE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
PROPOSE_DISABLE_VALIDATOR_NODE command completed successfully	proposes disabling of the Validator node from the validator set
ADD_VALIDATOR_NODE command completed successfully	
there is at least one Validator node	
sufficient number of approvals is received	disable confirmed
insufficient number of approvals is received	disable in a pending state
PROPOSE_DISABLE_VALIDATOR_NODE command failed	does not proposes disabling of the Validator node from the validator set
ADD_VALIDATOR_NODE command failed	
there is not a one Validator node	
approvals is not received	disable rejected
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
info (Information/Notes)	optional(string)
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
PROPOSE_DISABLE_VALIDATOR_NODE command completed successfully	proposes disabling of the Validator node from the validator set
ADD_VALIDATOR_NODE command completed successfully	
there is at least one Validator node	
sufficient number of approvals is received	disable confirmed
insufficient number of approvals is received	disable in a pending state
PROPOSE_DISABLE_VALIDATOR_NODE command failed	does not proposes disabling of the Validator node from the validator set
ADD_VALIDATOR_NODE command failed	
there is not a one Validator node	
approvals is not received	disable rejected
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
info (Information/Notes)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
### APPROVE_DISABLE_VALIDATOR_NODE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
APPROVE_DISABLE_VALIDATOR_NODE command completed successfully	approves disabling of the Validator node by a Trustee
ADD_VALIDATOR_NODE command completed successfully	
PROPOSE_DISABLE_VALIDATOR_NODE command completed successfully	
sufficient number of Trustees approve	validator node is disabled
APPROVE_DISABLE_VALIDATOR_NODE command failed	does not approves disabling of the Validator node by a Trustee
ADD_VALIDATOR_NODE command failed	
PROPOSE_DISABLE_VALIDATOR_NODE command failed	
insufficient number of Trustees approve	validator node is not disabled
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
info (Information/Notes)	optional(string)
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
APPROVE_DISABLE_VALIDATOR_NODE command completed successfully	approves disabling of the Validator node by a Trustee
ADD_VALIDATOR_NODE command completed successfully	
PROPOSE_DISABLE_VALIDATOR_NODE command completed successfully	
sufficient number of Trustees approve	validator node is disabled
APPROVE_DISABLE_VALIDATOR_NODE command failed	does not approves disabling of the Validator node by a Trustee
ADD_VALIDATOR_NODE command failed	
PROPOSE_DISABLE_VALIDATOR_NODE command failed	
insufficient number of Trustees approve	validator node is not disabled
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
info (Information/Notes)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
### REJECT_DISABLE_VALIDATOR_NODE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
REJECT_DISABLE_VALIDATOR_NODE command completed successfully	rejects disabling of the Validator node by a Trustee
ADD_VALIDATOR_NODE command completed successfully	
PROPOSE_DISABLE_VALIDATOR_NODE command completed successfully	
remove the proposal	
disable validator proposal has only proposer's approval and no rejects	
number of rejects more than 1/3 of Trustees	validator node is reject
REJECT_DISABLE_VALIDATOR_NODE command failed	does not approves disabling of the Validator node by a Trustee
ADD_VALIDATOR_NODE command failed	
PROPOSE_DISABLE_VALIDATOR_NODE command failed	
remove the proposal	
certificate has not proposer's approval	
certificate has only proposer's approval and rejects	
number of rejects less than 1/3 of Trustees	validator node is not reject
number of rejects equals 1/3 of Trustees	validator node is not reject
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
info (Information/Notes)	optional(string)
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
REJECT_DISABLE_VALIDATOR_NODE command completed successfully	rejects disabling of the Validator node by a Trustee
ADD_VALIDATOR_NODE command completed successfully	
PROPOSE_DISABLE_VALIDATOR_NODE command completed successfully	
remove the proposal	
disable validator proposal has only proposer's approval and no rejects	
number of rejects more than 1/3 of Trustees	validator node is reject
REJECT_DISABLE_VALIDATOR_NODE command failed	does not approves disabling of the Validator node by a Trustee
ADD_VALIDATOR_NODE command failed	
PROPOSE_DISABLE_VALIDATOR_NODE command failed	
remove the proposal	
certificate has not proposer's approval	
certificate has only proposer's approval and rejects	
number of rejects less than 1/3 of Trustees	validator node is not reject
number of rejects equals 1/3 of Trustees	validator node is not reject
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
info (Information/Notes)	optional(string)
     * Positive:	
empty value	
text value format
MIN < length < MAX	
     * Negative:	
length > MAX	
### ENABLE_VALIDATOR_NODE	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
ENABLE_VALIDATOR_NODE command completed successfully	enables the Validator node (returns to the validator set)
ADD_VALIDATOR_NODE command completed successfully	
DISABLE_VALIDATOR_NODE command completed successfully	
ENABLE_VALIDATOR_NODE command failed	does not enables the Validator node (returns to the validator set)ustee
ADD_VALIDATOR_NODE command failed	
DISABLE_VALIDATOR_NODE command failed	
Role (Who can send)	
Trustee	error
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	
owner	
not owner	error
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
ENABLE_VALIDATOR_NODE command completed successfully	enables the Validator node (returns to the validator set)
ADD_VALIDATOR_NODE command completed successfully	
DISABLE_VALIDATOR_NODE command completed successfully	
ENABLE_VALIDATOR_NODE command failed	does not enables the Validator node (returns to the validator set)ustee
ADD_VALIDATOR_NODE command failed	
DISABLE_VALIDATOR_NODE command failed	
Role (Who can send)	
Trustee	error
Vendor 	error
VendorAdmin 	error
CertificationCenter 	error
NodeAdmin 	
owner	
not owner	error
### GET_VALIDATOR	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_VALIDATOR command completed successfully	gets a validator node
ADD_VALIDATOR_NODE command completed successfully	
there is at least one validator node	
GET_VALIDATOR command failed	does not gets a validator node
ADD_VALIDATOR_NODE command failed	
there is not a one validator node	
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
GET_VALIDATOR command completed successfully	gets a validator node
ADD_VALIDATOR_NODE command completed successfully	
there is at least one validator node	
GET_VALIDATOR command failed	does not gets a validator node
ADD_VALIDATOR_NODE command failed	
there is not a one validator node	
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
### GET_ALL_VALIDATORS	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_VALIDATORS command completed successfully	gets the list of all validator nodes from the store
ADD_VALIDATOR_NODE command completed successfully	
there is at least one validator node	
there are only active stored validator nodes	
there are only jailed stored validator nodes	
there are both active and jailed stored validator nodes	all stored validator nodes (active and jailed) will be returned by default
GET_ALL_VALIDATORS command failed	does not gets the list of all validator nodes from the store
ADD_VALIDATOR_NODE command failed	
there is not a one validator node	
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
GET_ALL_VALIDATORS command completed successfully	gets the list of all validator nodes from the store
ADD_VALIDATOR_NODE command completed successfully	
there is at least one validator node	
there are only active stored validator nodes	
there are only jailed stored validator nodes	
there are both active and jailed stored validator nodes	all stored validator nodes (active and jailed) will be returned by default
GET_ALL_VALIDATORS command failed	does not gets the list of all validator nodes from the store
ADD_VALIDATOR_NODE command failed	
there is not a one validator node	
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
### GET_PROPOSED_DISABLE_VALIDATOR	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_PROPOSED_DISABLE_VALIDATOR command completed successfully	gets a proposed validator node
ADD_VALIDATOR_NODE command completed successfully	
PROPOSED_DISABLE_VALIDATOR command completed successfully	
there is at least one proposed validator node	
GET_PROPOSED_DISABLE_VALIDATOR command failed	does not gets a proposed validator node
ADD_VALIDATOR_NODE command failed	
PROPOSED_DISABLE_VALIDATOR command failed	
there is not a one proposed validator node	
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
GET_PROPOSED_DISABLE_VALIDATOR command completed successfully	gets a proposed validator node
ADD_VALIDATOR_NODE command completed successfully	
PROPOSED_DISABLE_VALIDATOR command completed successfully	
there is at least one proposed validator node	
GET_PROPOSED_DISABLE_VALIDATOR command failed	does not gets a proposed validator node
ADD_VALIDATOR_NODE command failed	
PROPOSED_DISABLE_VALIDATOR command failed	
there is not a one proposed validator node	
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
### GET_ALL_PROPOSED_DISABLE_VALIDATORS	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_PROPOSED_DISABLE_VALIDATORS command completed successfully	gets a the list of all proposed disable validator nodes from the store
ADD_VALIDATOR_NODE command completed successfully	
PROPOSED_DISABLE_VALIDATOR command completed successfully	
there is at least one proposed validator node	
GET_ALL_PROPOSED_DISABLE_VALIDATORS command failed	does not gets a the list of all proposed disable validator nodes from the store
ADD_VALIDATOR_NODE command failed	
PROPOSED_DISABLE_VALIDATOR command failed	
there is not a one proposed validator node	
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
GET_ALL_PROPOSED_DISABLE_VALIDATORS command completed successfully	gets a the list of all proposed disable validator nodes from the store
ADD_VALIDATOR_NODE command completed successfully	
PROPOSED_DISABLE_VALIDATOR command completed successfully	
there is at least one proposed validator node	
GET_ALL_PROPOSED_DISABLE_VALIDATORS command failed	does not gets a the list of all proposed disable validator nodes from the store
ADD_VALIDATOR_NODE command failed	
PROPOSED_DISABLE_VALIDATOR command failed	
there is not a one proposed validator node	
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
### GET_REJECTED_DISABLE_VALIDATOR	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_REJECTED_DISABLE_VALIDATOR command completed successfully	gets a rejected validator node
ADD_VALIDATOR_NODE command completed successfully	
PROPOSED_DISABLE_VALIDATOR command completed successfully	
REJECTED_DISABLE_VALIDATOR command completed successfully	
there is at least one rejected validator node	
GET_REJECTED_DISABLE_VALIDATOR command failed	does not gets a rejected validator node
ADD_VALIDATOR_NODE command failed	
PROPOSED_DISABLE_VALIDATOR command failed	
REJECTED_DISABLE_VALIDATOR command failed	
there are not one rejected validator nodes	
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
GET_REJECTED_DISABLE_VALIDATOR command completed successfully	gets a rejected validator node
ADD_VALIDATOR_NODE command completed successfully	
PROPOSED_DISABLE_VALIDATOR command completed successfully	
REJECTED_DISABLE_VALIDATOR command completed successfully	
there is at least one rejected validator node	
GET_REJECTED_DISABLE_VALIDATOR command failed	does not gets a rejected validator node
ADD_VALIDATOR_NODE command failed	
PROPOSED_DISABLE_VALIDATOR command failed	
REJECTED_DISABLE_VALIDATOR command failed	
there are not one rejected validator nodes	
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
### GET_ALL_REJECTED_DISABLE_VALIDATORS	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_REJECTED_DISABLE_VALIDATOR command completed successfully	gets the list of all rejected disable validator nodes from the store
ADD_VALIDATOR_NODE command completed successfully	
PROPOSED_DISABLE_VALIDATOR command completed successfully	
REJECTED_DISABLE_VALIDATOR command completed successfully	
there is at least one rejected validator node	
GET_REJECTED_DISABLE_VALIDATOR command failed	does not gets the list of all rejected disable validator nodes from the store
ADD_VALIDATOR_NODE command failed	
PROPOSED_DISABLE_VALIDATOR command failed	
REJECTED_DISABLE_VALIDATOR command failed	
there are not one rejected validator nodes	
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
GET_REJECTED_DISABLE_VALIDATOR command completed successfully	gets the list of all rejected disable validator nodes from the store
ADD_VALIDATOR_NODE command completed successfully	
PROPOSED_DISABLE_VALIDATOR command completed successfully	
REJECTED_DISABLE_VALIDATOR command completed successfully	
there is at least one rejected validator node	
GET_REJECTED_DISABLE_VALIDATOR command failed	does not gets the list of all rejected disable validator nodes from the store
ADD_VALIDATOR_NODE command failed	
PROPOSED_DISABLE_VALIDATOR command failed	
REJECTED_DISABLE_VALIDATOR command failed	
there are not one rejected validator nodes	
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
### GET_DISABLED_VALIDATOR	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_DISABLED_VALIDATOR command completed successfully	gets a disabled validator node
ADD_VALIDATOR_NODE command completed successfully	
PROPOSED_DISABLE_VALIDATOR command completed successfully	
APPROVE_DISABLE_VALIDATOR_NODE command completed successfully	
there is at least one disabled validator node	
GET_REJECTED_DISABLE_VALIDATOR command failed	does not gets a disabled validator node
ADD_VALIDATOR_NODE command failed	
PROPOSED_DISABLE_VALIDATOR command failed	
APPROVE_DISABLE_VALIDATOR_NODE command failed	
there are not one disabled validator node	
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
GET_DISABLED_VALIDATOR command completed successfully	gets a disabled validator node
ADD_VALIDATOR_NODE command completed successfully	
PROPOSED_DISABLE_VALIDATOR command completed successfully	
APPROVE_DISABLE_VALIDATOR_NODE command completed successfully	
there is at least one disabled validator node	
GET_REJECTED_DISABLE_VALIDATOR command failed	does not gets a disabled validator node
ADD_VALIDATOR_NODE command failed	
PROPOSED_DISABLE_VALIDATOR command failed	
APPROVE_DISABLE_VALIDATOR_NODE command failed	
there are not one disabled validator node	
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
### GET_ALL_DISABLED_VALIDATORS	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_DISABLED_VALIDATORS command completed successfully	gets a the list of all disabled validator nodes from the store
ADD_VALIDATOR_NODE command completed successfully	
PROPOSED_DISABLE_VALIDATOR command completed successfully	
APPROVE_DISABLE_VALIDATOR_NODE command completed successfully	
there is at least one disabled validator node	
GET_ALL_DISABLED_VALIDATORS command failed	does not gets a the list of all disabled validator nodes from the store
ADD_VALIDATOR_NODE command failed	
PROPOSED_DISABLE_VALIDATOR command failed	
APPROVE_DISABLE_VALIDATOR_NODE command failed	
there are not one disabled validator node	
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
GET_ALL_DISABLED_VALIDATORS command completed successfully	gets a the list of all disabled validator nodes from the store
ADD_VALIDATOR_NODE command completed successfully	
PROPOSED_DISABLE_VALIDATOR command completed successfully	
APPROVE_DISABLE_VALIDATOR_NODE command completed successfully	
there is at least one disabled validator node	
GET_ALL_DISABLED_VALIDATORS command failed	does not gets a the list of all disabled validator nodes from the store
ADD_VALIDATOR_NODE command failed	
PROPOSED_DISABLE_VALIDATOR command failed	
APPROVE_DISABLE_VALIDATOR_NODE command failed	
there are not one disabled validator node	
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
### GET_LAST_VALIDATOR_POWER	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_LAST_VALIDATOR_POWER command completed successfully	gets a last validator node power
there is at least one validator node	
GET_LAST_VALIDATOR_POWER command failed	does not gets a last validator node power
there are not one validator node	
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
GET_LAST_VALIDATOR_POWER command completed successfully	gets a last validator node power
there is at least one validator node	
GET_LAST_VALIDATOR_POWER command failed	does not gets a last validator node power
there are not one validator node	
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
### GET_ALL_LAST_VALIDATORS_POWER	
#### CLI command	
CLI command send	
Valid command	
command exists/relevant	
Invalid command	
access is denied to execute command	
incorrect command syntax	
Сommand result	
GET_ALL_LAST_VALIDATORS_POWER command completed successfully	gets the list of all last validator nodes power from the store
there is at least one validator node	
GET_ALL_LAST_VALIDATORS_POWER command failed	does not gets the list of all last validator nodes power from the store
there are not one validator node	
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
GET_ALL_LAST_VALIDATORS_POWER command completed successfully	gets the list of all last validator nodes power from the store
there is at least one validator node	
GET_ALL_LAST_VALIDATORS_POWER command failed	does not gets the list of all last validator nodes power from the store
there are not one validator node	
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