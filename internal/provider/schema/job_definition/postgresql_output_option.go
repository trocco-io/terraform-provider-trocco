package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func PostgresqlOutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of destination PostgreSQL settings",
		Attributes: map[string]schema.Attribute{
			"database": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Database name",
			},
			"schema": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Schema name",
			},
			"table": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Table name",
			},
			"mode": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf("insert", "insert_direct", "truncate_insert", "replace", "merge"),
				},
				MarkdownDescription: "Transfer mode",
			},
			"default_time_zone": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Default time zone",
			},
			"postgresql_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "PostgreSQL connection ID",
			},
			"merge_keys": schema.SetAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Merge keys (only applicable if mode is 'merge')",
			},
			"retry_limit": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(12),
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Maximum number of retries. Default is 12.",
			},
			"retry_wait": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(1000),
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Initial wait time in milliseconds between retries. Default is 1000.",
			},
			"max_retry_wait": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(1800000),
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Maximum wait time in milliseconds between retries. Default is 1800000.",
			},
			"before_load": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "SQL statement to execute before loading data (not available when mode is 'replace').",
			},
			"after_load": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "SQL statement to execute after loading data.",
			},
		},
	}
}
