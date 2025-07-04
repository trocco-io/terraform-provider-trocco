package input_options

import (
	jobDefEntity "terraform-provider-trocco/internal/client/entity/job_definition"
	jobDefParams "terraform-provider-trocco/internal/client/parameter/job_definition"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Decoder struct {
	MatchName types.String `tfsdk:"match_name"`
}

func NewDecoder(decoder *jobDefEntity.Decoder) *Decoder {
	if decoder == nil {
		return nil
	}
	return &Decoder{
		MatchName: types.StringValue(decoder.MatchName),
	}
}

func (decoder *Decoder) ToDecoderInput() *jobDefParams.DecoderInput {
	if decoder == nil {
		return nil
	}
	return &jobDefParams.DecoderInput{
		MatchName: decoder.MatchName.ValueString(),
	}
}
