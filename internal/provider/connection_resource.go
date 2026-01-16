package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/connection"

	planModifier "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"
)

var (
	_ resource.Resource                = &connectionResource{}
	_ resource.ResourceWithConfigure   = &connectionResource{}
	_ resource.ResourceWithImportState = &connectionResource{}
)

type connectionResourceModel struct {
	// Common Fields
	ConnectionType  types.String `tfsdk:"connection_type"`
	ID              types.Int64  `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	Description     types.String `tfsdk:"description"`
	ResourceGroupID types.Int64  `tfsdk:"resource_group_id"`

	// BigQuery Fields
	ProjectID             types.String `tfsdk:"project_id"`
	ServiceAccountJSONKey types.String `tfsdk:"service_account_json_key"`

	// Snowflake Fields
	Host       types.String `tfsdk:"host"`
	UserName   types.String `tfsdk:"user_name"`
	Role       types.String `tfsdk:"role"`
	AuthMethod types.String `tfsdk:"auth_method"`
	Password   types.String `tfsdk:"password"`
	PrivateKey types.String `tfsdk:"private_key"`

	// GCS Fields
	ApplicationName     types.String `tfsdk:"application_name"`
	ServiceAccountEmail types.String `tfsdk:"service_account_email"`

	// MySQL Fields
	Port    types.Int64         `tfsdk:"port"`
	SSL     *connection.SSL     `tfsdk:"ssl"`
	Gateway *connection.Gateway `tfsdk:"gateway"`

	// PostgreSQL Fields
	Driver types.String `tfsdk:"driver"`

	// S3 Fields
	AWSAuthType   types.String              `tfsdk:"aws_auth_type"`
	AWSIAMUser    *connection.AWSIAMUser    `tfsdk:"aws_iam_user"`
	AWSAssumeRole *connection.AWSAssumeRole `tfsdk:"aws_assume_role"`

	// Salesforce Fields
	SecurityToken types.String `tfsdk:"security_token"`
	AuthEndPoint  types.String `tfsdk:"auth_end_point"`

	// Kintone Fields
	Domain            types.String `tfsdk:"domain"`
	LoginMethod       types.String `tfsdk:"login_method"`
	Token             types.String `tfsdk:"token"`
	Username          types.String `tfsdk:"username"`
	BasicAuthUsername types.String `tfsdk:"basic_auth_username"`
	BasicAuthPassword types.String `tfsdk:"basic_auth_password"`

	// SFTP Fields
	SecretKey             types.String `tfsdk:"secret_key"`
	SecretKeyPassphrase   types.String `tfsdk:"secret_key_passphrase"`
	UserDirectoryIsRoot   types.Bool   `tfsdk:"user_directory_is_root"`
	WindowsServer         types.Bool   `tfsdk:"windows_server"`
	SSHTunnelID           types.Int64  `tfsdk:"ssh_tunnel_id"`
	AWSPrivatelinkEnabled types.Bool   `tfsdk:"aws_privatelink_enabled"`
	// Databricks Fields
	ServerHostname      types.String `tfsdk:"server_hostname"`
	HttpPath            types.String `tfsdk:"http_path"`
	AuthType            types.String `tfsdk:"auth_type"`
	PersonalAccessToken types.String `tfsdk:"personal_access_token"`
	OAuth2ClientID      types.String `tfsdk:"oauth2_client_id"`
	OAuth2ClientSecret  types.String `tfsdk:"oauth2_client_secret"`
}

func (m *connectionResourceModel) ToCreateConnectionInput() *client.CreateConnectionInput {
	input := &client.CreateConnectionInput{
		// Common Fields
		Name:            m.Name.ValueString(),
		Description:     m.Description.ValueStringPointer(),
		ResourceGroupID: model.NewNullableInt64(m.ResourceGroupID),

		// BigQuery Fields
		ProjectID:             m.ProjectID.ValueStringPointer(),
		ServiceAccountJSONKey: m.ServiceAccountJSONKey.ValueStringPointer(),

		// Snowflake Fields
		Host:       m.Host.ValueStringPointer(),
		UserName:   m.UserName.ValueStringPointer(),
		Role:       m.Role.ValueStringPointer(),
		AuthMethod: m.AuthMethod.ValueStringPointer(),
		Password:   m.Password.ValueStringPointer(),
		PrivateKey: m.PrivateKey.ValueStringPointer(),

		// GCS Fields
		ApplicationName:     m.ApplicationName.ValueStringPointer(),
		ServiceAccountEmail: m.ServiceAccountEmail.ValueStringPointer(),

		// MySQL Fields
		Port: model.NewNullableInt64(m.Port),

		// Salesforce Fields
		SecurityToken: m.SecurityToken.ValueStringPointer(),
		AuthEndPoint:  m.AuthEndPoint.ValueStringPointer(),

		// S3 Fields
		AWSAuthType: m.AWSAuthType.ValueStringPointer(),

		// PostgreSQL Fields
		Driver: model.NewNullableString(m.Driver),

		// Kintone
		Domain:            m.Domain.ValueStringPointer(),
		LoginMethod:       m.LoginMethod.ValueStringPointer(),
		Token:             m.Token.ValueStringPointer(),
		Username:          model.NewNullableString(m.Username),
		BasicAuthUsername: model.NewNullableString(m.BasicAuthUsername),
		BasicAuthPassword: model.NewNullableString(m.BasicAuthPassword),

		// SFTP Fields
		SecretKey:             m.SecretKey.ValueStringPointer(),
		SecretKeyPassphrase:   m.SecretKeyPassphrase.ValueStringPointer(),
		UserDirectoryIsRoot:   m.UserDirectoryIsRoot.ValueBoolPointer(),
		WindowsServer:         m.WindowsServer.ValueBoolPointer(),
		SSHTunnelID:           model.NewNullableInt64(m.SSHTunnelID),
		AWSPrivatelinkEnabled: m.AWSPrivatelinkEnabled.ValueBoolPointer(),
		// Databricks Fields
		ServerHostname:      m.ServerHostname.ValueStringPointer(),
		HttpPath:            m.HttpPath.ValueStringPointer(),
		AuthType:            m.AuthType.ValueStringPointer(),
		PersonalAccessToken: model.NewNullableString(m.PersonalAccessToken),
		OAuth2ClientID:      model.NewNullableString(m.OAuth2ClientID),
		OAuth2ClientSecret:  model.NewNullableString(m.OAuth2ClientSecret),
	}

	// SSL Fields
	if m.SSL != nil {
		input.SSL = model.NewNullableBool(types.BoolValue(true))
		input.SSLCA = m.SSL.CA.ValueStringPointer()
		input.SSLCert = m.SSL.Cert.ValueStringPointer()
		input.SSLClientCa = m.SSL.Cert.ValueStringPointer()
		input.SSLKey = m.SSL.Key.ValueStringPointer()
		input.SSLClientKey = m.SSL.Key.ValueStringPointer()
		input.SSLMode = model.NewNullableString(m.SSL.SSLMode)
	} else {
		input.SSL = model.NewNullableBool(types.BoolValue(false))
	}

	// Gateway Fields
	if m.Gateway != nil {
		input.GatewayEnabled = model.NewNullableBool(types.BoolValue(true))
		input.GatewayHost = m.Gateway.Host.ValueStringPointer()
		input.GatewayPort = model.NewNullableInt64(m.Gateway.Port)
		input.GatewayUserName = m.Gateway.UserName.ValueStringPointer()
		input.GatewayPassword = m.Gateway.Password.ValueStringPointer()
		input.GatewayKey = m.Gateway.Key.ValueStringPointer()
		input.GatewayKeyPassphrase = m.Gateway.KeyPassphrase.ValueStringPointer()
	} else {
		input.GatewayEnabled = model.NewNullableBool(types.BoolValue(false))
	}

	// AWS IAM User Fields
	if m.AWSIAMUser != nil {
		input.AWSAccessKeyID = m.AWSIAMUser.AccessKeyID.ValueStringPointer()
		input.AWSSecretAccessKey = m.AWSIAMUser.SecretAccessKey.ValueStringPointer()
	}

	// AWS Assume Role Fields
	if m.AWSAssumeRole != nil {
		input.AWSAssumeRoleAccountID = m.AWSAssumeRole.AccountID.ValueStringPointer()
		input.AWSAssumeRoleName = m.AWSAssumeRole.AccountRoleName.ValueStringPointer()
	}

	return input
}

func (m *connectionResourceModel) ToUpdateConnectionInput() *client.UpdateConnectionInput {
	input := &client.UpdateConnectionInput{
		// Common Fields
		Name:            m.Name.ValueStringPointer(),
		Description:     m.Description.ValueStringPointer(),
		ResourceGroupID: model.NewNullableInt64(m.ResourceGroupID),

		// BigQuery Fields
		ProjectID:             m.ProjectID.ValueStringPointer(),
		ServiceAccountJSONKey: m.ServiceAccountJSONKey.ValueStringPointer(),

		// Snowflake Fields
		Host:       m.Host.ValueStringPointer(),
		UserName:   m.UserName.ValueStringPointer(),
		Role:       m.Role.ValueStringPointer(),
		AuthMethod: m.AuthMethod.ValueStringPointer(),
		Password:   m.Password.ValueStringPointer(),
		PrivateKey: m.PrivateKey.ValueStringPointer(),

		// GCS Fields
		ApplicationName:     m.ApplicationName.ValueStringPointer(),
		ServiceAccountEmail: m.ServiceAccountEmail.ValueStringPointer(),

		// MySQL Fields
		Port: model.NewNullableInt64(m.Port),

		// Salesforce Fields
		SecurityToken: m.SecurityToken.ValueStringPointer(),
		AuthEndPoint:  m.AuthEndPoint.ValueStringPointer(),

		// S3 Fields
		AWSAuthType: m.AWSAuthType.ValueStringPointer(),

		// PostgreSQL Fields
		Driver: model.NewNullableString(m.Driver),

		// Kintone
		Domain:            m.Domain.ValueStringPointer(),
		LoginMethod:       m.LoginMethod.ValueStringPointer(),
		Token:             m.Token.ValueStringPointer(),
		Username:          model.NewNullableString(m.Username),
		BasicAuthUsername: model.NewNullableString(m.BasicAuthUsername),
		BasicAuthPassword: model.NewNullableString(m.BasicAuthPassword),

		// SFTP Fields
		SecretKey:             m.SecretKey.ValueStringPointer(),
		SecretKeyPassphrase:   m.SecretKeyPassphrase.ValueStringPointer(),
		UserDirectoryIsRoot:   m.UserDirectoryIsRoot.ValueBoolPointer(),
		WindowsServer:         m.WindowsServer.ValueBoolPointer(),
		SSHTunnelID:           model.NewNullableInt64(m.SSHTunnelID),
		AWSPrivatelinkEnabled: m.AWSPrivatelinkEnabled.ValueBoolPointer(),
		// Databricks Fields
		ServerHostname:      m.ServerHostname.ValueStringPointer(),
		HttpPath:            m.HttpPath.ValueStringPointer(),
		AuthType:            m.AuthType.ValueStringPointer(),
		PersonalAccessToken: model.NewNullableString(m.PersonalAccessToken),
		OAuth2ClientID:      model.NewNullableString(m.OAuth2ClientID),
		OAuth2ClientSecret:  model.NewNullableString(m.OAuth2ClientSecret),
	}

	// SSL Fields
	if m.SSL != nil {
		input.SSL = model.NewNullableBool(types.BoolValue(true))
		input.SSLCA = m.SSL.CA.ValueStringPointer()
		input.SSLCert = m.SSL.Cert.ValueStringPointer()
		input.SSLKey = m.SSL.Key.ValueStringPointer()
		input.SSLClientCa = m.SSL.Cert.ValueStringPointer()
		input.SSLClientKey = m.SSL.Key.ValueStringPointer()
		input.SSLMode = model.NewNullableString(m.SSL.SSLMode)
	} else {
		input.SSL = model.NewNullableBool(types.BoolValue(false))
	}

	// Gateway Fields
	if m.Gateway != nil {
		input.GatewayEnabled = model.NewNullableBool(types.BoolValue(true))
		input.GatewayHost = m.Gateway.Host.ValueStringPointer()
		input.GatewayPort = model.NewNullableInt64(m.Gateway.Port)
		input.GatewayUserName = m.Gateway.UserName.ValueStringPointer()
		input.GatewayPassword = m.Gateway.Password.ValueStringPointer()
		input.GatewayKey = m.Gateway.Key.ValueStringPointer()
		input.GatewayKeyPassphrase = m.Gateway.KeyPassphrase.ValueStringPointer()
	} else {
		input.GatewayEnabled = model.NewNullableBool(types.BoolValue(false))
	}

	// AWS IAM User Fields
	if m.AWSIAMUser != nil {
		input.AWSAccessKeyID = m.AWSIAMUser.AccessKeyID.ValueStringPointer()
		input.AWSSecretAccessKey = m.AWSIAMUser.SecretAccessKey.ValueStringPointer()
	}

	// AWS Assume Role Fields
	if m.AWSAssumeRole != nil {
		input.AWSAssumeRoleAccountID = m.AWSAssumeRole.AccountID.ValueStringPointer()
		input.AWSAssumeRoleName = m.AWSAssumeRole.AccountRoleName.ValueStringPointer()
	}

	return input
}

type connectionResource struct {
	client *client.TroccoClient
}

func NewConnectionResource() resource.Resource {
	return &connectionResource{}
}

func (r *connectionResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = fmt.Sprintf("%s_connection", req.ProviderTypeName)
}

func (r *connectionResource) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}

	c, ok := req.ProviderData.(*client.TroccoClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = c
}

var supportedConnectionTypes = []string{
	"bigquery",
	"snowflake",
	"gcs",
	"google_spreadsheets",
	"mysql",
	"salesforce",
	"s3",
	"postgresql",
	"google_analytics4",
	"kintone",
	"sftp",
	"databricks",
}

func (r *connectionResource) Schema(
	ctx context.Context,
	req resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO connection resource.",
		Attributes: map[string]schema.Attribute{
			// Common Fields
			"connection_type": schema.StringAttribute{
				MarkdownDescription: fmt.Sprintf(
					"The type of the connection. It must be one of %s.",
					strings.Join(
						lo.Map(supportedConnectionTypes, func(s string, _ int) string {
							return fmt.Sprintf("`%s`", s)
						}),
						", ",
					),
				),
				Required: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf(supportedConnectionTypes...),
				},
			},
			"id": schema.Int64Attribute{
				MarkdownDescription: "The ID of the connection.",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the connection.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtMost(255),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the connection.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"resource_group_id": schema.Int64Attribute{
				MarkdownDescription: "The ID of the resource group the connection belongs to.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},

			// BigQuery Fields
			"project_id": schema.StringAttribute{
				MarkdownDescription: "BigQuery, GCS: A GCP project ID.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"service_account_json_key": schema.StringAttribute{
				MarkdownDescription: "BigQuery, Google Sheets, Google Analytics4: A GCP service account key.",
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},

			// Snowflake Fields
			"host": schema.StringAttribute{
				MarkdownDescription: "Snowflake, PostgreSQL: The host of a (Snowflake, PostgreSQL) account.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"user_name": schema.StringAttribute{
				MarkdownDescription: "Snowflake, PostgreSQL: The name of a (Snowflake, PostgreSQL) user.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"role": schema.StringAttribute{
				MarkdownDescription: "Snowflake: A role attached to the Snowflake user.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"auth_method": schema.StringAttribute{
				MarkdownDescription: "Snowflake: The authentication method for the Snowflake user. It must be one of `key_pair` or `user_password`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("key_pair", "user_password"),
				},
			},
			"password": schema.StringAttribute{
				MarkdownDescription: "Snowflake, PostgreSQL: The password for the (Snowflake, PostgreSQL) user.",
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"private_key": schema.StringAttribute{
				MarkdownDescription: "Snowflake: A private key for the Snowflake user.",
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},

			// GCS Fields
			"application_name": schema.StringAttribute{
				MarkdownDescription: "GCS: Application name.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"service_account_email": schema.StringAttribute{
				MarkdownDescription: "GCS: A GCP service account email.",
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},

			// MySQL Fields
			"port": schema.Int64Attribute{
				MarkdownDescription: "MySQL, PostgreSQL: The port of the (MySQL, PostgreSQL) server.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
					int64validator.AtMost(65535),
				},
			},

			"ssl": schema.SingleNestedAttribute{
				MarkdownDescription: "MySQL, PostgreSQL: SSL configuration.",
				Optional:            true,
				Attributes: map[string]schema.Attribute{
					"ca": schema.StringAttribute{
						MarkdownDescription: "MySQL, PostgreSQL: CA certificate",
						Optional:            true,
						Sensitive:           true,
						Validators: []validator.String{
							stringvalidator.UTF8LengthAtLeast(1),
						},
						Computed: true,
						Default:  stringdefault.StaticString(""),
					},
					"cert": schema.StringAttribute{
						MarkdownDescription: "MySQL, PostgreSQL: Certificate (CRT file)",
						Optional:            true,
						Sensitive:           true,
						Validators: []validator.String{
							stringvalidator.UTF8LengthAtLeast(1),
						},
						Computed: true,
						Default:  stringdefault.StaticString(""),
					},
					"key": schema.StringAttribute{
						MarkdownDescription: "MySQL, PostgreSQL: Key (KEY file)",
						Optional:            true,
						Sensitive:           true,
						Validators: []validator.String{
							stringvalidator.UTF8LengthAtLeast(1),
						},
						Computed: true,
						Default:  stringdefault.StaticString(""),
					},
					"ssl_mode": schema.StringAttribute{
						MarkdownDescription: "PostgreSQL: SSL connection mode.",
						Optional:            true,
						Validators: []validator.String{
							stringvalidator.OneOf("require", "verify-ca"),
						},
					},
				},
			},
			"gateway": schema.SingleNestedAttribute{
				MarkdownDescription: "MySQL, PostgreSQL: Whether to connect via SSH",
				Optional:            true,
				Attributes: map[string]schema.Attribute{
					"host": schema.StringAttribute{
						MarkdownDescription: "MySQL, PostgreSQL: SSH Host",
						Optional:            true,
						Sensitive:           true,
						Validators: []validator.String{
							stringvalidator.UTF8LengthAtLeast(1),
						},
					},
					"port": schema.Int64Attribute{
						MarkdownDescription: "MySQL, PostgreSQL: SSH Port",
						Optional:            true,
						Sensitive:           true,
						Validators: []validator.Int64{
							int64validator.AtLeast(1),
							int64validator.AtMost(65535),
						},
					},
					"user_name": schema.StringAttribute{
						MarkdownDescription: "MySQL, PostgreSQL: SSH User",
						Optional:            true,
						Sensitive:           true,
						Validators: []validator.String{
							stringvalidator.UTF8LengthAtLeast(1),
						},
					},
					"password": schema.StringAttribute{
						MarkdownDescription: "MySQL, PostgreSQL, Kintone: SSH Password",
						Optional:            true,
						Computed:            true,
						Sensitive:           true,
						Default:             stringdefault.StaticString(""),
					},
					"key": schema.StringAttribute{
						MarkdownDescription: "MySQL, PostgreSQL: SSH Private Key",
						Optional:            true,
						Computed:            true,
						Sensitive:           true,
						Default:             stringdefault.StaticString(""),
					},
					"key_passphrase": schema.StringAttribute{
						MarkdownDescription: "MySQL, PostgreSQL: SSH Private Key Passphrase",
						Optional:            true,
						Computed:            true,
						Sensitive:           true,
						Default:             stringdefault.StaticString(""),
					},
				},
			},

			// Salesforce Fields
			"security_token": schema.StringAttribute{
				MarkdownDescription: "Salesforce: Security token.",
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"auth_end_point": schema.StringAttribute{
				MarkdownDescription: "Salesforce: Authentication endpoint.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},

			// S3 Fields
			"aws_auth_type": schema.StringAttribute{
				MarkdownDescription: "S3: The authentication type for the S3 connection. It must be one of `iam_user` or `assume_role`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("iam_user", "assume_role"),
				},
			},
			"aws_iam_user": schema.SingleNestedAttribute{
				MarkdownDescription: "S3: IAM User configuration.",
				Optional:            true,
				Attributes: map[string]schema.Attribute{
					"access_key_id": schema.StringAttribute{
						MarkdownDescription: "S3: The access key ID for the S3 connection.",
						Optional:            true,
						Validators: []validator.String{
							stringvalidator.UTF8LengthAtLeast(1),
						},
					},
					"secret_access_key": schema.StringAttribute{
						MarkdownDescription: "S3: The secret access key for the S3 connection.",
						Optional:            true,
						Sensitive:           true,
						Validators: []validator.String{
							stringvalidator.UTF8LengthAtLeast(1),
						},
					},
				},
			},
			"aws_assume_role": schema.SingleNestedAttribute{
				MarkdownDescription: "S3: AssumeRole configuration.",
				Optional:            true,
				Attributes: map[string]schema.Attribute{
					"account_id": schema.StringAttribute{
						MarkdownDescription: "S3: The account ID for the AssumeRole configuration.",
						Optional:            true,
						Validators: []validator.String{
							stringvalidator.UTF8LengthAtLeast(1),
						},
					},
					"role_name": schema.StringAttribute{
						MarkdownDescription: "S3: The account role name for the AssumeRole configuration.",
						Optional:            true,
						Validators: []validator.String{
							stringvalidator.UTF8LengthAtLeast(1),
						},
					},
				},
			},

			// PostgreSQL Fields
			"driver": schema.StringAttribute{
				MarkdownDescription: `Snowflake, MySQL, PostgreSQL: The name of a Database driver.
  - MySQL: null, mysql_connector_java_5_1_49
  - Snowflake: null, snowflake_jdbc_3_14_2, snowflake_jdbc_3_17_0,
  - PostgreSQL: postgresql_42_5_1, postgresql_9_4_1205_jdbc41
`,
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf(
						// MySQL
						"mysql_connector_java_5_1_49",
						// Snowflake
						"snowflake_jdbc_3_14_2",
						"snowflake_jdbc_3_17_0",
						// PostgreSQL
						"postgresql_42_5_1",
						"postgresql_9_4_1205_jdbc41",
					),
				},
			},

			// Kintone Fields
			"domain": schema.StringAttribute{
				MarkdownDescription: "Kintone: Domain.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"login_method": schema.StringAttribute{
				MarkdownDescription: "Kintone: Login Method",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("username_and_password", "token"),
				},
			},
			"username": schema.StringAttribute{
				MarkdownDescription: "Kintone: The name of a user.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"token": schema.StringAttribute{
				MarkdownDescription: "Kintone: Token.",
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"basic_auth_username": schema.StringAttribute{
				MarkdownDescription: "Kintone: Basic Auth Username",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"basic_auth_password": schema.StringAttribute{
				MarkdownDescription: "Kintone: Basic Auth Password",
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			// Databricks Fields
			"server_hostname": schema.StringAttribute{
				MarkdownDescription: "Databricks: The host of a (Databricks) account.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"http_path": schema.StringAttribute{
				MarkdownDescription: "Databricks: The HTTP Path for the Databricks connection.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"auth_type": schema.StringAttribute{
				MarkdownDescription: "Databricks: The Auth Type for the Databricks connection. It must be one of `pat` or `oauth-m2m`.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.OneOf("pat", "oauth-m2m"),
				},
			},
			"personal_access_token": schema.StringAttribute{
				MarkdownDescription: "Databricks: The Personal Access Token for the Databricks connection.",
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"oauth2_client_id": schema.StringAttribute{
				MarkdownDescription: "Databricks: The OAuth2 Client ID for the Databricks connection.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"oauth2_client_secret": schema.StringAttribute{
				MarkdownDescription: "Databricks: The OAuth2 Client Secret for the Databricks connection.",
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			// SFTP Fields
			"secret_key": schema.StringAttribute{
				MarkdownDescription: "SFTP: RSA private key for authentication.",
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"secret_key_passphrase": schema.StringAttribute{
				MarkdownDescription: "SFTP: Passphrase for the RSA private key.",
				Optional:            true,
				Sensitive:           true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"user_directory_is_root": schema.BoolAttribute{
				MarkdownDescription: "SFTP: Whether the user directory is root. Default is true.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					planModifier.ConditionalBooleanDefault(true, "sftp"),
				},
			},
			"windows_server": schema.BoolAttribute{
				MarkdownDescription: "SFTP: Whether the server is a Windows server. Default is false.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					planModifier.ConditionalBooleanDefault(false, "sftp"),
				},
			},
			"ssh_tunnel_id": schema.Int64Attribute{
				MarkdownDescription: "SFTP: SSH tunnel ID. Required when aws_privatelink_enabled is true.",
				Optional:            true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"aws_privatelink_enabled": schema.BoolAttribute{
				MarkdownDescription: "SFTP: Whether AWS PrivateLink is enabled. Default is false.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					planModifier.ConditionalBooleanDefault(false, "sftp"),
				},
			},
		},
	}
}

func (r *connectionResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	plan := &connectionResourceModel{}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	conn, err := r.client.CreateConnection(
		plan.ConnectionType.ValueString(),
		plan.ToCreateConnectionInput(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating connection",
			fmt.Sprintf("Unable to create connection, got error: %s", err),
		)
		return
	}

	newState := connectionResourceModel{
		// Common Fields
		ConnectionType:  plan.ConnectionType,
		ID:              types.Int64Value(conn.ID),
		Name:            types.StringPointerValue(conn.Name),
		Description:     types.StringPointerValue(conn.Description),
		ResourceGroupID: types.Int64PointerValue(conn.ResourceGroupID),

		// BigQuery Fields
		ProjectID:             types.StringPointerValue(conn.ProjectID),
		ServiceAccountJSONKey: plan.ServiceAccountJSONKey,

		// Snowflake Fields
		Host:       types.StringPointerValue(conn.Host),
		UserName:   types.StringPointerValue(conn.UserName),
		Role:       types.StringPointerValue(conn.Role),
		AuthMethod: types.StringPointerValue(conn.AuthMethod),
		Password:   plan.Password,
		PrivateKey: plan.PrivateKey,

		// GCS Fields
		ApplicationName:     types.StringPointerValue(conn.ApplicationName),
		ServiceAccountEmail: plan.ServiceAccountEmail,

		// MySQL Fields
		Port: types.Int64PointerValue(conn.Port),

		// SSL Fields
		SSL: plan.SSL,

		// Gateway Fields
		Gateway: plan.Gateway,

		// Salesforce Fields
		SecurityToken: plan.SecurityToken,
		AuthEndPoint:  types.StringPointerValue(conn.AuthEndPoint),

		// S3 Fields
		AWSAuthType:   types.StringPointerValue(conn.AWSAuthType),
		AWSIAMUser:    plan.AWSIAMUser,
		AWSAssumeRole: connection.NewAWSAssumeRole(conn),

		// PostgreSQL Fields
		Driver: types.StringPointerValue(conn.Driver),

		// Kintone
		Domain:            types.StringPointerValue(conn.Domain),
		LoginMethod:       types.StringPointerValue(conn.LoginMethod),
		Token:             plan.Token,
		Username:          types.StringPointerValue(conn.Username),
		BasicAuthUsername: types.StringPointerValue(conn.BasicAuthUsername),
		BasicAuthPassword: plan.BasicAuthPassword,

		// SFTP Fields
		SecretKey:             plan.SecretKey,
		SecretKeyPassphrase:   plan.SecretKeyPassphrase,
		UserDirectoryIsRoot:   types.BoolPointerValue(conn.UserDirectoryIsRoot),
		WindowsServer:         types.BoolPointerValue(conn.WindowsServer),
		SSHTunnelID:           types.Int64PointerValue(conn.SSHTunnelID),
		AWSPrivatelinkEnabled: types.BoolPointerValue(conn.AWSPrivatelinkEnabled),
		// Databricks Fields
		ServerHostname:      types.StringPointerValue(conn.ServerHostname),
		HttpPath:            types.StringPointerValue(conn.HttpPath),
		AuthType:            types.StringPointerValue(conn.AuthType),
		PersonalAccessToken: plan.PersonalAccessToken,
		OAuth2ClientID:      plan.OAuth2ClientID,
		OAuth2ClientSecret:  plan.OAuth2ClientSecret,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (m *connectionResourceModel) NewGateway() *connection.Gateway {
	if m.Gateway == nil {
		return nil
	}
	return &connection.Gateway{
		Host:          m.Gateway.Host,
		Port:          m.Gateway.Port,
		UserName:      m.Gateway.UserName,
		Password:      m.Gateway.Password,
		Key:           m.Gateway.Key,
		KeyPassphrase: m.Gateway.KeyPassphrase,
	}
}

func (r *connectionResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	state := &connectionResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan := &connectionResourceModel{}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	connection, err := r.client.UpdateConnection(
		state.ConnectionType.ValueString(),
		state.ID.ValueInt64(),
		plan.ToUpdateConnectionInput(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating connection",
			fmt.Sprintf("Unable to update connection, got error: %s", err),
		)
		return
	}

	newState := connectionResourceModel{
		// Common Fields
		ConnectionType:  state.ConnectionType,
		ID:              types.Int64Value(connection.ID),
		Name:            types.StringPointerValue(connection.Name),
		Description:     types.StringPointerValue(connection.Description),
		ResourceGroupID: types.Int64PointerValue(connection.ResourceGroupID),

		// BigQuery Fields
		ProjectID:             types.StringPointerValue(connection.ProjectID),
		ServiceAccountJSONKey: plan.ServiceAccountJSONKey,

		// Snowflake Fields
		Host:       types.StringPointerValue(connection.Host),
		UserName:   types.StringPointerValue(connection.UserName),
		Role:       types.StringPointerValue(connection.Role),
		AuthMethod: types.StringPointerValue(connection.AuthMethod),
		Password:   plan.Password,
		PrivateKey: plan.PrivateKey,

		// GCS Fields
		ApplicationName:     types.StringPointerValue(connection.ApplicationName),
		ServiceAccountEmail: plan.ServiceAccountEmail,

		// MySQL Fields
		Port:    types.Int64PointerValue(connection.Port),
		SSL:     plan.SSL,
		Gateway: plan.Gateway,

		// Salesforce Fields
		SecurityToken: plan.SecurityToken,
		AuthEndPoint:  types.StringPointerValue(connection.AuthEndPoint),

		// S3 Fields
		AWSAuthType:   types.StringPointerValue(connection.AWSAuthType),
		AWSIAMUser:    plan.AWSIAMUser,
		AWSAssumeRole: plan.AWSAssumeRole,

		// PostgreSQL Fields
		Driver: types.StringPointerValue(connection.Driver),

		// Kintone
		Domain:            types.StringPointerValue(connection.Domain),
		LoginMethod:       types.StringPointerValue(connection.LoginMethod),
		Token:             plan.Token,
		Username:          types.StringPointerValue(connection.Username),
		BasicAuthUsername: types.StringPointerValue(connection.BasicAuthUsername),
		BasicAuthPassword: plan.BasicAuthPassword,

		// SFTP Fields
		SecretKey:             plan.SecretKey,
		SecretKeyPassphrase:   plan.SecretKeyPassphrase,
		UserDirectoryIsRoot:   types.BoolPointerValue(connection.UserDirectoryIsRoot),
		WindowsServer:         types.BoolPointerValue(connection.WindowsServer),
		SSHTunnelID:           types.Int64PointerValue(connection.SSHTunnelID),
		AWSPrivatelinkEnabled: types.BoolPointerValue(connection.AWSPrivatelinkEnabled),
		// Databricks Fields
		ServerHostname:      types.StringPointerValue(connection.ServerHostname),
		HttpPath:            types.StringPointerValue(connection.HttpPath),
		AuthType:            types.StringPointerValue(connection.AuthType),
		PersonalAccessToken: plan.PersonalAccessToken,
		OAuth2ClientID:      plan.OAuth2ClientID,
		OAuth2ClientSecret:  plan.OAuth2ClientSecret,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *connectionResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	state := &connectionResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	conn, err := r.client.GetConnection(
		state.ConnectionType.ValueString(),
		state.ID.ValueInt64(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading connection",
			fmt.Sprintf("Unable to read connection, got error: %s", err),
		)
		return
	}

	newState := connectionResourceModel{
		// Common Fields
		ConnectionType:  state.ConnectionType,
		ID:              types.Int64Value(conn.ID),
		Name:            types.StringPointerValue(conn.Name),
		Description:     types.StringPointerValue(conn.Description),
		ResourceGroupID: types.Int64PointerValue(conn.ResourceGroupID),

		// BigQuery Fields
		ProjectID:             types.StringPointerValue(conn.ProjectID),
		ServiceAccountJSONKey: state.ServiceAccountJSONKey,

		// Snowflake Fields
		Host:       types.StringPointerValue(conn.Host),
		UserName:   types.StringPointerValue(conn.UserName),
		Role:       types.StringPointerValue(conn.Role),
		AuthMethod: types.StringPointerValue(conn.AuthMethod),
		Password:   state.Password,
		PrivateKey: state.PrivateKey,

		// GCS Fields
		ApplicationName:     types.StringPointerValue(conn.ApplicationName),
		ServiceAccountEmail: state.ServiceAccountEmail,

		// MySQL Fields
		Port:    types.Int64PointerValue(conn.Port),
		SSL:     state.SSL,
		Gateway: state.Gateway,

		// Salesforce Fields
		SecurityToken: state.SecurityToken,
		AuthEndPoint:  types.StringPointerValue(conn.AuthEndPoint),

		// S3 Fields
		AWSAuthType:   types.StringPointerValue(conn.AWSAuthType),
		AWSIAMUser:    state.AWSIAMUser,
		AWSAssumeRole: connection.NewAWSAssumeRole(conn),

		// PostgreSQL Fields
		Driver: types.StringPointerValue(conn.Driver),

		// Kintone Fields
		Domain:            types.StringPointerValue(conn.Domain),
		LoginMethod:       types.StringPointerValue(conn.LoginMethod),
		Token:             state.Token,
		Username:          types.StringPointerValue(conn.Username),
		BasicAuthUsername: types.StringPointerValue(conn.BasicAuthUsername),
		BasicAuthPassword: state.BasicAuthPassword,

		// SFTP Fields
		SecretKey:             state.SecretKey,
		SecretKeyPassphrase:   state.SecretKeyPassphrase,
		UserDirectoryIsRoot:   types.BoolPointerValue(conn.UserDirectoryIsRoot),
		WindowsServer:         types.BoolPointerValue(conn.WindowsServer),
		SSHTunnelID:           types.Int64PointerValue(conn.SSHTunnelID),
		AWSPrivatelinkEnabled: types.BoolPointerValue(conn.AWSPrivatelinkEnabled),
		// Databricks Fields
		ServerHostname:      types.StringPointerValue(conn.ServerHostname),
		HttpPath:            types.StringPointerValue(conn.HttpPath),
		AuthType:            types.StringPointerValue(conn.AuthType),
		PersonalAccessToken: state.PersonalAccessToken,
		OAuth2ClientID:      types.StringPointerValue(conn.OAuth2ClientID),
		OAuth2ClientSecret:  state.OAuth2ClientSecret,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *connectionResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	s := &connectionResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, s)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.DeleteConnection(
		s.ConnectionType.ValueString(),
		s.ID.ValueInt64(),
	); err != nil {
		resp.Diagnostics.AddError(
			"Deleting connection",
			fmt.Sprintf("Unable to delete connection, got error: %s", err),
		)
		return
	}
}

func (r *connectionResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	// We must write custom logic because the Read method requires two attributes to refresh.
	// For more information, see https://developer.hashicorp.com/terraform/plugin/framework/resources/import#multiple-attributes.

	idParts := strings.Split(req.ID, ",")
	if len(idParts) != 2 || idParts[0] == "" || idParts[1] == "" {
		resp.Diagnostics.AddError(
			"Importing connection",
			fmt.Sprintf("Expected import identifier with format: connection_type,connection_id. Got: %q", req.ID),
		)
		return
	}

	connectionType := idParts[0]

	connectionID, err := strconv.ParseInt(idParts[1], 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing connection",
			fmt.Sprintf("Failed to parse connection ID: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("connection_type"), connectionType)...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), connectionID)...)
}

func (r *connectionResource) ValidateConfig(
	ctx context.Context,
	req resource.ValidateConfigRequest,
	resp *resource.ValidateConfigResponse,
) {
	plan := &connectionResourceModel{}
	resp.Diagnostics.Append(req.Config.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	switch plan.ConnectionType.ValueString() {
	case "bigquery":
		validateRequiredString(plan.ServiceAccountJSONKey, "service_account_json_key", "BigQuery", resp)
		validateRequiredString(plan.ProjectID, "project_id", "BigQuery", resp)
	case "snowflake":
		validateRequiredString(plan.Host, "host", "Snowflake", resp)
		validateRequiredString(plan.UserName, "user_name", "Snowflake", resp)
		validateRequiredString(plan.AuthMethod, "auth_method", "Snowflake", resp)
		if plan.AuthMethod.ValueString() == "key_pair" {
			validateRequiredString(plan.PrivateKey, "private_key", "Snowflake", resp)
		}
		if plan.AuthMethod.ValueString() == "user_password" {
			validateRequiredString(plan.Password, "password", "Snowflake", resp)
		}
		validateStringAgainstPatterns(plan.Driver, "driver", "Snowflake", resp, "snowflake_jdbc_3_14_2", "snowflake_jdbc_3_17_0")
	case "gcs":
		validateRequiredString(plan.ApplicationName, "application_name", "GCS", resp)
		validateRequiredString(plan.ServiceAccountEmail, "service_account_email", "GCS", resp)
		validateRequiredString(plan.ProjectID, "project_id", "GCS", resp)
	case "google_spreadsheets":
		validateRequiredString(plan.ServiceAccountJSONKey, "service_account_json_key", "Google Sheets", resp)
	case "mysql":
		validateRequiredString(plan.Host, "host", "MySQL", resp)
		validateRequiredInt(plan.Port, "port", "MySQL", resp)
		validateRequiredString(plan.UserName, "user_name", "MySQL", resp)
		validateRequiredString(plan.Password, "password", "MySQL", resp)
		validateStringAgainstPatterns(plan.Driver, "driver", "MySQL", resp, "mysql_connector_java_5_1_49")
		if plan.Gateway != nil {
			validateRequiredString(plan.Gateway.Host, "gateway.host", "MySQL", resp)
			validateRequiredInt(plan.Gateway.Port, "gateway.port", "MySQL", resp)
			validateRequiredString(plan.Gateway.UserName, "gateway.user_name", "MySQL", resp)
		}
	case "salesforce":
		validateRequiredString(plan.AuthMethod, "auth_method", "Salesforce", resp)
		if plan.AuthMethod.ValueString() != "user_password" {
			resp.Diagnostics.AddError(
				"auth_method",
				"auth_method must be 'user_password' for Salesforce connection.",
			)
		}
		validateRequiredString(plan.UserName, "user_name", "Salesforce", resp)
		validateRequiredString(plan.Password, "password", "Salesforce", resp)
		validateRequiredString(plan.SecurityToken, "security_token", "Salesforce", resp)
		validateRequiredString(plan.AuthEndPoint, "auth_end_point", "Salesforce", resp)
	case "s3":
		validateRequiredString(plan.AWSAuthType, "aws_auth_type", "S3", resp)
		if plan.AWSAssumeRole != nil && plan.AWSIAMUser != nil {
			resp.Diagnostics.AddError(
				"aws_auth_type",
				"`aws_assume_role` and `aws_iam_user` cannot be used together for S3 connection.",
			)
		}
		switch plan.AWSAuthType.ValueString() {
		case "iam_user":
			if plan.AWSIAMUser == nil {
				resp.Diagnostics.AddError(
					"aws_iam_user",
					"aws_iam_user is required for S3 connection with aws_auth_type `iam_user`.",
				)
			} else {
				validateRequiredString(plan.AWSIAMUser.AccessKeyID, "aws_iam_user.access_key_id", "S3", resp)
				validateRequiredString(plan.AWSIAMUser.SecretAccessKey, "aws_iam_user.secret_access_key", "S3", resp)
			}
		case "assume_role":
			if plan.AWSAssumeRole == nil {
				resp.Diagnostics.AddError(
					"aws_assume_role",
					"aws_assume_role is required for S3 connection with aws_auth_type `assume_role`.",
				)
			} else {
				validateRequiredString(plan.AWSAssumeRole.AccountID, "aws_assume_role.account_id", "S3", resp)
				validateRequiredString(plan.AWSAssumeRole.AccountRoleName, "aws_assume_role.account_role_name", "S3", resp)
			}
		}
	case "postgresql":
		validateRequiredString(plan.Host, "host", "PostgreSQL", resp)
		validateRequiredInt(plan.Port, "port", "PostgreSQL", resp)
		validateRequiredString(plan.UserName, "user_name", "PostgreSQL", resp)
		validateRequiredString(plan.Driver, "driver", "PostgreSQL", resp)
		validateStringAgainstPatterns(plan.Driver, "driver", "PostgreSQL", resp, "postgresql_42_5_1", "postgresql_9_4_1205_jdbc41")
		if plan.Gateway != nil {
			validateRequiredString(plan.Gateway.Host, "gateway.host", "PostgreSQL", resp)
			validateRequiredInt(plan.Gateway.Port, "gateway.port", "PostgreSQL", resp)
			validateRequiredString(plan.Gateway.UserName, "gateway.user_name", "PostgreSQL", resp)
		}
	case "google_analytics4":
		validateRequiredString(plan.ServiceAccountJSONKey, "service_account_json_key", "Google Analytics4", resp)
	case "kintone":
		validateRequiredString(plan.Domain, "domain", "Kintone", resp)
		validateRequiredString(plan.LoginMethod, "login_method", "Kintone", resp)
		validateStringAgainstPatterns(plan.LoginMethod, "login_method", "Kintone", resp, "username_and_password", "token")
		switch plan.LoginMethod.ValueString() {
		case "username_and_password":
			if plan.Password.IsNull() {
				resp.Diagnostics.AddError(
					"password",
					"password is required for Kintone connection with login_method `username_and_password`.",
				)
			}
			if plan.Username.IsNull() {
				resp.Diagnostics.AddError(
					"username",
					"username is required for Kintone connection with login_method `username_and_password`.",
				)
			}
			if !plan.Token.IsNull() {
				resp.Diagnostics.AddError(
					"token",
					"token should not be set when login_method is `username_and_password`.",
				)
			}
		case "token":
			if plan.Token.IsNull() {
				resp.Diagnostics.AddError(
					"token",
					"token is required for Kintone connection with login_method `token`.",
				)
			}
			if !plan.Password.IsNull() {
				resp.Diagnostics.AddError(
					"password",
					"password should not be set when login_method is `token`.",
				)
			}
		}
	case "sftp":
		validateRequiredString(plan.Host, "host", "SFTP", resp)
		validateRequiredInt(plan.Port, "port", "SFTP", resp)
		validateRequiredString(plan.UserName, "user_name", "SFTP", resp)
		if plan.AWSPrivatelinkEnabled.ValueBool() {
			validateRequiredInt(plan.SSHTunnelID, "ssh_tunnel_id", "SFTP", resp)
		}
	case "databricks":
		validateRequiredString(plan.ServerHostname, "server_hostname", "Databricks", resp)
		validateRequiredString(plan.HttpPath, "http_path", "Databricks", resp)
		validateRequiredString(plan.AuthType, "auth_type", "Databricks", resp)
		validateStringAgainstPatterns(plan.AuthType, "auth_type", "Databricks", resp, "pat", "oauth-m2m")
		switch plan.AuthType.ValueString() {
		case "pat":
			if plan.PersonalAccessToken.IsNull() {
				resp.Diagnostics.AddError(
					"personal_access_token",
					"personal_access_token is required for Databricks connection with auth_type `pat`.",
				)
			}
		case "oauth-m2m":
			if plan.OAuth2ClientID.IsNull() {
				resp.Diagnostics.AddError(
					"oauth2_client_id",
					"oauth2_client_id is required for Databricks connection with auth_type `oauth-m2m`.",
				)
			}
			if plan.OAuth2ClientSecret.IsNull() {
				resp.Diagnostics.AddError(
					"oauth2_client_secret",
					"oauth2_client_secret is required for Databricks connection with auth_type `oauth-m2m`.",
				)
			}
		}
	}
}

func validateStringAgainstPatterns(field types.String, fieldName, connectionType string, resp *resource.ValidateConfigResponse, patterns ...string) {
	if field.IsNull() {
		return
	}

	for _, pattern := range patterns {
		if field.ValueString() == pattern {
			return
		}
	}

	resp.Diagnostics.AddError(
		fieldName,
		fmt.Sprintf("%s: `%s` is invalid for %s connection. Valid values are: %s",
			fieldName,
			field.ValueString(),
			connectionType,
			strings.Join(patterns, ", "),
		),
	)
}

func validateRequiredString(field types.String, fieldName, connectionType string, resp *resource.ValidateConfigResponse) {
	if field.IsNull() {
		resp.Diagnostics.AddError(
			fieldName,
			fmt.Sprintf("%s is required for %s connection.", fieldName, connectionType),
		)
	}
}

func validateRequiredInt(field types.Int64, fieldName, connectionType string, resp *resource.ValidateConfigResponse) {
	if field.IsNull() {
		resp.Diagnostics.AddError(
			fieldName,
			fmt.Sprintf("%s is required for %s connection.", fieldName, connectionType),
		)
	}
}
