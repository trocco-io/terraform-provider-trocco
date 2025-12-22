package job_definition

import (
	planModifier "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func DatabricksOutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of destination Databricks settings",
		Attributes: map[string]schema.Attribute{
			"databricks_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of Databricks connection",
			},
			"catalog_name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Databricks catalog name",
			},
			"schema_name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Databricks schema name",
			},
			"table": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Table name",
			},
			"batch_size": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Batch size for data transfer.",
			},
			"mode": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf("insert", "insert_direct", "truncate_insert", "replace", "merge"),
				},
				MarkdownDescription: "Write mode. One of `insert`, `insert_direct`, `truncate_insert`, `replace`, `merge`",
			},
			"default_time_zone": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Default time zone for timestamp without time zone",
			},
			"databricks_output_option_column_options": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column name",
						},
						"type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Data type",
						},
						"value_type": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Value type",
						},
						"timestamp_format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Timestamp format",
						},
						"timezone": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Time zone",
						},
					},
					PlanModifiers: []planmodifier.Object{
						&planModifier.DatabricksOutputOptionColumnPlanModifier{},
					},
				},
			},
			"databricks_output_option_merge_keys": schema.SetAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Merge keys (only applicable if mode is 'merge')",
			},
		},
		PlanModifiers: []planmodifier.Object{
			&planModifier.DatabricksOutputOptionPlanModifier{},
		},
	}
}
