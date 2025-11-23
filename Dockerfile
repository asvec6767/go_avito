FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

EXPOSE $PORT

RUN go build -o go-avito cmd/go_avito_api/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app .

CMD [ "./go-avito" ]