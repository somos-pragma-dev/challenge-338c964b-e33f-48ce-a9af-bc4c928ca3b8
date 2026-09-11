FROM golang:1.22 AS builder
WORKDIR /app
COPY..
RUN go build -o payment./src/payment

FROM alpine:latest
COPY --from=builder /app/payment /usr/local/bin/payment
EXPOSE 50051
CMD ["payment"]