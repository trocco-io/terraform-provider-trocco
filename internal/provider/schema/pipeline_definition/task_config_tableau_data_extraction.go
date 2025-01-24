package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func TableauDataExtractionTaskConfig() schema.Attribute {
	return schema.SingleNestedAttribute{
		MarkdownDescription: "The task configuration for the tableau data extraction task.",
		Optional:            true,
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the task",
				Required:            true,
			},
			"connection_id": schema.Int64Attribute{
				MarkdownDescription: "The connection id to use for the task",
				Required:            true,
			},
			"task_id": schema.StringAttribute{
				MarkdownDescription: "The Tableau task ID. You can get with the [Tableau API](https://help.tableau.com/current/api/rest_api/en-us/REST/rest_api_ref.htm#list_extract_refresh_tasks1).",
				Required:            true,
			},
		},
	}
}
