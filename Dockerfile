FROM golang:1.20

RUN mkdir /app
COPY . /app
WORKDIR /app

EXPOSE 443

RUN go build -o main .
CMD ["/app/main"]