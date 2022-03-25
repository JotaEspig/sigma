FROM golang:latest

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

ENV PORT 8080

EXPOSE 8080

RUN go build -o sigma main.go

CMD ["./sigma"]