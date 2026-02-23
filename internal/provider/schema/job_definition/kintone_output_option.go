package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func KintoneOutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of destination Kintone settings",
		Attributes: map[string]schema.Attribute{
			"kintone_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "Kintone connection ID",
			},
			"app_id": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Kintone app ID",
			},
			"guest_space_id": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Guest space ID",
			},
			"mode": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.OneOf("insert", "update", "upsert"),
				},
				Default:             stringdefault.StaticString("insert"),
				MarkdownDescription: "Transfer mode. One of `insert`, `update`, `upsert`",
			},
			"update_key": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Update key (only applicable if mode is 'update' or 'upsert')",
			},
			"ignore_nulls": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Whether to ignore NULL values",
			},
			"reduce_key": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Reduce key for deduplication",
			},
			"chunk_size": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				Default:             int64default.StaticInt64(100),
				MarkdownDescription: "Chunk size",
			},
			"kintone_output_option_column_options": schema.ListNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Column name",
						},
						"field_code": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Field code",
						},
						"type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"SINGLE_LINE_TEXT",
									"MULTI_LINE_TEXT",
									"RICH_TEXT",
									"NUMBER",
									"CHECK_BOX",
									"RADIO_BUTTON",
									"MULTI_SELECT",
									"DROP_DOWN",
									"USER_SELECT",
									"ORGANIZATION_SELECT",
									"GROUP_SELECT",
									"DATE",
									"TIME",
									"DATETIME",
									"LINK",
									"SUBTABLE",
								),
							},
							MarkdownDescription: "Field type",
						},
						"timezone": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Timezone (only applicable if type is 'DATE', 'TIME' or 'DATETIME')",
						},
						"sort_column": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Sort column (only applicable if type is 'SUBTABLE')",
						},
					},
				},
			},
		},
	}
}
