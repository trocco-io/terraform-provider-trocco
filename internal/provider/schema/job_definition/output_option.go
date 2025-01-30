package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	planmodifier2 "terraform-provider-trocco/internal/provider/planmodifier"
)

func OutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Required: true,
		Attributes: map[string]schema.Attribute{
			"bigquery_output_option": BigqueryOutputOptionSchema(),
			// "snowflake_output_option": SnowflakeOutputOptionSchema(),
		},
		PlanModifiers: []planmodifier.Object{
			&planmodifier2.OutputOptionPlanModifier{},
		},
	}
}
