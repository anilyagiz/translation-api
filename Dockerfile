# Build stage
FROM golang:1.18-alpine AS build

WORKDIR /app
COPY . .

RUN go build -o main .

# Runtime stage
FROM alpine:latest

WORKDIR /root/
COPY --from=build /app/main .

EXPOSE 8080

CMD ["./main"]
