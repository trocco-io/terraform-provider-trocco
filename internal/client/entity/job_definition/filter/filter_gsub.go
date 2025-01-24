package filter

type FilterGsub struct {
	ColumnName string `json:"column_name"`
	Pattern    string `json:"pattern"`
	To         string `json:"to"`
}
