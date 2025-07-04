package job_definitions

import (
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/model"
	inputOptions "terraform-provider-trocco/internal/provider/model/job_definition/input_option"

	"github.com/hashicorp/terraform-plugin-framework/diag"
)

type InputOption struct {
	MySQLInputOption              *inputOptions.MySQLInputOption              `tfsdk:"mysql_input_option"`
	GcsInputOption                *inputOptions.GcsInputOption                `tfsdk:"gcs_input_option"`
	SnowflakeInputOption          *inputOptions.SnowflakeInputOption          `tfsdk:"snowflake_input_option"`
	SalesforceInputOption         *inputOptions.SalesforceInputOption         `tfsdk:"salesforce_input_option"`
	GoogleSpreadsheetsInputOption *inputOptions.GoogleSpreadsheetsInputOption `tfsdk:"google_spreadsheets_input_option"`
	S3InputOption                 *inputOptions.S3InputOption                 `tfsdk:"s3_input_option"`
	BigqueryInputOption           *inputOptions.BigqueryInputOption           `tfsdk:"bigquery_input_option"`
	PostgreSQLInputOption         *inputOptions.PostgreSQLInputOption         `tfsdk:"postgresql_input_option"`
	GoogleAnalytics4InputOption   *inputOptions.GoogleAnalytics4InputOption   `tfsdk:"google_analytics4_input_option"`
	HTTPInputOption               *inputOptions.HTTPInputOption               `tfsdk:"http_input_option"`
	KintoneInputOption            *inputOptions.KintoneInputOption            `tfsdk:"kintone_input_option"`
	YahooAdsApiYssInputOption     *inputOptions.YahooAdsApiYssInputOption     `tfsdk:"yahoo_ads_api_yss_input_option"`
}

func NewInputOption(inputOption client.InputOption, previous *InputOption) (*InputOption, diag.Diagnostics) {
	var previousHTTPInputOption *inputOptions.HTTPInputOption
	if previous != nil {
		previousHTTPInputOption = previous.HTTPInputOption
	}
	httpInputOption, diags := inputOptions.NewHTTPInputOption(inputOption.HTTPInputOption, previousHTTPInputOption)
	return &InputOption{
		GcsInputOption:                inputOptions.NewGcsInputOption(inputOption.GcsInputOption),
		MySQLInputOption:              inputOptions.NewMysqlInputOption(inputOption.MySQLInputOption),
		SnowflakeInputOption:          inputOptions.NewSnowflakeInputOption(inputOption.SnowflakeInputOption),
		SalesforceInputOption:         inputOptions.NewSalesforceInputOption(inputOption.SalesforceInputOption),
		GoogleSpreadsheetsInputOption: inputOptions.NewGoogleSpreadsheetsInputOption(inputOption.GoogleSpreadsheetsInputOption),
		S3InputOption:                 inputOptions.NewS3InputOption(inputOption.S3InputOption),
		BigqueryInputOption:           inputOptions.NewBigqueryInputOption(inputOption.BigqueryInputOption),
		PostgreSQLInputOption:         inputOptions.NewPostgreSQLInputOption(inputOption.PostgreSQLInputOption),
		GoogleAnalytics4InputOption:   inputOptions.NewGoogleAnalytics4InputOption(inputOption.GoogleAnalytics4InputOption),
		HTTPInputOption:               httpInputOption,
		KintoneInputOption:            inputOptions.NewKintoneInputOption(inputOption.KintoneInputOption),
		YahooAdsApiYssInputOption:     inputOptions.NewYahooAdsApiYssInputOption(inputOption.YahooAdsApiYssInputOption),
	}, diags
}

func (o InputOption) ToInput() (client.InputOptionInput, diag.Diagnostics) {
	var diags diag.Diagnostics

	httpInput, d := o.HTTPInputOption.ToInput()
	diags.Append(d...)

	return client.InputOptionInput{
		GcsInputOption:                model.WrapObject(o.GcsInputOption.ToInput()),
		MySQLInputOption:              model.WrapObject(o.MySQLInputOption.ToInput()),
		SnowflakeInputOption:          model.WrapObject(o.SnowflakeInputOption.ToInput()),
		SalesforceInputOption:         model.WrapObject(o.SalesforceInputOption.ToInput()),
		GoogleSpreadsheetsInputOption: model.WrapObject(o.GoogleSpreadsheetsInputOption.ToInput()),
		S3InputOption:                 model.WrapObject(o.S3InputOption.ToInput()),
		BigqueryInputOption:           model.WrapObject(o.BigqueryInputOption.ToInput()),
		PostgreSQLInputOption:         model.WrapObject(o.PostgreSQLInputOption.ToInput()),
		GoogleAnalytics4InputOption:   model.WrapObject(o.GoogleAnalytics4InputOption.ToInput()),
		HTTPInputOption:               model.WrapObject(httpInput),
		KintoneInputOption:            model.WrapObject(o.KintoneInputOption.ToInput()),
		YahooAdsApiYssInputOption:     model.WrapObject(o.YahooAdsApiYssInputOption.ToInput()),
	}, diags
}

func (o InputOption) ToUpdateInput() (*client.UpdateInputOptionInput, diag.Diagnostics) {
	var diags diag.Diagnostics
	httpInput, d := o.HTTPInputOption.ToUpdateInput()
	diags.Append(d...)

	return &client.UpdateInputOptionInput{
		GcsInputOption:                model.WrapObject(o.GcsInputOption.ToUpdateInput()),
		MySQLInputOption:              model.WrapObject(o.MySQLInputOption.ToUpdateInput()),
		SnowflakeInputOption:          model.WrapObject(o.SnowflakeInputOption.ToUpdateInput()),
		SalesforceInputOption:         model.WrapObject(o.SalesforceInputOption.ToUpdateInput()),
		GoogleSpreadsheetsInputOption: model.WrapObject(o.GoogleSpreadsheetsInputOption.ToUpdateInput()),
		S3InputOption:                 model.WrapObject(o.S3InputOption.ToUpdateInput()),
		BigqueryInputOption:           model.WrapObject(o.BigqueryInputOption.ToUpdateInput()),
		PostgreSQLInputOption:         model.WrapObject(o.PostgreSQLInputOption.ToUpdateInput()),
		GoogleAnalytics4InputOption:   model.WrapObject(o.GoogleAnalytics4InputOption.ToUpdateInput()),
		HTTPInputOption:               model.WrapObject(httpInput),
		KintoneInputOption:            model.WrapObject(o.KintoneInputOption.ToUpdateInput()),
		YahooAdsApiYssInputOption:     model.WrapObject(o.YahooAdsApiYssInputOption.ToUpdateInput()),
	}, diags
}
