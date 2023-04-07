# Use an official Golang runtime as a parent image
FROM golang:1.16.3-alpine3.13

RUN go version
ENV GOPATH=/

COPY ./ ./

# install psql
RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres.sh executable
RUN chmod +x wait-for-postgres.sh

# build go app
RUN go mod download
RUN go build -o apiserver ./cmd/apiserver/main.go

CMD ["./apiserver"]