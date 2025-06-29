package pipeline_definition

import (
	"terraform-provider-trocco/internal/provider/custom_type"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func RedshiftDatacheckTaskConfigSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "The task configuration for the datacheck task.",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the datacheck task",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"connection_id": schema.Int64Attribute{
				MarkdownDescription: "The connection id to use for the datacheck task",
				Required:            true,
			},
			"query": schema.StringAttribute{
				MarkdownDescription: "The query to run for the datacheck task",
				Optional:            true,
				CustomType:          custom_type.TrimmedStringType{},
			},
			"operator": schema.StringAttribute{
				MarkdownDescription: "The operator to use for the datacheck task",
				Required:            true,
			},
			"query_result": schema.Int64Attribute{
				MarkdownDescription: "The query result to use for the datacheck task",
				Required:            true,
			},
			"accepts_null": schema.BoolAttribute{
				MarkdownDescription: "Whether the datacheck task accepts null values",
				Required:            true,
			},
			"database": schema.StringAttribute{
				MarkdownDescription: "The database to use for the datacheck task",
				Optional:            true,
			},
			"custom_variables": CustomVariablesSchema(),
		},
	}
}
