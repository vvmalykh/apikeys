# Use a Golang base image
FROM golang:1.21.0 as builder

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /app

# Initialize a new module and download dependencies
RUN go mod init helloworld && \
    go mod tidy

# Copy only the necessary source code files
COPY . .

# Build the application
RUN CGO_ENABLED=0 go build -o /bin/api-key-service

# Use a minimal image to run the service
FROM alpine:latest

# Copy the compiled service from the build stage
COPY --from=builder /bin/api-key-service /bin/api-key-service

# Run the service
CMD ["/bin/api-key-service"]