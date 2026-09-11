@echo off
rem Compila la librería compartida para Windows (.dll) usando mingw-w64
rem Requiere mingw-w64-gcc instalado y en el PATH.
setlocal
cd /d %~dp0..\src\core
set GOOS=windows
set GOARCH=amd64
set CGO_ENABLED=1
set CC=x86_64-w64-mingw32-gcc
go build -buildmode=c-shared -o ..\..\lib\validador.dll .
echo OK: lib\validador.dll
endlocal