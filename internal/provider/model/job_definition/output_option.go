package job_definitions

import (
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/model"
	output_options "terraform-provider-trocco/internal/provider/model/job_definition/output_option"
)

type OutputOption struct {
	BigQueryOutputOption  *output_options.BigQueryOutputOption  `tfsdk:"bigquery_output_option"`
	SnowflakeOutputOption *output_options.SnowflakeOutputOption `tfsdk:"snowflake_output_option"`
}

func NewOutputOption(outputOption client.OutputOption) *OutputOption {
	return &OutputOption{
		BigQueryOutputOption:  output_options.NewBigQueryOutputOption(outputOption.BigQueryOutputOption),
		SnowflakeOutputOption: output_options.NewSnowflakeOutputOption(outputOption.SnowflakeOutputOption),
	}
}

func (o OutputOption) ToInput() client.OutputOptionInput {
	return client.OutputOptionInput{
		BigQueryOutputOption:  model.WrapObject(o.BigQueryOutputOption.ToInput()),
		SnowflakeOutputOption: model.WrapObject(o.SnowflakeOutputOption.ToInput()),
	}
}

func (o OutputOption) ToUpdateInput() *client.UpdateOutputOptionInput {
	return &client.UpdateOutputOptionInput{
		BigQueryOutputOption:  model.WrapObject(o.BigQueryOutputOption.ToUpdateInput()),
		SnowflakeOutputOption: model.WrapObject(o.SnowflakeOutputOption.ToUpdateInput()),
	}
}
