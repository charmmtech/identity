# syntax=docker/dockerfile:1.6

FROM golang:1.25-alpine AS builder

LABEL maintainer="Joseph Akitoye <josephakitoye@gmail.com>"

# Install dependencies
RUN apk add --no-cache git bash build-base ca-certificates

WORKDIR /app

# ✅ Proper Go env for private modules
ENV GO111MODULE=on
ENV GOPRIVATE=buf.build
ENV GONOSUMDB=buf.build
ENV GOPROXY=https://proxy.golang.org,direct

# Copy go mod first (for caching)
COPY go.mod go.sum ./

# ✅ Authenticate to Buf using BuildKit secret
RUN --mount=type=secret,id=buf_token \
    mkdir -p /root && \
    echo "machine buf.build login token password $(cat /run/secrets/buf_token)" > /root/.netrc && \
    go mod download && \
    go mod verify && \
    rm /root/.netrc

# Copy source code
COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main cmd/server/main.go

# --- Final Stage ---
FROM alpine:latest
LABEL com.docker.compose.project=microservice

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
