#!/bin/bash
docker run --net=party -e CGO_ENABLED=0 --rm=true -i -v $(pwd):/go/src/smartcrop -w /go/src/smartcrop golang:1.9-alpine go build main.go
