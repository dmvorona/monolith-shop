FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git build-base

COPY go.mod ./
COPY go.sum ./
COPY controllers ./
COPY db ./
COPY routes ./
COPY models ./
RUN go mod download

COPY . ./

RUN go build -o products services/products/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/products .
COPY --from=builder /app/templates ./templates
COPY --from=builder /app/static ./static

# Needed for SQLite
RUN apk add --no-cache sqlite-libs

EXPOSE 8081

CMD ["./products"]
