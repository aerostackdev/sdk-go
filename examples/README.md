# Aerostack Go SDK Examples

Examples for `github.com/aerostackdev/sdks/packages/go`.

## Prerequisites

```bash
go get github.com/aerostackdev/sdks/packages/go
```

## Available Examples

| Example | Description | Pattern |
|---------|-------------|---------|
| [**HTTP Server**](./http_server.go) | Usage with `net/http` stdlib. | Standard Library |
| [**Gin Middleware**](./gin_middleware.go) | Integration with Gin Web Framework. | Middleware / Context Injection |

## Usage

### Running Examples

1. Initialize a Go module if needed:
   ```bash
   go mod init example
   go get github.com/aerostackdev/sdks/packages/go
   go get github.com/gin-gonic/gin 
   ```

2. Run example:
   ```bash
   go run examples/http_server.go
   ```
