set -xe

go env -w GO111MODULE=on
go get

go build -o bin/application application.go