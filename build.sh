#!/usr/bin/env bash -e

set -o pipefail

# Simple packaging of zookeepercli
#
# Requires fpm: https://github.com/jordansissel/fpm
#

platform=$(uname -s)
release_version="1.0.10"
release_dir=/tmp/zookeepercli
rm -rf ${release_dir:?}/*
mkdir -p $release_dir

pushd "$(dirname "$0")"
for f in $(find . -name "*.go"); do go fmt $f; done

go build -o $release_dir/zookeepercli || exit 1

if [ "$platform" = "Linux" ]; then
  pushd "$release_dir"
  # rpm packaging
  fpm -v "${release_version}" -f -s dir -t rpm -n zookeepercli -C $release_dir --prefix=/usr/bin .
  fpm -v "${release_version}" -f -s dir -t deb -n zookeepercli -C $release_dir --prefix=/usr/bin .
  popd
fi

echo "---"
echo "Done. Find releases in $release_dir"

popd
