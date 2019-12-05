#!/usr/bin/env bash
d="$( cd "$( dirname "$0" )" || exit 1; cd ..; pwd )"
set -x
set -e

export GO111MODULE=on
USE_DOCKER_TINYGO=0
TINYGO_VERSION=0.8.0
WASM_EXECUTABLE=wasm

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

wasm_available() {
  command -v wasm > /dev/null 2>&1 && {
    return 0
  }

  vendor_interpreter="$d/vendor/WebAssembly/spec/interpreter/wasm"
  if [ -x "$vendor_interpreter" ]; then
    WASM_EXECUTABLE="$vendor_interpreter"
    return 0
  fi
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

build_with_wasm() {
  dir=$1
  src=$2

  cd "$dir"
  $WASM_EXECUTABLE -d "$dir/$src" -o "$dir/${src/%\.wast/.wasm}"
  cd -
}

strip() {
  dir=$1
  src=$2

  wasm=${src/%.rs/.wasm}
  out=${wasm/%.wasm/-stripped.wasm}
  "$d/cmd/wastrip/wastrip" -o "$dir/$out" "$dir/$wasm"
}

go mod download

build_with_go "$d/cmd/wadisasm"
build_with_go "$d/cmd/wadump"
build_with_go "$d/cmd/wastrip"
build_with_go "$d/cmd/wax"

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

if wasm_available; then
  spectestdir="$d/vendor/WebAssembly/spec/test/core"
  build_with_wasm "$spectestdir" address.wast
  build_with_wasm "$spectestdir" block.wast
  build_with_wasm "$spectestdir" i32.wast
  build_with_wasm "$spectestdir" i64.wast

  "$d/scripts/generate-standard-tests.sh"
else
  echo "skipping building spec tests written in wast, as 'wasm' is not available"
fi