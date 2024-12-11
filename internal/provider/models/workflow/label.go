package workflow

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NewLabels(ens []int64) []types.Int64 {
	if ens == nil {
		return nil
	}

	var mds []types.Int64
	for _, en := range ens {
		mds = append(mds, types.Int64Value(en))
	}

	// If no labels are present, the API returns an empty array but the provider should set `null`.
	if len(mds) == 0 {
		return nil
	}

	return mds
}
