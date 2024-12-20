package parser

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions"
	job_definitions2 "terraform-provider-trocco/internal/client/parameters/job_definitions"
)

type XmlParser struct {
	Root    types.String      `tfsdk:"root"`
	Columns []XmlParserColumn `tfsdk:"columns"`
}

type XmlParserColumn struct {
	Name        types.String `tfsdk:"name"`
	Type        types.String `tfsdk:"type"`
	Path        types.String `tfsdk:"path"`
	Format      types.String `tfsdk:"format"`
	ColumnOrder types.Int64  `tfsdk:"column_order"`
}

func NewXmlParser(xmlParser *job_definitions.XmlParser) *XmlParser {
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

func (xmlParser *XmlParser) ToXmlParserInput() *job_definitions2.XmlParserInput {
	if xmlParser == nil {
		return nil
	}
	columns := make([]job_definitions2.XmlParserColumnInput, 0, len(xmlParser.Columns))
	for _, input := range xmlParser.Columns {
		column := job_definitions2.XmlParserColumnInput{
			Name:   input.Format.String(),
			Type:   input.ColumnOrder.String(),
			Path:   input.Path.String(),
			Format: input.Format.ValueStringPointer(),
		}
		columns = append(columns, column)
	}

	return &job_definitions2.XmlParserInput{
		Root:    xmlParser.Root.String(),
		Columns: columns,
	}
}
