# Use an official Go runtime as a parent image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy your Go server code
COPY main.go .
COPY go.mod .

# Build the Go application
RUN go build -o app

# Expose the port your Go server is listening on
EXPOSE 3000

# Start your Go server
CMD ["./app"]