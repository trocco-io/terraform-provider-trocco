package input_options

import (
	jobDefinitionEntities "terraform-provider-trocco/internal/client/entity/job_definition"
	parmas "terraform-provider-trocco/internal/client/parameter/job_definition"

	"github.com/hashicorp/terraform-plugin-framework/types"
)

type Decoder struct {
	MatchName types.String `tfsdk:"match_name"`
}

func NewDecoder(decoder *jobDefinitionEntities.Decoder, previous *Decoder) *Decoder {
	if decoder == nil {
		// The TROCCO API normalizes a decoder with an empty match_name to null,
		// so keep the previous empty decoder to avoid inconsistent results after apply.
		if previous != nil && previous.MatchName.ValueString() == "" {
			return previous
		}
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
