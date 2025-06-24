---
page_title: "trocco_pipeline_definition Resource - trocco"
subcategory: ""
description: |-
  Provides a TROCCO pipeline definition resource.
---

# trocco_pipeline_definition (Resource)

Provides a TROCCO pipeline definition resource.

## Example Usage

### Minimum

```terraform
resource "trocco_pipeline_definition" "minimum" {
  name = "minimum"
}
```

### Task Dependencies

```terraform
resource "trocco_pipeline_definition" "task_dependencies" {
  name = "task_dependencies"

  tasks = [
    {
      key  = "trocco_transfer_first"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 1
      }
    },
    {
      key  = "trocco_transfer_second"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 2
      }
    },
  ]


  task_dependencies = [
    {
      source      = "trocco_transfer_first"
      destination = "trocco_transfer_second"
    },
  ]
}
```

### Labels

```terraform
resource "trocco_pipeline_definition" "labels" {
  name = "labels"

  labels = [
    "foo",
    "bar",
  ]
}
```

### Notifications

```terraform
resource "trocco_pipeline_definition" "notifications" {
  name = "notifications"

  notifications = [
    {
      type             = "job_execution"
      destination_type = "slack"
      notify_when      = "finished"

      slack_config = {
        notification_id = 1
        message         = "The quick brown fox jumps over the lazy dog."
      }
    },
    {
      type             = "job_time_alert"
      destination_type = "email"
      time             = 5

      email_config = {
        notification_id = 1
        message         = "The quick brown fox jumps over the lazy dog."
      }
    },
  ]
}
```

### Tasks

#### TROCCO Transfer

```terraform
resource "trocco_pipeline_definition" "trocco_transfer" {
  name = "trocco_transfer"

  tasks = [
    {
      key  = "trocco_transfer"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 1
      }
    }
  ]
}
```

##### Custom Variable Loop Examples

```terraform
resource "trocco_pipeline_definition" "trocco_transfer_with_custom_variable_loop_with_bigquery_config" {
  name = "trocco_transfer_with_custom_variable_loop_with_bigquery_config"

  tasks = [
    {
      key  = "trocco_transfer"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 1
        custom_variable_loop = {
          type = "bigquery"
          bigquery_config = {
            connection_id = 1
            query         = "select foo, bar from sample"
            variables = [
              "$foo$",
              "$bar$"
            ]
          }
        }
      }
    }
  ]
}
```

```terraform
resource "trocco_pipeline_definition" "trocco_transfer_with_custom_variable_loop_with_period_config" {
  name = "trocco_transfer_with_custom_variable_loop_with_period_config"

  tasks = [
    {
      key  = "trocco_transfer"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 1
        custom_variable_loop = {
          type = "period"
          period_config = {
            interval  = "day"
            time_zone = "Asia/Tokyo"
            from = {
              value = 7
              unit  = "day"
            }
            to = {
              value = 1
              unit  = "day"
            }
            variables = [
              {
                name = "$date$"
                offset = {
                  value = 0
                  unit  = "day"
                }
              }
            ]
          }
        }
      }
    }
  ]
}
```

```terraform
resource "trocco_pipeline_definition" "trocco_transfer_with_custom_variable_loop_with_redshift_config" {
  name = "trocco_transfer_with_custom_variable_loop_with_redshift_config"

  tasks = [
    {
      key  = "trocco_transfer"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 1
        custom_variable_loop = {
          type = "redshift"
          redshift_config = {
            connection_id = 1
            query         = "select foo, bar from sample"
            variables = [
              "$foo$",
              "$bar$"
            ]
            database = "dev"
          }
        }
      }
    }
  ]
}
```

```terraform
resource "trocco_pipeline_definition" "trocco_transfer_with_custom_variable_loop_with_snowflake_config" {
  name = "trocco_transfer_with_custom_variable_loop_with_snowflake_config"

  tasks = [
    {
      key  = "trocco_transfer"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 1
        custom_variable_loop = {
          type = "snowflake"
          snowflake_config = {
            connection_id = 1
            query         = "select foo, bar from sample"
            variables = [
              "$foo$",
              "$bar$"
            ]
            warehouse = "COMPUTE_WH"
          }
        }
      }
    }
  ]
}
```

```terraform
resource "trocco_pipeline_definition" "trocco_transfer_with_custom_variable_loop_with_string_config" {
  name = "trocco_transfer_with_custom_variable_loop_with_string_config"

  tasks = [
    {
      key  = "trocco_transfer"
      type = "trocco_transfer"

      trocco_transfer_config = {
        definition_id = 1

        custom_variable_loop = {
          type = "string"
          string_config = {
            variables = [
              {
                name   = "$foo$"
                values = ["a", "b"]
              },
              {
                name = "$bar$"
                values : ["", "c"]
              }
            ]
          }
        }
      }
    }
  ]
}
```


#### TROCCO Transfer Bulk

```terraform
resource "trocco_pipeline_definition" "trocco_transfer_bulk" {
  name = "trocco_transfer_bulk"

  tasks = [
    {
      key  = "trocco_transfer_bulk"
      type = "trocco_transfer_bulk"

      trocco_transfer_bulk_config = {
        definition_id                 = 1
        is_parallel_execution_allowed = false
        is_stopped_on_errors          = true
        max_errors                    = 1
      }
    }
  ]
}
```

#### Datamart

##### Azure Synapse Analytics

```terraform
resource "trocco_pipeline_definition" "trocco_azure_synapse_analytics_datamart" {
  name = "trocco_azure_synapse_analytics_datamart"

  tasks = [
    {
      key  = "trocco_azure_synapse_analytics_datamart"
      type = "trocco_azure_synapse_analytics_datamart"

      trocco_azure_synapse_analytics_datamart_config = {
        definition_id = 26
      }
    }
  ]
}
```

##### BigQuery

```terraform
resource "trocco_pipeline_definition" "trocco_bigquery_datamart" {
  name = "trocco_bigquery_datamart"

  tasks = [
    {
      key  = "trocco_bigquery_datamart"
      type = "trocco_bigquery_datamart"

      trocco_bigquery_datamart_config = {
        definition_id = 1
      }
    }
  ]
}
```

##### Redshift

```terraform
resource "trocco_pipeline_definition" "trocco_redshift_datamart" {
  name = "trocco_redshift_datamart"

  tasks = [
    {
      key  = "trocco_redshift_datamart"
      type = "trocco_redshift_datamart"

      trocco_redshift_datamart_config = {
        definition_id = 1
      }
    }
  ]
}
```

##### Snowflake

```terraform
resource "trocco_pipeline_definition" "trocco_snowflake_datamart" {
  name = "trocco_snowflake_datamart"

  tasks = [
    {
      key  = "trocco_snowflake_datamart"
      type = "trocco_snowflake_datamart"

      trocco_snowflake_datamart_config = {
        definition_id = 1
      }
    }
  ]
}
```

#### Data Check

##### BigQuery

```terraform
resource "trocco_pipeline_definition" "bigquery_data_check" {
  name = "bigquery_data_check"

  tasks = [
    {
      key  = "bigquery_data_check"
      type = "bigquery_data_check"

      bigquery_data_check_config = {
        name          = "Example"
        connection_id = 1
        query         = "SELECT COUNT(id) FROM examples"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
      }
    }
  ]
}

resource "trocco_pipeline_definition" "bigquery_data_check_with_custom_variables" {
  name = "bigquery_data_check_with_custom_variables"

  tasks = [
    {
      key  = "bigquery_data_check_with_custom_variables"
      type = "bigquery_data_check"

      bigquery_data_check_config = {
        name          = "Example"
        connection_id = 1
        query         = "SELECT COUNT(id) FROM examples"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false

        custom_variables = [
          {
            name  = "$string$"
            type  = "string"
            value = "foo"
          },
          {
            name      = "$timestamp$"
            type      = "timestamp"
            quantity  = 1,
            unit      = "hour"
            direction = "ago"
            format    = "%Y-%m-%d %H:%M:%S"
            time_zone = "Asia/Tokyo"
          },
        ]
      }
    }
  ]
}
```

##### Redshift

```terraform
resource "trocco_pipeline_definition" "redshift_data_check" {
  name = "redshift_data_check"

  tasks = [
    {
      key  = "redshift_data_check"
      type = "redshift_data_check"

      redshift_data_check_config = {
        name          = "Example"
        connection_id = 1
        query         = "SELECT COUNT(id) FROM examples"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
        database      = "EXAMPLE"
      }
    }
  ]
}

resource "trocco_pipeline_definition" "redshift_data_check_with_custom_variables" {
  name = "redshift_data_check_with_custom_variables"

  tasks = [
    {
      key  = "redshift_data_check_with_custom_variables"
      type = "redshift_data_check"

      redshift_data_check_config = {
        name          = "Example"
        connection_id = 1
        query         = "SELECT COUNT(id) FROM examples"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false

        custom_variables = [
          {
            name  = "$string$"
            type  = "string"
            value = "foo"
          },
          {
            name      = "$timestamp$"
            type      = "timestamp"
            quantity  = 1,
            unit      = "hour"
            direction = "ago"
            format    = "%Y-%m-%d %H:%M:%S"
            time_zone = "Asia/Tokyo"
          },
        ]
      }
    }
  ]
}
```

##### Snowflake

```terraform
resource "trocco_pipeline_definition" "snowflake_data_check" {
  name = "snowflake_data_check"

  tasks = [
    {
      key  = "snowflake_data_check"
      type = "snowflake_data_check"

      snowflake_data_check_config = {
        name          = "Example"
        connection_id = 1
        query         = "SELECT COUNT(id) FROM examples"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false
        warehouse     = "EXAMPLE"
      }
    }
  ]
}

resource "trocco_pipeline_definition" "snowflake_data_check_with_custom_variables" {
  name = "snowflake_data_check_with_custom_variables"

  tasks = [
    {
      key  = "snowflake_data_check_with_custom_variables"
      type = "snowflake_data_check"

      snowflake_data_check_config = {
        name          = "Example"
        connection_id = 1
        query         = "SELECT COUNT(id) FROM examples"
        operator      = "equal"
        query_result  = 1
        accepts_null  = false

        custom_variables = [
          {
            name  = "$string$"
            type  = "string"
            value = "foo"
          },
          {
            name      = "$timestamp$"
            type      = "timestamp"
            quantity  = 1,
            unit      = "hour"
            direction = "ago"
            format    = "%Y-%m-%d %H:%M:%S"
            time_zone = "Asia/Tokyo"
          },
        ]
      }
    }
  ]
}
```

#### HTTP Request

```terraform
resource "trocco_pipeline_definition" "http_request" {
  name = "http_request"

  tasks = [
    {
      key  = "http_request"
      type = "http_request"

      http_request_config = {
        name = "Example"

        http_method = "GET"

        url = "https://example.com"

        request_headers = [
          { key : "Authorization", value : "Bearer example", masking : true },
          { key : "Content-Type", value : "application/json", masking : false },
        ]

        request_parameters = [
          { key : "foo", value : "bar", masking : true },
        ]
      }
    }
  ]
}
```

#### Slack Notify

```terraform
resource "trocco_pipeline_definition" "slack_notify" {
  name = "slack_notify"

  tasks = [
    {
      key  = "slack_notify"
      type = "slack_notify"

      slack_notification_config = {
        name          = "Example"
        connection_id = 1
        message       = "The quick brown fox jumps over the lazy dog."
        ignore_error  = false
      }
    }
  ]
}
```

#### Tableau Extract

```terraform
resource "trocco_pipeline_definition" "tableau_extract" {
  name = "tableau_extract"

  tasks = [
    {
      key  = "tableau_extract"
      type = "tableau_extract"

      tableau_data_extraction_config = {
        name          = "Example"
        connection_id = 1
        task_id       = "57f1fdc6-aef7-4d4d-a38a-3b73ed529de3"
      }
    }
  ]
}
```

#### dbt

```terraform
resource "trocco_pipeline_definition" "trocco_dbt" {
  name = "trocco_dbt"

  tasks = [
    {
      key  = "trocco_dbt"
      type = "trocco_dbt"

      trocco_dbt_config = {
        definition_id = 1
      }
    }
  ]
}
```

#### TROCCO Pipeline

```terraform
resource "trocco_pipeline_definition" "trocco_pipeline" {
  name = "trocco_pipeline"

  tasks = [
    {
      key  = "trocco_pipeline"
      type = "trocco_pipeline"

      trocco_pipeline_config = {
        definition_id = 1
      }
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) The name of the pipeline definition

### Optional

- `description` (String) The description of the pipeline definition
- `execution_timeout` (Number) The maximum time in minutes that the pipeline can run
- `is_concurrent_execution_skipped` (Boolean) Weather to skip execution of the pipeline if it is already running
- `is_stopped_on_errors` (Boolean) Weather to stop the pipeline if any task fails
- `labels` (Set of String) The labels of the pipeline definition
- `max_retries` (Number) The maximum number of retries that the pipeline can have
- `max_task_parallelism` (Number) The maximum number of tasks that the pipeline can run in parallel
- `min_retry_interval` (Number) The minimum time in minutes between retries
- `notifications` (Attributes Set) The notifications of the pipeline definition (see [below for nested schema](#nestedatt--notifications))
- `resource_group_id` (Number) The resource group ID of the pipeline definition
- `schedules` (Attributes Set) The schedules of the pipeline definition (see [below for nested schema](#nestedatt--schedules))
- `task_dependencies` (Attributes Set) The task dependencies of the workflow. (see [below for nested schema](#nestedatt--task_dependencies))
- `tasks` (Attributes List) The tasks of the workflow. (see [below for nested schema](#nestedatt--tasks))

### Read-Only

- `id` (Number) The ID of the pipeline definition

<a id="nestedatt--notifications"></a>
### Nested Schema for `notifications`

Required:

- `destination_type` (String) The destination type of the notification
- `type` (String) The type of the notification

Optional:

- `email_config` (Attributes) The email configuration of the notification (see [below for nested schema](#nestedatt--notifications--email_config))
- `notify_when` (String) When to notify
- `slack_config` (Attributes) The slack configuration of the notification (see [below for nested schema](#nestedatt--notifications--slack_config))
- `time` (Number) The time of the notification

<a id="nestedatt--notifications--email_config"></a>
### Nested Schema for `notifications.email_config`

Required:

- `message` (String) The message of the notification
- `notification_id` (Number) The notification id


<a id="nestedatt--notifications--slack_config"></a>
### Nested Schema for `notifications.slack_config`

Required:

- `message` (String) The message of the notification
- `notification_id` (Number) The notification id



<a id="nestedatt--schedules"></a>
### Nested Schema for `schedules`

Required:

- `frequency` (String) The frequency of the schedule
- `minute` (Number) The minute of the schedule
- `time_zone` (String) The time zone of the schedule

Optional:

- `day` (Number) The day of the schedule
- `day_of_week` (Number) The day of the week of the schedule
- `hour` (Number) The hour of the schedule


<a id="nestedatt--task_dependencies"></a>
### Nested Schema for `task_dependencies`

Required:

- `destination` (String) The destination task key.
- `source` (String) The source task key.


<a id="nestedatt--tasks"></a>
### Nested Schema for `tasks`

Required:

- `key` (String) The key of the task.
- `type` (String) The type of the task.

Optional:

- `bigquery_data_check_config` (Attributes) The datacheck task config of the pipeline definition (see [below for nested schema](#nestedatt--tasks--bigquery_data_check_config))
- `http_request_config` (Attributes) The task configuration for the HTTP request task. (see [below for nested schema](#nestedatt--tasks--http_request_config))
- `redshift_data_check_config` (Attributes) The task configuration for the datacheck task. (see [below for nested schema](#nestedatt--tasks--redshift_data_check_config))
- `slack_notification_config` (Attributes) The task configuration for the slack notification task. (see [below for nested schema](#nestedatt--tasks--slack_notification_config))
- `snowflake_data_check_config` (Attributes) The task configuration for the datacheck task. (see [below for nested schema](#nestedatt--tasks--snowflake_data_check_config))
- `tableau_data_extraction_config` (Attributes) The task configuration for the tableau data extraction task. (see [below for nested schema](#nestedatt--tasks--tableau_data_extraction_config))
- `trocco_azure_synapse_analytics_datamart_config` (Attributes) The task configuration for the datamart task. (see [below for nested schema](#nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config))
- `trocco_bigquery_datamart_config` (Attributes) The task configuration for the datamart task. (see [below for nested schema](#nestedatt--tasks--trocco_bigquery_datamart_config))
- `trocco_dbt_config` (Attributes) The task configuration for the trocco dbt task. (see [below for nested schema](#nestedatt--tasks--trocco_dbt_config))
- `trocco_pipeline_config` (Attributes) The task configuration for the trocco pipeline task. (see [below for nested schema](#nestedatt--tasks--trocco_pipeline_config))
- `trocco_redshift_datamart_config` (Attributes) The task configuration for the datamart task. (see [below for nested schema](#nestedatt--tasks--trocco_redshift_datamart_config))
- `trocco_snowflake_datamart_config` (Attributes) The task configuration for the datamart task. (see [below for nested schema](#nestedatt--tasks--trocco_snowflake_datamart_config))
- `trocco_transfer_bulk_config` (Attributes) The task configuration for the trocco transfer bulk task. (see [below for nested schema](#nestedatt--tasks--trocco_transfer_bulk_config))
- `trocco_transfer_config` (Attributes) The task configuration for the trocco transfer task. (see [below for nested schema](#nestedatt--tasks--trocco_transfer_config))

Read-Only:

- `task_identifier` (Number) The task identifier.

<a id="nestedatt--tasks--bigquery_data_check_config"></a>
### Nested Schema for `tasks.bigquery_data_check_config`

Required:

- `connection_id` (Number) The connection id of the datacheck task
- `name` (String) The name of the datacheck task

Optional:

- `accepts_null` (Boolean) Whether the datacheck task accepts null
- `custom_variables` (Attributes Set) The custom variables of the pipeline definition (see [below for nested schema](#nestedatt--tasks--bigquery_data_check_config--custom_variables))
- `operator` (String) The operator of the datacheck task
- `query` (String) The query of the datacheck task
- `query_result` (Number) The query result of the datacheck task

<a id="nestedatt--tasks--bigquery_data_check_config--custom_variables"></a>
### Nested Schema for `tasks.bigquery_data_check_config.custom_variables`

Required:

- `name` (String) The name of the custom variable
- `type` (String) The type of the custom variable

Optional:

- `direction` (String) The direction of the custom variable
- `format` (String) The format of the custom variable
- `quantity` (Number) The quantity of the custom variable
- `time_zone` (String) The time zone of the custom variable
- `unit` (String) The unit of the custom variable
- `value` (String) The value of the custom variable



<a id="nestedatt--tasks--http_request_config"></a>
### Nested Schema for `tasks.http_request_config`

Required:

- `http_method` (String) The HTTP method to use for the request
- `name` (String) The name of the task
- `url` (String) The URL to send the request to

Optional:

- `connection_id` (Number) The connection id to use for the task
- `custom_variables` (Attributes Set) The custom variables of the pipeline definition (see [below for nested schema](#nestedatt--tasks--http_request_config--custom_variables))
- `request_body` (String) The body of the request
- `request_headers` (Attributes List) The headers to send with the request (see [below for nested schema](#nestedatt--tasks--http_request_config--request_headers))
- `request_parameters` (Attributes List) (see [below for nested schema](#nestedatt--tasks--http_request_config--request_parameters))

<a id="nestedatt--tasks--http_request_config--custom_variables"></a>
### Nested Schema for `tasks.http_request_config.custom_variables`

Required:

- `name` (String) The name of the custom variable
- `type` (String) The type of the custom variable

Optional:

- `direction` (String) The direction of the custom variable
- `format` (String) The format of the custom variable
- `quantity` (Number) The quantity of the custom variable
- `time_zone` (String) The time zone of the custom variable
- `unit` (String) The unit of the custom variable
- `value` (String) The value of the custom variable


<a id="nestedatt--tasks--http_request_config--request_headers"></a>
### Nested Schema for `tasks.http_request_config.request_headers`

Required:

- `key` (String) The key of the header
- `value` (String, Sensitive) The value of the header

Optional:

- `masking` (Boolean) Whether to mask the value of the header


<a id="nestedatt--tasks--http_request_config--request_parameters"></a>
### Nested Schema for `tasks.http_request_config.request_parameters`

Required:

- `key` (String) The key of the parameter
- `value` (String, Sensitive) The value of the parameter

Optional:

- `masking` (Boolean) Whether to mask the value of the parameter



<a id="nestedatt--tasks--redshift_data_check_config"></a>
### Nested Schema for `tasks.redshift_data_check_config`

Required:

- `connection_id` (Number) The connection id to use for the datacheck task
- `name` (String) The name of the datacheck task

Optional:

- `accepts_null` (Boolean) Whether the datacheck task accepts null values
- `custom_variables` (Attributes Set) The custom variables of the pipeline definition (see [below for nested schema](#nestedatt--tasks--redshift_data_check_config--custom_variables))
- `database` (String) The database to use for the datacheck task
- `operator` (String) The operator to use for the datacheck task
- `query` (String) The query to run for the datacheck task
- `query_result` (Number) The query result to use for the datacheck task

<a id="nestedatt--tasks--redshift_data_check_config--custom_variables"></a>
### Nested Schema for `tasks.redshift_data_check_config.custom_variables`

Required:

- `name` (String) The name of the custom variable
- `type` (String) The type of the custom variable

Optional:

- `direction` (String) The direction of the custom variable
- `format` (String) The format of the custom variable
- `quantity` (Number) The quantity of the custom variable
- `time_zone` (String) The time zone of the custom variable
- `unit` (String) The unit of the custom variable
- `value` (String) The value of the custom variable



<a id="nestedatt--tasks--slack_notification_config"></a>
### Nested Schema for `tasks.slack_notification_config`

Required:

- `connection_id` (Number) The connection id to use for the task
- `ignore_error` (Boolean) Whether to ignore errors
- `message` (String) The message to send
- `name` (String) The name of the task


<a id="nestedatt--tasks--snowflake_data_check_config"></a>
### Nested Schema for `tasks.snowflake_data_check_config`

Required:

- `connection_id` (Number) The connection id to use for the datacheck task
- `name` (String) The name of the datacheck task

Optional:

- `accepts_null` (Boolean) Whether the datacheck task accepts null values
- `custom_variables` (Attributes Set) The custom variables of the pipeline definition (see [below for nested schema](#nestedatt--tasks--snowflake_data_check_config--custom_variables))
- `operator` (String) The operator to use for the datacheck task
- `query` (String) The query to run for the datacheck task
- `query_result` (Number) The query result to use for the datacheck task
- `warehouse` (String) The warehouse to use for the datacheck task

<a id="nestedatt--tasks--snowflake_data_check_config--custom_variables"></a>
### Nested Schema for `tasks.snowflake_data_check_config.custom_variables`

Required:

- `name` (String) The name of the custom variable
- `type` (String) The type of the custom variable

Optional:

- `direction` (String) The direction of the custom variable
- `format` (String) The format of the custom variable
- `quantity` (Number) The quantity of the custom variable
- `time_zone` (String) The time zone of the custom variable
- `unit` (String) The unit of the custom variable
- `value` (String) The value of the custom variable



<a id="nestedatt--tasks--tableau_data_extraction_config"></a>
### Nested Schema for `tasks.tableau_data_extraction_config`

Required:

- `connection_id` (Number) The connection id to use for the task
- `name` (String) The name of the task
- `task_id` (String) The Tableau task ID. You can get with the [Tableau API](https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref.htm#list_extract_refresh_tasks1).


<a id="nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config"></a>
### Nested Schema for `tasks.trocco_azure_synapse_analytics_datamart_config`

Required:

- `definition_id` (Number) The definition id to use for the datamart task

Optional:

- `custom_variable_loop` (Attributes) The custom variable loop of the pipeline definition (see [below for nested schema](#nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop))

<a id="nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop"></a>
### Nested Schema for `tasks.trocco_azure_synapse_analytics_datamart_config.custom_variable_loop`

Required:

- `type` (String) The type of the custom variable loop. Allowed values: "string", "period", "bigquery", "snowflake", "redshift".

Optional:

- `bigquery_config` (Attributes) BigQuery custom variabe loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--bigquery_config))
- `is_parallel_execution_allowed` (Boolean) Whether parallel execution is allowed
- `is_stopped_on_errors` (Boolean) Whether the loop is stopped on errors
- `max_errors` (Number) The maximum number of errors
- `period_config` (Attributes) Period custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--period_config))
- `redshift_config` (Attributes) Redshift custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--redshift_config))
- `snowflake_config` (Attributes) Snowflake custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--snowflake_config))
- `string_config` (Attributes) String custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--string_config))

<a id="nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--bigquery_config"></a>
### Nested Schema for `tasks.trocco_azure_synapse_analytics_datamart_config.custom_variable_loop.bigquery_config`

Required:

- `connection_id` (Number) BigQuery connection ID
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded


<a id="nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--period_config"></a>
### Nested Schema for `tasks.trocco_azure_synapse_analytics_datamart_config.custom_variable_loop.period_config`

Required:

- `from` (Attributes) Start of the loop (see [below for nested schema](#nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--period_config--from))
- `interval` (String) Interval of the loop
- `time_zone` (String) Timezone of the configuration
- `to` (Attributes) End of the loop (see [below for nested schema](#nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--period_config--to))
- `variables` (Attributes List) Custom variables to be expanded (see [below for nested schema](#nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--period_config--variables))

<a id="nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--period_config--from"></a>
### Nested Schema for `tasks.trocco_azure_synapse_analytics_datamart_config.custom_variable_loop.period_config.from`

Required:

- `unit` (String) Unit
- `value` (Number) Value


<a id="nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--period_config--to"></a>
### Nested Schema for `tasks.trocco_azure_synapse_analytics_datamart_config.custom_variable_loop.period_config.to`

Required:

- `unit` (String) Unit
- `value` (Number) Value


<a id="nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--period_config--variables"></a>
### Nested Schema for `tasks.trocco_azure_synapse_analytics_datamart_config.custom_variable_loop.period_config.variables`

Required:

- `name` (String) Name of custom variable
- `offset` (Attributes) Offset on custom variable expanded (see [below for nested schema](#nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--period_config--variables--offset))

<a id="nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--period_config--variables--offset"></a>
### Nested Schema for `tasks.trocco_azure_synapse_analytics_datamart_config.custom_variable_loop.period_config.variables.offset`

Required:

- `unit` (String) Unit
- `value` (Number) Value




<a id="nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--redshift_config"></a>
### Nested Schema for `tasks.trocco_azure_synapse_analytics_datamart_config.custom_variable_loop.redshift_config`

Required:

- `connection_id` (Number) Redshift connection ID
- `database` (String) Redshift database
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded


<a id="nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--snowflake_config"></a>
### Nested Schema for `tasks.trocco_azure_synapse_analytics_datamart_config.custom_variable_loop.snowflake_config`

Required:

- `connection_id` (Number) Snowflake connection ID
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded
- `warehouse` (String) Snowflake warehouse


<a id="nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--string_config"></a>
### Nested Schema for `tasks.trocco_azure_synapse_analytics_datamart_config.custom_variable_loop.string_config`

Required:

- `variables` (Attributes List) Custom variables (see [below for nested schema](#nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--string_config--variables))

<a id="nestedatt--tasks--trocco_azure_synapse_analytics_datamart_config--custom_variable_loop--string_config--variables"></a>
### Nested Schema for `tasks.trocco_azure_synapse_analytics_datamart_config.custom_variable_loop.string_config.variables`

Required:

- `name` (String) Custom variable name
- `values` (List of String) Custom variable values





<a id="nestedatt--tasks--trocco_bigquery_datamart_config"></a>
### Nested Schema for `tasks.trocco_bigquery_datamart_config`

Required:

- `definition_id` (Number) The definition id to use for the datamart task

Optional:

- `custom_variable_loop` (Attributes) The custom variable loop of the pipeline definition (see [below for nested schema](#nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop))

<a id="nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop"></a>
### Nested Schema for `tasks.trocco_bigquery_datamart_config.custom_variable_loop`

Required:

- `type` (String) The type of the custom variable loop. Allowed values: "string", "period", "bigquery", "snowflake", "redshift".

Optional:

- `bigquery_config` (Attributes) BigQuery custom variabe loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--bigquery_config))
- `is_parallel_execution_allowed` (Boolean) Whether parallel execution is allowed
- `is_stopped_on_errors` (Boolean) Whether the loop is stopped on errors
- `max_errors` (Number) The maximum number of errors
- `period_config` (Attributes) Period custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--period_config))
- `redshift_config` (Attributes) Redshift custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--redshift_config))
- `snowflake_config` (Attributes) Snowflake custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--snowflake_config))
- `string_config` (Attributes) String custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--string_config))

<a id="nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--bigquery_config"></a>
### Nested Schema for `tasks.trocco_bigquery_datamart_config.custom_variable_loop.bigquery_config`

Required:

- `connection_id` (Number) BigQuery connection ID
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded


<a id="nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--period_config"></a>
### Nested Schema for `tasks.trocco_bigquery_datamart_config.custom_variable_loop.period_config`

Required:

- `from` (Attributes) Start of the loop (see [below for nested schema](#nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--period_config--from))
- `interval` (String) Interval of the loop
- `time_zone` (String) Timezone of the configuration
- `to` (Attributes) End of the loop (see [below for nested schema](#nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--period_config--to))
- `variables` (Attributes List) Custom variables to be expanded (see [below for nested schema](#nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--period_config--variables))

<a id="nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--period_config--from"></a>
### Nested Schema for `tasks.trocco_bigquery_datamart_config.custom_variable_loop.period_config.from`

Required:

- `unit` (String) Unit
- `value` (Number) Value


<a id="nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--period_config--to"></a>
### Nested Schema for `tasks.trocco_bigquery_datamart_config.custom_variable_loop.period_config.to`

Required:

- `unit` (String) Unit
- `value` (Number) Value


<a id="nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--period_config--variables"></a>
### Nested Schema for `tasks.trocco_bigquery_datamart_config.custom_variable_loop.period_config.variables`

Required:

- `name` (String) Name of custom variable
- `offset` (Attributes) Offset on custom variable expanded (see [below for nested schema](#nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--period_config--variables--offset))

<a id="nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--period_config--variables--offset"></a>
### Nested Schema for `tasks.trocco_bigquery_datamart_config.custom_variable_loop.period_config.variables.offset`

Required:

- `unit` (String) Unit
- `value` (Number) Value




<a id="nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--redshift_config"></a>
### Nested Schema for `tasks.trocco_bigquery_datamart_config.custom_variable_loop.redshift_config`

Required:

- `connection_id` (Number) Redshift connection ID
- `database` (String) Redshift database
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded


<a id="nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--snowflake_config"></a>
### Nested Schema for `tasks.trocco_bigquery_datamart_config.custom_variable_loop.snowflake_config`

Required:

- `connection_id` (Number) Snowflake connection ID
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded
- `warehouse` (String) Snowflake warehouse


<a id="nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--string_config"></a>
### Nested Schema for `tasks.trocco_bigquery_datamart_config.custom_variable_loop.string_config`

Required:

- `variables` (Attributes List) Custom variables (see [below for nested schema](#nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--string_config--variables))

<a id="nestedatt--tasks--trocco_bigquery_datamart_config--custom_variable_loop--string_config--variables"></a>
### Nested Schema for `tasks.trocco_bigquery_datamart_config.custom_variable_loop.string_config.variables`

Required:

- `name` (String) Custom variable name
- `values` (List of String) Custom variable values





<a id="nestedatt--tasks--trocco_dbt_config"></a>
### Nested Schema for `tasks.trocco_dbt_config`

Required:

- `definition_id` (Number) The definition id to use for the trocco dbt task


<a id="nestedatt--tasks--trocco_pipeline_config"></a>
### Nested Schema for `tasks.trocco_pipeline_config`

Required:

- `definition_id` (Number) The definition id to use for the trocco pipeline task

Optional:

- `custom_variable_loop` (Attributes) The custom variable loop of the pipeline definition (see [below for nested schema](#nestedatt--tasks--trocco_pipeline_config--custom_variable_loop))

<a id="nestedatt--tasks--trocco_pipeline_config--custom_variable_loop"></a>
### Nested Schema for `tasks.trocco_pipeline_config.custom_variable_loop`

Required:

- `type` (String) The type of the custom variable loop. Allowed values: "string", "period", "bigquery", "snowflake", "redshift".

Optional:

- `bigquery_config` (Attributes) BigQuery custom variabe loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--bigquery_config))
- `is_parallel_execution_allowed` (Boolean) Whether parallel execution is allowed
- `is_stopped_on_errors` (Boolean) Whether the loop is stopped on errors
- `max_errors` (Number) The maximum number of errors
- `period_config` (Attributes) Period custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--period_config))
- `redshift_config` (Attributes) Redshift custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--redshift_config))
- `snowflake_config` (Attributes) Snowflake custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--snowflake_config))
- `string_config` (Attributes) String custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--string_config))

<a id="nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--bigquery_config"></a>
### Nested Schema for `tasks.trocco_pipeline_config.custom_variable_loop.bigquery_config`

Required:

- `connection_id` (Number) BigQuery connection ID
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded


<a id="nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--period_config"></a>
### Nested Schema for `tasks.trocco_pipeline_config.custom_variable_loop.period_config`

Required:

- `from` (Attributes) Start of the loop (see [below for nested schema](#nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--period_config--from))
- `interval` (String) Interval of the loop
- `time_zone` (String) Timezone of the configuration
- `to` (Attributes) End of the loop (see [below for nested schema](#nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--period_config--to))
- `variables` (Attributes List) Custom variables to be expanded (see [below for nested schema](#nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--period_config--variables))

<a id="nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--period_config--from"></a>
### Nested Schema for `tasks.trocco_pipeline_config.custom_variable_loop.period_config.from`

Required:

- `unit` (String) Unit
- `value` (Number) Value


<a id="nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--period_config--to"></a>
### Nested Schema for `tasks.trocco_pipeline_config.custom_variable_loop.period_config.to`

Required:

- `unit` (String) Unit
- `value` (Number) Value


<a id="nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--period_config--variables"></a>
### Nested Schema for `tasks.trocco_pipeline_config.custom_variable_loop.period_config.variables`

Required:

- `name` (String) Name of custom variable
- `offset` (Attributes) Offset on custom variable expanded (see [below for nested schema](#nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--period_config--variables--offset))

<a id="nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--period_config--variables--offset"></a>
### Nested Schema for `tasks.trocco_pipeline_config.custom_variable_loop.period_config.variables.offset`

Required:

- `unit` (String) Unit
- `value` (Number) Value




<a id="nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--redshift_config"></a>
### Nested Schema for `tasks.trocco_pipeline_config.custom_variable_loop.redshift_config`

Required:

- `connection_id` (Number) Redshift connection ID
- `database` (String) Redshift database
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded


<a id="nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--snowflake_config"></a>
### Nested Schema for `tasks.trocco_pipeline_config.custom_variable_loop.snowflake_config`

Required:

- `connection_id` (Number) Snowflake connection ID
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded
- `warehouse` (String) Snowflake warehouse


<a id="nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--string_config"></a>
### Nested Schema for `tasks.trocco_pipeline_config.custom_variable_loop.string_config`

Required:

- `variables` (Attributes List) Custom variables (see [below for nested schema](#nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--string_config--variables))

<a id="nestedatt--tasks--trocco_pipeline_config--custom_variable_loop--string_config--variables"></a>
### Nested Schema for `tasks.trocco_pipeline_config.custom_variable_loop.string_config.variables`

Required:

- `name` (String) Custom variable name
- `values` (List of String) Custom variable values





<a id="nestedatt--tasks--trocco_redshift_datamart_config"></a>
### Nested Schema for `tasks.trocco_redshift_datamart_config`

Required:

- `definition_id` (Number) The definition id to use for the datamart task

Optional:

- `custom_variable_loop` (Attributes) The custom variable loop of the pipeline definition (see [below for nested schema](#nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop))

<a id="nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop"></a>
### Nested Schema for `tasks.trocco_redshift_datamart_config.custom_variable_loop`

Required:

- `type` (String) The type of the custom variable loop. Allowed values: "string", "period", "bigquery", "snowflake", "redshift".

Optional:

- `bigquery_config` (Attributes) BigQuery custom variabe loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--bigquery_config))
- `is_parallel_execution_allowed` (Boolean) Whether parallel execution is allowed
- `is_stopped_on_errors` (Boolean) Whether the loop is stopped on errors
- `max_errors` (Number) The maximum number of errors
- `period_config` (Attributes) Period custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--period_config))
- `redshift_config` (Attributes) Redshift custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--redshift_config))
- `snowflake_config` (Attributes) Snowflake custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--snowflake_config))
- `string_config` (Attributes) String custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--string_config))

<a id="nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--bigquery_config"></a>
### Nested Schema for `tasks.trocco_redshift_datamart_config.custom_variable_loop.bigquery_config`

Required:

- `connection_id` (Number) BigQuery connection ID
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded


<a id="nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--period_config"></a>
### Nested Schema for `tasks.trocco_redshift_datamart_config.custom_variable_loop.period_config`

Required:

- `from` (Attributes) Start of the loop (see [below for nested schema](#nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--period_config--from))
- `interval` (String) Interval of the loop
- `time_zone` (String) Timezone of the configuration
- `to` (Attributes) End of the loop (see [below for nested schema](#nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--period_config--to))
- `variables` (Attributes List) Custom variables to be expanded (see [below for nested schema](#nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--period_config--variables))

<a id="nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--period_config--from"></a>
### Nested Schema for `tasks.trocco_redshift_datamart_config.custom_variable_loop.period_config.from`

Required:

- `unit` (String) Unit
- `value` (Number) Value


<a id="nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--period_config--to"></a>
### Nested Schema for `tasks.trocco_redshift_datamart_config.custom_variable_loop.period_config.to`

Required:

- `unit` (String) Unit
- `value` (Number) Value


<a id="nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--period_config--variables"></a>
### Nested Schema for `tasks.trocco_redshift_datamart_config.custom_variable_loop.period_config.variables`

Required:

- `name` (String) Name of custom variable
- `offset` (Attributes) Offset on custom variable expanded (see [below for nested schema](#nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--period_config--variables--offset))

<a id="nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--period_config--variables--offset"></a>
### Nested Schema for `tasks.trocco_redshift_datamart_config.custom_variable_loop.period_config.variables.offset`

Required:

- `unit` (String) Unit
- `value` (Number) Value




<a id="nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--redshift_config"></a>
### Nested Schema for `tasks.trocco_redshift_datamart_config.custom_variable_loop.redshift_config`

Required:

- `connection_id` (Number) Redshift connection ID
- `database` (String) Redshift database
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded


<a id="nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--snowflake_config"></a>
### Nested Schema for `tasks.trocco_redshift_datamart_config.custom_variable_loop.snowflake_config`

Required:

- `connection_id` (Number) Snowflake connection ID
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded
- `warehouse` (String) Snowflake warehouse


<a id="nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--string_config"></a>
### Nested Schema for `tasks.trocco_redshift_datamart_config.custom_variable_loop.string_config`

Required:

- `variables` (Attributes List) Custom variables (see [below for nested schema](#nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--string_config--variables))

<a id="nestedatt--tasks--trocco_redshift_datamart_config--custom_variable_loop--string_config--variables"></a>
### Nested Schema for `tasks.trocco_redshift_datamart_config.custom_variable_loop.string_config.variables`

Required:

- `name` (String) Custom variable name
- `values` (List of String) Custom variable values





<a id="nestedatt--tasks--trocco_snowflake_datamart_config"></a>
### Nested Schema for `tasks.trocco_snowflake_datamart_config`

Required:

- `definition_id` (Number) The definition id to use for the datamart task

Optional:

- `custom_variable_loop` (Attributes) The custom variable loop of the pipeline definition (see [below for nested schema](#nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop))

<a id="nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop"></a>
### Nested Schema for `tasks.trocco_snowflake_datamart_config.custom_variable_loop`

Required:

- `type` (String) The type of the custom variable loop. Allowed values: "string", "period", "bigquery", "snowflake", "redshift".

Optional:

- `bigquery_config` (Attributes) BigQuery custom variabe loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--bigquery_config))
- `is_parallel_execution_allowed` (Boolean) Whether parallel execution is allowed
- `is_stopped_on_errors` (Boolean) Whether the loop is stopped on errors
- `max_errors` (Number) The maximum number of errors
- `period_config` (Attributes) Period custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--period_config))
- `redshift_config` (Attributes) Redshift custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--redshift_config))
- `snowflake_config` (Attributes) Snowflake custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--snowflake_config))
- `string_config` (Attributes) String custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--string_config))

<a id="nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--bigquery_config"></a>
### Nested Schema for `tasks.trocco_snowflake_datamart_config.custom_variable_loop.bigquery_config`

Required:

- `connection_id` (Number) BigQuery connection ID
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded


<a id="nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--period_config"></a>
### Nested Schema for `tasks.trocco_snowflake_datamart_config.custom_variable_loop.period_config`

Required:

- `from` (Attributes) Start of the loop (see [below for nested schema](#nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--period_config--from))
- `interval` (String) Interval of the loop
- `time_zone` (String) Timezone of the configuration
- `to` (Attributes) End of the loop (see [below for nested schema](#nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--period_config--to))
- `variables` (Attributes List) Custom variables to be expanded (see [below for nested schema](#nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--period_config--variables))

<a id="nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--period_config--from"></a>
### Nested Schema for `tasks.trocco_snowflake_datamart_config.custom_variable_loop.period_config.from`

Required:

- `unit` (String) Unit
- `value` (Number) Value


<a id="nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--period_config--to"></a>
### Nested Schema for `tasks.trocco_snowflake_datamart_config.custom_variable_loop.period_config.to`

Required:

- `unit` (String) Unit
- `value` (Number) Value


<a id="nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--period_config--variables"></a>
### Nested Schema for `tasks.trocco_snowflake_datamart_config.custom_variable_loop.period_config.variables`

Required:

- `name` (String) Name of custom variable
- `offset` (Attributes) Offset on custom variable expanded (see [below for nested schema](#nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--period_config--variables--offset))

<a id="nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--period_config--variables--offset"></a>
### Nested Schema for `tasks.trocco_snowflake_datamart_config.custom_variable_loop.period_config.variables.offset`

Required:

- `unit` (String) Unit
- `value` (Number) Value




<a id="nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--redshift_config"></a>
### Nested Schema for `tasks.trocco_snowflake_datamart_config.custom_variable_loop.redshift_config`

Required:

- `connection_id` (Number) Redshift connection ID
- `database` (String) Redshift database
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded


<a id="nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--snowflake_config"></a>
### Nested Schema for `tasks.trocco_snowflake_datamart_config.custom_variable_loop.snowflake_config`

Required:

- `connection_id` (Number) Snowflake connection ID
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded
- `warehouse` (String) Snowflake warehouse


<a id="nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--string_config"></a>
### Nested Schema for `tasks.trocco_snowflake_datamart_config.custom_variable_loop.string_config`

Required:

- `variables` (Attributes List) Custom variables (see [below for nested schema](#nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--string_config--variables))

<a id="nestedatt--tasks--trocco_snowflake_datamart_config--custom_variable_loop--string_config--variables"></a>
### Nested Schema for `tasks.trocco_snowflake_datamart_config.custom_variable_loop.string_config.variables`

Required:

- `name` (String) Custom variable name
- `values` (List of String) Custom variable values





<a id="nestedatt--tasks--trocco_transfer_bulk_config"></a>
### Nested Schema for `tasks.trocco_transfer_bulk_config`

Required:

- `definition_id` (Number) The definition id to use for the trocco transfer bulk task

Optional:

- `is_parallel_execution_allowed` (Boolean) Whether the task is allowed to run in parallel
- `is_stopped_on_errors` (Boolean) Whether the task should stop on errors
- `max_errors` (Number) The maximum number of errors allowed before the task stops


<a id="nestedatt--tasks--trocco_transfer_config"></a>
### Nested Schema for `tasks.trocco_transfer_config`

Required:

- `definition_id` (Number) The definition id to use for the trocco transfer task

Optional:

- `custom_variable_loop` (Attributes) The custom variable loop of the pipeline definition (see [below for nested schema](#nestedatt--tasks--trocco_transfer_config--custom_variable_loop))

<a id="nestedatt--tasks--trocco_transfer_config--custom_variable_loop"></a>
### Nested Schema for `tasks.trocco_transfer_config.custom_variable_loop`

Required:

- `type` (String) The type of the custom variable loop. Allowed values: "string", "period", "bigquery", "snowflake", "redshift".

Optional:

- `bigquery_config` (Attributes) BigQuery custom variabe loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_transfer_config--custom_variable_loop--bigquery_config))
- `is_parallel_execution_allowed` (Boolean) Whether parallel execution is allowed
- `is_stopped_on_errors` (Boolean) Whether the loop is stopped on errors
- `max_errors` (Number) The maximum number of errors
- `period_config` (Attributes) Period custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_transfer_config--custom_variable_loop--period_config))
- `redshift_config` (Attributes) Redshift custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_transfer_config--custom_variable_loop--redshift_config))
- `snowflake_config` (Attributes) Snowflake custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_transfer_config--custom_variable_loop--snowflake_config))
- `string_config` (Attributes) String custom variable loop configuration (see [below for nested schema](#nestedatt--tasks--trocco_transfer_config--custom_variable_loop--string_config))

<a id="nestedatt--tasks--trocco_transfer_config--custom_variable_loop--bigquery_config"></a>
### Nested Schema for `tasks.trocco_transfer_config.custom_variable_loop.bigquery_config`

Required:

- `connection_id` (Number) BigQuery connection ID
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded


<a id="nestedatt--tasks--trocco_transfer_config--custom_variable_loop--period_config"></a>
### Nested Schema for `tasks.trocco_transfer_config.custom_variable_loop.period_config`

Required:

- `from` (Attributes) Start of the loop (see [below for nested schema](#nestedatt--tasks--trocco_transfer_config--custom_variable_loop--period_config--from))
- `interval` (String) Interval of the loop
- `time_zone` (String) Timezone of the configuration
- `to` (Attributes) End of the loop (see [below for nested schema](#nestedatt--tasks--trocco_transfer_config--custom_variable_loop--period_config--to))
- `variables` (Attributes List) Custom variables to be expanded (see [below for nested schema](#nestedatt--tasks--trocco_transfer_config--custom_variable_loop--period_config--variables))

<a id="nestedatt--tasks--trocco_transfer_config--custom_variable_loop--period_config--from"></a>
### Nested Schema for `tasks.trocco_transfer_config.custom_variable_loop.period_config.from`

Required:

- `unit` (String) Unit
- `value` (Number) Value


<a id="nestedatt--tasks--trocco_transfer_config--custom_variable_loop--period_config--to"></a>
### Nested Schema for `tasks.trocco_transfer_config.custom_variable_loop.period_config.to`

Required:

- `unit` (String) Unit
- `value` (Number) Value


<a id="nestedatt--tasks--trocco_transfer_config--custom_variable_loop--period_config--variables"></a>
### Nested Schema for `tasks.trocco_transfer_config.custom_variable_loop.period_config.variables`

Required:

- `name` (String) Name of custom variable
- `offset` (Attributes) Offset on custom variable expanded (see [below for nested schema](#nestedatt--tasks--trocco_transfer_config--custom_variable_loop--period_config--variables--offset))

<a id="nestedatt--tasks--trocco_transfer_config--custom_variable_loop--period_config--variables--offset"></a>
### Nested Schema for `tasks.trocco_transfer_config.custom_variable_loop.period_config.variables.offset`

Required:

- `unit` (String) Unit
- `value` (Number) Value




<a id="nestedatt--tasks--trocco_transfer_config--custom_variable_loop--redshift_config"></a>
### Nested Schema for `tasks.trocco_transfer_config.custom_variable_loop.redshift_config`

Required:

- `connection_id` (Number) Redshift connection ID
- `database` (String) Redshift database
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded


<a id="nestedatt--tasks--trocco_transfer_config--custom_variable_loop--snowflake_config"></a>
### Nested Schema for `tasks.trocco_transfer_config.custom_variable_loop.snowflake_config`

Required:

- `connection_id` (Number) Snowflake connection ID
- `query` (String) Query to expand custom variables
- `variables` (List of String) Custom variables to be expanded
- `warehouse` (String) Snowflake warehouse


<a id="nestedatt--tasks--trocco_transfer_config--custom_variable_loop--string_config"></a>
### Nested Schema for `tasks.trocco_transfer_config.custom_variable_loop.string_config`

Required:

- `variables` (Attributes List) Custom variables (see [below for nested schema](#nestedatt--tasks--trocco_transfer_config--custom_variable_loop--string_config--variables))

<a id="nestedatt--tasks--trocco_transfer_config--custom_variable_loop--string_config--variables"></a>
### Nested Schema for `tasks.trocco_transfer_config.custom_variable_loop.string_config.variables`

Required:

- `name` (String) Custom variable name
- `values` (List of String) Custom variable values








## Import

Import is supported using the following syntax:

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import trocco_pipeline_definition (Resource). For example:

```terraform
import {
  id = 1
  to = trocco_pipeline_definition.example
}
```

Using the [`terraform import` command](https://developer.hashicorp.com/terraform/cli/commands/import):

```shell
terraform import trocco_pipeline_definition.example <id>
```

