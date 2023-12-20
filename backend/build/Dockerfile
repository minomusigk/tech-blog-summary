### Base image and build
FROM golang:1.21.5-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY ./app/ ./app
RUN go build -ldflags "-s -w" -o /tmp/main -a ./app/cmd/server.go

### local開発環境向け
FROM builder as local

COPY --from=builder /tmp/main /app/main
CMD ["/app/main"]