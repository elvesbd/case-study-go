FROM golang:1.15.3-alpine3.12

EXPOSE 9000

RUN apk update \
  && apk add --no-cache \
  mysql-client \
  build-base

RUN mkdir /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
COPY ./grpc_entryponit.sh /usr/local/bin/grpc_entryponit.sh
RUN /bin/chmod +x /usr/local/bin/grpc_entryponit.sh

RUN go build cmd/main.go
RUN mv main /usr/local/bin/

CMD ["main"]
ENTRYPOINT ["grpc_entryponit.sh"]