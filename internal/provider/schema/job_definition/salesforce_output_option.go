package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	planModifier "terraform-provider-trocco/internal/provider/planmodifier"
)

func SalesforceOutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of destination Salesforce settings",
		Attributes: map[string]schema.Attribute{
			"object": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Object name",
			},
			"action_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Transfer mode",
				Validators: []validator.String{
					stringvalidator.OneOf("insert", "upsert", "update"),
				},
				Default: stringdefault.StaticString("insert"),
			},
			"api_version": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				MarkdownDescription: "Api version",
				Default:             stringdefault.StaticString("54.0"),
			},
			"upsert_key": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Upsert key. If action_type is 'upsert', this field can be set.",
			},
			"ignore_nulls": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Update processing when NULL is included. Even if true, the record update process itself is performed.",
			},
			"throw_if_failed": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "Status of records that could not be sent",
			},
			"salesforce_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "Salesforce connection ID. Only connection information with authentication method user_password can be selected.",
			},
		},
		PlanModifiers: []planmodifier.Object{
			&planModifier.SalesforceOutputOptionPlanModifier{},
		},
	}
}
