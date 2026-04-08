package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func YahooAdsApiYdnInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of source yahoo_ads_api_ydn",
		Attributes: map[string]schema.Attribute{
			"yahoo_ads_api_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of yahoo_ads_api connection",
			},
			"target": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Data retrieval target. Either \"report\" or \"stats\"",
				Validators: []validator.String{
					stringvalidator.OneOf("report", "stats"),
				},
			},
			"account_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Yahoo Display Ads account ID",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"base_account_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Base account ID (required for POST)",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"report_type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Report type. Valid only when target=\"stats\". One of: CAMPAIGN, ADGROUP, AD",
				Validators: []validator.String{
					stringvalidator.OneOf("CAMPAIGN", "ADGROUP", "AD"),
				},
			},
			"start_date": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Start date. Format: YYYYMMDD or custom variable (e.g., $start_date$)",
			},
			"end_date": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "End date. Format: YYYYMMDD or custom variable (e.g., $end_date$)",
			},
			"include_deleted": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Include deleted ads. Valid only when target=\"report\"",
			},
			"input_option_columns": schema.ListNestedAttribute{
				Optional:            true,
				MarkdownDescription: "List of columns to be retrieved and their types",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Column name",
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
						"type": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Column type",
							Validators: []validator.String{
								stringvalidator.OneOf("boolean", "long", "timestamp", "double", "string", "json"),
							},
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Column format (for timestamp type)",
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
	}
}
