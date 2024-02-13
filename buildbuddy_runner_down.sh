#!/usr/bin/env bash
set -euo pipefail
# Change to the WORKSPACE directory
cd "$BUILD_WORKSPACE_DIRECTORY"
curl https://keybase.io/codecovsecurity/pgp_keys.asc | gpg --no-default-keyring --keyring trustedkeys.gpg --import # One-time step

if [[ "$OSTYPE" == "darwin"* ]]; then
    # download Codecov CLI
    curl -Os https://cli.codecov.io/latest/macos/codecov

    # integrity check
    curl -Os https://cli.codecov.io/latest/macos/codecov
    curl -Os https://cli.codecov.io/latest/macos/codecov.SHA256SUM
    curl -Os https://cli.codecov.io/latest/macos/codecov.SHA256SUM.sig
    gpgv codecov.SHA256SUM.sig codecov.SHA256SUM
else
    apt-get install gnupg2
    apt-get remove gnupg
    ln -s /usr/bin/gpg2 /usr/bin/gpg
    # download Codecov CLI
    curl -Os https://cli.codecov.io/latest/linux/codecov

    # integrity check
    curl -Os https://cli.codecov.io/latest/linux/codecov
    curl -Os https://cli.codecov.io/latest/linux/codecov.SHA256SUM
    curl -Os https://cli.codecov.io/latest/linux/codecov.SHA256SUM.sig
    gpgv codecov.SHA256SUM.sig codecov.SHA256SUM
fi

shasum -a 256 -c codecov.SHA256SUM
chmod +x codecov
./codecov --help

export CODECOV_TOKEN=$TESTING_CODECOV_TOKEN

./codecov --verbose upload-process \
    --fail-on-error \
    --dir bazel-testlogs
