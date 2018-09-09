#!/bin/bash
export GOPATH=~/go/
echo "Delete existing repositories..."

rm -rf $GOPATH/src/github.com
rm -rf $GOPATH/src/golang.org
rm -rf $GOPATH/src/google.golang.org

echo "Downloading Negroni"
go get -u -x github.com/urfave/negroni

echo "Downloading go-sqlite3"
go get -u -x github.com/mattn/go-sqlite3

echo "Downloading Gorilla"
go get github.com/gorilla/mux

echo "Downloading go-gorp"
go get github.com/go-gorp/gorp

echo "Download Gorilla Sessions"
go get github.com/gorilla/sessions

echo "Downloading bcrypt"
go get golang.org/x/crypto/bcrypt

echo "Downloading App engine sdk"
go get google.golang.org/appengine

