package client

import (
	"fmt"
	"net/http"
	"terraform-provider-trocco/internal/client/parameters/job_definitions/filter"
	"terraform-provider-trocco/internal/client/parameters/job_definitions/input_options"
	"terraform-provider-trocco/internal/client/parameters/job_definitions/output_options"
)

type JobDefinition struct {
	ID                        int64                             `json:"id"`
	Name                      string                            `json:"name"`
	Description               *string                           `json:"description"`
	ResourceGroupId           *bool                             `json:"resource_group_id"`
	IsRunnableConcurrently    *bool                             `json:"is_runnable_concurrently"`
	RetryLimit                int64                             `json:"retry_limit"`
	ResourceEnhancement       *string                           `json:"resource_enhancement"`
	FilterColumns             []filter.FilterColumn             `json:"filter_columns"`
	FilterRows                *filter.FilterRows                `json:"filter_rows"`
	FilterMasks               []filter.FilterMask               `json:"filter_masks"`
	FilterAddTime             *filter.FilterAddTime             `json:"filter_add_time"`
	FilterGsub                []filter.FilterGsub               `json:"filter_gsub"`
	FilterStringTransforms    []filter.FilterStringTransform    `json:"filter_string_transforms"`
	FilterHashes              []filter.FilterHash               `json:"filter_hashes"`
	FilterUnixTimeConversions []filter.FilterUnixTimeConversion `json:"filter_unixtime_conversions"`
	InputOptionType           string                            `json:"input_option_type"`
	InputOption               InputOptionInput                  `json:"input_option"`
	OutputOptionType          string                            `json:"output_option_type"`
	OutputOption              OutputOptionInput                 `json:"output_option"`
}

type CreateJobDefinitionInput struct {
	Name                      string                            `json:"name"`
	Description               *string                           `json:"description,omitempty"`
	ResourceGroupId           *bool                             `json:"resource_group_id,omitempty"`
	IsRunnableConcurrently    *bool                             `json:"is_runnable_concurrently"`
	RetryLimit                int64                             `json:"retry_limit"`
	ResourceEnhancement       *string                           `json:"resource_enhancement,omitempty"`
	FilterColumns             []filter.FilterColumnInput        `json:"filter_columns"`
	FilterRows                *filter.FilterRows                `json:"filter_rows,omitempty"`
	FilterMasks               []filter.FilterMaskInput          `json:"filter_masks"`
	FilterAddTime             *filter.FilterAddTime             `json:"filter_add_time,omitempty"`
	FilterGsub                []filter.FilterGsub               `json:"filter_gsub"`
	FilterStringTransforms    []filter.FilterStringTransform    `json:"filter_string_transforms"`
	FilterHashes              []filter.FilterHash               `json:"filter_hashes"`
	FilterUnixTimeConversions []filter.FilterUnixTimeConversion `json:"filter_unixtime_conversions"`
	InputOptionType           string                            `json:"input_option_type"`
	InputOption               InputOptionInput                  `json:"input_option"`
	OutputOptionType          string                            `json:"output_option_type"`
	OutputOption              OutputOptionInput                 `json:"output_option"`
}

type UpdateJobDefinitionInput struct {
	Name                      *string                            `json:"name,omitempty"`
	Description               *string                            `json:"description,omitempty"`
	ResourceGroupId           *bool                              `json:"resource_group_id,omitempty"`
	IsRunnableConcurrently    *bool                              `json:"is_runnable_concurrently,omitempty"`
	RetryLimit                *int64                             `json:"retry_limit,omitempty"`
	ResourceEnhancement       *string                            `json:"resource_enhancement,omitempty"`
	FilterColumns             *[]filter.FilterColumnInput        `json:"filter_columns,omitempty"`
	FilterRows                *filter.FilterRows                 `json:"filter_rows,omitempty"`
	FilterMasks               *[]filter.FilterMaskInput          `json:"filter_masks,omitempty"`
	FilterAddTime             *filter.FilterAddTime              `json:"filter_add_time,omitempty"`
	FilterGsub                *[]filter.FilterGsub               `json:"filter_gsub,omitempty"`
	FilterStringTransforms    *[]filter.FilterStringTransform    `json:"filter_string_transforms,omitempty"`
	FilterHashes              *[]filter.FilterHash               `json:"filter_hashes,omitempty"`
	FilterUnixTimeConversions *[]filter.FilterUnixTimeConversion `json:"filter_unixtime_conversions,omitempty"`
	InputOption               *UpdateInputOptionInput            `json:"input_option,omitempty"`
	OutputOption              *UpdateOutputOptionInput           `json:"output_option,omitempty"`
}

type InputOption struct {
	MySQLInputOption *input_options.MySQLInputOption `json:"mysql_input_option"`
	GcsInputOption   *input_options.GcsInputOption   `json:"gcs_input_option"`
}

type InputOptionInput struct {
	MySQLInputOption *input_options.MySQLInputOptionInput `json:"mysql_input_option,omitempty"`
	GcsInputOption   *input_options.GcsInputOptionInput   `json:"gcs_input_option,omitempty"`
}

type UpdateInputOptionInput struct {
	MySQLInputOption *input_options.UpdateMySQLInputOptionInput `json:"mysql_input_option,omitempty"`
	GcsInputOption   *input_options.UpdateGcsInputOptionInput   `json:"gcs_input_option,omitempty"`
}

type OutputOption struct {
	BigQueryOutputOption *output_options.BigQueryOutputOption `json:"bigquery_output_option"`
}

type OutputOptionInput struct {
	BigQueryOutputOption *output_options.BigQueryOutputOptionInput `json:"bigquery_output_option,omitempty"`
}

type UpdateOutputOptionInput struct {
	BigQueryOutputOption *output_options.UpdateBigQueryOutputOptionInput `json:"bigquery_output_option,omitempty"`
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
