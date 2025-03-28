---
page_title: "trocco_resource_group Resource - trocco"
subcategory: ""
description: |-
  Provides a TROCCO resource_group resource.
---

# trocco_resource_group (Resource)

Provides a TROCCO resource_group resource.

## Example Usage

```terraform
resource "trocco_resource_group" "example" {
  name        = "resource group name"
  description = "description"
  teams = [
    {
      team_id = 1
      role    = "administrator"
    },
    {
      team_id = 2
      role    = "operator"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the resource group.
- `teams` (Attributes Set) The team roles of the resource group. (see [below for nested schema](#nestedatt--teams))

### Optional

- `description` (String) The description of the resource group.

### Read-Only

- `id` (Number) The ID of the resource group.

<a id="nestedatt--teams"></a>
### Nested Schema for `teams`

Required:

- `role` (String) The role of the team. Valid values are `administrator`, `editor`, `operator`, `viewer`.
- `team_id` (Number) The team ID of the role.




## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import trocco_resource_group (Resource). For example:

```terraform
import {
  id = 1
  to = trocco_resource_group.example
}
```

Using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import):

```shell
terraform import trocco_resource.import_resource_group <resource_group_id>
```

