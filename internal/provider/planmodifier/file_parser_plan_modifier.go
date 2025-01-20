package planmodifier

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strconv"
)

var _ planmodifier.Object = &FileParserPlanModifier{}

type FileParserPlanModifier struct{}

func (d *FileParserPlanModifier) Description(ctx context.Context) string {
	return "Modifier for validating schedule attributes"
}

func (d *FileParserPlanModifier) MarkdownDescription(ctx context.Context) string {
	return d.Description(ctx)
}

func (d *FileParserPlanModifier) PlanModifyObject(ctx context.Context, req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse) {
	var csvParser types.Object
	var jsonlParser types.Object
	var ltsvParser types.Object
	var excelParser types.Object
	var xmlParser types.Object
	var jsonpathParser types.Object
	var parquetParser types.Object

	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("csv_parser"), &csvParser)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("jsonl_parser"), &jsonlParser)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("ltsv_parser"), &ltsvParser)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("excel_parser"), &excelParser)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("xml_parser"), &xmlParser)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("jsonpath_parser"), &jsonpathParser)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("parquet_parser"), &parquetParser)...)
	if resp.Diagnostics.HasError() {
		return
	}

	nonNilParserCount := countNonNil(csvParser, jsonlParser, ltsvParser, excelParser, xmlParser, jsonpathParser, parquetParser)

	if nonNilParserCount > 1 {
		addFileParserAttributeError(req, resp, strconv.Itoa(nonNilParserCount)+"number of parser is excessive. please specify only one parser")
	}

	var lastPath types.String
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("last_path"), &lastPath)...)
	if resp.Diagnostics.HasError() {
		return
	}
	var incrementalLoadingEnabled types.Bool
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, req.Path.AtName("incremental_loading_enabled"), &incrementalLoadingEnabled)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !incrementalLoadingEnabled.ValueBool() && !lastPath.IsNull() {
		addFileParserAttributeError(req, resp, "last_path is only valid when incremental_loading_enabled is true")
	}
}

func addFileParserAttributeError(req planmodifier.ObjectRequest, resp *planmodifier.ObjectResponse, message string) {
	resp.Diagnostics.AddAttributeError(
		req.Path,
		"File Parser Validation Error",
		fmt.Sprintf("Attribute %s %s", req.Path, message),
	)
}

func countNonNil(vars ...types.Object) int {
	count := 0
	for _, v := range vars {
		if !v.IsNull() {
			count++
		}
	}
	return count
}
