package parser

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions"
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
