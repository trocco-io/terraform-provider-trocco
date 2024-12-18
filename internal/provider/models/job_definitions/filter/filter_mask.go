package filter

import "github.com/hashicorp/terraform-plugin-framework/types"

type FilterMask struct {
	Name       types.String `tfsdk:"name"`
	MaskType   types.Int32  `tfsdk:"mask_type"`
	Length     types.Int64  `tfsdk:"length"`
	Pattern    types.String `tfsdk:"pattern"`
	StartIndex types.Int64  `tfsdk:"start_index"`
	EndIndex   types.Int64  `tfsdk:"end_index"`
}
