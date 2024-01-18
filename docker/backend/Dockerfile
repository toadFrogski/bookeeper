FROM golang:1.21.3-bookworm as builder
WORKDIR /build
RUN apt-get update && apt-get install -y git

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . ./
RUN go build -o ./app ./cmd/web

# Create final image
FROM alpine
WORKDIR /
COPY --from=builder /build/app .
EXPOSE 8080
CMD ["./app"]