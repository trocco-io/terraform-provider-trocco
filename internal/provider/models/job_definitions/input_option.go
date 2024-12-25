package job_definitions

import (
	"terraform-provider-trocco/internal/client"
	input_options2 "terraform-provider-trocco/internal/client/parameters/job_definitions/input_options"
	"terraform-provider-trocco/internal/provider/models/job_definitions/input_options"
)

type InputOption struct {
	MySQLInputOption *input_options.MySQLInputOption `tfsdk:"mysql_input_option"`
	// GcsInputOption   *input_options.GcsInputOption   `tfsdk:"gcs_input_option"`
}

func NewInputOption(inputOption client.InputOption) *InputOption {
	return &InputOption{
		//GcsInputOption:   input_options.NewGcsInputOption(inputOption.GcsInputOption),
		MySQLInputOption: input_options.NewMysqlInputOption(inputOption.MySQLInputOption),
	}
}

func (inputOption InputOption) ToInput() client.InputOptionInput {
	return client.InputOptionInput{
		// GcsInputOption:   inputOption.GcsInputOption.ToInput(),
		MySQLInputOption: func() *input_options2.MySQLInputOptionInput {
			if inputOption.MySQLInputOption == nil {
				return nil
			}
			return inputOption.MySQLInputOption.ToInput()
		}(),
	}
}

func (inputOption InputOption) ToUpdateInput() *client.UpdateInputOptionInput {
	return &client.UpdateInputOptionInput{
		// GcsInputOption:   inputOption.GcsInputOption.ToInput(),
		MySQLInputOption: func() *input_options2.UpdateMySQLInputOptionInput {
			if inputOption.MySQLInputOption == nil {
				return nil
			}
			return inputOption.MySQLInputOption.ToUpdateInput()
		}(),
	}
}
