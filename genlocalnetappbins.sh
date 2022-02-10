set -euo pipefail

LOCALNET_DIR=".localnet"

docker build -f Dockerfile-build -t dcld-build .
docker container create --name dcld-build-inst dcld-build

for node_name in node0 node1 node2 node3 observer0 lightclient0; do
    if [[ -d "${LOCALNET_DIR}/${node_name}" ]]; then
        mkdir -p ${LOCALNET_DIR}/${node_name}/cosmovisor/genesis/bin/
        docker cp dcld-build-inst:/go/bin/dcld ${LOCALNET_DIR}/${node_name}/cosmovisor/genesis/bin/
    fi
done

docker rm dcld-build-inst
