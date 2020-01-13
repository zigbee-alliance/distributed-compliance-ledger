############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Git is required for fetching the dependencies,
# make is required for building.
RUN apk update && apk add --no-cache git make

WORKDIR /go/src/zb-ledger
COPY . .

RUN make

############################
# STEP 2 build an image
############################
FROM alpine:latest

COPY --from=builder /go/bin/zbld /usr/bin/zbld
COPY --from=builder /go/bin/zblcli /usr/bin/zblcli

VOLUME /root/.zbld
VOLUME /root/.zblcli

EXPOSE 26656 26657

STOPSIGNAL SIGTERM
