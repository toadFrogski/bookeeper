FROM golang:alpine

WORKDIR /work

RUN apk add curl docker && \
  curl -sSf https://atlasgo.sh | sh -s -- -y

ENTRYPOINT [ "atlas" ]