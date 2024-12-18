package models

import "github.com/hashicorp/terraform-plugin-framework/types"

type LabelModel struct {
	ID   types.Int64  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}
