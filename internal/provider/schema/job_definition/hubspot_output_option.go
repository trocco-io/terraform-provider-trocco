package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func HubspotOutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of destination HubSpot settings",
		Attributes: map[string]schema.Attribute{
			"hubspot_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "HubSpot connection ID",
			},
			"object_type": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Object type. Standard objects: `contact`, `company`, `deal`, `product`, `ticket`, `line_item`, `quote`, `subscription`. Engagement objects: `call`, `email`, `meeting`, `note`, `postal_mail`, `task`. Custom objects are also supported",
			},
			"mode": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("merge"),
				Validators: []validator.String{
					stringvalidator.OneOf("merge", "insert"),
				},
				MarkdownDescription: "Transfer mode. `merge`: Upsert (update if exists, insert if not). `insert`: Insert only. Note: For `subscription` object type, mode is always `merge`",
			},
			"upsert_key": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Upsert key. Required when mode is `merge` and object_type is not `subscription`. Not used when mode is `insert` or object_type is `subscription`",
			},
			"number_of_parallels": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(1),
				Validators: []validator.Int64{
					int64validator.Between(1, 10),
				},
				MarkdownDescription: "Number of parallel processes. Must be between 1 and 10",
			},
			"associations": schema.ListNestedAttribute{
				Optional: true,
				Validators: []validator.List{
					listvalidator.SizeAtLeast(0),
				},
				MarkdownDescription: "Association settings. Only available for engagement objects (`call`, `email`, `meeting`, `note`, `postal_mail`, `task`). Supported associations by object type: `call/email/meeting/note/postal_mail`: contact, company, deal, ticket; `task`: contact, company, deal, ticket, quote",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"to_object_type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Target object type for association",
						},
						"from_object_key": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Source key (column name in transfer data)",
						},
						"to_object_key": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Target key (HubSpot property name)",
						},
					},
				},
			},
		},
	}
}
