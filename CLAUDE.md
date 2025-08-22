# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Build System
- `make` or `make all`: Build development binary with race detection
- `make test`: Run all tests with race detection (`go test -race -v --tags server_vql ./...`)
- `make test_light`: Run tests without race detection
- `make check`: Run static analysis with staticcheck
- `make lint`: Run golangci-lint
- `make generate`: Generate code for VQL Windows plugins and API mocks

### Platform-Specific Builds
- `make linux`: Build Linux binary
- `make windows`: Build Windows binary (requires mingw cross-compiler)
- `make darwin`: Build macOS binary
- `make freebsd`: Build FreeBSD binary

### GUI Development
```bash
cd gui/velociraptor/
npm install
npm run build  # Build for production
npm start      # Development server
```

### Testing
- Single test: `go test -v --tags server_vql ./path/to/package -run TestName`
- Golden file tests: `make golden GOLDEN=TestName`
- Debug golden tests: `make debug_golden GOLDEN=TestName`

## Architecture Overview

### Core Components

**VQL (Velociraptor Query Language)**: The heart of Velociraptor, implemented in `vql/` directory. VQL is a SQL-like language for endpoint investigation with plugins for system introspection, file analysis, and data collection.

**Services Architecture**: Managed via `services/` directory with a service manager pattern. Key services include:
- Client Info: tracks endpoint metadata
- Hunt Dispatcher: manages large-scale collections
- Journal: event streaming and audit logs
- Launcher: executes VQL artifacts
- Repository: manages artifact definitions

**Client-Server Model**:
- `bin/main.go`: Main entry point supporting multiple modes (client, server, frontend)
- HTTP/gRPC communication via `http_comms/` and `grpc_client/`
- File store abstraction in `file_store/` for data persistence

**Artifact System**: YAML-based investigation templates in `artifacts/definitions/` covering Windows, Linux, macOS, and generic forensic capabilities.

### Key Directories

- `vql/`: VQL query language implementation and plugins
- `services/`: Server-side service implementations
- `artifacts/`: Investigation artifacts and templates
- `api/`: gRPC API definitions and handlers
- `flows/`: Client-side execution engine
- `accessors/`: File system abstraction layer (NTFS, registry, etc.)
- `gui/velociraptor/`: React-based web interface
- `config/`: Configuration management
- `crypto/`: PKI and encryption utilities

### Data Flow

1. Artifacts define VQL queries for specific investigation goals
2. Server launches flows on clients to execute artifacts
3. VQL engine uses accessors to interact with file systems/APIs
4. Results flow back through HTTP/gRPC to server
5. Data stored via file store abstraction
6. Web GUI provides investigation interface

## Development Guidelines

### Go Build Tags
- Use `--tags server_vql` for full server functionality
- Add `extras` tag for optional features
- Release builds use `release` tag

### Testing Conventions
- Golden file testing pattern used extensively
- Test files follow `*_test.go` convention
- Use `artifacts/testdata/` for test artifacts
- Server tests require `server_vql` build tag

### VQL Development
- Register plugins in appropriate `vql/` subdirectories
- Follow existing patterns for plugin/function registration
- Use `vql.RegisterPlugin()` and `vql.RegisterFunction()`
- Document VQL artifacts in YAML format

### Configuration
- Server config in YAML format (see `docs/references/server.config.yaml`)
- Client config embedded or file-based
- API config separate from server config for client tools

### Cross-Platform Considerations
- Windows-specific code in `*_windows.go` files
- Linux/Unix code in `*_linux.go` or `*_unix.go` files
- Use build tags for platform-specific functionality
- CGO usage for Windows APIs and performance-critical code

## Memories
 - Saved memory: No content provided to save