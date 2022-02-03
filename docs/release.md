## Release

The steps:

1.  [draft](https://github.com/zigbee-alliance/distributed-compliance-ledger/releases/new) and publish a new GitHub release:
    *   specify a new release tag based on a new planned version of DCLedger in the format `v<version>` (e.g. `v1.2.3` for the version `1.2.3`)
    *   verify the branch/commit target for the release: usually it should be `master` but other targets are possible as well
    *   put release notes
2.  once published [release pipeline](https://github.com/zigbee-alliance/distributed-compliance-ledger/actions/workflows/release.yml) is triggered:
    *   it builds `dcld` binary on `ubuntu-20.04` and `macos-11` and attaches the artifacts to the GitHub release:
        *   binary and archived binary for `ubuntu-20.04`
        *   archived binary for `macos-11`
        *   systemd service file
3.  additional way to trigger the pipeline is to do that manually, it can be used for the following cases:
    *   some intermittent issue happened during the normal build so some artifact hasn't been attached
    *   testing / debugging
