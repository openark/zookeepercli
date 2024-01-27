#!/bin/sh -e
# Build deb package for zookeepercli
# Sergei Kraev <skraev@tradingview.com>
#

echo "========= Preparing packing start =========";
PKGNAME="zookeepercli"
VER=$(git describe --long --tags --always --abbrev=10 | sed 's/^[^0-9]//ig')
ARCH=$(arch | sed s/aarch64/arm64/ | sed s/x86_64/amd64/)
RDIR="$(pwd)/pkgs_out/"

rm -rf "${RDIR}"
mkdir -p "$RDIR"

OUTTYPES="rpm deb apk tar zip"

for OUTTYPE in $OUTTYPES; do
    echo "========= Packing ${OUTTYPE} start =========";
    fpm \
        --output-type "${OUTTYPE}" \
        --input-type dir \
        --force \
        \
        --name "${PKGNAME}" \
        --package "${RDIR}" \
        --version "${VER}" \
        --architecture "${ARCH}" \
        --maintainer 'Shlomi Noach <shlomi-noach@github.com>' \
        --url 'https://github.com/openark/zookeepercli' \
        --description 'Zookeeper client console' \
        --license 'Apache 2.0' \
        --category 'universe/net' \
        --no-depends --no-auto-depends \
        --prefix /usr/local/bin \
        --chdir "./bin" .
done
