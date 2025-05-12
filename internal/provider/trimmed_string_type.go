package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

var _ basetypes.StringTypable = trimmedStringType{}

type trimmedStringType struct {
	basetypes.StringType
}

func (t trimmedStringType) Equal(o attr.Type) bool {
	other, ok := o.(trimmedStringType)

	if !ok {
		return false
	}

	return t.StringType.Equal(other.StringType)
}

func (t trimmedStringType) String() string {
	return "trimmedStringType"
}

func (t trimmedStringType) ValueFromString(ctx context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	value := trimmedStringValue{
		StringValue: in,
	}

	return value, nil
}

func (t trimmedStringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
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

func (t trimmedStringType) ValueType(ctx context.Context) attr.Value {
	return trimmedStringValue{}
}

var _ basetypes.StringValuable = trimmedStringValue{}

type trimmedStringValue struct {
	basetypes.StringValue
}

func (v trimmedStringValue) Equal(o attr.Value) bool {
	other, ok := o.(trimmedStringValue)

	if !ok {
		return false
	}

	return v.StringValue.Equal(other.StringValue)
}

func (v trimmedStringValue) Type(ctx context.Context) attr.Type {
	return trimmedStringType{}
}

func (v trimmedStringValue) StringSemanticEquals(ctx context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics

	newValue, ok := newValuable.(trimmedStringValue)

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
