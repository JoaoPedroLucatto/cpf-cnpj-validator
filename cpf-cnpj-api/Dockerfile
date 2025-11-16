FROM golang:1.24-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main /app/cmd/main.go

FROM gcr.io/distroless/static-debian11
COPY --from=builder /app/main /main
CMD ["/main"]
