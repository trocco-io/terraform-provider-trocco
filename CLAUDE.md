# CLAUDE.md

## Overview

This repository provides a Terraform provider for [TROCCO](https://trocco.io), a cloud ETL service.

## Dependencies

- Terraform Plugin Framework

## Commands

### Tests

```sh
# Run all acceptance tests.
make testacc

# Run specific acceptance tests.
make testacc TESTARGS="-run TestAccSome"

# Run unit tests.
go test -v -cover ./...
```

## Rules

- **MUST** Format Go code using `go fmt`

- Commit Messages
    - **MUST** Follows Conventional Commits
- PR Titles
    - **MUST** Follows Conventional Commits
