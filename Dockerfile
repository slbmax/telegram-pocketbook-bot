FROM golang:1.18-alpine as buildbase

COPY . /github.com/slbmax/telegram-pocketbook-bot/
WORKDIR /github.com/slbmax/telegram-pocketbook-bot/

RUN GOOS=linux go build -o ./.bin/bot ./main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/slbmax/telegram-pocketbook-bot/.bin/bot .

CMD ["./bot"]
