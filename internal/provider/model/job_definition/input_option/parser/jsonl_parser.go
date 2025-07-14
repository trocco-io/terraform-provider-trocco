package parser

import (
	"context"
	jobDefinitionEntities "terraform-provider-trocco/internal/client/entity/job_definition"
	jobDefinitionParameters "terraform-provider-trocco/internal/client/parameter/job_definition"
	"terraform-provider-trocco/internal/provider/model"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type JsonlParser struct {
	StopOnInvalidRecord types.Bool   `tfsdk:"stop_on_invalid_record"`
	DefaultTimeZone     types.String `tfsdk:"default_time_zone"`
	Newline             types.String `tfsdk:"newline"`
	Charset             types.String `tfsdk:"charset"`
	Columns             types.List   `tfsdk:"columns"`
}

type JsonlParserColumn struct {
	Name     types.String `tfsdk:"name"`
	Type     types.String `tfsdk:"type"`
	TimeZone types.String `tfsdk:"time_zone"`
	Format   types.String `tfsdk:"format"`
}

func NewJsonlParser(ctx context.Context, jsonlParser *jobDefinitionEntities.JsonlParser) *JsonlParser {
	if jsonlParser == nil {
		return nil
	}

	columnElements := make([]JsonlParserColumn, 0, len(jsonlParser.Columns))
	for _, input := range jsonlParser.Columns {
		column := JsonlParserColumn{
			Name:     types.StringValue(input.Name),
			Type:     types.StringValue(input.Type),
			TimeZone: types.StringPointerValue(input.TimeZone),
			Format:   types.StringPointerValue(input.Format),
		}
		columnElements = append(columnElements, column)
	}

	columns, diags := types.ListValueFrom(
		ctx,
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":      types.StringType,
				"type":      types.StringType,
				"time_zone": types.StringType,
				"format":    types.StringType,
			},
		},
		columnElements,
	)
	if diags.HasError() {
		return nil
	}

	return &JsonlParser{
		StopOnInvalidRecord: types.BoolPointerValue(jsonlParser.StopOnInvalidRecord),
		DefaultTimeZone:     types.StringValue(jsonlParser.DefaultTimeZone),
		Newline:             types.StringPointerValue(jsonlParser.Newline),
		Charset:             types.StringPointerValue(jsonlParser.Charset),
		Columns:             columns,
	}
}

func (jsonlParser *JsonlParser) ToJsonlParserInput(ctx context.Context) *jobDefinitionParameters.JsonlParserInput {
	if jsonlParser == nil {
		return nil
	}

	var columnElements []JsonlParserColumn
	diags := jsonlParser.Columns.ElementsAs(ctx, &columnElements, false)
	if diags.HasError() {
		return nil
	}

	columns := make([]jobDefinitionParameters.JsonlParserColumnInput, 0, len(columnElements))
	for _, input := range columnElements {
		column := jobDefinitionParameters.JsonlParserColumnInput{
			Name:     input.Name.ValueString(),
			Type:     input.Type.ValueString(),
			TimeZone: input.TimeZone.ValueStringPointer(),
			Format:   input.Format.ValueStringPointer(),
		}
		columns = append(columns, column)
	}

	return &jobDefinitionParameters.JsonlParserInput{
		StopOnInvalidRecord: jsonlParser.StopOnInvalidRecord.ValueBool(),
		DefaultTimeZone:     jsonlParser.DefaultTimeZone.ValueString(),
		Newline:             model.NewNullableString(jsonlParser.Newline),
		Charset:             model.NewNullableString(jsonlParser.Charset),
		Columns:             columns,
	}
}
