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
	ID                       int64   `json:"id"`                                   // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone
	Name                     *string `json:"name"`                                 // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone
	Description              *string `json:"description"`                          // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone
	ResourceGroupID          *int64  `json:"resource_group_id"`                    // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone
	ProjectID                *string `json:"project_id"`                           // bigquery, gcs
	IsOAuth                  *bool   `json:"is_oauth"`                             // bigquery, gcs, google_spreadsheets, google_analytics4 (read-only)
	HasServiceAccountJSONKey *bool   `json:"has_service_account_json_key"`         // bigquery, gcs, google_spreadsheets, google_analytics4 (read-only)
	GoogleOAuth2CredentialID *int64  `json:"google_oauth2_credential_id"`          // bigquery, gcs, google_spreadsheets, google_analytics4 (read-only)
	Host                     *string `json:"host"`                                 // snowflake, mysql, postgresql, sftp
	UserName                 *string `json:"user_name"`                            // snowflake, mysql, postgresql, salesforce, sftp
	Role                     *string `json:"role"`                                 // snowflake
	AuthMethod               *string `json:"auth_method"`                          // snowflake
	AWSPrivatelinkEnabled    *bool   `json:"aws_privatelink_enabled"`              // snowflake, sftp (read-only)
	Driver                   *string `json:"driver"`                               // mysql, postgresql, snowflake
	ApplicationName          *string `json:"application_name"`                     // gcs
	ServiceAccountEmail      *string `json:"service_account_email"`                // gcs
	Port                     *int64  `json:"port"`                                 // mysql, postgresql, sftp
	SSL                      *bool   `json:"ssl"`                                  // mysql, postgresql
	GatewayEnabled           *bool   `json:"gateway_enabled"`                      // mysql, postgresql
	AuthEndPoint             *string `json:"auth_end_point"`                       // salesforce
	AWSAuthType              *string `json:"aws_auth_type,omitempty"`              // s3
	AWSAccessKeyID           *string `json:"aws_access_key_id,omitempty"`          // s3
	AWSSecretAccessKey       *string `json:"aws_secret_access_key,omitempty"`      // s3
	AWSAssumeRoleAccountID   *string `json:"aws_assume_role_account_id,omitempty"` // s3
	AWSAssumeRoleName        *string `json:"aws_assume_role_name,omitempty"`       // s3
	Domain                   *string `json:"domain"`                               // kintone
	LoginMethod              *string `json:"login_method"`                         // kintone
	Username                 *string `json:"username"`                             // kintone
	BasicAuthUsername        *string `json:"basic_auth_username"`                  // kintone
	SecretKey                *string `json:"secret_key"`                           // sftp
	SecretKeyPassphrase      *string `json:"secret_key_passphrase"`                // sftp
	UserDirectoryIsRoot      *bool   `json:"user_directory_is_root"`               // sftp
	WindowsServer            *bool   `json:"windows_server"`                       // sftp
	SSHTunnelID              *int64  `json:"ssh_tunnel_id"`                        // sftp
}

type GetConnectionsInput struct {
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type CreateConnectionInput struct {
	Name                   string                    `json:"name"`                                 // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone, sftp
	Description            *string                   `json:"description,omitempty"`                // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone, sftp
	ResourceGroupID        *parameter.NullableInt64  `json:"resource_group_id,omitempty"`          // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone, sftp
	ProjectID              *string                   `json:"project_id,omitempty"`                 // bigquery, gcs
	ServiceAccountJSONKey  *string                   `json:"service_account_json_key,omitempty"`   // bigquery, gcs, google_spreadsheets, google_analytics4
	Host                   *string                   `json:"host,omitempty"`                       // snowflake, mysql, postgresql, sftp
	UserName               *string                   `json:"user_name,omitempty"`                  // snowflake, mysql, postgresql, salesforce, sftp
	Role                   *string                   `json:"role,omitempty"`                       // snowflake
	AuthMethod             *string                   `json:"auth_method,omitempty"`                // snowflake
	Password               *string                   `json:"password,omitempty"`                   // snowflake, mysql, postgresql, salesforce, kintone, sftp
	PrivateKey             *string                   `json:"private_key,omitempty"`                // snowflake
	ApplicationName        *string                   `json:"application_name,omitempty"`           // gcs
	ServiceAccountEmail    *string                   `json:"service_account_email,omitempty"`      // gcs
	Port                   *parameter.NullableInt64  `json:"port,omitempty"`                       // mysql, postgresql, sftp
	SSL                    *parameter.NullableBool   `json:"ssl,omitempty"`                        // mysql, postgresql
	SSLCA                  *string                   `json:"ssl_ca,omitempty"`                     // mysql, postgresql
	SSLCert                *string                   `json:"ssl_cert,omitempty"`                   // mysql, postgresql
	SSLKey                 *string                   `json:"ssl_key,omitempty"`                    // mysql, postgresql
	GatewayEnabled         *parameter.NullableBool   `json:"gateway_enabled,omitempty"`            // mysql, postgresql
	GatewayHost            *string                   `json:"gateway_host,omitempty"`               // mysql, postgresql
	GatewayPort            *parameter.NullableInt64  `json:"gateway_port,omitempty"`               // mysql, postgresql
	GatewayUserName        *string                   `json:"gateway_user_name,omitempty"`          // mysql, postgresql
	GatewayPassword        *string                   `json:"gateway_password,omitempty"`           // mysql, postgresql
	GatewayKey             *string                   `json:"gateway_key,omitempty"`                // mysql, postgresql
	GatewayKeyPassphrase   *string                   `json:"gateway_key_passphrase,omitempty"`     // mysql, postgresql
	SecurityToken          *string                   `json:"security_token,omitempty"`             // salesforce
	AuthEndPoint           *string                   `json:"auth_end_point,omitempty"`             // salesforce
	AWSAuthType            *string                   `json:"aws_auth_type,omitempty"`              // s3
	AWSAccessKeyID         *string                   `json:"aws_access_key_id,omitempty"`          // s3
	AWSSecretAccessKey     *string                   `json:"aws_secret_access_key,omitempty"`      // s3
	AWSAssumeRoleAccountID *string                   `json:"aws_assume_role_account_id,omitempty"` // s3
	AWSAssumeRoleName      *string                   `json:"aws_assume_role_name,omitempty"`       // s3
	SSLClientCa            *string                   `json:"ssl_client_ca,omitempty"`              // postgresql
	SSLClientKey           *string                   `json:"ssl_client_key,omitempty"`             // postgresql
	SSLMode                *parameter.NullableString `json:"ssl_mode,omitempty"`                   // postgresql
	Driver                 *parameter.NullableString `json:"driver,omitempty"`                     // mysql, postgresql, snowflake
	Domain                 *string                   `json:"domain,omitempty"`                     // kintone
	LoginMethod            *string                   `json:"login_method,omitempty"`               // kintone
	Token                  *string                   `json:"token,omitempty"`                      // kintone
	Username               *parameter.NullableString `json:"username,omitempty"`                   // kintone
	BasicAuthUsername      *parameter.NullableString `json:"basic_auth_username,omitempty"`        // kintone
	BasicAuthPassword      *parameter.NullableString `json:"basic_auth_password,omitempty"`        // kintone
	SecretKey              *string                   `json:"secret_key,omitempty"`                 // sftp
	SecretKeyPassphrase    *string                   `json:"secret_key_passphrase,omitempty"`      // sftp
	UserDirectoryIsRoot    *parameter.NullableBool   `json:"user_directory_is_root,omitempty"`     // sftp
	WindowsServer          *parameter.NullableBool   `json:"windows_server,omitempty"`             // sftp
	SSHTunnelID            *parameter.NullableInt64  `json:"ssh_tunnel_id,omitempty"`              // sftp
	AWSPrivatelinkEnabled  *parameter.NullableBool   `json:"aws_privatelink_enabled,omitempty"`    // sftp
}

type UpdateConnectionInput struct {
	Name                   *string                   `json:"name,omitempty"`                       // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone, sftp
	Description            *string                   `json:"description,omitempty"`                // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone, sftp
	ResourceGroupID        *parameter.NullableInt64  `json:"resource_group_id,omitempty"`          // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone, sftp
	ProjectID              *string                   `json:"project_id,omitempty"`                 // bigquery, gcs
	ServiceAccountJSONKey  *string                   `json:"service_account_json_key"`             // bigquery, gcs, google_spreadsheets, google_analytics4
	Host                   *string                   `json:"host,omitempty"`                       // snowflake, mysql, postgresql, sftp
	UserName               *string                   `json:"user_name,omitempty"`                  // snowflake, mysql, postgresql, salesforce, sftp
	Role                   *string                   `json:"role,omitempty"`                       // snowflake
	AuthMethod             *string                   `json:"auth_method,omitempty"`                // snowflake
	Password               *string                   `json:"password,omitempty"`                   // snowflake, mysql, postgresql, salesforce, kintone, sftp
	PrivateKey             *string                   `json:"private_key,omitempty"`                // snowflake
	ApplicationName        *string                   `json:"application_name,omitempty"`           // gcs
	ServiceAccountEmail    *string                   `json:"service_account_email,omitempty"`      // gcs
	Port                   *parameter.NullableInt64  `json:"port,omitempty"`                       // mysql, postgresql, sftp
	SSL                    *parameter.NullableBool   `json:"ssl,omitempty"`                        // mysql, postgresql
	SSLCA                  *string                   `json:"ssl_ca,omitempty"`                     // mysql, postgresql
	SSLCert                *string                   `json:"ssl_cert,omitempty"`                   // mysql, postgresql
	SSLKey                 *string                   `json:"ssl_key,omitempty"`                    // mysql, postgresql
	GatewayEnabled         *parameter.NullableBool   `json:"gateway_enabled,omitempty"`            // mysql, postgresql
	GatewayHost            *string                   `json:"gateway_host,omitempty"`               // mysql, postgresql
	GatewayPort            *parameter.NullableInt64  `json:"gateway_port,omitempty"`               // mysql, postgresql
	GatewayUserName        *string                   `json:"gateway_user_name,omitempty"`          // mysql, postgresql
	GatewayPassword        *string                   `json:"gateway_password,omitempty"`           // mysql, postgresql
	GatewayKey             *string                   `json:"gateway_key,omitempty"`                // mysql, postgresql
	GatewayKeyPassphrase   *string                   `json:"gateway_key_passphrase,omitempty"`     // mysql, postgresql
	SecurityToken          *string                   `json:"security_token,omitempty"`             // salesforce
	AuthEndPoint           *string                   `json:"auth_end_point,omitempty"`             // salesforce
	AWSAuthType            *string                   `json:"aws_auth_type,omitempty"`              // s3
	AWSAccessKeyID         *string                   `json:"aws_access_key_id,omitempty"`          // s3
	AWSSecretAccessKey     *string                   `json:"aws_secret_access_key,omitempty"`      // s3
	AWSAssumeRoleAccountID *string                   `json:"aws_assume_role_account_id,omitempty"` // s3
	AWSAssumeRoleName      *string                   `json:"aws_assume_role_name,omitempty"`       // s3
	SSLClientCa            *string                   `json:"ssl_client_ca,omitempty"`              // postgresql
	SSLClientKey           *string                   `json:"ssl_client_key,omitempty"`             // postgresql
	SSLMode                *parameter.NullableString `json:"ssl_mode,omitempty"`                   // postgresql
	Driver                 *parameter.NullableString `json:"driver,omitempty"`                     // mysql, postgresql, snowflake
	Domain                 *string                   `json:"domain,omitempty"`                     // kintone
	LoginMethod            *string                   `json:"login_method,omitempty"`               // kintone
	Token                  *string                   `json:"token,omitempty"`                      // kintone
	Username               *parameter.NullableString `json:"username,omitempty"`                   // kintone
	BasicAuthUsername      *parameter.NullableString `json:"basic_auth_username,omitempty"`        // kintone
	BasicAuthPassword      *parameter.NullableString `json:"basic_auth_password,omitempty"`        // kintone
	SecretKey              *string                   `json:"secret_key,omitempty"`                 // sftp
	SecretKeyPassphrase    *string                   `json:"secret_key_passphrase,omitempty"`      // sftp
	UserDirectoryIsRoot    *parameter.NullableBool   `json:"user_directory_is_root,omitempty"`     // sftp
	WindowsServer          *parameter.NullableBool   `json:"windows_server,omitempty"`             // sftp
	SSHTunnelID            *parameter.NullableInt64  `json:"ssh_tunnel_id,omitempty"`              // sftp
	AWSPrivatelinkEnabled  *parameter.NullableBool   `json:"aws_privatelink_enabled,omitempty"`    // sftp
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
