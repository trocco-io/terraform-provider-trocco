package workflow

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func NewLabels(ens []string, returnsNilIfEmpty bool) []types.String {
	if ens == nil {
		return nil
	}

	if returnsNilIfEmpty && len(ens) == 0 {
		return nil
	}

	mds := []types.String{}
	for _, en := range ens {
		mds = append(mds, types.StringValue(en))
	}

	return mds
}
