set -euo pipefail

DOCKERFILE="${1:-Dockerfile-build}"
IMAGE_TAG="${2:-dcld-build}"
VERSION_DIR="${3:-genesis}"
MAINNET_STABLE_VERSION="${4:-""}"
LOCALNET_DIR=".localnet"

docker build -f ${DOCKERFILE} -t ${IMAGE_TAG} .
docker container create --name ${IMAGE_TAG}-inst ${IMAGE_TAG}

if [ -n "$MAINNET_STABLE_VERSION" ]; then
    wget "https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/download/${MAINNET_STABLE_VERSION}/dcld"
    chmod ugo+x dcld 
fi

for node_name in node0 node1 node2 node3 observer0 lightclient0; do
    if [[ -d "${LOCALNET_DIR}/${node_name}" ]]; then
        mkdir -p ${LOCALNET_DIR}/${node_name}/cosmovisor/${VERSION_DIR}/bin/
        if [ -n "$MAINNET_STABLE_VERSION" ]; then
            cp dcld ${LOCALNET_DIR}/${node_name}/cosmovisor/${VERSION_DIR}/bin/
        else
            docker cp ${IMAGE_TAG}-inst:/go/bin/dcld ${LOCALNET_DIR}/${node_name}/cosmovisor/${VERSION_DIR}/bin/
        fi
    fi
done

if [ -n "$MAINNET_STABLE_VERSION" ]; then
    rm dcld
fi

docker rm ${IMAGE_TAG}-inst
