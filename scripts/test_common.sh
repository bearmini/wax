#!/bin/bash

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
