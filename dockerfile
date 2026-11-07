# Первый этап
FROM golang:1.27.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o aurelia-reads main.go

# Второй этап
FROM alpine:3.24.3

WORKDIR /app

COPY --from=builder /app/aurelia-reads ./aurelia-reads

CMD ["./aurelia-reads"]