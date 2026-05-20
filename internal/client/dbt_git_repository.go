package client

import (
	"fmt"
	"net/http"
	"terraform-provider-trocco/internal/client/entity"
)

const dbtGitRepositoryBasePath = "/api/dbt_git_repositories"

type DbtGitRepository = entity.DbtGitRepository

// Get a DbtGitRepository

func (client *TroccoClient) GetDbtGitRepository(id int64) (*DbtGitRepository, error) {
	path := fmt.Sprintf(dbtGitRepositoryBasePath+"/%d", id)
	output := new(DbtGitRepository)
	if err := client.do(http.MethodGet, path, nil, output); err != nil {
		return nil, err
	}
	return output, nil
}

// Create a DbtGitRepository

type CreateDbtGitRepositoryInput struct {
	Name            string  `json:"name"`
	Description     *string `json:"description,omitempty"`
	AdapterType     string  `json:"adapter_type"`
	DbtVersion      string  `json:"dbt_version"`
	URL             string  `json:"url"`
	Branch          string  `json:"branch"`
	Subdirectory    *string `json:"subdirectory,omitempty"`
	ResourceGroupID *int64  `json:"resource_group_id,omitempty"`
}

func (client *TroccoClient) CreateDbtGitRepository(input *CreateDbtGitRepositoryInput) (*DbtGitRepository, error) {
	output := new(DbtGitRepository)
	if err := client.do(http.MethodPost, dbtGitRepositoryBasePath, input, output); err != nil {
		return nil, err
	}
	return output, nil
}

// Update a DbtGitRepository

// Note: adapter_type cannot be modified after creation and is intentionally omitted.
type UpdateDbtGitRepositoryInput struct {
	Name            *string `json:"name,omitempty"`
	Description     *string `json:"description,omitempty"`
	DbtVersion      *string `json:"dbt_version,omitempty"`
	URL             *string `json:"url,omitempty"`
	Branch          *string `json:"branch,omitempty"`
	Subdirectory    *string `json:"subdirectory,omitempty"`
	ResourceGroupID *int64  `json:"resource_group_id,omitempty"`
}

func (client *TroccoClient) UpdateDbtGitRepository(id int64, input *UpdateDbtGitRepositoryInput) (*DbtGitRepository, error) {
	path := fmt.Sprintf(dbtGitRepositoryBasePath+"/%d", id)
	output := new(DbtGitRepository)
	if err := client.do(http.MethodPatch, path, input, output); err != nil {
		return nil, err
	}
	return output, nil
}

// Delete a DbtGitRepository

func (client *TroccoClient) DeleteDbtGitRepository(id int64) error {
	path := fmt.Sprintf(dbtGitRepositoryBasePath+"/%d", id)
	return client.do(http.MethodDelete, path, nil, nil)
}
