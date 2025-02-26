FROM golang:alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://raw.githubusercontent.com/pressly/goose/master/install.sh | \
 GOOSE_INSTALL=/app sh -s v3.24.1

# Run stage
FROM alpine
WORKDIR /app
# Copy the goose binary
COPY --from=builder /app/bin/goose /usr/local/bin/goose
# Copy your application binary
COPY --from=builder /app/main .
COPY .env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration
RUN chmod +x /app/start.sh /app/wait-for.sh
# Verify Goose installation
RUN goose --version
EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]