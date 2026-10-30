package client

import (
	"fmt"
	"net/http"
	"terraform-provider-trocco/internal/client/entity"
	"terraform-provider-trocco/internal/client/parameter"
)

const dbtJobDefinitionBasePath = "/api/dbt_job_definitions"

type DbtJobDefinition = entity.DbtJobDefinition
type CreateDbtJobDefinitionInput = parameter.CreateDbtJobDefinitionInput
type UpdateDbtJobDefinitionInput = parameter.UpdateDbtJobDefinitionInput

func (client *TroccoClient) GetDbtJobDefinition(id int64) (*DbtJobDefinition, error) {
	path := fmt.Sprintf(dbtJobDefinitionBasePath+"/%d", id)
	output := new(DbtJobDefinition)
	if err := client.do(http.MethodGet, path, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (client *TroccoClient) CreateDbtJobDefinition(input *CreateDbtJobDefinitionInput) (*DbtJobDefinition, error) {
	output := new(DbtJobDefinition)
	if err := client.do(http.MethodPost, dbtJobDefinitionBasePath, input, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (client *TroccoClient) UpdateDbtJobDefinition(id int64, input *UpdateDbtJobDefinitionInput) (*DbtJobDefinition, error) {
	path := fmt.Sprintf(dbtJobDefinitionBasePath+"/%d", id)
	output := new(DbtJobDefinition)
	if err := client.do(http.MethodPatch, path, input, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (client *TroccoClient) DeleteDbtJobDefinition(id int64) error {
	path := fmt.Sprintf(dbtJobDefinitionBasePath+"/%d", id)
	return client.do(http.MethodDelete, path, nil, nil)
}
