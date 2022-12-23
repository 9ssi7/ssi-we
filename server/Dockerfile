FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download
COPY ./ ./
RUN go build ./main.go

FROM alpine:latest AS runner

ARG PORT=8080

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE $PORT

CMD ["./main"]