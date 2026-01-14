package job_definition

import (
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

func MongodbInputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of source MongoDB",
		Attributes: map[string]schema.Attribute{
			"database": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Database name",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"collection": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Collection name",
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"query": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Query. Required when incremental_loading_enabled is false.",
			},
			"incremental_loading_enabled": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(false),
				MarkdownDescription: "Transfer method. true: Differential transfer (only incremental data from the previous transfer), false: Transfer using query",
			},
			"incremental_columns": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Column to determine incremental data. The value of the column specified here is saved in 'Last Transferred Record' for each transfer. From the second transfer onwards, only records where the value of 'Column for Determining Incremental Data' is greater than the previous transfer value (= 'Last Transferred Record') will be transferred. To specify multiple columns, separate them with commas.",
			},
			"last_record": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "Last record transferred. During differential updates, data newer than the value specified here will be transferred. If the form is blank, the transfer will start from the beginning. Unless there is a special reason, do not change this value. Duplicate data may occur.",
			},
			"mongodb_connection_id": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of MongoDB connection",
			},
			"input_option_columns": schema.ListNestedAttribute{
				Required:            true,
				MarkdownDescription: "Column information",
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
							MarkdownDescription: "Format for timestamp columns",
						},
						"timezone": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Timezone for timestamp columns",
						},
					},
				},
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
				},
			},
			"custom_variable_settings": CustomVariableSettingsSchema(),
		},
	}
}
