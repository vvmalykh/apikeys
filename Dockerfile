# Use a Golang base image
FROM golang:1.22.5 as builder

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy only the necessary source code files
COPY . .

# Migrations
RUN wget -q -O migrate.linux-amd64.tar.gz https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz && \
    tar xvzf migrate.linux-amd64.tar.gz && \
    mv migrate.linux-amd64 /usr/local/bin/migrate && \
    ls -al /usr/local/bin

# Build the application
RUN CGO_ENABLED=0 go build -o /bin/api-key-service ./cmd/apikeys

# Use a minimal image to run the service
FROM alpine:latest

# Copy the compiled service from the build stage
COPY --from=builder /bin/api-key-service /bin/api-key-service
# Copy the migrate executable
COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate

# Run the service
CMD ["/bin/api-key-service"]
