#!/bin/bash
set -e

echo "Get key info for Jack"
zblcli keys show jack

echo "Get account info for Jack"
zblcli query account $(zblcli keys show jack -a)

echo "Assign Vendor role to Jack"
echo "test1234" | zblcli tx authz assign-role $(zblcli keys show jack -a) "vendor" --from jack --yes

sleep 5

vid=$RANDOM
pid=$RANDOM
echo "Add Model with VID: ${vid} PID: ${pid}"
echo "test1234" | zblcli tx modelinfo add-model $vid $pid "Device #1" "Device Description" "SKU12FS" "1.0" "2.0" true --from jack --yes

sleep 5

echo "Get Model with VID: ${vid} PID: ${pid}"
zblcli query modelinfo model $vid $pid

echo "Get all model infos"
zblcli query modelinfo all-models

echo "Get all vendors"
zblcli query modelinfo vendors

echo "Get Vendor Models with VID: ${vid}"
zblcli query modelinfo vendor-models $vid

echo "Update Model with VID: ${vid} PID: ${pid}"
echo "test1234" | zblcli tx modelinfo update-model $vid $pid "New Device Description" true --from jack --yes

sleep 5

echo "Get Model with VID: ${vid} PID: ${pid}"
zblcli query modelinfo model $vid $pid