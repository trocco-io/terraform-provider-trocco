package client

import (
	"fmt"
	"net/http"
	"terraform-provider-trocco/internal/client/entity"
	"terraform-provider-trocco/internal/client/parameter"
)

const customConnectorInputBasePath = "/api/custom_connector_inputs"

type CustomConnectorInput = entity.CustomConnectorInput
type CreateCustomConnectorInputInput = parameter.CreateCustomConnectorInputInput
type UpdateCustomConnectorInputInput = parameter.UpdateCustomConnectorInputInput

func (client *TroccoClient) GetCustomConnectorInput(id int64) (*CustomConnectorInput, error) {
	path := fmt.Sprintf(customConnectorInputBasePath+"/%d", id)
	output := new(CustomConnectorInput)
	if err := client.do(http.MethodGet, path, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (client *TroccoClient) CreateCustomConnectorInput(input *CreateCustomConnectorInputInput) (*CustomConnectorInput, error) {
	output := new(CustomConnectorInput)
	if err := client.do(http.MethodPost, customConnectorInputBasePath, input, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (client *TroccoClient) UpdateCustomConnectorInput(id int64, input *UpdateCustomConnectorInputInput) (*CustomConnectorInput, error) {
	path := fmt.Sprintf(customConnectorInputBasePath+"/%d", id)
	output := new(CustomConnectorInput)
	if err := client.do(http.MethodPatch, path, input, output); err != nil {
		return nil, err
	}
	return output, nil
}

func (client *TroccoClient) DeleteCustomConnectorInput(id int64) error {
	path := fmt.Sprintf(customConnectorInputBasePath+"/%d", id)
	return client.do(http.MethodDelete, path, nil, nil)
}
