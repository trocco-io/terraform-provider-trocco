package job_definitions

import (
	"context"
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/model"
	outputOptionModel "terraform-provider-trocco/internal/provider/model/job_definition/output_option"
)

type OutputOption struct {
	BigQueryOutputOption           *outputOptionModel.BigQueryOutputOption           `tfsdk:"bigquery_output_option"`
	SnowflakeOutputOption          *outputOptionModel.SnowflakeOutputOption          `tfsdk:"snowflake_output_option"`
	SalesforceOutputOption         *outputOptionModel.SalesforceOutputOption         `tfsdk:"salesforce_output_option"`
	GoogleSpreadsheetsOutputOption *outputOptionModel.GoogleSpreadsheetsOutputOption `tfsdk:"google_spreadsheets_output_option"`
	DatabricksOutputOption         *outputOptionModel.DatabricksOutputOption         `tfsdk:"databricks_output_option"`
}

func NewOutputOption(ctx context.Context, outputOption client.OutputOption) *OutputOption {
	return &OutputOption{
		BigQueryOutputOption:           outputOptionModel.NewBigQueryOutputOption(ctx, outputOption.BigQueryOutputOption),
		SnowflakeOutputOption:          outputOptionModel.NewSnowflakeOutputOption(ctx, outputOption.SnowflakeOutputOption),
		SalesforceOutputOption:         outputOptionModel.NewSalesforceOutputOption(outputOption.SalesforceOutputOption),
		GoogleSpreadsheetsOutputOption: outputOptionModel.NewGoogleSpreadsheetsOutputOption(ctx, outputOption.GoogleSpreadsheetsOutputOption),
		DatabricksOutputOption:         outputOptionModel.NewDatabricksOutputOption(outputOption.DatabricksOutputOption),
	}
}

func (o OutputOption) ToInput(ctx context.Context) client.OutputOptionInput {
	return client.OutputOptionInput{
		BigQueryOutputOption:           model.WrapObject(o.BigQueryOutputOption.ToInput(ctx)),
		SnowflakeOutputOption:          model.WrapObject(o.SnowflakeOutputOption.ToInput(ctx)),
		SalesforceOutputOption:         model.WrapObject(o.SalesforceOutputOption.ToInput()),
		GoogleSpreadsheetsOutputOption: model.WrapObject(o.GoogleSpreadsheetsOutputOption.ToInput(ctx)),
		DatabricksOutputOption:         model.WrapObject(o.DatabricksOutputOption.ToInput()),
	}
}

func (o OutputOption) ToUpdateInput(ctx context.Context) *client.UpdateOutputOptionInput {
	return &client.UpdateOutputOptionInput{
		BigQueryOutputOption:           model.WrapObject(o.BigQueryOutputOption.ToUpdateInput(ctx)),
		SnowflakeOutputOption:          model.WrapObject(o.SnowflakeOutputOption.ToUpdateInput(ctx)),
		SalesforceOutputOption:         model.WrapObject(o.SalesforceOutputOption.ToUpdateInput()),
		GoogleSpreadsheetsOutputOption: model.WrapObject(o.GoogleSpreadsheetsOutputOption.ToUpdateInput(ctx)),
		DatabricksOutputOption:         model.WrapObject(o.DatabricksOutputOption.ToUpdateInput()),
	}
}
