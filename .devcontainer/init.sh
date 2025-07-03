#!/bin/bash

echo "run .devcontainer/init/sh"
echo "install packages"
go clean -cache
go clean -modcache
# VS Code Go 開発関連
go install golang.org/x/tools/gopls@latest
go install github.com/cweill/gotests/gotests@v1.6.0
go install github.com/fatih/gomodifytags@v1.17.0
go install github.com/josharian/impl@v1.4.0
go install github.com/haya14busa/goplay/cmd/goplay@v1.0.0
go install github.com/go-delve/delve/cmd/dlv@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
go install golang.org/x/tools/cmd/goimports@latest
# Protocol Buffer 関連
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
echo "finish installing packages"
echo "pwd"
pwd
PB_REL="https://github.com/protocolbuffers/protobuf/releases"
curl -LO $PB_REL/download/v30.2/protoc-30.2-linux-x86_64.zip
unzip protoc-30.2-linux-x86_64.zip -d $HOME/.local
rm protoc-30.2-linux-x86_64.zip
echo 'export PATH="$PATH:$HOME/.local/bin"' >> ~/.bashrc

