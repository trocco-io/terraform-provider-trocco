package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/model"
	jobdefschema "terraform-provider-trocco/internal/provider/schema/job_definition"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/objectvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

const (
	dbtJobDefinitionDefaultThreads = 1
	dbtJobDefinitionDefaultTarget  = "trocco_default"
)

var dbtCommandValues = []string{
	"build",
	"deps",
	"run",
	"run-operation",
	"seed",
	"snapshot",
	"source-freshness",
	"test",
}

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
			"adapter_type": schema.StringAttribute{
				Computed:            true,
				MarkdownDescription: "The adapter type derived from the linked dbt Git repository (one of `bigquery`, `snowflake`, `redshift`).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"dbt_git_repository_id": schema.Int64Attribute{
				Required:            true,
				MarkdownDescription: "The ID of the dbt Git repository to associate with this job definition. Changing it to a repository with a different adapter type is rejected by the server.",
			},
			"threads": schema.Int64Attribute{
				Optional: true,
				Computed: true,
				Default:  int64default.StaticInt64(dbtJobDefinitionDefaultThreads),
				Validators: []validator.Int64{
					int64validator.Between(1, 16),
				},
				MarkdownDescription: "Number of dbt threads (1-16). Defaults to `1`.",
			},
			"target": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Default:             stringdefault.StaticString(dbtJobDefinitionDefaultTarget),
				MarkdownDescription: "dbt profile target name. Defaults to `trocco_default`.",
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
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"command": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf(dbtCommandValues...),
							},
							MarkdownDescription: fmt.Sprintf("The dbt subcommand. One of %v.", dbtCommandValues),
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
				MarkdownDescription: "Ordered list of dbt commands to run.",
			},
			"custom_variable_settings": jobdefschema.CustomVariableSettingsSchema(),
		},
	}
}

func (r *dbtJobDefinitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan model.DbtJobDefinitionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := plan.ToCreateInput()
	def, err := r.client.CreateDbtJobDefinition(&input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating dbt job definition",
			fmt.Sprintf("Unable to create dbt job definition, got error: %s", err),
		)
		return
	}

	state := model.NewDbtJobDefinitionModel(def)
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

	state := model.NewDbtJobDefinitionModel(def)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *dbtJobDefinitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, prior model.DbtJobDefinitionModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	if resp.Diagnostics.HasError() {
		return
	}

	input := plan.ToUpdateInput()
	def, err := r.client.UpdateDbtJobDefinition(prior.ID.ValueInt64(), &input)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating dbt job definition",
			fmt.Sprintf("Unable to update dbt job definition, got error: %s", err),
		)
		return
	}

	state := model.NewDbtJobDefinitionModel(def)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
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
