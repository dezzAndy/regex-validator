#!/usr/bin/env bash
# Compila la librería compartida para Linux (.so)
set -euo pipefail

cd "$(dirname "$0")/../src/core"
CGO_ENABLED=1 go build -buildmode=c-shared -o ../../lib/libvalidador.so .
echo "OK: lib/libvalidador.so"