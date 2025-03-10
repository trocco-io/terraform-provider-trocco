---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "trocco_label Resource - trocco"
subcategory: ""
description: |-
  Provides a TROCCO label resource.
---

# trocco_label (Resource)

Provides a TROCCO label resource.

## Example Usage

```terraform
resource "trocco_label" "example" {
  name        = "label name"
  description = "label description"
  color       = "#FF0000"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `color` (String) The color of the label. It must be in format #RRGGBB or #RGB.
- `name` (String) The name of the label. It must be at least 1 character.

### Optional

- `description` (String) The description of the label.

### Read-Only

- `id` (Number) The ID of the label.

## Import

Import is supported using the following syntax:

```shell
terraform import trocco_label.example <label_id>
```
