# Step 1: Build the Go app
FROM golang:1.23.2 AS builder
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build cmd/main.go

ENV PORT=8888
ENV POSTGRES_URI="postgres://postgres:password@postgres-container:5432/postgres?sslmode=disable"


#COPY --from=builder /app/main .
EXPOSE 8888
CMD ["./main"]

