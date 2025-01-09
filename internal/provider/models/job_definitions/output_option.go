package job_definitions

import (
	"terraform-provider-trocco/internal/client"
	output_options2 "terraform-provider-trocco/internal/client/parameters/job_definitions/output_options"
	"terraform-provider-trocco/internal/provider/models/job_definitions/output_options"
)

type OutputOption struct {
	BigQueryOutputOption *output_options.BigQueryOutputOption `tfsdk:"bigquery_output_option"`
}

func NewOutputOption(outputOption client.OutputOption) *OutputOption {
	return &OutputOption{
		BigQueryOutputOption: output_options.NewBigQueryOutputOption(outputOption.BigQueryOutputOption),
	}
}

func (outputOption OutputOption) ToInput() client.OutputOptionInput {
	return client.OutputOptionInput{
		BigQueryOutputOption: func() *output_options2.BigQueryOutputOptionInput {
			if outputOption.BigQueryOutputOption == nil {
				return nil
			}
			return outputOption.BigQueryOutputOption.ToInput()
		}(),
	}
}

func (outputOption OutputOption) ToUpdateInput() *client.UpdateOutputOptionInput {
	return &client.UpdateOutputOptionInput{
		BigQueryOutputOption: func() *output_options2.UpdateBigQueryOutputOptionInput {
			if outputOption.BigQueryOutputOption == nil {
				return nil
			}
			return outputOption.BigQueryOutputOption.ToUpdateInput()
		}(),
	}
}
