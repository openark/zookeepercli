#!/bin/bash

# Simple packaging of zookeepercli
#
# Requires fpm: https://github.com/jordansissel/fpm
#

release_version="1.0.10"
release_dir=/tmp/zookeepercli
rm -rf $release_dir/*
mkdir -p $release_dir

cd  $(dirname $0)
for f in $(find . -name "*.go"); do go fmt $f; done

go build -o $release_dir/zookeepercli

if [[ $? -ne 0 ]] ; then
	exit 1
fi

cd $release_dir
# rpm packaging
fpm -v "${release_version}" -f -s dir -t rpm -n zookeepercli -C $release_dir --prefix=/usr/bin .
fpm -v "${release_version}" -f -s dir -t deb -n zookeepercli -C $release_dir --prefix=/usr/bin .

echo "---"
echo "Done. Find releases in $release_dir"
