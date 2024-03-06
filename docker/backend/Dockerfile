FROM golang:1.21.3-bookworm as builder
WORKDIR /build
COPY . .

RUN apt-get update && apt-get install -y git

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download && \
    go install github.com/google/wire/cmd/wire@latest && \
    wire gen ./modules/*

# Build
WORKDIR /build/cmd/web
RUN CGO_ENABLED=0 go build -o app-binary .

# Create final image
FROM alpine as final
WORKDIR /app
COPY --from=builder /build/cmd/web/app-binary .
ENTRYPOINT ["/app/app-binary"]