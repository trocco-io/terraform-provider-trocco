package input_options

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"terraform-provider-trocco/internal/client/entities/job_definitions"
	job_definitions2 "terraform-provider-trocco/internal/client/parameters/job_definitions"
)

type Decoder struct {
	MatchName types.String `tfsdk:"match_name"`
}

func NewDecoder(decoder *job_definitions.Decoder) *Decoder {
	if decoder == nil {
		return nil
	}
	return &Decoder{
		MatchName: types.StringValue(decoder.MatchName),
	}
}

func (decoder *Decoder) ToDecoderInput() *job_definitions2.DecoderInput {
	if decoder == nil {
		return nil
	}
	return &job_definitions2.DecoderInput{
		MatchName: decoder.MatchName.ValueString(),
	}
}

func ToDecoderModel(decoder *job_definitions.Decoder) *Decoder {
	if decoder == nil {
		return nil
	}
	return &Decoder{
		MatchName: types.StringValue(decoder.MatchName),
	}
}
