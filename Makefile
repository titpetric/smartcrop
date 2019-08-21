.PHONY: all build test docker push

IMAGE := registry.dev.rtvslo.si/library/smartcrop
GOOS := linux
GOARCH := amd64

all:
	drone exec

build:
	go fmt ./...
	GOOS=$(GOOS) GOARCH=$(GOARCH) CGO_ENABLED=0 go build -o build/smartcrop-$(GOOS)-$(GOARCH) cmd/smartcrop/*.go

test: build
	find ./tests -type f -name '*.jpg' | xargs -n1 -I{} ./build/smartcrop-$(GOOS)-$(GOARCH) -output={}.json {}

docker:
	docker build --rm --no-cache -t $(IMAGE) .

push:
	docker push $(IMAGE)
