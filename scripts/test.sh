#!/usr/bin/env bash
d="$( cd "$( dirname "$0" )" || exit 1; cd ..; pwd )"
set -x
set -e

source "$d/scripts/test_common.sh"

wax="$d/cmd/wax/wax"

a="$( $wax -f "add" -a "i32:123" -a "i32:234" "$d/examples/go/add/main.wasm" )"; assert_equal "$a" "0:i32:0x00000165 357 357"
a="$( $wax -f "sub" -a "i32:123" -a "i32:23"  "$d/examples/go/add/main.wasm" )"; assert_equal "$a" "0:i32:0x00000064 100 100"
a="$( $wax -f "sub" -a "i32:123" -a "i32:234" "$d/examples/go/add/main.wasm" )"; assert_equal "$a" "0:i32:0xffffff91 4294967185 -111"
a="$( $wax -f "mul" -a "i32:13"  -a "i32:17"  "$d/examples/go/add/main.wasm" )"; assert_equal "$a" "0:i32:0x000000dd 221 221"
a="$( $wax -f "div" -a "i32:10"  -a "i32:3"   "$d/examples/go/add/main.wasm" )"; assert_equal "$a" "0:i32:0x00000003 3 3"

a="$( $wax -f "fib" -a "i32:20" "$d/examples/go/fib/main.wasm" )"; assert_equal "$a" "0:i32:0x00001a6d 6765 6765"

a="$( $wax -f "add" -a "i32:123" -a "i32:234" "$d/examples/rust/add/main-stripped.wasm" )"; assert_equal "$a" "0:i32:0x00000165 357 357"
a="$( $wax -f "sub" -a "i32:123" -a "i32:23"  "$d/examples/rust/add/main-stripped.wasm" )"; assert_equal "$a" "0:i32:0x00000064 100 100"
a="$( $wax -f "sub" -a "i32:123" -a "i32:234" "$d/examples/rust/add/main-stripped.wasm" )"; assert_equal "$a" "0:i32:0xffffff91 4294967185 -111"
a="$( $wax -f "mul" -a "i32:13"  -a "i32:17"  "$d/examples/rust/add/main-stripped.wasm" )"; assert_equal "$a" "0:i32:0x000000dd 221 221"
a="$( $wax -f "div" -a "i32:10"  -a "i32:3" "$d/examples/rust/add/main-stripped.wasm" )"; assert_equal "$a" "0:i32:0x00000003 3 3"

a="$( $wax -f "block_test" -a "i32:10" -a "i32:20" "$d/examples/wat/block_test/main.wasm" )"; assert_equal "$a" "0:i32:0x00000001 1 1"
a="$( $wax -f "block_test" -a "i32:0"  -a "i32:20" "$d/examples/wat/block_test/main.wasm" )"; assert_equal "$a" "0:i32:0x00000001 1 1"
a="$( $wax -f "block_test" -a "i32:-1" -a "i32:20" "$d/examples/wat/block_test/main.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "block_test" -a "i32:10" -a "i32:49" "$d/examples/wat/block_test/main.wasm" )"; assert_equal "$a" "0:i32:0x00000001 1 1"
a="$( $wax -f "block_test" -a "i32:10" -a "i32:50" "$d/examples/wat/block_test/main.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"

a="$( $wax -f "loop3" "$d/examples/wat/loop_test/main.wasm" )"; assert_equal "$a" "0:i32:0x00000003 3 3"

set +e
a="$( $wax -f "infinite_loop" -t 3 "$d/examples/go/loop/main.wasm"   2>&1 )"; assert_contains "$a" "context deadline exceeded"
a="$( $wax -f "infinite_loop" -t 3 "$d/examples/rust/loop/main.wasm" 2>&1 )"; assert_contains "$a" "context deadline exceeded"