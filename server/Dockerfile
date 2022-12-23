FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download && go mod verify
RUN go build -o main ./main.go

FROM alpine:latest AS runner

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]