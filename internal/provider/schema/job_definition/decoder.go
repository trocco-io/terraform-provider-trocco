package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

func DecoderSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional: true,
		Attributes: map[string]schema.Attribute{
			"match_name": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Relative path after decompression (regular expression). If not entered, all data in the compressed file will be transferred.",
			},
		},
	}
}
