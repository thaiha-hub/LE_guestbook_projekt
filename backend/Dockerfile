# Using Go lang base image -
FROM registry.access.redhat.com/ubi10/go-toolset:10.0 AS builder

USER 0
WORKDIR /app

# Copy go mod files first (better caching)
COPY go.mod ./
RUN go mod tidy

# Copy source code
COPY . .

# Build the binary
RUN go build -o guestbook-api .

# ---------- Stage 2: Runtime ----------
FROM registry.access.redhat.com/ubi10-minimal:10.0

WORKDIR /app

# Copy the binary from builder
COPY --from=builder /app/guestbook-api .

# Expose application port
EXPOSE 8080

# Run the application
CMD ["./guestbook-api"]

