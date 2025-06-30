# CLAUDE.md

## Overview

This repository provides a Terraform provider for [TROCCO](https://trocco.io), a cloud ETL service.

## Dependencies

- Terraform Plugin Framework

## Commands

### Foramt

```sh
# Format all Go files.
go fmt ./...

# Format a specific Go file.
go fmt [FILE]
```

### Test

```sh
# Run all acceptance tests.
make testacc

# Run specific acceptance tests.
make testacc TESTARGS="-run TestAccSome"

# Run unit tests.
go test -v -cover ./...
```

## Instructions

- **MUST** Use English in files and pull requests
