---
page_title: "trocco_bigquery_datamart_definition Resource - trocco"
subcategory: ""
description: |-
  The datamart definition resource allows you to create, read, update, and delete a datamart definition.
---

# trocco_bigquery_datamart_definition (Resource)

The datamart definition resource allows you to create, read, update, and delete a datamart definition.

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

- `bigquery_connection_id` (Number)
- `is_runnable_concurrently` (Boolean) Whether or not to run a job if another job with the same data mart definition is running at the time the job is run.
- `name` (String) It must be less than 256 characters
- `query` (String)
- `query_mode` (String) The following query modes are supported: `insert`, `query`

### Optional

- `before_load` (String) Valid for `insert` mode
- `clustering_fields` (List of String) Valid for `insert` mode. At most 4 fields can be specified.
- `custom_variable_settings` (Attributes List) (see [below for nested schema](#nestedatt--custom_variable_settings))
- `description` (String) It must be at least 1 character
- `destination_dataset` (String) Required for `insert` mode
- `destination_table` (String) Required for `insert` mode
- `labels` (Attributes Set) (see [below for nested schema](#nestedatt--labels))
- `location` (String) Valid for `query` mode
- `notifications` (Attributes Set) (see [below for nested schema](#nestedatt--notifications))
- `partitioning` (String) The following partitioning types are supported: `ingestion_time`, `time_unit_column`. Valid for `insert` mode
- `partitioning_field` (String) Required when `partitioning` is `time_unit_column`
- `partitioning_time` (String) The following partitioning time units are supported: `DAY`, `HOUR`, `MONTH`, `YEAR`. Valid for `insert` mode. Required when `partitioning` is set
- `resource_group_id` (Number) Resource group ID to which the datamart definition belongs
- `schedules` (Attributes Set) (see [below for nested schema](#nestedatt--schedules))
- `write_disposition` (String) The following write dispositions are supported: `append`, `truncate`. Required for `insert` mode

### Read-Only

- `id` (Number) The ID of this resource.

<a id="nestedatt--custom_variable_settings"></a>
### Nested Schema for `custom_variable_settings`

Required:

- `name` (String) It must start and end with `$`
- `type` (String) The following types are supported: `string`, `timestamp`, `timestamp_runtime`

Optional:

- `direction` (String) The following directions are supported: `ago`, `later`. Required for `timestamp` and `timestamp_runtime` types.
- `format` (String) Required for `timestamp` and `timestamp_runtime` types.
- `quantity` (Number) Required for `timestamp` and `timestamp_runtime` types.
- `time_zone` (String) Required for `timestamp` and `timestamp_runtime` types.
- `unit` (String) The following units are supported: `hour`, `date`, `month`. Required for `timestamp` and `timestamp_runtime` types.
- `value` (String) Required for `string` type.


<a id="nestedatt--labels"></a>
### Nested Schema for `labels`

Required:

- `name` (String)

Read-Only:

- `id` (Number)


<a id="nestedatt--notifications"></a>
### Nested Schema for `notifications`

Required:

- `destination_type` (String) The following destination types are supported: `slack`, `email`
- `message` (String)
- `notification_type` (String) The following notification types are supported: `job`, `record`

Optional:

- `email_id` (Number) Required when `destination_type` is `email`
- `notify_when` (String) The following notify when types are supported: `finished`, `failed`. Required for `job` notification type
- `record_count` (Number) Required for `record` notification type
- `record_operator` (String) The following record operators are supported: `above`, `below`. Required for `record` notification type
- `slack_channel_id` (Number) Required when `destination_type` is `slack`


<a id="nestedatt--schedules"></a>
### Nested Schema for `schedules`

Required:

- `frequency` (String) The following frequencies are supported: `hourly`, `daily`, `weekly`, `monthly`
- `minute` (Number)
- `time_zone` (String) Time zone to calculate the schedule time

Optional:

- `day` (Number) Required for `monthly` schedule
- `day_of_week` (Number) Sunday - Saturday is represented as 0 - 6. Required for `weekly` schedule
- `hour` (Number) Required for `daily`, `weekly`, and `monthly` schedules




## Import

Import is supported using the following syntax:

```shell
terraform import trocco_bigquery_datamart_definition.example <id>
```

