# Fork

CosmosSDK is forked at the version v0.37.4 in order to implement multiproofs.

For development:
* Fork zigbee-alliance/cosmos-sdk
* Checkout `multiproofs` branch
* Replace `github.com/cosmos/cosmos-sdk => github.com/zigbee-alliance/cosmos-sdk multiproofs`
to `github.com/cosmos/cosmos-sdk => local/path/to/the/repo` in `go.mod`

# Proofs

How proofs in Tendermint / Cosmos SDK works.

* You call `cliCtx.QueryStore(key, storeName)`
* On client side CosmosSDK sends `abci.RequestQuery` where `Path=/store/[storeName]/key`, `Data=binaryEncodedKey`. See:
`client/context/query.go:query:91`.
* On server side CosmosSDK routes request depending on route prefix. See: `baseapp/baseapp.go:Query:452`.
* If prefix is `/store` cosmos casts rootmultistore to `Queriable`, cuts `/store` prefix from the `req.Path` and sends
`abci.RequestQuery` directly to the rootmultistore.
* Root store cuts part with the store name and routes query to the store. At this moment `req.Path=/key`,
`req.Data=encoded key`

* Depending on route (see `store/iavl/store.go`):
    * `/key` - store responds with abci.ResponseQuery where `res.Key=encoded key`, `res.Value=encoded value`,
    `res.Proof=iavl value operator | iavl absence operator`. In operator `key=encoded key`. Tendermint defines
    proof operators and operator chain validation algorithm (see `tendermint/crypto/merkle/proof_test.go`
    to understand). `tendermint/iavl` package provides operators for `proof_iavl_absence.go` and `proof_iavl_value.go`.
    * `/range` (new) - encoded `RangeReq` is expected as `req.Key`. This value is returned in `res.Key`. `Resp.Value` is
    encoded `RangeRes` structure that holds response keys and values. `ProofOpIAVLRange` (new) is used as proof operator.
    `res.Key` is also used in `proofOperator.Key` and `keyPath` during validation.
    * `/subspace` - not interesting.

* Multistore adds proof operator to the chain.
* On client side Cosmos sdk validates the response. See: `client/context/query.go:verifyProof:145`. Key path is
constructed as `[storeName]/[url encoded resp.Key bytes]` `res.Proof` contains proofOperator chain. Value is taken from
`res.Value`. In range proof imlementation res.

We store `query` in `res.Key` and `(keys, values)` in `res.Value` for the following reasons:
* Key path (see `ProofRuntime.Verify` and `crypto/merkle/proof_key_path.go`) doesn't support multikeys
* It's impossible to validate response without query since we use "start with this key and take this amount" queries, not "get those keys".
* `res.Key` and `res.Value` are `byte[]`, not `byte[][]`. See: `ResponseQuery`
