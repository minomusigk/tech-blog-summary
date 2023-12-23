FROM golang:1.21.5-alpine AS local

WORKDIR /go/src/app
RUN apk add --no-cache --update-cache \
    alpine-sdk
RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download && go mod verify

CMD ["air", "-c", "./build/.air.toml"]