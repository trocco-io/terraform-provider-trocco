package filter

type FilterMaskInput struct {
	Name       string  `json:"name"`
	MaskType   string  `json:"mask_type"`
	Length     *int64  `json:"length,omitempty"`
	Pattern    *string `json:"pattern,omitempty"`
	StartIndex *int64  `json:"start_index,omitempty"`
	EndIndex   *int64  `json:"end_index,omitempty"`
}
