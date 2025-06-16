FROM golang:1.24-alpine AS builder

WORKDIR /app

RUN apk add --no-cache git build-base

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o users main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/users .

# Needed for SQLite
RUN apk add --no-cache sqlite-libs

EXPOSE 8082

CMD ["./users"]
