FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY server.go .

COPY ./content ./content

RUN go build -o server .

FROM scratch

WORKDIR /app

USER 1001:1001

COPY --from=builder /app/server .
COPY --from=builder /app/content ./content

EXPOSE 8080

CMD ["./server"]