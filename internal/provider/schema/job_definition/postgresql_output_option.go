package job_definition

import (
	planModifier "terraform-provider-trocco/internal/provider/planmodifier"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func PostgresqlOutputOptionSchema() schema.Attribute {
	return schema.SingleNestedAttribute{
		Optional:            true,
		MarkdownDescription: "Attributes of destination PostgreSQL settings",
		Attributes: map[string]schema.Attribute{
			"database": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Database name",
			},
			"schema": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Schema name",
			},
			"table": schema.StringAttribute{
				Required:            true,
				MarkdownDescription: "Table name",
			},
			"mode": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("insert"),
				Validators: []validator.String{
					stringvalidator.OneOf("insert", "insert_direct", "truncate_insert", "replace", "merge"),
				},
				MarkdownDescription: "Transfer mode",
			},
			"default_time_zone": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString("UTC"),
				MarkdownDescription: "Default time zone",
			},
			"postgresql_connection_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "PostgreSQL connection ID",
			},
			"postgresql_output_option_merge_keys": schema.SetAttribute{
				Optional:            true,
				ElementType:         types.StringType,
				MarkdownDescription: "Merge keys (only applicable if mode is 'merge')",
			},
		},
		PlanModifiers: []planmodifier.Object{
			&planModifier.PostgresqlOutputOptionPlanModifier{},
		},
	}
}
