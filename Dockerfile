FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o apiserver ./cmd/apiserver/main.go

RUN apk add --no-cache postgresql-client
RUN wget -q https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz && \
    tar -zxvf migrate.linux-amd64.tar.gz && \
    mv migrate.linux-amd64 /usr/bin/migrate && \
    rm -f migrate.linux-amd64.tar.gz

EXPOSE 8080

CMD ["sh", "-c", "migrate -path migrations -database postgres://$DB_USER:$DB_PASSWORD@$DB_HOST:$DB_PORT/$DB_NAME?sslmode=disable up && ./apiserver"]

