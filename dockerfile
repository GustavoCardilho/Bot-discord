FROM golang:latest

WORKDIR /app

# COPY go.mod and go.sum files to cache dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the application
RUN go build -o main ./infra

# Run the application
CMD ["./main"]