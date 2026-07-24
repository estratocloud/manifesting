ARG GO_VERSION=1.24
FROM golang:${GO_VERSION}-alpine AS dev

RUN go install golang.org/x/tools/cmd/goimports@v0.24

WORKDIR /app

ENTRYPOINT ["sh"]


FROM scratch AS production
ARG BIN=manifesting
COPY --chmod=755 ${BIN} /manifesting
WORKDIR /app
ENTRYPOINT ["/manifesting"]
