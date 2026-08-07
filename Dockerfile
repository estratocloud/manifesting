ARG GO_VERSION=1.24
FROM golang:${GO_VERSION}-alpine AS dev

RUN go install golang.org/x/tools/cmd/goimports@v0.24
RUN wget -O- -nv https://golangci-lint.run/install.sh | sh -s v2.12.2

WORKDIR /app

ENTRYPOINT ["sh"]


FROM scratch AS production
ARG BIN=manifesting
COPY --chmod=755 ${BIN} /manifesting
WORKDIR /app
ENTRYPOINT ["/manifesting"]
