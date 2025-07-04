package pipeline_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NewLabels(ctx context.Context, ens []string, returnsNilIfEmpty bool) types.Set {
	if ens == nil || (returnsNilIfEmpty && len(ens) == 0) {
		return types.SetNull(types.StringType)
	}

	set, diags := types.SetValueFrom(ctx, types.StringType, ens)
	if diags.HasError() {
		return types.SetNull(types.StringType)
	}
	return set
}
