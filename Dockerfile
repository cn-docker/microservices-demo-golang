###############################################################################
# Build Go Binary
###############################################################################
FROM golang:1.20-alpine3.18 AS GO_BUILD
LABEL maintainer="Julian Nonino <noninojulian@gmail.com"

# Set the Current Working Directory inside the container
WORKDIR /app

# We copy everything in the root to the working directory in the container
COPY . .

# Build the Go app
RUN go build -o microservices-demo

###############################################################################
# Backend Docker Image
###############################################################################
FROM alpine:3.18
LABEL maintainer="Julian Nonino <noninojulian@gmail.com"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the binary from the build stage
COPY --from=GO_BUILD /app/microservices-demo .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by the build stage
CMD ["/app/microservices-demo"]
