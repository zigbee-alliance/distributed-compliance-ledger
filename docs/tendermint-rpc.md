# Tendermint RPC

[Tendermint RPC](https://docs.tendermint.com/v0.34/rpc/#) can be used to query application's state and/or [subscribe to Tendermint WebSocket events](https://docs.tendermint.com/v0.34/tendermint-core/subscription.html).   
The default RPC listen address is tcp://0.0.0.0:26657. See the [docs](https://docs.tendermint.com/v0.34/rpc/#) to learn how to configure and customize.

## WebSocket Events

Full list of Tendermint WebSocket events can be found [here](https://pkg.go.dev/github.com/tendermint/tendermint/types#pkg-constants).

### Subscribe

Subscribing to Tendermint WebSocket events can be done using any WS clients like [wscat](https://github.com/websockets/wscat), [Postman](https://www.postman.com/), etc. The following is an example of subscribing to a new `tx` (transaction) event using wscat:

1. Connect to the server:

```bash
wscat -c ws://<node-ip>:<port>/websocket
```

In cases of [Test Net](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/deployment/persistent_chains/testnet-2.0/testnet-2.0-csa-endpoints.md) and [Main Net](https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/deployment/persistent_chains/main-net/main-net-csa-endpoints.md) observer nodes, `node-ip`s are `on.test-net.dcl.csa-iot.org` and `on.dcl.csa-iot.org`, respectively, having the same ports `26657`

2. Subscribe to the `tx` event in the wscat command line (refer to the [docs](https://docs.tendermint.com/v0.34/rpc/#) to learn the syntax):

```json
{"jsonrpc":"2.0","method":"subscribe","id":0,"params":{"query":"tm.event='Tx'"}}
```

You should get the following response if everything is OK:

```json
{
  "jsonrpc": "2.0",
  "id": 0,
  "result": {}
}
```

3. Wait for the transaction

Consider performing any transaction in another terminal tab to broadcast a `tx` event for the test purposes. For example, the following transaction can be sent in case of a local pool (for Test Net or Main Net pools `dcld` needs to be configured properly and write access needs to be given, see https://github.com/zigbee-alliance/distributed-compliance-ledger/blob/master/docs/how-to.md#cli-configuration)

```bash
dcld tx auth propose-add-account \
--address=cosmos1sdg5vkpz9urcemmw67lnxvzhhuveqe6v70ex2l \
--pubkey="{\"@type\":\"/cosmos.crypto.secp256k1.PubKey\",\"key\":\"A4m98FT/tgsMWZuVBlavWmawwXrvv/nUMhDlU8QsHlOM\"}" \
--roles=NodeAdmin \
--from=jack
```

4. Confirm that the subscriber received the event. For the transaction example above:

```json
{
  "jsonrpc": "2.0",
  "id": 0,
  "result": {
    "query": "tm.event='Tx'",
    "data": {
      "type": "tendermint/event/Tx",
      "value": {
        "TxResult": {
          "height": "1684",
          "tx": "CocCCoQCCkgvemlnYmVlYWxsaWFuY2UuZGlzdHJpYnV0ZWRjb21wbGlhbmNlbGVkZ2VyLmRjbGF1dGguTXNnUHJvcG9zZUFkZEFjY291bnQStwEKLWNvc21vczE0cG1nMnhnMzd0enI2a3VxNHV4dnp3a3p1ZnRldGQ3dDhhMnN1dRItY29zbW9zMXNkZzV2a3B6OXVyY2VtbXc2N2xueHZ6aGh1dmVxZTZ2NzBleDJsGkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohA4m98FT/tgsMWZuVBlavWmawwXrvv/nUMhDlU8QsHlOMIglOb2RlQWRtaW44lo/YoAYSWApQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohApEHaI8+Wj64IQdR/WZDMaWHNd53+9G19zNcWjrAiXMbEgQKAggBGAISBBDAmgwaQJCg0T8phCo70H7VpL2M1C0kQix4LeVOjkXWgK3EHmFobXkBGLH2jJiKTHakVWUO2yj2Tdpc4YoTDw2eKe6UZj8=",
          "result": {
            "data": "CkoKSC96aWdiZWVhbGxpYW5jZS5kaXN0cmlidXRlZGNvbXBsaWFuY2VsZWRnZXIuZGNsYXV0aC5Nc2dQcm9wb3NlQWRkQWNjb3VudA==",
            "log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeAddAccount\"}]}]}]",
            "gas_used": "25760",
            "events": [
              {
                "type": "tx",
                "attributes": [
                  {
                    "key": "YWNjX3NlcQ==",
                    "value": "Y29zbW9zMTRwbWcyeGczN3R6cjZrdXE0dXh2endrenVmdGV0ZDd0OGEyc3V1LzI=",
                    "index": true
                  }
                ]
              },
              {
                "type": "tx",
                "attributes": [
                  {
                    "key": "c2lnbmF0dXJl",
                    "value": "a0tEUlB5bUVLanZRZnRXa3ZZelVMU1JDTEhndDVVNk9SZGFBcmNRZVlXaHRlUUVZc2ZhTW1JcE1kcVJWWlE3YktQWk4ybHpoaWhNUERaNHA3cFJtUHc9PQ==",
                    "index": true
                  }
                ]
              },
              {
                "type": "message",
                "attributes": [
                  {
                    "key": "YWN0aW9u",
                    "value": "L3ppZ2JlZWFsbGlhbmNlLmRpc3RyaWJ1dGVkY29tcGxpYW5jZWxlZGdlci5kY2xhdXRoLk1zZ1Byb3Bvc2VBZGRBY2NvdW50",
                    "index": true
                  }
                ]
              }
            ]
          }
        }
      }
    },
    "events": {
      "tm.event": [
        "Tx"
      ],
      "tx.hash": [
        "AF831BBE5BF23545447BB2D1B65E49F99B5A22035F7E1FEB3E4BDC8D243F4115"
      ],
      "tx.height": [
        "1684"
      ],
      "tx.acc_seq": [
        "cosmos14pmg2xg37tzr6kuq4uxvzwkzuftetd7t8a2suu/2"
      ],
      "tx.signature": [
        "kKDRPymEKjvQftWkvYzULSRCLHgt5U6ORdaArcQeYWhteQEYsfaMmIpMdqRVZQ7bKPZN2lzhihMPDZ4p7pRmPw=="
      ],
      "message.action": [
        "/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeAddAccount"
      ]
    }
  }
}
```

### Decoding The Transaction String 

The details of transaction payload encoding can be found [here](https://docs.cosmos.network/main/core/encoding#transaction-encoding). Thus, the transaction payload/string in the example above can be decoded as follows:

```bash
dcld tx decode CocCCoQCCkgvemlnYmVlYWxsaWFuY2UuZGlzdHJpYnV0ZWRjb21wbGlhbmNlbGVkZ2VyLmRjbGF1dGguTXNnUHJvcG9zZUFkZEFjY291bnQStwEKLWNvc21vczE0cG1nMnhnMzd0enI2a3VxNHV4dnp3a3p1ZnRldGQ3dDhhMnN1dRItY29zbW9zMXNkZzV2a3B6OXVyY2VtbXc2N2xueHZ6aGh1dmVxZTZ2NzBleDJsGkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohA4m98FT/tgsMWZuVBlavWmawwXrvv/nUMhDlU8QsHlOMIglOb2RlQWRtaW44lo/YoAYSWApQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohApEHaI8+Wj64IQdR/WZDMaWHNd53+9G19zNcWjrAiXMbEgQKAggBGAISBBDAmgwaQJCg0T8phCo70H7VpL2M1C0kQix4LeVOjkXWgK3EHmFobXkBGLH2jJiKTHakVWUO2yj2Tdpc4YoTDw2eKe6UZj8=
```

which gives the result:

```json
{"body":{"messages":[{"@type":"/zigbeealliance.distributedcomplianceledger.dclauth.MsgProposeAddAccount","signer":"cosmos14pmg2xg37tzr6kuq4uxvzwkzuftetd7t8a2suu","address":"cosmos1sdg5vkpz9urcemmw67lnxvzhhuveqe6v70ex2l","pubKey":{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"A4m98FT/tgsMWZuVBlavWmawwXrvv/nUMhDlU8QsHlOM"},"roles":["NodeAdmin"],"vendorID":0,"info":"","time":"1679165334"}],"memo":"","timeout_height":"0","extension_options":[],"non_critical_extension_options":[]},"auth_info":{"signer_infos":[{"public_key":{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"ApEHaI8+Wj64IQdR/WZDMaWHNd53+9G19zNcWjrAiXMb"},"mode_info":{"single":{"mode":"SIGN_MODE_DIRECT"}},"sequence":"2"}],"fee":{"amount":[],"gas_limit":"200000","payer":"","granter":""}},"signatures":["kKDRPymEKjvQftWkvYzULSRCLHgt5U6ORdaArcQeYWhteQEYsfaMmIpMdqRVZQ7bKPZN2lzhihMPDZ4p7pRmPw=="]}
```

Detailed info about subscribing can be found [here](https://docs.tendermint.com/v0.34/rpc/#/Websocket/subscribe).

### Unsubscribe

Unsubscribing is done using a similar command to [subscription](#subscribe):

```json
{"jsonrpc":"2.0","method":"unsubscribe","id":0,"params":{"query":"tm.event='Tx'"}}
```

## Querying Application Components

There are many endpoints available to query various types of information about the nodes. For example, a transaction can be fetched by its hash using any web client like curl as follows:

```bash
curl -X GET "http://<node-ip>:<port>/tx?hash=0x74504B24ED59A424E436656E5E9A11034C7A7C7ED3BE7C3CDEA1ED387EF62967&prove=true" -H "accept: application/json"
```

It returns a response as follows:

```json
{
  "jsonrpc": "2.0",
  "id": -1,
  "result": {
    "hash": "74504B24ED59A424E436656E5E9A11034C7A7C7ED3BE7C3CDEA1ED387EF62967",
    "height": "1692",
    "index": 0,
    "tx_result": {
      "code": 0,
      "data": "CkoKSC96aWdiZWVhbGxpYW5jZS5kaXN0cmlidXRlZGNvbXBsaWFuY2VsZWRnZXIuZGNsYXV0aC5Nc2dBcHByb3ZlQWRkQWNjb3VudA==",
      "log": "[{\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"/zigbeealliance.distributedcomplianceledger.dclauth.MsgApproveAddAccount\"}]}]}]",
      "info": "",
      "gas_wanted": "0",
      "gas_used": "32100",
      "events": [
        {
          "type": "tx",
          "attributes": [
            {
              "key": "YWNjX3NlcQ==",
              "value": "Y29zbW9zMXlnejM3am4wdmFkdmV2M3dyODduZ2RkdDl6cG1hcHd3ZzhwMDRrLzQ=",
              "index": true
            }
          ]
        },
        {
          "type": "tx",
          "attributes": [
            {
              "key": "c2lnbmF0dXJl",
              "value": "VHltY0YwbE9VNzJTUW1VRDg4ZEdmdldMOUxXcWYzRXgwYldjNEtLdExKSkdoV2hqdTJkbUs1L3BrNHFlQjcwL2wvSm80Mit3MlNoUHc4WUpGemR4M2c9PQ==",
              "index": true
            }
          ]
        },
        {
          "type": "message",
          "attributes": [
            {
              "key": "YWN0aW9u",
              "value": "L3ppZ2JlZWFsbGlhbmNlLmRpc3RyaWJ1dGVkY29tcGxpYW5jZWxlZGdlci5kY2xhdXRoLk1zZ0FwcHJvdmVBZGRBY2NvdW50",
              "index": true
            }
          ]
        }
      ],
      "codespace": ""
    },
    "tx": "CrMBCrABCkgvemlnYmVlYWxsaWFuY2UuZGlzdHJpYnV0ZWRjb21wbGlhbmNlbGVkZ2VyLmRjbGF1dGguTXNnQXBwcm92ZUFkZEFjY291bnQSZAotY29zbW9zMXlnejM3am4wdmFkdmV2M3dyODduZ2RkdDl6cG1hcHd3ZzhwMDRrEi1jb3Ntb3Mxc2RnNXZrcHo5dXJjZW1tdzY3bG54dnpoaHV2ZXFlNnY3MGV4MmwgwY/YoAYSWApQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAv+avtRfIGPCkbToqtJJUO8yX8NsAeW4MPkg3ADUEM0zEgQKAggBGAQSBBDAmgwaQE8pnBdJTlO9kkJlA/PHRn71i/S1qn9xMdG1nOCirSySRoVoY7tnZiuf6ZOKnge9P5fyaONvsNkoT8PGCRc3cd4=",
    "proof": {
      "root_hash": "BB905252C0221042D19C66A8215A7ABF53873C773975E56BEAF946C3B1DC2512",
      "data": "CrMBCrABCkgvemlnYmVlYWxsaWFuY2UuZGlzdHJpYnV0ZWRjb21wbGlhbmNlbGVkZ2VyLmRjbGF1dGguTXNnQXBwcm92ZUFkZEFjY291bnQSZAotY29zbW9zMXlnejM3am4wdmFkdmV2M3dyODduZ2RkdDl6cG1hcHd3ZzhwMDRrEi1jb3Ntb3Mxc2RnNXZrcHo5dXJjZW1tdzY3bG54dnpoaHV2ZXFlNnY3MGV4MmwgwY/YoAYSWApQCkYKHy9jb3Ntb3MuY3J5cHRvLnNlY3AyNTZrMS5QdWJLZXkSIwohAv+avtRfIGPCkbToqtJJUO8yX8NsAeW4MPkg3ADUEM0zEgQKAggBGAQSBBDAmgwaQE8pnBdJTlO9kkJlA/PHRn71i/S1qn9xMdG1nOCirSySRoVoY7tnZiuf6ZOKnge9P5fyaONvsNkoT8PGCRc3cd4=",
      "proof": {
        "total": "1",
        "index": "0",
        "leaf_hash": "u5BSUsAiEELRnGaoIVp6v1OHPHc5deVr6vlGw7HcJRI=",
        "aunts": [
          
        ]
      }
    }
  }
}
```

The `proof` property can be excluded from the response by setting the `prove` property in the query string to `false` (or just by omitting it).  
