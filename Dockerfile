# Use the official Golang image to create a build artifact.
FROM golang:1.21.1-alpine as builder

# Copy local code to the container image.
WORKDIR /app
COPY . ./

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server cmd/server/main.go

# Use the official Alpine image for a lean production container.
FROM alpine:3.14

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /server

# Run the web service on container startup.
CMD ["/server"]