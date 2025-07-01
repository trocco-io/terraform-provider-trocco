package job_definitions

import (
	"context"
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/model"
	input_options "terraform-provider-trocco/internal/provider/model/job_definition/input_option"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type InputOption struct {
	MySQLInputOption              *input_options.MySQLInputOption              `tfsdk:"mysql_input_option"`
	GcsInputOption                *input_options.GcsInputOption                `tfsdk:"gcs_input_option"`
	SnowflakeInputOption          *input_options.SnowflakeInputOption          `tfsdk:"snowflake_input_option"`
	SalesforceInputOption         *input_options.SalesforceInputOption         `tfsdk:"salesforce_input_option"`
	GoogleSpreadsheetsInputOption *input_options.GoogleSpreadsheetsInputOption `tfsdk:"google_spreadsheets_input_option"`
	S3InputOption                 *input_options.S3InputOption                 `tfsdk:"s3_input_option"`
	BigqueryInputOption           *input_options.BigqueryInputOption           `tfsdk:"bigquery_input_option"`
	PostgreSQLInputOption         *input_options.PostgreSQLInputOption         `tfsdk:"postgresql_input_option"`
	GoogleAnalytics4InputOption   *input_options.GoogleAnalytics4InputOption   `tfsdk:"google_analytics4_input_option"`
	HttpInputOption               *input_options.HttpInputOption               `tfsdk:"http_input_option"`
	KintoneInputOption            *input_options.KintoneInputOption            `tfsdk:"kintone_input_option"`
	YahooAdsApiYssInputOption     *input_options.YahooAdsApiYssInputOption     `tfsdk:"yahoo_ads_api_yss_input_option"`
}

func NewInputOption(ctx context.Context, inputOption client.InputOption, previous *InputOption) (*InputOption, diag.Diagnostics) {
	var previousHttpInputOption *input_options.HttpInputOption
	if previous != nil {
		previousHttpInputOption = previous.HttpInputOption
	}
	httpInputOption, diags := input_options.NewHttpInputOption(inputOption.HttpInputOption, previousHttpInputOption)
	return &InputOption{
		GcsInputOption:                input_options.NewGcsInputOption(ctx, inputOption.GcsInputOption),
		MySQLInputOption:              input_options.NewMysqlInputOption(ctx, inputOption.MySQLInputOption),
		SnowflakeInputOption:          input_options.NewSnowflakeInputOption(ctx, inputOption.SnowflakeInputOption),
		SalesforceInputOption:         input_options.NewSalesforceInputOption(ctx, inputOption.SalesforceInputOption),
		GoogleSpreadsheetsInputOption: input_options.NewGoogleSpreadsheetsInputOption(ctx, inputOption.GoogleSpreadsheetsInputOption),
		S3InputOption:                 input_options.NewS3InputOption(ctx, inputOption.S3InputOption),
		BigqueryInputOption:           input_options.NewBigqueryInputOption(ctx, inputOption.BigqueryInputOption),
		PostgreSQLInputOption:         input_options.NewPostgreSQLInputOption(ctx, inputOption.PostgreSQLInputOption),
		GoogleAnalytics4InputOption:   input_options.NewGoogleAnalytics4InputOption(ctx, inputOption.GoogleAnalytics4InputOption),
		HttpInputOption:               httpInputOption,
		KintoneInputOption:            input_options.NewKintoneInputOption(ctx, inputOption.KintoneInputOption),
		YahooAdsApiYssInputOption:     input_options.NewYahooAdsApiYssInputOption(ctx, inputOption.YahooAdsApiYssInputOption),
	}, diags
}

func (o InputOption) ToInput(ctx context.Context) (client.InputOptionInput, diag.Diagnostics) {
	var diags diag.Diagnostics

	httpInput, d := o.HttpInputOption.ToInput()
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
	}, diags
}

func (o InputOption) ToUpdateInput(ctx context.Context) (*client.UpdateInputOptionInput, diag.Diagnostics) {
	var diags diag.Diagnostics
	httpInput, d := o.HttpInputOption.ToUpdateInput()
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
	}, diags
}
