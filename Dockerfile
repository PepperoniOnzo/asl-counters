# syntax=docker/dockerfile:1

FROM golang:1.22.3

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go mod download

RUN go build -o /asl-counters

EXPOSE 8080

CMD [ "/asl-counters" ]