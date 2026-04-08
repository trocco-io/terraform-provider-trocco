package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func RedshiftInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of source redshift",
		Attributes: map[string]schema.Attribute{
			"redshift_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "ID of Redshift connection",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"database": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Database name",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"schema": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("public"),
				MarkdownDescription: "Schema name",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"query": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "SQL query to fetch data",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"fetch_rows": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(10000),
				MarkdownDescription: "Number of records processed by the cursor at one time",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"connect_timeout": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(300),
				MarkdownDescription: "Connection timeout (sec)",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"socket_timeout": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(1800),
				MarkdownDescription: "Socket timeout (sec)",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
	}
}
