source integration_tests/cli/common.sh

LOCALNET_DIR=".localnet"
SED_EXT=

  # patch configs properly by having all values >= 1 sec, otherwise headers may start having time from the future and light client verification will fail
  # if we patch config to have new blocks created in less than 1 sec, the min time in a time header is still 1 sec.
  # So, new blocks started to be from the future.
patch_consensus_config() {
  local NODE_CONFIGS="$(find "$LOCALNET_DIR" -type f -name "config.toml" -wholename "*node*")"

  for NODE_CONFIG in ${NODE_CONFIGS}; do
    sed -i $SED_EXT 's/timeout_propose = "3s"/timeout_propose = "1s"/g' "${NODE_CONFIG}"
    #sed -i $SED_EXT 's/timeout_prevote = "1s"/timeout_prevote = "1s"/g' "${NODE_CONFIG}"
    #sed -i $SED_EXT 's/timeout_precommit = "1s"/timeout_precommit = "1s"/g' "${NODE_CONFIG}"
    sed -i $SED_EXT 's/timeout_commit = "5s"/timeout_commit = "1s"/g' "${NODE_CONFIG}"
  done
}

make install image localnet_rebuild
patch_consensus_config
make localnet_start
wait_for_height 2 20

