package pipeline_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/samber/lo"

	we "terraform-provider-trocco/internal/client/entity/pipeline_definition"
	p "terraform-provider-trocco/internal/client/parameter"
	wp "terraform-provider-trocco/internal/client/parameter/pipeline_definition"
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

func NewCustomVariableLoop(ctx context.Context, en *we.CustomVariableLoop) *CustomVariableLoop {
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
		md.StringConfig = NewStringCustomVariableLoopConfig(ctx, en.StringConfig)
	}
	if en.PeriodConfig != nil {
		md.PeriodConfig = NewPeriodCustomVariableLoopConfig(ctx, en.PeriodConfig)
	}
	if en.BigqueryConfig != nil {
		md.BigqueryConfig = NewBigqueryCustomVariableLoopConfig(ctx, en.BigqueryConfig)
	}
	if en.SnowflakeConfig != nil {
		md.SnowflakeConfig = NewSnowflakeCustomVariableLoopConfig(ctx, en.SnowflakeConfig)
	}
	if en.RedshiftConfig != nil {
		md.RedshiftConfig = NewRedshiftCustomVariableLoopConfig(ctx, en.RedshiftConfig)
	}

	return md
}

func (c *CustomVariableLoop) ToInput(ctx context.Context) wp.CustomVariableLoop {
	i := wp.CustomVariableLoop{
		Type:                       c.Type.ValueString(),
		IsParallelExecutionAllowed: &p.NullableBool{Valid: !c.IsParallelExecutionAllowed.IsNull(), Value: c.IsParallelExecutionAllowed.ValueBool()},
		IsStoppedOnErrors:          &p.NullableBool{Valid: !c.IsStoppedOnErrors.IsNull(), Value: c.IsStoppedOnErrors.ValueBool()},
		MaxErrors:                  &p.NullableInt64{Valid: !c.MaxErrors.IsNull(), Value: c.MaxErrors.ValueInt64()},
	}

	if c.StringConfig != nil {
		i.StringConfig = lo.ToPtr(c.StringConfig.ToInput(ctx))
	}
	if c.PeriodConfig != nil {
		i.PeriodConfig = lo.ToPtr(c.PeriodConfig.ToInput(ctx))
	}
	if c.BigqueryConfig != nil {
		i.BigqueryConfig = lo.ToPtr(c.BigqueryConfig.ToInput(ctx))
	}
	if c.SnowflakeConfig != nil {
		i.SnowflakeConfig = lo.ToPtr(c.SnowflakeConfig.ToInput(ctx))
	}
	if c.RedshiftConfig != nil {
		i.RedshiftConfig = lo.ToPtr(c.RedshiftConfig.ToInput(ctx))
	}

	return i
}

func CustomVariableLoopAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"type":                          types.StringType,
		"is_parallel_execution_allowed": types.BoolType,
		"is_stopped_on_errors":          types.BoolType,
		"max_errors":                    types.Int64Type,
		"string_config": types.ObjectType{
			AttrTypes: StringCustomVariableLoopConfigAttrTypes(),
		},
		"period_config": types.ObjectType{
			AttrTypes: PeriodCustomVariableLoopConfigAttrTypes(),
		},
		"bigquery_config": types.ObjectType{
			AttrTypes: BigqueryCustomVariableLoopConfigAttrTypes(),
		},
		"snowflake_config": types.ObjectType{
			AttrTypes: SnowflakeCustomVariableLoopConfigAttrTypes(),
		},
		"redshift_config": types.ObjectType{
			AttrTypes: RedshiftCustomVariableLoopConfigAttrTypes(),
		},
	}
}
