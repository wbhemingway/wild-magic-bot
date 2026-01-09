# Use the official Golang image to create a build artifact.
# This is known as a multi-stage build.
FROM golang:1.22-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
# CGO_ENABLED=0 is required for a static build that can run in a minimal container
# -o /wms-bot specifies the output file name
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /wms-bot ./cmd/wms-bot

# ---

# Start from a new, minimal image
FROM alpine:latest

# The wms-bot only needs a root certificate bundle to communicate with Discord APIs if making outbound requests.
# It's good practice to include it.
RUN apk --no-cache add ca-certificates

# Copy the pre-built binary from the "builder" stage
COPY --from=builder /wms-bot /wms-bot

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
# The executable is run directly, not through a shell
ENTRYPOINT ["/wms-bot"]
