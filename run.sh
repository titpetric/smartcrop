#!/bin/bash
docker run --net=party -e GOOS=linux -e GOARCH=386 -e CGO_ENABLED=0 --rm=true -i -v $(pwd):/go/src/smartcrop -w /go/src/smartcrop golang:1.9-alpine go build -o main32 main.go
docker run --net=party -e GOOS=linux -e GOARCH=amd64 -e CGO_ENABLED=0 --rm=true -i -v $(pwd):/go/src/smartcrop -w /go/src/smartcrop golang:1.9-alpine go build -o main64 main.go
