package workflow

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func NewTroccoTransferTaskConfigAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"definition_id": schema.Int64Attribute{
				Required: true,
			},
			"custom_variable_loop": NewCustomVariableLoopAttribute(),
		},
	}
}

func NewTroccoTransferBulkTaskConfigAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"definition_id": schema.Int64Attribute{
				Required: true,
			},
		},
	}
}

func NewDBTTaskConfigAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"definition_id": schema.Int64Attribute{
				Required: true,
			},
		},
	}
}

func NewTroccoAgentTaskConfigAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"definition_id": schema.Int64Attribute{
				Required: true,
			},
		},
	}
}

func NewBigQueryDatamartTaskConfigAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"definition_id": schema.Int64Attribute{
				Required: true,
			},
			"custom_variable_loop": NewCustomVariableLoopAttribute(),
		},
	}
}

func NewRedshiftDatamartTaskConfigAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"definition_id": schema.Int64Attribute{
				Required: true,
			},
			"custom_variable_loop": NewCustomVariableLoopAttribute(),
		},
	}
}

func NewSnowflakeDatamartTaskConfigAttribute() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"definition_id": schema.Int64Attribute{
				Required: true,
			},
			"custom_variable_loop": NewCustomVariableLoopAttribute(),
		},
	}
}
