package job_definitions

import "github.com/hashicorp/terraform-plugin-framework/types"

type Decoder struct {
	MatchName types.String `tfsdk:"match_name"`
}
