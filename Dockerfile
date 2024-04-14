# Dockerfile

FROM golang:1.22.1 AS build

WORKDIR /app

COPY . .

RUN go build -o vanity

CMD ["./vanity"]
