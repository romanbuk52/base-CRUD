#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/crud-server
COPY . .
RUN go get -d -v ./...
RUN apk add --update gcc musl-dev
RUN go build -o /bin ./...

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /go/src/crud-server
COPY --from=builder /bin/main ./bin/main
LABEL Name=crudserver Version=0.0.1
EXPOSE 8081
ENTRYPOINT ./bin/main

