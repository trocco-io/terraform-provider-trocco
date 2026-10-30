package client

import (
	"fmt"
	"net/http"
	"terraform-provider-trocco/internal/client/entity"
	"terraform-provider-trocco/internal/client/parameter"
)

const customConnectorOutputBasePath = "/api/custom_connector_outputs"

type CustomConnectorOutput = entity.CustomConnectorOutput
type CreateCustomConnectorOutputInput = parameter.CreateCustomConnectorOutputInput
type UpdateCustomConnectorOutputInput = parameter.UpdateCustomConnectorOutputInput

func (client *TroccoClient) GetCustomConnectorOutput(id int64) (*CustomConnectorOutput, error) {
	path := fmt.Sprintf(customConnectorOutputBasePath+"/%d", id)
	output := new(CustomConnectorOutput)
	if err := client.do(http.MethodGet, path, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (client *TroccoClient) CreateCustomConnectorOutput(input *CreateCustomConnectorOutputInput) (*CustomConnectorOutput, error) {
	output := new(CustomConnectorOutput)
	if err := client.do(http.MethodPost, customConnectorOutputBasePath, input, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (client *TroccoClient) UpdateCustomConnectorOutput(id int64, input *UpdateCustomConnectorOutputInput) (*CustomConnectorOutput, error) {
	path := fmt.Sprintf(customConnectorOutputBasePath+"/%d", id)
	output := new(CustomConnectorOutput)
	if err := client.do(http.MethodPatch, path, input, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (client *TroccoClient) DeleteCustomConnectorOutput(id int64) error {
	path := fmt.Sprintf(customConnectorOutputBasePath+"/%d", id)
	return client.do(http.MethodDelete, path, nil, nil)
}
