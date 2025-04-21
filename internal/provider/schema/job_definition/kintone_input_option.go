package job_definition

import (
	planmodifier2 "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func KintoneInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of source kintone",
		Attributes: map[string]schema.Attribute{
			"app_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "app id",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"guest_space_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "guest space id",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "If you want to use all record loading, specify it.",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"expand_subtable": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				MarkdownDescription: "If enabled and the target Kintone app contains a table, data will be transferred by table row instead of per record.",
			},
			"kintone_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of kintone connection",
			},
			"input_option_columns": schema.ListNestedAttribute{
				Required:            true,
				MarkdownDescription: "List of columns to be retrieved and their types",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
							MarkdownDescription: "Column name",
						},
						"type": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column type",
							Validators: []validator.String{
								stringvalidator.OneOf("boolean", "long", "timestamp", "double", "string", "json"),
							},
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Column format",
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
		PlanModifiers: []planmodifier.Object{
			&planmodifier2.KintoneInputOptionPlanModifier{},
		},
	}
}
