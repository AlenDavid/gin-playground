FROM golang:1.19-alpine

WORKDIR /var/run/app

COPY . .

RUN go install cmd/server/server.go

CMD ["server"]
