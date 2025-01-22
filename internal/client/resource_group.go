package client

import (
	"fmt"
	"net/http"
	"net/url"
)

const resourceGroupBasePath = "/api/resource_groups"

type ResourceGroup struct {
	ID          int64        `json:"id"`
	Name        string       `json:"name"`
	Description *string      `json:"description"`
	CreatedAt   string       `json:"created_at"`
	UpdatedAt   string       `json:"updated_at"`
}

type ResourceGroupWithTeams struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Teams       []ResourceGroupPermission `json:"teams"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
}

type ResourceGroupPermission struct {
	TeamID int64  `json:"team_id"`
	Role   string `json:"role"`
}

// List of ResourceGroups

type ListResouceGroupInput struct {
	limit  *int
	cursor *string
}

func (input *ListResouceGroupInput) SetLimit(limit int) {
	input.limit = &limit
}

func (input *ListResouceGroupInput) SetCursor(cursor string) {
	input.cursor = &cursor
}

type ListResouceGroupOutput struct {
	Items      []ResourceGroupWithTeams `json:"items"`
	NextCursor *string              `json:"next_cursor"`
}

const MaxListResourceGroupsLimit = 100

func (client *TroccoClient) ListResourceGroups(input *ListResouceGroupInput) (*ListResouceGroupOutput, error) {
	params := url.Values{}
	if input != nil && input.limit != nil {
		if *input.limit < 1 || *input.limit > MaxListResourceGroupsLimit {
			return nil, fmt.Errorf("limit must be between 1 and %d", MaxListResourceGroupsLimit)
		}
		params.Add("limit", fmt.Sprintf("%d", *input.limit))
	}
	if input != nil && input.cursor != nil {
		params.Add("cursor", *input.cursor)
	}
	path := fmt.Sprintf(resourceGroupBasePath+"?%s", params.Encode())
	output := new(ListResouceGroupOutput)
	err := client.do(http.MethodGet, path, nil, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Get a Team

func (client *TroccoClient) GetResourceGroup(id int64) (*ResourceGroupWithTeams, error) {
	path := fmt.Sprintf("%s/%d", resourceGroupBasePath, id)
	output := new(ResourceGroupWithTeams)
	err := client.do(http.MethodGet, path, nil, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Create

type CreateResourceGroupInput struct {
	Name        string        `json:"name"`
	Description *string       `json:"description,omitempty"`
	Teams     []TeamRoleInput `json:"teams"`
}

type TeamRoleInput struct {
	TeamID int64  `json:"team_id"`
	Role   string `json:"role"`
}

func (client *TroccoClient) CreateResourceGroup(input *CreateResourceGroupInput) (*ResourceGroupWithTeams, error) {
	output := new(ResourceGroupWithTeams)
	err := client.do(http.MethodPost, resourceGroupBasePath, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Update

type UpdateResourceGroupInput struct {
	Name        string        `json:"name"`
	Description *string       `json:"description,omitempty"`
	Teams     []TeamRoleInput `json:"teams"`
}

func (client *TroccoClient) UpdateResourceGroup(id int64, input *UpdateResourceGroupInput) (*ResourceGroupWithTeams, error) {
	path := fmt.Sprintf("%s/%d", resourceGroupBasePath, id)
	output := new(ResourceGroupWithTeams)
	err := client.do(http.MethodPatch, path, input, output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Delete

func (client *TroccoClient) DeleteResourceGroup(id int64) error {
	path := fmt.Sprintf("%s/%d", resourceGroupBasePath, id)
	return client.do(http.MethodDelete, path, nil, nil)
}
