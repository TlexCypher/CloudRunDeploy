FROM golang:1.23

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app main.go

EXPOSE 8080

CMD ["./app"]
