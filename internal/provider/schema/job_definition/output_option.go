package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func OutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Required: true,
		Attributes: map[string]schema.Attribute{
			"bigquery_output_option": BigqueryOutputOptionSchema(),
		},
	}
}
