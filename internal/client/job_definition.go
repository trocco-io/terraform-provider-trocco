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
	jobDefinitions "terraform-provider-trocco/internal/client/parameter/job_definition"
	filterParameters "terraform-provider-trocco/internal/client/parameter/job_definition/filter"
	inputOptions "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	outputOptions "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"
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
	Notifications             []jobDefinitions.JobDefinitionNotificationInput                `json:"notifications"`
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
	Notifications             *[]jobDefinitions.JobDefinitionNotificationInput               `json:"notifications,omitempty"`
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
	HTTPInputOption               *inputOptionEntities.HTTPInputOption               `json:"http_input_option"`
	KintoneInputOption            *inputOptionEntities.KintoneInputOption            `json:"kintone_input_option"`
	YahooAdsApiYssInputOption     *inputOptionEntities.YahooAdsApiYssInputOption     `json:"yahoo_ads_api_yss_input_option"`
}

type InputOptionInput struct {
	MySQLInputOption              *parameter.NullableObject[inputOptions.MySQLInputOptionInput]              `json:"mysql_input_option,omitempty"`
	GcsInputOption                *parameter.NullableObject[inputOptions.GcsInputOptionInput]                `json:"gcs_input_option,omitempty"`
	SnowflakeInputOption          *parameter.NullableObject[inputOptions.SnowflakeInputOptionInput]          `json:"snowflake_input_option,omitempty"`
	SalesforceInputOption         *parameter.NullableObject[inputOptions.SalesforceInputOptionInput]         `json:"salesforce_input_option,omitempty"`
	GoogleSpreadsheetsInputOption *parameter.NullableObject[inputOptions.GoogleSpreadsheetsInputOptionInput] `json:"google_spreadsheets_input_option,omitempty"`
	S3InputOption                 *parameter.NullableObject[inputOptions.S3InputOptionInput]                 `json:"s3_input_option,omitempty"`
	BigqueryInputOption           *parameter.NullableObject[inputOptions.BigqueryInputOptionInput]           `json:"bigquery_input_option,omitempty"`
	PostgreSQLInputOption         *parameter.NullableObject[inputOptions.PostgreSQLInputOptionInput]         `json:"postgresql_input_option,omitempty"`
	GoogleAnalytics4InputOption   *parameter.NullableObject[inputOptions.GoogleAnalytics4InputOptionInput]   `json:"google_analytics4_input_option,omitempty"`
	HTTPInputOption               *parameter.NullableObject[inputOptions.HTTPInputOptionInput]               `json:"http_input_option,omitempty"`
	KintoneInputOption            *parameter.NullableObject[inputOptions.KintoneInputOptionInput]            `json:"kintone_input_option,omitempty"`
	YahooAdsApiYssInputOption     *parameter.NullableObject[inputOptions.YahooAdsApiYssInputOptionInput]     `json:"yahoo_ads_api_yss_input_option,omitempty"`
}

type UpdateInputOptionInput struct {
	MySQLInputOption              *parameter.NullableObject[inputOptions.UpdateMySQLInputOptionInput]              `json:"mysql_input_option,omitempty"`
	GcsInputOption                *parameter.NullableObject[inputOptions.UpdateGcsInputOptionInput]                `json:"gcs_input_option,omitempty"`
	SnowflakeInputOption          *parameter.NullableObject[inputOptions.UpdateSnowflakeInputOptionInput]          `json:"snowflake_input_option,omitempty"`
	SalesforceInputOption         *parameter.NullableObject[inputOptions.UpdateSalesforceInputOptionInput]         `json:"salesforce_input_option,omitempty"`
	GoogleSpreadsheetsInputOption *parameter.NullableObject[inputOptions.UpdateGoogleSpreadsheetsInputOptionInput] `json:"google_spreadsheets_input_option,omitempty"`
	S3InputOption                 *parameter.NullableObject[inputOptions.UpdateS3InputOptionInput]                 `json:"s3_input_option,omitempty"`
	BigqueryInputOption           *parameter.NullableObject[inputOptions.UpdateBigqueryInputOptionInput]           `json:"bigquery_input_option,omitempty"`
	PostgreSQLInputOption         *parameter.NullableObject[inputOptions.UpdatePostgreSQLInputOptionInput]         `json:"postgresql_input_option,omitempty"`
	GoogleAnalytics4InputOption   *parameter.NullableObject[inputOptions.UpdateGoogleAnalytics4InputOptionInput]   `json:"google_analytics4_input_option,omitempty"`
	HTTPInputOption               *parameter.NullableObject[inputOptions.UpdateHTTPInputOptionInput]               `json:"http_input_option,omitempty"`
	KintoneInputOption            *parameter.NullableObject[inputOptions.UpdateKintoneInputOptionInput]            `json:"kintone_input_option,omitempty"`
	YahooAdsApiYssInputOption     *parameter.NullableObject[inputOptions.UpdateYahooAdsApiYssInputOptionInput]     `json:"yahoo_ads_api_yss_input_option,omitempty"`
}

type OutputOption struct {
	BigQueryOutputOption           *outputOptionEntities.BigQueryOutputOption           `json:"bigquery_output_option"`
	SnowflakeOutputOption          *outputOptionEntities.SnowflakeOutputOption          `json:"snowflake_output_option"`
	SalesforceOutputOption         *outputOptionEntities.SalesforceOutputOption         `json:"salesforce_output_option"`
	GoogleSpreadsheetsOutputOption *outputOptionEntities.GoogleSpreadsheetsOutputOption `json:"google_spreadsheets_output_option"`
}

type OutputOptionInput struct {
	BigQueryOutputOption           *parameter.NullableObject[outputOptions.BigQueryOutputOptionInput]           `json:"bigquery_output_option,omitempty"`
	SnowflakeOutputOption          *parameter.NullableObject[outputOptions.SnowflakeOutputOptionInput]          `json:"snowflake_output_option,omitempty"`
	SalesforceOutputOption         *parameter.NullableObject[outputOptions.SalesforceOutputOptionInput]         `json:"salesforce_output_option,omitempty"`
	GoogleSpreadsheetsOutputOption *parameter.NullableObject[outputOptions.GoogleSpreadsheetsOutputOptionInput] `json:"google_spreadsheets_output_option,omitempty"`
}

type UpdateOutputOptionInput struct {
	BigQueryOutputOption           *parameter.NullableObject[outputOptions.UpdateBigQueryOutputOptionInput]           `json:"bigquery_output_option,omitempty"`
	SnowflakeOutputOption          *parameter.NullableObject[outputOptions.UpdateSnowflakeOutputOptionInput]          `json:"snowflake_output_option,omitempty"`
	SalesforceOutputOption         *parameter.NullableObject[outputOptions.UpdateSalesforceOutputOptionInput]         `json:"salesforce_output_option,omitempty"`
	GoogleSpreadsheetsOutputOption *parameter.NullableObject[outputOptions.UpdateGoogleSpreadsheetsOutputOptionInput] `json:"google_spreadsheets_output_option,omitempty"`
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
