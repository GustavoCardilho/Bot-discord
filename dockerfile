FROM golang:latest AS builder

WORKDIR /app

# COPY go.mod and go.sum files to cache dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Remove the existing 'src' directory
RUN rm -rf ./src

# Build the application
RUN go build -o main ./src

# Use a smaller base image for the final image
FROM scratch

WORKDIR /app

# Copy only the built binary from the builder stage
COPY --from=builder /app/main .

# Run the application
CMD ["./main"]