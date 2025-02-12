package client

import (
	"fmt"
	"net/http"
	"net/url"
	"terraform-provider-trocco/internal/client/parameter"
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

	// GCS Fields
	ApplicationName     *string `json:"application_name"`
	ServiceAccountEmail *string `json:"service_account_email"`

	// MySQL Fields
	Port           *int64 `json:"port"`
	SSL            *bool  `json:"ssl"`
	GatewayEnabled *bool  `json:"gateway_enabled"`

	// S3 Fields
	AWSAuthType            *string `json:"aws_auth_type,omitempty"`
	AWSAccessKeyID         *string `json:"aws_access_key_id,omitempty"`
	AWSSecretAccessKey     *string `json:"aws_secret_access_key,omitempty"`
	AWSAssumeRoleAccountID *string `json:"aws_assume_role_account_id,omitempty"`
	AWSAssumeRoleName      *string `json:"aws_assume_role_name,omitempty"`

	// Salesforce Fields
	AuthEndPoint *string `json:"auth_end_point"`
}

type GetConnectionsInput struct {
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type CreateConnectionInput struct {
	// Common Fields
	Name            string                   `json:"name"`
	Description     *string                  `json:"description,omitempty"`
	ResourceGroupID *parameter.NullableInt64 `json:"resource_group_id,omitempty"`

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

	// GCS Fields
	ApplicationName     *string `json:"application_name,omitempty"`
	ServiceAccountEmail *string `json:"service_account_email,omitempty"`

	// MySQL Fields
	Port                 *parameter.NullableInt64 `json:"port,omitempty"`
	SSL                  *parameter.NullableBool  `json:"ssl,omitempty"`
	SSLCA                *string                  `json:"ssl_ca,omitempty"`
	SSLCert              *string                  `json:"ssl_cert,omitempty"`
	SSLKey               *string                  `json:"ssl_key,omitempty"`
	GatewayEnabled       *parameter.NullableBool  `json:"gateway_enabled,omitempty"`
	GatewayHost          *string                  `json:"gateway_host,omitempty"`
	GatewayPort          *parameter.NullableInt64 `json:"gateway_port,omitempty"`
	GatewayUserName      *string                  `json:"gateway_user_name,omitempty"`
	GatewayPassword      *string                  `json:"gateway_password,omitempty"`
	GatewayKey           *string                  `json:"gateway_key,omitempty"`
	GatewayKeyPassphrase *string                  `json:"gateway_key_passphrase,omitempty"`

	// S3 Fields
	AWSAuthType            *string `json:"aws_auth_type,omitempty"`
	AWSAccessKeyID         *string `json:"aws_access_key_id,omitempty"`
	AWSSecretAccessKey     *string `json:"aws_secret_access_key,omitempty"`
	AWSAssumeRoleAccountID *string `json:"aws_assume_role_account_id,omitempty"`
	AWSAssumeRoleName      *string `json:"aws_assume_role_name,omitempty"`

	// Salesforce Fields
	SecurityToken *string `json:"security_token,omitempty"`
	AuthEndPoint  *string `json:"auth_end_point,omitempty"`
}

type UpdateConnectionInput struct {
	// Common Fields
	Name            *string                  `json:"name,omitempty"`
	Description     *string                  `json:"description,omitempty"`
	ResourceGroupID *parameter.NullableInt64 `json:"resource_group_id,omitempty"`

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

	// GCS Fields
	ApplicationName     *string `json:"application_name,omitempty"`
	ServiceAccountEmail *string `json:"service_account_email,omitempty"`

	// MySQL Fields
	Port                 *parameter.NullableInt64 `json:"port,omitempty"`
	SSL                  *parameter.NullableBool  `json:"ssl,omitempty"`
	SSLCA                *string                  `json:"ssl_ca,omitempty"`
	SSLCert              *string                  `json:"ssl_cert,omitempty"`
	SSLKey               *string                  `json:"ssl_key,omitempty"`
	GatewayEnabled       *parameter.NullableBool  `json:"gateway_enabled,omitempty"`
	GatewayHost          *string                  `json:"gateway_host,omitempty"`
	GatewayPort          *parameter.NullableInt64 `json:"gateway_port,omitempty"`
	GatewayUserName      *string                  `json:"gateway_user_name,omitempty"`
	GatewayPassword      *string                  `json:"gateway_password,omitempty"`
	GatewayKey           *string                  `json:"gateway_key,omitempty"`
	GatewayKeyPassphrase *string                  `json:"gateway_key_passphrase,omitempty"`

	// S3 Fields
	AWSAuthType            *string `json:"aws_auth_type,omitempty"`
	AWSAccessKeyID         *string `json:"aws_access_key_id,omitempty"`
	AWSSecretAccessKey     *string `json:"aws_secret_access_key,omitempty"`
	AWSAssumeRoleAccountID *string `json:"aws_assume_role_account_id,omitempty"`
	AWSAssumeRoleName      *string `json:"aws_assume_role_name,omitempty"`

	// Salesforce Fields
	SecurityToken *string `json:"security_token,omitempty"`
	AuthEndPoint  *string `json:"auth_end_point,omitempty"`
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
