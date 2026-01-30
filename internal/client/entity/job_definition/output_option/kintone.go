package output_option

type KintoneOutputOption struct {
	KintoneConnectionID              int64                             `json:"kintone_connection_id"`
	AppID                            string                            `json:"app_id"`
	GuestSpaceID                     *string                           `json:"guest_space_id"`
	Mode                             string                            `json:"mode"`
	UpdateKey                        *string                           `json:"update_key"`
	IgnoreNulls                      bool                              `json:"ignore_nulls"`
	ReduceKey                        *string                           `json:"reduce_key"`
	ChunkSize                        int64                             `json:"chunk_size"`
	KintoneOutputOptionColumnOptions []KintoneOutputOptionColumnOption `json:"kintone_output_option_column_options"`
}

type KintoneOutputOptionColumnOption struct {
	Name       string  `json:"name"`
	FieldCode  string  `json:"field_code"`
	Type       string  `json:"type"`
	Timezone   *string `json:"timezone"`
	SortColumn *string `json:"sort_column"`
}
