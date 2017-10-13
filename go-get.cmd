@ECHO off
echo Installing go modules for Compiler
echo Installing ANTLR4 Runtime for Go
go get github.com/antlr/antlr4/runtime/Go/antlr
cd %GOPATH%/src/github.com/antlr/antlr4
echo Setting ANTLR4 runtime to release 4.7
git checkout tags/4.7
echo Done!
