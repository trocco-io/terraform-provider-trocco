package filter

type FilterColumnInput struct {
	Name                     string                  `json:"name"`
	Src                      string                  `json:"src"`
	Type                     string                  `json:"type"`
	Default                  *string                 `json:"default,omitempty"`
	Format                   *string                 `json:"format"`
	JSONExpandEnabled        bool                    `json:"json_expand_enabled"`
	JSONExpandKeepBaseColumn bool                    `json:"json_expand_keep_base_column"`
	JSONExpandColumns        []JSONExpandColumnInput `json:"json_expand_columns,omitempty"`
}

type JSONExpandColumnInput struct {
	Name     string  `json:"name"`
	JSONPath string  `json:"json_path"`
	Type     string  `json:"type"`
	Format   *string `json:"format,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}
