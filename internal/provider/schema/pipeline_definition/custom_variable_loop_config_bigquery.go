package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func BigqueryCustomVariableLoopConfig() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "BigQuery custom variabe loop configuration",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"connection_id": schema.Int64Attribute{
				MarkdownDescription: "BigQuery connection ID",
				Required:            true,
			},
			"query": schema.StringAttribute{
				MarkdownDescription: "Query to expand custom variables",
				Required:            true,
			},
			"variables": schema.SetAttribute{
				MarkdownDescription: "Custom variables to be expanded",
				Required:            true,
				ElementType:         types.StringType,
				Validators: []validator.Set{
					setvalidator.SizeAtLeast(1),
				},
			},
		},
	}
}
