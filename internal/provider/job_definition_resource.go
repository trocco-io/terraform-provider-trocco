package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/int32validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strconv"
	"terraform-provider-trocco/internal/client"
	job_definitions3 "terraform-provider-trocco/internal/client/entities/job_definitions"
	"terraform-provider-trocco/internal/client/parameters"
	job_definitions2 "terraform-provider-trocco/internal/client/parameters/job_definitions"
	filter2 "terraform-provider-trocco/internal/client/parameters/job_definitions/filter"
	"terraform-provider-trocco/internal/provider/models"
	"terraform-provider-trocco/internal/provider/models/job_definitions"
	"terraform-provider-trocco/internal/provider/models/job_definitions/filter"
	"terraform-provider-trocco/internal/provider/models/job_definitions/input_options"
)

var (
	_ resource.Resource                = &jobDefinitionResource{}
	_ resource.ResourceWithConfigure   = &jobDefinitionResource{}
	_ resource.ResourceWithImportState = &jobDefinitionResource{}
)

func NewJobDefinitionResource() resource.Resource {
	return &jobDefinitionResource{}
}

type jobDefinitionResource struct {
	client *client.TroccoClient
}

func (r *jobDefinitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Importing job definition",
			fmt.Sprintf("Unable to parse id, got error: %s", err),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (r *jobDefinitionResource) Metadata(
	ctx context.Context,
	req resource.MetadataRequest,
	resp *resource.MetadataResponse,
) {
	resp.TypeName = fmt.Sprintf("%s_job_definition", req.ProviderTypeName)
}

func (r *jobDefinitionResource) Configure(
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

func (r *jobDefinitionResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Provides a TROCCO job definitions.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:            true,
				MarkdownDescription: "The ID of the job definition",
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtMost(255),
				},
				MarkdownDescription: "Name of the job definition. It must be less than 256 characters",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Description of the job definition. It must be at least 1 character",
			},
			"resource_group_id": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
				MarkdownDescription: "ID of the resource group to which the job definition belongs",
			},
			"resource_enhancement": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.OneOf("medium", "custom_spec", "large", "xlarge"),
				},
				MarkdownDescription: "Resource size to be used when executing the job. If not specified, the resource size specified in the transfer settings is applied. The value that can be specified varies depending on the connector. (This parameter is available only in the Professional plan.",
			},
			"is_runnable_concurrently": schema.BoolAttribute{
				Required:            true,
				MarkdownDescription: "Specifies whether or not to run a job if another job with the same job definition is running at the time the job is run",
			},
			"retry_limit": schema.Int64Attribute{
				Required: true,
				Validators: []validator.Int64{
					int64validator.Between(0, 10),
				},
				MarkdownDescription: "Maximum number of retries. if set 0, the job will not be retried",
			},
			"input_option_type": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Input option type.",
			},
			"input_option": schema.MapNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"mysql_input_option": schema.MapNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"database": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.UTF8LengthAtLeast(1),
										},
									},
									"table": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.UTF8LengthAtLeast(1),
										},
									},
									"query": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.UTF8LengthAtLeast(1),
										},
									},
									"incremental_columns": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.UTF8LengthAtLeast(1),
										},
									},
									"last_record": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.UTF8LengthAtLeast(1),
										},
									},
									"incremental_loading_enabled": schema.BoolAttribute{
										Required: true,
									},
									"fetch_rows": schema.Int64Attribute{
										Required: true,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
									"connect_timeout": schema.Int64Attribute{
										Required: true,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
									"socket_timeout": schema.Int64Attribute{
										Required: true,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
									"default_time_zone": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.UTF8LengthAtLeast(1),
										},
									},
									"use_legacy_datetime_code": schema.BoolAttribute{
										Optional: true,
									},
									"mysql_connection_id": schema.Int64Attribute{
										Required: true,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
									"input_option_columns": schema.ListNestedAttribute{
										Required: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Required: true,
													Validators: []validator.String{
														stringvalidator.UTF8LengthAtLeast(1),
													},
												},
												"type": schema.StringAttribute{
													Required: true,
													Validators: []validator.String{
														stringvalidator.OneOf("boolean", "long", "timestamp", "double", "string", "json"),
													},
												},
											},
										},
										Validators: []validator.List{
											listvalidator.SizeAtLeast(1),
										},
									},
									"custom_variable_settings": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Required: true,
													Validators: []validator.String{
														wrappingDollarValidator{},
													},
													MarkdownDescription: "Custom variable name. It must start and end with `$`",
												},
												"type": schema.StringAttribute{
													Required: true,
													Validators: []validator.String{
														stringvalidator.OneOf("string", "timestamp", "timestamp_runtime"),
													},
													MarkdownDescription: "Custom variable type. The following types are supported: `string`, `timestamp`, `timestamp_runtime`",
												},
												"value": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Fixed string which will replace variables at runtime. Required in `string` type",
												},
												"quantity": schema.Int32Attribute{
													Optional: true,
													Validators: []validator.Int32{
														int32validator.AtLeast(0),
													},
													MarkdownDescription: "Quantity used to calculate diff from context_time. Required in `timestamp` and `timestamp_runtime` types",
												},
												"unit": schema.StringAttribute{
													Optional: true,
													Validators: []validator.String{
														stringvalidator.OneOf("hour", "date", "month"),
													},
													MarkdownDescription: "Time unit used to calculate diff from context_time. The following units are supported: `hour`, `date`, `month`. Required in `timestamp` and `timestamp_runtime` types",
												},
												"direction": schema.StringAttribute{
													Optional: true,
													Validators: []validator.String{
														stringvalidator.OneOf("ago", "later"),
													},
													MarkdownDescription: "Direction of the diff from context_time. The following directions are supported: `ago`, `later`. Required in `timestamp` and `timestamp_runtime` types",
												},
												"format": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Format used to replace variables. Required in `timestamp` and `timestamp_runtime` types",
												},
												"time_zone": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Time zone used to format the timestamp. Required in `timestamp` and `timestamp_runtime` types",
												},
											},
											PlanModifiers: []planmodifier.Object{
												&customVariableSettingPlanModifier{},
											},
										},
									},
								},
							},
						},
						// TODO: GCS
					},
				},
			},
			"output_option_type": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				MarkdownDescription: "Output option type.",
			},
			"output_option": schema.MapNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"bigquery_output_option": schema.MapNestedAttribute{
							Optional: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"dataset": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.UTF8LengthAtLeast(1),
										},
									},
									"table": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.UTF8LengthAtLeast(1),
										},
									},
									"mode": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf("append", "append_direct", "replace", "delete_in_advance", "merge"),
										},
									},
									"auto_create_dataset": schema.BoolAttribute{
										Required: true,
									},
									"auto_create_table": schema.BoolAttribute{
										Required: true,
									},
									"open_timeout_sec": schema.Int64Attribute{
										Required: true,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
									"timeout_sec": schema.Int64Attribute{
										Required: true,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
									"send_timeout_sec": schema.Int64Attribute{
										Required: true,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
									"read_timeout_sec": schema.Int64Attribute{
										Required: true,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
									"retries": schema.Int64Attribute{
										Required: true,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
									"partitioning_type": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf("ingestion_time", "time_unit_column"),
										},
									},
									"time_partitioning_type": schema.StringAttribute{
										Optional: true,
										Validators: []validator.String{
											stringvalidator.OneOf("HOUR", "DAY", "MONTH", "YEAR"),
										},
									},
									"time_partitioning_field": schema.StringAttribute{
										Optional: true,
									},
									"time_partitioning_expiration_ms": schema.Int64Attribute{
										Optional: true,
									},
									"time_partitioning_require_partition_filter": schema.BoolAttribute{
										Optional: true,
									},
									"location": schema.StringAttribute{
										Optional: true,
									},
									"template_table": schema.StringAttribute{
										Optional: true,
									},
									"bigquery_connection_id": schema.Int64Attribute{
										Required: true,
										Validators: []validator.Int64{
											int64validator.AtLeast(1),
										},
									},
									"before_load": schema.StringAttribute{
										Optional: true,
									},
									"bigquery_output_option_column_options": schema.ListNestedAttribute{
										Optional: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Required: true,
												},
												"type": schema.StringAttribute{
													Required: true,
												},
												"mode": schema.StringAttribute{
													Required: true,
												},
												"timestamp_format": schema.StringAttribute{
													Optional: true,
												},
												"timezone": schema.StringAttribute{
													Optional: true,
												},
												"description": schema.StringAttribute{
													Optional: true,
												},
											},
										},
									},
									"bigquery_output_option_clustering_fields": schema.ListAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
									"bigquery_output_option_merge_keys": schema.ListAttribute{
										Optional:    true,
										ElementType: types.StringType,
									},
								},
							},
						},
					},
				},
			},
			"schedules": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"frequency": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("hourly", "daily", "weekly", "monthly"),
							},
							MarkdownDescription: "Frequency of automatic execution. The following frequencies are supported: `hourly`, `daily`, `weekly`, `monthly`",
						},
						"minute": schema.Int32Attribute{
							Required: true,
							Validators: []validator.Int32{
								int32validator.Between(0, 59),
							},
							MarkdownDescription: "Value of minute. Required for all schedules",
						},
						"hour": schema.Int32Attribute{
							Optional: true,
							Validators: []validator.Int32{
								int32validator.Between(0, 23),
							},
							MarkdownDescription: "Value of hour. Required in `daily`, `weekly`, and `monthly` schedules",
						},
						"day_of_week": schema.Int32Attribute{
							Optional: true,
							Validators: []validator.Int32{
								int32validator.Between(0, 6),
							},
							MarkdownDescription: "Value of day of week. Sunday - Saturday is represented as 0 - 6. Required in `weekly` schedule",
						},
						"day": schema.Int32Attribute{
							Optional: true,
							Validators: []validator.Int32{
								int32validator.Between(1, 31),
							},
							MarkdownDescription: "Value of day. Required in `monthly` schedule",
						},
						"time_zone": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "Time zone to be used for calculation",
						},
					},
					PlanModifiers: []planmodifier.Object{
						&schedulePlanModifier{},
					},
				},
				MarkdownDescription: "Schedules to be attached to the job definition",
			},
			"notifications": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"destination_type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("slack", "email"),
							},
							MarkdownDescription: "Destination service where the notification will be sent. The following types are supported: `slack`, `email`",
						},
						"slack_channel_id": schema.Int64Attribute{
							Optional: true,
							Validators: []validator.Int64{
								int64validator.AtLeast(1),
							},
							MarkdownDescription: "ID of the slack channel used to send notifications. Required when `destination_type` is `slack`",
						},
						"email_id": schema.Int64Attribute{
							Optional: true,
							Validators: []validator.Int64{
								int64validator.AtLeast(1),
							},
							MarkdownDescription: "ID of the email used to send notifications. Required when `destination_type` is `email`",
						},
						"notification_type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("job", "record"),
							},
							MarkdownDescription: "Category of condition. The following types are supported: `job`, `record`",
						},
						"notify_when": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOf("finished", "failed"),
							},
							MarkdownDescription: "Specifies the job status that trigger a notification. The following types are supported: `finished`, `failed`. Required when `notification_type` is `job`",
						},
						"record_count": schema.Int64Attribute{
							Optional:            true,
							MarkdownDescription: "The number of records to be used for condition. Required when `notification_type` is `record`",
						},
						"record_operator": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOf("above", "below"),
							},
							MarkdownDescription: "Operator to be used for condition. The following operators are supported: `above`, `below`. Required when `notification_type` is `record`",
						},
						"record_type": schema.StringAttribute{
							Optional: true,
							Validators: []validator.String{
								stringvalidator.OneOf("transfer", "skipped"),
							},
							MarkdownDescription: "Condition for number of records to be notified",
						},
						"message": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "The message to be sent with the notification",
						},
					},
				},
				MarkdownDescription: "Notifications to be attached to the job definition",
			},
			"labels": schema.SetNestedAttribute{
				Optional: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:            true,
							MarkdownDescription: "The ID of the label",
						},
						"name": schema.StringAttribute{
							Required:            true,
							MarkdownDescription: "The name of the label",
						},
					},
				},
				MarkdownDescription: "Labels to be attached to the job definition",
			},
		},
	}
}

type jobDefinitionResourceModel struct {
	ID                        types.Int64                                  `tfsdk:"id"`
	Name                      types.String                                 `tfsdk:"name"`
	Description               types.String                                 `tfsdk:"description"`
	ResourceGroupID           types.Int64                                  `tfsdk:"resource_group_id"`
	IsRunnableConcurrently    types.Bool                                   `tfsdk:"is_runnable_concurrently"`
	RetryLimit                types.Int64                                  `tfsdk:"retry_limit"`
	ResourceEnhancement       types.String                                 `tfsdk:"resource_enhancement"`
	InputOptionType           types.String                                 `tfsdk:"input_option_type"`
	InputOption               job_definitions.InputOption                  `tfsdk:"input_option"`
	OutputOptionType          types.String                                 `tfsdk:"output_option_type"`
	OutputOption              job_definitions.OutputOption                 `tfsdk:"output_option"`
	FilterColumns             []filter.FilterColumn                        `tfsdk:"filter_columns"`
	FilterRows                *filter.FilterRows                           `tfsdk:"filter_rows"`
	FilterMasks               []filter.FilterMask                          `tfsdk:"filter_masks"`
	FilterAddTime             *filter.FilterAddTime                        `tfsdk:"filter_add_time"`
	FilterGsub                []filter.FilterGsub                          `tfsdk:"filter_gsub"`
	FilterStringTransforms    []filter.FilterStringTransform               `tfsdk:"filter_string_transforms"`
	FilterHashes              []filter.FilterHash                          `tfsdk:"filter_hashes"`
	FilterUnixTimeConversions []filter.FilterUnixTimeConversion            `tfsdk:"filter_unixtime_conversions"`
	Notifications             *[]job_definitions.JobDefinitionNotification `tfsdk:"notifications"`
	Schedules                 *[]models.Schedule                           `tfsdk:"schedules"`
	Labels                    *[]models.LabelModel                         `tfsdk:"labels"`
}

func (model *jobDefinitionResourceModel) ToCreateJobDefinitionInput() *client.CreateJobDefinitionInput {
	var labels []string
	if model.Labels != nil {
		for _, l := range *model.Labels {
			labels = append(labels, l.Name.ValueString())
		}
	}
	var notifications []job_definitions2.JobDefinitionNotificationInput
	if model.Notifications != nil {
		for _, n := range *model.Notifications {
			notifications = append(notifications, n.ToInput())
		}
	}
	var schedules []parameters.ScheduleInput
	if model.Schedules != nil {
		for _, s := range *model.Schedules {
			schedules = append(schedules, s.ToInput())
		}
	}

	var filterColumns []filter2.FilterColumnInput
	if model.FilterColumns != nil {
		for _, f := range model.FilterColumns {
			filterColumns = append(filterColumns, f.ToInput())
		}
	}

	var filterMasks []filter2.FilterMaskInput
	for _, f := range model.FilterMasks {
		filterMasks = append(filterMasks, f.ToInput())
	}

	var filterGsub []filter2.FilterGsubInput
	for _, f := range model.FilterGsub {
		filterGsub = append(filterGsub, f.ToInput())
	}

	var filterStringTransforms []filter2.FilterStringTransformInput
	for _, f := range model.FilterStringTransforms {
		filterStringTransforms = append(filterStringTransforms, f.ToInput())
	}

	var filterHashes []filter2.FilterHashInput
	for _, f := range model.FilterHashes {
		filterHashes = append(filterHashes, f.ToInput())
	}

	var filterUnixTimeconversions []filter2.FilterUnixTimeConversionInput
	for _, f := range model.FilterUnixTimeConversions {
		filterUnixTimeconversions = append(filterUnixTimeconversions, f.ToInput())
	}

	return &client.CreateJobDefinitionInput{
		Name:                      model.Name.ValueString(),
		Description:               model.Description.ValueStringPointer(),
		ResourceGroupID:           newNullableFromTerraformInt64(model.ResourceGroupID),
		IsRunnableConcurrently:    model.IsRunnableConcurrently.ValueBoolPointer(),
		RetryLimit:                model.RetryLimit.ValueInt64(),
		ResourceEnhancement:       model.ResourceEnhancement.ValueStringPointer(),
		FilterColumns:             filterColumns,
		FilterRows:                model.FilterRows.ToInput(),
		FilterMasks:               filterMasks,
		FilterAddTime:             model.FilterAddTime.ToInput(),
		FilterGsub:                filterGsub,
		FilterStringTransforms:    filterStringTransforms,
		FilterHashes:              filterHashes,
		FilterUnixTimeConversions: filterUnixTimeconversions,
		InputOptionType:           model.InputOptionType.ValueString(),
		InputOption:               model.InputOption.ToInput(),
		OutputOptionType:          model.OutputOptionType.ValueString(),
		OutputOption:              model.OutputOption.ToInput(),
		Labels:                    labels,
		Schedules:                 schedules,
		Notifications:             notifications,
	}
}

func (r *jobDefinitionResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	//TODO implement me
	panic("implement me")
}

func (r *jobDefinitionResource) Create(
	ctx context.Context,
	req resource.CreateRequest,
	resp *resource.CreateResponse,
) {
	plan := &jobDefinitionResourceModel{}
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	jobDefinition, err := r.client.CreateJobDefinition(
		plan.ToCreateJobDefinitionInput(),
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Creating job definition",
			fmt.Sprintf("Unable to create job definition, got error: %s", err),
		)
		return
	}

	var mysqlInputOptionColumns []input_options.InputOptionColumn
	for _, m := range jobDefinition.InputOption.MySQLInputOption.InputOptionColumns {
		mysqlInputOptionColumns = append(mysqlInputOptionColumns, input_options.InputOptionColumn{
			Name: types.StringValue(m.Name),
			Type: types.StringValue(m.Type),
		})
	}

	newState := jobDefinitionResourceModel{
		ID:                        types.Int64Value(jobDefinition.ID),
		Name:                      types.StringValue(jobDefinition.Name),
		Description:               types.StringPointerValue(jobDefinition.Description),
		ResourceGroupID:           types.Int64PointerValue(jobDefinition.ResourceGroupID),
		IsRunnableConcurrently:    types.BoolPointerValue(jobDefinition.IsRunnableConcurrently),
		RetryLimit:                types.Int64Value(jobDefinition.RetryLimit),
		ResourceEnhancement:       types.StringPointerValue(jobDefinition.ResourceEnhancement),
		InputOptionType:           types.StringValue(jobDefinition.InputOptionType),
		InputOption:               job_definitions.NewInputOption(jobDefinition.InputOption),
		OutputOptionType:          types.StringValue(jobDefinition.OutputOptionType),
		OutputOption:              job_definitions.NewOutputOption(jobDefinition.OutputOption),
		FilterColumns:             filter.NewFilterColumns(jobDefinition.FilterColumns),
		FilterRows:                filter.NewFilterRows(jobDefinition.FilterRows),
		FilterMasks:               filter.NewFilterMasks(jobDefinition.FilterMasks),
		FilterAddTime:             filter.NewFilterAddTime(jobDefinition.FilterAddTime),
		FilterGsub:                filter.NewFilterGsub(jobDefinition.FilterGsub),
		FilterStringTransforms:    filter.NewFilterStringTransforms(jobDefinition.FilterStringTransforms),
		FilterHashes:              filter.NewFilterHashes(jobDefinition.FilterHashes),
		FilterUnixTimeConversions: filter.NewFilterUnixTimeConversions(jobDefinition.FilterUnixTimeConversions),
		Notifications:             job_definitions.NewJobDefinitionNotifications(jobDefinition.Notifications),
		Schedules:                 models.NewSchedules(jobDefinition.Schedules),
		Labels:                    models.NewLabels(jobDefinition.Labels),
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func toJobDefinitionNotificationModel(notifications *[]job_definitions3.JobDefinitionNotification) *[]job_definitions.JobDefinitionNotification {
	if notifications == nil {
		return nil
	}
	outputs := make([]job_definitions.JobDefinitionNotification, 0, len(*notifications))
	for _, input := range *notifications {
		outputs = append(outputs, job_definitions.JobDefinitionNotification{
			DestinationType:  types.StringValue(input.DestinationType),
			SlackChannelID:   types.Int64PointerValue(input.SlackChannelID),
			EmailID:          types.Int64PointerValue(input.EmailID),
			NotificationType: types.StringValue(input.NotificationType),
			NotifyWhen:       types.StringPointerValue(input.NotifyWhen),
			Message:          types.StringValue(input.Message),
			RecordCount:      types.Int64PointerValue(input.RecordCount),
			RecordOperator:   types.StringPointerValue(input.RecordOperator),
			RecordType:       types.StringPointerValue(input.RecordType),
			Minutes:          types.Int64PointerValue(input.Minutes),
		})
	}
	return &outputs
}

func (r *jobDefinitionResource) Read(
	ctx context.Context,
	req resource.ReadRequest,
	resp *resource.ReadResponse,
) {
	state := &jobDefinitionResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	jobDefinition, err := r.client.GetJobDefinition(state.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Reading job definition",
			fmt.Sprintf("Unable to read job definition, got error: %s", err),
		)
		return
	}

	newState := jobDefinitionResourceModel{
		ID:                        types.Int64Value(jobDefinition.ID),
		Name:                      types.StringValue(jobDefinition.Name),
		Description:               types.StringPointerValue(jobDefinition.Description),
		ResourceGroupID:           types.Int64PointerValue(jobDefinition.ResourceGroupID),
		IsRunnableConcurrently:    types.BoolPointerValue(jobDefinition.IsRunnableConcurrently),
		RetryLimit:                types.Int64Value(jobDefinition.RetryLimit),
		ResourceEnhancement:       types.StringPointerValue(jobDefinition.ResourceEnhancement),
		InputOptionType:           types.StringValue(jobDefinition.InputOptionType),
		InputOption:               job_definitions.NewInputOption(jobDefinition.InputOption),
		OutputOptionType:          types.StringValue(jobDefinition.OutputOptionType),
		OutputOption:              job_definitions.NewOutputOption(jobDefinition.OutputOption),
		FilterColumns:             filter.NewFilterColumns(jobDefinition.FilterColumns),
		FilterRows:                filter.NewFilterRows(jobDefinition.FilterRows),
		FilterMasks:               filter.NewFilterMasks(jobDefinition.FilterMasks),
		FilterAddTime:             filter.NewFilterAddTime(jobDefinition.FilterAddTime),
		FilterGsub:                filter.NewFilterGsub(jobDefinition.FilterGsub),
		FilterStringTransforms:    filter.NewFilterStringTransforms(jobDefinition.FilterStringTransforms),
		FilterHashes:              filter.NewFilterHashes(jobDefinition.FilterHashes),
		FilterUnixTimeConversions: filter.NewFilterUnixTimeConversions(jobDefinition.FilterUnixTimeConversions),
		Notifications:             job_definitions.NewJobDefinitionNotifications(jobDefinition.Notifications),
		Schedules:                 models.NewSchedules(jobDefinition.Schedules),
		Labels:                    models.NewLabels(jobDefinition.Labels),
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, newState)...)
}

func (r *jobDefinitionResource) Delete(
	ctx context.Context,
	req resource.DeleteRequest,
	resp *resource.DeleteResponse,
) {
	s := &jobDefinitionResourceModel{}
	resp.Diagnostics.Append(req.State.Get(ctx, s)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if err := r.client.DeleteJobDefinition(s.ID.ValueInt64()); err != nil {
		resp.Diagnostics.AddError(
			"Deleting job definition",
			fmt.Sprintf("Unable to delete job definition, got error: %s", err),
		)
		return
	}
}