package workflow

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type TaskDependency struct {
	Source      types.String `tfsdk:"source"`
	Destination types.String `tfsdk:"destination"`
}
