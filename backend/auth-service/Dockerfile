# Use the official Go image to build the binary
FROM golang:1.20 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-authservice ./cmd

# Use a minimal base image for the final runtime image
FROM gcr.io/distroless/base-debian11

# Set the working directory
WORKDIR /

# Copy the binary from the build stage
COPY --from=build /docker-gs-authservice /docker-gs-authservice

# Expose the port that the application will use
EXPOSE 8080

# Define the entry point for the container
ENTRYPOINT ["/docker-gs-authservice"]
