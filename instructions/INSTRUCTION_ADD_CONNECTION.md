# Instruction: Add New Connection Type

This document provides step-by-step instructions for adding CRUD (Create, Read, Update, Delete) support for a new connection type.

## Files to Modify

1. `internal/client/connection.go` - API client structs
2. `internal/provider/connection_resource.go` - Terraform resource implementation
3. `docs/resources/connection.md` - Documentation
4. `examples/resources/trocco_connection/{connection_type}.tf` - Example (new file)

---

## Step 1: Update `internal/client/connection.go`

### 1.1 Add Fields to `Connection` Struct (Response)

Add fields that the API returns:

```go
type Connection struct {
    // ... existing fields ...

    // YourConnectionType Fields
    YourField1 *string `json:"your_field_1"` // your_connection_type
    YourField2 *int64  `json:"your_field_2"` // your_connection_type
}
```

### 1.2 Add Fields to `CreateConnectionInput` Struct (Create Request)

Add fields for creating a connection:

```go
type CreateConnectionInput struct {
    // ... existing fields ...

    // YourConnectionType Fields
    YourField1 *string                   `json:"your_field_1,omitempty"` // your_connection_type
    YourField2 *parameter.NullableInt64  `json:"your_field_2,omitempty"` // your_connection_type
    SecretKey  *parameter.NullableString `json:"secret_key,omitempty"`   // your_connection_type
}
```

**Type Guidelines:**
- `string` - required fields
- `*string` - optional string fields
- `*parameter.NullableInt64` - optional int64 that can be null
- `*parameter.NullableString` - optional string that can be explicitly null
- `*parameter.NullableBool` - optional bool that can be null

### 1.3 Add Fields to `UpdateConnectionInput` Struct (Update Request)

Add same fields for updating:

```go
type UpdateConnectionInput struct {
    // ... existing fields ...

    // YourConnectionType Fields
    YourField1 *string                   `json:"your_field_1,omitempty"` // your_connection_type
    YourField2 *parameter.NullableInt64  `json:"your_field_2,omitempty"` // your_connection_type
    SecretKey  *parameter.NullableString `json:"secret_key,omitempty"`   // your_connection_type
}
```

---

## Step 2: Update `internal/provider/connection_resource.go`

### 2.1 Add to `supportedConnectionTypes` List

```go
var supportedConnectionTypes = []string{
    "bigquery",
    // ... existing types ...
    "your_connection_type", // Add here
}
```

### 2.2 Add Fields to `connectionResourceModel` Struct

```go
type connectionResourceModel struct {
    // ... existing fields ...

    // YourConnectionType Fields
    YourField1 types.String `tfsdk:"your_field_1"`
    YourField2 types.Int64  `tfsdk:"your_field_2"`
    SecretKey  types.String `tfsdk:"secret_key"`
}
```

### 2.3 Update `ToCreateConnectionInput` Method

Map Terraform model to API input (for Create):

```go
func (m *connectionResourceModel) ToCreateConnectionInput() *client.CreateConnectionInput {
    input := &client.CreateConnectionInput{
        // ... existing fields ...

        // YourConnectionType Fields
        YourField1: m.YourField1.ValueStringPointer(),
        YourField2: model.NewNullableInt64(m.YourField2),
        SecretKey:  model.NewNullableString(m.SecretKey),
    }

    // ... rest of method ...
    return input
}
```

### 2.4 Update `ToUpdateConnectionInput` Method

Map Terraform model to API input (for Update):

```go
func (m *connectionResourceModel) ToUpdateConnectionInput() *client.UpdateConnectionInput {
    input := &client.UpdateConnectionInput{
        // ... existing fields ...

        // YourConnectionType Fields
        YourField1: m.YourField1.ValueStringPointer(),
        YourField2: model.NewNullableInt64(m.YourField2),
        SecretKey:  model.NewNullableString(m.SecretKey),
    }

    // ... rest of method ...
    return input
}
```

### 2.5 Add Schema Attributes

Define Terraform resource schema:

```go
func (r *connectionResource) Schema(...) {
    resp.Schema = schema.Schema{
        Attributes: map[string]schema.Attribute{
            // ... existing attributes ...

            // YourConnectionType Fields
            "your_field_1": schema.StringAttribute{
                MarkdownDescription: "YourConnectionType: Description of field 1.",
                Optional:            true, // or Required: true
                Validators: []validator.String{
                    stringvalidator.UTF8LengthAtLeast(1),
                },
            },
            "your_field_2": schema.Int64Attribute{
                MarkdownDescription: "YourConnectionType: Description of field 2.",
                Optional:            true,
                Validators: []validator.Int64{
                    int64validator.AtLeast(1),
                },
            },
            "secret_key": schema.StringAttribute{
                MarkdownDescription: "YourConnectionType: Secret key for authentication.",
                Optional:            true,
                Sensitive:           true, // Mark sensitive fields
                Validators: []validator.String{
                    stringvalidator.UTF8LengthAtLeast(1),
                },
            },
        },
    }
}
```

### 2.6 Update `Create` Method (CREATE)

Map API response to Terraform state after creation:

```go
func (r *connectionResource) Create(...) {
    // ... existing logic ...

    newState := connectionResourceModel{
        // ... existing fields ...

        // YourConnectionType Fields
        YourField1: types.StringPointerValue(conn.YourField1),
        YourField2: types.Int64PointerValue(conn.YourField2),
        SecretKey:  plan.SecretKey, // Use plan for sensitive fields
    }

    resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
```

### 2.7 Update `Read` Method (READ)

Map API response to Terraform state when reading:

```go
func (r *connectionResource) Read(...) {
    // ... existing logic ...

    newState := connectionResourceModel{
        // ... existing fields ...

        // YourConnectionType Fields
        YourField1: types.StringPointerValue(conn.YourField1),
        YourField2: types.Int64PointerValue(conn.YourField2),
        SecretKey:  state.SecretKey, // Use state for sensitive fields
    }

    resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
```

### 2.8 Update `Update` Method (UPDATE)

Map API response to Terraform state after update:

```go
func (r *connectionResource) Update(...) {
    // ... existing logic ...

    newState := connectionResourceModel{
        // ... existing fields ...

        // YourConnectionType Fields
        YourField1: types.StringPointerValue(connection.YourField1),
        YourField2: types.Int64PointerValue(connection.YourField2),
        SecretKey:  plan.SecretKey, // Use plan for sensitive fields
    }

    resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}
```

**Note:** The `Delete` and `ImportState` methods do NOT need modification.

### 2.9 Add Validation in `ValidateConfig` Method

Validate connection configuration before apply:

```go
func (r *connectionResource) ValidateConfig(...) {
    plan := &connectionResourceModel{}
    resp.Diagnostics.Append(req.Config.Get(ctx, &plan)...)
    if resp.Diagnostics.HasError() {
        return
    }

    switch plan.ConnectionType.ValueString() {
    // ... existing cases ...

    case "your_connection_type":
        // Validate required fields
        validateRequiredString(plan.YourField1, "your_field_1", "YourConnectionType", resp)
        validateRequiredString(plan.SecretKey, "secret_key", "YourConnectionType", resp)

        // Validate conditional requirements (if applicable)
        if plan.YourField1.ValueString() == "special_mode" {
            validateRequiredInt(plan.YourField2, "your_field_2", "YourConnectionType", resp)
        }
    }
}
```

---

## Step 3: Update `docs/resources/connection.md`

Add example usage section:

```markdown
### YourConnectionType

```terraform
resource "trocco_connection" "your_connection_type" {
  connection_type = "your_connection_type"

  name        = "YourConnectionType Example"
  description = "This is a YourConnectionType connection example"

  your_field_1 = "value1"
  your_field_2 = 12345
  secret_key   = "your_secret_here"
}
```
```

**Note:** Schema documentation is auto-generated by `tfplugindocs`.

---

## Step 4: Create `examples/resources/trocco_connection/{connection_type}.tf`

Create example file: `examples/resources/trocco_connection/your_connection_type.tf`

```terraform
# Basic example
resource "trocco_connection" "your_connection_type_basic" {
  connection_type = "your_connection_type"

  name         = "YourConnectionType Example"
  description  = "This is a basic example"

  your_field_1 = "value1"
  secret_key   = "your_secret_here"
}

# Advanced example with optional fields
resource "trocco_connection" "your_connection_type_advanced" {
  connection_type = "your_connection_type"

  name              = "YourConnectionType Advanced"
  description       = "Advanced example with all fields"
  resource_group_id = 1

  your_field_1 = "value1"
  your_field_2 = 54321
  secret_key   = "your_secret_here"
}
```

---

## Step 5: Format & Test

### 5.1 Format Code

```bash
# Format Go files
golangci-lint run --fix

# Format Terraform files
terraform fmt examples/resources/trocco_connection/your_connection_type.tf
```

### 5.2 Run Tests

```bash
# Run unit tests
go test -v -cover ./...

# Run acceptance tests (if you have API access)
export TROCCO_API_KEY="your_api_key"
make testacc TESTARGS="-run TestAccConnectionResource_YourConnectionType"
```

---

## Quick Reference: Type Conversion Methods

| Terraform Type → API Type | Method |
|---------------------------|--------|
| `types.String` (required) → `string` | `.ValueString()` |
| `types.String` (optional) → `*string` | `.ValueStringPointer()` |
| `types.String` (nullable) → `*parameter.NullableString` | `model.NewNullableString(field)` |
| `types.Int64` (nullable) → `*parameter.NullableInt64` | `model.NewNullableInt64(field)` |
| `types.Bool` (nullable) → `*parameter.NullableBool` | `model.NewNullableBool(field)` |

| API Type → Terraform Type | Method |
|---------------------------|--------|
| `string` → `types.String` | `types.StringValue(value)` |
| `*string` → `types.String` | `types.StringPointerValue(ptr)` |
| `int64` → `types.Int64` | `types.Int64Value(value)` |
| `*int64` → `types.Int64` | `types.Int64PointerValue(ptr)` |

---

## Important Notes

### Sensitive Fields
- Mark as `Sensitive: true` in schema
- Use plan/state values (not API response) in Create/Read/Update
- API typically doesn't return sensitive values

### Required vs Optional Fields
- **Required**: Use `Required: true` in schema
- **Optional**: Use `Optional: true` in schema
- Validate required fields in `ValidateConfig` method

### Field Naming Convention
- **Go struct**: `PascalCase` (e.g., `YourField1`)
- **Terraform**: `snake_case` (e.g., `your_field_1`)
- **JSON API**: `snake_case` (e.g., `your_field_1`)

### Common Validators
```go
// String validators
stringvalidator.UTF8LengthAtLeast(1)
stringvalidator.UTF8LengthAtMost(255)
stringvalidator.OneOf("option1", "option2")

// Int64 validators
int64validator.AtLeast(1)
int64validator.AtMost(65535)
```

---

## Checklist

Before committing your changes:

- [ ] Updated `Connection` struct in `internal/client/connection.go`
- [ ] Updated `CreateConnectionInput` struct
- [ ] Updated `UpdateConnectionInput` struct
- [ ] Added to `supportedConnectionTypes` list
- [ ] Updated `connectionResourceModel` struct
- [ ] Updated `ToCreateConnectionInput` method
- [ ] Updated `ToUpdateConnectionInput` method
- [ ] Added schema attributes with validators
- [ ] Updated `Create` method
- [ ] Updated `Read` method
- [ ] Updated `Update` method
- [ ] Added validation in `ValidateConfig` method
- [ ] Added documentation to `docs/resources/connection.md`
- [ ] Created example file `examples/resources/trocco_connection/{type}.tf`
- [ ] Formatted code with `golangci-lint run --fix`
- [ ] Formatted Terraform files with `terraform fmt`
- [ ] Tested CRUD operations manually or with acceptance tests
