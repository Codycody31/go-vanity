# First stage: build the executable with CGO disabled for compatibility with Alpine
FROM golang:1.22.1 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the local code to the container's workspace
COPY . .

# Build the Go app with CGO disabled
RUN CGO_ENABLED=0 go build -o vanity

# Second stage: setup the runtime environment
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary from the build stage to the current stage
COPY --from=build /app/vanity .

# Ensure the binary is executable
RUN chmod +x vanity

# Set the binary as the container's entry point
ENTRYPOINT ["./vanity"]
