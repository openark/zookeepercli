#!/bin/sh -e
# Build deb package for zookeepercli
# Sergei Kraev <skraev@tradingview.com>
#

echo "========= Preparing packing start =========";
apk update
apk add --no-cache \
    git \
    ruby \
    rpm-dev \
    tar \
    zip

git config --global --add safe.directory "$(pwd)"
gem install fpm
