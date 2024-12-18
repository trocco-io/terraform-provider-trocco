package output_options

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/parameters"
)

type BigQueryOutputOption struct {
	CustomVariableSettings                 []parameters.CustomVariableSetting `tfsdk:"custom_variable_settings"`
	Dataset                                types.String                       `tfsdk:"dataset"`
	Table                                  types.String                       `tfsdk:"table"`
	AutoCreateDataset                      types.Bool                         `tfsdk:"auto_create_dataset"`
	AutoCreateTable                        types.Bool                         `tfsdk:"auto_create_table"`
	OpenTimeoutSec                         types.Int64                        `tfsdk:"open_timeout_sec"`
	TimeoutSec                             types.Int64                        `tfsdk:"timeout_sec"`
	SendTimeoutSec                         types.Int64                        `tfsdk:"send_timeout_sec"`
	ReadTimeoutSec                         types.Int64                        `tfsdk:"read_timeout_sec"`
	Retries                                types.Int64                        `tfsdk:"retries"`
	Mode                                   types.String                       `tfsdk:"mode"`
	PartitioningType                       types.String                       `tfsdk:"partitioning_type"`
	TimePartitioningType                   types.String                       `tfsdk:"time_partitioning_type"`
	TimePartitioningField                  types.String                       `tfsdk:"time_partitioning_field"`
	TimePartitioningExpirationMs           types.Int64                        `tfsdk:"time_partitioning_expiration_ms"`
	TimePartitioningRequirePartitionFilter types.Bool                         `tfsdk:"time_partitioning_require_partition_filter"`
	Location                               types.String                       `tfsdk:"location"`
	TemplateTable                          types.String                       `tfsdk:"template_table"`
	BigQueryConnectionID                   types.Int64                        `tfsdk:"bigquery_connection_id"`
	BeforeLoad                             types.String                       `tfsdk:"before_load"`
	BigQueryOutputOptionColumnOptions      []bigQueryOutputOptionColumnOption `tfsdk:"bigquery_output_option_column_options"`
	BigQueryOutputOptionClusteringFields   []types.String                     `tfsdk:"bigquery_output_option_clustering_fields"`
	BigQueryOutputOptionMergeKeys          []types.String                     `tfsdk:"bigquery_output_option_merge_keys"`
}

type bigQueryOutputOptionColumnOption struct {
	Name            types.String `tfsdk:"name"`
	Type            types.String `tfsdk:"type"`
	Mode            types.String `tfsdk:"mode"`
	TimestampFormat types.String `tfsdk:"timestamp_format"`
	Timezone        types.String `tfsdk:"timezone"`
	Description     types.String `tfsdk:"description"`
}
