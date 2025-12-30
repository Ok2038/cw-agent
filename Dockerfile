# GoReleaser Dockerfile - uses pre-built binary
FROM gcr.io/distroless/static-debian12:nonroot

# Copy the pre-built binary from GoReleaser
COPY cw-agent /cw-agent

# Set user
USER nonroot:nonroot

# Set entrypoint
ENTRYPOINT ["/cw-agent"]

# Default command
CMD ["start", "-c", "/etc/certwatch/certwatch.yaml"]
