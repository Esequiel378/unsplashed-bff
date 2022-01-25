FROM golang:1.17-alpine AS builder

LABEL maintainer="Esequiel <esequielalbornoz7@gmail.com>"

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o unsplashed

FROM alpine

RUN apk --no-cache add ca-certificates

WORKDIR /

COPY .env ./
COPY --from=builder /app/unsplashed /unsplashed

EXPOSE 80

ENTRYPOINT ["/unsplashed"]
