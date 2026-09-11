#!/usr/bin/env bash
# Compila la librería compartida para Windows (.dll) usando mingw-w64
# Requiere: sudo pacman -S mingw-w64-gcc
set -euo pipefail

cd "$(dirname "$0")/../src/core"
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -buildmode=c-shared -o ../../lib/validador.dll .
echo "OK: lib/validador.dll"