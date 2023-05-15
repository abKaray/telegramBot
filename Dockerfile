# Use an official Golang runtime as a parent image
FROM golang:1.19.4-alpine3.17

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the binary executable
RUN go build -o main ./cmd/main.go

# Expose port 8080 for the container
EXPOSE 8080

# Run the command to start the server
CMD ["./main"]