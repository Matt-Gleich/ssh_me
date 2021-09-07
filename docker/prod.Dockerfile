FROM golang:1.17.0-alpine3.14 AS builder

# Meta data:
LABEL maintainer="email@mattglei.ch"
LABEL description="👋 Meet me via ssh!"

# Copying over all the files:
COPY . /usr/src/app
WORKDIR /usr/src/app

# Installing dependencies/
RUN go get -v -t -d ./...

# Build the app
RUN go build -o app .

# hadolint ignore=DL3006,DL3007
FROM alpine:latest
WORKDIR /
COPY --from=builder /usr/src/app/app .
CMD ["./app"]
