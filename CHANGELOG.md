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
