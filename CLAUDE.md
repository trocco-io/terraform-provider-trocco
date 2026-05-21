# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

This repository provides a Terraform provider for [TROCCO](https://trocco.io), a cloud ETL service. The provider is built using the Terraform Plugin Framework and allows managing TROCCO resources via their API (available for paid plans only).

## Architecture

The codebase follows a clean separation of concerns:

- **API Client Layer** (`internal/client/`): Handles communication with TROCCO API
  - `entity/`: API response entities
  - `parameter/`: API request parameters
- **Provider Layer** (`internal/provider/`): Terraform provider implementation
  - `model/`: Terraform resource/data source models
  - `schema/`: Terraform schema definitions
  - `validator/`: Custom validation logic
  - `planmodifier/`: Custom plan modification logic

Resources communicate through: Terraform Configuration → Provider → Client → TROCCO API

## Key Commands

### Formatting

```sh
# Format all Go files (REQUIRED before commits)
golangci-lint run --fix

# Alternative Go formatting
go fmt ./...

# Format Terraform configuration files (REQUIRED)
terraform fmt
```

### Testing

```sh
# Run all acceptance tests
make testacc

# Run specific acceptance tests
make testacc TESTARGS="-run TestAccConnectionResource"

# Run unit tests
go test -v -cover ./...

# Generate HTML coverage report
make cover-html
```

### Development

```sh
# Build the provider
go build -o terraform-provider-trocco

# Install provider locally for testing
go install .
```

## Important Conventions

1. **Language**: Use English in all files, code comments, and pull requests
2. **Commit Messages**: Follow Conventional Commits format (e.g., `feat:`, `fix:`, `chore:`)
3. **Log Messages**: Use lowercase letters for all log messages
4. **Testing**: Write acceptance tests for all resources and data sources
5. **Validation**: Follow patterns in `docs/styleguide.md` for validators

## Resource Naming Conventions

When working with resources, use these terminology mappings:
- Connection = 接続情報
- Job Definition = 転送設定
- Datamart Definition = データマート定義
- Pipeline Definition = パイプライン定義
- Resource Group = リソースグループ
- Notification Destination = 通知先

## Testing Guidelines

1. Acceptance tests require `TF_ACC=1` environment variable
2. Use `TROCCO_TEST_URL` for custom API endpoint in tests
3. Test configurations go in `internal/provider/testdata/`
4. Follow existing test patterns using `testAccProtoV6ProviderFactories`

## Code Style Requirements

The project uses golangci-lint with these key linters:
- `errorlint`: Proper error handling
- `gosec`: Security issues
- `ineffassign`: Ineffective assignments
- `staticcheck`: Static analysis
- `testifylint`: Test code quality

Always run `golangci-lint run --fix` before committing Go code.

## Provider Development Patterns

1. **Schema Definition**: Define schemas in `internal/provider/schema/`
2. **Model Definition**: Create corresponding models in `internal/provider/model/`
3. **Validation**: Implement validators in `internal/provider/validator/`
4. **Plan Modifiers**: Use plan modifiers for computed fields or special update logic
5. **Client Integration**: Keep API logic in `internal/client/`, not in provider code

## Common Development Tasks

### Adding a New Resource

1. Create schema in `internal/provider/schema/[resource]_schema.go`
2. Create model in `internal/provider/model/[resource]_model.go`
3. Implement resource in `internal/provider/[resource]_resource.go`
4. Add client methods in `internal/client/`
5. Write acceptance tests in `internal/provider/[resource]_resource_test.go`
6. Add documentation in `docs/resources/[resource].md`

### Running Specific Tests

```sh
# Run a single test function
make testacc TESTARGS="-run TestAccConnectionResource_basic"

# Run all tests for a resource
make testacc TESTARGS="-run TestAccConnectionResource"

# Run with verbose output
make testacc TESTARGS="-v -run TestAccConnectionResource"
```

## Debugging

1. Enable debug logging: `export TF_LOG=DEBUG`
2. Use `tflog.Debug(ctx, "message", map[string]any{"key": "value"})` for logging
3. Provider logs use lowercase messages consistently

## CI/CD

GitHub Actions runs on all PRs:
- Go formatting check
- golangci-lint
- Unit tests
- Acceptance tests (on merge to main)

Ensure all checks pass before requesting review.