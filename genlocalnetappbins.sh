set -euo pipefail

DOCKERFILE="${1:-Dockerfile-build}"
IMAGE_TAG="${2:-dcld-build}"
VERSION_DIR="${3:-genesis}"
MAINNET_STABLE_VERSION="${4:-""}"
LOCALNET_DIR=".localnet"

if env | grep GOCOVER=1; then
    docker build --build-arg "GOCOVER=1" -f ${DOCKERFILE} -t ${IMAGE_TAG} .
else
    docker build --build-arg "GOCOVER=" -f ${DOCKERFILE} -t ${IMAGE_TAG} .
fi

docker container create --name ${IMAGE_TAG}-inst ${IMAGE_TAG}

for node_name in node0 node1 node2 node3 observer0 lightclient0; do
    if [[ -d "${LOCALNET_DIR}/${node_name}" ]]; then
        mkdir -p ${LOCALNET_DIR}/${node_name}/cosmovisor/${VERSION_DIR}/bin/
        mkdir -p ${LOCALNET_DIR}/${node_name}/gocover
        if [ -n "$MAINNET_STABLE_VERSION" ]; then
            echo "cp ${MAINNET_STABLE_VERSION} ${LOCALNET_DIR}/${node_name}/cosmovisor/${VERSION_DIR}/bin/"
            cp ${MAINNET_STABLE_VERSION} ${LOCALNET_DIR}/${node_name}/cosmovisor/${VERSION_DIR}/bin/dcld
        else
            docker cp ${IMAGE_TAG}-inst:/go/bin/dcld ${LOCALNET_DIR}/${node_name}/cosmovisor/${VERSION_DIR}/bin/
        fi
    fi
done

docker rm ${IMAGE_TAG}-inst
