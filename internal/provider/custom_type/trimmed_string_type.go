package custom_type

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ basetypes.StringTypable = TrimmedStringType{}

type TrimmedStringType struct {
	basetypes.StringType
}

func (t TrimmedStringType) Equal(o attr.Type) bool {
	other, ok := o.(TrimmedStringType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t TrimmedStringType) String() string {
	return "TrimmedStringType"
}

func (t TrimmedStringType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	value := TrimmedStringValue{
		StringValue: in,
	}

	return value, nil
}

func (t TrimmedStringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)

	if err != nil {
		return nil, err
	}

	stringValue, ok := attrValue.(basetypes.StringValue)

	if !ok {
		return nil, fmt.Errorf("unexpected value type of %T", attrValue)
	}

	stringValuable, diags := t.ValueFromString(ctx, stringValue)

	if diags.HasError() {
		return nil, fmt.Errorf("unexpected error converting StringValue to StringValuable: %v", diags)
	}

	return stringValuable, nil
}

func (t TrimmedStringType) ValueType(ctx context.Context) attr.Value {
	return TrimmedStringValue{}
}

var _ basetypes.StringValuable = TrimmedStringValue{}

type TrimmedStringValue struct {
	basetypes.StringValue
}

func (v TrimmedStringValue) Equal(o attr.Value) bool {
	other, ok := o.(TrimmedStringValue)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v TrimmedStringValue) Type(ctx context.Context) attr.Type {
	return TrimmedStringType{}
}

func (v TrimmedStringValue) StringSemanticEquals(ctx context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(TrimmedStringValue)

	if !ok {
		diags.AddError(
			"Semantic Equality Check Error",
			"An unexpected value type was received while performing semantic equality checks. "+
				"Please report this to the provider developers.\n\n"+
				"Expected Value Type: "+fmt.Sprintf("%T", v)+"\n"+
				"Got Value Type: "+fmt.Sprintf("%T", newValuable),
		)

		return false, diags
	}

	priorTrimmed := strings.TrimSpace(v.ValueString())
	newTrimmed := strings.TrimSpace(newValue.ValueString())

	return priorTrimmed == newTrimmed, diags
}
