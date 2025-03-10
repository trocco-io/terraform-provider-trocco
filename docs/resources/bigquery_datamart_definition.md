---
page_title: "trocco_bigquery_datamart_definition Resource - trocco"
subcategory: ""
description: |-
  Provides a TROCCO datamart definitions for Google BigQuery resource.
---

# trocco_bigquery_datamart_definition (Resource)

Provides a TROCCO datamart definitions for Google BigQuery resource.

## Example Usage

### Minimum

```terraform
resource "trocco_bigquery_datamart_definition" "minimum" {
  name                     = "example_minimum"
  is_runnable_concurrently = false
  bigquery_connection_id   = 1
  query                    = "SELECT * FROM tables"
  query_mode               = "insert"
  destination_dataset      = "dist_datasets"
  destination_table        = "dist_tables"
  write_disposition        = "append"
}
```

### With Optional Fields

```terraform
resource "trocco_bigquery_datamart_definition" "with_optionals" {
  name                     = "example_with_optionals"
  description              = "This is an example with optional fields"
  is_runnable_concurrently = false
  resource_group_id        = 1
  custom_variable_settings = [
    {
      name  = "$string$",
      type  = "string",
      value = "foo",
    },
    {
      name      = "$timestamp$",
      type      = "timestamp",
      quantity  = 1,
      unit      = "hour",
      direction = "ago",
      format    = "%Y-%m-%d %H:%M:%S",
      time_zone = "Asia/Tokyo",
    }
  ]
  bigquery_connection_id = 1
  query                  = "SELECT * FROM tables"
  query_mode             = "insert"
  destination_dataset    = "dist_datasets"
  destination_table      = "dist_tables"
  write_disposition      = "append"
}
```

### Insert Mode

```terraform
resource "trocco_bigquery_datamart_definition" "insert_mode" {
  name                     = "example_insert_mode"
  is_runnable_concurrently = false
  bigquery_connection_id   = 1
  query                    = "SELECT * FROM tables"
  query_mode               = "insert"
  destination_dataset      = "dist_datasets"
  destination_table        = "dist_tables"
  write_disposition        = "append"
  before_load              = "DELETE FROM tables WHERE created_at < '2024-01-01'"
  partitioning             = "time_unit_column"
  partitioning_time        = "DAY"
  partitioning_field       = "created_at"
  clustering_fields        = ["id", "name"]
}
```

### Query Mode

```terraform
resource "trocco_bigquery_datamart_definition" "query_mode" {
  name                     = "example_query_mode"
  is_runnable_concurrently = false
  bigquery_connection_id   = 1
  query                    = "SELECT * FROM tables"
  query_mode               = "query"
  location                 = "asia-northeast1"
}
```

### With Schedules

```terraform
resource "trocco_bigquery_datamart_definition" "with_schedules" {
  name                     = "example_with_schedules"
  is_runnable_concurrently = false
  bigquery_connection_id   = 1
  query                    = "SELECT * FROM tables"
  query_mode               = "insert"
  destination_dataset      = "dist_datasets"
  destination_table        = "dist_tables"
  write_disposition        = "append"
  schedules = [
    {
      frequency = "hourly"
      minute    = 0
      time_zone = "Asia/Tokyo"
    },
    {
      frequency = "daily"
      hour      = 0
      minute    = 0
      time_zone = "Asia/Tokyo"
    },
    {
      frequency   = "weekly"
      day_of_week = 0
      hour        = 0
      minute      = 0
      time_zone   = "Asia/Tokyo"
    },
    {
      frequency = "monthly"
      day       = 1
      hour      = 0
      minute    = 0
      time_zone = "Asia/Tokyo"
    }
  ]
}
```

### With Notifications

```terraform
resource "trocco_bigquery_datamart_definition" "with_notifications" {
  name                     = "example_with_notifications"
  is_runnable_concurrently = false
  bigquery_connection_id   = 1
  query                    = "SELECT * FROM tables"
  query_mode               = "insert"
  destination_dataset      = "dist_datasets"
  destination_table        = "dist_tables"
  write_disposition        = "append"
  notifications = [
    {
      destination_type  = "slack"
      slack_channel_id  = 1
      notification_type = "job"
      notify_when       = "finished"
      message           = "@here Job finished."
    },
    {
      destination_type  = "email"
      email_id          = 1
      notification_type = "record"
      record_count      = 100
      record_operator   = "below"
      message           = "Record count is below 100."
    }
  ]
}
```

### With Labels

```terraform
resource "trocco_bigquery_datamart_definition" "with_labels" {
  name                     = "example_with_labels"
  is_runnable_concurrently = false
  bigquery_connection_id   = 1
  query                    = "SELECT * FROM tables"
  query_mode               = "insert"
  destination_dataset      = "dist_datasets"
  destination_table        = "dist_tables"
  write_disposition        = "append"
  labels = [
    {
      name = "test_label_1"
    },
    {
      name = "test_label_2"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `bigquery_connection_id` (Number) ID of the BigQuery connection which is used to communicate with Google BigQuery
- `is_runnable_concurrently` (Boolean) Specifies whether or not to run a job if another job with the same datamart definition is running at the time the job is run
- `name` (String) Name of the datamart definition. It must be less than 256 characters
- `query` (String) Query to be executed.
- `query_mode` (String) The following query modes are supported: `insert`, `query`. You can simply specify the query and the destination table in insert mode. In query mode, you can write and execute any DML/DDL statement

### Optional

- `before_load` (String) The query to be executed before loading the data into the destination table. Available only in `insert` mode
- `clustering_fields` (List of String) Column names to be used for clustering. At most 4 fields can be specified. Available only in `insert` mode
- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--custom_variable_settings))
- `description` (String) Description of the datamart definition. It must be at least 1 character
- `destination_dataset` (String) Destination dataset where the query result will be inserted. Required in `insert` mode
- `destination_table` (String) Destination table where the query result will be inserted. Required in `insert` mode
- `labels` (Attributes Set) Labels to be attached to the datamart definition (see [below for nested schema](#nestedatt--labels))
- `location` (String) The location where the query will be executed. If not specified, the location is automatically determined by Google BigQuery. Available only in `query` mode
- `notifications` (Attributes Set) Notifications to be attached to the datamart definition (see [below for nested schema](#nestedatt--notifications))
- `partitioning` (String) The following partitioning types are supported: `ingestion_time`, `time_unit_column`. In the case of `ingestion_time`, partitions are cut based on TROCCO's job execution time. In the case of `time_unit_column`, partitioning is done based on the reference column. Available only in `insert` mode
- `partitioning_field` (String) Column name to be used for partitioning. Required when `partitioning` is `time_unit_column`
- `partitioning_time` (String) The granularity of table partitioning. The following units are supported: `DAY`, `HOUR`, `MONTH`, `YEAR`. Required when `partitioning` is set
- `resource_group_id` (Number) ID of the resource group to which the datamart definition belongs
- `schedules` (Attributes Set) Schedules to be attached to the datamart definition (see [below for nested schema](#nestedatt--schedules))
- `write_disposition` (String) The following write dispositions are supported: `append`, `truncate`. In the case of `append`, the result of the query execution is appended after the records of the existing table. In the case of `truncate`, records in the existing table are deleted and replaced with the results of the query execution. Required in `insert` mode

### Read-Only

- `id` (Number) The ID of the datamart definition

<a id="nestedatt--custom_variable_settings"></a>
### Nested Schema for `custom_variable_settings`

Required:

- `name` (String) Custom variable name. It must start and end with `$`
- `type` (String) Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types
- `format` (String) Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types
- `quantity` (Number) Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types
- `time_zone` (String) Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types
- `unit` (String) Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types
- `value` (String) Fixed string which will replace variables at runtime. Required in `string` type


<a id="nestedatt--labels"></a>
### Nested Schema for `labels`

Required:

- `name` (String) The name of the label

Read-Only:

- `id` (Number) The ID of the label


<a id="nestedatt--notifications"></a>
### Nested Schema for `notifications`

Required:

- `destination_type` (String) Destination service where the notification will be sent. The following types are supported: `slack`, `email`
- `message` (String) The message to be sent with the notification
- `notification_type` (String) Category of condition. The following types are supported: `job`, `record`

Optional:

- `email_id` (Number) ID of the email used to send notifications. Required when `destination_type` is `email`
- `notify_when` (String) Specifies the job status that trigger a notification. The following types are supported: `finished`, `failed`. Required when `notification_type` is `job`
- `record_count` (Number) The number of records to be used for condition. Required when `notification_type` is `record`
- `record_operator` (String) Operator to be used for condition. The following operators are supported: `above`, `below`. Required when `notification_type` is `record`
- `slack_channel_id` (Number) ID of the slack channel used to send notifications. Required when `destination_type` is `slack`


<a id="nestedatt--schedules"></a>
### Nested Schema for `schedules`

Required:

- `frequency` (String) Frequency of automatic execution. The following frequencies are supported: `hourly`, `daily`, `weekly`, `monthly`
- `minute` (Number) Value of minute. Required for all schedules
- `time_zone` (String) Time zone to be used for calculation

Optional:

- `day` (Number) Value of day. Required in `monthly` schedule
- `day_of_week` (Number) Value of day of week. Sunday - Saturday is represented as 0 - 6. Required in `weekly` schedule
- `hour` (Number) Value of hour. Required in `daily`, `weekly`, and `monthly` schedules




## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import trocco_bigquery_datamart_definition (Resource). For example:

```terraform
import {
  id = 1
  to = trocco_bigquery_datamart_definition.example
}
```

Using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import):

```shell
terraform import trocco_bigquery_datamart_definition.example <id>
```

