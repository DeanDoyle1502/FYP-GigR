# Use the official Go image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Set up Go proxy
ENV GOPROXY=https://proxy.golang.org,direct

# Copy the Go module files
COPY go.mod go.sum ./

# Copy the rest of the application files
COPY . ./

# Build the Go application
RUN go build -o app

# Expose the application port
EXPOSE 8080

# Start the application
CMD ["./app"]
