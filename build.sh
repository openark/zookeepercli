#!/bin/bash -e

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
find . -name "*.go" -exec go fmt {} \;

GOPATH="$(pwd)"
export GOPATH
printf '%s\n' "getting github.com/outbrain/golib/log"
go get github.com/outbrain/golib/log
printf '%s\n' "getting github.com/samuel/go-zookeeper/zk"
go get github.com/samuel/go-zookeeper/zk
printf '%s\n' "building ./src/github.com/outbrain/zookeepercli/main.go"
go build -o $release_dir/zookeepercli ./src/github.com/outbrain/zookeepercli/main.go

if [[ $? -ne 0 ]] ; then
	exit 1
fi

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
