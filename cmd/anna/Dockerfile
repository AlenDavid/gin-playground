FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main ./cmd/anna/anna.go

FROM scratch

WORKDIR /app

COPY --from=builder /app/main /app

CMD [ "./main" ]
