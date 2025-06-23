## 0.17.0
FEATURES:
- `trocco_job_definition` resource:
  - Added support for `http` input option with comprehensive HTTP data fetching capabilities
    - Supports various HTTP methods (GET, POST), pagination (offset, cursor), and multiple parsers (CSV, JSONL, JSONPath, LTSV, Excel, XML)
    - Includes request/response configuration options like headers, parameters, timeouts, and retry settings
- `trocco_pipeline_definition` resource:
  - Changed `notifications` from `List` to `Set` type to improve consistency and avoid ordering issues

CHORE:
- Updated Go version from 1.21 to 1.23.0
- Updated multiple dependencies including:
  - terraform-plugin-docs: 0.19.4 → 0.21.0
  - terraform-plugin-framework: 1.11.0 → 1.15.0
  - terraform-plugin-go: 0.23.0 → 0.27.0
  - Various other dependency updates for improved security and performance
- Added comprehensive validation for HTTP input options
- Enhanced documentation with HTTP input option examples

## 0.16.0
FEATURES:
- `trocco_job_definition` resource:
  - Changed `bigquery_output_option_clustering_fields` and `bigquery_output_option_merge_keys` from `List` to `Set` type
  - Made `bigquery_output_option_clustering_fields` and `bigquery_output_option_merge_keys` optional instead of required
  - Added validation for `bigquery_output_option_merge_keys` to be required only when `mode` is `merge`
- Added new plan modifiers:
  - `EmptyListForNull` and `EmptySetForNull` to avoid unnecessary diffs

CHORE:
- Added review workflow documentation
- Added `.gitignore` for logs directory

## 0.15.2
CHORE:
- Add validation for `before_load` in `trocco_bigquery_datamart_definition` resource
  - `before_load` is only available when `write_disposition` is "append"
- Use custom type TrimmedStringValue in `pipeline_definition` resource
  - `description` field

## 0.15.1
CHORE:
- Add tests
- Fix GitHub Actions timeout settings
- Update fields to use TrimmedStringValue custom type
    - `before_load` in `trocco_bigquery_datamart_definition` resource
    - `message` in `trocco_bigquery_datamart_definition.notifications` resource
    - `message` in `trocco_job_definition.notifications` resource
    - `message` in `trocco_pipeline_definition.notifications` resource

## 0.15.0
FEATURES:
- Added support for `yahoo_ads_api_yss` input option in `trocco_job_definition` resource.

CHORE:
- Use custom type TrimStringValue in `pipeline_definition` resource.
    - `bigquery_data_check_config.query`
    - `redshift_data_check_config.query`
    - `snowflake_data_check_config.query`
- Add tests & coverage report.
- Fix connection types in documentation.
- Fix import block example in documentation.

## 0.14.0
FEATURES:
- Added `kintone` input in `trocco_job_definition` resource.

CHORE:
- Add pinact-action

## 0.13.0
FEATURES:
- Added `trocco_notification_destination` resource.

CHORE:
- Fix `custom_variable_settings` in `trocco_bigquery_datamart_definition` resource.

## 0.12.0
FEATURES:
- Added `kintone` type for `trocco_connection` resource.
- Added `google_analytics4` input for `trocco_job_definition` resource.

CHORE:
- Fix validation for importing `trocco_user` resource.

## 0.11.0
FEATURES:
- Added `google_analytics4` type for `trocco_connection` resource.

CHORE:
- Add Import block examples.
- Ignore target for e2e and tests.

## 0.10.0
FEATURES:
- Added `postgresql` input in `trocco_job_definition` resource.

CHORE:
- Added key to `trocco_pipeline_definition` resource with terraform import.
- Fix version setting in UserAgent

## 0.9.0
FEATURES:
- Added `postgresql` type for `trocco_connection` resource.
- Added `google_spreadsheets` input/output in `trocco_job_definition` resource.
- Added `driver` for `trocco_connection` resource in `snowflake` and `mysql`.

CHORE:
- Fix datamart label change
- Fix examples in sample label color
- Update documentation & examples
- Added E2E testings with GitHub Actions
- Fix problems on importing trocco_job_definition

## 0.8.0
FEATURES:
- Added `salesforce` and `google_spreadsheets` type for `trocco_connection` resource.
- Added `salesforce` input/output in `trocco_job_definition` resource.

CHORE:
- Remove `auto_create_table` option from `bigquery` output for `trocco_job_definition` resource.
- Update documentation & examples

## 0.7.0
FEATURES:
- Added `snowflake` input/output in `trocco_job_definition` resource.
- Added `s3` type for `trocco_connection` resource.

## 0.6.0
FEATURES:
- Added `resource_group` resource.
- Added `label` resource.
- Added `mysql` type for `trocco_connection` resource.

## 0.5.0
FEATURES:
- Added `trocco_job_definition` resource.
- Added `trocco_pipeline_definition` resource.

CHORE:
- Suppressed unnecessary diffs.

## 0.4.0
FEATURES:
- Supported `gcs` type for `trocco_connection` resource.

CHORE:
- Improved error messages.

## 0.3.0
FEATURES:
- Added `trocco_team` resource.

## 0.2.1
CHORE:
- Set provider version in the user agent for API calls.

## 0.2.0
FEATURES:
- Added `trocco_connection` resource.
- Added `trocco_user` resource.

## 0.1.4
CHORE:
- Reduction in the number of TROCCO API calls

## 0.1.3
CHORE:
- Updated terraform-plugin-framework to v1.11.0 from v1.10.0

## 0.1.2
CHORE:
- Refined error messages when API error occur.

## 0.1.1
FEATURES:
- Added `trocco_bigquery_datamart_definition` resource.
