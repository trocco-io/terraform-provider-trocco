package provider

import (
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccJobDefinitionResourceMysqlToBigQuery(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "trocco_connection" "test_mysql" {
  connection_type = "mysql"
  name            = "MySQL Example"
  host            = "db.example.com"
  port            = 65535
  user_name       = "root"
  password        = "password"
}
resource "trocco_connection" "test_bq" {
  connection_type = "bigquery"
  name        = "BigQuery Example"
  project_id               = "example"
  service_account_json_key = <<JSON
  {
    "type": "service_account",
    "project_id": "example-project-id",
    "private_key_id": "example-private-key-id",
    "private_key":"-----BEGIN PRIVATE KEY-----\n..."
  }
  JSON
}
resource "trocco_team" "test" {
  name        = "test"
  members     = [
    {
      user_id = 10626
      role    = "team_admin"
    },
  ]
}
resource "trocco_resource_group" "test" {
  name        = "test"
  description = "test"
  teams     = [
    {
      team_id = trocco_team.test.id
      role    = "administrator"
    },
  ]
}
resource "trocco_job_definition" "mysql_to_bigquery" {
  name                        = "test job_definition"
  description                 = "test description"
  resource_enhancement        = "large"
  resource_group_id           = trocco_resource_group.test.id
  retry_limit                 = 1
  is_runnable_concurrently    = true
  filter_columns              = [
    {
      default                      = ""
      format = "%Y"
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "id"
      src                          = "id"
      type                         = "long"
    },
    {
      default                      = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "name"
      src                          = "name"
      type                         = "string"
    },
    {
      default                      = ""
      format                       = ""
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "created_at"
      src                          = "created_at"
      type                         = "timestamp"
    },
    {
      default                      = ""
      json_expand_enabled          = true
      json_expand_keep_base_column = true
      name                         = "jsontekitou"
      src                          = ""
      type                         = "json"
      json_expand_columns          = [
        {
          json_path = "path"
          name      = "json_col"
          format = "%Y"
          timezone  = "UTC/ETC"
          type      = "string"
        },
      ]
    }
  ]
  filter_gsub                 = [
    {
      column_name = "regex_col"
      pattern     = "/regex/"
      to          = "replace_string"
    },
  ]
  filter_hashes               = [
    {
      name = "hash_col"
    },
  ]
  filter_masks                = [
    {
      length    = 9
      mask_type = "all"
      name      = "mask_all_string"
    },
    {
      length    = 10
      mask_type = "email"
      name      = "mask_email"
    },
    {
      mask_type = "regex"
      name      = "mask_regex"
      pattern   = "/regex/"
    },
    {
      start_index = 2
      end_index   = 2
      length      = 10
      mask_type   = "all"
      name        = "mail_partial_string"
    },
  ]
  filter_rows                 = {
    condition             = "or"
    filter_row_conditions = [
      {
        argument = "2"
        column   = "bbb"
        operator = "greater_equal"
      },
    ]
  }

  filter_string_transforms    = [
    {
      column_name = "transforms"
      type        = "normalize_nfkc"
    },
  ]
  filter_unixtime_conversions = [
    {
      column_name       = "unix_to_time"
      datetime_format   = "%Y-%m-%d %H:%M:%S.%N"
      datetime_timezone = "Asia/Tokyo"
      kind              = "unixtime_to_timestamp"
      unixtime_unit     = "second"
    },
    {
      column_name       = "unix_to_time_str"
      datetime_format   = "%Y-%m-%d %H:%M:%S.%N %z"
      datetime_timezone = "Etc/UTC"
      kind              = "unixtime_to_string"
      unixtime_unit     = "second"
    },
    {
      column_name       = "time_to_unix"
      datetime_format   = "%Y-%m-%d %H:%M:%S.%N %Z"
      datetime_timezone = "Etc/UTC"
      kind              = "timestamp_to_unixtime"
      unixtime_unit     = "second"
    },
    {
      column_name       = "time_str_to_unix"
      datetime_format   = "%Y-%m-%d %H:%M:%S.%N %z"
      datetime_timezone = "Etc/UTC"
      kind              = "string_to_unixtime"
      unixtime_unit     = "second"
    },
  ]
  input_option_type           = "mysql"
  input_option                = {
    mysql_input_option = {
      connect_timeout             = 300
      database                    = "test_database"
      fetch_rows                  = 1000
      incremental_loading_enabled = false
      table = "test_table"
      default_time_zone           = "Asia/Tokyo"
      use_legacy_datetime_code    = false
      input_option_columns        = [
        {
          name = "id"
          type = "long"
        },
        {
          name = "name"
          type = "string"
        },
        {
          name = "email"
          type = "string"
        },
        {
          name = "created_at"
          type = "timestamp"
        },
      ]
      mysql_connection_id         = trocco_connection.test_mysql.id
      query                       = <<-EOT
                select
                    *
                from
      	          example_table;
      EOT
      socket_timeout              = 1801
    }
  }
  output_option_type          = "bigquery"
  output_option               = {
    bigquery_output_option = {
      dataset                                    = "test_dataset"
      table                                      = "test_table"
      mode                                       = "append"
      auto_create_dataset                        = true
      timeout_sec                                = 300
      open_timeout_sec                           = 300
      read_timeout_sec                           = 300
      send_timeout_sec                           = 300
      retries                                    = 2
      bigquery_connection_id                     = trocco_connection.test_bq.id
      location                                   = "us-west1"
      bigquery_output_option_clustering_fields   = []
      bigquery_output_option_column_options      = []
      bigquery_output_option_merge_keys          = []
    }
  }
  # please create labels if testing in local environment
  # see https://trocco.io/labels#side-nav-labels
  labels = [
    {
      name = "label1"
    },
    {
      name = "label2"
    },
  ]
}

				`,
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("trocco_job_definition.mysql_to_bigquery", "name", "test job_definition"),
					resource.TestCheckResourceAttr("trocco_job_definition.mysql_to_bigquery", "description", "test description"),
					resource.TestCheckResourceAttr("trocco_job_definition.mysql_to_bigquery", "resource_enhancement", "large"),
					resource.TestCheckResourceAttr("trocco_job_definition.mysql_to_bigquery", "retry_limit", "1"),
					resource.TestCheckResourceAttr("trocco_job_definition.mysql_to_bigquery", "is_runnable_concurrently", "true"),
				),
			},
			{
				ResourceName:            "trocco_job_definition.mysql_to_bigquery",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					jobDefinitionId := s.RootModule().Resources["trocco_job_definition.mysql_to_bigquery"].Primary.ID

					return jobDefinitionId, nil
				},
			},
		},
	})
}

func TestAccJobDefinitionResourceS3ToSnowflake(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + `
resource "trocco_connection" "s3" {
  connection_type = "s3"
  name        = "S3 Example"
  description = "This is a AWS S3 connection example"
  aws_auth_type = "iam_user"
  aws_iam_user = {
    access_key_id     = "YOUR_ACCESS_KEY_ID"
    secret_access_key = "YOUR_SECRET_ACCESS_KEY"
  }
}

resource "trocco_connection" "snowflake" {
  connection_type = "snowflake"

  name        = "Snowflake Example"
  description = "This is a Snowflake connection example"

  host        = "exmaple.snowflakecomputing.com"
  auth_method = "user_password"
  user_name   = "dummy_name"
  password    = "dummy_password"
}


resource "trocco_job_definition" "s3_test" {
  description              = ""
  filter_columns           = [
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "name"
      src                          = "name"
      type                         = "string"
    },
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "url"
      src                          = "url"
      type                         = "string"
    },
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "test"
      src                          = "test"
      type                         = "string"
    },
    {
      default                      = null
      json_expand_enabled          = false
      json_expand_keep_base_column = false
      name                         = "asf"
      src                          = "asf"
      type                         = "string"
    },
  ]
  input_option             = {
    s3_input_option = {
      bucket                      = "test_bucket"
      csv_parser                  = {
        allow_extra_columns     = false
        allow_optional_columns  = false
        charset                 = "UTF-8"
        columns                 = [
          {
            name = "name"
            type = "string"
          },
          {
            name = "url"
            type = "string"
          },
          {
            name = "test"
            type = "string"
          },
          {
            name = "asf"
            type = "string"
          },
        ]
        comment_line_marker     = ""
        default_date            = "1970-01-01"
        default_time_zone       = "UTC"
        delimiter               = ","
        escape                  = "\""
        max_quoted_size_limit   = 131072
        newline                 = "CRLF"
        null_string             = ""
        null_string_enabled     = false
        quote                   = "\""
        quotes_in_quoted_fields = "ACCEPT_ONLY_RFC4180_ESCAPED"
        skip_header_lines       = 1
        stop_on_invalid_record  = true
        trim_if_not_quoted      = false
      }
      decompression_type          = "default"
      incremental_loading_enabled = false
      is_skip_header_line         = false
      path_match_pattern          = ""
      path_prefix                 = "dev/000.00.csv"
      region                      = "ap-northeast-1"
      s3_connection_id            = trocco_connection.s3.id
      stop_when_file_not_found    = false
    }
  }
  input_option_type        = "s3"
  is_runnable_concurrently = false
  name                     = "s3 to snowflake"
  output_option            = {
    snowflake_output_option = {
      batch_size              = 50
      database                = "test_database"
      default_time_zone       = "UTC"
      delete_stage_on_error   = false
      empty_field_as_null     = true
      max_retry_wait          = 1800000
      mode                    = "insert"
      retry_limit             = 12
      retry_wait              = 1000
      schema                  = "PUBLIC"
      snowflake_connection_id = trocco_connection.snowflake.id
      table                   = "ewaoiiowe"
      warehouse               = "COMPUTE_WH"
    }
  }
  output_option_type       = "snowflake"
  resource_enhancement     = "custom_spec"
  retry_limit              = 0
}


				`,
			},
			{
				ResourceName:            "trocco_job_definition.s3_test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					jobDefinitionId := s.RootModule().Resources["trocco_job_definition.s3_test"].Primary.ID
					return jobDefinitionId, nil
				},
			},
		},
	})
}

func TestAccJobDefinitionResourceGoogleAnalytics4ToSnowflake(t *testing.T) {
	resourceName := "trocco_job_definition.ga4_to_snowflake"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextile("../../examples/testdata/job_definition/google_analytics4_to_snowflake/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "GA4 to Snowflake"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.time_series", "dateHour"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.start_date", "2daysAgo"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.end_date", "1daysAgo"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.google_analytics4_input_option_dimensions.0.name", "yyyymm"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.google_analytics4_input_option_dimensions.0.expression", "{\n  \"concatenate\": {\n    \"dimensionNames\": [\"year\",\"month\"],\n    \"delimiter\": \"-\"\n  }\n}\n"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.google_analytics4_input_option_metrics.0.name", "totalUsers"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.google_analytics4_input_option_metrics.0.expression", ""),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.incremental_loading_enabled", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.retry_limit", "5"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.retry_sleep", "2"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.raise_on_other_row", "false"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.limit_of_rows", "10000"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.input_option_columns.0.name", "date_hour"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.input_option_columns.0.type", "timestamp"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.input_option_columns.1.name", "yyyymm"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.input_option_columns.1.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.input_option_columns.2.name", "total_users"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.input_option_columns.2.type", "long"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.input_option_columns.3.name", "property_id"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.input_option_columns.3.type", "string"),
				),
			},
			// ImportState testing
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					jobDefinitionId := s.RootModule().Resources[resourceName].Primary.ID
					return jobDefinitionId, nil
				},
			},
			// Update testing with null dimension
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextile("../../examples/testdata/job_definition/google_analytics4_to_snowflake/update_dimension_null.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "GA4 to Snowflake"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.time_series", "dateHour"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.start_date", "2daysAgo"),
					resource.TestCheckResourceAttr(resourceName, "input_option.google_analytics4_input_option.google_analytics4_input_option_dimensions.#", "0"),
				),
			},
		},
	})
}

func TestAccJobDefinitionResourceGoogleAnalytics4ToSnowflakeInvalid(t *testing.T) {
	resourceName := "trocco_job_definition.ga4_to_snowflake"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextile("../../examples/testdata/job_definition/google_analytics4_to_snowflake/update_dimension_too_many.tf"),
				ExpectError:  regexp.MustCompile(`list must contain at most 8 elements, got: 9`),
			},
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextile("../../examples/testdata/job_definition/google_analytics4_to_snowflake/update_metrics_required.tf"),
				ExpectError:  regexp.MustCompile(`"google_analytics4_input_option_metrics" is required.`),
			},
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextile("../../examples/testdata/job_definition/google_analytics4_to_snowflake/update_metrics_too_many.tf"),
				ExpectError:  regexp.MustCompile(`list must contain at least 1 elements and at most 10 elements, got: 11`),
			},
		},
	})
}

func TestAccJobDefinitionResourceKintoneToSnowflake(t *testing.T) {
	resourceName := "trocco_job_definition.kintone_to_snowflake"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextile("../../examples/testdata/job_definition/kintone_to_snowflake/create.tf"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "name", "kintone to snowflake"),
					resource.TestCheckResourceAttr(resourceName, "retry_limit", "0"),
					resource.TestCheckResourceAttr(resourceName, "resource_enhancement", "medium"),
					resource.TestCheckResourceAttr(resourceName, "output_option.snowflake_output_option.batch_size", "50"),
					resource.TestCheckResourceAttr(resourceName, "output_option.snowflake_output_option.database", "test_database"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.app_id", "123"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.expand_subtable", "false"),
					resource.TestCheckNoResourceAttr(resourceName, "input_option.kintone_input_option.guest_space_id"),
					resource.TestCheckNoResourceAttr(resourceName, "input_option.kintone_input_option.query"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.input_option_columns.#", "2"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.input_option_columns.0.name", "duration"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.input_option_columns.0.type", "string"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.input_option_columns.1.name", "date"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.input_option_columns.1.type", "timestamp"),
					resource.TestCheckResourceAttr(resourceName, "input_option.kintone_input_option.input_option_columns.1.format", "%Y%m%d"),
				),
			},
			// ImportState testing
			{
				ResourceName:            resourceName,
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
				ImportStateIdFunc: func(s *terraform.State) (string, error) {
					jobDefinitionId := s.RootModule().Resources[resourceName].Primary.ID
					return jobDefinitionId, nil
				},
			},
		},
	})
}

func TestAccJobDefinitionResourceKintoneToSnowflakeInvalid(t *testing.T) {
	resourceName := "trocco_job_definition.kintone_to_snowflake"
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				ResourceName: resourceName,
				Config:       providerConfig + LoadTextile("../../examples/testdata/job_definition/kintone_to_snowflake/update_app_id_required.tf"),
				ExpectError:  regexp.MustCompile(`Missing Configuration for Required Attribute`),
			},
		},
	})
}
