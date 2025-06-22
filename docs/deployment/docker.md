# Docker

Containerization and deployment strategies for your CLI application.

## Docker Support

The CLI template includes Docker support for containerized deployments and distribution.

## Dockerfile

### Current Configuration

The project includes a minimal Dockerfile:

```dockerfile
FROM scratch

COPY cli-template /usr/local/bin/cli-template

ENTRYPOINT ["/usr/local/bin/cli-template"]
```

This creates a minimal container with just the static binary.

### Multi-stage Build

For a complete build process, use a multi-stage Dockerfile:

```dockerfile
# Build stage
FROM golang:1.23-alpine AS builder

# Install build dependencies
RUN apk add --no-cache git ca-certificates

# Set working directory
WORKDIR /app

# Copy go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build \
    -ldflags "-s -w -X main.version=${VERSION} -X main.commit=${COMMIT}" \
    -o cli-template ./main.go

# Final stage
FROM scratch

# Copy CA certificates for HTTPS
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary
COPY --from=builder /app/cli-template /usr/local/bin/cli-template

# Set entrypoint
ENTRYPOINT ["/usr/local/bin/cli-template"]
```

## Building Images

### Local Build

```bash
# Basic build
docker build -t cli-template:latest .

# Build with version
docker build --build-arg VERSION=v1.0.0 -t cli-template:v1.0.0 .

# Multi-platform build
docker buildx build --platform linux/amd64,linux/arm64 -t cli-template:latest .
```

### Development Build

```dockerfile
# Dockerfile.dev
FROM golang:1.23-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Install development tools
RUN go install github.com/cosmtrek/air@latest

CMD ["air"]
```

```bash
# Development container
docker build -f Dockerfile.dev -t cli-template:dev .
docker run -v $(pwd):/app cli-template:dev
```

## Container Registry

### GitHub Container Registry

Images are automatically pushed to GHCR via GoReleaser:

```yaml
dockers:
  - image_templates:
      - "ghcr.io/mavogel/cli-template:{{ .Tag }}"
      - "ghcr.io/mavogel/cli-template:latest"
    dockerfile: Dockerfile
```

### Usage

```bash
# Pull from registry
docker pull ghcr.io/mavogel/cli-template:latest

# Run container
docker run --rm ghcr.io/mavogel/cli-template:latest --help

# Interactive mode
docker run -it --rm ghcr.io/mavogel/cli-template:latest
```

## Docker Compose

### Basic Compose

```yaml
# docker-compose.yml
version: '3.8'

services:
  cli-template:
    image: ghcr.io/mavogel/cli-template:latest
    command: ["--help"]
    
  cli-template-dev:
    build:
      context: .
      dockerfile: Dockerfile.dev
    volumes:
      - .:/app
      - go-cache:/go/pkg/mod
    working_dir: /app

volumes:
  go-cache:
```

### With Configuration

```yaml
version: '3.8'

services:
  cli-template:
    image: ghcr.io/mavogel/cli-template:latest
    volumes:
      - ./config:/config:ro
      - ./data:/data
    environment:
      - CONFIG_PATH=/config/app.yaml
      - LOG_LEVEL=debug
    command: ["hello", "--name", "Docker"]
```

## Kubernetes Deployment

### Basic Deployment

```yaml
# k8s/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: cli-template
spec:
  replicas: 1
  selector:
    matchLabels:
      app: cli-template
  template:
    metadata:
      labels:
        app: cli-template
    spec:
      containers:
      - name: cli-template
        image: ghcr.io/mavogel/cli-template:latest
        command: ["cli-template"]
        args: ["--help"]
        resources:
          requests:
            memory: "64Mi"
            cpu: "250m"
          limits:
            memory: "128Mi"
            cpu: "500m"
```

### ConfigMap

```yaml
# k8s/configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: cli-template-config
data:
  app.yaml: |
    log_level: info
    output_format: json
```

### Job

```yaml
# k8s/job.yaml
apiVersion: batch/v1
kind: Job
metadata:
  name: cli-template-job
spec:
  template:
    spec:
      containers:
      - name: cli-template
        image: ghcr.io/mavogel/cli-template:latest
        command: ["cli-template"]
        args: ["hello", "--name", "Kubernetes"]
      restartPolicy: Never
  backoffLimit: 4
```

## Security Considerations

### Image Scanning

```bash
# Scan with Docker Scout
docker scout cves ghcr.io/mavogel/cli-template:latest

# Scan with Trivy
trivy image ghcr.io/mavogel/cli-template:latest
```

### Distroless Images

Use distroless for better security:

```dockerfile
FROM gcr.io/distroless/static:nonroot

COPY cli-template /usr/local/bin/cli-template
USER nonroot:nonroot

ENTRYPOINT ["/usr/local/bin/cli-template"]
```

### Minimal Base Images

```dockerfile
# Alpine-based
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY cli-template /usr/local/bin/cli-template
ENTRYPOINT ["/usr/local/bin/cli-template"]

# Scratch (minimal)
FROM scratch
COPY ca-certificates.crt /etc/ssl/certs/
COPY cli-template /usr/local/bin/cli-template
ENTRYPOINT ["/usr/local/bin/cli-template"]
```

## Best Practices

### Image Optimization

- **Multi-stage builds**: Reduce final image size
- **Static linking**: Avoid runtime dependencies
- **Minimal base**: Use scratch or distroless
- **Layer caching**: Optimize Dockerfile order

```dockerfile
# Good layer order
COPY go.mod go.sum ./    # Changes rarely
RUN go mod download      # Cache dependencies
COPY . .                 # Changes frequently
RUN go build ...         # Build step
```

### Security

- **Non-root user**: Run as non-privileged user
- **Read-only filesystem**: Use read-only containers
- **Secrets management**: Use proper secret handling
- **Regular updates**: Keep base images updated

```dockerfile
FROM alpine:latest
RUN adduser -D -s /bin/sh appuser
USER appuser
COPY --chown=appuser:appuser cli-template /usr/local/bin/
```

### Configuration

- **Environment variables**: Use for configuration
- **Volume mounts**: For persistent data
- **Health checks**: Monitor container health
- **Resource limits**: Set appropriate limits

```dockerfile
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD ["/usr/local/bin/cli-template", "health"] || exit 1
```

## Troubleshooting

### Common Issues

**Binary not found**:
```dockerfile
# Ensure correct path
COPY cli-template /usr/local/bin/cli-template
RUN chmod +x /usr/local/bin/cli-template
```

**Permission denied**:
```dockerfile
# Set executable permissions
RUN chmod +x /usr/local/bin/cli-template
```

**SSL certificate errors**:
```dockerfile
# Copy CA certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
```

### Debugging

```bash
# Interactive shell
docker run -it --rm --entrypoint /bin/sh alpine:latest

# Check binary
docker run --rm ghcr.io/mavogel/cli-template:latest --version

# Mount for debugging
docker run -it --rm -v $(pwd):/debug ghcr.io/mavogel/cli-template:latest
```

### Build Issues

```bash
# Clear build cache
docker system prune -a

# Build with no cache
docker build --no-cache -t cli-template:latest .

# Check build context
docker build --progress=plain -t cli-template:latest .
```

## Advanced Usage

### Init Containers

```yaml
apiVersion: v1
kind: Pod
spec:
  initContainers:
  - name: setup
    image: ghcr.io/mavogel/cli-template:latest
    command: ['cli-template', 'setup']
  containers:
  - name: main
    image: ghcr.io/mavogel/cli-template:latest
    command: ['cli-template', 'run']
```

### Sidecar Pattern

```yaml
containers:
- name: main-app
  image: main-application:latest
- name: cli-sidecar
  image: ghcr.io/mavogel/cli-template:latest
  command: ['cli-template', 'monitor']
```

### Batch Processing

```bash
# Process multiple files
docker run --rm -v $(pwd)/data:/data \
  ghcr.io/mavogel/cli-template:latest \
  process --input /data --output /data/results
```