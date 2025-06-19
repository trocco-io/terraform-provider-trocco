# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

### Build & Install

```bash
# Build the provider
go build -o terraform-provider-trocco

# Install the provider locally (Linux/macOS)
mkdir -p ~/.terraform.d/plugins/registry.terraform.io/trocco-io/trocco/[version]/[platform]/
cp terraform-provider-trocco ~/.terraform.d/plugins/registry.terraform.io/trocco-io/trocco/[version]/[platform]/

# Generate docs
go generate
```

### Testing

```bash
# Run acceptance tests (requires API credentials)
# Set environment variables:
# TROCCO_API_KEY - Your TROCCO API key
# TROCCO_REGION - The region to use (japan, india, korea)
make testacc

# Run specific test
make testacc TESTARGS="-run TestAccConnectionResource"

# Generate test coverage report
make cover-html
```

### Development

```bash
# Format code
go fmt ./...

# Vet code
go vet ./...

# Generate documentation
go generate
```

## Code Architecture

The terraform-provider-trocco is a Terraform provider for TROCCO, a data integration platform. It allows users to manage TROCCO resources via Terraform.

### Key Components

1. **Provider Structure**
   - `main.go` - Entry point that serves the provider
   - `internal/provider/provider.go` - Defines the provider configuration and available resources
   - `version/version.go` - Contains version information

2. **Resources**
   - The provider exposes multiple resources that map to TROCCO entities:
     - `connection` - Data source connections (BigQuery, MySQL, S3, etc.)
     - `job_definition` - ETL job configurations
     - `pipeline_definition` - Pipeline workflows
     - `bigquery_datamart_definition` - BigQuery datamart settings
     - `notification_destination` - Where to send notifications
     - `team`, `user`, `resource_group`, `label` - Organization resources

3. **Internal Structure**
   - `internal/client/` - API client implementation for TROCCO
   - `internal/provider/model/` - Data models for resources
   - `internal/provider/schema/` - Terraform schema definitions
   - `internal/provider/planmodifier/` - Custom plan modifiers
   - `internal/provider/validator/` - Custom validators

4. **Documentation**
   - `docs/` - Generated documentation
   - `examples/` - Example Terraform configurations

### Authentication

The provider requires an API key for authentication, which can be provided in two ways:
1. Via the `TROCCO_API_KEY` environment variable
2. Via the `api_key` provider configuration

Regional configuration is also required, defaulting to "japan" if not specified.

### Resource Management Pattern

Each resource follows a similar pattern:
1. Define schema (terraform attributes)
2. Map terraform attributes to API models
3. Implement CRUD operations:
   - Create: Map terraform config to API call
   - Read: Get current state from API
   - Update: Map terraform config changes to API call 
   - Delete: Remove resource via API

Custom validators and plan modifiers handle special cases like required fields or defaults.