# Build stage
FROM golang:1.25 AS builder
WORKDIR /app
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -o timeservice main.go

#Final stage
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/timeservice .
EXPOSE 6080
CMD ["./timeservice"]

