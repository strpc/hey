ARG GO_VERSION=1.19
FROM golang:${GO_VERSION}-alpine as builder

RUN apk add --no-cache --update git && apk add build-base

WORKDIR /build
COPY ./go.mod ./go.sum ./
RUN go mod download -x

COPY ./ .

RUN CGO_ENABLED=0 GOOS=linux go build -o ./hey


FROM alpine:3.17

RUN adduser -u 1000 -h /app -D -g "" user  \
    && chown -hR user: /app

WORKDIR /hey

ENV PATH=/hey:$PATH
COPY --from=builder --chown=user:user /build/hey .

USER user

CMD ["./hey"]
