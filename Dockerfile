FROM golang:1.17-buster AS builder

LABEL maintainer="Esequiel <esequielalbornoz7@gmail.com>"

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o unsplashed

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY .env ./
COPY --from=builder /app/unsplashed /unsplashed

USER nonroot:nonroot

EXPOSE 80

ENTRYPOINT ["/unsplashed"]
