package workflow

import (
	we "terraform-provider-trocco/internal/client/entities/pipeline_definition"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type CustomVariable struct {
	Name      types.String `tfsdk:"name"`
	Type      types.String `tfsdk:"type"`
	Value     types.String `tfsdk:"value"`
	Quantity  types.Int64  `tfsdk:"quantity"`
	Unit      types.String `tfsdk:"unit"`
	Direction types.String `tfsdk:"direction"`
	Format    types.String `tfsdk:"format"`
	TimeZone  types.String `tfsdk:"time_zone"`
}

func NewCustomVariables(ens []we.CustomVariable) []CustomVariable {
	if len(ens) == 0 {
		// If no custom variables are present, the API returns an empty array but the provider should set `null`.
		return nil
	}

	var mds []CustomVariable
	for _, en := range ens {
		mds = append(mds, NewCustomVariable(en))
	}

	return mds
}

func NewCustomVariable(en we.CustomVariable) CustomVariable {
	return CustomVariable{
		Name:      types.StringPointerValue(en.Name),
		Type:      types.StringPointerValue(en.Type),
		Value:     types.StringPointerValue(en.Value),
		Quantity:  types.Int64PointerValue(en.Quantity),
		Unit:      types.StringPointerValue(en.Unit),
		Direction: types.StringPointerValue(en.Direction),
		Format:    types.StringPointerValue(en.Format),
		TimeZone:  types.StringPointerValue(en.TimeZone),
	}
}