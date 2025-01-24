package job_definitions

import (
	"terraform-provider-trocco/internal/client"
	"terraform-provider-trocco/internal/provider/model"
	input_options "terraform-provider-trocco/internal/provider/model/job_definition/input_option"
)

type InputOption struct {
	MySQLInputOption *input_options.MySQLInputOption `tfsdk:"mysql_input_option"`
	GcsInputOption   *input_options.GcsInputOption   `tfsdk:"gcs_input_option"`
}

func NewInputOption(inputOption client.InputOption) *InputOption {
	return &InputOption{
		GcsInputOption:   input_options.NewGcsInputOption(inputOption.GcsInputOption),
		MySQLInputOption: input_options.NewMysqlInputOption(inputOption.MySQLInputOption),
	}
}

func (o InputOption) ToInput() client.InputOptionInput {
	return client.InputOptionInput{
		GcsInputOption:   model.WrapObject(o.GcsInputOption.ToInput()),
		MySQLInputOption: model.WrapObject(o.MySQLInputOption.ToInput()),
	}
}

func (o InputOption) ToUpdateInput() *client.UpdateInputOptionInput {
	return &client.UpdateInputOptionInput{
		GcsInputOption:   model.WrapObject(o.GcsInputOption.ToUpdateInput()),
		MySQLInputOption: model.WrapObject(o.MySQLInputOption.ToUpdateInput()),
	}
}
