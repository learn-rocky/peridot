#!/usr/bin/env bash

set -o errexit

source hack/bazel_setup.sh

$BAZEL_B $($BAZEL_QR query "//... except attr(tags, 'manual', //...) except //vendor/...")
