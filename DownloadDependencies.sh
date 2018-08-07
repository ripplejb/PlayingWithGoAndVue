#!/bin/bash
export GOPATH=$(pwd)
echo "Deleting existing ..."

rm -rf $GOPATH/src/github.com

echo "Downloading Negroni"
go get -u -x github.com/urfave/negroni

echo "Downloading go-sqlite3"
go get -u -x github.com/mattn/go-sqlite3
