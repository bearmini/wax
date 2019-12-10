#!/usr/bin/env bash
d="$( cd "$( dirname "$0" )" || exit 1; cd ..; pwd )"
set -x
set -e

export GO111MODULE=on
USE_DOCKER_TINYGO=0
TINYGO_VERSION=0.8.0

tinygo_available() {
  command -v tinygo > /dev/null 2>&1 && {
    return 0
  }
  command -v docker > /dev/null 2>&1 && {
    docker pull tinygo/tinygo:$TINYGO_VERSION
    USE_DOCKER_TINYGO=1
    return 0
  }
  return 1
}

rustc_available() {
  command -v rustc > /dev/null 2>&1 && {
    return 0
  }
  return 1
}

wat2wasm_available() {
  command -v wat2wasm > /dev/null 2>&1 && {
    return 0
  }
  return 1
}

build_with_go() {
  dir=$1

  cd "$dir"
  go build
  cd -
}

build_with_tinygo() {
  dir=$1
  src=$2

  cd "$dir"
  if [ $USE_DOCKER_TINYGO -eq 0 ]; then
    tinygo build -target wasm -o "$dir/${src/.go/.wasm}" "$dir/$src"
  else
    docker run --rm -v "$dir":/src tinygo/tinygo:$TINYGO_VERSION tinygo build -target wasm -o "/src/${src/.go/.wasm}" "/src/$src"
  fi
  cd -
}

build_with_rustc() {
  dir=$1
  src=$2

  cd "$dir"
  rustc --crate-type=cdylib --target wasm32-unknown-unknown -O "$src"
  strip "$dir" "$src"
  cd -
}

build_with_wat2wasm() {
  dir=$1
  src=$2

  cd "$dir"
  wat2wasm "$src"
  cd -
}

strip() {
  dir=$1
  src=$2

  wasm=${src/%.rs/.wasm}
  out=${wasm/%.wasm/-stripped.wasm}
  "$d/cmd/wastrip/wastrip" -o "$dir/$out" "$dir/$wasm"
}

cd "$d" && {
  go mod download
  git submodule update --init
}

build_with_go "$d/cmd/wadisasm"
build_with_go "$d/cmd/wadump"
build_with_go "$d/cmd/wastrip"
build_with_go "$d/cmd/wax"
build_with_go "$d/tools/std-test"

wasm_bin="$d/vendor/WebAssembly/spec/interpreter/wasm"
std_test_bin="$d/tools/std-test/std-test"
std_test_wast_dir="$d/vendor/WebAssembly/spec/test/core"
#"$std_test_bin" -d "$std_test_wast_dir"                            -w "$wasm_bin"
"$std_test_bin" -i "$std_test_wast_dir/address.wast"                -w "$wasm_bin"
"$std_test_bin" -i "$std_test_wast_dir/align.wast"                  -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/binary-leb128.wast"          -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/binary.wast"                 -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/block.wast"                  -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/br_if.wast"                  -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/br.wast"                     -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/break-drop.wast"             -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/call_indirect.wast"          -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/call.wast"                   -w "$wasm_bin"
"$std_test_bin" -i "$std_test_wast_dir/comments.wast"               -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/const.wast"                  -w "$wasm_bin"
"$std_test_bin" -i "$std_test_wast_dir/conversions.wast"            -w "$wasm_bin"
"$std_test_bin" -i "$std_test_wast_dir/custom.wast"                 -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/data.wast"                   -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/elem.wast"                   -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/endianness.wast"             -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/exports.wast"                -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/f32_bitwise.wast"            -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/f32_cmp.wast"                -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/f32.wast"                    -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/f64_bitwise.wast"            -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/f64_cmp.wast"                -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/f64.wast"                    -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/fac.wast"                    -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/float_exprs.wast"            -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/float_literals.wast"         -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/float_memory.wast"           -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/float_misc.wast"             -w "$wasm_bin"
"$std_test_bin" -i "$std_test_wast_dir/forward.wast"                -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/func_ptrs.wast"              -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/func.wast"                   -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/globals.wast"                -w "$wasm_bin"
"$std_test_bin" -i "$std_test_wast_dir/i32.wast"                    -w "$wasm_bin"
"$std_test_bin" -i "$std_test_wast_dir/i64.wast"                    -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/if.wast"                     -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/imports.wast"                -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/inline-module.wast"          -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/init_exprs.wast"             -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/labels.wast"                 -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/left-to-right.wast"          -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/linking.wast"                -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/load.wast"                   -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/local_get.wast"              -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/local_set.wast"              -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/local_tee.wast"              -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/loop.wast"                   -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/memory_grow.wast"            -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/memory_redundancy.wast"      -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/memory_size.wast"            -w "$wasm_bin"
"$std_test_bin" -i "$std_test_wast_dir/names.wast"                  -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/nop.wast"                    -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/return.wast"                 -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/select.wast"                 -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/skip-stack-guard-page.wast"  -w "$wasm_bin"
"$std_test_bin" -i "$std_test_wast_dir/stack.wast"                  -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/start.wast"                  -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/store.wast"                  -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/switch.wast"                 -w "$wasm_bin"
"$std_test_bin" -i "$std_test_wast_dir/token.wast"                  -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/traps.wast"                  -w "$wasm_bin"
"$std_test_bin" -i "$std_test_wast_dir/type.wast"                   -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/unreachable.wast"            -w "$wasm_bin"
"$std_test_bin" -i "$std_test_wast_dir/unreached-invalid.wast"      -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/unwind.wast"                 -w "$wasm_bin"
"$std_test_bin" -i "$std_test_wast_dir/utf8-custom-section-id.wast" -w "$wasm_bin"
"$std_test_bin" -i "$std_test_wast_dir/utf8-import-field.wast"      -w "$wasm_bin"
"$std_test_bin" -i "$std_test_wast_dir/utf8-import-module.wast"     -w "$wasm_bin"
#"$std_test_bin" -i "$std_test_wast_dir/utf8-invalid-encoding.wast"  -w "$wasm_bin"

if tinygo_available; then
  build_with_tinygo "$d/examples/go/add" main.go
  build_with_tinygo "$d/examples/go/fib" main.go
  build_with_tinygo "$d/examples/go/loop" main.go
  build_with_tinygo "$d/examples/go/string" main.go
else
  echo "skipping building examples written in go, as 'tinygo' is not available"
fi

if rustc_available; then
  build_with_rustc "$d/examples/rust/add" main.rs
  build_with_rustc "$d/examples/rust/loop" main.rs
else
  echo "skipping building examples written in rust, as 'rustc' is not available"
fi

if wat2wasm_available; then
  build_with_wat2wasm "$d/examples/wat/block_test" main.wat
  build_with_wat2wasm "$d/examples/wat/loop_test" main.wat
else
  echo "skipping building examples written in wat, as 'wat2wasm' is not available"
fi
