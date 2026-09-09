package notification_destination

import "github.com/hashicorp/terraform-plugin-framework/types"

type HTTPConfig struct {
	Name        types.String   `tfsdk:"name"`
	URL         types.String   `tfsdk:"url"`
	Description types.String   `tfsdk:"description"`
	Headers     []HTTPKeyValue `tfsdk:"headers"`
	QueryParams []HTTPKeyValue `tfsdk:"query_params"`
}

type HTTPKeyValue struct {
	Key     types.String `tfsdk:"key"`
	Value   types.String `tfsdk:"value"`
	Masking types.Bool   `tfsdk:"masking"`
}
