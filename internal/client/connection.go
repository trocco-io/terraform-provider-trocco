package client

import (
	"fmt"
	"net/http"
	"net/url"
)

type ConnectionList struct {
	Connections []*Connection `json:"connections"`
	NextCursor  string        `json:"next_cursor"`
}

type Connection struct {
	// Common Fields
	ID              int64   `json:"id"`
	Name            *string `json:"name"`
	Description     *string `json:"description"`
	ResourceGroupID *int64  `json:"resource_group_id"`

	// BigQuery Fields
	ProjectID                *string `json:"project_id"`
	IsOAuth                  *bool   `json:"is_oauth"`
	HasServiceAccountJSONKey *bool   `json:"has_service_account_json_key"`
	GoogleOAuth2CredentialID *int64  `json:"google_oauth2_credential_id"`

	// Snowflake Fields
	Host                  *string `json:"host"`
	UserName              *string `json:"user_name"`
	Role                  *string `json:"role"`
	AuthMethod            *string `json:"auth_method"`
	AWSPrivateLinkEnabled *bool   `json:"aws_privatelink_enabled"`
	Driver                *string `json:"driver"`
}

type GetConnectionsInput struct {
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type CreateConnectionInput struct {
	// Common Fields
	Name            string         `json:"name"`
	Description     *string        `json:"description,omitempty"`
	ResourceGroupID *NullableInt64 `json:"resource_group_id,omitempty"`

	// BigQuery Fields
	ProjectID             *string `json:"project_id,omitempty"`
	ServiceAccountJSONKey *string `json:"service_account_json_key,omitempty"`

	// Snowflake Fields
	Host       *string `json:"host,omitempty"`
	UserName   *string `json:"user_name,omitempty"`
	Role       *string `json:"role,omitempty"`
	AuthMethod *string `json:"auth_method,omitempty"`
	Password   *string `json:"password,omitempty"`
	PrivateKey *string `json:"private_key,omitempty"`
}

type UpdateConnectionInput struct {
	// Common Fields
	Name            *string        `json:"name,omitempty"`
	Description     *string        `json:"description,omitempty"`
	ResourceGroupID *NullableInt64 `json:"resource_group_id,omitempty"`

	// BigQuery Fields
	ProjectID             *string `json:"project_id,omitempty"`
	ServiceAccountJSONKey *string `json:"service_account_json_key"`

	// Snowflake Fields
	Host       *string `json:"host,omitempty"`
	UserName   *string `json:"user_name,omitempty"`
	Role       *string `json:"role,omitempty"`
	AuthMethod *string `json:"auth_method,omitempty"`
	Password   *string `json:"password,omitempty"`
	PrivateKey *string `json:"private_key,omitempty"`
}

func (c *TroccoClient) GetConnections(connectionType string, in *GetConnectionsInput) (*ConnectionList, error) {
	params := url.Values{}
	if in != nil {
		if in.Limit != 0 {
			params.Add("limit", fmt.Sprintf("%d", in.Limit))
		}

		if in.Cursor != "" {
			params.Add("cursor", in.Cursor)
		}
	}

	out := &ConnectionList{}
	if err := c.do(
		http.MethodGet,
		fmt.Sprintf("/api/connections/%s/?%s", connectionType, params.Encode()),
		nil,
		out,
	); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) GetConnection(connectionType string, id int64) (*Connection, error) {
	out := &Connection{}
	if err := c.do(
		http.MethodGet,
		fmt.Sprintf("/api/connections/%s/%d", connectionType, id),
		nil,
		out,
	); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) CreateConnection(connectionType string, in *CreateConnectionInput) (*Connection, error) {
	out := &Connection{}
	if err := c.do(
		http.MethodPost,
		fmt.Sprintf("/api/connections/%s", connectionType),
		in,
		out,
	); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) UpdateConnection(connectionType string, id int64, in *UpdateConnectionInput) (*Connection, error) {
	out := &Connection{}
	if err := c.do(
		http.MethodPatch,
		fmt.Sprintf("/api/connections/%s/%d", connectionType, id),
		in,
		out,
	); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *TroccoClient) DeleteConnection(connectionType string, id int64) error {
	return c.do(
		http.MethodDelete,
		fmt.Sprintf("/api/connections/%s/%d", connectionType, id),
		nil,
		nil,
	)
}
