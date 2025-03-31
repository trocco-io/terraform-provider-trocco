---
page_title: "trocco_notification_destination Resource - trocco"
subcategory: ""
description: |-
  Provides a TROCCO notification destination resource.
---

# trocco_notification_destination (Resource)

Provides a TROCCO notification destination resource.

## Example Usage

### Email

```terraform
resource "trocco_notification_destination" "email" {
  type = "email"

  email_config {
    email = "notify@example.com"
  }
}
```

### SlackChannel

```terraform
resource "trocco_notification_destination" "slack_channel" {
  type = "slack_channel"

  slack_channel_config {
    channel     = "#general"
    webhook_url = "https://hooks.slack.com/services/XXXX/YYYY/ZZZZ"
  }
}
```

## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import trocco_notification_destination (Resource). For example:

**Note:** After importing a resource of `slack_channel`, if you run `terraform plan` or `terraform apply` again, Terraform will often show that the resource has changed.

```terraform
import {
  id = 1
  to = trocco_notification_destination.example
}
```

Using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import):

```shell
terraform import trocco_notification_destination.example <type>,<id>
```
