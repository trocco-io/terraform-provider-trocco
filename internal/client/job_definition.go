package client

import (
	"fmt"
	"net/http"
	"terraform-provider-trocco/internal/client/entity"
	jobDefinitionEntities "terraform-provider-trocco/internal/client/entity/job_definition"
	filterEntities "terraform-provider-trocco/internal/client/entity/job_definition/filter"
	inputOptionEntities "terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	outputOptionEntities "terraform-provider-trocco/internal/client/entity/job_definition/output_option"
	"terraform-provider-trocco/internal/client/parameter"
	jobDefinitionParameters "terraform-provider-trocco/internal/client/parameter/job_definition"
	filterParameters "terraform-provider-trocco/internal/client/parameter/job_definition/filter"
	inputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	outputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"
)

type JobDefinition struct {
	ID                        int64                                             `json:"id"`
	Name                      string                                            `json:"name"`
	Description               *string                                           `json:"description"`
	ResourceGroupID           *int64                                            `json:"resource_group_id"`
	IsRunnableConcurrently    *bool                                             `json:"is_runnable_concurrently"`
	RetryLimit                int64                                             `json:"retry_limit"`
	ResourceEnhancement       *string                                           `json:"resource_enhancement"`
	FilterColumns             []filterEntities.FilterColumn                     `json:"filter_columns"`
	FilterRows                *filterEntities.FilterRows                        `json:"filter_rows"`
	FilterMasks               []filterEntities.FilterMask                       `json:"filter_masks"`
	FilterAddTime             *filterEntities.FilterAddTime                     `json:"filter_add_time"`
	FilterGsub                []filterEntities.FilterGsub                       `json:"filter_gsub"`
	FilterStringTransforms    []filterEntities.FilterStringTransform            `json:"filter_string_transforms"`
	FilterHashes              []filterEntities.FilterHash                       `json:"filter_hashes"`
	FilterUnixTimeConversions []filterEntities.FilterUnixTimeConversion         `json:"filter_unixtime_conversions"`
	InputOptionType           string                                            `json:"input_option_type"`
	InputOption               InputOption                                       `json:"input_option"`
	OutputOptionType          string                                            `json:"output_option_type"`
	OutputOption              OutputOption                                      `json:"output_option"`
	Labels                    []entity.Label                                    `json:"labels"`
	Schedules                 []entity.Schedule                                 `json:"schedules"`
	Notifications             []jobDefinitionEntities.JobDefinitionNotification `json:"notifications"`
}

type CreateJobDefinitionInput struct {
	Name                      string                                                         `json:"name"`
	Description               *parameter.NullableString                                      `json:"description,omitempty"`
	ResourceGroupID           *parameter.NullableInt64                                       `json:"resource_group_id,omitempty"`
	IsRunnableConcurrently    bool                                                           `json:"is_runnable_concurrently"`
	RetryLimit                int64                                                          `json:"retry_limit"`
	ResourceEnhancement       *string                                                        `json:"resource_enhancement,omitempty"`
	FilterColumns             []filterParameters.FilterColumnInput                           `json:"filter_columns"`
	FilterRows                *parameter.NullableObject[filterParameters.FilterRowsInput]    `json:"filter_rows,omitempty"`
	FilterMasks               []filterParameters.FilterMaskInput                             `json:"filter_masks"`
	FilterAddTime             *parameter.NullableObject[filterParameters.FilterAddTimeInput] `json:"filter_add_time,omitempty"`
	FilterGsub                []filterParameters.FilterGsubInput                             `json:"filter_gsub"`
	FilterStringTransforms    []filterParameters.FilterStringTransformInput                  `json:"filter_string_transforms"`
	FilterHashes              []filterParameters.FilterHashInput                             `json:"filter_hashes"`
	FilterUnixTimeConversions []filterParameters.FilterUnixTimeConversionInput               `json:"filter_unixtime_conversions"`
	InputOptionType           string                                                         `json:"input_option_type"`
	InputOption               InputOptionInput                                               `json:"input_option"`
	OutputOptionType          string                                                         `json:"output_option_type"`
	OutputOption              OutputOptionInput                                              `json:"output_option"`
	Labels                    []string                                                       `json:"labels"`
	Schedules                 []parameter.ScheduleInput                                      `json:"schedules"`
	Notifications             []jobDefinitionParameters.JobDefinitionNotificationInput       `json:"notifications"`
}

type UpdateJobDefinitionInput struct {
	Name                      *string                                                        `json:"name,omitempty"`
	Description               *parameter.NullableString                                      `json:"description,omitempty"`
	ResourceGroupID           *parameter.NullableInt64                                       `json:"resource_group_id,omitempty"`
	IsRunnableConcurrently    *bool                                                          `json:"is_runnable_concurrently,omitempty"`
	RetryLimit                *int64                                                         `json:"retry_limit,omitempty"`
	ResourceEnhancement       *string                                                        `json:"resource_enhancement,omitempty"`
	FilterColumns             *[]filterParameters.FilterColumnInput                          `json:"filter_columns,omitempty"`
	FilterRows                *parameter.NullableObject[filterParameters.FilterRowsInput]    `json:"filter_rows,omitempty"`
	FilterMasks               *[]filterParameters.FilterMaskInput                            `json:"filter_masks,omitempty"`
	FilterAddTime             *parameter.NullableObject[filterParameters.FilterAddTimeInput] `json:"filter_add_time,omitempty"`
	FilterGsub                *[]filterParameters.FilterGsubInput                            `json:"filter_gsub,omitempty"`
	FilterStringTransforms    *[]filterParameters.FilterStringTransformInput                 `json:"filter_string_transforms,omitempty"`
	FilterHashes              *[]filterParameters.FilterHashInput                            `json:"filter_hashes,omitempty"`
	FilterUnixTimeConversions *[]filterParameters.FilterUnixTimeConversionInput              `json:"filter_unixtime_conversions,omitempty"`
	InputOption               *UpdateInputOptionInput                                        `json:"input_option,omitempty"`
	OutputOption              *UpdateOutputOptionInput                                       `json:"output_option,omitempty"`
	Labels                    *[]string                                                      `json:"labels,omitempty"`
	Schedules                 *[]parameter.ScheduleInput                                     `json:"schedules,omitempty"`
	Notifications             *[]jobDefinitionParameters.JobDefinitionNotificationInput      `json:"notifications,omitempty"`
}

type InputOption struct {
	MySQLInputOption              *inputOptionEntities.MySQLInputOption              `json:"mysql_input_option"`
	GcsInputOption                *inputOptionEntities.GcsInputOption                `json:"gcs_input_option"`
	SnowflakeInputOption          *inputOptionEntities.SnowflakeInputOption          `json:"snowflake_input_option"`
	SalesforceInputOption         *inputOptionEntities.SalesforceInputOption         `json:"salesforce_input_option"`
	GoogleSpreadsheetsInputOption *inputOptionEntities.GoogleSpreadsheetsInputOption `json:"google_spreadsheets_input_option"`
	S3InputOption                 *inputOptionEntities.S3InputOption                 `json:"s3_input_option"`
	BigqueryInputOption           *inputOptionEntities.BigqueryInputOption           `json:"bigquery_input_option"`
	PostgreSQLInputOption         *inputOptionEntities.PostgreSQLInputOption         `json:"postgresql_input_option"`
	GoogleAnalytics4InputOption   *inputOptionEntities.GoogleAnalytics4InputOption   `json:"google_analytics4_input_option"`
	HttpInputOption               *inputOptionEntities.HttpInputOption               `json:"http_input_option"`
	KintoneInputOption            *inputOptionEntities.KintoneInputOption            `json:"kintone_input_option"`
	YahooAdsApiYssInputOption     *inputOptionEntities.YahooAdsApiYssInputOption     `json:"yahoo_ads_api_yss_input_option"`
	DatabricksInputOption         *inputOptionEntities.DatabricksInputOption         `json:"databricks_input_option"`
}

type InputOptionInput struct {
	MySQLInputOption              *parameter.NullableObject[inputOptionParameters.MySQLInputOptionInput]              `json:"mysql_input_option,omitempty"`
	GcsInputOption                *parameter.NullableObject[inputOptionParameters.GcsInputOptionInput]                `json:"gcs_input_option,omitempty"`
	SnowflakeInputOption          *parameter.NullableObject[inputOptionParameters.SnowflakeInputOptionInput]          `json:"snowflake_input_option,omitempty"`
	SalesforceInputOption         *parameter.NullableObject[inputOptionParameters.SalesforceInputOptionInput]         `json:"salesforce_input_option,omitempty"`
	GoogleSpreadsheetsInputOption *parameter.NullableObject[inputOptionParameters.GoogleSpreadsheetsInputOptionInput] `json:"google_spreadsheets_input_option,omitempty"`
	S3InputOption                 *parameter.NullableObject[inputOptionParameters.S3InputOptionInput]                 `json:"s3_input_option,omitempty"`
	BigqueryInputOption           *parameter.NullableObject[inputOptionParameters.BigqueryInputOptionInput]           `json:"bigquery_input_option,omitempty"`
	PostgreSQLInputOption         *parameter.NullableObject[inputOptionParameters.PostgreSQLInputOptionInput]         `json:"postgresql_input_option,omitempty"`
	GoogleAnalytics4InputOption   *parameter.NullableObject[inputOptionParameters.GoogleAnalytics4InputOptionInput]   `json:"google_analytics4_input_option,omitempty"`
	HttpInputOption               *parameter.NullableObject[inputOptionParameters.HttpInputOptionInput]               `json:"http_input_option,omitempty"`
	KintoneInputOption            *parameter.NullableObject[inputOptionParameters.KintoneInputOptionInput]            `json:"kintone_input_option,omitempty"`
	YahooAdsApiYssInputOption     *parameter.NullableObject[inputOptionParameters.YahooAdsApiYssInputOptionInput]     `json:"yahoo_ads_api_yss_input_option,omitempty"`
	DatabricksInputOption         *parameter.NullableObject[inputOptionParameters.DatabricksInputOptionInput]         `json:"databricks_input_option,omitempty"`
}

type UpdateInputOptionInput struct {
	MySQLInputOption              *parameter.NullableObject[inputOptionParameters.UpdateMySQLInputOptionInput]              `json:"mysql_input_option,omitempty"`
	GcsInputOption                *parameter.NullableObject[inputOptionParameters.UpdateGcsInputOptionInput]                `json:"gcs_input_option,omitempty"`
	SnowflakeInputOption          *parameter.NullableObject[inputOptionParameters.UpdateSnowflakeInputOptionInput]          `json:"snowflake_input_option,omitempty"`
	SalesforceInputOption         *parameter.NullableObject[inputOptionParameters.UpdateSalesforceInputOptionInput]         `json:"salesforce_input_option,omitempty"`
	GoogleSpreadsheetsInputOption *parameter.NullableObject[inputOptionParameters.UpdateGoogleSpreadsheetsInputOptionInput] `json:"google_spreadsheets_input_option,omitempty"`
	S3InputOption                 *parameter.NullableObject[inputOptionParameters.UpdateS3InputOptionInput]                 `json:"s3_input_option,omitempty"`
	BigqueryInputOption           *parameter.NullableObject[inputOptionParameters.UpdateBigqueryInputOptionInput]           `json:"bigquery_input_option,omitempty"`
	PostgreSQLInputOption         *parameter.NullableObject[inputOptionParameters.UpdatePostgreSQLInputOptionInput]         `json:"postgresql_input_option,omitempty"`
	GoogleAnalytics4InputOption   *parameter.NullableObject[inputOptionParameters.UpdateGoogleAnalytics4InputOptionInput]   `json:"google_analytics4_input_option,omitempty"`
	HttpInputOption               *parameter.NullableObject[inputOptionParameters.UpdateHttpInputOptionInput]               `json:"http_input_option,omitempty"`
	KintoneInputOption            *parameter.NullableObject[inputOptionParameters.UpdateKintoneInputOptionInput]            `json:"kintone_input_option,omitempty"`
	YahooAdsApiYssInputOption     *parameter.NullableObject[inputOptionParameters.UpdateYahooAdsApiYssInputOptionInput]     `json:"yahoo_ads_api_yss_input_option,omitempty"`
	DatabricksInputOption         *parameter.NullableObject[inputOptionParameters.UpdateDatabricksInputOptionInput]         `json:"databricks_input_option,omitempty"`
}

type OutputOption struct {
	BigQueryOutputOption           *outputOptionEntities.BigQueryOutputOption           `json:"bigquery_output_option"`
	SnowflakeOutputOption          *outputOptionEntities.SnowflakeOutputOption          `json:"snowflake_output_option"`
	SalesforceOutputOption         *outputOptionEntities.SalesforceOutputOption         `json:"salesforce_output_option"`
	GoogleSpreadsheetsOutputOption *outputOptionEntities.GoogleSpreadsheetsOutputOption `json:"google_spreadsheets_output_option"`
	DatabricksOutputOption         *outputOptionEntities.DatabricksOutputOption         `json:"databricks_output_option"`
	PostgresqlOutputOption         *outputOptionEntities.PostgresqlOutputOption         `json:"postgresql_output_option"`
}

type OutputOptionInput struct {
	BigQueryOutputOption           *parameter.NullableObject[outputOptionParameters.BigQueryOutputOptionInput]           `json:"bigquery_output_option,omitempty"`
	SnowflakeOutputOption          *parameter.NullableObject[outputOptionParameters.SnowflakeOutputOptionInput]          `json:"snowflake_output_option,omitempty"`
	SalesforceOutputOption         *parameter.NullableObject[outputOptionParameters.SalesforceOutputOptionInput]         `json:"salesforce_output_option,omitempty"`
	GoogleSpreadsheetsOutputOption *parameter.NullableObject[outputOptionParameters.GoogleSpreadsheetsOutputOptionInput] `json:"google_spreadsheets_output_option,omitempty"`
	DatabricksOutputOption         *parameter.NullableObject[outputOptionParameters.DatabricksOutputOptionInput]         `json:"databricks_output_option,omitempty"`
	PostgresqlOutputOption         *parameter.NullableObject[outputOptionParameters.PostgresqlOutputOptionInput]         `json:"postgresql_output_option,omitempty"`
}

type UpdateOutputOptionInput struct {
	BigQueryOutputOption           *parameter.NullableObject[outputOptionParameters.UpdateBigQueryOutputOptionInput]           `json:"bigquery_output_option,omitempty"`
	SnowflakeOutputOption          *parameter.NullableObject[outputOptionParameters.UpdateSnowflakeOutputOptionInput]          `json:"snowflake_output_option,omitempty"`
	SalesforceOutputOption         *parameter.NullableObject[outputOptionParameters.UpdateSalesforceOutputOptionInput]         `json:"salesforce_output_option,omitempty"`
	GoogleSpreadsheetsOutputOption *parameter.NullableObject[outputOptionParameters.UpdateGoogleSpreadsheetsOutputOptionInput] `json:"google_spreadsheets_output_option,omitempty"`
	DatabricksOutputOption         *parameter.NullableObject[outputOptionParameters.UpdateDatabricksOutputOptionInput]         `json:"databricks_output_option,omitempty"`
	PostgresqlOutputOption         *parameter.NullableObject[outputOptionParameters.UpdatePostgresqlOutputOptionInput]         `json:"postgresql_output_option,omitempty"`
}

func (c *TroccoClient) CreateJobDefinition(in *CreateJobDefinitionInput) (*JobDefinition, error) {
	out := &JobDefinition{}
	if err := c.do(
		http.MethodPost,
		"/api/job_definitions",
		in,
		out,
	); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) UpdateJobDefinition(id int64, in *UpdateJobDefinitionInput) (*JobDefinition, error) {
	out := &JobDefinition{}
	if err := c.do(
		http.MethodPatch,
		fmt.Sprintf("/api/job_definitions/%d", id),
		in,
		out,
	); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) DeleteJobDefinition(id int64) error {
	return c.do(
		http.MethodDelete,
		fmt.Sprintf("/api/job_definitions/%d", id),
		nil,
		nil,
	)
}

func (c *TroccoClient) GetJobDefinition(id int64) (*JobDefinition, error) {
	out := &JobDefinition{}
	if err := c.do(
		http.MethodGet,
		fmt.Sprintf("/api/job_definitions/%d", id),
		nil,
		out,
	); err != nil {
		return nil, err
	}
	return out, nil
}
