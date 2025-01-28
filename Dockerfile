FROM golang:1.23.4-bullseye AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /main .

FROM alpine:3.21 AS runner

WORKDIR /app

COPY --from=builder /main /app/main
COPY --from=builder /app/assets/static /app/assets/static/
COPY --from=builder /app/templates /app/templates/

USER 1000

EXPOSE 8080
CMD ["./main"]