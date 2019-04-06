#!/bin/bash
echo "Downloading Negroni"
govendor fetch github.com/urfave/negroni

echo "Downloading Gorilla"
govendor fetch github.com/gorilla/mux

echo "Download Gorilla Sessions"
govendor fetch github.com/gorilla/sessions

echo "Downloading bcrypt"
govendor fetch golang.org/x/crypto/bcrypt

echo "Downloading App engine sdk"
govendor fetch google.golang.org/appengine

echo "Downloading Datastore SDK"
govendor fetch cloud.google.com/go/datastore

