package input_options

import (
	"terraform-provider-trocco/internal/client/entity/job_definition/input_option"
	param "terraform-provider-trocco/internal/client/parameter/job_definition/input_option"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type YahooAdsApiYssInputOption struct {
	AccountID               types.String                      `tfsdk:"account_id"`
	BaseAccountID           types.String                      `tfsdk:"base_account_id"`
	Service                 types.String                      `tfsdk:"service"`
	ExcludeZeroImpressions  types.Bool                        `tfsdk:"exclude_zero_impressions"`
	ReportType              types.String                      `tfsdk:"report_type"`
	StartDate               types.String                      `tfsdk:"start_date"`
	EndDate                 types.String                      `tfsdk:"end_date"`
	YahooAdsApiConnectionID types.Int64                       `tfsdk:"yahoo_ads_api_connection_id"`
	InputOptionColumns      []YahooAdsApiYssInputOptionColumn `tfsdk:"input_option_columns"`
	CustomVariableSettings  *[]model.CustomVariableSetting    `tfsdk:"custom_variable_settings"`
}

type YahooAdsApiYssInputOptionColumn struct {
	Name   types.String `tfsdk:"name"`
	Type   types.String `tfsdk:"type"`
	Format types.String `tfsdk:"format"`
}

func NewYahooAdsApiYssInputOption(inputOption *input_option.YahooAdsApiYssInputOption) *YahooAdsApiYssInputOption {
	if inputOption == nil {
		return nil
	}

	return &YahooAdsApiYssInputOption{
		AccountID:               types.StringValue(inputOption.AccountID),
		BaseAccountID:           types.StringValue(inputOption.BaseAccountID),
		Service:                 types.StringValue(inputOption.Service),
		ExcludeZeroImpressions:  types.BoolValue(inputOption.ExcludeZeroImpressions),
		ReportType:              types.StringPointerValue(inputOption.ReportType),
		StartDate:               types.StringPointerValue(inputOption.StartDate),
		EndDate:                 types.StringPointerValue(inputOption.EndDate),
		YahooAdsApiConnectionID: types.Int64Value(inputOption.YahooAdsApiConnectionID),
		InputOptionColumns:      newYahooAdsApiYssInputOptionColumns(inputOption.InputOptionColumns),
		CustomVariableSettings:  model.NewCustomVariableSettings(inputOption.CustomVariableSettings),
	}
}

func newYahooAdsApiYssInputOptionColumns(inputOptionColumns []input_option.YahooAdsApiYssInputOptionColumn) []YahooAdsApiYssInputOptionColumn {
	if inputOptionColumns == nil {
		return nil
	}
	columns := make([]YahooAdsApiYssInputOptionColumn, 0, len(inputOptionColumns))
	for _, input := range inputOptionColumns {
		column := YahooAdsApiYssInputOptionColumn{
			Name:   types.StringValue(input.Name),
			Type:   types.StringValue(input.Type),
			Format: types.StringPointerValue(input.Format),
		}
		columns = append(columns, column)
	}
	return columns
}

func (inputOption *YahooAdsApiYssInputOption) ToInput() *param.YahooAdsApiYssInputOptionInput {
	if inputOption == nil {
		return nil
	}

	return &param.YahooAdsApiYssInputOptionInput{
		AccountID:               inputOption.AccountID.ValueString(),
		BaseAccountID:           inputOption.BaseAccountID.ValueString(),
		Service:                 model.NewNullableString(inputOption.Service),
		ExcludeZeroImpressions:  inputOption.ExcludeZeroImpressions.ValueBool(),
		ReportType:              model.NewNullableString(inputOption.ReportType),
		StartDate:               model.NewNullableString(inputOption.StartDate),
		EndDate:                 model.NewNullableString(inputOption.EndDate),
		YahooAdsApiConnectionID: inputOption.YahooAdsApiConnectionID.ValueInt64(),
		InputOptionColumns:      toYahooAdsApiYssInputOptionColumnsInput(inputOption.InputOptionColumns),
		CustomVariableSettings:  model.ToCustomVariableSettingInputs(inputOption.CustomVariableSettings),
	}
}

func (inputOption *YahooAdsApiYssInputOption) ToUpdateInput() *param.UpdateYahooAdsApiYssInputOptionInput {
	if inputOption == nil {
		return nil
	}

	return &param.UpdateYahooAdsApiYssInputOptionInput{
		AccountID:               model.NewNullableString(inputOption.AccountID),
		BaseAccountID:           model.NewNullableString(inputOption.BaseAccountID),
		Service:                 model.NewNullableString(inputOption.Service),
		ExcludeZeroImpressions:  model.NewNullableBool(inputOption.ExcludeZeroImpressions),
		ReportType:              model.NewNullableString(inputOption.ReportType),
		StartDate:               model.NewNullableString(inputOption.StartDate),
		EndDate:                 model.NewNullableString(inputOption.EndDate),
		YahooAdsApiConnectionID: model.NewNullableInt64(inputOption.YahooAdsApiConnectionID),
		InputOptionColumns:      toYahooAdsApiYssInputOptionColumnsInput(inputOption.InputOptionColumns),
		CustomVariableSettings:  model.ToCustomVariableSettingInputs(inputOption.CustomVariableSettings),
	}
}

func toYahooAdsApiYssInputOptionColumnsInput(columns []YahooAdsApiYssInputOptionColumn) []param.YahooAdsApiYssInputOptionColumn {
	if columns == nil {
		return nil
	}
	result := make([]param.YahooAdsApiYssInputOptionColumn, 0, len(columns))
	for _, column := range columns {
		result = append(result, param.YahooAdsApiYssInputOptionColumn{
			Name:   column.Name.ValueString(),
			Type:   column.Type.ValueString(),
			Format: column.Format.ValueStringPointer(),
		})
	}
	return result
}
