#!/bin/bash
if [ ! -d "vendor/github.com/muesli/smartcrop" ]; then
	gvt fetch github.com/muesli/smartcrop
fi
CGO_ENABLED=0 go build main.go
