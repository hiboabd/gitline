FROM golang:1.16-alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    && update-ca-certificates

EXPOSE 2000

# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY server/go.mod .
COPY server/go.sum .
RUN go mod download

# Copy the code into the container
COPY server/ .

# Build the application
RUN go build -o mock-server .

# Move to /go/bin directory as the place for resulting binary folder
WORKDIR /go/bin

# Copy binary from build to main folder
RUN cp /build/mock-server  .

# Build a small image
FROM scratch

COPY --from=builder /go/bin/mock-server /
COPY server/response.json /
# Command to run
ENTRYPOINT ["/mock-server"]