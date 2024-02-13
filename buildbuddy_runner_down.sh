#!/usr/bin/env bash
set -euo pipefail
# Change to the WORKSPACE directory
cd "$BUILD_WORKSPACE_DIRECTORY"

if [[ "$OSTYPE" == "darwin"* ]]; then
    # download Codecov CLI
    curl https://keybase.io/codecovsecurity/pgp_keys.asc | gpg --no-default-keyring --keyring trustedkeys.gpg --import # One-time step
    curl -Os https://cli.codecov.io/latest/macos/codecov

    # integrity check
    curl -Os https://cli.codecov.io/latest/macos/codecov
    curl -Os https://cli.codecov.io/latest/macos/codecov.SHA256SUM
    curl -Os https://cli.codecov.io/latest/macos/codecov.SHA256SUM.sig
else
    apt-get update && apt-get install gnupg
    curl https://keybase.io/codecovsecurity/pgp_keys.asc | gpg --no-default-keyring --keyring trustedkeys.gpg --import # One-time step
    # download Codecov CLI
    curl -Os https://cli.codecov.io/latest/linux/codecov

    # integrity check
    curl -Os https://cli.codecov.io/latest/linux/codecov
    curl -Os https://cli.codecov.io/latest/linux/codecov.SHA256SUM
    curl -Os https://cli.codecov.io/latest/linux/codecov.SHA256SUM.sig
fi

gpgv codecov.SHA256SUM.sig codecov.SHA256SUM
shasum -a 256 -c codecov.SHA256SUM
chmod +x codecov
./codecov --help

export CODECOV_TOKEN=$TESTING_CODECOV_TOKEN

./codecov --verbose upload-process \
    --fail-on-error \
    --dir bazel-testlogs
