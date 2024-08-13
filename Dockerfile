FROM golang:1.23 as builder

# Create and change to the app directory.
WORKDIR /app

COPY . .

# Build the command inside the container.
RUN CGO_ENABLED=0 GO111MODULE=on GOOS=linux go build -v -o main cmd/geo-api/main.go

# Use a Docker multi-stage build to create a lean production image.
FROM alpine
RUN apk add --no-cache ca-certificates

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/main /main

# Run the web service on container startup.
CMD ["/main"]
