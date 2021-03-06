FROM golang:1.16-alpine as base
WORKDIR /app

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    && update-ca-certificates

FROM node:16.13.0-alpine3.14 as asset-env

WORKDIR /app

COPY web/assets web/assets
COPY webpack.config.js .
COPY package.json .
COPY yarn.lock .

RUN yarn install
RUN yarn build

### Development with hot reload and debugger
FROM base AS dev
WORKDIR /app

COPY --from=asset-env /app/web/static web/static
COPY --from=asset-env /app/web/assets/images web/static/images

RUN go get -u github.com/cosmtrek/air && go install github.com/go-delve/delve/cmd/dlv@latest
EXPOSE 3001
EXPOSE 2346

ENTRYPOINT ["air"]

FROM base as build-env

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -a -installsuffix cgo -o /go/bin/gitline

FROM alpine:3.14

WORKDIR /go/bin

RUN apk --update --no-cache add \
    ca-certificates \
    && rm -rf /var/cache/apk/*
RUN apk --no-cache add tzdata

COPY --from=build-env /go/bin/gitline gitline
COPY --from=build-env /app/web/templates web/templates
COPY --from=asset-env /app/web/static web/static
ENTRYPOINT ["./gitline"]
CMD ["./gitline"]