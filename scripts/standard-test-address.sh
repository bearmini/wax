#!/bin/bash
d="$( cd "$( dirname "$0" )" || exit 1; cd ..; pwd )"
wax="$d/cmd/wax/wax"
spec_core_test="$d/vendor/WebAssembly/spec/test/core"
source "$d/scripts/test_common.sh"
set -x


a="$( $wax -f "8u_good1" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000061 97 97"
a="$( $wax -f "8u_good2" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000061 97 97"
a="$( $wax -f "8u_good3" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000062 98 98"
a="$( $wax -f "8u_good4" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000063 99 99"
a="$( $wax -f "8u_good5" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x0000007a 122 122"
a="$( $wax -f "8s_good1" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000061 97 97"
a="$( $wax -f "8s_good2" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000061 97 97"
a="$( $wax -f "8s_good3" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000062 98 98"
a="$( $wax -f "8s_good4" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000063 99 99"
a="$( $wax -f "8s_good5" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x0000007a 122 122"
a="$( $wax -f "16u_good1" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00006261 25185 25185"
a="$( $wax -f "16u_good2" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00006261 25185 25185"
a="$( $wax -f "16u_good3" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00006362 25442 25442"
a="$( $wax -f "16u_good4" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00006463 25699 25699"
a="$( $wax -f "16u_good5" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x0000007a 122 122"
a="$( $wax -f "16s_good1" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00006261 25185 25185"
a="$( $wax -f "16s_good2" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00006261 25185 25185"
a="$( $wax -f "16s_good3" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00006362 25442 25442"
a="$( $wax -f "16s_good4" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00006463 25699 25699"
a="$( $wax -f "16s_good5" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x0000007a 122 122"
a="$( $wax -f "32_good1" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x64636261 1684234849 1684234849"
a="$( $wax -f "32_good2" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x64636261 1684234849 1684234849"
a="$( $wax -f "32_good3" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x65646362 1701077858 1701077858"
a="$( $wax -f "32_good4" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x66656463 1717920867 1717920867"
a="$( $wax -f "32_good5" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x0000007a 122 122"
a="$( $wax -f "8u_good1" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8u_good2" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8u_good3" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8u_good4" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8u_good5" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8s_good1" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8s_good2" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8s_good3" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8s_good4" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8s_good5" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16u_good1" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16u_good2" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16u_good3" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16u_good4" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16u_good5" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16s_good1" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16s_good2" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16s_good3" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16s_good4" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16s_good5" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "32_good1" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "32_good2" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "32_good3" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "32_good4" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "32_good5" -a "i32:65507"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8u_good1" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8u_good2" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8u_good3" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8u_good4" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8u_good5" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8s_good1" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8s_good2" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8s_good3" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8s_good4" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "8s_good5" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16u_good1" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16u_good2" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16u_good3" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16u_good4" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16u_good5" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16s_good1" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16s_good2" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16s_good3" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16s_good4" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "16s_good5" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "32_good1" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "32_good2" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "32_good3" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
a="$( $wax -f "32_good4" -a "i32:65508"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i32:0x00000000 0 0"
set +e
a="$( $wax -f "32_good5" -a "i32:65508"      "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "8u_bad" -a "i32:0"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "8s_bad" -a "i32:0"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "16u_bad" -a "i32:0"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "16s_bad" -a "i32:0"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "32_bad" -a "i32:0"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "8u_bad" -a "i32:1"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "8s_bad" -a "i32:1"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "16u_bad" -a "i32:1"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "16s_bad" -a "i32:1"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "32_bad" -a "i32:1"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
a="$( $wax -f "8u_good1" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000061 97 97"
a="$( $wax -f "8u_good2" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000061 97 97"
a="$( $wax -f "8u_good3" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000062 98 98"
a="$( $wax -f "8u_good4" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000063 99 99"
a="$( $wax -f "8u_good5" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x000000000000007a 122 122"
a="$( $wax -f "8s_good1" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000061 97 97"
a="$( $wax -f "8s_good2" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000061 97 97"
a="$( $wax -f "8s_good3" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000062 98 98"
a="$( $wax -f "8s_good4" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000063 99 99"
a="$( $wax -f "8s_good5" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x000000000000007a 122 122"
a="$( $wax -f "16u_good1" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000006261 25185 25185"
a="$( $wax -f "16u_good2" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000006261 25185 25185"
a="$( $wax -f "16u_good3" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000006362 25442 25442"
a="$( $wax -f "16u_good4" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000006463 25699 25699"
a="$( $wax -f "16u_good5" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x000000000000007a 122 122"
a="$( $wax -f "16s_good1" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000006261 25185 25185"
a="$( $wax -f "16s_good2" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000006261 25185 25185"
a="$( $wax -f "16s_good3" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000006362 25442 25442"
a="$( $wax -f "16s_good4" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000006463 25699 25699"
a="$( $wax -f "16s_good5" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x000000000000007a 122 122"
a="$( $wax -f "32u_good1" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000064636261 1684234849 1684234849"
a="$( $wax -f "32u_good2" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000064636261 1684234849 1684234849"
a="$( $wax -f "32u_good3" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000065646362 1701077858 1701077858"
a="$( $wax -f "32u_good4" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000066656463 1717920867 1717920867"
a="$( $wax -f "32u_good5" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x000000000000007a 122 122"
a="$( $wax -f "32s_good1" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000064636261 1684234849 1684234849"
a="$( $wax -f "32s_good2" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000064636261 1684234849 1684234849"
a="$( $wax -f "32s_good3" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000065646362 1701077858 1701077858"
a="$( $wax -f "32s_good4" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000066656463 1717920867 1717920867"
a="$( $wax -f "32s_good5" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x000000000000007a 122 122"
a="$( $wax -f "64_good1" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x6867666564636261 7523094288207667809 7523094288207667809"
a="$( $wax -f "64_good2" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x6867666564636261 7523094288207667809 7523094288207667809"
a="$( $wax -f "64_good3" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x6968676665646362 7595434461045744482 7595434461045744482"
a="$( $wax -f "64_good4" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x6a69686766656463 7667774633883821155 7667774633883821155"
a="$( $wax -f "64_good5" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x000000000000007a 122 122"
a="$( $wax -f "8u_good1" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8u_good2" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8u_good3" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8u_good4" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8u_good5" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8s_good1" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8s_good2" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8s_good3" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8s_good4" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8s_good5" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16u_good1" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16u_good2" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16u_good3" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16u_good4" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16u_good5" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16s_good1" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16s_good2" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16s_good3" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16s_good4" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16s_good5" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32u_good1" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32u_good2" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32u_good3" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32u_good4" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32u_good5" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32s_good1" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32s_good2" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32s_good3" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32s_good4" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32s_good5" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "64_good1" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "64_good2" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "64_good3" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "64_good4" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "64_good5" -a "i32:65503"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8u_good1" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8u_good2" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8u_good3" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8u_good4" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8u_good5" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8s_good1" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8s_good2" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8s_good3" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8s_good4" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "8s_good5" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16u_good1" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16u_good2" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16u_good3" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16u_good4" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16u_good5" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16s_good1" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16s_good2" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16s_good3" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16s_good4" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "16s_good5" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32u_good1" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32u_good2" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32u_good3" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32u_good4" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32u_good5" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32s_good1" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32s_good2" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32s_good3" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32s_good4" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "32s_good5" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "64_good1" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "64_good2" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "64_good3" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
a="$( $wax -f "64_good4" -a "i32:65504"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:i64:0x0000000000000000 0 0"
set +e
a="$( $wax -f "64_good5" -a "i32:65504"      "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "8u_bad" -a "i32:0"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "8s_bad" -a "i32:0"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "16u_bad" -a "i32:0"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "16s_bad" -a "i32:0"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "32u_bad" -a "i32:0"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "32s_bad" -a "i32:0"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "64_bad" -a "i32:0"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "8u_bad" -a "i32:1"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "8s_bad" -a "i32:1"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "16u_bad" -a "i32:1"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "16s_bad" -a "i32:1"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "32u_bad" -a "i32:0"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "32s_bad" -a "i32:0"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "64_bad" -a "i32:1"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
a="$( $wax -f "32_good1" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f32:0x00000000 0.000000"
a="$( $wax -f "32_good2" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f32:0x00000000 0.000000"
a="$( $wax -f "32_good3" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f32:0x00000000 0.000000"
a="$( $wax -f "32_good4" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f32:0x00000000 0.000000"
a="$( $wax -f "32_good5" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f32:0x7fc00000 NaN"
a="$( $wax -f "32_good1" -a "i32:65524"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f32:0x00000000 0.000000"
a="$( $wax -f "32_good2" -a "i32:65524"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f32:0x00000000 0.000000"
a="$( $wax -f "32_good3" -a "i32:65524"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f32:0x00000000 0.000000"
a="$( $wax -f "32_good4" -a "i32:65524"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f32:0x00000000 0.000000"
a="$( $wax -f "32_good5" -a "i32:65524"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f32:0x00000000 0.000000"
a="$( $wax -f "32_good1" -a "i32:65525"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f32:0x00000000 0.000000"
a="$( $wax -f "32_good2" -a "i32:65525"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f32:0x00000000 0.000000"
a="$( $wax -f "32_good3" -a "i32:65525"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f32:0x00000000 0.000000"
a="$( $wax -f "32_good4" -a "i32:65525"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f32:0x00000000 0.000000"
set +e
a="$( $wax -f "32_good5" -a "i32:65525"      "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "32_bad" -a "i32:0"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "32_bad" -a "i32:1"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
a="$( $wax -f "64_good1" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f64:0x00000000 0.000000"
a="$( $wax -f "64_good2" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f64:0x00000000 0.000000"
a="$( $wax -f "64_good3" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f64:0x00000000 0.000000"
a="$( $wax -f "64_good4" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f64:0x00000000 0.000000"
a="$( $wax -f "64_good5" -a "i32:0"          "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f64:0x7ff8000000000001 NaN"
a="$( $wax -f "64_good1" -a "i32:65510"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f64:0x00000000 0.000000"
a="$( $wax -f "64_good2" -a "i32:65510"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f64:0x00000000 0.000000"
a="$( $wax -f "64_good3" -a "i32:65510"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f64:0x00000000 0.000000"
a="$( $wax -f "64_good4" -a "i32:65510"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f64:0x00000000 0.000000"
a="$( $wax -f "64_good5" -a "i32:65510"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f64:0x00000000 0.000000"
a="$( $wax -f "64_good1" -a "i32:65511"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f64:0x00000000 0.000000"
a="$( $wax -f "64_good2" -a "i32:65511"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f64:0x00000000 0.000000"
a="$( $wax -f "64_good3" -a "i32:65511"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f64:0x00000000 0.000000"
a="$( $wax -f "64_good4" -a "i32:65511"      "$spec_core_test/address.wasm" )"; assert_equal "$a" "0:f64:0x00000000 0.000000"
set +e
a="$( $wax -f "64_good5" -a "i32:65511"      "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "64_bad" -a "i32:0"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
set +e
a="$( $wax -f "64_bad" -a "i32:1"          "$spec_core_test/address.wasm" 2>&1 )"; assert_contains "$a" "out of bounds memory access"
set -e
