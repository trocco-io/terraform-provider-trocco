# CLAUDE.md

## Overview

This repository provides a Terraform provider for [TROCCO](https://trocco.io), a cloud ETL service.

## Dependencies

- Terraform Plugin Framework

## Commands

### Tests

```sh
# Run all acceptance tests.
TROCCO_TEST_URL=https://localhost:4000 TROCCO_API_KEY=**** make testacc

# Run specific acceptance tests.
TROCCO_TEST_URL=https://localhost:4000 TROCCO_API_KEY=**** make testacc \
   TESTARGS="-run TestAccSome"

# Run unit tests.
go test -v -cover ./...
```
