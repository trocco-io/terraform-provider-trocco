# `test-adhoc`

## Description

`test-adhoc` run ad hoc tests for the TROCCO Terraform provider to discover problems that cannot be predicted in advance and cannot be detected by unit tests or acceptance tests.

## Usage

```sh
/test-adhoc [options]
```

### Options

Option | Required | Description
--- | --- | ---
`--dry-run` | No | Generate Terraform HCL files but do not run tests.
`--property <property_names...>` |No | Generate Terraform HCL files only for specific properties.
`--resource <resource_types...>` | No | Generate Terraform HCL files only for specific resource types.
`--language <language>` | No | Use the specified language for outputs.
`--max-cases <number>` | No | Specify the maximum number of test cases to generate.
`--min-cases <number>` | No | Specify the minimum number o
f test cases to generate.

### Examples

```sh
/test-adhoc --resource trocco_pipeline_definition --property labels,notifications --language ja --max-cases 10 --min-cases 5
```

## Rules

### Resource Dependencies

Many resources have dependencies on other resources. For example a `trocco_job_definition` resource requires a `trocco_connection` resource.

You have 2 options for handling such dependencies.

#### Option 1: Create resources within the test

You can create required resources in Terraform HCL files and reference them.

```hcl
# trocco_job_definition.tf

resource "trocco_job_definition" "test" {
  input_option = {
    bigquery_input_option = {
      bigquery_connection_id = trocco_connection.test_bigquery.id
      
      # ...
    }
  }

  # ...
}

resource "trocco_connection" "test_dependency_bigquery" {
  # ...
}
```

#### Option 2: Use existing resources

You can reference the existing resources below by their IDs (DO NOT delete them).

- **Users**
    - ID = 10626
    - ID = 10652
- **Labels** 
    - ID = 5192
    - ID = 5193
- **Connections**
    - Redshift Connection ID = 256

### Working Directory

Working directories must follows the structure below.

```
tests/
└── adhoc/
    ├── test_19700101000000/
    |   ├── result.md
    |   ├── schema.json
    |   ├── provider.tf
    |   ├── trocco_connection.tf
    |   ├── trocco_label.tf
    |   ├── ...
    |   ├── <resource_type>.tf
    |── test_19710101000000/
    |   └── ...
    └── test_<YYYYMMDDHHmmss>/
        └── ...
```

## Instructions

### Steps

#### 1. Initialize a test environment

Execute the following commands exactly to initialize a test environment.

```sh
go install .

TEST_DIR="tests/adhoc/test_$(date +%Y%m%d%H%M%S)" && mkdir -p "${TEST_DIR}" && cd "${TEST_DIR}" && pwd

cat <<EOF > provider.tf
terraform {
  required_providers {
    trocco = {
      source = "registry.terraform.io/trocco-io/trocco"
    }
  }
}

provider "trocco" {
  region = "japan"
}
EOF

terraform init

terraform providers schema --json > schema.json
```

### 2. Generate Terraform HCL files

Generate Terraform HCL files randomly based on the `schema.json` file in the scope that passes the `terraform validate` command.

#### Available existing resources

You can use the following existing resources to generate Terraform HCL files (DO NOT delete them).

- Users
    - ID = 10626
    - ID = 10652
- Labels
    - ID = 5192, Nmae = "label1"
    - ID = 5193, Name = "label2"
- Connections
    - Redshift
        - ID = 256

### 3. Run Terraform commands

Run Terraform commands under random scenarios.

Typically, problems occur in the following scenarios.

1. Updates
    1. Run `terraform apply` to create resources
    2. Modify Terraform HCL files
    3. Run `terraform apply` again to update resources
1. Imports
    1. Run `terraform apply` to create resources
    2. Run `terraform state rm` to remove resources from a Terraform state
    3. Run `terraform import` to import resources

However, these scenarios are just examples, so you must try out actively other scenarios that might cause problems.

### 4. Clean up the test enviornment

Delete the created resources using `terraform destroy` command.

### 5. Create a report

Create a report in the `result.md` file.
