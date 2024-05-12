# stage: build go binary
FROM golang:1.22 AS build
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o binary ./cmd/rest

# stage: generate ca-certificates
FROM alpine AS alpine
RUN apk add --no-cache ca-certificates

# stage: final image
FROM scratch
WORKDIR /app
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /app/binary /app/binary
#COPY --from=build /app/.env /app/.env

ENTRYPOINT ["/app/binary"]