package output_options

import (
	"terraform-provider-trocco/internal/client/entity/job_definition/output_option"
	outputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DatabricksOutputOption struct {
	DatabricksConnectionID types.Int64  `tfsdk:"databricks_connection_id"`
	CatalogName            types.String `tfsdk:"catalog_name"`
	SchemaName             types.String `tfsdk:"schema_name"`
	Table                  types.String `tfsdk:"table"`
	BatchSize              types.Int64  `tfsdk:"batch_size"`
	Mode                   types.String `tfsdk:"mode"`
	DefaultTimeZone        types.String `tfsdk:"default_time_zone"`
}

func NewDatabricksOutputOption(entity *output_option.DatabricksOutputOption) *DatabricksOutputOption {
	if entity == nil {
		return nil
	}

	return &DatabricksOutputOption{
		DatabricksConnectionID: types.Int64Value(entity.DatabricksConnectionID),
		CatalogName:            types.StringValue(entity.CatalogName),
		SchemaName:             types.StringValue(entity.SchemaName),
		Table:                  types.StringValue(entity.Table),
		BatchSize:              types.Int64Value(entity.BatchSize),
		Mode:                   types.StringValue(entity.Mode),
		DefaultTimeZone:        types.StringValue(entity.DefaultTimeZone),
	}
}

func (o *DatabricksOutputOption) ToInput() *outputOptionParameters.DatabricksOutputOptionInput {
	if o == nil {
		return nil
	}

	return &outputOptionParameters.DatabricksOutputOptionInput{
		DatabricksConnectionID: o.DatabricksConnectionID.ValueInt64(),
		CatalogName:            o.CatalogName.ValueString(),
		SchemaName:             o.SchemaName.ValueString(),
		Table:                  o.Table.ValueString(),
		BatchSize:              o.BatchSize.ValueInt64(),
		Mode:                   o.Mode.ValueString(),
		DefaultTimeZone:        o.DefaultTimeZone.ValueString(),
	}
}

func (o *DatabricksOutputOption) ToUpdateInput() *outputOptionParameters.UpdateDatabricksOutputOptionInput {
	if o == nil {
		return nil
	}

	return &outputOptionParameters.UpdateDatabricksOutputOptionInput{
		DatabricksConnectionID: o.DatabricksConnectionID.ValueInt64Pointer(),
		CatalogName:            o.CatalogName.ValueStringPointer(),
		SchemaName:             o.SchemaName.ValueStringPointer(),
		Table:                  o.Table.ValueStringPointer(),
		BatchSize:              o.BatchSize.ValueInt64Pointer(),
		Mode:                   o.Mode.ValueStringPointer(),
		DefaultTimeZone:        o.DefaultTimeZone.ValueStringPointer(),
	}
}
