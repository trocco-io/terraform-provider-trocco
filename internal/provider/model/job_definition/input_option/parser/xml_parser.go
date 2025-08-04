package parser

import (
	"context"
	jobDefinitionEntities "terraform-provider-trocco/internal/client/entity/job_definition"
	jobDefinitionParameters "terraform-provider-trocco/internal/client/parameter/job_definition"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type XmlParser struct {
	Root    types.String `tfsdk:"root"`
	Columns types.List   `tfsdk:"columns"`
}

type XmlParserColumn struct {
	Name     types.String `tfsdk:"name"`
	Type     types.String `tfsdk:"type"`
	Path     types.String `tfsdk:"path"`
	Timezone types.String `tfsdk:"timezone"`
	Format   types.String `tfsdk:"format"`
}

func NewXmlParser(ctx context.Context, xmlParser *jobDefinitionEntities.XmlParser) *XmlParser {
	if xmlParser == nil {
		return nil
	}

	columnElements := make([]XmlParserColumn, 0, len(xmlParser.Columns))
	for _, input := range xmlParser.Columns {
		column := XmlParserColumn{
			Name:     types.StringValue(input.Name),
			Type:     types.StringValue(input.Type),
			Path:     types.StringValue(input.Path),
			Timezone: types.StringPointerValue(input.Timezone),
			Format:   types.StringPointerValue(input.Format),
		}
		columnElements = append(columnElements, column)
	}

	columns, diags := types.ListValueFrom(
		ctx,
		types.ObjectType{
			AttrTypes: map[string]attr.Type{
				"name":     types.StringType,
				"type":     types.StringType,
				"path":     types.StringType,
				"timezone": types.StringType,
				"format":   types.StringType,
			},
		},
		columnElements,
	)
	if diags.HasError() {
		return nil
	}

	return &XmlParser{
		Root:    types.StringValue(xmlParser.Root),
		Columns: columns,
	}
}

func (xmlParser *XmlParser) ToXmlParserInput(ctx context.Context) *jobDefinitionParameters.XmlParserInput {
	if xmlParser == nil {
		return nil
	}

	var columnElements []XmlParserColumn
	diags := xmlParser.Columns.ElementsAs(ctx, &columnElements, false)
	if diags.HasError() {
		return nil
	}

	columns := make([]jobDefinitionParameters.XmlParserColumnInput, 0, len(columnElements))
	for _, input := range columnElements {
		column := jobDefinitionParameters.XmlParserColumnInput{
			Name:     input.Name.ValueString(),
			Type:     input.Type.ValueString(),
			Path:     input.Path.ValueString(),
			Format:   input.Format.ValueStringPointer(),
			Timezone: input.Timezone.ValueStringPointer(),
		}
		columns = append(columns, column)
	}

	return &jobDefinitionParameters.XmlParserInput{
		Root:    xmlParser.Root.ValueString(),
		Columns: columns,
	}
}
