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

### Note

After importing a `slack_channel` resource, running `terraform plan` or `terraform apply` will result in the `webhook_url` being set to an empty string (`""`), which prevents the plan or apply from completing successfully. This happens because `webhook_url` is treated as sensitive data and is not returned by the API. Therefore, it defaults to an empty string in the state.
