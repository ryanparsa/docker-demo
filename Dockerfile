FROM golang:latest as builder
ENV CGO_ENABLED=0
WORKDIR /app

COPY . .

RUN go build -o main .

# Use a lightweight image for the final stage
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
