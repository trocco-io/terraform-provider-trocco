package job_definitions

import (
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/model"
	"terraform-provider-trocco/internal/provider/model/job_definitions/output_options"
)

type OutputOption struct {
	BigQueryOutputOption *output_options.BigQueryOutputOption `tfsdk:"bigquery_output_option"`
}

func NewOutputOption(outputOption client.OutputOption) *OutputOption {
	return &OutputOption{
		BigQueryOutputOption: output_options.NewBigQueryOutputOption(outputOption.BigQueryOutputOption),
	}
}

func (o OutputOption) ToInput() client.OutputOptionInput {
	return client.OutputOptionInput{
		BigQueryOutputOption: model.WrapObject(o.BigQueryOutputOption.ToInput()),
	}
}

func (o OutputOption) ToUpdateInput() *client.UpdateOutputOptionInput {
	return &client.UpdateOutputOptionInput{
		BigQueryOutputOption: model.WrapObject(o.BigQueryOutputOption.ToUpdateInput()),
	}
}
