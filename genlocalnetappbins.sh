set -euo pipefail

DOCKERFILE="${1:-Dockerfile-build}"
IMAGE_TAG="${2:-dcld-build}"
VERSION_DIR="${3:-genesis}"

LOCALNET_DIR=".localnet"

docker build -f ${DOCKERFILE} -t ${IMAGE_TAG} .
docker container create --name ${IMAGE_TAG}-inst ${IMAGE_TAG}

for node_name in node0 node1 node2 node3 observer0 lightclient0; do
    if [[ -d "${LOCALNET_DIR}/${node_name}" ]]; then
        mkdir -p ${LOCALNET_DIR}/${node_name}/cosmovisor/${VERSION_DIR}/bin/
        docker cp ${IMAGE_TAG}-inst:/go/bin/dcld ${LOCALNET_DIR}/${node_name}/cosmovisor/${VERSION_DIR}/bin/
    fi
done

docker rm ${IMAGE_TAG}-inst
