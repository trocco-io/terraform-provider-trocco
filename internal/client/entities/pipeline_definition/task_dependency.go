package workflow

type TaskDependency struct {
	Source      int64 `json:"source"`
	Destination int64 `json:"destination"`
}
