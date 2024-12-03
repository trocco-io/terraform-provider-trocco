package provider

import (
	"context"
	"fmt"
	"strconv"
	"terraform-provider-trocco/internal/client"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.Resource                = &workflowResource{}
	_ resource.ResourceWithConfigure   = &workflowResource{}
	_ resource.ResourceWithImportState = &workflowResource{}
)

type workflowResourceModel struct {
	ID               types.Int64                           `tfsdk:"id"`
	Name             types.String                          `tfsdk:"name"`
	Description      types.String                          `tfsdk:"description"`
	Tasks            []workflowResourceTaskModel           `tfsdk:"tasks"`
	TaskDependencies []workflowResourceTaskDependencyModel `tfsdk:"task_dependencies"`
}

type workflowResourceTaskModel struct {
	Key            types.String                    `tfsdk:"key"`
	TaskIdentifier types.Int64                     `tfsdk:"task_identifier"`
	Type           types.String                    `tfsdk:"type"`
	Config         workflowResourceTaskConfigModel `tfsdk:"config"`
}

type workflowResourceTaskConfigModel struct {
	ResourceID      types.Int64                                     `tfsdk:"resource_id"`
	Name            types.String                                    `tfsdk:"name"`
	Message         types.String                                    `tfsdk:"message"`
	Query           types.String                                    `tfsdk:"query"`
	Operator        types.String                                    `tfsdk:"operator"`
	QueryResult     types.Int64                                     `tfsdk:"query_result"`
	AcceptsNull     types.Bool                                      `tfsdk:"accepts_null"`
	Warehouse       types.String                                    `tfsdk:"warehouse"`
	Database        types.String                                    `tfsdk:"database"`
	TaskID          types.String                                    `tfsdk:"task_id"`
	CustomVariables []workflowResourceTaskCustomVariableConfigModel `tfsdk:"custom_variables"`

	HTTPMethod        types.String                                      `tfsdk:"http_method"`
	URL               types.String                                      `tfsdk:"url"`
	RequestBody       types.String                                      `tfsdk:"request_body"`
	RequestHeaders    []workflowResourceTaskRequestHeaderConfigModel    `tfsdk:"request_headers"`
	RequestParameters []workflowResourceTaskRequestParameterConfigModel `tfsdk:"request_parameters"`
}

type workflowResourceTaskCustomVariableConfigModel struct {
	Name      types.String `tfsdk:"name"`
	Type      types.String `tfsdk:"type"`
	Value     types.String `tfsdk:"value"`
	Quantity  types.Int64  `tfsdk:"quantity"`
	Unit      types.String `tfsdk:"unit"`
	Direction types.String `tfsdk:"direction"`
	Format    types.String `tfsdk:"format"`
	TimeZone  types.String `tfsdk:"time_zone"`
}

type workflowResourceTaskRequestHeaderConfigModel struct {
	Key     types.String `tfsdk:"key"`
	Value   types.String `tfsdk:"value"`
	Masking types.Bool   `tfsdk:"masking"`
}

type workflowResourceTaskRequestParameterConfigModel struct {
	Key     types.String `tfsdk:"key"`
	Value   types.String `tfsdk:"value"`
	Masking types.Bool   `tfsdk:"masking"`
}

type workflowResourceTaskDependencyModel struct {
	Source      types.String `tfsdk:"source"`
	Destination types.String `tfsdk:"destination"`
}

func (m *workflowResourceModel) ToCreateWorkflowInput() *client.CreateWorkflowInput {
	tasks := []client.WorkflowTaskInput{}
	for _, t := range m.Tasks {
		customVariables := []client.WorkflowTaskCustomVariableConfigInput{}
		for _, v := range t.Config.CustomVariables {
			customVariables = append(customVariables, client.WorkflowTaskCustomVariableConfigInput{
				Name:      v.Name.ValueStringPointer(),
				Type:      v.Type.ValueStringPointer(),
				Value:     v.Value.ValueStringPointer(),
				Quantity:  newNullableFromTerraformInt64(v.Quantity),
				Unit:      v.Unit.ValueStringPointer(),
				Direction: v.Direction.ValueStringPointer(),
				Format:    v.Format.ValueStringPointer(),
				TimeZone:  v.TimeZone.ValueStringPointer(),
			})
		}

		requestHeaders := []client.WorkflowTaskRequestHeaderConfigInput{}
		for _, h := range t.Config.RequestHeaders {
			requestHeaders = append(requestHeaders, client.WorkflowTaskRequestHeaderConfigInput{
				Key:     h.Key.ValueString(),
				Value:   h.Value.ValueString(),
				Masking: newNullableFromTerraformBool(h.Masking),
			})
		}

		requestParameters := []client.WorkflowTaskRequestParameterConfigInput{}
		for _, p := range t.Config.RequestParameters {
			requestParameters = append(requestParameters, client.WorkflowTaskRequestParameterConfigInput{
				Key:     p.Key.ValueString(),
				Value:   p.Value.ValueString(),
				Masking: newNullableFromTerraformBool(p.Masking),
			})
		}

		config := client.WorkflowTaskConfigInput{
			ResourceID:      newNullableFromTerraformInt64(t.Config.ResourceID),
			Name:            t.Config.Name.ValueStringPointer(),
			Message:         t.Config.Message.ValueStringPointer(),
			Query:           t.Config.Query.ValueStringPointer(),
			Operator:        t.Config.Operator.ValueStringPointer(),
			QueryResult:     newNullableFromTerraformInt64(t.Config.QueryResult),
			AcceptsNull:     t.Config.AcceptsNull.ValueBoolPointer(),
			Warehouse:       t.Config.Warehouse.ValueStringPointer(),
			Database:        t.Config.Database.ValueStringPointer(),
			TaskID:          t.Config.TaskID.ValueStringPointer(),
			CustomVarialbes: customVariables,

			HTTPMethod:        t.Config.HTTPMethod.ValueStringPointer(),
			URL:               t.Config.URL.ValueStringPointer(),
			RequestBody:       t.Config.RequestBody.ValueStringPointer(),
			RequestHeaders:    requestHeaders,
			RequestParameters: requestParameters,
		}

		tasks = append(tasks, client.WorkflowTaskInput{
			Key:            t.Key.ValueString(),
			TaskIdentifier: t.TaskIdentifier.ValueInt64(),
			Type:           t.Type.ValueString(),
			Config:         config,
		})
	}

	taskDependencies := []client.WorkflowTaskDependencyInput{}
	for _, d := range m.TaskDependencies {
		taskDependencies = append(taskDependencies, client.WorkflowTaskDependencyInput{
			Source:      d.Source.ValueString(),
			Destination: d.Destination.ValueString(),
		})
	}

	return &client.CreateWorkflowInput{
		Name:             m.Name.ValueString(),
		Description:      m.Description.ValueStringPointer(),
		Tasks:            tasks,
		TaskDependencies: taskDependencies,
	}
}

func (m *workflowResourceModel) ToUpdateWorkflowInput(state *workflowResourceModel) *client.UpdateWorkflowInput {
	stateTaskIdentifiers := map[string]int64{}
	for _, s := range state.Tasks {
		stateTaskIdentifiers[s.Key.ValueString()] = s.TaskIdentifier.ValueInt64()
	}

	tasks := []client.WorkflowTaskInput{}
	for _, t := range m.Tasks {
		identifier := stateTaskIdentifiers[t.Key.ValueString()]

		customVariables := []client.WorkflowTaskCustomVariableConfigInput{}
		for _, v := range t.Config.CustomVariables {
			customVariables = append(customVariables, client.WorkflowTaskCustomVariableConfigInput{
				Name:      v.Name.ValueStringPointer(),
				Type:      v.Type.ValueStringPointer(),
				Value:     v.Value.ValueStringPointer(),
				Quantity:  newNullableFromTerraformInt64(v.Quantity),
				Unit:      v.Unit.ValueStringPointer(),
				Direction: v.Direction.ValueStringPointer(),
				Format:    v.Format.ValueStringPointer(),
				TimeZone:  v.TimeZone.ValueStringPointer(),
			})
		}

		requestHeaders := []client.WorkflowTaskRequestHeaderConfigInput{}
		for _, h := range t.Config.RequestHeaders {
			requestHeaders = append(requestHeaders, client.WorkflowTaskRequestHeaderConfigInput{
				Key:     h.Key.ValueString(),
				Value:   h.Value.ValueString(),
				Masking: newNullableFromTerraformBool(h.Masking),
			})
		}

		requestParameters := []client.WorkflowTaskRequestParameterConfigInput{}
		for _, p := range t.Config.RequestParameters {
			requestParameters = append(requestParameters, client.WorkflowTaskRequestParameterConfigInput{
				Key:     p.Key.ValueString(),
				Value:   p.Value.ValueString(),
				Masking: newNullableFromTerraformBool(p.Masking),
			})
		}

		config := client.WorkflowTaskConfigInput{
			ResourceID:      newNullableFromTerraformInt64(t.Config.ResourceID),
			Name:            t.Config.Name.ValueStringPointer(),
			Message:         t.Config.Message.ValueStringPointer(),
			Query:           t.Config.Query.ValueStringPointer(),
			Operator:        t.Config.Operator.ValueStringPointer(),
			QueryResult:     newNullableFromTerraformInt64(t.Config.QueryResult),
			AcceptsNull:     t.Config.AcceptsNull.ValueBoolPointer(),
			Warehouse:       t.Config.Warehouse.ValueStringPointer(),
			Database:        t.Config.Database.ValueStringPointer(),
			TaskID:          t.Config.TaskID.ValueStringPointer(),
			CustomVarialbes: customVariables,

			HTTPMethod:        t.Config.HTTPMethod.ValueStringPointer(),
			URL:               t.Config.URL.ValueStringPointer(),
			RequestBody:       t.Config.RequestBody.ValueStringPointer(),
			RequestHeaders:    requestHeaders,
			RequestParameters: requestParameters,
		}

		tasks = append(tasks, client.WorkflowTaskInput{
			Key:            t.Key.ValueString(),
			TaskIdentifier: identifier,
			Type:           t.Type.ValueString(),
			Config:         config,
		})
	}

	taskDependencies := []client.WorkflowTaskDependencyInput{}
	for _, d := range m.TaskDependencies {
		taskDependencies = append(taskDependencies, client.WorkflowTaskDependencyInput{
			Source:      d.Source.ValueString(),
			Destination: d.Destination.ValueString(),
		})
	}

	return &client.UpdateWorkflowInput{
		Name:             m.Name.ValueStringPointer(),
		Description:      m.Description.ValueStringPointer(),
		Tasks:            tasks,
		TaskDependencies: taskDependencies,
	}
}

type workflowResource struct {
	client *client.TroccoClient
}

func NewWorkflowResource() resource.Resource {
	return &workflowResource{}
}

func (r *workflowResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = fmt.Sprintf("%s_workflow", req.ProviderTypeName)
}

func (r *workflowResource) Configure(
	ctx context.Context,
	req resource.ConfigureRequest,
	resp *resource.ConfigureResponse,
) {
	if req.ProviderData == nil {
		return
	}

	c, ok := req.ProviderData.(*client.TroccoClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.client = c
}

func (r *workflowResource) Schema(
	ctx context.Context,
	req resource.SchemaRequest,
	resp *resource.SchemaResponse,
) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO workflow resource.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				MarkdownDescription: "The ID of the workflow.",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the workflow.",
				Required:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtMost(255),
				},
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the workflow.",
				Optional:            true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
			},
			"tasks": schema.SetNestedAttribute{
				MarkdownDescription: "The tasks of the workflow.",
				Optional:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"key": schema.StringAttribute{
							Required: true,
						},
						"task_identifier": schema.Int64Attribute{
							Optional: true,
							Computed: true,
						},
						"type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"trocco_transfer",
									"slack_notify",
									"bigquery_data_check",
									"snowflake_data_check",
									"redshift_data_check",
									"tableau_extract",
									"http_request",
								),
							},
						},
						"config": schema.SingleNestedAttribute{
							Required: true,
							Attributes: map[string]schema.Attribute{
								"resource_id": schema.Int64Attribute{
									Optional: true,
								},
								"name": schema.StringAttribute{
									Optional: true,
								},
								"message": schema.StringAttribute{
									Optional: true,
								},
								"query": schema.StringAttribute{
									Optional: true,
								},
								"operator": schema.StringAttribute{
									Optional: true,
								},
								"query_result": schema.Int64Attribute{
									Optional: true,
								},
								"accepts_null": schema.BoolAttribute{
									Optional: true,
								},
								"warehouse": schema.StringAttribute{
									Optional: true,
								},
								"database": schema.StringAttribute{
									Optional: true,
								},
								"task_id": schema.StringAttribute{
									Optional: true,
								},
								"custom_variables": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"name": schema.StringAttribute{
												Required: true,
											},
											"type": schema.StringAttribute{
												Required: true,
											},
											"value": schema.StringAttribute{
												Optional: true,
											},
											"quantity": schema.Int64Attribute{
												Optional: true,
											},
											"unit": schema.StringAttribute{
												Optional: true,
											},
											"direction": schema.StringAttribute{
												Optional: true,
											},
											"format": schema.StringAttribute{
												Optional: true,
											},
											"time_zone": schema.StringAttribute{
												Optional: true,
											},
										},
									},
								},
								"http_method": schema.StringAttribute{
									Optional: true,
								},
								"url": schema.StringAttribute{
									Optional: true,
								},
								"request_body": schema.StringAttribute{
									Optional: true,
								},
								"request_headers": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Required: true,
											},
											"value": schema.StringAttribute{
												Required: true,
											},
											"masking": schema.BoolAttribute{
												Required: true,
											},
										},
									},
								},
								"request_parameters": schema.ListNestedAttribute{
									Optional: true,
									NestedObject: schema.NestedAttributeObject{
										Attributes: map[string]schema.Attribute{
											"key": schema.StringAttribute{
												Required: true,
											},
											"value": schema.StringAttribute{
												Required: true,
											},
											"masking": schema.BoolAttribute{
												Required: true,
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"task_dependencies": schema.SetNestedAttribute{
				MarkdownDescription: "The task dependencies of the workflow.",
				Required:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"source": schema.StringAttribute{
							Required: true,
						},
						"destination": schema.StringAttribute{
							Required: true,
						},
					},
				},
			},
		},
	}
}

func (r *workflowResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	plan := &workflowResourceModel{}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workflow, err := r.client.CreateWorkflow(
		plan.ToCreateWorkflowInput(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating workflow",
			fmt.Sprintf("Unable to create workflow, got error: %s", err),
		)
		return
	}

	tasks := []workflowResourceTaskModel{}
	for _, t := range workflow.Tasks {
		customVariables := []workflowResourceTaskCustomVariableConfigModel{}
		for _, v := range t.Config.CustomVariables {
			customVariables = append(customVariables, workflowResourceTaskCustomVariableConfigModel{
				Name:      types.StringPointerValue(v.Name),
				Type:      types.StringPointerValue(v.Type),
				Value:     types.StringPointerValue(v.Value),
				Quantity:  types.Int64PointerValue(v.Quantity),
				Unit:      types.StringPointerValue(v.Unit),
				Direction: types.StringPointerValue(v.Direction),
				Format:    types.StringPointerValue(v.Format),
				TimeZone:  types.StringPointerValue(v.TimeZone),
			})
		}
		if len(customVariables) == 0 {
			customVariables = nil
		}

		requestHeaders := []workflowResourceTaskRequestHeaderConfigModel{}
		for _, h := range t.Config.RequestHeaders {
			requestHeaders = append(requestHeaders, workflowResourceTaskRequestHeaderConfigModel{
				Key:     types.StringValue(h.Key),
				Value:   types.StringValue(h.Value),
				Masking: types.BoolValue(h.Masking),
			})
		}
		if len(requestHeaders) == 0 {
			requestHeaders = nil
		}

		requestParameters := []workflowResourceTaskRequestParameterConfigModel{}
		for _, p := range t.Config.RequestParameters {
			requestParameters = append(requestParameters, workflowResourceTaskRequestParameterConfigModel{
				Key:     types.StringValue(p.Key),
				Value:   types.StringValue(p.Value),
				Masking: types.BoolValue(p.Masking),
			})
		}
		if len(requestParameters) == 0 {
			requestParameters = nil
		}

		tasks = append(tasks, workflowResourceTaskModel{
			Key:            types.StringValue(t.Key),
			TaskIdentifier: types.Int64Value(t.TaskIdentifier),
			Type:           types.StringValue(t.Type),
			Config: workflowResourceTaskConfigModel{
				ResourceID:      types.Int64PointerValue(t.Config.ResourceID),
				Name:            types.StringPointerValue(t.Config.Name),
				Message:         types.StringPointerValue(t.Config.Message),
				Query:           types.StringPointerValue(t.Config.Query),
				Operator:        types.StringPointerValue(t.Config.Operator),
				QueryResult:     types.Int64PointerValue(t.Config.QueryResult),
				AcceptsNull:     types.BoolPointerValue(t.Config.AcceptsNull),
				Warehouse:       types.StringPointerValue(t.Config.Warehouse),
				Database:        types.StringPointerValue(t.Config.Database),
				TaskID:          types.StringPointerValue(t.Config.TaskID),
				CustomVariables: customVariables,

				HTTPMethod:        types.StringPointerValue(t.Config.HTTPMethod),
				URL:               types.StringPointerValue(t.Config.URL),
				RequestBody:       types.StringPointerValue(t.Config.RequestBody),
				RequestHeaders:    requestHeaders,
				RequestParameters: requestParameters,
			},
		})
	}

	keys := map[int64]types.String{}
	for _, t := range workflow.Tasks {
		keys[t.TaskIdentifier] = types.StringValue(t.Key)
	}

	taskDependencies := []workflowResourceTaskDependencyModel{}
	for _, d := range workflow.TaskDependencies {
		taskDependencies = append(taskDependencies, workflowResourceTaskDependencyModel{
			Source:      keys[d.Source],
			Destination: keys[d.Destination],
		})
	}

	newState := workflowResourceModel{
		ID:               types.Int64Value(workflow.ID),
		Name:             types.StringPointerValue(workflow.Name),
		Description:      types.StringPointerValue(workflow.Description),
		Tasks:            tasks,
		TaskDependencies: taskDependencies,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *workflowResource) Update(
	ctx context.Context,
	req resource.UpdateRequest,
	resp *resource.UpdateResponse,
) {
	state := &workflowResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan := &workflowResourceModel{}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workflow, err := r.client.UpdateWorkflow(
		state.ID.ValueInt64(),
		plan.ToUpdateWorkflowInput(state),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Updating workflow",
			fmt.Sprintf("Unable to update workflow, got error: %s", err),
		)
		return
	}

	tasks := []workflowResourceTaskModel{}
	for _, t := range workflow.Tasks {
		customVariables := []workflowResourceTaskCustomVariableConfigModel{}
		for _, v := range t.Config.CustomVariables {
			customVariables = append(customVariables, workflowResourceTaskCustomVariableConfigModel{
				Name:      types.StringPointerValue(v.Name),
				Type:      types.StringPointerValue(v.Type),
				Value:     types.StringPointerValue(v.Value),
				Quantity:  types.Int64PointerValue(v.Quantity),
				Unit:      types.StringPointerValue(v.Unit),
				Direction: types.StringPointerValue(v.Direction),
				Format:    types.StringPointerValue(v.Format),
				TimeZone:  types.StringPointerValue(v.TimeZone),
			})
		}
		if len(customVariables) == 0 {
			customVariables = nil
		}

		requestHeaders := []workflowResourceTaskRequestHeaderConfigModel{}
		for _, h := range t.Config.RequestHeaders {
			requestHeaders = append(requestHeaders, workflowResourceTaskRequestHeaderConfigModel{
				Key:     types.StringValue(h.Key),
				Value:   types.StringValue(h.Value),
				Masking: types.BoolValue(h.Masking),
			})
		}
		if len(requestHeaders) == 0 {
			requestHeaders = nil
		}

		requestParameters := []workflowResourceTaskRequestParameterConfigModel{}
		for _, p := range t.Config.RequestParameters {
			requestParameters = append(requestParameters, workflowResourceTaskRequestParameterConfigModel{
				Key:     types.StringValue(p.Key),
				Value:   types.StringValue(p.Value),
				Masking: types.BoolValue(p.Masking),
			})
		}
		if len(requestParameters) == 0 {
			requestParameters = nil
		}

		tasks = append(tasks, workflowResourceTaskModel{
			Key:            types.StringValue(t.Key),
			TaskIdentifier: types.Int64Value(t.TaskIdentifier),
			Type:           types.StringValue(t.Type),
			Config: workflowResourceTaskConfigModel{
				ResourceID:      types.Int64PointerValue(t.Config.ResourceID),
				Name:            types.StringPointerValue(t.Config.Name),
				Message:         types.StringPointerValue(t.Config.Message),
				Query:           types.StringPointerValue(t.Config.Query),
				Operator:        types.StringPointerValue(t.Config.Operator),
				QueryResult:     types.Int64PointerValue(t.Config.QueryResult),
				AcceptsNull:     types.BoolPointerValue(t.Config.AcceptsNull),
				Warehouse:       types.StringPointerValue(t.Config.Warehouse),
				Database:        types.StringPointerValue(t.Config.Database),
				TaskID:          types.StringPointerValue(t.Config.TaskID),
				CustomVariables: customVariables,

				HTTPMethod:        types.StringPointerValue(t.Config.HTTPMethod),
				URL:               types.StringPointerValue(t.Config.URL),
				RequestBody:       types.StringPointerValue(t.Config.RequestBody),
				RequestHeaders:    requestHeaders,
				RequestParameters: requestParameters,
			},
		})
	}

	keys := map[int64]types.String{}
	for _, t := range workflow.Tasks {
		keys[t.TaskIdentifier] = types.StringValue(t.Key)
	}

	taskDependencies := []workflowResourceTaskDependencyModel{}
	for _, d := range workflow.TaskDependencies {
		taskDependencies = append(taskDependencies, workflowResourceTaskDependencyModel{
			Source:      keys[d.Source],
			Destination: keys[d.Destination],
		})
	}

	newState := workflowResourceModel{
		ID:               types.Int64Value(workflow.ID),
		Name:             types.StringPointerValue(workflow.Name),
		Description:      types.StringPointerValue(workflow.Description),
		Tasks:            tasks,
		TaskDependencies: taskDependencies,
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *workflowResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	state := &workflowResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	workflow, err := r.client.GetWorkflow(
		state.ID.ValueInt64(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading workflow",
			fmt.Sprintf("Unable to read workflow, got error: %s", err),
		)
		return
	}

	stateKeys := map[int64]string{}
	for _, s := range state.Tasks {
		stateKeys[s.TaskIdentifier.ValueInt64()] = s.Key.ValueString()
	}

	tasks := []workflowResourceTaskModel{}
	for _, t := range workflow.Tasks {
		key := stateKeys[t.TaskIdentifier]

		customVariables := []workflowResourceTaskCustomVariableConfigModel{}
		for _, v := range t.Config.CustomVariables {
			customVariables = append(customVariables, workflowResourceTaskCustomVariableConfigModel{
				Name:      types.StringPointerValue(v.Name),
				Type:      types.StringPointerValue(v.Type),
				Value:     types.StringPointerValue(v.Value),
				Quantity:  types.Int64PointerValue(v.Quantity),
				Unit:      types.StringPointerValue(v.Unit),
				Direction: types.StringPointerValue(v.Direction),
				Format:    types.StringPointerValue(v.Format),
				TimeZone:  types.StringPointerValue(v.TimeZone),
			})
		}
		if len(customVariables) == 0 {
			customVariables = nil
		}

		requestHeaders := []workflowResourceTaskRequestHeaderConfigModel{}
		for _, h := range t.Config.RequestHeaders {
			requestHeaders = append(requestHeaders, workflowResourceTaskRequestHeaderConfigModel{
				Key:     types.StringValue(h.Key),
				Value:   types.StringValue(h.Value),
				Masking: types.BoolValue(h.Masking),
			})
		}
		if len(requestHeaders) == 0 {
			requestHeaders = nil
		}

		requestParameters := []workflowResourceTaskRequestParameterConfigModel{}
		for _, p := range t.Config.RequestParameters {
			requestParameters = append(requestParameters, workflowResourceTaskRequestParameterConfigModel{
				Key:     types.StringValue(p.Key),
				Value:   types.StringValue(p.Value),
				Masking: types.BoolValue(p.Masking),
			})
		}
		if len(requestParameters) == 0 {
			requestParameters = nil
		}

		tasks = append(tasks, workflowResourceTaskModel{
			Key:            types.StringValue(key),
			TaskIdentifier: types.Int64Value(t.TaskIdentifier),
			Type:           types.StringValue(t.Type),
			Config: workflowResourceTaskConfigModel{
				ResourceID:      types.Int64PointerValue(t.Config.ResourceID),
				Name:            types.StringPointerValue(t.Config.Name),
				Message:         types.StringPointerValue(t.Config.Message),
				Query:           types.StringPointerValue(t.Config.Query),
				Operator:        types.StringPointerValue(t.Config.Operator),
				QueryResult:     types.Int64PointerValue(t.Config.QueryResult),
				AcceptsNull:     types.BoolPointerValue(t.Config.AcceptsNull),
				Warehouse:       types.StringPointerValue(t.Config.Warehouse),
				Database:        types.StringPointerValue(t.Config.Database),
				TaskID:          types.StringPointerValue(t.Config.TaskID),
				CustomVariables: customVariables,

				HTTPMethod:        types.StringPointerValue(t.Config.HTTPMethod),
				URL:               types.StringPointerValue(t.Config.URL),
				RequestBody:       types.StringPointerValue(t.Config.RequestBody),
				RequestHeaders:    requestHeaders,
				RequestParameters: requestParameters,
			},
		})
	}

	newState := workflowResourceModel{
		ID:          types.Int64Value(workflow.ID),
		Name:        types.StringPointerValue(workflow.Name),
		Description: types.StringPointerValue(workflow.Description),
		Tasks:       tasks,
		// create/update のときは string
		// read のときは int64
		TaskDependencies: state.TaskDependencies,
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *workflowResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	s := &workflowResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, s)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.DeleteWorkflow(
		s.ID.ValueInt64(),
	); err != nil {
		resp.Diagnostics.AddError(
			"Deleting workflow",
			fmt.Sprintf("Unable to delete workflow, got error: %s", err),
		)
		return
	}
}

func (r *workflowResource) ImportState(
	ctx context.Context,
	req resource.ImportStateRequest,
	resp *resource.ImportStateResponse,
) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing workflow",
			fmt.Sprintf("Unable to parse id, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
