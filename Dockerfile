FROM golang:1.23.3-alpine AS builder

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o /api ./cmd/api/main.go

# Final stage
FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata

RUN adduser -D -g '' appuser

WORKDIR /app

COPY --from=builder /api .

USER appuser

EXPOSE 80

CMD ["./api"]
