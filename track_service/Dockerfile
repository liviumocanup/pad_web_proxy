FROM golang:1.21

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o main ./main.go

EXPOSE 50052

CMD ["./main"]
