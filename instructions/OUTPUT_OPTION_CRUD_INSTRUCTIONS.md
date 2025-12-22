# Output Option CRUD Implementation Instructions

This guide provides step-by-step instructions for implementing CRUD operations for new output options in the Terraform Provider for Trocco.

## Overview

When adding a new output option for a connection type, you need to implement the following layers:
1. **Client Layer**: API entities and parameters
2. **Provider Layer**: Models, schemas, and plan modifiers
3. **Tests**: Unit tests and integration tests
4. **Examples**: Usage examples and documentation

## File Structure

Replace `{connection_type}` with the actual connection type name (e.g., `bigquery`, `snowflake`, `s3`, etc.)
Replace `{input}` and `{output}` with actual source and destination types (e.g., `mysql_to_bigquery`)

### 1. Implementation Files

#### Client Layer
- `internal/client/entity/job_definition/output_option/{connection_type}.go` - API response entity
- `internal/client/parameter/job_definition/output_option/{connection_type}.go` - API request parameters
- `internal/client/job_definition.go` - Update to include new output option

#### Provider Layer  
- `internal/provider/model/job_definition/output_option/{connection_type}.go` - Terraform model
- `internal/provider/model/job_definition/output_option.go` - Update to register new output option
- `internal/provider/schema/job_definition/{connection_type}_output_option.go` - Terraform schema
- `internal/provider/schema/job_definition/output_option.go` - Update to include new schema
- `internal/provider/planmodifier/{connection_type}_output_option_plan_modifier.go` - Custom plan modifiers
- `internal/provider/planmodifier/{connection_type}_output_option_column_plan_modifier.go` - Column-specific plan modifiers

#### Test Data
- `internal/provider/testdata/job_definition/{input}_to_{output}/create.tf` - Integration test data

### 2. Unit Tests
- `internal/client/job_definition_test.go` - Update with new output option tests (optional)
- `internal/provider/testdata/connection/{connection_type}_create.tf` - Connection test data

### 3. Examples
- `examples/resources/trocco_job_definition/{source}_to_{destination}.tf` - Job definition examples
- `examples/resources/trocco_job_definition/output_options/{connection_type}_output_option.tf` - Output option specific examples

## Implementation Steps

### Step 1: Create Client Entity

Create `internal/client/entity/job_definition/output_option/{connection_type}.go`:

```go
package output_option

type {ConnectionType}OutputOption struct {
    // Basic fields
    TableName        *string `json:"table_name,omitempty"`
    DatabaseName     *string `json:"database_name,omitempty"`
    WriteDisposition *string `json:"write_disposition,omitempty"`
    
    // Connection-specific fields
    ProjectID        *string `json:"project_id,omitempty"` // For BigQuery
    Dataset          *string `json:"dataset,omitempty"`    // For BigQuery
    Warehouse        *string `json:"warehouse,omitempty"`  // For Snowflake
    
    // Column mapping
    Columns []Column `json:"columns,omitempty"`
}

type Column struct {
    Name        *string `json:"name,omitempty"`
    Type        *string `json:"type,omitempty"`
    Description *string `json:"description,omitempty"`
    // Add more column-specific fields as needed
}
```

### Step 2: Create Client Parameters

Create `internal/client/parameter/job_definition/output_option/{connection_type}.go`:

```go
package output_option

type {ConnectionType}OutputOptionInput struct {
    // Basic fields
    TableName        *string `json:"table_name,omitempty"`
    DatabaseName     *string `json:"database_name,omitempty"`
    WriteDisposition *string `json:"write_disposition,omitempty"`
    
    // Connection-specific fields
    ProjectID        *string `json:"project_id,omitempty"`
    Dataset          *string `json:"dataset,omitempty"`
    Warehouse        *string `json:"warehouse,omitempty"`
    
    // Column mapping
    Columns []ColumnInput `json:"columns,omitempty"`
}

type Update{ConnectionType}OutputOptionInput struct {
    // Same structure as input but for update requests
    TableName        *string `json:"table_name,omitempty"`
    DatabaseName     *string `json:"database_name,omitempty"`
    WriteDisposition *string `json:"write_disposition,omitempty"`
    ProjectID        *string `json:"project_id,omitempty"`
    Dataset          *string `json:"dataset,omitempty"`
    Warehouse        *string `json:"warehouse,omitempty"`
    Columns          []ColumnInput `json:"columns,omitempty"`
}

type ColumnInput struct {
    Name        *string `json:"name,omitempty"`
    Type        *string `json:"type,omitempty"`
    Description *string `json:"description,omitempty"`
}
```

### Step 3: Update Job Definition Client

Update `internal/client/job_definition.go`:

```go
// Add import
import outputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"

// Add field to JobDefinitionInput struct
{ConnectionType}OutputOption *outputOptionParameters.{ConnectionType}OutputOptionInput `json:"{connection_type}_output_option,omitempty"`

// Add field to UpdateJobDefinitionInput struct  
{ConnectionType}OutputOption *outputOptionParameters.Update{ConnectionType}OutputOptionInput `json:"{connection_type}_output_option,omitempty"`
```

### Step 4: Create Provider Model

Create `internal/provider/model/job_definition/output_option/{connection_type}.go`:

```go
package output_option

import (
    "context"
    "terraform-provider-trocco/internal/client/entity/job_definition/output_option"
    "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"
    
    "github.com/hashicorp/terraform-plugin-framework/types"
)

type {ConnectionType}OutputOption struct {
    TableName        types.String `tfsdk:"table_name"`
    DatabaseName     types.String `tfsdk:"database_name"`
    WriteDisposition types.String `tfsdk:"write_disposition"`
    ProjectID        types.String `tfsdk:"project_id"`
    Dataset          types.String `tfsdk:"dataset"`
    Warehouse        types.String `tfsdk:"warehouse"`
    Columns          types.List   `tfsdk:"columns"` // List of Column objects
}

type Column struct {
    Name        types.String `tfsdk:"name"`
    Type        types.String `tfsdk:"type"`
    Description types.String `tfsdk:"description"`
}

func New{ConnectionType}OutputOption(ctx context.Context, outputOption *output_option.{ConnectionType}OutputOption) *{ConnectionType}OutputOption {
    if outputOption == nil {
        return nil
    }
    
    // Convert columns
    columns := make([]Column, 0)
    if outputOption.Columns != nil {
        for _, col := range outputOption.Columns {
            columns = append(columns, Column{
                Name:        types.StringPointerValue(col.Name),
                Type:        types.StringPointerValue(col.Type),
                Description: types.StringPointerValue(col.Description),
            })
        }
    }
    
    columnsList, diags := types.ListValueFrom(ctx, types.ObjectType{
        AttrTypes: map[string]attr.Type{
            "name":        types.StringType,
            "type":        types.StringType,
            "description": types.StringType,
        },
    }, columns)
    
    if diags.HasError() {
        // Handle error appropriately
        columnsList = types.ListNull(types.ObjectType{})
    }
    
    return &{ConnectionType}OutputOption{
        TableName:        types.StringPointerValue(outputOption.TableName),
        DatabaseName:     types.StringPointerValue(outputOption.DatabaseName),
        WriteDisposition: types.StringPointerValue(outputOption.WriteDisposition),
        ProjectID:        types.StringPointerValue(outputOption.ProjectID),
        Dataset:          types.StringPointerValue(outputOption.Dataset),
        Warehouse:        types.StringPointerValue(outputOption.Warehouse),
        Columns:          columnsList,
    }
}

func (outputOption *{ConnectionType}OutputOption) ToInput(ctx context.Context) *outputOptionParameters.{ConnectionType}OutputOptionInput {
    if outputOption == nil {
        return nil
    }
    
    // Convert columns
    var columns []outputOptionParameters.ColumnInput
    if !outputOption.Columns.IsNull() && !outputOption.Columns.IsUnknown() {
        var columnModels []Column
        outputOption.Columns.ElementsAs(ctx, &columnModels, false)
        
        for _, col := range columnModels {
            columns = append(columns, outputOptionParameters.ColumnInput{
                Name:        col.Name.ValueStringPointer(),
                Type:        col.Type.ValueStringPointer(),
                Description: col.Description.ValueStringPointer(),
            })
        }
    }
    
    return &outputOptionParameters.{ConnectionType}OutputOptionInput{
        TableName:        outputOption.TableName.ValueStringPointer(),
        DatabaseName:     outputOption.DatabaseName.ValueStringPointer(),
        WriteDisposition: outputOption.WriteDisposition.ValueStringPointer(),
        ProjectID:        outputOption.ProjectID.ValueStringPointer(),
        Dataset:          outputOption.Dataset.ValueStringPointer(),
        Warehouse:        outputOption.Warehouse.ValueStringPointer(),
        Columns:          columns,
    }
}

func (outputOption *{ConnectionType}OutputOption) ToUpdateInput(ctx context.Context) *outputOptionParameters.Update{ConnectionType}OutputOptionInput {
    if outputOption == nil {
        return nil
    }
    
    // Convert columns
    var columns []outputOptionParameters.ColumnInput
    if !outputOption.Columns.IsNull() && !outputOption.Columns.IsUnknown() {
        var columnModels []Column
        outputOption.Columns.ElementsAs(ctx, &columnModels, false)
        
        for _, col := range columnModels {
            columns = append(columns, outputOptionParameters.ColumnInput{
                Name:        col.Name.ValueStringPointer(),
                Type:        col.Type.ValueStringPointer(),
                Description: col.Description.ValueStringPointer(),
            })
        }
    }
    
    return &outputOptionParameters.Update{ConnectionType}OutputOptionInput{
        TableName:        outputOption.TableName.ValueStringPointer(),
        DatabaseName:     outputOption.DatabaseName.ValueStringPointer(),
        WriteDisposition: outputOption.WriteDisposition.ValueStringPointer(),
        ProjectID:        outputOption.ProjectID.ValueStringPointer(),
        Dataset:          outputOption.Dataset.ValueStringPointer(),
        Warehouse:        outputOption.Warehouse.ValueStringPointer(),
        Columns:          columns,
    }
}
```

### Step 5: Update Output Option Model

Update `internal/provider/model/job_definition/output_option.go`:

```go
// Add import
import outputOptionSchema "terraform-provider-trocco/internal/provider/schema/job_definition"

// Add field to OutputOption struct
{ConnectionType}OutputOption *{ConnectionType}OutputOption `tfsdk:"{connection_type}_output_option"`

// Update NewOutputOption function
func NewOutputOption(ctx context.Context, outputOption *outputOptionEntities.OutputOption) *OutputOption {
    // Add case for new output option
    {connection_type}OutputOption := New{ConnectionType}OutputOption(ctx, outputOption.{ConnectionType}OutputOption)
    
    return &OutputOption{
        // ... existing fields
        {ConnectionType}OutputOption: {connection_type}OutputOption,
    }
}

// Update ToInput method
func (outputOption *OutputOption) ToInput(ctx context.Context) *outputOptionParameters.OutputOptionInput {
    return &outputOptionParameters.OutputOptionInput{
        // ... existing fields
        {ConnectionType}OutputOption: outputOption.{ConnectionType}OutputOption.ToInput(ctx),
    }
}

// Update ToUpdateInput method
func (outputOption *OutputOption) ToUpdateInput(ctx context.Context) *outputOptionParameters.UpdateOutputOptionInput {
    return &outputOptionParameters.UpdateOutputOptionInput{
        // ... existing fields  
        {ConnectionType}OutputOption: outputOption.{ConnectionType}OutputOption.ToUpdateInput(ctx),
    }
}

// Update Attributes method
func (outputOption *OutputOption) Attributes() map[string]schema.Attribute {
    attributes := map[string]schema.Attribute{
        // ... existing attributes
        "{connection_type}_output_option": outputOptionSchema.{ConnectionType}OutputOptionAttribute(),
    }
    return attributes
}
```

### Step 6: Create Schema

Create `internal/provider/schema/job_definition/{connection_type}_output_option.go`:

```go
package job_definition

import (
    "github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
    "github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
    "github.com/hashicorp/terraform-plugin-framework/schema/validator"
    "github.com/hashicorp/terraform-plugin-framework/types"
    planModifier "terraform-provider-trocco/internal/provider/planmodifier"
)

func {ConnectionType}OutputOptionAttribute() schema.Attribute {
    return schema.SingleNestedAttribute{
        MarkdownDescription: "{ConnectionType} output option configuration.",
        Optional:            true,
        Attributes: map[string]schema.Attribute{
            "table_name": schema.StringAttribute{
                MarkdownDescription: "Name of the destination table.",
                Required:            true,
                Validators: []validator.String{
                    stringvalidator.UTF8LengthAtLeast(1),
                },
            },
            "database_name": schema.StringAttribute{
                MarkdownDescription: "Name of the destination database/schema.",
                Optional:            true,
                Validators: []validator.String{
                    stringvalidator.UTF8LengthAtLeast(1),
                },
            },
            "write_disposition": schema.StringAttribute{
                MarkdownDescription: "Write disposition. Valid values: `append`, `truncate`, `replace`.",
                Optional:            true,
                Computed:            true,
                Default:             stringdefault.StaticString("append"),
                Validators: []validator.String{
                    stringvalidator.OneOf("append", "truncate", "replace"),
                },
            },
            // Connection-specific fields
            "project_id": schema.StringAttribute{
                MarkdownDescription: "BigQuery: GCP Project ID.",
                Optional:            true,
                Validators: []validator.String{
                    stringvalidator.UTF8LengthAtLeast(1),
                },
            },
            "dataset": schema.StringAttribute{
                MarkdownDescription: "BigQuery: Dataset name.",
                Optional:            true,
                Validators: []validator.String{
                    stringvalidator.UTF8LengthAtLeast(1),
                },
            },
            "warehouse": schema.StringAttribute{
                MarkdownDescription: "Snowflake: Warehouse name.",
                Optional:            true,
                Validators: []validator.String{
                    stringvalidator.UTF8LengthAtLeast(1),
                },
            },
            "columns": schema.ListNestedAttribute{
                MarkdownDescription: "Column definitions for the output table.",
                Optional:            true,
                NestedObject: schema.NestedAttributeObject{
                    Attributes: map[string]schema.Attribute{
                        "name": schema.StringAttribute{
                            MarkdownDescription: "Column name.",
                            Required:            true,
                            Validators: []validator.String{
                                stringvalidator.UTF8LengthAtLeast(1),
                            },
                        },
                        "type": schema.StringAttribute{
                            MarkdownDescription: "Column data type.",
                            Required:            true,
                            Validators: []validator.String{
                                stringvalidator.UTF8LengthAtLeast(1),
                            },
                        },
                        "description": schema.StringAttribute{
                            MarkdownDescription: "Column description.",
                            Optional:            true,
                            Validators: []validator.String{
                                stringvalidator.UTF8LengthAtLeast(1),
                            },
                        },
                    },
                },
                Validators: []validator.List{
                    listvalidator.SizeAtLeast(1),
                },
                PlanModifiers: []planmodifier.List{
                    planModifier.{ConnectionType}OutputOptionColumnPlanModifier(),
                },
            },
        },
        PlanModifiers: []planmodifier.Object{
            planModifier.{ConnectionType}OutputOptionPlanModifier(),
        },
    }
}
```

### Step 7: Update Output Option Schema

Update `internal/provider/schema/job_definition/output_option.go`:

```go
// Add to Attributes function
"{connection_type}_output_option": {ConnectionType}OutputOptionAttribute(),
```

### Step 8: Create Plan Modifiers

Create `internal/provider/planmodifier/{connection_type}_output_option_plan_modifier.go`:

```go
package planmodifier

import (
    "context"
    
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

type {ConnectionType}OutputOptionPlanModifier struct{}

func (d *{ConnectionType}OutputOptionPlanModifier) Description(ctx context.Context) string {
    return "Custom plan modifier for {ConnectionType} output options."
}

func (d *{ConnectionType}OutputOptionPlanModifier) MarkdownDescription(ctx context.Context) string {
    return d.Description(ctx)
}

func (d *{ConnectionType}OutputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
    // Implement custom validation or modification logic
    // For example: validate project_id is set when using BigQuery
    
    if req.PlanValue.IsNull() {
        return
    }
    
    // Add custom logic here
}

func {ConnectionType}OutputOptionPlanModifier() planmodifier.Object {
    return &{ConnectionType}OutputOptionPlanModifier{}
}
```

Create `internal/provider/planmodifier/{connection_type}_output_option_column_plan_modifier.go`:

```go
package planmodifier

import (
    "context"
    
    "github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
)

type {ConnectionType}OutputOptionColumnPlanModifier struct{}

func (d *{ConnectionType}OutputOptionColumnPlanModifier) Description(ctx context.Context) string {
    return "Custom plan modifier for {ConnectionType} output option columns."
}

func (d *{ConnectionType}OutputOptionColumnPlanModifier) MarkdownDescription(ctx context.Context) string {
    return d.Description(ctx)
}

func (d *{ConnectionType}OutputOptionColumnPlanModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
    // Implement custom validation for columns
    // For example: validate column types are supported by the destination
    
    if req.PlanValue.IsNull() {
        return
    }
    
    // Add column validation logic here
}

func {ConnectionType}OutputOptionColumnPlanModifier() planmodifier.List {
    return &{ConnectionType}OutputOptionColumnPlanModifier{}
}
```

### Step 9: Create Test Data

Create `internal/provider/testdata/job_definition/{input}_to_{output}/create.tf`:

```hcl
resource "trocco_job_definition" "test" {
  name                = "test-job"
  connection_id       = 123
  destination_id      = 456
  
  output_option {
    {connection_type}_output_option {
      table_name        = "test_table"
      database_name     = "test_db"
      write_disposition = "append"
      project_id        = "test-project"  # For BigQuery
      dataset           = "test_dataset"  # For BigQuery
      warehouse         = "test_warehouse" # For Snowflake
      
      columns = [
        {
          name        = "id"
          type        = "INTEGER"
          description = "Primary key"
        },
        {
          name        = "name"
          type        = "STRING"
          description = "User name"
        }
      ]
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
  
  # Add connection-specific fields based on connection type
  project_id               = "test-project"        # For BigQuery/GCS
  service_account_json_key = "test-service-key"    # For BigQuery
  host                     = "test.snowflake.com"  # For Snowflake
  user_name               = "testuser"             # For Snowflake
  auth_method             = "user_password"        # For Snowflake
  password                = "testpass"             # For Snowflake
}
```

### Step 10: Update Unit Tests (Optional)

Update `internal/client/job_definition_test.go`:

```go
func Test{ConnectionType}OutputOption(t *testing.T) {
    // Add test cases for the new output option
    t.Run("Create with {ConnectionType} output option", func(t *testing.T) {
        // Implement test
    })
    
    t.Run("Update {ConnectionType} output option", func(t *testing.T) {
        // Implement test
    })
    
    t.Run("Create with {ConnectionType} output option columns", func(t *testing.T) {
        // Test column handling
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
  
  output_option {
    {connection_type}_output_option {
      table_name        = "destination_table"
      database_name     = "analytics"
      write_disposition = "append"
      
      # Connection-specific fields
      project_id = "my-gcp-project"    # For BigQuery
      dataset    = "my_dataset"        # For BigQuery
      warehouse  = "compute_wh"        # For Snowflake
      
      columns = [
        {
          name        = "user_id"
          type        = "INTEGER"
          description = "Unique user identifier"
        },
        {
          name        = "created_at"
          type        = "TIMESTAMP"
          description = "Record creation timestamp"
        },
        {
          name        = "email"
          type        = "STRING"
          description = "User email address"
        }
      ]
    }
  }
}
```

Create `examples/resources/trocco_job_definition/output_options/{connection_type}_output_option.tf`:

```hcl
# {ConnectionType} Output Option Example
resource "trocco_job_definition" "{connection_type}_output_example" {
  name                = "{ConnectionType} Output Option Example"
  connection_id       = trocco_connection.source.id
  destination_id      = trocco_connection.{connection_type}.id
  
  output_option {
    {connection_type}_output_option {
      table_name        = "example_output_table"
      database_name     = "analytics_db"
      write_disposition = "truncate"
      
      # BigQuery specific fields
      project_id = "my-gcp-project"
      dataset    = "analytics"
      
      # Snowflake specific fields  
      warehouse = "compute_wh"
      
      # Column definitions
      columns = [
        {
          name        = "id"
          type        = "INTEGER"
          description = "Primary key"
        },
        {
          name        = "timestamp"
          type        = "TIMESTAMP"
          description = "Event timestamp"
        },
        {
          name        = "event_name"
          type        = "STRING"
          description = "Name of the event"
        },
        {
          name        = "user_properties"
          type        = "JSON"
          description = "User properties as JSON"
        }
      ]
    }
  }
}

# Example destination connection
resource "trocco_connection" "{connection_type}" {
  connection_type = "{connection_type}"
  name           = "{ConnectionType} Destination Connection"
  description    = "Example {connection_type} destination connection"
  
  # Connection-specific configuration
  project_id               = "my-gcp-project"
  service_account_json_key = file("path/to/service-account-key.json")
}
```

### Step 12: Create E2E Tests

Create comprehensive end-to-end tests to validate the output option implementation.

**Important**: All E2E tests must be added to `internal/provider/job_definition_resource_test.go` following the existing test patterns in that file.

#### A. Unit Tests for Client Layer

Update `internal/client/job_definition_test.go` with output option tests (optional - for complex validation logic):

```go
func Test{ConnectionType}OutputOption(t *testing.T) {
    t.Run("Create with {ConnectionType} output option", func(t *testing.T) {
        requestBody := client.CreateJobDefinitionInput{
            Name:           "Test {ConnectionType} Job",
            ConnectionID:   123,
            DestinationID:  456,
            OutputOption: &client.OutputOptionInput{
                {ConnectionType}OutputOption: model.WrapObject(&outputOptionParameters.{ConnectionType}OutputOptionInput{
                    // Add required fields
                }),
            },
        }
        
        // Test API call and response validation
        assertJobDefinitionCreated(t, requestBody)
    })
}
```

#### B. E2E Integration Tests in `internal/provider/job_definition_resource_test.go`

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
                    
                    // Check output option fields
                    resource.TestCheckResourceAttrSet("trocco_job_definition.test", "output_option.{destination}_output_option.connection_id"),
                    resource.TestCheckResourceAttr("trocco_job_definition.test", "output_option.{destination}_output_option.table", "test_table"),
                    resource.TestCheckResourceAttr("trocco_job_definition.test", "output_option.{destination}_output_option.mode", "append"),
                ),
            },
            // Step 2: Update
            {
                Config: testAccJobDefinitionConfig_{source}To{destination}Update(),
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckResourceAttr("trocco_job_definition.test", "name", "test-{source}-to-{destination}-updated"),
                    resource.TestCheckResourceAttr("trocco_job_definition.test", "output_option.{destination}_output_option.mode", "replace"),
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

##### Pattern 3: Test for Multiple Formatter Types (if applicable)

For output options with multiple formatters (e.g., CSV, JSONL):

```go
func TestAccJobDefinitionResource_{Source}To{Destination}CSV(t *testing.T) {
    resource.Test(t, resource.TestCase{
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: testAccJobDefinitionConfig_{source}To{destination}CSV(),
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckResourceAttr("trocco_job_definition.test", "output_option.{destination}_output_option.csv_formatter.delimiter", ","),
                    resource.TestCheckResourceAttr("trocco_job_definition.test", "output_option.{destination}_output_option.csv_formatter.header_line", "true"),
                ),
            },
        },
    })
}

func TestAccJobDefinitionResource_{Source}To{Destination}JSONL(t *testing.T) {
    resource.Test(t, resource.TestCase{
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config: testAccJobDefinitionConfig_{source}To{destination}JSONL(),
                Check: resource.ComposeAggregateTestCheckFunc(
                    resource.TestCheckResourceAttr("trocco_job_definition.test", "output_option.{destination}_output_option.jsonl_formatter.encoding", "UTF-8"),
                    resource.TestCheckResourceAttr("trocco_job_definition.test", "output_option.{destination}_output_option.jsonl_formatter.newline", "LF"),
                ),
            },
        },
    })
}
```

##### Pattern 4: Validation Tests

Test plan modifier validation:

```go
func TestAccJobDefinitionResource_{Source}To{Destination}Validation(t *testing.T) {
    resource.Test(t, resource.TestCase{
        ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
        Steps: []resource.TestStep{
            {
                Config:      testAccJobDefinitionConfig_{source}To{destination}InvalidFormatter(),
                ExpectError: regexp.MustCompile("either csv_formatter or jsonl_formatter must be provided"),
            },
            {
                Config:      testAccJobDefinitionConfig_{source}To{destination}MissingRequired(),
                ExpectError: regexp.MustCompile("connection_id is required"),
            },
        },
    })
}
```

#### C. Test Data Files

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
    }
  }
  
  output_option = {
    {destination}_output_option = {
      # Destination-specific configuration
      # Include all required fields
      # Include optional fields to test full coverage
    }
  }
}
```

#### D. Running E2E Tests

Execute tests with:

```bash
# Run all job definition tests
go test -v ./internal/provider -run TestAccJobDefinitionResource

# Run specific test
go test -v ./internal/provider -run TestAccJobDefinitionResource_{Source}To{Destination}

# Run with timeout for long tests
go test -v -timeout 30m ./internal/provider -run TestAccJobDefinitionResource_{Source}To{Destination}
```

#### E. Test Coverage Requirements

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
   - Nested objects (formatters, configurations)
   - Lists of objects (columns, custom variables)
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
- [ ] Plan modifiers implement proper validation for connection-specific fields
- [ ] Column plan modifiers validate column types and constraints
- [ ] Unit tests cover create, read, update, and delete operations
- [ ] Test data includes valid Terraform configurations with columns
- [ ] Examples demonstrate all available options including columns
- [ ] Documentation strings are clear and descriptive
- [ ] Error handling is implemented for column operations
- [ ] Null/empty value handling is considered for all fields
- [ ] Integration tests pass with the new output option
- [ ] **E2E tests added to `internal/provider/job_definition_resource_test.go`**
- [ ] **E2E tests cover CREATE, READ, UPDATE, DELETE operations**
- [ ] **E2E tests include import state verification**
- [ ] **E2E test data files created in `internal/provider/testdata/job_definition/`**
- [ ] E2E tests cover all formatter types (if applicable: CSV, JSONL, etc.)
- [ ] E2E tests validate plan modifier logic
- [ ] E2E tests check custom variable handling (if applicable)
- [ ] Performance tests pass with large configurations
- [ ] Validation tests ensure proper error messages
- [ ] Tests validate formatter type consistency (if applicable)
- [ ] Tests check required field validation
- [ ] Tests verify sensitive field handling
- [ ] Tests confirm state correctly reflects API responses

## Common Patterns

### Handling Column Lists
```go
// In entity
Columns []Column `json:"columns,omitempty"`

// In model  
Columns types.List `tfsdk:"columns"`

// In schema
"columns": schema.ListNestedAttribute{
    NestedObject: schema.NestedAttributeObject{
        Attributes: map[string]schema.Attribute{
            "name": schema.StringAttribute{
                Required: true,
            },
            "type": schema.StringAttribute{
                Required: true,
            },
        },
    },
},
```

### Handling Write Dispositions
```go
// Common write disposition validation
Validators: []validator.String{
    stringvalidator.OneOf("append", "truncate", "replace", "merge"),
},
Default: stringdefault.StaticString("append"),
```

### Connection-Specific Fields
```go
// Use conditional validation in plan modifiers
func (d *BigQueryOutputOptionPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
    // Validate that project_id and dataset are provided for BigQuery
    if req.PlanValue.IsNull() {
        return
    }
    
    var plan BigQueryOutputOption
    resp.Diagnostics.Append(req.PlanValue.As(ctx, &plan, basetypes.ObjectAsOptions{})...)
    
    if plan.ProjectID.IsNull() {
        resp.Diagnostics.AddAttributeError(
            req.Path.AtName("project_id"),
            "Missing required field",
            "project_id is required for BigQuery output options",
        )
    }
}
```

### Column Type Validation
```go
// Validate column types are supported by destination
func (d *BigQueryOutputOptionColumnPlanModifier) PlanModifyList(ctx context.Context, req planmodifier.ListRequest, resp *planmodifier.ListResponse) {
    if req.PlanValue.IsNull() {
        return
    }
    
    var columns []Column
    req.PlanValue.ElementsAs(ctx, &columns, false)
    
    validTypes := []string{"STRING", "INTEGER", "FLOAT", "BOOLEAN", "TIMESTAMP", "DATE", "JSON"}
    
    for i, col := range columns {
        if !slices.Contains(validTypes, col.Type.ValueString()) {
            resp.Diagnostics.AddAttributeError(
                req.Path.AtListIndex(i).AtName("type"),
                "Invalid column type",
                fmt.Sprintf("Column type %s is not supported for BigQuery", col.Type.ValueString()),
            )
        }
    }
}
```

## Troubleshooting

### Common Issues

1. **Column Conversion Errors**: Ensure proper handling of List types and nested objects
2. **Plan Modifier Dependencies**: Make sure object and list plan modifiers work together
3. **Connection-Specific Validation**: Use plan modifiers to enforce connection-specific requirements
4. **Write Disposition Defaults**: Ensure appropriate defaults for each connection type
5. **Column Type Validation**: Validate column types against destination capabilities

### Debug Tips

1. Use `terraform plan -out=plan.tfplan && terraform show -json plan.tfplan` to debug plan issues
2. Test column operations with various data types
3. Validate nested object serialization/deserialization
4. Test with empty column lists and single columns
5. Use plan modifiers to add helpful error messages

## Notes

- Replace all `{connection_type}` placeholders with the actual connection type name
- Replace all `{ConnectionType}` placeholders with the PascalCase version
- Replace `{input}` and `{output}` with actual source and destination types
- Follow existing code patterns in the repository
- Ensure backward compatibility when updating existing files
- Add appropriate error handling and validation for columns
- Consider performance implications of large column lists
- Document any connection-specific limitations or requirements