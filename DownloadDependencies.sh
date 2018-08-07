#!/bin/bash
export GOPATH=$(pwd)
echo "Delete existing repositories..."

rm -rf $GOPATH/src/github.com

echo "Downloading Negroni"
go get -u -x github.com/urfave/negroni

echo "Downloading go-sqlite3"
go get -u -x github.com/mattn/go-sqlite3
