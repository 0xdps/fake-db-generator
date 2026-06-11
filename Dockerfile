FROM golang:1.22-alpine AS builder

WORKDIR /build

# Install build dependencies
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Copy go files
COPY golang/go.mod golang/go.sum ./
RUN go mod download

# Copy source
COPY golang/*.go ./

# Build binary
RUN CGO_ENABLED=1 go build -ldflags="-s -w" -o fakestack .

# Final stage
FROM alpine:latest

RUN apk add --no-cache sqlite-libs

COPY --from=builder /build/fakestack /usr/local/bin/fakestack

WORKDIR /data

ENTRYPOINT ["fakestack"]
CMD ["-h"]
