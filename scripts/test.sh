#!/usr/bin/env bash
d="$( cd "$( dirname "$0" )" || exit 1; cd ..; pwd )"
set -x
set -e

assert_equal() {
    if [ "$1" != "$2" ]; then
        echo "Expected: $2"
        echo "Actual:   $1"
        echo "Diff:"
        diff <(echo "$1") <(echo "$2")
        exit 1
    fi
}

assert_contains() {
    if ! [[ "$1" == *"$2"* ]]; then
        echo "Expected to be contained: $2"
        echo "Actual:                   $1"
        exit 1
    fi
}

wax="$d/cmd/wax/wax"

a="$( $wax -f "add" -a "i32:123" -a "i32:234" "$d/examples/go/add/main.wasm" )"
assert_equal "$a" "0:i32:357"

a="$( $wax -f "sub" -a "i32:123" -a "i32:23" "$d/examples/go/add/main.wasm" )"
assert_equal "$a" "0:i32:100"

a="$( $wax -f "sub" -a "i32:123" -a "i32:234" "$d/examples/go/add/main.wasm" )"
assert_equal "$a" "0:i32:4294967185" # -111

a="$( $wax -f "mul" -a "i32:13" -a "i32:17" "$d/examples/go/add/main.wasm" )"
assert_equal "$a" "0:i32:221"

a="$( $wax -f "div" -a "i32:10" -a "i32:3" "$d/examples/go/add/main.wasm" )"
assert_equal "$a" "0:i32:3"

a="$( $wax -f "fib" -a "i32:20" "$d/examples/go/fib/main.wasm" )"
assert_equal "$a" "0:i32:6765"

a="$( $wax -f "add" -a "i32:123" -a "i32:234" "$d/examples/rust/add/main-stripped.wasm" )"
assert_equal "$a" "0:i32:357"

a="$( $wax -f "sub" -a "i32:123" -a "i32:23" "$d/examples/rust/add/main-stripped.wasm" )"
assert_equal "$a" "0:i32:100"

a="$( $wax -f "sub" -a "i32:123" -a "i32:234" "$d/examples/rust/add/main-stripped.wasm" )"
assert_equal "$a" "0:i32:4294967185" # -111

a="$( $wax -f "mul" -a "i32:13" -a "i32:17" "$d/examples/rust/add/main-stripped.wasm" )"
assert_equal "$a" "0:i32:221"

a="$( $wax -f "div" -a "i32:10" -a "i32:3" "$d/examples/rust/add/main-stripped.wasm" )"
assert_equal "$a" "0:i32:3"

a="$( $wax -f "block_test" -a "i32:10" -a "i32:20" "$d/examples/wat/block_test/main.wasm" )"
assert_equal "$a" "0:i32:1"

a="$( $wax -f "block_test" -a "i32:0" -a "i32:20" "$d/examples/wat/block_test/main.wasm" )"
assert_equal "$a" "0:i32:1"

a="$( $wax -f "block_test" -a "i32:-1" -a "i32:20" "$d/examples/wat/block_test/main.wasm" )"
assert_equal "$a" "0:i32:0"

a="$( $wax -f "block_test" -a "i32:10" -a "i32:49" "$d/examples/wat/block_test/main.wasm" )"
assert_equal "$a" "0:i32:1"

a="$( $wax -f "block_test" -a "i32:10" -a "i32:50" "$d/examples/wat/block_test/main.wasm" )"
assert_equal "$a" "0:i32:0"

a="$( $wax -f "loop3" "$d/examples/wat/loop_test/main.wasm" )"
assert_equal "$a" "0:i32:3"

set +e
a="$( $wax -f "infinite_loop" -t 3 "$d/examples/go/loop/main.wasm" 2>&1 )"
assert_contains "$a" "context deadline exceeded"

a="$( $wax -f "infinite_loop" -t 3 "$d/examples/rust/loop/main.wasm" 2>&1 )"
assert_contains "$a" "context deadline exceeded"