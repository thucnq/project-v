#!/usr/bin/env bash

if ! [ -x "$(command -v CompileDaemon)" ]; then
    echo '---------------------------------------------------'
    echo '> CompileDaemon is not installed.                 <'
    echo '> Run the following command to install the binary <'
    echo '> go get github.com/githubnemo/CompileDaemon      <'
    echo '---------------------------------------------------'
    export GO111MODULE=off
    go get github.com/githubnemo/CompileDaemon
    export GO111MODULE=on
fi

trap "rm main; exit" SIGHUP SIGINT SIGTERM

CompileDaemon -log-prefix=false -build="go build -x -mod=mod ./cmd/server/main.go ./cmd/server/bootstrap.go ./cmd/server/server.go ./cmd/server/i18n.go" -command="./main" -exclude-dir=".git" -exclude-dir="cmd/client"  -exclude-dir=".idea" -exclude-dir="vendor" -color
