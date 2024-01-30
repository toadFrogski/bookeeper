FROM golang:1.21.3-bookworm as builder

WORKDIR /app

RUN apt-get update && apt-get install -y git && \
    go install github.com/swaggo/swag/cmd/swag@latest

ENTRYPOINT [ "/go/bin/swag" ]
CMD [ "init" ]