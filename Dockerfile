# Standalone build: docker build -t gyro . && docker run --rm -e GYRO_TOKEN gyro server list
FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ARG VERSION=dev
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w -X main.version=${VERSION}" -o /out/gyro ./cmd/gyro

FROM alpine:3.22
RUN apk add --no-cache ca-certificates openssh-client tzdata && adduser -D -u 10001 gyro
COPY --from=build /out/gyro /usr/local/bin/gyro
USER gyro
ENTRYPOINT ["/usr/local/bin/gyro"]
