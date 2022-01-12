source integration_tests/cli/common.sh

LOCALNET_DIR=".localnet"
SED_EXT=

patch_consensus_config() {
  local NODE_CONFIGS="$(find "$LOCALNET_DIR" -type f -name "config.toml" -wholename "*node*")"

  for NODE_CONFIG in ${NODE_CONFIGS}; do
    sed -i $SED_EXT 's/timeout_propose = "3s"/timeout_propose = "500ms"/g' "${NODE_CONFIG}"
    sed -i $SED_EXT 's/timeout_prevote = "1s"/timeout_prevote = "500ms"/g' "${NODE_CONFIG}"
    sed -i $SED_EXT 's/timeout_precommit = "1s"/timeout_precommit = "500ms"/g' "${NODE_CONFIG}"
    sed -i $SED_EXT 's/timeout_commit = "5s"/timeout_commit = "500ms"/g' "${NODE_CONFIG}"
  done
}

make localnet_clean
make install
make image
make localnet_init
patch_consensus_config
make localnet_start
wait_for_height 2 20

