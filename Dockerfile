# Step 1: Build the Go app
FROM golang:1.23.2 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build cmd/main.go

#COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]

