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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"strconv"
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/client/parameters"
	job_definitions2 "terraform-provider-trocco/internal/client/parameters/job_definitions"
	filter2 "terraform-provider-trocco/internal/client/parameters/job_definitions/filter"
	"terraform-provider-trocco/internal/provider/models"
	"terraform-provider-trocco/internal/provider/models/job_definitions"
	"terraform-provider-trocco/internal/provider/models/job_definitions/filter"
	planmodifier2 "terraform-provider-trocco/internal/provider/planmodifier"
	validator2 "terraform-provider-trocco/internal/provider/validator"
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
	resp.TypeName = req.ProviderTypeName + "_job_definition"
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
				Computed: true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
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
				Optional:            true,
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
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				MarkdownDescription: "Resource size to be used when executing the job. If not specified, the resource size specified in the transfer settings is applied. The value that can be specified varies depending on the connector. (This parameter is available only in the Professional plan.",
				Validators: []validator.String{
					stringvalidator.OneOf("medium", "custom_spec", "large", "xlarge"),
				},
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
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Input option type.",
			},
			"input_option": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"mysql_input_option": schema.SingleNestedAttribute{
						Optional:            true,
						MarkdownDescription: "Attributes of source mysql",
						Attributes: map[string]schema.Attribute{
							"database": schema.StringAttribute{
								Required:            true,
								MarkdownDescription: "database name",
								Validators: []validator.String{
									stringvalidator.UTF8LengthAtLeast(1),
								},
							},
							"table": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "table name. If you want to use incremental loading, specify it.",
							},
							"query": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "If you want to use all record loading, specify it.",
							},
							"incremental_columns": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "Columns to determine incremental data",
							},
							"last_record": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "Last record transferred. The value of the column specified here is stored in “Last Transferred Record” for each transfer, and for the second and subsequent transfers, only records for which the value of the “Column for Determining Incremental Data” is greater than the value of the previous transfer (= “Last Transferred Record”) are transferred. If you wish to specify multiple columns, specify them separated by commas. If not specified, the primary key is used.",
							},
							"incremental_loading_enabled": schema.BoolAttribute{
								Required:            true,
								MarkdownDescription: "If it is true, to be incremental loading. If it is false, to be all record loading",
							},
							"fetch_rows": schema.Int64Attribute{
								Required:            true,
								MarkdownDescription: "Number of records processed by the cursor at one time",
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},
							"connect_timeout": schema.Int64Attribute{
								Required: true,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
								MarkdownDescription: "Connection timeout (sec)",
							},
							"socket_timeout": schema.Int64Attribute{
								Required: true,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
								MarkdownDescription: "Socket timeout (seconds)",
							},
							"default_time_zone": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "Default time zone. enter the server-side time zone setting for MySQL. If the time zone is set to Japan, enter “Asia/Tokyo”.",
							},
							"use_legacy_datetime_code": schema.BoolAttribute{
								Optional:            true,
								MarkdownDescription: "Legacy time code setting. setting the useLegacyDatetimeCode option in the JDBC driver",
							},
							"mysql_connection_id": schema.Int64Attribute{
								Required: true,
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
								MarkdownDescription: "ID of MySQL connection",
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
												validator2.WrappingDollarValidator{},
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
										&planmodifier2.CustomVariableSettingPlanModifier{},
									},
								},
							},
						},
						PlanModifiers: []planmodifier.Object{
							&MysqlInputOptionPlanModifier{},
						},
					},
					"gcs_input_option": schema.SingleNestedAttribute{
						Optional:            true,
						MarkdownDescription: "Attributes about source GCS",
						Attributes: map[string]schema.Attribute{
							"bucket": schema.StringAttribute{
								Required:            true,
								MarkdownDescription: "Bucket name",
								Validators: []validator.String{
									stringvalidator.UTF8LengthAtLeast(1),
								},
							},
							"path_prefix": schema.StringAttribute{
								Required: true,
								Validators: []validator.String{
									stringvalidator.UTF8LengthAtLeast(1),
								},
								MarkdownDescription: "Path prefix",
							},
							"incremental_loading_enabled": schema.BoolAttribute{
								Required:            true,
								MarkdownDescription: "If it is true, to be incremental loading. If it is false, to be all record loading",
							},
							"last_path": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "Last path transferred. It is only enabled when incremental loading is true. When updating differences, data behind in lexicographic order from the path specified here is transferred. If the form is blank, the data is transferred from the beginning. Do not change this value unless there is a special reason. Duplicate data may occur.",
							},
							"stop_when_file_not_found": schema.BoolAttribute{
								Required:            true,
								MarkdownDescription: "Flag whether the transfer should continue if the file does not exist in the specified path",
							},
							"gcs_connection_id": schema.Int64Attribute{
								Required:            true,
								MarkdownDescription: "Id of GCS connection",
								Validators: []validator.Int64{
									int64validator.AtLeast(1),
								},
							},
							"decompression_type": schema.StringAttribute{
								Optional:            true,
								MarkdownDescription: "Decompression type",
								Validators: []validator.String{
									stringvalidator.OneOf("gzip", "bzip2", "zip", "targz"),
								},
							},
							"parquet_parser": schema.SingleNestedAttribute{
								Optional:            true,
								MarkdownDescription: "For files in parquet format, this parameter is required.",
								Attributes: map[string]schema.Attribute{
									"columns": schema.ListNestedAttribute{
										Required:            true,
										MarkdownDescription: "List of columns to be retrieved and their types",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Required:            true,
													MarkdownDescription: "Column name",
												},
												"type": schema.StringAttribute{
													Required:            true,
													Validators:          []validator.String{stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean", "json")},
													MarkdownDescription: "Column type",
												},
												"format": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Format of the column.",
												},
											},
										},
									},
								},
							},
							"jsonpath_parser": schema.SingleNestedAttribute{
								Optional:            true,
								MarkdownDescription: "For files in jsonpath format, this parameter is required.",
								Attributes: map[string]schema.Attribute{
									"root": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "JSONPath",
									},
									"default_time_zone": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "Default time zone",
									},
									"columns": schema.ListNestedAttribute{
										Required: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Required:            true,
													MarkdownDescription: "Column name",
												},
												"type": schema.StringAttribute{
													Required:            true,
													Validators:          []validator.String{stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean", "json")},
													MarkdownDescription: "Column type",
												},
												"time_zone": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "time zone",
												},
												"format": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Format of the column.",
												},
											},
										},
									},
								},
							},
							"xml_parser": schema.SingleNestedAttribute{
								Optional:            true,
								MarkdownDescription: "For files in xml format, this parameter is required.",
								Attributes: map[string]schema.Attribute{
									"root": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "Root element",
									},
									"columns": schema.ListNestedAttribute{
										Required: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Required:            true,
													MarkdownDescription: "Column name",
												},
												"type": schema.StringAttribute{
													Required:            true,
													Validators:          []validator.String{stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean", "json")},
													MarkdownDescription: "Column type",
												},
												"path": schema.StringAttribute{
													Required:            true,
													MarkdownDescription: "XPath",
												},
												"timezone": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "time zone",
												},
												"format": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Format of the column.",
												},
											},
										},
									},
								},
							},
							"excel_parser": schema.SingleNestedAttribute{
								MarkdownDescription: "For files in excel format, this parameter is required.",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"default_time_zone": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "Default time zone",
									},
									"sheet_name": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "Sheet name",
									},
									"skip_header_lines": schema.Int64Attribute{
										Required:            true,
										MarkdownDescription: "Number of header lines to skip",
									},
									"columns": schema.ListNestedAttribute{
										MarkdownDescription: "List of columns to be retrieved and their types",
										Required:            true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Required:            true,
													MarkdownDescription: "Column name",
												},
												"type": schema.StringAttribute{
													Required:            true,
													Validators:          []validator.String{stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean", "json")},
													MarkdownDescription: "Column type",
												},
												"format": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Format of the column.",
												},
												"formula_handling": schema.StringAttribute{
													Required:            true,
													Validators:          []validator.String{stringvalidator.OneOf("cashed_value", "evaluate")},
													MarkdownDescription: "Formula handling",
												},
											},
										},
									},
								},
							},
							"ltsv_parser": schema.SingleNestedAttribute{
								MarkdownDescription: "For files in LTSV format, this parameter is required.",
								Optional:            true,
								Attributes: map[string]schema.Attribute{
									"newline": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Newline character",
									},
									"charset": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Character set",
									},
									"columns": schema.ListNestedAttribute{
										Required:            true,
										MarkdownDescription: "List of columns to be retrieved and their types",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Required:            true,
													MarkdownDescription: "Column name",
												},
												"type": schema.StringAttribute{
													Required:            true,
													Validators:          []validator.String{stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean")},
													MarkdownDescription: "Column type",
												},
												"format": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Format of the column.",
												},
											},
										},
									},
								},
							},
							"jsonl_parser": schema.SingleNestedAttribute{
								Optional:            true,
								MarkdownDescription: "For files in JSONL format, this parameter is required",
								Attributes: map[string]schema.Attribute{
									"stop_on_invalid_record": schema.BoolAttribute{
										Required:            true,
										MarkdownDescription: "Flag whether the transfer should stop if an invalid record is found",
									},
									"default_time_zone": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "Default time zone",
									},
									"newline": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Newline character",
									},
									"charset": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Character set",
									},
									"columns": schema.ListNestedAttribute{
										Required:            true,
										MarkdownDescription: "List of columns to be retrieved and their types",
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Required:            true,
													MarkdownDescription: "Column name",
												},
												"type": schema.StringAttribute{
													Required:            true,
													MarkdownDescription: "Column type",
													Validators:          []validator.String{stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean")},
												},
												"time_zone": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "time zone",
												},
												"format": schema.StringAttribute{
													Optional:            true,
													MarkdownDescription: "Format of the column",
												},
											},
										},
									},
								},
							},
							"csv_parser": schema.SingleNestedAttribute{
								Optional:            true,
								MarkdownDescription: "For files in CSV format, this parameter is required",
								Attributes: map[string]schema.Attribute{
									"delimiter": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "Delimiter",
									},
									"quote": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Quote character",
									},
									"escape": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Escape character",
									},
									"skip_header_lines": schema.Int64Attribute{
										Required:            true,
										MarkdownDescription: "Number of header lines to skip",
									},
									"null_string_enabled": schema.BoolAttribute{
										Required:            true,
										MarkdownDescription: "Flag whether or not to set the string to be replaced by NULL",
									},
									"null_string": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Replacement source string to be converted to NULL",
									},
									"trim_if_not_quoted": schema.BoolAttribute{
										Required:            true,
										MarkdownDescription: "Flag whether or not to remove spaces from the value if it is not quoted",
									},
									"quotes_in_quoted_fields": schema.StringAttribute{
										Required:            true,
										MarkdownDescription: "Processing method for irregular quarts",
										Validators:          []validator.String{stringvalidator.OneOf("ACCEPT_ONLY_RFC4180_ESCAPED", "ACCEPT_STRAY_QUOTES_ASSUMING_NO_DELIMITERS_IN_FIELDS")},
									},
									"comment_line_marker": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Comment line marker. Skip if this character is at the beginning of a line",
									},
									"allow_optional_columns": schema.BoolAttribute{
										Required:            true,
										MarkdownDescription: "If true, NULL-complete the missing columns. If false, treat as invalid record.",
									},
									"allow_extra_columns": schema.BoolAttribute{
										Required:            true,
										MarkdownDescription: "If true, ignore the column. If false, treat as invalid record.",
									},
									"max_quoted_size_limit": schema.Int64Attribute{
										Required:            true,
										MarkdownDescription: "Maximum amount of data that can be enclosed in quotation marks.",
									},
									"stop_on_invalid_record": schema.BoolAttribute{
										Required:            true,
										MarkdownDescription: "Flag whether or not to abort the transfer if an invalid record is found.",
									},
									"default_time_zone": schema.StringAttribute{
										Required: true,
									},
									"default_date": schema.StringAttribute{
										Required: true,
									},
									"newline": schema.StringAttribute{
										Required: true,
									},
									"charset": schema.StringAttribute{
										Optional: true,
									},
									"columns": schema.ListNestedAttribute{
										Required: true,
										NestedObject: schema.NestedAttributeObject{
											Attributes: map[string]schema.Attribute{
												"name": schema.StringAttribute{
													Required: true,
												},
												"type": schema.StringAttribute{
													Required:   true,
													Validators: []validator.String{stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean")},
												},
												"format": schema.StringAttribute{
													Optional: true,
												},
												"date": schema.StringAttribute{
													Optional: true,
												},
											},
										},
									},
								},
							},
							"decoder": schema.SingleNestedAttribute{
								Optional: true,
								Attributes: map[string]schema.Attribute{
									"match_name": schema.StringAttribute{
										Optional:            true,
										MarkdownDescription: "Relative path after decompression (regular expression). If not entered, all data in the compressed file will be transferred.",
									},
								},
							},
							"custom_variable_settings": schema.ListNestedAttribute{
								Optional: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Required: true,
											Validators: []validator.String{
												validator2.WrappingDollarValidator{},
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
										&planmodifier2.CustomVariableSettingPlanModifier{},
									},
								},
							},
						},
						PlanModifiers: []planmodifier.Object{
							&fileParserPlanModifier{},
							&gcsInputOptionPlanModifier{},
						},
					},
				},
			},
			"output_option_type": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.UTF8LengthAtLeast(1),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				MarkdownDescription: "Output option type.",
			},
			"output_option": schema.SingleNestedAttribute{
				Required: true,
				Attributes: map[string]schema.Attribute{
					"bigquery_output_option": schema.SingleNestedAttribute{
						Optional:            true,
						MarkdownDescription: "Attributes of destination BigQuery settings",
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
							"location": schema.StringAttribute{
								Required: true,
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
								Required: true,
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
								Required:    true,
								ElementType: types.StringType,
							},
							"bigquery_output_option_merge_keys": schema.ListAttribute{
								Required:    true,
								ElementType: types.StringType,
							},
							"custom_variable_settings": schema.ListNestedAttribute{
								Optional: true,
								NestedObject: schema.NestedAttributeObject{
									Attributes: map[string]schema.Attribute{
										"name": schema.StringAttribute{
											Required: true,
											Validators: []validator.String{
												validator2.WrappingDollarValidator{},
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
										&planmodifier2.CustomVariableSettingPlanModifier{},
									},
								},
							},
						},
						PlanModifiers: []planmodifier.Object{
							&bigqueryOutputOptionPlanModifier{},
						},
					},
				},
			},
			"filter_columns": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
						"src": schema.StringAttribute{
							Required: true,
						},
						"type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("string", "long", "timestamp", "double", "boolean", "json"),
							},
						},
						"default": schema.StringAttribute{
							Optional: true,
						},
						"format": schema.StringAttribute{
							Optional: true,
						},
						"json_expand_enabled": schema.BoolAttribute{
							Required: true,
						},
						"json_expand_keep_base_column": schema.BoolAttribute{
							Required: true,
						},
						"json_expand_columns": schema.ListNestedAttribute{
							Required: true,
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"name": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.UTF8LengthAtLeast(1),
										},
									},
									"json_path": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.UTF8LengthAtLeast(1),
										},
									},
									"type": schema.StringAttribute{
										Required: true,
										Validators: []validator.String{
											stringvalidator.OneOf("boolean", "long", "timestamp", "string"),
										},
									},
									"format": schema.StringAttribute{
										Optional: true,
									},
									"timezone": schema.StringAttribute{
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"filter_rows": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"condition": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf("and", "or"),
						},
					},
					"filter_row_conditions": schema.ListNestedAttribute{
						Required: true,
						NestedObject: schema.NestedAttributeObject{
							Attributes: map[string]schema.Attribute{
								"column": schema.StringAttribute{
									Required: true,
									Validators: []validator.String{
										stringvalidator.UTF8LengthAtLeast(1),
									},
								},
								"operator": schema.StringAttribute{
									Required: true,
									Validators: []validator.String{
										stringvalidator.OneOf("greater", "greater_equal", "less", "less_equal", "equal", "not_equal", "start_with", "end_with", "include", "is_null", "is_not_null", "regexp"),
									},
								},
								"argument": schema.StringAttribute{
									Required: true,
									Validators: []validator.String{
										stringvalidator.UTF8LengthAtLeast(1),
									},
								},
							},
						},
					},
				},
			},
			"filter_masks": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
						"mask_type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("all", "email", "regex", "substring"),
							},
						},
						"length": schema.Int64Attribute{
							Optional: true,
						},
						"pattern": schema.StringAttribute{
							Optional: true,
						},
						"start_index": schema.Int64Attribute{
							Optional: true,
						},
						"end_index": schema.Int64Attribute{
							Optional: true,
						},
					},
				},
				MarkdownDescription: "Filter masks to be attached to the job definition",
			},
			"filter_add_time": schema.SingleNestedAttribute{
				Optional: true,
				Attributes: map[string]schema.Attribute{
					"column_name": schema.StringAttribute{
						Required: true,
					},
					"type": schema.StringAttribute{
						Required: true,
						Validators: []validator.String{
							stringvalidator.OneOf("timestamp", "string"),
						},
					},
					"timestamp_format": schema.StringAttribute{
						Optional: true,
						Computed: true,
					},
					"time_zone": schema.StringAttribute{
						Optional: true,
						Computed: true,
					},
				},
			},
			"filter_gsub": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"column_name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
						"pattern": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
						"to": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
					},
				},
				MarkdownDescription: "Filter gsub to be attached to the job definition",
			},
			"filter_string_transforms": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"column_name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
						"type": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("normalize_nfkc"),
							},
						},
					},
				},
			},
			"filter_hashes": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
					},
				},
			},
			"filter_unixtime_conversions": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"column_name": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.UTF8LengthAtLeast(1),
							},
						},
						"kind": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("unixtime_to_timestamp", "unixtime_to_string", "timestamp_to_unixtime", "string_to_unixtime"),
							},
						},
						"unixtime_unit": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf("second", "millisecond", "microsecond", "nanosecond"),
							},
						},
						"datetime_format": schema.StringAttribute{
							Required: true,
						},
						"datetime_timezone": schema.StringAttribute{
							Required: true,
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
						&planmodifier2.SchedulePlanModifier{},
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
								stringvalidator.OneOf("job", "record", "exec_time"),
							},
							MarkdownDescription: "Category of condition. The following types are supported: `job`, `record`, `exec_time`",
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
						"minutes": schema.Int64Attribute{
							Optional: true,
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
	ID                        types.Int64                                 `tfsdk:"id"`
	Name                      types.String                                `tfsdk:"name"`
	Description               types.String                                `tfsdk:"description"`
	ResourceGroupID           types.Int64                                 `tfsdk:"resource_group_id"`
	IsRunnableConcurrently    types.Bool                                  `tfsdk:"is_runnable_concurrently"`
	RetryLimit                types.Int64                                 `tfsdk:"retry_limit"`
	ResourceEnhancement       types.String                                `tfsdk:"resource_enhancement"`
	InputOptionType           types.String                                `tfsdk:"input_option_type"`
	InputOption               *job_definitions.InputOption                `tfsdk:"input_option"`
	OutputOptionType          types.String                                `tfsdk:"output_option_type"`
	OutputOption              *job_definitions.OutputOption               `tfsdk:"output_option"`
	FilterColumns             []filter.FilterColumn                       `tfsdk:"filter_columns"`
	FilterRows                *filter.FilterRows                          `tfsdk:"filter_rows"`
	FilterMasks               []filter.FilterMask                         `tfsdk:"filter_masks"`
	FilterAddTime             *filter.FilterAddTime                       `tfsdk:"filter_add_time"`
	FilterGsub                []filter.FilterGsub                         `tfsdk:"filter_gsub"`
	FilterStringTransforms    []filter.FilterStringTransform              `tfsdk:"filter_string_transforms"`
	FilterHashes              []filter.FilterHash                         `tfsdk:"filter_hashes"`
	FilterUnixTimeConversions []filter.FilterUnixTimeConversion           `tfsdk:"filter_unixtime_conversions"`
	Notifications             []job_definitions.JobDefinitionNotification `tfsdk:"notifications"`
	Schedules                 []models.Schedule                           `tfsdk:"schedules"`
	Labels                    []models.LabelModel                         `tfsdk:"labels"`
}

func (model *jobDefinitionResourceModel) ToCreateJobDefinitionInput() *client.CreateJobDefinitionInput {
	var labels []string
	if model.Labels != nil {
		for _, l := range model.Labels {
			labels = append(labels, l.Name.ValueString())
		}
	}
	var notifications []job_definitions2.JobDefinitionNotificationInput
	if model.Notifications != nil {
		for _, n := range model.Notifications {
			notifications = append(notifications, n.ToInput())
		}
	}
	var schedules []parameters.ScheduleInput
	if model.Schedules != nil {
		for _, s := range model.Schedules {
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
		Description:               models.NewNullableString(model.Description),
		ResourceGroupID:           models.NewNullableInt64(model.ResourceGroupID),
		IsRunnableConcurrently:    model.IsRunnableConcurrently.ValueBool(),
		RetryLimit:                model.RetryLimit.ValueInt64(),
		ResourceEnhancement:       model.ResourceEnhancement.ValueStringPointer(),
		FilterColumns:             filterColumns,
		FilterRows:                models.NewNullableObject(model.FilterRows.ToInput()),
		FilterMasks:               filterMasks,
		FilterAddTime:             models.NewNullableObject(model.FilterAddTime.ToInput()),
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
	state := &jobDefinitionResourceModel{}
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)
	if response.Diagnostics.HasError() {
		return
	}

	plan := &jobDefinitionResourceModel{}
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)
	if response.Diagnostics.HasError() {
		return
	}

	jobDefinition, err := r.client.UpdateJobDefinition(
		state.ID.ValueInt64(),
		plan.ToUpdateJobDefinitionInput(),
	)
	if err != nil {
		response.Diagnostics.AddError(
			"Updating job definition",
			fmt.Sprintf("Unable to update job definition, got error: %s", err),
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
	response.Diagnostics.Append(response.State.Set(ctx, newState)...)
}

func (model *jobDefinitionResourceModel) ToUpdateJobDefinitionInput() *client.UpdateJobDefinitionInput {
	labels := []string{}
	if model.Labels != nil {
		for _, l := range model.Labels {
			labels = append(labels, l.Name.ValueString())
		}
	}
	notifications := []job_definitions2.JobDefinitionNotificationInput{}
	if model.Notifications != nil {
		for _, n := range model.Notifications {
			notifications = append(notifications, n.ToInput())
		}
	}

	schedules := []parameters.ScheduleInput{}
	if model.Schedules != nil {
		for _, s := range model.Schedules {
			schedules = append(schedules, s.ToInput())
		}
	}

	filterColumns := []filter2.FilterColumnInput{}
	if model.FilterColumns != nil {
		for _, f := range model.FilterColumns {
			filterColumns = append(filterColumns, f.ToInput())
		}
	}

	filterMasks := []filter2.FilterMaskInput{}
	for _, f := range model.FilterMasks {
		filterMasks = append(filterMasks, f.ToInput())
	}

	filterGsub := []filter2.FilterGsubInput{}
	for _, f := range model.FilterGsub {
		filterGsub = append(filterGsub, f.ToInput())
	}

	filterStringTransforms := []filter2.FilterStringTransformInput{}
	for _, f := range model.FilterStringTransforms {
		filterStringTransforms = append(filterStringTransforms, f.ToInput())
	}

	filterHashes := []filter2.FilterHashInput{}
	for _, f := range model.FilterHashes {
		filterHashes = append(filterHashes, f.ToInput())
	}

	filterUnixTimeconversions := []filter2.FilterUnixTimeConversionInput{}
	for _, f := range model.FilterUnixTimeConversions {
		filterUnixTimeconversions = append(filterUnixTimeconversions, f.ToInput())
	}

	return &client.UpdateJobDefinitionInput{
		Name:                      model.Name.ValueStringPointer(),
		Description:               models.NewNullableString(model.Description),
		ResourceGroupID:           models.NewNullableInt64(model.ResourceGroupID),
		IsRunnableConcurrently:    model.IsRunnableConcurrently.ValueBoolPointer(),
		RetryLimit:                model.RetryLimit.ValueInt64Pointer(),
		ResourceEnhancement:       model.ResourceEnhancement.ValueStringPointer(),
		FilterColumns:             &filterColumns,
		FilterRows:                models.NewNullableObject(model.FilterRows.ToInput()),
		FilterMasks:               &filterMasks,
		FilterAddTime:             models.NewNullableObject(model.FilterAddTime.ToInput()),
		FilterGsub:                &filterGsub,
		FilterStringTransforms:    &filterStringTransforms,
		FilterHashes:              &filterHashes,
		FilterUnixTimeConversions: &filterUnixTimeconversions,
		InputOption:               model.InputOption.ToUpdateInput(),
		OutputOption:              model.OutputOption.ToUpdateInput(),
		Labels:                    &labels,
		Schedules:                 &schedules,
		Notifications:             &notifications,
	}
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
