package job_definitions

import (
	"context"
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

func NewOutputOption(ctx context.Context, outputOption client.OutputOption) *OutputOption {
	return &OutputOption{
		BigQueryOutputOption:           output_options.NewBigQueryOutputOption(ctx, outputOption.BigQueryOutputOption),
		SnowflakeOutputOption:          output_options.NewSnowflakeOutputOption(ctx, outputOption.SnowflakeOutputOption),
		SalesforceOutputOption:         output_options.NewSalesforceOutputOption(outputOption.SalesforceOutputOption),
		GoogleSpreadsheetsOutputOption: output_options.NewGoogleSpreadsheetsOutputOption(ctx, outputOption.GoogleSpreadsheetsOutputOption),
	}
}

func (o OutputOption) ToInput(ctx context.Context) client.OutputOptionInput {
	return client.OutputOptionInput{
		BigQueryOutputOption:           model.WrapObject(o.BigQueryOutputOption.ToInput(ctx)),
		SnowflakeOutputOption:          model.WrapObject(o.SnowflakeOutputOption.ToInput(ctx)),
		SalesforceOutputOption:         model.WrapObject(o.SalesforceOutputOption.ToInput()),
		GoogleSpreadsheetsOutputOption: model.WrapObject(o.GoogleSpreadsheetsOutputOption.ToInput(ctx)),
	}
}

func (o OutputOption) ToUpdateInput(ctx context.Context) *client.UpdateOutputOptionInput {
	return &client.UpdateOutputOptionInput{
		BigQueryOutputOption:           model.WrapObject(o.BigQueryOutputOption.ToUpdateInput(ctx)),
		SnowflakeOutputOption:          model.WrapObject(o.SnowflakeOutputOption.ToUpdateInput(ctx)),
		SalesforceOutputOption:         model.WrapObject(o.SalesforceOutputOption.ToUpdateInput()),
		GoogleSpreadsheetsOutputOption: model.WrapObject(o.GoogleSpreadsheetsOutputOption.ToUpdateInput(ctx)),
	}
}
