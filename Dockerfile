#FROM golang:1.22 as golang
#
#RUN mkdir -p /
#
#WORKDIR /
#
#COPY . .
#
#RUN make build
##
##EXPOSE 8888
##
###ENTRYPOINT ["/rest-app"]
##CMD ["./rest-app"]

# syntax = docker/dockerfile:1
#ARG GOLANG=golang:1.22
#
#
#
#FROM ${GOLANG} AS base
#
#FROM ${GOLANG} AS builder
#
#WORKDIR /src
#COPY ./ ./
#
#RUN --mount=type=cache,target=/go/pkg/mod \
#    --mount=type=cache,target=/root/.cache/go-build \
#    CGO_ENABLED=0 GOOS=linux go build -o boilerplate_svc ./cmd/rest
#
#
#
##FROM alpine AS alpine
#
##RUN apk add --no-cache ca-certificates
#
#
#
#FROM scratch
#
##COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
#COPY --from=builder /src/app/boilerplate_svc /usr/local/bin/boilerplate_svc
#
#ENTRYPOINT ["/usr/local/bin/boilerplate_svc"]
#
##CMD ["./usr/local/bin/boilerplate_svc"]

FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o binary ./cmd/rest

ENTRYPOINT ["/app/binary"]