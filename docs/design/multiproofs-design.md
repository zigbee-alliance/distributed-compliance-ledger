# Multiproofs design

## How queries in Cosmos work

- Client (CIL or rest) makes a `QueryStore` or `QueryWithData`  call to CosmosSDK.
- Cosmos encodes the request and sends it to a Tendermint node using ABCI
- On the node's side request is decoded by Cosmos
- If the request is of type `QueryWithData`:
  - Cosmos resolves route and sends it to the required handler (controller)
  - A handler can access storage using Cosmos's `store` API
  - The handler returns whatever it wants
- If the request is of type `QueryStore`:
  - Cosmos directly requests `store` for the proof
- Cosmos encodes the responce and returns it to Tendermint
- Tendermint returns the response to the client
- On the client's side, Cosmos decodes the response and returns it to the caller

## Options

`Store` is Cosmos's abstraction above data structures used to persist data (iavl, multisrore, ...). Those structures in their turn use Tendermint's `tm-db` abstraction above databases (leveldb, memdb, ...).

Currently, `iavl` implementation is used to store domain entities in `dc-ledger`. The implementation contains the method for getting range proof `GetRangeWithProof`. It doesn't contain the method for checking that the range proof contains sequential elements.

There are two options on how to bring multiproof to dc-ledger:

- Extend the `store` interface and Cosmos API.
  - The flow will be like in `QueryStore` use case.
  - \- It requires more changes in Cosmos.
- Extend the `store` interface only.
  - The flow will be like in `QueryWithData` use case.
  - Handlers will be able to access multiproof API via `store`.
  - \  It requires fewer changes in Cosmos.
  - \- Probably there will be code duplicates in the validation logic.

I recommend the second option because of the listed pros and cons. Also, we will be able to take the implementation and extend it to the first option if necessary.

The described solution will primarily work for paginated results (`get all` is a particular case). For filtered results, we will have to maintain dedicated indexes.

As a result of the proposed implementation, each response with a page of entities will also have multi proof that will prove that each of the entity is present in the ledger.

It WILL NOT prove offset. To prove it we will have to return hashes of all the entities preceding to the page.

To prove that returned entities are sequential we will have to implement an additional check. There is no such check in the IAVL implementation.

## Questions

- Do we need to prove total, offset and that all element form a sequential group?