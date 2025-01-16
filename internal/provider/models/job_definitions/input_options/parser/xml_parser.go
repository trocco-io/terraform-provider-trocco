package parser

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions"
	params "terraform-provider-trocco/internal/client/parameters/job_definitions"
)

type XmlParser struct {
	Root    types.String      `tfsdk:"root"`
	Columns []XmlParserColumn `tfsdk:"columns"`
}

type XmlParserColumn struct {
	Name     types.String `tfsdk:"name"`
	Type     types.String `tfsdk:"type"`
	Path     types.String `tfsdk:"path"`
	Timezone types.String `tfsdk:"timezone"`
	Format   types.String `tfsdk:"format"`
}

func NewXmlParser(xmlParser *job_definitions.XmlParser) *XmlParser {
	if xmlParser == nil {
		return nil
	}
	columns := make([]XmlParserColumn, 0, len(xmlParser.Columns))
	for _, input := range xmlParser.Columns {
		column := XmlParserColumn{
			Name:     types.StringValue(input.Name),
			Type:     types.StringValue(input.Type),
			Path:     types.StringValue(input.Path),
			Timezone: types.StringPointerValue(input.Timezone),
			Format:   types.StringPointerValue(input.Format),
		}
		columns = append(columns, column)
	}

	return &XmlParser{
		Root:    types.StringValue(xmlParser.Root),
		Columns: columns,
	}
}

func (xmlParser *XmlParser) ToXmlParserInput() *params.XmlParserInput {
	if xmlParser == nil {
		return nil
	}
	columns := make([]params.XmlParserColumnInput, 0, len(xmlParser.Columns))
	for _, input := range xmlParser.Columns {
		column := params.XmlParserColumnInput{
			Name:     input.Name.ValueString(),
			Type:     input.Type.ValueString(),
			Path:     input.Path.ValueString(),
			Format:   input.Format.ValueStringPointer(),
			Timezone: input.Timezone.ValueStringPointer(),
		}
		columns = append(columns, column)
	}

	return &params.XmlParserInput{
		Root:    xmlParser.Root.ValueString(),
		Columns: columns,
	}
}

func ToXmlParserModel(xmlParser *job_definitions.XmlParser) *XmlParser {
	if xmlParser == nil {
		return nil
	}
	columns := make([]XmlParserColumn, 0, len(xmlParser.Columns))
	for _, input := range xmlParser.Columns {
		column := XmlParserColumn{
			Name: types.StringValue(input.Name),
			Type: types.StringValue(input.Type),
		}
		columns = append(columns, column)
	}

	return &XmlParser{
		Root:    types.StringValue(xmlParser.Root),
		Columns: columns,
	}
}
