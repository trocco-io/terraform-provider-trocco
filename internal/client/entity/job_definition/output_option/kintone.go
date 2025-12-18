package output_option

import "strconv"

type KintoneOutputOption struct {
	KintoneConnectionID              int64                             `json:"kintone_connection_id"`
	AppID                            string                            `json:"app_id"`
	GuestSpaceID                     *int64                            `json:"guest_space_id"`
	Mode                             string                            `json:"mode"`
	UpdateKey                        *string                           `json:"update_key"`
	IgnoreNulls                      bool                              `json:"ignore_nulls"`
	ReduceKey                        *string                           `json:"reduce_key"`
	ChunkSize                        int64                             `json:"chunk_size"`
	KintoneOutputOptionColumnOptions []KintoneOutputOptionColumnOption `json:"kintone_output_option_column_options"`
}

func (k *KintoneOutputOption) GetAppIDAsInt64() int64 {
	if k.AppID == "" {
		return 0
	}
	id, _ := strconv.ParseInt(k.AppID, 10, 64)
	return id
}

type KintoneOutputOptionColumnOption struct {
	Name       string  `json:"name"`
	FieldCode  string  `json:"field_code"`
	Type       string  `json:"type"`
	Timezone   *string `json:"timezone"`
	SortColumn *string `json:"sort_column"`
}
