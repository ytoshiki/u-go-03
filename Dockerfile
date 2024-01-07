FROM golang:1.21-alpine

WORKDIR /app

COPY . .

RUN go build -o /u-go-03

CMD ["/u-go-03"]