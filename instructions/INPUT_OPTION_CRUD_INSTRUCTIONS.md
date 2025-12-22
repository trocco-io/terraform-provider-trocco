# Input Option CRUD Implementation Instructions

This guide provides step-by-step instructions for implementing CRUD operations for new input options in the Terraform Provider for Trocco.

## Overview

When adding a new input option for a connection type, you need to implement the following layers:
1. **Client Layer**: API entities and parameters
2. **Provider Layer**: Models, schemas, and plan modifiers
3. **Tests**: Unit tests and integration tests
4. **Examples**: Usage examples and documentation

## File Structure

Replace `{connection_type}` with the actual connection type name (e.g., `mysql`, `postgresql`, `s3`, etc.)

### 1. Implementation Files

#### Client Layer
- `internal/client/entity/job_definition/input_option/{connection_type}.go` - API response entity
- `internal/client/parameter/job_definition/input_option/{connection_type}.go` - API request parameters
- `internal/client/job_definition.go` - Update to include new input option

#### Provider Layer  
- `internal/provider/model/job_definition/input_option/{connection_type}.go` - Terraform model
- `internal/provider/model/job_definition/input_option.go` - Update to register new input option
- `internal/provider/schema/job_definition/{connection_type}_input_option.go` - Terraform schema
- `internal/provider/schema/job_definition/input_option.go` - Update to include new schema
- `internal/provider/planmodifier/{connection_type}_input_option_plan_modifier.go` - Custom plan modifiers

#### Test Data
- `internal/provider/testdata/job_definition/{input}_to_{output}/create.tf` - Integration test data

### 2. Unit Tests
- `internal/client/job_definition_test.go` - Update with new input option tests
- `internal/provider/testdata/connection/{connection_type}_create.tf` - Connection test data

### 3. Examples
- `examples/resources/trocco_job_definition/{source}_to_{destination}.tf` - Job definition examples
- `examples/resources/trocco_job_definition/input_options/{connection_type}_input_option.tf` - Input option specific examples

## Implementation Steps

### Step 1: Create Client Entity

Create `internal/client/entity/job_definition/input_option/{connection_type}.go`:

```go
package input_option

type {ConnectionType}InputOption struct {
    // Define fields based on API response
    FieldName *string `json:"field_name,omitempty"`
    // Add more fields as needed
}
```

### Step 2: Create Client Parameters

Create `internal/client/parameter/job_definition/input_option/{connection_type}.go`:

```go
package input_option

type {ConnectionType}InputOptionInput struct {
    // Define fields for API requests
    FieldName *string `json:"field_name,omitempty"`
    // Add more fields as needed
}

type Update{ConnectionType}InputOptionInput struct {
    // Define fields for update requests
    FieldName *string `json:"field_name,omitempty"`
    // Add more fields as needed
}
```

### Step 3: Update Job Definition Client

Update `internal/client/job_definition.go`:

```go
// Add import
import inputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"

// Add field to JobDefinitionInput structs
{ConnectionType}InputOption *inputOptionParameters.{ConnectionType}InputOptionInput `json:"{connection_type}_input_option,omitempty"`

// Add field to UpdateJobDefinitionInput structs  
{ConnectionType}InputOption *inputOptionParameters.Update{ConnectionType}InputOptionInput `json:"{connection_type}_input_option,omitempty"`
```

### Step 4: Create Provider Model

Create `internal/provider/model/job_definition/input_option/{connection_type}.go`:

```go
package input_option

import (
    "context"
    "terraform-provider-trocco/internal/client/entity/job_definition/input_option"
    "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
    
    "github.com/hashicorp/terraform-plugin-framework/types"
)

type {ConnectionType}InputOption struct {
    FieldName types.String `tfsdk:"field_name"`
    // Add more fields as needed
}

func New{ConnectionType}InputOption(ctx context.Context, inputOption *input_option.{ConnectionType}InputOption) *{ConnectionType}InputOption {
    if inputOption == nil {
        return nil
    }
    
    return &{ConnectionType}InputOption{
        FieldName: types.StringPointerValue(inputOption.FieldName),
        // Map other fields
    }
}

func (inputOption *{ConnectionType}InputOption) ToInput(ctx context.Context) *inputOptionParameters.{ConnectionType}InputOptionInput {
    if inputOption == nil {
        return nil
    }
    
    return &inputOptionParameters.{ConnectionType}InputOptionInput{
        FieldName: inputOption.FieldName.ValueStringPointer(),
        // Map other fields
    }
}

func (inputOption *{ConnectionType}InputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.Update{ConnectionType}InputOptionInput {
    if inputOption == nil {
        return nil
    }
    
    return &inputOptionParameters.Update{ConnectionType}InputOptionInput{
        FieldName: inputOption.FieldName.ValueStringPointer(),
        // Map other fields
    }
}
```

### Step 5: Update Input Option Model

Update `internal/provider/model/job_definition/input_option.go`:

```go
// Add import
import inputOptionSchema "terraform-provider-trocco/internal/provider/schema/job_definition"

// Add field to InputOption struct
{ConnectionType}InputOption *{ConnectionType}InputOption `tfsdk:"{connection_type}_input_option"`

// Update NewInputOption function
func NewInputOption(ctx context.Context, inputOption *inputOptionEntities.InputOption) *InputOption {
    // Add case for new input option
    {connection_type}InputOption := New{ConnectionType}InputOption(ctx, inputOption.{ConnectionType}InputOption)
    
    return &InputOption{
        // ... existing fields
        {ConnectionType}InputOption: {connection_type}InputOption,
    }
}

// Update ToInput method
func (inputOption *InputOption) ToInput(ctx context.Context) *inputOptionParameters.InputOptionInput {
    return &inputOptionParameters.InputOptionInput{
        // ... existing fields
        {ConnectionType}InputOption: inputOption.{ConnectionType}InputOption.ToInput(ctx),
    }
}

// Update ToUpdateInput method
func (inputOption *InputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateInputOptionInput {
    return &inputOptionParameters.UpdateInputOptionInput{
        // ... existing fields  
        {ConnectionType}InputOption: inputOption.{ConnectionType}InputOption.ToUpdateInput(ctx),
    }
}

// Update Attributes method
func (inputOption *InputOption) Attributes() map[string]schema.Attribute {
    attributes := map[string]schema.Attribute{
        // ... existing attributes
        "{connection_type}_input_option": inputOptionSchema.{ConnectionType}InputOptionAttribute(),
    }
    return attributes
}
```

### Step 6: Create Schema

Create `internal/provider/schema/job_definition/{connection_type}_input_option.go`:

```go
package job_definition

import (
    "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
    "github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func {ConnectionType}InputOptionAttribute() schema.Attribute {
    return schema.SingleNestedAttribute{
        MarkdownDescription: "{ConnectionType} input option configuration.",
        Optional:            true,
        Attributes: map[string]schema.Attribute{
            "field_name": schema.StringAttribute{
                MarkdownDescription: "Description of the field.",
                Optional:            true,
                PlanModifiers: []planmodifier.String{
                    stringplanmodifier.RequiresReplace(),
                },
                Validators: []validator.String{
                    stringvalidator.UTF8LengthAtLeast(1),
                },
            },
            // Add more attributes as needed
        },
    }
}
```

### Step 7: Update Input Option Schema

Update `internal/provider/schema/job_definition/input_option.go`:

```go
// Add to Attributes function
"{connection_type}_input_option": {ConnectionType}InputOptionAttribute(),
```

### Step 8: Create Plan Modifier (if needed)

Create `internal/provider/planmodifier/{connection_type}_input_option_plan_modifier.go`:

```go
package planmodifier

import (
    "context"
    
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

type {ConnectionType}InputOptionPlanModifier struct{}

func (d *{ConnectionType}InputOptionPlanModifier) Description(ctx context.Context) string {
    return "Custom plan modifier for {ConnectionType} input options."
}

func (d *{ConnectionType}InputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
    return d.Description(ctx)
}

func (d *{ConnectionType}InputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
    // Implement custom logic if needed
}

func {ConnectionType}InputOptionPlanModifier() planmodifier.Object {
    return &{ConnectionType}InputOptionPlanModifier{}
}
```

### Step 9: Create Test Data

Create `internal/provider/testdata/job_definition/{input}_to_{output}/create.tf`:

```hcl
resource "trocco_job_definition" "test" {
  name                = "test-job"
  connection_id       = 123
  destination_id      = 456
  
  input_option {
    {connection_type}_input_option {
      field_name = "test_value"
      // Add more fields as needed
    }
  }
}
```

Create `internal/provider/testdata/connection/{connection_type}_create.tf`:

```hcl
resource "trocco_connection" "test" {
  connection_type = "{connection_type}"
  name           = "Test {ConnectionType} Connection"
  description    = "Test connection for {connection_type}"
  
  # Add connection-specific fields
  host      = "example.com"
  port      = 3306
  user_name = "testuser"
  password  = "testpass"
}
```

### Step 10: Update Unit Tests

Update `internal/client/job_definition_test.go`:

```go
func Test{ConnectionType}InputOption(t *testing.T) {
    // Add test cases for the new input option
    t.Run("Create with {ConnectionType} input option", func(t *testing.T) {
        // Implement test
    })
    
    t.Run("Update {ConnectionType} input option", func(t *testing.T) {
        // Implement test
    })
}
```

### Step 11: Create Examples

Create `examples/resources/trocco_job_definition/{source}_to_{destination}.tf`:

```hcl
resource "trocco_job_definition" "example" {
  name                = "Example Job"
  connection_id       = trocco_connection.source.id
  destination_id      = trocco_connection.destination.id
  
  input_option {
    {connection_type}_input_option {
      field_name = "example_value"
      // Add more example fields
    }
  }
}
```

Create `examples/resources/trocco_job_definition/input_options/{connection_type}_input_option.tf`:

```hcl
# {ConnectionType} Input Option Example
resource "trocco_job_definition" "{connection_type}_example" {
  name                = "{ConnectionType} Input Option Example"
  connection_id       = trocco_connection.{connection_type}.id
  destination_id      = trocco_connection.destination.id
  
  input_option {
    {connection_type}_input_option {
      field_name = "example_value"
      // Demonstrate all available options
      # field2 = "another_value"
      # field3 = true
    }
  }
}

# Example connection for reference
resource "trocco_connection" "{connection_type}" {
  connection_type = "{connection_type}"
  name           = "{ConnectionType} Connection"
  description    = "Example {connection_type} connection"
  
  # Add connection-specific configuration
  host      = "example.com"
  port      = 3306
  user_name = "username"
  password  = "password"
}
```

### Step 12: Create E2E Tests

Create comprehensive end-to-end tests to validate the input option implementation.

**Important**: All E2E tests must be added to `internal/provider/job_definition_resource_test.go` following the existing test patterns in that file.

#### A. E2E Integration Tests in `internal/provider/job_definition_resource_test.go`

**Required**: Add E2E tests to `internal/provider/job_definition_resource_test.go` following these patterns:

##### Pattern 1: Basic CRUD Test

Add a test function following the naming pattern `TestAccJobDefinitionResource_{Source}To{Destination}`:

```go
func TestAccJobDefinitionResource_{Source}To{Destination}(t *testing.T) {
    resource.Test(t, resource.TestCase{
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            // Step 1: Create
            {
                Config: testAccJobDefinitionConfig_{source}To{destination}(),
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckResourceAttr("trocco_job_definition.test", "name", "test-{source}-to-{destination}"),
                    resource.TestCheckResourceAttr("trocco_job_definition.test", "input_option_type", "{source}"),
                    resource.TestCheckResourceAttr("trocco_job_definition.test", "output_option_type", "{destination}"),
                    
                    // Check input option fields
                    resource.TestCheckResourceAttrSet("trocco_job_definition.test", "input_option.{source}_input_option.connection_id"),
                    resource.TestCheckResourceAttr("trocco_job_definition.test", "input_option.{source}_input_option.table", "test_table"),
                    resource.TestCheckResourceAttr("trocco_job_definition.test", "input_option.{source}_input_option.query", "SELECT * FROM test"),
                ),
            },
            // Step 2: Update
            {
                Config: testAccJobDefinitionConfig_{source}To{destination}Update(),
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckResourceAttr("trocco_job_definition.test", "name", "test-{source}-to-{destination}-updated"),
                    resource.TestCheckResourceAttr("trocco_job_definition.test", "input_option.{source}_input_option.query", "SELECT * FROM test WHERE id > 0"),
                ),
            },
            // Step 3: Import
            {
                ResourceName:      "trocco_job_definition.test",
                ImportState:       true,
                ImportStateVerify: true,
                ImportStateVerifyIgnore: []string{
                    "input_option.{source}_input_option.password", // Sensitive fields
                },
            },
        },
    })
}
```

##### Pattern 2: Test Configuration Helper Functions

Add helper functions that return Terraform configurations using testdata files:

```go
func testAccJobDefinitionConfig_{source}To{destination}() string {
    // Use testdata file
    config, err := os.ReadFile("testdata/job_definition/{source}_to_{destination}/create.tf")
    if err != nil {
        panic(fmt.Sprintf("Failed to read testdata: %v", err))
    }
    return string(config)
}

func testAccJobDefinitionConfig_{source}To{destination}Update() string {
    config, err := os.ReadFile("testdata/job_definition/{source}_to_{destination}/update.tf")
    if err != nil {
        panic(fmt.Sprintf("Failed to read testdata: %v", err))
    }
    return string(config)
}
```

#### B. Test Data Files

**Required**: Create test data files in `internal/provider/testdata/job_definition/{source}_to_{destination}/`:

1. **create.tf** - Initial resource configuration
2. **update.tf** - Updated resource configuration (for update tests)
3. **invalid_*.tf** - Invalid configurations (for validation tests)

Example structure for `create.tf`:

```hcl
resource "trocco_job_definition" "test" {
  name           = "test-{source}-to-{destination}"
  description    = "E2E test for {source} to {destination}"
  
  input_option_type  = "{source}"
  output_option_type = "{destination}"
  
  filter_columns = [
    {
      name                         = "id"
      src                          = "id"
      type                         = "long"
      default                      = ""
      has_parser                   = false
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      json_expand_columns          = null
    }
  ]
  
  input_option = {
    {source}_input_option = {
      # Source-specific configuration
      # Include all required fields
      # Include optional fields to test full coverage
    }
  }
  
  output_option = {
    {destination}_output_option = {
      # Destination-specific configuration
    }
  }
}
```

#### C. Running E2E Tests

Execute tests with:

```bash
# Run all job definition tests
go test -v ./internal/provider -run TestAccJobDefinitionResource

# Run specific test
go test -v ./internal/provider -run TestAccJobDefinitionResource_{Source}To{Destination}

# Run with timeout for long tests
go test -v -timeout 30m ./internal/provider -run TestAccJobDefinitionResource_{Source}To{Destination}
```

#### D. Test Coverage Requirements

Ensure your E2E tests cover:

1. **Basic CRUD Operations**
   - Create resource with all required fields
   - Read and verify all attributes
   - Update mutable fields
   - Import existing resource
   - Delete resource

2. **Field Validation**
   - Required fields throw errors when missing
   - Optional fields have correct defaults
   - Enum fields reject invalid values
   - Numeric fields respect min/max constraints

3. **Complex Nested Structures**
   - Nested objects (parsers, configurations)
   - Lists of objects (columns, filters)
   - Conditional fields based on parent values

4. **Plan Modifiers**
   - Validation logic works correctly
   - Computed fields are set properly
   - RequiresReplace is triggered appropriately

5. **State Management**
   - State correctly reflects API response
   - Sensitive fields are handled properly
   - Null vs empty distinction is maintained

## Validation Checklist

Before submitting your implementation, ensure:

- [ ] All files are created with proper naming conventions
- [ ] Entity, parameter, and model structs have proper JSON tags
- [ ] Schema attributes have appropriate validators and plan modifiers
- [ ] Plan modifiers are implemented if custom logic is needed
- [ ] Unit tests cover create, read, update, and delete operations
- [ ] Test data includes valid Terraform configurations
- [ ] Examples demonstrate all available options
- [ ] Documentation strings are clear and descriptive
- [ ] Error handling is implemented where appropriate
- [ ] Null/empty value handling is considered
- [ ] Integration tests pass with the new input option
- [ ] **E2E tests added to `internal/provider/job_definition_resource_test.go`**
- [ ] **E2E tests cover CREATE, READ, UPDATE, DELETE operations**
- [ ] **E2E tests include import state verification**
- [ ] **E2E test data files created in `internal/provider/testdata/job_definition/`**
- [ ] E2E tests validate plan modifier logic
- [ ] E2E tests verify sensitive field handling
- [ ] E2E tests confirm state correctly reflects API responses

## Common Patterns

### Handling Optional Fields
```go
// In entity
Field *string `json:"field,omitempty"`

// In model
Field types.String `tfsdk:"field"`

// In conversion
Field: types.StringPointerValue(entity.Field),
```

### Handling Arrays/Lists
```go
// In entity
Items []ItemEntity `json:"items,omitempty"`

// In model  
Items types.List `tfsdk:"items"`

// In schema
"items": schema.ListNestedAttribute{
    NestedObject: schema.NestedAttributeObject{
        Attributes: map[string]schema.Attribute{
            // Item attributes
        },
    },
},
```

### Handling Complex Nested Objects
```go
// Create separate struct for nested objects
type NestedObject struct {
    Field types.String `tfsdk:"field"`
}

// Use SingleNestedAttribute in schema
"nested": schema.SingleNestedAttribute{
    Attributes: map[string]schema.Attribute{
        "field": schema.StringAttribute{
            // Configuration
        },
    },
},
```

## Troubleshooting

### Common Issues

1. **Import Cycles**: Ensure proper package structure and avoid circular imports
2. **Type Mismatches**: Use appropriate `types.*` from Terraform Framework
3. **Null Handling**: Always check for null values before conversion
4. **Schema Validation**: Ensure schema attributes match model fields exactly
5. **Plan Modifier Issues**: Use appropriate plan modifier types (String, Object, etc.)

### Debug Tips

1. Use `fmt.Printf()` for debugging during development
2. Check Terraform logs with `TF_LOG=DEBUG`
3. Validate JSON serialization/deserialization manually
4. Test with minimal configurations first
5. Use `terraform plan` to verify schema changes

## Notes

- Replace all `{connection_type}` placeholders with the actual connection type name
- Replace all `{ConnectionType}` placeholders with the PascalCase version
- Replace `{input}` and `{output}` with actual source and destination types
- Follow existing code patterns in the repository
- Ensure backward compatibility when updating existing files
- Add appropriate error handling and validation
- Document any special requirements or limitations
