package filter

type FilterColumn struct {
	Name                     string             `json:"name"`
	Src                      string             `json:"src"`
	Type                     string             `json:"type"`
	Default                  *string            `json:"default"`
	HasParser                bool               `json:"has_parser"`
	Format                   *string            `json:"format"`
	JSONExpandEnabled        bool               `json:"json_expand_enabled"`
	JSONExpandKeepBaseColumn bool               `json:"json_expand_keep_base_column"`
	JSONExpandColumns        []JSONExpandColumn `json:"json_expand_columns"`
}

type JSONExpandColumn struct {
	Name     string  `json:"name"`
	JSONPath string  `json:"json_path"`
	Type     string  `json:"type"`
	Format   *string `json:"format"`
	Timezone *string `json:"timezone"`
}
