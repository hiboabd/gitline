# multistage build
# build our helloworld binary to run on alpine linux
# copy the compiled binary into an alpine base image
# without the extra golang tools

FROM golang:1.16-alpine AS builder
RUN mkdir /build
ADD go.mod go.sum hello.go /build/
WORKDIR /build
RUN go build

# swithc to different base image
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
# copy the binary from our builder image
COPY --from=builder /build/gitline /app/
COPY views/ /app/views
WORKDIR /app
CMD ["./gitline"]