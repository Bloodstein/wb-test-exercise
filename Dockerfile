FROM golang:1.17.5 AS builder

RUN go version
RUN apt-get install git

COPY ./ /tg-offc-mgr
WORKDIR /tg-offc-mgr

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./bin/app ./cmd/main.go

# легковесный контейнер linux с бинарником
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /tg-offc-mgr/bin/app .
COPY --from=0 /tg-offc-mgr/config ./config
COPY --from=0 /tg-offc-mgr/.env .

EXPOSE 8000

CMD ["./app"]