# Run Light Client Proxy

Light Client Proxy can be used if there are no trusted Full Nodes (Validator or Observers) a client can connect to.

It can be a proxy for CLI or direct requests from code done via CometBFT RPC.

Please note, that CLI can use a Light Client proxy only for single-value query requests.
A Full Node (Validator or Observer) should be used for multi-value query requests and write requests.

See the following links for details about a Light Client:

- <https://docs.cometbft.com/v0.37/core/light-client>

## Running Light Client Proxy - Short

```bash
dcld light <chain-id> -p tcp://<primary-host>:26657 -w tcp://<witness1-host>:26657,tcp://<witness2-host>:26657
dcld config chain-id <chain-id>
dcld config node tcp://<light-client-proxy-host>:8888
```

## Running Light Client Proxy - Detailed

See <https://docs.cometbft.com/v0.37/core/light-client> for details

### 1. Choose Semi-trusted or Non-trusted Nodes for Connection

Light Client must be connected to one Primary Full Node (Validator or Observer) and
a number of Witness Full Nodes (Validators or Observers).

### 2. Start a Light Client Proxy Server

```bash
dcld light <chain-id> -p tcp://<primary-host>:26657 -w tcp://<witness1-host>:26657,tcp://<witness2-host>:26657
```

Light Client Proxy is started at port `8888` by default. It can be changed with `--laddr` argument.

Light Client Proxy will write some information (about headers) to disk. Default location is `~/.cometbft-light`.
It can be changed with `--dir` argument.

Light Client needs initial height and hash it trusts. If no `--height` and `--hash` arguments provided,
it will try to calculate it automatically. Basically it will do the steps from Item 4.

If you want to specify the trusted height and hash manually, you can use  `--height` and `--hash` arguments.

```bash
dcld light <chain-id> -p tcp://<primary-host>:26657 -w tcp://<witness1-host>:26657,tcp://<witness2-host>:26657 --height <height> --hash <hash>
```

Please look at Item 4 for a possible way how to obtain them.

### 3. Connect CLI to a Light Client Proxy Server

```bash
dcld config chain-id <chain-id>
dcld config node tcp://<light-client-proxy-host>:8888
```

### 4. How to Obtain Trusted Height & Hash

This is needed only if you want to specify the trusted height and hash manually.

One way to obtain a semi-trusted hash & height is to query multiple full nodes and compare their hashes.

For the first node (for example Primary Node):

```bash
curl -s http(s)://<node host>:26657/commit | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```

Result Example:

```json
{
  "height": "<height>",
  "hash": "<hash>"
}
```

For other nodes (for example Witness Nodes):

```bash
curl -s http(s)://<node host>:26657/commit?height=<height> | jq "{height: .result.signed_header.header.height, hash: .result.signed_header.commit.block_id.hash}"
```

Where `<height>` obtained from the first run.

If result is the same on all nodes, then obtained `<height>` and `<hash>` can be used as parameters for the light client at the next step.
