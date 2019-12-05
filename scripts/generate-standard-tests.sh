#!/bin/bash
d="$( cd "$( dirname "$0" )" || exit 1; cd ..; pwd )"
set -e
set -x

go run "$d/tools/standard-test-conv" -i "$d/vendor/WebAssembly/spec/test/core/address.wast" > "$d/scripts/standard-test-address.sh"
go run "$d/tools/standard-test-conv" -i "$d/vendor/WebAssembly/spec/test/core/i32.wast" > "$d/scripts/standard-test-i32.sh"
go run "$d/tools/standard-test-conv" -i "$d/vendor/WebAssembly/spec/test/core/i64.wast" > "$d/scripts/standard-test-i64.sh"
