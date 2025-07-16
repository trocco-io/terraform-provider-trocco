package pipeline_definition

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NewLabels(ctx context.Context, ens []string, previousIsNull bool) types.Set {
	if ens == nil {
		return types.SetNull(types.StringType)
	}

	if previousIsNull && len(ens) == 0 {
		return types.SetNull(types.StringType)
	}

	set, diags := types.SetValueFrom(ctx, types.StringType, ens)
	if diags.HasError() {
		return types.SetNull(types.StringType)
	}
	return set
}
