package output_options

import "terraform-provider-trocco/internal/client/parameter"

type KintoneOutputOptionInput struct {
	KintoneConnectionID              int64                                                               `json:"kintone_connection_id"`
	AppID                            string                                                              `json:"app_id"`
	GuestSpaceID                     *string                                                             `json:"guest_space_id,omitempty"`
	Mode                             string                                                              `json:"mode"`
	UpdateKey                        *string                                                             `json:"update_key,omitempty"`
	IgnoreNulls                      bool                                                                `json:"ignore_nulls"`
	ReduceKey                        *string                                                             `json:"reduce_key,omitempty"`
	ChunkSize                        int64                                                               `json:"chunk_size"`
	KintoneOutputOptionColumnOptions *parameter.NullableObjectList[KintoneOutputOptionColumnOptionInput] `json:"kintone_output_option_column_options,omitempty"`
}

type UpdateKintoneOutputOptionInput struct {
	KintoneConnectionID              *int64                                                              `json:"kintone_connection_id,omitempty"`
	AppID                            *string                                                             `json:"app_id,omitempty"`
	GuestSpaceID                     *string                                                             `json:"guest_space_id,omitempty"`
	Mode                             *string                                                             `json:"mode,omitempty"`
	UpdateKey                        *string                                                             `json:"update_key,omitempty"`
	IgnoreNulls                      *bool                                                               `json:"ignore_nulls,omitempty"`
	ReduceKey                        *string                                                             `json:"reduce_key,omitempty"`
	ChunkSize                        *int64                                                              `json:"chunk_size,omitempty"`
	KintoneOutputOptionColumnOptions *parameter.NullableObjectList[KintoneOutputOptionColumnOptionInput] `json:"kintone_output_option_column_options,omitempty"`
}

type KintoneOutputOptionColumnOptionInput struct {
	Name       string  `json:"name"`
	FieldCode  string  `json:"field_code"`
	Type       string  `json:"type"`
	Timezone   *string `json:"timezone,omitempty"`
	SortColumn *string `json:"sort_column,omitempty"`
}
