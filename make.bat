@echo off

set TARGET=calculator.exe
set COMMAND="%~1"

if %COMMAND%=="" goto build
if %COMMAND%=="build" goto build

if %COMMAND%=="clean" goto clean

if %COMMAND%=="test" goto test

echo "Invalid command specificed"
echo "Usage: make [build|clean|test]"

goto end

:build
go build -o ./%TARGET%
goto end

:clean
if exist .\%TARGET% del /Q .\%TARGET%
goto end

:test
go test ./...
goto end

:end