FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY server.go .

RUN go build -o server .

#FROM scratch
#
#WORKDIR /app
#
#COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]