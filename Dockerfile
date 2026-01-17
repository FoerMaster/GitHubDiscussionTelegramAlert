FROM golang:1.25.5

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/telegram-alert

EXPOSE 8080

CMD ["/app/bin/telegram-alert"]