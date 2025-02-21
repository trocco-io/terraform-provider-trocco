package job_definitions

import (
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/model"
	output_options "terraform-provider-trocco/internal/provider/model/job_definition/output_option"
)

type OutputOption struct {
	BigQueryOutputOption           *output_options.BigQueryOutputOption           `tfsdk:"bigquery_output_option"`
	SnowflakeOutputOption          *output_options.SnowflakeOutputOption          `tfsdk:"snowflake_output_option"`
	SalesforceOutputOption         *output_options.SalesforceOutputOption         `tfsdk:"salesforce_output_option"`
	GoogleSpreadsheetsOutputOption *output_options.GoogleSpreadsheetsOutputOption `tfsdk:"google_spreadsheets_output_option"`
}

func NewOutputOption(outputOption client.OutputOption) *OutputOption {
	return &OutputOption{
		BigQueryOutputOption:           output_options.NewBigQueryOutputOption(outputOption.BigQueryOutputOption),
		SnowflakeOutputOption:          output_options.NewSnowflakeOutputOption(outputOption.SnowflakeOutputOption),
		SalesforceOutputOption:         output_options.NewSalesforceOutputOption(outputOption.SalesforceOutputOption),
		GoogleSpreadsheetsOutputOption: output_options.NewGoogleSpreadsheetsOutputOption(outputOption.GoogleSpreadsheetsOutputOption),
	}
}

func (o OutputOption) ToInput() client.OutputOptionInput {
	return client.OutputOptionInput{
		BigQueryOutputOption:           model.WrapObject(o.BigQueryOutputOption.ToInput()),
		SnowflakeOutputOption:          model.WrapObject(o.SnowflakeOutputOption.ToInput()),
		SalesforceOutputOption:         model.WrapObject(o.SalesforceOutputOption.ToInput()),
		GoogleSpreadsheetsOutputOption: model.WrapObject(o.GoogleSpreadsheetsOutputOption.ToInput()),
	}
}

func (o OutputOption) ToUpdateInput() *client.UpdateOutputOptionInput {
	return &client.UpdateOutputOptionInput{
		BigQueryOutputOption:           model.WrapObject(o.BigQueryOutputOption.ToUpdateInput()),
		SnowflakeOutputOption:          model.WrapObject(o.SnowflakeOutputOption.ToUpdateInput()),
		SalesforceOutputOption:         model.WrapObject(o.SalesforceOutputOption.ToUpdateInput()),
		GoogleSpreadsheetsOutputOption: model.WrapObject(o.GoogleSpreadsheetsOutputOption.ToUpdateInput()),
	}
}
