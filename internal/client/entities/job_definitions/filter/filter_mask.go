package filter

type FilterMask struct {
	Name       string  `json:"name"`
	MaskType   int32   `json:"mask_type"`
	Length     *int64  `json:"length"`
	Pattern    *string `json:"pattern"`
	StartIndex *int64  `json:"start_index"`
	EndIndex   *int64  `json:"end_index"`
}
