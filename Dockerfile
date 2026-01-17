FROM golang:1.25.5

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin/telegram-alert

ENV PORT=8080

EXPOSE ${PORT}

CMD ["/app/bin/telegram-alert"]