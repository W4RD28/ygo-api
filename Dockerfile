FROM golang:1.22.4-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /ygo-api cmd/ygo-api/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /ygo-api .

EXPOSE 8080

CMD ["./ygo-api"]