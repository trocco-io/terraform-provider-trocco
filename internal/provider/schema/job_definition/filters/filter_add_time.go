package filters

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func FilterAddTimeSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Transfer Date Column Setting",
		Attributes: map[string]schema.Attribute{
			"column_name": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Column name",
			},
			"type": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.OneOf("timestamp", "string"),
				},
				MarkdownDescription: "Column type",
			},
			"timestamp_format": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Timestamp format",
			},
			"time_zone": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Time zone",
			},
		},
	}
}
