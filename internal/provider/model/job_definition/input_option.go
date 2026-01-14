package job_definitions

import (
	"context"
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/model"
	inputOptionModel "terraform-provider-trocco/internal/provider/model/job_definition/input_option"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type InputOption struct {
	MySQLInputOption              *inputOptionModel.MySQLInputOption              `tfsdk:"mysql_input_option"`
	GcsInputOption                *inputOptionModel.GcsInputOption                `tfsdk:"gcs_input_option"`
	SnowflakeInputOption          *inputOptionModel.SnowflakeInputOption          `tfsdk:"snowflake_input_option"`
	SalesforceInputOption         *inputOptionModel.SalesforceInputOption         `tfsdk:"salesforce_input_option"`
	GoogleSpreadsheetsInputOption *inputOptionModel.GoogleSpreadsheetsInputOption `tfsdk:"google_spreadsheets_input_option"`
	S3InputOption                 *inputOptionModel.S3InputOption                 `tfsdk:"s3_input_option"`
	BigqueryInputOption           *inputOptionModel.BigqueryInputOption           `tfsdk:"bigquery_input_option"`
	PostgreSQLInputOption         *inputOptionModel.PostgreSQLInputOption         `tfsdk:"postgresql_input_option"`
	GoogleAnalytics4InputOption   *inputOptionModel.GoogleAnalytics4InputOption   `tfsdk:"google_analytics4_input_option"`
	HttpInputOption               *inputOptionModel.HttpInputOption               `tfsdk:"http_input_option"`
	KintoneInputOption            *inputOptionModel.KintoneInputOption            `tfsdk:"kintone_input_option"`
	YahooAdsApiYssInputOption     *inputOptionModel.YahooAdsApiYssInputOption     `tfsdk:"yahoo_ads_api_yss_input_option"`
	DatabricksInputOption         *inputOptionModel.DatabricksInputOption         `tfsdk:"databricks_input_option"`
	MongoDBInputOption            *inputOptionModel.MongoDBInputOption            `tfsdk:"mongodb_input_option"`
}

func NewInputOption(ctx context.Context, inputOption client.InputOption, previous *InputOption) (*InputOption, diag.Diagnostics) {
	var previousHttpInputOption *inputOptionModel.HttpInputOption
	if previous != nil {
		previousHttpInputOption = previous.HttpInputOption
	}
	httpInputOption, diags := inputOptionModel.NewHttpInputOption(ctx, inputOption.HttpInputOption, previousHttpInputOption)
	return &InputOption{
		GcsInputOption:                inputOptionModel.NewGcsInputOption(ctx, inputOption.GcsInputOption),
		MySQLInputOption:              inputOptionModel.NewMysqlInputOption(ctx, inputOption.MySQLInputOption),
		SnowflakeInputOption:          inputOptionModel.NewSnowflakeInputOption(ctx, inputOption.SnowflakeInputOption),
		SalesforceInputOption:         inputOptionModel.NewSalesforceInputOption(ctx, inputOption.SalesforceInputOption),
		GoogleSpreadsheetsInputOption: inputOptionModel.NewGoogleSpreadsheetsInputOption(ctx, inputOption.GoogleSpreadsheetsInputOption),
		S3InputOption:                 inputOptionModel.NewS3InputOption(ctx, inputOption.S3InputOption),
		BigqueryInputOption:           inputOptionModel.NewBigqueryInputOption(ctx, inputOption.BigqueryInputOption),
		PostgreSQLInputOption:         inputOptionModel.NewPostgreSQLInputOption(ctx, inputOption.PostgreSQLInputOption),
		GoogleAnalytics4InputOption:   inputOptionModel.NewGoogleAnalytics4InputOption(ctx, inputOption.GoogleAnalytics4InputOption),
		HttpInputOption:               httpInputOption,
		KintoneInputOption:            inputOptionModel.NewKintoneInputOption(ctx, inputOption.KintoneInputOption),
		YahooAdsApiYssInputOption:     inputOptionModel.NewYahooAdsApiYssInputOption(ctx, inputOption.YahooAdsApiYssInputOption),
		DatabricksInputOption:         inputOptionModel.NewDatabricksInputOption(ctx, inputOption.DatabricksInputOption),
		MongoDBInputOption:            inputOptionModel.NewMongodbInputOption(ctx, inputOption.MongoDBInputOption),
	}, diags
}

func (o InputOption) ToInput(ctx context.Context) (client.InputOptionInput, diag.Diagnostics) {
	var diags diag.Diagnostics

	httpInput, d := o.HttpInputOption.ToInput(ctx)
	diags.Append(d...)

	return client.InputOptionInput{
		GcsInputOption:                model.WrapObject(o.GcsInputOption.ToInput(ctx)),
		MySQLInputOption:              model.WrapObject(o.MySQLInputOption.ToInput(ctx)),
		SnowflakeInputOption:          model.WrapObject(o.SnowflakeInputOption.ToInput(ctx)),
		SalesforceInputOption:         model.WrapObject(o.SalesforceInputOption.ToInput(ctx)),
		GoogleSpreadsheetsInputOption: model.WrapObject(o.GoogleSpreadsheetsInputOption.ToInput(ctx)),
		S3InputOption:                 model.WrapObject(o.S3InputOption.ToInput(ctx)),
		BigqueryInputOption:           model.WrapObject(o.BigqueryInputOption.ToInput(ctx)),
		PostgreSQLInputOption:         model.WrapObject(o.PostgreSQLInputOption.ToInput(ctx)),
		GoogleAnalytics4InputOption:   model.WrapObject(o.GoogleAnalytics4InputOption.ToInput(ctx)),
		HttpInputOption:               model.WrapObject(httpInput),
		KintoneInputOption:            model.WrapObject(o.KintoneInputOption.ToInput(ctx)),
		YahooAdsApiYssInputOption:     model.WrapObject(o.YahooAdsApiYssInputOption.ToInput(ctx)),
		DatabricksInputOption:         model.WrapObject(o.DatabricksInputOption.ToInput(ctx)),
		MongoDBInputOption:            model.WrapObject(o.MongoDBInputOption.ToInput(ctx)),
	}, diags
}

func (o InputOption) ToUpdateInput(ctx context.Context) (*client.UpdateInputOptionInput, diag.Diagnostics) {
	var diags diag.Diagnostics
	httpInput, d := o.HttpInputOption.ToUpdateInput(ctx)
	diags.Append(d...)

	return &client.UpdateInputOptionInput{
		GcsInputOption:                model.WrapObject(o.GcsInputOption.ToUpdateInput(ctx)),
		MySQLInputOption:              model.WrapObject(o.MySQLInputOption.ToUpdateInput(ctx)),
		SnowflakeInputOption:          model.WrapObject(o.SnowflakeInputOption.ToUpdateInput(ctx)),
		SalesforceInputOption:         model.WrapObject(o.SalesforceInputOption.ToUpdateInput(ctx)),
		GoogleSpreadsheetsInputOption: model.WrapObject(o.GoogleSpreadsheetsInputOption.ToUpdateInput(ctx)),
		S3InputOption:                 model.WrapObject(o.S3InputOption.ToUpdateInput(ctx)),
		BigqueryInputOption:           model.WrapObject(o.BigqueryInputOption.ToUpdateInput(ctx)),
		PostgreSQLInputOption:         model.WrapObject(o.PostgreSQLInputOption.ToUpdateInput(ctx)),
		GoogleAnalytics4InputOption:   model.WrapObject(o.GoogleAnalytics4InputOption.ToUpdateInput(ctx)),
		HttpInputOption:               model.WrapObject(httpInput),
		KintoneInputOption:            model.WrapObject(o.KintoneInputOption.ToUpdateInput(ctx)),
		YahooAdsApiYssInputOption:     model.WrapObject(o.YahooAdsApiYssInputOption.ToUpdateInput(ctx)),
		DatabricksInputOption:         model.WrapObject(o.DatabricksInputOption.ToUpdateInput(ctx)),
		MongoDBInputOption:            model.WrapObject(o.MongoDBInputOption.ToUpdateInput(ctx)),
	}, diags
}
