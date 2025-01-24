package filter

type FilterGsubInput struct {
	ColumnName string `json:"column_name"`
	Pattern    string `json:"pattern"`
	To         string `json:"to"`
}
