FROM alpine:latest AS builder

# Install Go from Alpine package + build tools
RUN apk add --no-cache \
    bash \
    build-base \
    ca-certificates \
    git \
    go

# Set Go version
ENV GO_VERSION=1.24.2

# Clone Go source
RUN git clone https://go.googlesource.com/go /root/src/go.googlesource.com/go \
    && cd /root/src/go.googlesource.com/go \
    && git checkout go${GO_VERSION}

# Build Go from source using Alpine's Go as bootstrap
WORKDIR /root/src/go.googlesource.com/go/src
RUN ./make.bash

# Set newly built Go as the default
ENV GOROOT=/root/src/go.googlesource.com/go
ENV PATH=$GOROOT/bin:$PATH

# Build app
WORKDIR /root/schemacafe
COPY . .
RUN go mod tidy
RUN go build -o /bin/server cmd/server/main.go

# Copy to fresh image
FROM scratch
COPY --from=builder /bin/server /bin/server
ENTRYPOINT ["/bin/server"]
