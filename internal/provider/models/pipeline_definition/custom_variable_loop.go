package pipeline_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"

	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"
	p "terraform-provider-trocco/internal/client/parameters"
	wp "terraform-provider-trocco/internal/client/parameters/pipeline_definition"
)

type CustomVariableLoop struct {
	Type                       types.String `tfsdk:"type"`
	IsParallelExecutionAllowed types.Bool   `tfsdk:"is_parallel_execution_allowed"`
	IsStoppedOnErrors          types.Bool   `tfsdk:"is_stopped_on_errors"`
	MaxErrors                  types.Int64  `tfsdk:"max_errors"`

	StringConfig    *StringCustomVariableLoopConfig    `tfsdk:"string_config"`
	PeriodConfig    *PeriodCustomVariableLoopConfig    `tfsdk:"period_config"`
	BigqueryConfig  *BigqueryCustomVariableLoopConfig  `tfsdk:"bigquery_config"`
	SnowflakeConfig *SnowflakeCustomVariableLoopConfig `tfsdk:"snowflake_config"`
	RedshiftConfig  *RedshiftCustomVariableLoopConfig  `tfsdk:"redshift_config"`
}

func NewCustomVariableLoop(en *we.CustomVariableLoop) *CustomVariableLoop {
	if en == nil {
		return nil
	}

	md := &CustomVariableLoop{
		Type:                       types.StringValue(en.Type),
		IsParallelExecutionAllowed: types.BoolPointerValue(en.IsParallelExecutionAllowed),
		IsStoppedOnErrors:          types.BoolPointerValue(en.IsStoppedOnErrors),
		MaxErrors:                  types.Int64PointerValue(en.MaxErrors),
	}

	if en.StringConfig != nil {
		md.StringConfig = NewStringCustomVariableLoopConfig(en.StringConfig)
	}
	if en.PeriodConfig != nil {
		md.PeriodConfig = NewPeriodCustomVariableLoopConfig(en.PeriodConfig)
	}
	if en.BigqueryConfig != nil {
		md.BigqueryConfig = NewBigqueryCustomVariableLoopConfig(en.BigqueryConfig)
	}
	if en.SnowflakeConfig != nil {
		md.SnowflakeConfig = NewSnowflakeCustomVariableLoopConfig(en.SnowflakeConfig)
	}
	if en.RedshiftConfig != nil {
		md.RedshiftConfig = NewRedshiftCustomVariableLoopConfig(en.RedshiftConfig)
	}

	return md
}

func (c *CustomVariableLoop) ToInput() wp.CustomVariableLoop {
	i := wp.CustomVariableLoop{
		Type:                       c.Type.ValueString(),
		IsParallelExecutionAllowed: &p.NullableBool{Valid: !c.IsParallelExecutionAllowed.IsNull(), Value: c.IsParallelExecutionAllowed.ValueBool()},
		IsStoppedOnErrors:          &p.NullableBool{Valid: !c.IsStoppedOnErrors.IsNull(), Value: c.IsStoppedOnErrors.ValueBool()},
		MaxErrors:                  &p.NullableInt64{Valid: !c.MaxErrors.IsNull(), Value: c.MaxErrors.ValueInt64()},
	}

	if c.StringConfig != nil {
		i.StringConfig = lo.ToPtr(c.StringConfig.ToInput())
	}
	if c.PeriodConfig != nil {
		i.PeriodConfig = lo.ToPtr(c.PeriodConfig.ToInput())
	}
	if c.BigqueryConfig != nil {
		i.BigqueryConfig = lo.ToPtr(c.BigqueryConfig.ToInput())
	}
	if c.SnowflakeConfig != nil {
		i.SnowflakeConfig = lo.ToPtr(c.SnowflakeConfig.ToInput())
	}
	if c.RedshiftConfig != nil {
		i.RedshiftConfig = lo.ToPtr(c.RedshiftConfig.ToInput())
	}

	return i
}