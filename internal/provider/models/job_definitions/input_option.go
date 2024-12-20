package job_definitions

import (
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/models/job_definitions/input_options"
)

type InputOption struct {
	MySQLInputOption *input_options.MySQLInputOption `tfsdk:"mysql_input_option"`
	GcsInputOption   *input_options.GcsInputOption   `tfsdk:"gcs_input_option"`
}

func NewInputOption(inputOption client.InputOption) InputOption {
	return InputOption{
		GcsInputOption:   input_options.NewGcsInputOption(inputOption.GcsInputOption),
		MySQLInputOption: input_options.NewMysqlInputOption(inputOption.MySQLInputOption),
	}
}

func (inputOption InputOption) ToInput() client.InputOptionInput {
	return client.InputOptionInput{
		GcsInputOption:   inputOption.GcsInputOption.ToInput(),
		MySQLInputOption: inputOption.MySQLInputOption.ToInput(),
	}
}
