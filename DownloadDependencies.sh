#!/bin/bash
export GOPATH=$(pwd)
echo "Delete existing repositories..."

rm -rf $GOPATH/src/github.com

echo "Downloading Negroni"
go get -u -x github.com/urfave/negroni

echo "Downloading go-sqlite3"
go get -u -x github.com/mattn/go-sqlite3

echo "Downloading Gorilla"
go get github.com/gorilla/mux

echo "Downloading go-gorp"
go get github.com/gorilla/mux

echo "Downloading negroni-sessions"
go get github.com/GoIncremental/negroni-sessions

echo "Downloading bcrypt"
go get golang.org/x/crypto/bcrypt

