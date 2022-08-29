FROM golang:1.18-alpine as buildbase

COPY . /github.com/slbmax/telegram-pocketbook-bot/
WORKDIR /github.com/slbmax/telegram-pocketbook-bot/

RUN GOOS=linux go build -o /usr/local/bin/telegram-pocketbook-bot ./main.go

ENTRYPOINT ["telegram-pocketbook-bot"]
