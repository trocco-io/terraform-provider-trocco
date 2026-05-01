package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func MarketoInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes about source Marketo Engage",
		Attributes: map[string]schema.Attribute{
			"marketo_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "Id of Marketo connection",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"target": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Target data type",
				Validators: []validator.String{
					stringvalidator.OneOf("lead", "activity", "campaign", "all_lead_with_list_id", "program", "program_members", "custom_object", "list", "activity_type", "folder"),
				},
			},
			// Date/Time Parameters (lead/activity only)
			"from_date": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Start date in ISO 8601 format (lead/activity only)",
			},
			"end_date": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "End date in ISO 8601 format (lead/activity only)",
			},
			"use_updated_at": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Use updated_at for filtering (lead only)",
			},
			"polling_interval_second": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Polling interval in seconds (lead/activity only)",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"bulk_job_timeout_second": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Bulk job timeout in seconds (lead/activity only)",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			// Activity Parameters
			"activity_type_ids": schema.ListAttribute{
				ElementType:         types.Int64Type,
				Optional:            true,
				MarkdownDescription: "Activity type IDs (activity only)",
			},
			// Custom Object Parameters
			"custom_object_api_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Custom object API name (custom_object only)",
			},
			"custom_object_filter_type": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Custom object filter type (custom_object only)",
			},
			"custom_object_filter_from_value": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Custom object filter range start (custom_object only)",
			},
			"custom_object_filter_to_value": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Custom object filter range end (custom_object only)",
			},
			"custom_object_fields": schema.ListNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Custom object fields (custom_object only)",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Field name",
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
					},
				},
			},
			// List/Program Parameters
			"list_ids": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "List IDs comma-separated (all_lead_with_list_id only)",
			},
			"program_ids": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Program IDs comma-separated (program_members only)",
			},
			// Folder Parameters
			"root_id": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "Root folder/program ID (folder only)",
			},
			"root_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Root type (folder only)",
				Validators: []validator.String{
					stringvalidator.OneOf("folder", "program"),
				},
			},
			"max_depth": schema.Int64Attribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Max folder depth (folder only)",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"workspace": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Workspace name (folder only)",
			},
			// Column Definitions
			"input_option_columns": schema.ListNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Column definitions",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column name",
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
						"type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column data type",
							Validators: []validator.String{
								stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean", "json"),
							},
						},
					},
				},
			},
			// Filter Columns (lead/all_lead_with_list_id only)
			"marketo_input_option_filter_columns": schema.ListNestedAttribute{
				Optional:            true,
				MarkdownDescription: "Filter columns (lead/all_lead_with_list_id only)",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Filter column name",
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
					},
				},
			},
			// Custom Variables
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
	}
}
