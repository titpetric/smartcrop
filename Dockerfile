FROM golang:1.12-alpine as builder

WORKDIR /smartcrop

ADD . .

# should match everything in .drone.yml

RUN apk --no-cache add make git
RUN make build

FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates tzdata

ENV TZ Europe/Ljubljana

RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY --from=builder /smartcrop/build/* /app/

ENTRYPOINT ["/app/smartcrop-linux-amd64"]
CMD ["-mode=http"]