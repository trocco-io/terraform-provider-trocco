package input_options

import (
	"context"
	"fmt"
	inputOptionEntities "terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	inputOptionParameters "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definition/common"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type YahooAdsApiYdnInputOption struct {
	YahooAdsApiConnectionID          types.Int64  `tfsdk:"yahoo_ads_api_connection_id"`
	Target                           types.String `tfsdk:"target"`
	AccountID                        types.String `tfsdk:"account_id"`
	BaseAccountID                    types.String `tfsdk:"base_account_id"`
	ReportType                       types.String `tfsdk:"report_type"`
	StartDate                        types.String `tfsdk:"start_date"`
	EndDate                          types.String `tfsdk:"end_date"`
	IncludeDeleted                   types.Bool   `tfsdk:"include_deleted"`
	YahooAdsApiYdnInputOptionColumns types.List   `tfsdk:"input_option_columns"`
	CustomVariableSettings           types.List   `tfsdk:"custom_variable_settings"`
}

type YahooAdsApiYdnInputOptionColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func NewYahooAdsApiYdnInputOption(ctx context.Context, inputOption *inputOptionEntities.YahooAdsApiYdnInputOption) *YahooAdsApiYdnInputOption {
	if inputOption == nil {
		return nil
	}

	result := &YahooAdsApiYdnInputOption{
		YahooAdsApiConnectionID: types.Int64Value(inputOption.YahooAdsApiConnectionID),
		Target:                  types.StringValue(inputOption.Target),
		AccountID:               types.StringValue(inputOption.AccountID),
		BaseAccountID:           types.StringPointerValue(inputOption.BaseAccountID),
		ReportType:              types.StringPointerValue(inputOption.ReportType),
		StartDate:               types.StringValue(inputOption.StartDate),
		EndDate:                 types.StringValue(inputOption.EndDate),
		IncludeDeleted:          types.BoolValue(inputOption.IncludeDeleted),
	}

	columns, err := newYahooAdsApiYdnInputOptionColumns(ctx, inputOption.YahooAdsApiYdnInputOptionColumns)
	if err != nil {
		return nil
	}
	result.YahooAdsApiYdnInputOptionColumns = columns

	customVarSettings, err := common.ConvertCustomVariableSettingsToList(ctx, inputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVarSettings

	return result
}

func newYahooAdsApiYdnInputOptionColumns(
	ctx context.Context,
	columns []inputOptionEntities.YahooAdsApiYdnInputOptionColumn,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: YahooAdsApiYdnInputOptionColumn{}.attrTypes(),
	}

	if columns == nil {
		return types.ListNull(objectType), nil
	}

	result := make([]YahooAdsApiYdnInputOptionColumn, 0, len(columns))
	for _, column := range columns {
		col := YahooAdsApiYdnInputOptionColumn{
			Name:   types.StringValue(column.Name),
			Type:   types.StringValue(column.Type),
			Format: types.StringPointerValue(column.Format),
		}
		result = append(result, col)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, result)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert columns to ListValue: %v", diags)
	}
	return listValue, nil
}

func (YahooAdsApiYdnInputOptionColumn) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name":   types.StringType,
		"type":   types.StringType,
		"format": types.StringType,
	}
}

func (inputOption *YahooAdsApiYdnInputOption) ToInput(ctx context.Context) *inputOptionParameters.YahooAdsApiYdnInputOptionInput {
	if inputOption == nil {
		return nil
	}

	var columnValues []YahooAdsApiYdnInputOptionColumn
	if !inputOption.YahooAdsApiYdnInputOptionColumns.IsNull() && !inputOption.YahooAdsApiYdnInputOptionColumns.IsUnknown() {
		diags := inputOption.YahooAdsApiYdnInputOptionColumns.ElementsAs(ctx, &columnValues, false)
		if diags.HasError() {
			return nil
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, inputOption.CustomVariableSettings)

	result := &inputOptionParameters.YahooAdsApiYdnInputOptionInput{
		YahooAdsApiConnectionID:          inputOption.YahooAdsApiConnectionID.ValueInt64(),
		Target:                           inputOption.Target.ValueString(),
		AccountID:                        inputOption.AccountID.ValueString(),
		BaseAccountID:                    model.NewNullableString(inputOption.BaseAccountID),
		ReportType:                       model.NewNullableString(inputOption.ReportType),
		StartDate:                        inputOption.StartDate.ValueString(),
		EndDate:                          inputOption.EndDate.ValueString(),
		YahooAdsApiYdnInputOptionColumns: toYahooAdsApiYdnInputOptionColumnsInput(columnValues),
		CustomVariableSettings:           model.ToCustomVariableSettingInputs(customVarSettings),
	}

	// Only include_deleted is valid when target is "report"
	if inputOption.Target.ValueString() == "report" {
		result.IncludeDeleted = model.NewNullableBool(inputOption.IncludeDeleted)
	}

	return result
}

func (inputOption *YahooAdsApiYdnInputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateYahooAdsApiYdnInputOptionInput {
	if inputOption == nil {
		return nil
	}

	var columnValues []YahooAdsApiYdnInputOptionColumn
	if !inputOption.YahooAdsApiYdnInputOptionColumns.IsNull() {
		if !inputOption.YahooAdsApiYdnInputOptionColumns.IsUnknown() {
			diags := inputOption.YahooAdsApiYdnInputOptionColumns.ElementsAs(ctx, &columnValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			columnValues = []YahooAdsApiYdnInputOptionColumn{}
		}
	} else {
		columnValues = []YahooAdsApiYdnInputOptionColumn{}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, inputOption.CustomVariableSettings)

	result := &inputOptionParameters.UpdateYahooAdsApiYdnInputOptionInput{
		YahooAdsApiConnectionID:          model.NewNullableInt64(inputOption.YahooAdsApiConnectionID),
		Target:                           model.NewNullableString(inputOption.Target),
		AccountID:                        model.NewNullableString(inputOption.AccountID),
		BaseAccountID:                    model.NewNullableString(inputOption.BaseAccountID),
		ReportType:                       model.NewNullableString(inputOption.ReportType),
		StartDate:                        model.NewNullableString(inputOption.StartDate),
		EndDate:                          model.NewNullableString(inputOption.EndDate),
		YahooAdsApiYdnInputOptionColumns: toYahooAdsApiYdnInputOptionColumnsInput(columnValues),
		CustomVariableSettings:           model.ToCustomVariableSettingInputs(customVarSettings),
	}

	// Only include_deleted is valid when target is "report"
	if inputOption.Target.ValueString() == "report" {
		result.IncludeDeleted = model.NewNullableBool(inputOption.IncludeDeleted)
	}

	return result
}

func toYahooAdsApiYdnInputOptionColumnsInput(columns []YahooAdsApiYdnInputOptionColumn) []inputOptionParameters.YahooAdsApiYdnInputOptionColumn {
	if columns == nil {
		return nil
	}
	result := make([]inputOptionParameters.YahooAdsApiYdnInputOptionColumn, 0, len(columns))
	for _, column := range columns {
		result = append(result, inputOptionParameters.YahooAdsApiYdnInputOptionColumn{
			Name:   column.Name.ValueString(),
			Type:   column.Type.ValueString(),
			Format: column.Format.ValueStringPointer(),
		})
	}
	return result
}
