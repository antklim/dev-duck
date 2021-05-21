# 1. Build
FROM golang:alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o build/ ./...

# 2. Run
FROM gcr.io/distroless/base AS devduck
COPY --from=builder /build/build/devduck /
ENTRYPOINT ["/devduck"]

FROM gcr.io/distroless/base AS devduckauth
COPY --from=builder /build/build/devduckauth /
ENTRYPOINT ["/devduckauth"]