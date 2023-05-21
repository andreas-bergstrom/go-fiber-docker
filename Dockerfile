FROM golang:1.20-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o hello_world .

FROM scratch

COPY --from=builder ["/build/hello_world", "/"]

ENTRYPOINT ["/hello_world"]