package filter

import "github.com/hashicorp/terraform-plugin-framework/types"

type FilterHash struct {
	Name types.String `tfsdk:"name"`
}
