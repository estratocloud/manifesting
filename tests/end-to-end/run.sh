#!/bin/sh

set -e

path=/app/tests/end-to-end

rm -rf $path/.generated

go run . --working-dir=$path

diff -r $path/expected $path/.generated

echo "End-to-End test completed successfully"
