FROM golang:1.21.3-bookworm as builder
WORKDIR /build

RUN apt-get update && apt-get install -y git

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
COPY . .
WORKDIR /build/cmd/web
RUN CGO_ENABLED=0 go build -o app-binary .

# Create final image
FROM alpine
WORKDIR /app
COPY --from=builder /build/cmd/web/app-binary .
EXPOSE 80
ENTRYPOINT ["/app/app-binary"]