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

type FacebookAdsInsightsInputOption struct {
	FacebookAdsInsightsConnectionID types.Int64  `tfsdk:"facebook_ads_insights_connection_id"`
	AdAccountID                     types.String `tfsdk:"ad_account_id"`
	Level                           types.String `tfsdk:"level"`
	TimeRangeSince                  types.String `tfsdk:"time_range_since"`
	TimeRangeUntil                  types.String `tfsdk:"time_range_until"`
	UseUnifiedAttributionSetting    types.Bool   `tfsdk:"use_unified_attribution_setting"`
	Fields                          types.List   `tfsdk:"fields"`
	Breakdowns                      types.List   `tfsdk:"breakdowns"`
	ActionAttributionWindows        types.List   `tfsdk:"action_attribution_windows"`
	ActionBreakdowns                types.List   `tfsdk:"action_breakdowns"`
	CustomVariableSettings          types.List   `tfsdk:"custom_variable_settings"`
}

type FacebookAdsInsightsField struct {
	Name types.String `tfsdk:"name"`
}

type FacebookAdsInsightsBreakdown struct {
	Name types.String `tfsdk:"name"`
}

type FacebookAdsInsightsAttrWindow struct {
	Name types.String `tfsdk:"name"`
}

type FacebookAdsInsightsActionBreakdown struct {
	Name types.String `tfsdk:"name"`
}

func (FacebookAdsInsightsField) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": types.StringType,
	}
}

func (FacebookAdsInsightsBreakdown) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": types.StringType,
	}
}

func (FacebookAdsInsightsAttrWindow) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": types.StringType,
	}
}

func (FacebookAdsInsightsActionBreakdown) attrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"name": types.StringType,
	}
}

func NewFacebookAdsInsightsInputOption(ctx context.Context, inputOption *inputOptionEntities.FacebookAdsInsightsInputOption) *FacebookAdsInsightsInputOption {
	if inputOption == nil {
		return nil
	}

	result := &FacebookAdsInsightsInputOption{
		FacebookAdsInsightsConnectionID: types.Int64Value(inputOption.FacebookAdsInsightsConnectionID),
		AdAccountID:                     types.StringValue(inputOption.AdAccountID),
		Level:                           types.StringValue(inputOption.Level),
		TimeRangeSince:                  types.StringValue(inputOption.TimeRangeSince),
		TimeRangeUntil:                  types.StringValue(inputOption.TimeRangeUntil),
		UseUnifiedAttributionSetting:    types.BoolValue(inputOption.UseUnifiedAttributionSetting),
	}

	fields, err := newFacebookAdsInsightsFields(ctx, inputOption.Fields)
	if err != nil {
		return nil
	}
	result.Fields = fields

	breakdowns, err := newFacebookAdsInsightsBreakdowns(ctx, inputOption.Breakdowns)
	if err != nil {
		return nil
	}
	result.Breakdowns = breakdowns

	actionAttrWindows, err := newFacebookAdsInsightsAttrWindows(ctx, inputOption.ActionAttributionWindows)
	if err != nil {
		return nil
	}
	result.ActionAttributionWindows = actionAttrWindows

	actionBreakdowns, err := newFacebookAdsInsightsActionBreakdowns(ctx, inputOption.ActionBreakdowns)
	if err != nil {
		return nil
	}
	result.ActionBreakdowns = actionBreakdowns

	customVariableSettings, err := common.ConvertCustomVariableSettingsToList(ctx, inputOption.CustomVariableSettings)
	if err != nil {
		return nil
	}
	result.CustomVariableSettings = customVariableSettings

	return result
}

func newFacebookAdsInsightsFields(
	ctx context.Context,
	entityFields []inputOptionEntities.FacebookAdsInsightsField,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: FacebookAdsInsightsField{}.attrTypes(),
	}

	if entityFields == nil {
		return types.ListNull(objectType), nil
	}

	fields := make([]FacebookAdsInsightsField, 0, len(entityFields))
	for _, input := range entityFields {
		field := FacebookAdsInsightsField{
			Name: types.StringValue(input.Name),
		}
		fields = append(fields, field)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, fields)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert facebook ads insights fields to ListValue: %v", diags)
	}
	return listValue, nil
}

func newFacebookAdsInsightsBreakdowns(
	ctx context.Context,
	entityBreakdowns []inputOptionEntities.FacebookAdsInsightsBreakdown,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: FacebookAdsInsightsBreakdown{}.attrTypes(),
	}

	if entityBreakdowns == nil {
		return types.ListNull(objectType), nil
	}

	breakdowns := make([]FacebookAdsInsightsBreakdown, 0, len(entityBreakdowns))
	for _, input := range entityBreakdowns {
		breakdown := FacebookAdsInsightsBreakdown{
			Name: types.StringValue(input.Name),
		}
		breakdowns = append(breakdowns, breakdown)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, breakdowns)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert facebook ads insights breakdowns to ListValue: %v", diags)
	}
	return listValue, nil
}

func newFacebookAdsInsightsAttrWindows(
	ctx context.Context,
	entityWindows []inputOptionEntities.FacebookAdsInsightsAttrWindow,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: FacebookAdsInsightsAttrWindow{}.attrTypes(),
	}

	if entityWindows == nil {
		return types.ListNull(objectType), nil
	}

	windows := make([]FacebookAdsInsightsAttrWindow, 0, len(entityWindows))
	for _, input := range entityWindows {
		window := FacebookAdsInsightsAttrWindow{
			Name: types.StringValue(input.Name),
		}
		windows = append(windows, window)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, windows)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert facebook ads insights action attribution windows to ListValue: %v", diags)
	}
	return listValue, nil
}

func newFacebookAdsInsightsActionBreakdowns(
	ctx context.Context,
	entityBreakdowns []inputOptionEntities.FacebookAdsInsightsActionBreakdown,
) (types.List, error) {
	objectType := types.ObjectType{
		AttrTypes: FacebookAdsInsightsActionBreakdown{}.attrTypes(),
	}

	if entityBreakdowns == nil {
		return types.ListNull(objectType), nil
	}

	breakdowns := make([]FacebookAdsInsightsActionBreakdown, 0, len(entityBreakdowns))
	for _, input := range entityBreakdowns {
		breakdown := FacebookAdsInsightsActionBreakdown{
			Name: types.StringValue(input.Name),
		}
		breakdowns = append(breakdowns, breakdown)
	}

	listValue, diags := types.ListValueFrom(ctx, objectType, breakdowns)
	if diags.HasError() {
		return types.ListNull(objectType), fmt.Errorf("failed to convert facebook ads insights action breakdowns to ListValue: %v", diags)
	}
	return listValue, nil
}

func (inputOption *FacebookAdsInsightsInputOption) ToInput(ctx context.Context) *inputOptionParameters.FacebookAdsInsightsInputOptionInput {
	if inputOption == nil {
		return nil
	}

	var fieldValues []FacebookAdsInsightsField
	if !inputOption.Fields.IsNull() && !inputOption.Fields.IsUnknown() {
		diags := inputOption.Fields.ElementsAs(ctx, &fieldValues, false)
		if diags.HasError() {
			return nil
		}
	}

	var breakdownValues []FacebookAdsInsightsBreakdown
	if !inputOption.Breakdowns.IsNull() && !inputOption.Breakdowns.IsUnknown() {
		diags := inputOption.Breakdowns.ElementsAs(ctx, &breakdownValues, false)
		if diags.HasError() {
			return nil
		}
	}

	var attrWindowValues []FacebookAdsInsightsAttrWindow
	if !inputOption.ActionAttributionWindows.IsNull() && !inputOption.ActionAttributionWindows.IsUnknown() {
		diags := inputOption.ActionAttributionWindows.ElementsAs(ctx, &attrWindowValues, false)
		if diags.HasError() {
			return nil
		}
	}

	var actionBreakdownValues []FacebookAdsInsightsActionBreakdown
	if !inputOption.ActionBreakdowns.IsNull() && !inputOption.ActionBreakdowns.IsUnknown() {
		diags := inputOption.ActionBreakdowns.ElementsAs(ctx, &actionBreakdownValues, false)
		if diags.HasError() {
			return nil
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, inputOption.CustomVariableSettings)

	return &inputOptionParameters.FacebookAdsInsightsInputOptionInput{
		FacebookAdsInsightsConnectionID: inputOption.FacebookAdsInsightsConnectionID.ValueInt64(),
		AdAccountID:                     inputOption.AdAccountID.ValueString(),
		Level:                           inputOption.Level.ValueString(),
		TimeRangeSince:                  inputOption.TimeRangeSince.ValueString(),
		TimeRangeUntil:                  inputOption.TimeRangeUntil.ValueString(),
		UseUnifiedAttributionSetting:    inputOption.UseUnifiedAttributionSetting.ValueBool(),
		Fields:                          toFacebookAdsInsightsFieldsInput(fieldValues),
		Breakdowns:                      toFacebookAdsInsightsBreakdownsInput(breakdownValues),
		ActionAttributionWindows:        toFacebookAdsInsightsAttrWindowsInput(attrWindowValues),
		ActionBreakdowns:                toFacebookAdsInsightsActionBreakdownsInput(actionBreakdownValues),
		CustomVariableSettings:          model.ToCustomVariableSettingInputs(customVarSettings),
	}
}

func (inputOption *FacebookAdsInsightsInputOption) ToUpdateInput(ctx context.Context) *inputOptionParameters.UpdateFacebookAdsInsightsInputOptionInput {
	if inputOption == nil {
		return nil
	}

	var fieldValues []FacebookAdsInsightsField
	if !inputOption.Fields.IsNull() {
		if !inputOption.Fields.IsUnknown() {
			diags := inputOption.Fields.ElementsAs(ctx, &fieldValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			fieldValues = []FacebookAdsInsightsField{}
		}
	} else {
		fieldValues = []FacebookAdsInsightsField{}
	}

	var breakdownValues []FacebookAdsInsightsBreakdown
	if !inputOption.Breakdowns.IsNull() {
		if !inputOption.Breakdowns.IsUnknown() {
			diags := inputOption.Breakdowns.ElementsAs(ctx, &breakdownValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			breakdownValues = []FacebookAdsInsightsBreakdown{}
		}
	}

	var attrWindowValues []FacebookAdsInsightsAttrWindow
	if !inputOption.ActionAttributionWindows.IsNull() {
		if !inputOption.ActionAttributionWindows.IsUnknown() {
			diags := inputOption.ActionAttributionWindows.ElementsAs(ctx, &attrWindowValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			attrWindowValues = []FacebookAdsInsightsAttrWindow{}
		}
	}

	var actionBreakdownValues []FacebookAdsInsightsActionBreakdown
	if !inputOption.ActionBreakdowns.IsNull() {
		if !inputOption.ActionBreakdowns.IsUnknown() {
			diags := inputOption.ActionBreakdowns.ElementsAs(ctx, &actionBreakdownValues, false)
			if diags.HasError() {
				return nil
			}
		} else {
			actionBreakdownValues = []FacebookAdsInsightsActionBreakdown{}
		}
	}

	customVarSettings := common.ExtractCustomVariableSettings(ctx, inputOption.CustomVariableSettings)

	result := &inputOptionParameters.UpdateFacebookAdsInsightsInputOptionInput{
		FacebookAdsInsightsConnectionID: model.NewNullableInt64(inputOption.FacebookAdsInsightsConnectionID),
		AdAccountID:                     model.NewNullableString(inputOption.AdAccountID),
		Level:                           model.NewNullableString(inputOption.Level),
		TimeRangeSince:                  model.NewNullableString(inputOption.TimeRangeSince),
		TimeRangeUntil:                  model.NewNullableString(inputOption.TimeRangeUntil),
		UseUnifiedAttributionSetting:    inputOption.UseUnifiedAttributionSetting.ValueBoolPointer(),
		CustomVariableSettings:          model.ToCustomVariableSettingInputs(customVarSettings),
	}

	if fieldValues != nil {
		fields := toFacebookAdsInsightsFieldsInput(fieldValues)
		result.Fields = &fields
	}

	if breakdownValues != nil {
		breakdowns := toFacebookAdsInsightsBreakdownsInput(breakdownValues)
		result.Breakdowns = &breakdowns
	}

	if attrWindowValues != nil {
		windows := toFacebookAdsInsightsAttrWindowsInput(attrWindowValues)
		result.ActionAttributionWindows = &windows
	}

	if actionBreakdownValues != nil {
		actionBreakdowns := toFacebookAdsInsightsActionBreakdownsInput(actionBreakdownValues)
		result.ActionBreakdowns = &actionBreakdowns
	}

	return result
}

func toFacebookAdsInsightsFieldsInput(fields []FacebookAdsInsightsField) []inputOptionParameters.FacebookAdsInsightsField {
	if fields == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.FacebookAdsInsightsField, 0, len(fields))
	for _, field := range fields {
		inputs = append(inputs, inputOptionParameters.FacebookAdsInsightsField{
			Name: field.Name.ValueString(),
		})
	}
	return inputs
}

func toFacebookAdsInsightsBreakdownsInput(breakdowns []FacebookAdsInsightsBreakdown) []inputOptionParameters.FacebookAdsInsightsBreakdown {
	if breakdowns == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.FacebookAdsInsightsBreakdown, 0, len(breakdowns))
	for _, breakdown := range breakdowns {
		inputs = append(inputs, inputOptionParameters.FacebookAdsInsightsBreakdown{
			Name: breakdown.Name.ValueString(),
		})
	}
	return inputs
}

func toFacebookAdsInsightsAttrWindowsInput(windows []FacebookAdsInsightsAttrWindow) []inputOptionParameters.FacebookAdsInsightsAttrWindow {
	if windows == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.FacebookAdsInsightsAttrWindow, 0, len(windows))
	for _, window := range windows {
		inputs = append(inputs, inputOptionParameters.FacebookAdsInsightsAttrWindow{
			Name: window.Name.ValueString(),
		})
	}
	return inputs
}

func toFacebookAdsInsightsActionBreakdownsInput(breakdowns []FacebookAdsInsightsActionBreakdown) []inputOptionParameters.FacebookAdsInsightsActionBreakdown {
	if breakdowns == nil {
		return nil
	}

	inputs := make([]inputOptionParameters.FacebookAdsInsightsActionBreakdown, 0, len(breakdowns))
	for _, breakdown := range breakdowns {
		inputs = append(inputs, inputOptionParameters.FacebookAdsInsightsActionBreakdown{
			Name: breakdown.Name.ValueString(),
		})
	}
	return inputs
}
