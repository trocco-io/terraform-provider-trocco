package pipeline_definition

import (
	"terraform-provider-trocco/internal/provider/custom_type"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func BigqueryDatacheckTaskConfig() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "The datacheck task config of the pipeline definition",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the datacheck task",
				Required:            true,
			},
			"connection_id": schema.Int64Attribute{
				MarkdownDescription: "The connection id of the datacheck task",
				Required:            true,
			},
			"query": schema.StringAttribute{
				MarkdownDescription: "The query of the datacheck task",
				Optional:            true,
				CustomType:          custom_type.TrimmedStringType{},
			},
			"operator": schema.StringAttribute{
				MarkdownDescription: "The operator of the datacheck task",
				Required:            true,
			},
			"query_result": schema.Int64Attribute{
				MarkdownDescription: "The query result of the datacheck task",
				Required:            true,
			},
			"accepts_null": schema.BoolAttribute{
				MarkdownDescription: "Whether the datacheck task accepts null",
				Required:            true,
			},
			"custom_variables": CustomVariables(),
		},
	}
}
