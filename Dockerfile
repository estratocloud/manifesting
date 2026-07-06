ARG GO_VERSION=1.24
FROM golang:${GO_VERSION}-alpine

RUN go install golang.org/x/tools/cmd/goimports@v0.24

WORKDIR /app

ENTRYPOINT ["sh"]
