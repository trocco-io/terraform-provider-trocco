# Styleguide

## Naming Conventions

- Use names for import aliases and variables that are descriptive and consistent in the project

## Validations

If you need to validate a single field, use validators.

```go
schema.StringAttribute{
    Required: true,
    Validators: []validator.String{
        stringvalidator.OneOf("email", "slack_channel"),
        stringvalidator.UTF8LengthAtLeast(1),
    },
}
```

If you need to validate a single field with complex logic, create a custom validator.

```go
type ExampleValidator struct{}

func (v ExampleValidator) ValidateString(
    ctx context.Context,
    req validator.StringRequest,
    resp *validator.StringResponse,
) {
    // ...
}
```

If you need to validate consistency among multiple fields, use the `ValidateConfig()` method.

```go
func (r *ExampleResource) ValidateConfig(
    ctx context.Context,
    req resource.ValidateConfigRequest,
    resp *resource.ValidateConfigResponse,
) {
    plan := &ExampleModel{}
    if resp.Diagnostics.Append(req.Config.Get(ctx, &plan)...); resp.Diagnostics.HasError() {
        return
    }

    if !plan.Foo.IsNull() {
        if plan.Bar.IsNull() {
            resp.Diagnostics.AddError(fieldName, "bar is required when foo is set")
        }
    }
}
```
