package client

import (
	"fmt"
	"net/http"
	"net/url"
	"terraform-provider-trocco/internal/client/entity"
)

const labelBasePath = "/api/labels"

type Label = entity.Label

// List of Labels

type ListLabelsInput struct {
	Limit                *int
	Cursor               *string
	DatamartDefinitionID *int64
	JobDefinitionID      *int64
	JobDefinitionBulkID  *int64
	PipelineDefinitionID *int64
}

func (input *ListLabelsInput) SetLimit(limit int) {
	input.Limit = &limit
}

func (input *ListLabelsInput) SetCursor(cursor string) {
	input.Cursor = &cursor
}

func (input *ListLabelsInput) SetDatamartDefinitionID(id int64) {
	input.DatamartDefinitionID = &id
}

func (input *ListLabelsInput) SetJobDefinitionID(id int64) {
	input.JobDefinitionID = &id
}

func (input *ListLabelsInput) SetJobDefinitionBulkID(id int64) {
	input.JobDefinitionBulkID = &id
}

func (input *ListLabelsInput) SetPipelineDefinitionID(id int64) {
	input.PipelineDefinitionID = &id
}

type ListLabelsOutput struct {
	Items      []Label `json:"items"`
	NextCursor *string `json:"next_cursor"`
}

const MaxListLabelsLimit = 200

func (client *TroccoClient) ListLabels(input *ListLabelsInput) (*ListLabelsOutput, error) {
	params := url.Values{}
	if input != nil && input.Limit != nil {
		if *input.Limit < 1 || *input.Limit > MaxListLabelsLimit {
			return nil, fmt.Errorf("limit must be between 1 and %d", MaxListLabelsLimit)
		}
		params.Add("limit", fmt.Sprintf("%d", *input.Limit))
	}
	if input != nil && input.Cursor != nil {
		params.Add("cursor", *input.Cursor)
	}
	if input != nil && input.DatamartDefinitionID != nil {
		params.Add("datamart_definition_id", fmt.Sprintf("%d", *input.DatamartDefinitionID))
	}
	if input != nil && input.JobDefinitionID != nil {
		params.Add("job_definition_id", fmt.Sprintf("%d", *input.JobDefinitionID))
	}
	if input != nil && input.JobDefinitionBulkID != nil {
		params.Add("job_definition_bulk_id", fmt.Sprintf("%d", *input.JobDefinitionBulkID))
	}
	if input != nil && input.PipelineDefinitionID != nil {
		params.Add("pipeline_definition_id", fmt.Sprintf("%d", *input.PipelineDefinitionID))
	}
	path := fmt.Sprintf(labelBasePath+"?%s", params.Encode())
	output := new(ListLabelsOutput)
	err := client.do(http.MethodGet, path, nil, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Get a Label

func (client *TroccoClient) GetLabel(id int64) (*Label, error) {
	path := fmt.Sprintf(labelBasePath+"/%d", id)
	output := new(Label)
	err := client.do(http.MethodGet, path, nil, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Create a Label

type CreateLabelInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Color       string  `json:"color"`
}

func (client *TroccoClient) CreateLabel(input *CreateLabelInput) (*Label, error) {
	output := new(Label)
	err := client.do(http.MethodPost, labelBasePath, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Update a Label

type UpdateLabelInput struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Color       *string `json:"color,omitempty"`
}

func (client *TroccoClient) UpdateLabel(id int64, input *UpdateLabelInput) (*Label, error) {
	path := fmt.Sprintf(labelBasePath+"/%d", id)
	output := new(Label)
	err := client.do(http.MethodPatch, path, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Delete a Label

func (client *TroccoClient) DeleteLabel(id int64) error {
	path := fmt.Sprintf(labelBasePath+"/%d", id)
	return client.do(http.MethodDelete, path, nil, nil)
}
