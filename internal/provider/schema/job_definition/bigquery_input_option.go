package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func BigqueryInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes about source bigquery",
		Attributes: map[string]schema.Attribute{
			"bigquery_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "Id of bigquery connection",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"gcs_uri": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "GCS URI",
			},
			"gcs_uri_format": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("bucket"),
				Validators: []validator.String{
					stringvalidator.OneOf("bucket", "custom_path"),
				},
				MarkdownDescription: "Format of GCS URI",
			},
			"query": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Query",
			},
			"temp_dataset": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Temporary dataset name",
			},
			"is_standard_sql": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "Flag whether standard SQL is enabled",
			},
			"cleanup_gcs_files": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "Flag whether temporary GCS files should be cleaned up",
			},
			"file_format": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("CSV"),
				Validators: []validator.String{
					stringvalidator.OneOf("CSV", "NEWLINE_DELIMITED_JSON"),
				},
				MarkdownDescription: "File format of temporary GCS files",
			},
			"location": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("US"),
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Location of bigquery job",
			},
			"cache": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "Flag whether query cache is enabled",
			},
			"bigquery_job_wait_second": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				Default:             int64default.StaticInt64(600),
				MarkdownDescription: "Wait time in seconds until bigquery job is completed",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"columns": schema.ListNestedAttribute{
				Required:            true,
				MarkdownDescription: "List of columns to be retrieved and their types",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
							MarkdownDescription: "Column name",
						},
						"type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column type.",
							Validators: []validator.String{
								stringvalidator.OneOf("boolean", "long", "timestamp", "double", "string", "json"),
							},
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "format",
						},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"decoder":                  DecoderSchema(),
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
	}
}
