package input_options

import (
	job_definitions "terraform-provider-trocco/internal/client/entity/job_definition"
	parmas "terraform-provider-trocco/internal/client/parameter/job_definitions"

	"github.com/hashicorp/terraform-plugin-framework/types"
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

func (decoder *Decoder) ToDecoderInput() *parmas.DecoderInput {
	if decoder == nil {
		return nil
	}
	return &parmas.DecoderInput{
		MatchName: decoder.MatchName.ValueString(),
	}
}
