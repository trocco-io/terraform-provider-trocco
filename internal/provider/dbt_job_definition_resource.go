package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/client/parameter"
	"terraform-provider-trocco/internal/provider/model"
	troccoPlanModifier "terraform-provider-trocco/internal/provider/planmodifier"
	troccoValidator "terraform-provider-trocco/internal/provider/validator"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &dbtJobDefinitionResource{}
	_ resource.ResourceWithConfigure   = &dbtJobDefinitionResource{}
	_ resource.ResourceWithImportState = &dbtJobDefinitionResource{}
)

func NewDbtJobDefinitionResource() resource.Resource {
	return &dbtJobDefinitionResource{}
}

type dbtJobDefinitionResource struct {
	client *client.TroccoClient
}

func (r *dbtJobDefinitionResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_dbt_job_definition"
}

func (r *dbtJobDefinitionResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*client.TroccoClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.TroccoClient, got: %T", req.ProviderData),
		)
		return
	}
	r.client = c
}

func (r *dbtJobDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO dbt job definition resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The ID of the dbt job definition.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
					stringvalidator.UTF8LengthAtMost(255),
				},
				MarkdownDescription: "The name of the dbt job definition.",
			},
			"description": schema.StringAttribute{
				Optional:            true,
				MarkdownDescription: "The description of the dbt job definition.",
			},
			"resource_group_id": schema.Int64Attribute{
				Optional:            true,
				MarkdownDescription: "The ID of the resource group that the dbt job definition belongs to.",
			},
			"dbt_git_repository_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "The ID of the dbt Git repository to associate with this job definition. Changing it to a repository with a different adapter type is rejected by the server.",
			},
			"threads": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Validators: []validator.Int64{
					int64validator.Between(1, 16),
				},
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Number of dbt threads (1-16). When omitted, the server applies its default.",
			},
			"target": schema.StringAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "dbt profile target name. When omitted, the server applies its default.",
			},
			"bigquery_setting": schema.SingleNestedAttribute{
				Optional: true,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(
						path.MatchRelative().AtParent().AtName("snowflake_setting"),
						path.MatchRelative().AtParent().AtName("redshift_setting"),
					),
				},
				Attributes: map[string]schema.Attribute{
					"connection_id": schema.Int64Attribute{
						Required:            true,
						MarkdownDescription: "The ID of the BigQuery connection.",
					},
					"dataset": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "BigQuery dataset.",
					},
					"location": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "BigQuery location.",
					},
				},
				MarkdownDescription: "BigQuery adapter setting. Exactly one of `bigquery_setting` / `snowflake_setting` / `redshift_setting` must be set, matching the adapter type of the linked dbt Git repository.",
			},
			"snowflake_setting": schema.SingleNestedAttribute{
				Optional: true,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(
						path.MatchRelative().AtParent().AtName("bigquery_setting"),
						path.MatchRelative().AtParent().AtName("redshift_setting"),
					),
				},
				Attributes: map[string]schema.Attribute{
					"connection_id": schema.Int64Attribute{
						Required:            true,
						MarkdownDescription: "The ID of the Snowflake connection.",
					},
					"warehouse": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "Snowflake warehouse name.",
					},
					"database": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "Snowflake database name.",
					},
					"schema": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "Snowflake schema name.",
					},
					"role": schema.StringAttribute{
						Optional:            true,
						MarkdownDescription: "Snowflake role name.",
					},
				},
				MarkdownDescription: "Snowflake adapter setting.",
			},
			"redshift_setting": schema.SingleNestedAttribute{
				Optional: true,
				Validators: []validator.Object{
					objectvalidator.ConflictsWith(
						path.MatchRelative().AtParent().AtName("bigquery_setting"),
						path.MatchRelative().AtParent().AtName("snowflake_setting"),
					),
				},
				Attributes: map[string]schema.Attribute{
					"connection_id": schema.Int64Attribute{
						Required:            true,
						MarkdownDescription: "The ID of the Redshift connection.",
					},
					"database": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "Redshift database name.",
					},
					"schema": schema.StringAttribute{
						Required:            true,
						MarkdownDescription: "Redshift schema name.",
					},
				},
				MarkdownDescription: "Redshift adapter setting.",
			},
			"commands": schema.ListNestedAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"command": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
							MarkdownDescription: "The dbt subcommand (e.g. `run`, `test`, `build`). Refer to the TROCCO documentation for the currently supported values.",
						},
						"value": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Argument for the subcommand (e.g. model selector for `run`).",
						},
						"options": schema.ListNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"key": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "Option key (e.g. `--vars`).",
									},
									"value": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Option value.",
									},
								},
							},
							MarkdownDescription: "Command options.",
						},
					},
				},
				MarkdownDescription: "Ordered list of dbt commands to run. Set to `[]` to clear.",
			},
			"custom_variable_settings": schema.ListNestedAttribute{
				Optional: true,
				Computed: true,
				PlanModifiers: []planmodifier.List{
					listplanmodifier.UseStateForUnknown(),
				},
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								troccoValidator.WrappingDollarValidator{},
							},
							MarkdownDescription: "Custom variable name. It must start and end with `$`",
						},
						"type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("string", "timestamp", "timestamp_runtime"),
							},
							MarkdownDescription: "Custom variable type.",
						},
						"value": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Fixed string. Required when `type` is `string`.",
						},
						"quantity": schema.Int64Attribute{
							Optional: true,
							Validators: []validator.Int64{
								int64validator.AtLeast(0),
							},
							MarkdownDescription: "Quantity. Required when `type` is timestamp/timestamp_runtime.",
						},
						"unit": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOf("hour", "date", "month"),
							},
							MarkdownDescription: "Time unit. Required when `type` is timestamp/timestamp_runtime.",
						},
						"direction": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOf("ago", "later"),
							},
							MarkdownDescription: "Direction. Required when `type` is timestamp/timestamp_runtime.",
						},
						"format": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Format. Required when `type` is timestamp/timestamp_runtime.",
						},
						"time_zone": schema.StringAttribute{
							Optional:            true,
							MarkdownDescription: "Time zone. Required when `type` is timestamp/timestamp_runtime.",
						},
					},
					PlanModifiers: []planmodifier.Object{
						&troccoPlanModifier.CustomVariableSettingPlanModifier{},
					},
				},
				MarkdownDescription: "Custom variable settings. Set to `[]` to clear.",
			},
		},
	}
}

func (r *dbtJobDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan model.DbtJobDefinitionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	commands, diags := extractDbtCommandInputs(ctx, plan.Commands)
	resp.Diagnostics.Append(diags...)
	cvSettings, diags := extractDbtCustomVariableSettingInputs(ctx, plan.CustomVariableSettings)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := parameter.CreateDbtJobDefinitionInput{
		Name:                   plan.Name.ValueString(),
		Description:            model.NewNullableString(plan.Description),
		ResourceGroupID:        model.NewNullableInt64(plan.ResourceGroupID),
		DbtGitRepositoryID:     plan.DbtGitRepositoryID.ValueInt64(),
		Commands:               commands,
		CustomVariableSettings: cvSettings,
	}
	if !plan.Threads.IsNull() && !plan.Threads.IsUnknown() {
		input.SetThreads(plan.Threads.ValueInt64())
	}
	if !plan.Target.IsNull() && !plan.Target.IsUnknown() {
		input.SetTarget(plan.Target.ValueString())
	}
	if plan.BigquerySetting != nil {
		input.SetBigquerySetting(buildDbtBigquerySettingInput(plan.BigquerySetting))
	}
	if plan.SnowflakeSetting != nil {
		input.SetSnowflakeSetting(buildDbtSnowflakeSettingInput(plan.SnowflakeSetting))
	}
	if plan.RedshiftSetting != nil {
		input.SetRedshiftSetting(buildDbtRedshiftSettingInput(plan.RedshiftSetting))
	}

	def, err := r.client.CreateDbtJobDefinition(&input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating dbt job definition",
			fmt.Sprintf("Unable to create dbt job definition, got error: %s", err),
		)
		return
	}

	state, err := model.NewDbtJobDefinitionModel(ctx, def)
	if err != nil {
		resp.Diagnostics.AddError("Converting dbt job definition", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *dbtJobDefinitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var prior model.DbtJobDefinitionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	if resp.Diagnostics.HasError() {
		return
	}

	def, err := r.client.GetDbtJobDefinition(prior.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading dbt job definition",
			fmt.Sprintf("Unable to read dbt job definition, got error: %s", err),
		)
		return
	}

	state, err := model.NewDbtJobDefinitionModel(ctx, def)
	if err != nil {
		resp.Diagnostics.AddError("Converting dbt job definition", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *dbtJobDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, prior model.DbtJobDefinitionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	if resp.Diagnostics.HasError() {
		return
	}

	commands, diags := extractDbtCommandInputs(ctx, plan.Commands)
	resp.Diagnostics.Append(diags...)
	cvSettings, diags := extractDbtCustomVariableSettingInputs(ctx, plan.CustomVariableSettings)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := parameter.UpdateDbtJobDefinitionInput{
		Description:            model.NewNullableString(plan.Description),
		ResourceGroupID:        model.NewNullableInt64(plan.ResourceGroupID),
		Commands:               commands,
		CustomVariableSettings: cvSettings,
	}
	if !plan.Name.IsNull() && !plan.Name.IsUnknown() {
		input.SetName(plan.Name.ValueString())
	}
	if !plan.DbtGitRepositoryID.IsNull() && !plan.DbtGitRepositoryID.IsUnknown() {
		input.SetDbtGitRepositoryID(plan.DbtGitRepositoryID.ValueInt64())
	}
	if !plan.Threads.IsNull() && !plan.Threads.IsUnknown() {
		input.SetThreads(plan.Threads.ValueInt64())
	}
	if !plan.Target.IsNull() && !plan.Target.IsUnknown() {
		input.SetTarget(plan.Target.ValueString())
	}
	if plan.BigquerySetting != nil {
		input.SetBigquerySetting(buildDbtBigquerySettingInput(plan.BigquerySetting))
	}
	if plan.SnowflakeSetting != nil {
		input.SetSnowflakeSetting(buildDbtSnowflakeSettingInput(plan.SnowflakeSetting))
	}
	if plan.RedshiftSetting != nil {
		input.SetRedshiftSetting(buildDbtRedshiftSettingInput(plan.RedshiftSetting))
	}

	def, err := r.client.UpdateDbtJobDefinition(prior.ID.ValueInt64(), &input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating dbt job definition",
			fmt.Sprintf("Unable to update dbt job definition, got error: %s", err),
		)
		return
	}

	state, err := model.NewDbtJobDefinitionModel(ctx, def)
	if err != nil {
		resp.Diagnostics.AddError("Converting dbt job definition", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func buildDbtBigquerySettingInput(s *model.DbtBigquerySettingModel) parameter.DbtBigquerySettingInput {
	return parameter.DbtBigquerySettingInput{
		ConnectionID: s.ConnectionID.ValueInt64(),
		Dataset:      s.Dataset.ValueString(),
		Location:     model.NewNullableString(s.Location),
	}
}

func buildDbtSnowflakeSettingInput(s *model.DbtSnowflakeSettingModel) parameter.DbtSnowflakeSettingInput {
	return parameter.DbtSnowflakeSettingInput{
		ConnectionID: s.ConnectionID.ValueInt64(),
		Warehouse:    s.Warehouse.ValueString(),
		Database:     s.Database.ValueString(),
		Schema:       s.Schema.ValueString(),
		Role:         model.NewNullableString(s.Role),
	}
}

func buildDbtRedshiftSettingInput(s *model.DbtRedshiftSettingModel) parameter.DbtRedshiftSettingInput {
	return parameter.DbtRedshiftSettingInput{
		ConnectionID: s.ConnectionID.ValueInt64(),
		Database:     s.Database.ValueString(),
		Schema:       s.Schema.ValueString(),
	}
}

// extractDbtCommandInputs converts the plan's types.List into a request payload.
// Null/Unknown plans (e.g. first apply with the attribute omitted) are mapped to
// an empty slice so the JSON payload always carries `commands: []`.
func extractDbtCommandInputs(ctx context.Context, list types.List) ([]parameter.DbtCommandInput, diag.Diagnostics) {
	if list.IsNull() || list.IsUnknown() {
		return []parameter.DbtCommandInput{}, nil
	}
	var commands []model.DbtCommandModel
	diags := list.ElementsAs(ctx, &commands, false)
	if diags.HasError() {
		return nil, diags
	}
	out := make([]parameter.DbtCommandInput, 0, len(commands))
	for _, c := range commands {
		cmd := parameter.DbtCommandInput{
			Command: c.Command.ValueString(),
		}
		if !c.Value.IsNull() && !c.Value.IsUnknown() {
			cmd.SetValue(c.Value.ValueString())
		}
		if len(c.Options) > 0 {
			opts := make([]parameter.DbtCommandOptionInput, 0, len(c.Options))
			for _, opt := range c.Options {
				o := parameter.DbtCommandOptionInput{
					Key: opt.Key.ValueString(),
				}
				if !opt.Value.IsNull() && !opt.Value.IsUnknown() {
					o.SetValue(opt.Value.ValueString())
				}
				opts = append(opts, o)
			}
			cmd.SetOptions(opts)
		}
		out = append(out, cmd)
	}
	return out, nil
}

func extractDbtCustomVariableSettingInputs(ctx context.Context, list types.List) ([]parameter.CustomVariableSettingInput, diag.Diagnostics) {
	if list.IsNull() || list.IsUnknown() {
		return []parameter.CustomVariableSettingInput{}, nil
	}
	var settings []model.CustomVariableSetting
	diags := list.ElementsAs(ctx, &settings, false)
	if diags.HasError() {
		return nil, diags
	}
	inputs := model.ToCustomVariableSettingInputs(&settings)
	if inputs == nil {
		return []parameter.CustomVariableSettingInput{}, nil
	}
	return *inputs, nil
}

func (r *dbtJobDefinitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state model.DbtJobDefinitionModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.DeleteDbtJobDefinition(state.ID.ValueInt64()); err != nil {
		resp.Diagnostics.AddError(
			"Deleting dbt job definition",
			fmt.Sprintf("Unable to delete dbt job definition, got error: %s", err),
		)
		return
	}
}

func (r *dbtJobDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing dbt job definition",
			fmt.Sprintf("Unable to parse id, got error: %s", err),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
