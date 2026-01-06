package output_options

import (
	"context"
	"fmt"
	"terraform-provider-trocco/internal/client/entity/job_definition/output_option"
	outputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/output_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type PostgresqlOutputOption struct {
	Database                        types.String `tfsdk:"database"`
	Schema                          types.String `tfsdk:"schema"`
	Table                           types.String `tfsdk:"table"`
	Mode                            types.String `tfsdk:"mode"`
	DefaultTimeZone                 types.String `tfsdk:"default_time_zone"`
	PostgresqlConnectionId          types.Int64  `tfsdk:"postgresql_connection_id"`
	PostgresqlOutputOptionMergeKeys types.Set    `tfsdk:"postgresql_output_option_merge_keys"`
}

func NewPostgresqlOutputOption(ctx context.Context, postgresqlOutputOption *output_option.PostgresqlOutputOption) *PostgresqlOutputOption {
	if postgresqlOutputOption == nil {
		return nil
	}

	result := &PostgresqlOutputOption{
		Database:               types.StringValue(postgresqlOutputOption.Database),
		Schema:                 types.StringValue(postgresqlOutputOption.Schema),
		Table:                  types.StringValue(postgresqlOutputOption.Table),
		Mode:                   types.StringPointerValue(postgresqlOutputOption.Mode),
		DefaultTimeZone:        types.StringPointerValue(postgresqlOutputOption.DefaultTimeZone),
		PostgresqlConnectionId: types.Int64Value(postgresqlOutputOption.PostgresqlConnectionId),
	}

	PostgresqlOutputOptionMergeKeys, err := newPostgresqlOutputOptionMergeKeys(ctx, postgresqlOutputOption.PostgresqlOutputOptionMergeKeys)
	if err != nil {
		return nil
	}
	result.PostgresqlOutputOptionMergeKeys = PostgresqlOutputOptionMergeKeys

	return result
}

func newPostgresqlOutputOptionMergeKeys(ctx context.Context, mergeKeys []string) (types.Set, error) {
	if len(mergeKeys) > 0 {
		values := make([]types.String, len(mergeKeys))
		for i, v := range mergeKeys {
			values[i] = types.StringValue(v)
		}
		setValue, diags := types.SetValueFrom(ctx, types.StringType, values)
		if diags.HasError() {
			return types.SetNull(types.StringType), fmt.Errorf("failed to convert to SetValue: %v", diags)
		}
		return setValue, nil
	}
	return types.SetNull(types.StringType), nil
}

func (postgresqlOutputOption *PostgresqlOutputOption) ToInput(ctx context.Context) *outputOptionParameters.PostgresqlOutputOptionInput {
	if postgresqlOutputOption == nil {
		return nil
	}

	var mergeKeys *[]string
	if !postgresqlOutputOption.PostgresqlOutputOptionMergeKeys.IsNull() && !postgresqlOutputOption.PostgresqlOutputOptionMergeKeys.IsUnknown() {
		var mergeKeyValues []types.String
		diags := postgresqlOutputOption.PostgresqlOutputOptionMergeKeys.ElementsAs(ctx, &mergeKeyValues, false)
		if diags.HasError() {
			return nil
		}
		if len(mergeKeyValues) > 0 {
			mk := make([]string, 0, len(mergeKeyValues))
			for _, input := range mergeKeyValues {
				mk = append(mk, input.ValueString())
			}
			mergeKeys = &mk
		}
	}

	return &outputOptionParameters.PostgresqlOutputOptionInput{
		Database:                        postgresqlOutputOption.Database.ValueString(),
		Schema:                          postgresqlOutputOption.Schema.ValueString(),
		Table:                           postgresqlOutputOption.Table.ValueString(),
		Mode:                            model.NewNullableString(postgresqlOutputOption.Mode),
		DefaultTimeZone:                 model.NewNullableString(postgresqlOutputOption.DefaultTimeZone),
		PostgresqlConnectionId:          postgresqlOutputOption.PostgresqlConnectionId.ValueInt64(),
		PostgresqlOutputOptionMergeKeys: model.WrapObjectList(mergeKeys),
	}
}

func (postgresqlOutputOption *PostgresqlOutputOption) ToUpdateInput(ctx context.Context) *outputOptionParameters.UpdatePostgresqlOutputOptionInput {
	if postgresqlOutputOption == nil {
		return nil
	}

	var mergeKeys *[]string
	if !postgresqlOutputOption.PostgresqlOutputOptionMergeKeys.IsNull() && !postgresqlOutputOption.PostgresqlOutputOptionMergeKeys.IsUnknown() {
		var mergeKeyValues []types.String
		diags := postgresqlOutputOption.PostgresqlOutputOptionMergeKeys.ElementsAs(ctx, &mergeKeyValues, false)
		if diags.HasError() {
			return nil
		}

		if len(mergeKeyValues) > 0 {
			mk := make([]string, 0, len(mergeKeyValues))
			for _, input := range mergeKeyValues {
				mk = append(mk, input.ValueString())
			}
			mergeKeys = &mk
		}
	}

	return &outputOptionParameters.UpdatePostgresqlOutputOptionInput{
		Database:                        postgresqlOutputOption.Database.ValueStringPointer(),
		Schema:                          postgresqlOutputOption.Schema.ValueStringPointer(),
		Table:                           postgresqlOutputOption.Table.ValueStringPointer(),
		Mode:                            model.NewNullableString(postgresqlOutputOption.Mode),
		DefaultTimeZone:                 model.NewNullableString(postgresqlOutputOption.DefaultTimeZone),
		PostgresqlConnectionId:          postgresqlOutputOption.PostgresqlConnectionId.ValueInt64Pointer(),
		PostgresqlOutputOptionMergeKeys: model.WrapObjectList(mergeKeys),
	}
}
