package cli

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const CliBinaryName = "dcld"

// common constants
var node_p2p_port int = 26670
var node_client_port int = 26671
var chain_id string = "dclchain"
var node0conn string = "tcp://192.167.10.2:26657"
var docker_network string = "distributed-compliance-ledger_localnet"
var Passphrase string = "test1234"
var LOCALNET_DIR string = ".localnet"

// RED := `tput setaf 1`
// GREEN := `tput setaf 2`
// RESET := `tput sgr0`
var RED = ""
var GREEN = ""
var RESET = ""
var DEF_OUTPUT_MODE string = "json"

func Random_string(arg1 string, _length ...int) string {
	length := _length[1]

	if length < 1 {
		length = 6
	}

	__resultvar, _ := Execute("eval", arg1)

	if _, err := Execute("command", "-v", "shasum", "&>", "/dev/null"); err != nil {
		__resultvar, _ = Execute("bash", "-c", "'date +%s.%N | shasum | fold -w", strconv.Itoa(length), " | head -n 1")

	} else {
		__resultvar, _ = Execute("bash", "-c", "'date +%s.%N | sha1sum | fold -w", strconv.Itoa(length), " | head -n 1")

	}
	return string(__resultvar)
}

func _check_response(arg1 string, temp_mode ...string) bool {
	_result := arg1
	_expected_string := temp_mode[1]
	var _mode string

	if temp_mode[2] == "" {
		_mode = DEF_OUTPUT_MODE
	}

	if _mode == "json" {
		if _, err := Execute("-n", "$(echo", _result, " | jq | grep", _expected_string, "2>/dev/null)"); err != nil {
			fmt.Printf("true")
			return true
		}
	} else {
		if _, err := Execute("-n", "$(echo", _result+" | grep "+_expected_string+" 2>/dev/null)"); err != nil {
			fmt.Printf("true")
			return true
		}
	}
	fmt.Print("false")
	return false
}

func Check_response(arg1 string, temp_mode ...string) bool {
	_result := arg1[1]
	_expected_string := temp_mode[1]
	var _mode string

	if temp_mode[2] == "" {
		_mode = DEF_OUTPUT_MODE
	}

	if err := _check_response(string(_result), string(_expected_string), _mode); err != false {
		fmt.Printf(GREEN, "ERROR:", RESET, "command failed. The expected string:", string(_expected_string), " not found in the result:", string(_result))
		return true
	}
	return false
}

func Test_divider() {
	fmt.Printf("")
	fmt.Printf("--------------------------")
	fmt.Printf("")
}

func AddKeys(user string) (string, error) {

	resp, err := Execute("keys", "add", user)

	return resp, err
}

func ShowKeys(user string, arg string) (string, error) {

	resp, err := Execute("keys", "show", user, arg)

	return resp, err
}

func GetUserAddress(user string, passphrase string) (string, error) {
	keys, _ := ShowKeys(user, "-a")

	args := "echo" + passphrase + " | " + keys

	resp, err := Execute("bash", "-c", args)

	return resp, err
}

func GetUserPubKey(user string, passphrase string) (string, error) {
	keys, _ := ShowKeys(user, "-p")

	args := "echo" + passphrase + " | " + keys

	resp, err := Execute("bash", "-c", args)

	return resp, err
}

func Random() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r1.Intn(32767)
}

func Check_response_and_report(arg1, arg2, temp_mode string) {
	_result := arg1
	_expected_string := arg2
	var _mode string

	if temp_mode == "" {
		_mode = DEF_OUTPUT_MODE
	}

	Check_response(_result, _expected_string, _mode)
	fmt.Printf(GREEN + "SUCCESS: " + RESET + "Result contains expected substring: " + _expected_string)
}

func Create_new_vendor_account(arg1, arg2 string) {

	_name := arg1
	_vid := arg2

	Execute("bash", "-c", "echo ", Passphrase, "| dcld keys add ", _name)
	_address, _ := Execute("bash", "-c", "echo ", Passphrase, " | dcld keys show ", _name, "-a")
	_pubkey, _ := Execute("bash", "-c", "echo ", Passphrase, "| dcld keys show ", _name, "-p")

	fmt.Printf("Jack proposes account for \"" + _name + "\" with Vendor role")
	_result, _ := Execute("bash", "-c", "echo ", Passphrase, "| dcld tx auth propose-add-account --address=", _address, " --pubkey=", _pubkey, " --roles=Vendor --vid=", _vid, "--from jack --yes")
	Check_response(_result, "\"code\": 0")
}

func Create_model_and_version(arg1, arg2, arg3, arg4, arg5 string) {

	_vid := arg1
	_pid := arg2
	_softwareVersion := arg3
	_softwareVersionString := arg4
	_user_address := arg5
	result, _ := Execute("bash", "-c", "echo ", Passphrase, " | dcld tx model add-model --vid=", _vid, "--pid=", _pid, "--deviceTypeID=1 --productName=TestProduct --productLabel=TestingProductLabel --partNumber=1 --commissioningCustomFlow=0 --from=", _user_address, "--yes")
	Check_response(result, "\"code\": ", "0")
	result, _ = Execute("bashe", "-c", "echo ", Passphrase, " | dcld tx model add-model-version --cdVersionNumber=1 --maxApplicableSoftwareVersion=10 --minApplicableSoftwareVersion=1 --vid=", _vid, "--pid=", _pid, "--softwareVersion=", _softwareVersion, "--softwareVersionString=", _softwareVersionString, "--from=", _user_address, "--yes")
	Check_response(result, "\"code\": ", "0")
}

func Create_new_account(arg1, arg2 string) string {

	__resultvar := arg1
	var name string
	name = Random_string(name)
	__resultvar = __resultvar + name

	roles := arg2

	fmt.Printf("Account name: " + name)

	fmt.Printf("Generate key for " + name)
	Execute("bash", "-c", "echo ", Passphrase, "; echo ", Passphrase, ") | dcld keys add ", name)

	address, _ := Execute("bash", "-c", "echo ", Passphrase, "| dcld keys show ", name, "-a")
	pubkey, _ := Execute("bash", "-c", "echo ", Passphrase, "| dcld keys show ", name, "-p")

	fmt.Printf("Jack proposes account for \"" + name + "\" with roles: \"" + roles + "\"")
	result, _ := Execute("bash", "-c", "echo ", Passphrase, "| dcld tx auth propose-add-account --address=", address, " --pubkey=", pubkey, " --roles=", roles, "--from jack --yes")
	Check_response(result, "\"code\": ", "0")
	fmt.Printf(result)

	fmt.Printf("Alice approves account for \"" + name + "\" with roles: \"" + roles + "\"")
	result, _ = Execute("bash", "-c", "echo ", Passphrase, "| dcld tx auth approve-add-account --address=", address, " --from alice --yes")
	Check_response(result, "\"code\": ", "0")
	fmt.Printf(result)

	return result
}

func Response_does_not_contain(arg1, arg2, temp_mode string) bool {
	_result := arg1
	_unexpected_string := arg2
	var _mode string = temp_mode

	if temp_mode == "" {
		_mode = DEF_OUTPUT_MODE
	}

	if _check_response(_result, _unexpected_string, _mode) == true {
		fmt.Printf("ERROR: command failed. The unexpected string: "+_unexpected_string, "found in the result: "+_result)
		return true
	}

	fmt.Printf(GREEN + "SUCCESS: " + RESET + "Result does not contain unexpected substring: " + _unexpected_string)

	return false
}
