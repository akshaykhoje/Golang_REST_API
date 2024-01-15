# Use Alpine Linux as the base image
FROM alpine:3.18

# Set shell to sh
SHELL ["/bin/sh", "-c"]

# Create a non-root user
RUN adduser -D -u 1000 definitely_not_root

# Install necessary dependencies
RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
        curl

# Download and install Go 1.21.6
RUN curl -sSL https://dl.google.com/go/go1.21.6.linux-amd64.tar.gz | tar -C /usr/local -xz

# Set Go environment variables
ENV GOPATH=/go
ENV PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
ENV API_URL=http://127.0.0.1/

# Create a working directory
WORKDIR /api

# Copy the necessary files
COPY . .

# Download Go dependencies mentiond in the go.mod file
RUN go mod download

# Switch to the non-root user
USER definitely_not_root

# Expose port 8080 of container
EXPOSE 80

# Command to start the session with the shell
CMD ["sh"]
