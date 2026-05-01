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

type ReadPreferenceTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Connection struct {
	ID                               int64                  `json:"id"`                                   // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone, mongodb, google_drive, redshift
	Name                             *string                `json:"name"`                                 // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone, mongodb, google_drive, redshift
	Description                      *string                `json:"description"`                          // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone, mongodb, google_drive, redshift
	ResourceGroupID                  *int64                 `json:"resource_group_id"`                    // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone, mongodb, google_drive, redshift
	ProjectID                        *string                `json:"project_id"`                           // bigquery, gcs
	IsOAuth                          *bool                  `json:"is_oauth"`                             // bigquery, gcs, google_spreadsheets, google_analytics4, google_drive (read-only)
	IsWorkloadIdentityFederation     *bool                  `json:"is_workload_identity_federation"`      // bigquery
	HasServiceAccountJSONKey         *bool                  `json:"has_service_account_json_key"`         // bigquery, gcs, google_spreadsheets, google_analytics4, google_drive (read-only)
	GoogleOAuth2CredentialID         *int64                 `json:"google_oauth2_credential_id"`          // bigquery, gcs, google_spreadsheets, google_analytics4, google_drive (read-only)
	WorkloadIdentityFederationConfig interface{}            `json:"workload_identity_federation_config"`  // bigquery
	Host                             *string                `json:"host"`                                 // snowflake, mysql, postgresql, sftp, mongodb, redshift
	UserName                         *string                `json:"user_name"`                            // snowflake, mysql, postgresql, salesforce, sftp, mongodb, redshift
	Role                             *string                `json:"role"`                                 // snowflake
	AuthMethod                       *string                `json:"auth_method"`                          // snowflake, mongodb
	AWSPrivatelinkEnabled            *bool                  `json:"aws_privatelink_enabled"`              // snowflake, sftp, redshift (read-only)
	Driver                           *string                `json:"driver"`                               // mysql, postgresql, snowflake
	ApplicationName                  *string                `json:"application_name"`                     // gcs
	ServiceAccountEmail              *string                `json:"service_account_email"`                // gcs
	Port                             *int64                 `json:"port"`                                 // mysql, postgresql, sftp, mongodb, redshift
	SSL                              *bool                  `json:"ssl"`                                  // mysql, postgresql, redshift
	GatewayEnabled                   *bool                  `json:"gateway_enabled"`                      // mysql, postgresql, mongodb, redshift
	GatewayHost                      *string                `json:"gateway_host"`                         // redshift
	GatewayPort                      *int64                 `json:"gateway_port"`                         // redshift
	GatewayUserName                  *string                `json:"gateway_user_name"`                    // redshift
	AuthEndPoint                     *string                `json:"auth_end_point"`                       // salesforce
	AWSAuthType                      *string                `json:"aws_auth_type,omitempty"`              // s3
	AWSAccessKeyID                   *string                `json:"aws_access_key_id,omitempty"`          // s3, redshift
	AWSSecretAccessKey               *string                `json:"aws_secret_access_key,omitempty"`      // s3
	AWSAssumeRoleAccountID           *string                `json:"aws_assume_role_account_id,omitempty"` // s3
	AWSAssumeRoleName                *string                `json:"aws_assume_role_name,omitempty"`       // s3
	Domain                           *string                `json:"domain"`                               // kintone
	LoginMethod                      *string                `json:"login_method"`                         // kintone
	Username                         *string                `json:"username"`                             // kintone
	BasicAuthUsername                *string                `json:"basic_auth_username"`                  // kintone
	SecretKey                        *string                `json:"secret_key"`                           // sftp
	SecretKeyPassphrase              *string                `json:"secret_key_passphrase"`                // sftp
	UserDirectoryIsRoot              *bool                  `json:"user_directory_is_root"`               // sftp
	WindowsServer                    *bool                  `json:"windows_server"`                       // sftp
	SSHTunnelID                      *int64                 `json:"ssh_tunnel_id"`                        // sftp, redshift
	ServerHostname                   *string                `json:"server_hostname"`                      // databricks
	HttpPath                         *string                `json:"http_path"`                            // databricks
	AuthType                         *string                `json:"auth_type"`                            // databricks
	OAuth2ClientID                   *string                `json:"oauth2_client_id"`                     // databricks
	ConnectionStringFormat           *string                `json:"connection_string_format"`             // mongodb
	ReadPreference                   *string                `json:"read_preference"`                      // mongodb
	AuthSource                       *string                `json:"auth_source"`                          // mongodb
	ReplicaSet                       *string                `json:"replica_set"`                          // mongodb
	ReadPreferenceTags               *[][]ReadPreferenceTag `json:"read_preference_tags"`                 // mongodb
	StrictReadPreferenceTags         *bool                  `json:"strict_read_preference_tags"`          // mongodb
	AccountID                        *string                `json:"account_id"`                           // marketo
	ClientID                         *string                `json:"client_id"`                            // marketo
	HasClientSecret                  *bool                  `json:"has_client_secret"`                    // marketo (read-only)
	APIMaxCallCount                  *int64                 `json:"api_max_call_count"`                   // marketo
}

type GetConnectionsInput struct {
	Limit  int    `json:"limit"`
	Cursor string `json:"cursor"`
}

type CreateConnectionInput struct {
	Name                             string                    `json:"name"`                                          // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone, sftp, mongodb, google_drive, redshift
	Description                      *string                   `json:"description,omitempty"`                         // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone, sftp, mongodb, google_drive, redshift
	ResourceGroupID                  *parameter.NullableInt64  `json:"resource_group_id,omitempty"`                   // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone, sftp, mongodb, google_drive, redshift
	ProjectID                        *string                   `json:"project_id,omitempty"`                          // bigquery, gcs
	ServiceAccountJSONKey            *string                   `json:"service_account_json_key,omitempty"`            // bigquery, gcs, google_spreadsheets, google_analytics4, google_drive
	IsWorkloadIdentityFederation     *bool                     `json:"is_workload_identity_federation,omitempty"`     // bigquery
	WorkloadIdentityFederationConfig interface{}               `json:"workload_identity_federation_config,omitempty"` // bigquery
	Host                             *string                   `json:"host,omitempty"`                                // snowflake, mysql, postgresql, sftp, mongodb, redshift
	UserName                         *string                   `json:"user_name,omitempty"`                           // snowflake, mysql, postgresql, salesforce, sftp, mongodb, redshift
	Role                             *string                   `json:"role,omitempty"`                                // snowflake
	AuthMethod                       *string                   `json:"auth_method,omitempty"`                         // snowflake, mongodb
	Password                         *string                   `json:"password,omitempty"`                            // snowflake, mysql, postgresql, salesforce, kintone, sftp, mongodb, redshift
	PrivateKey                       *string                   `json:"private_key,omitempty"`                         // snowflake
	ApplicationName                  *string                   `json:"application_name,omitempty"`                    // gcs
	ServiceAccountEmail              *string                   `json:"service_account_email,omitempty"`               // gcs
	Port                             *parameter.NullableInt64  `json:"port,omitempty"`                                // mysql, postgresql, sftp, mongodb, redshift
	SSL                              *parameter.NullableBool   `json:"ssl,omitempty"`                                 // mysql, postgresql, redshift
	SSLCA                            *string                   `json:"ssl_ca,omitempty"`                              // mysql, postgresql
	SSLCert                          *string                   `json:"ssl_cert,omitempty"`                            // mysql, postgresql
	SSLKey                           *string                   `json:"ssl_key,omitempty"`                             // mysql, postgresql
	GatewayEnabled                   *parameter.NullableBool   `json:"gateway_enabled,omitempty"`                     // mysql, postgresql, mongodb, redshift
	GatewayHost                      *string                   `json:"gateway_host,omitempty"`                        // mysql, postgresql, mongodb, redshift
	GatewayPort                      *parameter.NullableInt64  `json:"gateway_port,omitempty"`                        // mysql, postgresql, mongodb, redshift
	GatewayUserName                  *string                   `json:"gateway_user_name,omitempty"`                   // mysql, postgresql, mongodb, redshift
	GatewayPassword                  *string                   `json:"gateway_password,omitempty"`                    // mysql, postgresql, mongodb, redshift
	GatewayKey                       *string                   `json:"gateway_key,omitempty"`                         // mysql, postgresql, mongodb, redshift
	GatewayKeyPassphrase             *string                   `json:"gateway_key_passphrase,omitempty"`              // mysql, postgresql, mongodb, redshift
	SecurityToken                    *string                   `json:"security_token,omitempty"`                      // salesforce
	AuthEndPoint                     *string                   `json:"auth_end_point,omitempty"`                      // salesforce
	AWSAuthType                      *string                   `json:"aws_auth_type,omitempty"`                       // s3
	AWSAccessKeyID                   *string                   `json:"aws_access_key_id,omitempty"`                   // s3, redshift
	AWSSecretAccessKey               *string                   `json:"aws_secret_access_key,omitempty"`               // s3, redshift
	AWSAssumeRoleAccountID           *string                   `json:"aws_assume_role_account_id,omitempty"`          // s3
	AWSAssumeRoleName                *string                   `json:"aws_assume_role_name,omitempty"`                // s3
	SSLClientCa                      *string                   `json:"ssl_client_ca,omitempty"`                       // postgresql
	SSLClientKey                     *string                   `json:"ssl_client_key,omitempty"`                      // postgresql
	SSLMode                          *parameter.NullableString `json:"ssl_mode,omitempty"`                            // postgresql
	Driver                           *parameter.NullableString `json:"driver,omitempty"`                              // mysql, postgresql, snowflake
	Domain                           *string                   `json:"domain,omitempty"`                              // kintone
	LoginMethod                      *string                   `json:"login_method,omitempty"`                        // kintone
	Token                            *string                   `json:"token,omitempty"`                               // kintone
	Username                         *parameter.NullableString `json:"username,omitempty"`                            // kintone
	BasicAuthUsername                *parameter.NullableString `json:"basic_auth_username,omitempty"`                 // kintone
	BasicAuthPassword                *parameter.NullableString `json:"basic_auth_password,omitempty"`                 // kintone
	SecretKey                        *string                   `json:"secret_key,omitempty"`                          // sftp
	SecretKeyPassphrase              *string                   `json:"secret_key_passphrase,omitempty"`               // sftp
	UserDirectoryIsRoot              *bool                     `json:"user_directory_is_root,omitempty"`              // sftp
	WindowsServer                    *bool                     `json:"windows_server,omitempty"`                      // sftp
	SSHTunnelID                      *parameter.NullableInt64  `json:"ssh_tunnel_id,omitempty"`                       // sftp, redshift
	AWSPrivatelinkEnabled            *bool                     `json:"aws_privatelink_enabled,omitempty"`             // sftp, redshift
	HttpPath                         *string                   `json:"http_path,omitempty"`                           // databricks
	AuthType                         *string                   `json:"auth_type,omitempty"`                           // databricks
	PersonalAccessToken              *parameter.NullableString `json:"personal_access_token,omitempty"`               // databricks
	OAuth2ClientID                   *parameter.NullableString `json:"oauth2_client_id,omitempty"`                    // databricks
	OAuth2ClientSecret               *parameter.NullableString `json:"oauth2_client_secret,omitempty"`                // databricks
	ServerHostname                   *string                   `json:"server_hostname,omitempty"`                     // databricks
	ConnectionStringFormat           *parameter.NullableString `json:"connection_string_format,omitempty"`            // mongodb
	ReadPreference                   *parameter.NullableString `json:"read_preference,omitempty"`                     // mongodb
	AuthSource                       *parameter.NullableString `json:"auth_source,omitempty"`                         // mongodb
	ReplicaSet                       *parameter.NullableString `json:"replica_set,omitempty"`                         // mongodb
	ReadPreferenceTags               *[][]ReadPreferenceTag    `json:"read_preference_tags,omitempty"`                // mongodb
	StrictReadPreferenceTags         *parameter.NullableBool   `json:"strict_read_preference_tags,omitempty"`         // mongodb
	AccountID                        *string                   `json:"account_id,omitempty"`                          // marketo
	ClientID                         *string                   `json:"client_id,omitempty"`                           // marketo
	ClientSecret                     *string                   `json:"client_secret,omitempty"`                       // marketo
	APIMaxCallCount                  *parameter.NullableInt64  `json:"api_max_call_count,omitempty"`                  // marketo
}

type UpdateConnectionInput struct {
	Name                             *string                   `json:"name,omitempty"`                                // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone, sftp, mongodb, google_drive, redshift
	Description                      *string                   `json:"description,omitempty"`                         // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone, sftp, mongodb, google_drive, redshift
	ResourceGroupID                  *parameter.NullableInt64  `json:"resource_group_id,omitempty"`                   // bigquery, snowflake, gcs, google_spreadsheets, mysql, salesforce, s3, postgresql, google_analytics4, kintone, sftp, mongodb, google_drive, redshift
	ProjectID                        *string                   `json:"project_id,omitempty"`                          // bigquery, gcs
	ServiceAccountJSONKey            *string                   `json:"service_account_json_key"`                      // bigquery, gcs, google_spreadsheets, google_analytics4, google_drive
	IsWorkloadIdentityFederation     *bool                     `json:"is_workload_identity_federation,omitempty"`     // bigquery
	WorkloadIdentityFederationConfig interface{}               `json:"workload_identity_federation_config,omitempty"` // bigquery
	Host                             *string                   `json:"host,omitempty"`                                // snowflake, mysql, postgresql, sftp, mongodb, redshift
	UserName                         *string                   `json:"user_name,omitempty"`                           // snowflake, mysql, postgresql, salesforce, sftp, mongodb, redshift
	Role                             *string                   `json:"role,omitempty"`                                // snowflake
	AuthMethod                       *string                   `json:"auth_method,omitempty"`                         // snowflake, mongodb
	Password                         *string                   `json:"password,omitempty"`                            // snowflake, mysql, postgresql, salesforce, kintone, sftp, mongodb, redshift
	PrivateKey                       *string                   `json:"private_key,omitempty"`                         // snowflake
	ApplicationName                  *string                   `json:"application_name,omitempty"`                    // gcs
	ServiceAccountEmail              *string                   `json:"service_account_email,omitempty"`               // gcs
	Port                             *parameter.NullableInt64  `json:"port,omitempty"`                                // mysql, postgresql, sftp, mongodb, redshift
	SSL                              *parameter.NullableBool   `json:"ssl,omitempty"`                                 // mysql, postgresql, redshift
	SSLCA                            *string                   `json:"ssl_ca,omitempty"`                              // mysql, postgresql
	SSLCert                          *string                   `json:"ssl_cert,omitempty"`                            // mysql, postgresql
	SSLKey                           *string                   `json:"ssl_key,omitempty"`                             // mysql, postgresql
	GatewayEnabled                   *parameter.NullableBool   `json:"gateway_enabled,omitempty"`                     // mysql, postgresql, mongodb, redshift
	GatewayHost                      *string                   `json:"gateway_host,omitempty"`                        // mysql, postgresql, mongodb, redshift
	GatewayPort                      *parameter.NullableInt64  `json:"gateway_port,omitempty"`                        // mysql, postgresql, mongodb, redshift
	GatewayUserName                  *string                   `json:"gateway_user_name,omitempty"`                   // mysql, postgresql, mongodb, redshift
	GatewayPassword                  *string                   `json:"gateway_password,omitempty"`                    // mysql, postgresql, mongodb, redshift
	GatewayKey                       *string                   `json:"gateway_key,omitempty"`                         // mysql, postgresql, mongodb, redshift
	GatewayKeyPassphrase             *string                   `json:"gateway_key_passphrase,omitempty"`              // mysql, postgresql, mongodb, redshift
	SecurityToken                    *string                   `json:"security_token,omitempty"`                      // salesforce
	AuthEndPoint                     *string                   `json:"auth_end_point,omitempty"`                      // salesforce
	AWSAuthType                      *string                   `json:"aws_auth_type,omitempty"`                       // s3
	AWSAccessKeyID                   *string                   `json:"aws_access_key_id,omitempty"`                   // s3, redshift
	AWSSecretAccessKey               *string                   `json:"aws_secret_access_key,omitempty"`               // s3, redshift
	AWSAssumeRoleAccountID           *string                   `json:"aws_assume_role_account_id,omitempty"`          // s3
	AWSAssumeRoleName                *string                   `json:"aws_assume_role_name,omitempty"`                // s3
	SSLClientCa                      *string                   `json:"ssl_client_ca,omitempty"`                       // postgresql
	SSLClientKey                     *string                   `json:"ssl_client_key,omitempty"`                      // postgresql
	SSLMode                          *parameter.NullableString `json:"ssl_mode,omitempty"`                            // postgresql
	Driver                           *parameter.NullableString `json:"driver,omitempty"`                              // mysql, postgresql, snowflake
	Domain                           *string                   `json:"domain,omitempty"`                              // kintone
	LoginMethod                      *string                   `json:"login_method,omitempty"`                        // kintone
	Token                            *string                   `json:"token,omitempty"`                               // kintone
	Username                         *parameter.NullableString `json:"username,omitempty"`                            // kintone
	BasicAuthUsername                *parameter.NullableString `json:"basic_auth_username,omitempty"`                 // kintone
	BasicAuthPassword                *parameter.NullableString `json:"basic_auth_password,omitempty"`                 // kintone
	SecretKey                        *string                   `json:"secret_key,omitempty"`                          // sftp
	SecretKeyPassphrase              *string                   `json:"secret_key_passphrase,omitempty"`               // sftp
	UserDirectoryIsRoot              *bool                     `json:"user_directory_is_root,omitempty"`              // sftp
	WindowsServer                    *bool                     `json:"windows_server,omitempty"`                      // sftp
	SSHTunnelID                      *parameter.NullableInt64  `json:"ssh_tunnel_id,omitempty"`                       // sftp, redshift
	AWSPrivatelinkEnabled            *bool                     `json:"aws_privatelink_enabled,omitempty"`             // sftp, redshift
	HttpPath                         *string                   `json:"http_path,omitempty"`                           // databricks
	AuthType                         *string                   `json:"auth_type,omitempty"`                           // databricks
	PersonalAccessToken              *parameter.NullableString `json:"personal_access_token,omitempty"`               // databricks
	OAuth2ClientID                   *parameter.NullableString `json:"oauth2_client_id,omitempty"`                    // databricks
	OAuth2ClientSecret               *parameter.NullableString `json:"oauth2_client_secret,omitempty"`                // databricks
	ServerHostname                   *string                   `json:"server_hostname,omitempty"`                     // databricks
	ConnectionStringFormat           *parameter.NullableString `json:"connection_string_format,omitempty"`            // mongodb
	ReadPreference                   *parameter.NullableString `json:"read_preference,omitempty"`                     // mongodb
	AuthSource                       *parameter.NullableString `json:"auth_source,omitempty"`                         // mongodb
	ReplicaSet                       *parameter.NullableString `json:"replica_set,omitempty"`                         // mongodb
	ReadPreferenceTags               *[][]ReadPreferenceTag    `json:"read_preference_tags,omitempty"`                // mongodb
	StrictReadPreferenceTags         *parameter.NullableBool   `json:"strict_read_preference_tags,omitempty"`         // mongodb
	AccountID                        *string                   `json:"account_id,omitempty"`                          // marketo
	ClientID                         *string                   `json:"client_id,omitempty"`                           // marketo
	ClientSecret                     *string                   `json:"client_secret,omitempty"`                       // marketo
	APIMaxCallCount                  *parameter.NullableInt64  `json:"api_max_call_count,omitempty"`                  // marketo
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
