FROM golang:1.22.6-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o rssagg

EXPOSE 8080

CMD ["./rssagg"]