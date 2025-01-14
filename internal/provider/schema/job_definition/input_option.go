package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func InputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Required: true,
		Attributes: map[string]schema.Attribute{
			"mysql_input_option": MysqlInputOptionSchema(),
			"gcs_input_option":   GcsInputOptionSchema(),
		},
	}
}
