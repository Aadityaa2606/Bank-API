# builder stage
FROM golang:alpine AS builder
WORKDIR /app
COPY . .

# build the app
RUN go build -o main main.go

# Install gooose migration tool
RUN apk add curl
RUN curl -L https://raw.githubusercontent.com/pressly/goose/master/install.sh | \
 GOOSE_INSTALL=/app sh -s v3.24.1

# binary app stage
FROM alpine
WORKDIR /app

# Copy the goose binary
COPY --from=builder /app/bin/goose /usr/local/bin/goose

# Copy application binary
COPY --from=builder /app/main .

# Copy other things
COPY .env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

# Make sh's executable
RUN chmod +x /app/start.sh /app/wait-for.sh

EXPOSE 8080
CMD ["/app/main"]
ENTRYPOINT [ "/app/start.sh" ]