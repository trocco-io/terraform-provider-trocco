## Review Steps

1. Get the changes made on the current branch with the `gh pr view` and `gh pr diff` commands
2. List review perspectives
3. Save the review perspectives to a file (`/logs/reviews/<subject>-perspectives.md`)
4. Review the changes
5. Save the review result to a file (`/logs/reviews/<subject>-result.md`)

## Basic Review Perspectives

### MUST: Check errors after calling `ElementAs()`

You MUST check errors after calling `ElementAs()`.

```go
// Good
func (r ExampleResource) Create(
    ctx context.Context,
    req resource.ReadRequest,
    resp *resource.ReadResponse,
) {
    var plan ExampleModel
    
    // ...

    diags := plan.ExampleAttribute.ElementsAs(ctx, &example, false)

    // You must check errors after calling `ElementAs()`.
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

// Bad
func (r ExampleResource) Create(
    ctx context.Context,
    req resource.ReadRequest,
    resp *resource.ReadResponse,
) {
    var plan ExampleModel
    
    // ...

    plan.ExampleAttribute.ElementsAs(ctx, &example, false)
}
```

### MUST: Check errors after　calling `GetAttribute()`

You MUST check errors after calling `GetAttribute()`.

```go
// Good
func (r ExampleResource) Read(
    ctx context.Context,
    req resource.ReadRequest,
    resp *resource.ReadResponse,
) {
    var name types.String
    diags := req.State.GetAttribute(ctx, path.Root("name"), &name)

    // You MUST check errors after calling　`GetAttribute()`.
    resp.Diagnostics.Append(diags...)
    if resp.Diagnostics.HasError() {
        return
    }
}

// Bad
func (r ExampleResource) Read(
    ctx context.Context,
    req resource.ReadRequest,
    resp *resource.ReadResponse,
) {
    var name types.String
    req.State.GetAttribute(ctx, path.Root("name"), &name)
}
```
