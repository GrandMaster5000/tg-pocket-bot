FROM golang:1.18.3-alpine3.16 AS builder

COPY . /github.com/GrandMaster5000/tg-bot-pocket/
WORKDIR /github.com/GrandMaster5000/tg-bot-pocket/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/GrandMaster5000/tg-bot-pocket/bin/bot .
COPY --from=0 /github.com/GrandMaster5000/tg-bot-pocket/configs configs/

EXPOSE 80

CMD ["./bot"]