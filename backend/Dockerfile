# Backend Containerfile 
# This builds and runs the guestbook API

# Stage 1: The Builder
# ======================================
# Uses the Red Hat Go Toolset to compile the app
FROM registry.access.redhat.com/ubi10/go-toolset:10.0 AS builder

# Set the working dir (standard for UBI images)
WORKDIR /opt/app-root/src

# Copy backend source code  to the container (main.go, and go.mod)
# We use --chown to set the file owner to user 1001, non-root user  and group 0, root group
COPY --chown=1001:0 . .

# Run the dependencies check as requested
RUN go mod tidy

# Build the application
# We output the binary named 'guestbook-api'
RUN go build -o guestbook-api .

# Stage 2: The Runtime
# Uses the Minimal UBI for a smaller, secure final image
# ==========================================
FROM registry.access.redhat.com/ubi10-minimal:10.0

# Set working directory for the runtime
WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /opt/app-root/src/guestbook-api .

# Set default Environment Variables
# These can be overridden by GitHub Actions or Kubernetes later
ENV PORT=8080 \
    DB_HOST=localhost \
    DB_PORT=5432 \
    DB_USER=guestbook \
    DB_PASSWORD=password \
    DB_NAME=guestbook \
    REDIS_HOST=localhost \
    REDIS_PORT=6379 \
    REDIS_PASSWORD=""

# Expose the port the app listens on
EXPOSE 8080

# Command to run the application
CMD ["./guestbook-api"]
