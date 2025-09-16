# Use lightweight debian os as the base image
FROM debian:stable-slim

# Install CA certificates (needed for SSL connections to Supabase)
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the executable
COPY fooddeliveryapp bin/fooddeliveryapp

# Run the executable
CMD ["/bin/fooddeliveryapp"]